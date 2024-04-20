# protoc-gen-go-pagesize

will generate constant values for the default and max for the page_size field within any defined messages.

https://google.aip.dev/158

## install

```
go install github.com/lcmaguire/protoc-gen-go-pagesize@latest
```

## What can you do with setters that you cant do idiomatically?

by having setters you can then create packages within your organization based upon any naming structures you have for your proto specs.

for example if you are following goolges aip's in particular 

- [AIP-132](https://google.aip.dev/132) for list endpoints 
- [AIP-158](https://google.aip.dev/158) for pagination.

you can make a package for pagination based upon the func `GetNextPageToken()` in list responses & `SetPageToken(in string)` generated for requests

see [grpcpagination](https://github.com/lcmaguire/grpcpagination) for an example
