// Code generated by protoc-gen-go-pagesize. DO NOT EDIT.
// source: example/example.proto
package example

const (

	// ListExampleDefaultPageSize if the user does not specify page_size (or specifies 0)
	//
	// 25 is the value that will be used by the API as specified in https://google.aip.dev/158.
	ListExampleDefaultPageSize int32 = 25

	// ListExampleMaxPageSize represents the maximum page_size for a ListExample
	//
	// if page_size is greater than 1000 it will be coerced down to 1000 as specified in https://google.aip.dev/158.
	ListExampleMaxPageSize int32 = 1000
)