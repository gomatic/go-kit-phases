This branch
[adds](https://github.com/gomatic/go-kit-phases/compare/02-service-basic-implementation...03-service-basic-tests)
[basic test stubs](../internal/api/moody/service/self/service_test.go)
for the
[gRPC service](../api/moody/self.proto)
.

Note the importance of writing tests early for this business logic. This package and the gRPC service definitions are
central to the entire service so is probably where most design and development should be focused. Much of the rest of
this project is devoted to supporting capabilities around the core business logic.

[previous branch](../../02-service-basic-implementation/docs#readme)
| [next branch](../../04-endpoints-stubs/docs#readme)
