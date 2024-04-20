# protoc-gen-go-pagesize

will generate constant values for the `default` and `max` for the page_size field within any defined messages.

https://google.aip.dev/158

## install

```
go install github.com/lcmaguire/protoc-gen-go-pagesize@latest
```

## Usage

```proto

// import the field descriptor.
import "pagesize/pagesize.proto";

message ListExample {
	// add in the field descriptor
    int32 page_size = 1 [(pagesize.field).default_page_size = 25, (pagesize.field).max_page_size = 1000];
}

```

output 

```go
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
```