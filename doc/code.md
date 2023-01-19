# About the code

> **Note:** The instructions on these page are about building this module, **not** how to use it.

As already mentioned, the majority of this module code is copied from the source
[`golang.org/x/tools/gopls`](https://github.com/golang/tools/tree/master/gopls) module. Some packages are copied
completely, with only `import` statements replaced appropriately, while from others we've introduced some additional
changes or copied just a small part of them.

## Completely copied packages

The following packages are completely copied from the original `golang.org/x/tools/gopls/internal` directory:

- [`lsp/protocol`](../lsp/protocol)
- [`lsp/safetoken`](../lsp/safetoken). Note that a test file from this package will add some unnecessary dependencies,
  so we should remove these before committing the code. These are all starting with `golang.org/x/tools` (this module
  does not require any `goolang.org/x/tools` packages).
- [`span`](../span)

After that we have replaced all `golang.org/x/tools/gopls/internal/...` occurrences with the corresponding
`github.com/peske/lsp-srv/...`.

The previous steps can be done by using a _copy tool_. The tool automatically copies packages from
`golang.org/x/tools/gopls/internal` into this repository. To use it, we first need to build it:

```bash
go build .
```

This will create an executable file `lsp-srv` in the working directory. The next step is to execute it to actually copy
the code:

```bash
./lsp-srv /path/to/golang.org/x/tools
```

The CLI argument is the local path of the original `golang.org/x/tools` module. The argument is optional, and if omitted
it will default to `$GOHOME/src/golang.org/x/tools`.

If the source package is located at its default location (`$GOHOME/src/golang.org/x/tools`), instead of the previous two
steps we can simply do `go generate`. It will do the same thing because [`main.go`](../main.go) file contains the
following lines:

```go
//go:generate go build .
//go:generate ./lsp-srv
```

The copy tool has one additional feature: replacing `golang.org/x/tools..` imports with the corresponding
`github.com/peske/..` imports. The feature is used in the following way:

```bash
./lsp-srv -r
```

When executed with `-r` flag, the tool will not copy any code, but will simply perform the replacements in all `*.go`
and `*.ts` files in all the directories from this repository, except for _hidden_ directories (the directories whose
name starts with a dot), to avoid messing up with the `.git` directory content. The feature is also added to
`go generate`:

```go
//go:generate ./lsp-srv -r
```

In other words, simple `go generate` execution will perform both copying the code from the source, and the replacement
in all the files, copied and custom.

## Partially copied packages

We've copied `lsp/helper` package, but there we've introduced two changes:

- Package name `main` is changed to `helper`;
- We've added a custom file `generator.go` there.

Additionally, we've cherry-picked some small parts of packages `lsp/cmd` and `lsp/lsprpc`. There's no a tool or simple
instructions how to copy these.
