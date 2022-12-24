# Usage

The first step is to create an instance of a type that implements
[`protocol.Server` interface](../protocol/tsserver.go#L8). You can create such a type in any way you like, but probably
the easiest and recommended way is to [use generator](./generator.md).

The next step is to create [`cmd.Serve` instance](../cmd/serve.go#L20), and set its `ServerFactory` field to a function
that returns `protocol.Server` instance. The other fields are representing some configuration values, and for `stdio`
communication they should be left at their default values.

Finally, call [`Run` method](../cmd/serve.go#L36) of the `cmd.Serve` instance created in the previous step.
