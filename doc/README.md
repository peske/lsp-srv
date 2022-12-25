# Usage

## Create a type

The first step is to create a type that will handle communication with the client, and a factory function that creates
an instance of this type, and returns a pointer to the instance. The function should accept three parameters:

- `protocol.ClientCloser` that represents the client;
- `context.Context` is the context in which the server runs;
- `func()` is the context cancellation function that you can call to cancel the context.

Here's a minimal implementation:

```go
package lsp

import (
	"context"
	
	"github.com/peske/lsp-srv/lsp/protocol"
)

// Server is the type mentioned in the instructions above.
type Server struct {
	client protocol.ClientCloser
	ctx    context.Context
	cancel func()
}

// NewServer is the factory function mentioned in the instructions above.
func NewServer(client protocol.ClientCloser, ctx context.Context, cancel func()) *Server {
	return &Server{
		client: client,
		ctx:    ctx,
		cancel: cancel,
    }
}
```

## Implement `protocol.Server` interface

The type that you've created in the previous step must implement `protocol.Server` interface. You can create all the
methods manually, but it will be much faster if you [use generator](./generator.md).

## (Optionally) Create Config

If you want your server to use `stdio` for the communication with the client, you don't have to configure anything, and
you can skip creating [`lsp_srv.Config`](../config.go) instance. But if you want communication to go via TCP for
example, you will have to create it.

## Start the server

Here's an example file that starts the server:

```go
package main

import (
	"log"
	
	lsp_srv "github.com/peske/lsp-srv"

	"github.com/yourgh/yourmodule/lsp"
)

func main() {
	var cfg *lsp_srv.Config
	// Optionally create the config

	// Here we assume that your factory function resides in `lsp` package, thus `lsp.NewServer`.
	if err := lsp_srv.Run(lsp.NewServer, cfg); err != nil {
		log.Fatal(err)
	}
}
```

## Examples

You can find a few usage examples in https://github.com/peske/lsp-example repository.
