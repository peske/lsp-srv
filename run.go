package lsp_srv

import (
	"context"

	"github.com/peske/lsp-srv/lsp/cmd"
	"github.com/peske/lsp-srv/lsp/protocol"
)

func Run(serverFactory func(protocol.ClientCloser, context.Context, func()) protocol.Server, config *Config) error {
	s := cmd.Serve{ServerFactory: serverFactory}
	if config != nil {
		s.Port = config.Port
		s.Address = config.Address
		s.IdleTimeout = config.IdleTimeout
	}
	return s.Run()
}
