# What?

A module provides a wireframe for implementing [Language Server protocol](https://langserver.org/) server in Go. Note
that it does not provide any actual end-user functionality, just a general wireframe: message types, JSON RPC 2.0
serialization, communication etc. In other words:

> By using this package you can develop LSP server in Go for any purpose / programming language, without having to take
  care about the underlying implementation of JSON RPC 2.0, connections, etc. You can focus on the actual features you
  need.

Our main intention is to create:

> **The simplest and cleanest way to start developing your own LSP server implementations.**

Usage of the module is explained in [the documentation](./doc/README.md).

> **Note:** You should also check [`github.com/peske/lsp-srv-ex`](https://github.com/peske/lsp-srv-ex) module which
> extends this one, and provides some useful additional features.

# How?

This module is created by copying some packages from
[`golang.org/x/tools/gopls`](https://github.com/golang/tools/tree/master/gopls) module, and adding some custom packages.
More details in `./doc/code.md`(./doc/code.md).

# Why?

Probably the best implementation of LSP server in Go is in the mentioned
[`golang.org/x/tools/gopls`](https://github.com/golang/tools/tree/master/gopls) module, but it suffers from two
problems:

- LSP related code is _hidden_ in an `internal` parts of the module (`golang.org/x/tools/gopls/internal`), so it cannot
  be accessed / used from an external code.
- General LSP-related features are not separated from Go-specific implementation, thus making it very hard to use for
  building an LSP server for some other language.

In this module we're stripping out only general wireframe parts, and making it accessible for the external code.

## Alternatives?

There are some notable alternatives:

- https://github.com/sourcegraph/go-lsp
- https://github.com/saibing/bingo
- https://github.com/tliron/glsp

But they are suffering some other drawbacks:

- Often not actively developed, not supporting the latest LSP versions.
- Some being aborted and archived in favor of mentioned `gopls`.
- Usually lacking a comprehensive documentation.
- In some cases the code quality and design decisions are not up to our standards.

> **Note:** We apologize to the maintainers of the mentioned alternatives for our criticism. We **do respect** your
> work, and everything said here is **just our opinion** that explains our reasoning for starting this project.

# License?

The same license as the original one - [BSD-3-Clause license](./LICENSE). Although almost all the code is created by the
authors of the original module (`The Go Authors`), we've changed copyright here to `Fat Dragon and authors` not to get
the credits, but to protect the original authors of any responsibility if there are any problems in the code that we've
changed. All credits should go to the authors of the original module.

# Version?

Current `main` branch is based on the original repository commit
[3e6f71b](https://github.com/golang/tools/commit/3e6f71bba4359aeb7a301d361ee3cf95e8799599) from January 17, 2023.
