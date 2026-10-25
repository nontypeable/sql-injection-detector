package main

import (
	"log"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("usage: sast <package>")
	}

	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadAllSyntax}, args...)
	if err != nil {
		log.Fatalf("load packages: %v", err)
	}
	if errCount := packages.PrintErrors(pkgs); errCount > 0 {
		log.Fatalf("package contains %d errors", errCount)
	}

	prog, ssaPkgs := ssautil.Packages(pkgs, ssa.InstantiateGenerics)
	prog.Build()

	for _, pkg := range ssaPkgs {
		if pkg == nil {
			log.Fatal("ssautil: nil package")
		}

		for _, member := range pkg.Members {
			fn, ok := member.(*ssa.Function)
			if !ok {
				continue
			}
			if _, err := fn.WriteTo(os.Stdout); err != nil {
				log.Printf("write %s: %v", fn.Name(), err)
			}
		}
	}
}
