package main

import (
	"bytes"
	"flag"
	"text/template"

	"github.com/lcmaguire/protoc-gen-go-pagesize/pagesize"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// this enables optional fields to be supported.
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		// for all files, generate setters for all fields for all messages.
		for _, file := range gen.Files {
			if !file.Generate || len(file.Messages) == 0 {
				continue
			}

			outPut := map[string]*pagesize.PageSizeValues{}
			for _, message := range file.Messages {
				for _, field := range message.Fields {
					if pageSize := findPageSize(field.Desc); pageSize != nil {
						messageName := message.Desc.Name()
						outPut[string(messageName)] = pageSize
					}
				}
			}

			if len(outPut) == 0 {
				continue
			}

			newFile := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".pb.pagesize.go", ".")
			newFile.P("// Code generated by protoc-gen-go-pagesize. DO NOT EDIT.")
			newFile.P("// source: ", file.GeneratedFilenamePrefix, ".proto")
			newFile.P("package " + file.GoPackageName)

			for messageName, pageSizeInfo := range outPut {
				p := PageSizeTemplate{
					MessageName:     messageName,
					DefaultPageSize: pageSizeInfo.DefaultPageSize,
					MaxPageSize:     pageSizeInfo.MaxPageSize,
				}
				newFile.P(p.executeTemplate())
			}
		}

		return nil
	})
}

func findPageSize(desc protoreflect.FieldDescriptor) *pagesize.PageSizeValues {
	num := pagesize.E_Field.TypeDescriptor().Number()

	var msg proto.Message
	proto.RangeExtensions(desc.Options(), func(typ protoreflect.ExtensionType, i interface{}) bool {
		if num != typ.TypeDescriptor().Number() {
			return true
		}
		msg, _ = i.(proto.Message)
		return false
	})

	if msg == nil {
		return nil
	}

	pageSize, _ := msg.(*pagesize.PageSizeValues)
	return pageSize
}

// PageSizeTemplate template used to generate
type PageSizeTemplate struct {
	MessageName     string
	DefaultPageSize int32
	MaxPageSize     int32
}

const tplate = `
	const (
	
		// {{.MessageName}}DefaultPageSize if the user does not specify page_size (or specifies 0)
		// 
		// {{.DefaultPageSize}} is the value that will be used by the API as specified in https://google.aip.dev/158.
		{{.MessageName}}DefaultPageSize int32 = {{.DefaultPageSize}}
	
		// {{.MessageName}}MaxPageSize represents the maximum page_size for a {{.MessageName}}
		//
		// if page_size is greater than {{.MaxPageSize}} it will be coerced down to {{.MaxPageSize}} as specified in https://google.aip.dev/158.
		{{.MessageName}}MaxPageSize int32 = {{.MaxPageSize}}
	)
`

func (p PageSizeTemplate) executeTemplate() string {
	templ, err := template.New("").Parse(tplate)
	if err != nil {
		panic(err)
	}
	buffy := bytes.NewBuffer([]byte{})
	if err := templ.Execute(buffy, p); err != nil {
		panic(err)
	}
	return buffy.String()
}
