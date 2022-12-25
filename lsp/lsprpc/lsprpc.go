// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lsprpc implements a jsonrpc2.StreamServer that may be used to
// serve the LSP on a jsonrpc2 channel.
package lsprpc

import (
	"context"
	"strings"

	"github.com/peske/x-tools-internal/event"
	"github.com/peske/x-tools-internal/jsonrpc2"

	"github.com/peske/lsp-srv/lsp/protocol"
)

// The StreamServer type is a jsonrpc2.StreamServer that handles incoming
// streams as a new LSP session, using a shared cache.
type StreamServer struct {
	Context               context.Context
	ContextCancellationFn func() // Context cancellation function
	serverFactory         func(protocol.ClientCloser, context.Context, func()) protocol.Server
}

// NewStreamServer creates a StreamServer using the shared cache. If
// withTelemetry is true, each session is instrumented with telemetry that
// records RPC statistics.
func NewStreamServer(serverFactory func(protocol.ClientCloser, context.Context, func()) protocol.Server) *StreamServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &StreamServer{
		Context:               ctx,
		ContextCancellationFn: cancel,
		serverFactory:         serverFactory,
	}
}

// ServeStream implements the jsonrpc2.StreamServer interface, by handling
// incoming streams using a new lsp server.
func (s *StreamServer) ServeStream(ctx context.Context, conn jsonrpc2.Conn) error {
	client := protocol.ClientDispatcher(conn)
	server := s.serverFactory(client, s.Context, s.ContextCancellationFn)
	// Clients may or may not send a shutdown message. Make sure the server is
	// shut down.
	// TODO(rFindley): this shutdown should perhaps be on a disconnected context.
	defer func() {
		if err := server.Shutdown(ctx); err != nil {
			event.Error(ctx, "error shutting down", err)
		}
	}()
	ctx = protocol.WithClient(ctx, client)
	conn.Go(ctx, protocol.Handlers(protocol.ServerHandler(server, jsonrpc2.MethodNotFound)))
	<-conn.Done()
	return conn.Err()
}

// ParseAddr parses the address of a gopls remote.
// TODO(rFindley): further document this syntax, and allow URI-style remote
// addresses such as "auto://...".
func ParseAddr(listen string) (network string, address string) {
	// Allow passing just -remote=auto, as a shorthand for using automatic remote
	// resolution.
	if listen == AutoNetwork {
		return AutoNetwork, ""
	}
	if parts := strings.SplitN(listen, ";", 2); len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "tcp", listen
}
