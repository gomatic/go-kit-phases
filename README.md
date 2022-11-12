# Introduction

> “Namespaces are one honking great idea–let’s do more of those!”
> - [The Zen of Python](https://www.python.org/dev/peps/pep-0020/), Item 19

> A good package starts with a good name
> - [The Zen of ~Python~ Go](https://dave.cheney.net/2020/02/23/the-zen-of-go#_a_good_package_starts_with_a_good_name)

This project began as just an attempt to refresh my understanding of [go-kit](https://gokit.io/) while thinking through
the organization and how to iteratively add features to a go-kit project. It evolved into a specialized layout designed
for **iterative development** of transports and endpoints, **isolation of implementation details**, and to be **a clear
reflection of protobuf** service definitions to make navigating and implementing the code more obvious and
straightforward based mostly on the folder hierarchy.

A primary design principle of this project is centered around exposing the specifics of the service, its methods, and
its types while retaining the great concepts and separation of concerns introduced by [go-kit](https://gokit.io/).

Additionally, within the package layout, import-aliases are utilized to standardize packages, making the code for the
packages highly similar even though the types being referenced are incompatible. This technique was chosen over using
interface types or reflection to keep the code more easily navigable and explicit.

```go
import (
  in "github.com/gomatic/go-kit-phases/internal/api/moody/transform/feeling"
  out "github.com/gomatic/go-kit-phases/internal/api/moody/transform/overall"
)

func Request(_ context.Context, q interface{}) (interface{}, error)  { return in.To(q) }
func Response(_ context.Context, p interface{}) (interface{}, error) { return out.To(p) }
```

This design allows each folder to be highly standardized and allowing patterns to standout, making this layout suitable
for code generation based entirely on the protobuf and the transport requirements of the service. And allowing for
the `endpoint` package to be almost the only code requiring domain-specific customizations. That is, a generation
command can take the protobuf and transport type(s) as input and fully generate everything except the implementations of
the endpoints.

# Getting started

Each branch documentation links to comparisons with the prior branch and explains the changes made.

[Start here](docs/) and follow along using the branches.
