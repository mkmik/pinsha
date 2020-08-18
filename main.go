// Pinsha is a simple tool that adds the image digest to an OCI image reference.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <image>\n", os.Args[0])
}

func mainE() error {
	if flag.NArg() == 0 {
		usage()
		return nil
	}

	ref := flag.Arg(0)
	if strings.Contains(ref, "@sha256:") {
		fmt.Println(ref)
		return nil
	}

	d, err := crane.Digest(ref)
	if err != nil {
		return err
	}
	fmt.Printf("%s@%s\n", ref, d)
	return nil
}

func main() {
	flag.Parse()

	if err := mainE(); err != nil {
		log.Fatal(err)
	}
}
