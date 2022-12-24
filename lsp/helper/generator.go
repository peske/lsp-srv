package helper

import (
	"flag"
	"log"
	"path"
	"path/filepath"
	"runtime"
)

func Generate() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	if *def == "" {
		*def = tsserver()
	}
	if *typ == "" {
		*typ = "Server"
	}
	if *use == "" || *out == "" {
		flag.PrintDefaults()
		return
	}
	// read the type definition and see what methods we're looking for
	doTypes()

	// parse the package and see which methods are defined
	doUses()

	output()
}

func tsserver() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller call failed")
	}
	return path.Join(filepath.Dir(file), "../protocol/tsserver.go")
}
