package lsp_srv

import "time"

type Config struct {
	// Port on which to run the server.
	Port int

	// Address on which to listen for remote connections. If prefixed by 'unix;', the subsequent address is assumed to
	// be a unix domain socket. Otherwise, TCP is used.
	Address string

	// IdleTimeout - shut down the server when there are no connected clients for this duration.
	IdleTimeout time.Duration
}
