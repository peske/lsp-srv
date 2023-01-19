package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/peske/x-tools-internal/_copy_tool"
)

var lspPackages = []string{
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
	// Get source path
	var src string
	var err error
	switch len(os.Args) {
	case 1:
		h := os.Getenv("GOHOME")
		if h == "" {
			log.Fatalln("Source directory not specified.")
		}
		src = filepath.Join(h, "src", "golang.org", "x", "tools")
	case 2:
		if os.Args[1] == "-r" {
			if err = replaceImports("."); err == nil {
				log.Println("Replace finished")
			} else {
				log.Fatalln(err)
			}
			return
		}
		if src, err = filepath.Abs(os.Args[1]); err != nil {
			log.Fatalf("Invalid source path '%s'", os.Args[1])
		}
	default:
		log.Fatalln("Invalid arguments.")
	}

	// Check source validity
	if err = checkSourceValidity(src); err != nil {
		log.Fatalln(err)
	}
	src = filepath.Join(src, "gopls", "internal")

	// Delete and copy new
	dst, _ := filepath.Abs(".") // won't error ever
	for _, p := range rootPackages {
		d := filepath.Join(dst, p)
		if err = os.RemoveAll(d); err != nil && !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		if err = _copy_tool.CopyDir(filepath.Join(src, p), d, replace); err != nil {
			log.Fatalln(err)
		}
	}
	src = filepath.Join(src, "lsp")
	dst = filepath.Join(dst, "lsp")
	for _, p := range lspPackages {
		d := filepath.Join(dst, p)
		if err = os.RemoveAll(d); err != nil && !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		if err = _copy_tool.CopyDir(filepath.Join(src, p), d, replace); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Copy finished")
}

func checkSourceValidity(src string) error {
	if err := _copy_tool.CheckDirectoryExists(src); err != nil {
		return err
	}
	src = filepath.Join(src, "gopls", "internal")
	if err := _copy_tool.CheckDirectoryExists(src); err != nil {
		return err
	}
	for _, p := range rootPackages {
		if err := _copy_tool.CheckDirectoryExists(filepath.Join(src, p)); err != nil {
			return err
		}
	}
	src = filepath.Join(src, "lsp")
	if err := _copy_tool.CheckDirectoryExists(src); err != nil {
		return err
	}
	for _, p := range lspPackages {
		if err := _copy_tool.CheckDirectoryExists(filepath.Join(src, p)); err != nil {
			return err
		}
	}
	return nil
}

func replaceImports(dir string) error {
	fds, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, fd := range fds {
		if strings.HasPrefix(fd.Name(), ".") {
			// We don't want to replace in hidden files/directories
			continue
		}
		f := filepath.Join(dir, fd.Name())
		if fd.IsDir() {
			if err = replaceImports(f); err != nil {
				return err
			}
			continue
		}
		if !strings.HasSuffix(strings.ToLower(f), ".go") && !strings.HasSuffix(strings.ToLower(f), ".ts") {
			// We want to replace only in *.go and *.ts files
			continue
		}
		var content []byte
		var fi os.FileInfo
		if content, err = os.ReadFile(f); err != nil {
			return err
		}
		if fi, err = os.Stat(f); err != nil {
			return err
		}
		content = _copy_tool.Replace(f, content)
		content = replace(f, content)
		if err = os.WriteFile(f, content, fi.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func replace(fp string, content []byte) []byte {
	if strings.HasSuffix(strings.ToLower(fp), ".md") {
		// We don't want to replace in Markdown files
		return content
	}
	str := string(content)
	for _, p := range lspPackages {
		str = strings.ReplaceAll(str,
			"golang.org/x/tools/gopls/internal/lsp/"+p,
			"github.com/peske/lsp-srv/lsp/"+p)
	}
	for _, p := range lspReplace {
		str = strings.ReplaceAll(str,
			"golang.org/x/tools/gopls/internal/lsp/"+p,
			"github.com/peske/lsp-srv/lsp/"+p)
	}
	for _, p := range rootPackages {
		str = strings.ReplaceAll(str,
			"golang.org/x/tools/gopls/internal/"+p,
			"github.com/peske/lsp-srv/"+p)
	}
	str = strings.ReplaceAll(str,
		"github.com/peske/lsp-srv/lsp/...",
		"github.com/peske/lsp-srv/lsp/...")
	return []byte(str)
}

//go:generate go build .
//go:generate ./lsp-srv
//go:generate ./lsp-srv -r
