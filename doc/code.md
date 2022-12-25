# About the code

The following packages are simply copied from the original `golang.org/x/tools/gopls/internal` directory:

- [`lsp/lsppos`](../lsp/lsppos)
- [`lsp/protocol`](../lsp/protocol)
- [`lsp/safetoken`](../lsp/safetoken). Note that a test file from this package will add some unnecessary dependencies,
  so we should remove these before committing the code. These are all starting with `golang.org/x/tools` (this module
  does not require any `goolang.org/x/tools` packages).
- [`span`](../span)

After that we have replaced all `golang.org/x/tools/gopls/internal/...` occurrences with the corresponding
`github.com/peske/lsp-srv/...`.

The previous steps can be done by using a [copy tool](../_copy_tool). The tool automatically copies packages from
`golang.org/x/tools/gopls/internal` into this repository. Usage:

- `cd` into `./_copy_tool` directory, and build the tool by executing `go build -o ../cptool`. Note that the executable
  is stored in the root directory of this repository.
- `cd` into the root directory of this repository and execute: `./cptool /path/to/golang.org/x/tools/gopls/internal`
  (change the source path appropriately).

Additionally, we've copied `lsp/helper` package, but there we've introduced two changes:

- Package name `main` is changed to `helper`;
- We've added a custom file `generator.go` there.

We've also used some other packages from the original repo: `lsp/cmd` and `lsp/lsprpc`, but for these we were
chery-picking only the parts that we need. There's no a tool or simple instructions how to copy these.

Finally, we've added some custom code in [`helper`](../helper) package. This is the place where our creativity shines :)
