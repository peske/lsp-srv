# Usage

The first step is to create an instance of a type that implements
[`protocol.Server` interface](../lsp/protocol/tsserver.go#L18). You can create such a type in any way you like, but
probably the easiest and recommended way is to [use generator](./generator.md).

The next step is to create [`cmd.Serve` instance](../lsp/cmd/serve.go#L21), and set its `ServerFactory` field to a
function that returns `protocol.Server` instance. The other fields are representing some configuration values, and for
`stdio` communication they should be left at their default values.

Finally, call [`Run` method](../lsp/cmd/serve.go#L37) of the `cmd.Serve` instance created in the previous step.

You can find a few usage examples in https://github.com/peske/lsp-example repository.
