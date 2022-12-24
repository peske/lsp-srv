package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/peske/x-tools-internal/_copy_tool/utils"
)

var lspPackages = []string{
	"lsppos",
	"protocol",
	"safetoken",
}

var rootPackages = []string{
	"span",
}

var lspReplace = []string{
	"cmd",
	"lsprpc",
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("invalid arguments")
	}
	var err error
	if os.Args[1] == "-r" {
		err = replaceImports(".")
	} else {
		err = copyPackages(os.Args[1])
	}
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Finished")
	}
}

func copyPackages(src string) error {
	utils.EnsureDir(src)
	for _, p := range lspPackages {
		utils.EnsureDir(filepath.Join(src, "lsp", p))
		utils.EnsureDir(filepath.Join("lsp", p))
	}
	for _, p := range rootPackages {
		utils.EnsureDir(filepath.Join(src, p))
		utils.EnsureDir(p)
	}
	for _, p := range lspPackages {
		dst := filepath.Join("lsp", p)
		if err := os.RemoveAll(dst); err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := utils.CopyDir(filepath.Join(src, "lsp", p), dst, replace); err != nil {
			return err
		}
	}
	for _, p := range rootPackages {
		if err := os.RemoveAll(p); err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := utils.CopyDir(filepath.Join(src, p), p, replace); err != nil {
			return err
		}
	}
	return nil
}

func replaceImports(dir string) error {
	var err error
	var fds []os.DirEntry
	if fds, err = os.ReadDir(dir); err != nil {
		return err
	}
	for _, fd := range fds {
		if fd.IsDir() {
			if err = replaceImports(fd.Name()); err != nil {
				return err
			}
			continue
		}
		c, err := os.ReadFile(fd.Name())
		if err != nil {
			return err
		}
		str := replace(string(c))
		if err = os.WriteFile(fd.Name(), []byte(str), 0700); err != nil {
			return err
		}
	}
	return nil
}

func replace(content string) string {
	for _, p := range utils.Packages {
		content = strings.Replace(content,
			"golang.org/x/tools/internal/"+p,
			"github.com/peske/x-tools-internal/"+p, -1)
	}
	for _, p := range lspPackages {
		content = strings.Replace(content,
			"golang.org/x/tools/gopls/internal/lsp/"+p,
			"github.com/peske/lsp-srv/lsp/"+p, -1)
	}
	for _, p := range lspReplace {
		content = strings.Replace(content,
			"golang.org/x/tools/gopls/internal/lsp/"+p,
			"github.com/peske/lsp-srv/lsp/"+p, -1)
	}
	for _, p := range rootPackages {
		content = strings.Replace(content,
			"golang.org/x/tools/gopls/internal/"+p,
			"github.com/peske/lsp-srv/"+p, -1)
	}
	content = strings.Replace(content,
		"golang.org/x/tools/gopls/internal/lsp/...",
		"github.com/peske/lsp-srv/lsp/...", -1)
	return content
}
