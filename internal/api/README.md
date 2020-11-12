# `api` package

The `api` package includes a folder/package for [each protobuf API](../../api).

Generally, a repository should _probably_ be restricted to the implementation of a single API. But to facilitate rapid
development, for example, for interactions between APIs, it may be valuable to development APIs together. The layout of
this repository helps ensure it will be relatively easy to subsequently split the APIs into their own repositories.

The overall folder hierarchy is:

```text
{{.API}}
├── endpoint
│   ├── endpoint.go
│   └── {{.Service}}
│       ├── set.go
│       ├── {{.Method}}
│       .   └── endpoint.go
│       .
│       └── {{.Method}}
│           └── ...
├── service
│   ├── service.go
│   └── {{.Service}}
│       ├── methods.go
│       └── service.go
├── transform
│   ├── transform.go
│   ├── {{.Message}}
│   │   └── server
│   │       ├── server.go
│   │       ├── {{.Transport}}
│   │       .   └── transform.go
│   │       .
│   │       └── {{.Transport}}
│   │           └── transform.go
│   .
│   .
│   └── {{.Message}}
│       └── ...
└── transport
    ├── transport.go
    ├── {{.Transport}}
    │   ├── instance.go
    │   ├── listen.go
    │   └── {{.Service}}
    .       └── server.go
    .
    └── {{.Transport}}
        └── ...
```
