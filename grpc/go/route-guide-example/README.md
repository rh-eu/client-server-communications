# [route-guide-example](https://github.com/grpc/grpc-go/tree/master/examples/route_guide)

## generate gRPC code

```bash
cd grpc/go/route-guide-example/pkg/
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/routeguide.proto
```
