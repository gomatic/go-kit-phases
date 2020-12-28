This [simply](https://github.com/gomatic/go-kit-phases/compare/06-modified-main...09-merge-servers) brings together all the servers for the final product.
The transport implementations will not collide when merged.
All other aspects of the service are common regardless of the transport (given this relatively simplistic implementation).
Thus, the value of this go-kit technique.
