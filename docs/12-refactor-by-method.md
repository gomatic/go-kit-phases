This refactor is intended to make managing endpoints and transports simpler. It does this by making endpoints somewhat cookie cutter in that they are relatively the same set of methods with all the same interfaces to support the same transports.

It also extracts the transformations from the transport package to make similarly cookie cutter transformations. Although the transformations are transport-specific, it's relatively trivial now to copy/paste transforms from on type to another instead of having lots of type-specific functions.

A nice side-effect of this is that the transport-definitions more clearly represet the request/response types of each service endpoint.

It's left to the implementation to actually handle encoders and decoders since they are service specific. They should probably reside in their own package. Of course, that was/is the transports package, but it feels like data transformations are not specifically the domain of a transport. These data transformations are more about interfacing between packages than they are specific to transports.
