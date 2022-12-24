# What?

The module provides a wireframe for implementing [Language Server protocol](https://langserver.org/) server in Go. Note
that it does not provide any actual end-user functionality, just a general wireframe: message types, JSON RPC 2.0
serialization, communication etc.

The module is based on the existing on the existing
[`golang.org/x/tools/gopls`](https://github.com/golang/tools/tree/master/gopls) module.

# How?

The following packages are simply copied from the original `golang.org/x/tools/gopls/internal` directory:

- [`lsp/lsppos`](./lsp/lsppos)
- [`lsp/protocol`](./lsp/protocol)
- [`lsp/safetoken`](./lsp/safetoken), except for `safetoken_test.go` file which is skipped to avoid adding unnecessary
  dependencies
- [`span`](./span)

After that we have replaced all `golang.org/x/tools/gopls/internal/...` occurrences with the corresponding
`github.com/peske/lsp-srv/...`.

The previous steps can be done by using a [copy tool](./_copy_tool). The tool automatically copies packages from
`golang.org/x/tools/gopls/internal` into this repository. Usage:

- `cd` into `./_copy_tool` directory, and build the tool by executing `go build -o ../cptool`. Note that the executable
  is stored in the root directory of this repository.
- `cd` into the root directory of this repository and execute: `./cptool /path/to/golang.org/x/tools/gopls/internal`
  (change the source path appropriately).

# Why?

The copied packages contain a very nice functionality which isn't accessible since they are `internal` in the original
module.

# License?

The same "BSD-3-Clause license" used by the original repository. Here I've changed _Copyright_ section only because we
have some additional code that may have some errors, and I didn't want anyone to blame the original developers ("The Go
Authors"). But they deserve all the credits for the code we're using here.

# Version?

Current `main` branch is based on the original repository commit
[eb70795](https://github.com/golang/tools/commit/eb70795aaccb8e6c9615c88085ef3414ba04b8c9) from December 17, 2022.
