package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/peske/lsp-srv/lsp/lsprpc"
	"github.com/peske/lsp-srv/lsp/protocol"
	"github.com/peske/x-tools-internal/fakenet"
	"github.com/peske/x-tools-internal/jsonrpc2"
)

// Serve is a struct that exposes the configurable parts of the LSP server as
// flags, in the right form for tool.Main to consume.
type Serve struct {
	// Port on which to run the server.
	Port int

	// Address on which to listen for remote connections. If prefixed by 'unix;', the subsequent address is assumed to
	// be a unix domain socket. Otherwise, TCP is used.
	Address string

	// IdleTimeout - shut down the server when there are no connected clients for this duration.
	IdleTimeout time.Duration

	ServerFactory func(protocol.ClientCloser) protocol.Server
}

// Run configures a server based on the flags, and then runs it.
// It blocks until the server shuts down.
func (s *Serve) Run(ctx context.Context) error {
	ss := lsprpc.NewStreamServer(s.ServerFactory)

	var network, addr string
	if s.Address != "" {
		network, addr = lsprpc.ParseAddr(s.Address)
	}
	if s.Port != 0 {
		network = "tcp"
		addr = fmt.Sprintf(":%v", s.Port)
	}
	if addr != "" {
		log.Printf("LSP daemon: listening on %s network, address %s...", network, addr)
		defer log.Printf("LSP daemon: exiting")
		return jsonrpc2.ListenAndServe(ctx, network, addr, ss, s.IdleTimeout)
	}
	stream := jsonrpc2.NewHeaderStream(fakenet.NewConn("stdio", os.Stdin, os.Stdout))
	conn := jsonrpc2.NewConn(stream)
	err := ss.ServeStream(ctx, conn)
	if errors.Is(err, io.EOF) {
		return nil
	}
	return err
}
