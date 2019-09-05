package main

import (
	"encoding/base32"
	"flag"
	"io"
	"log"
	"os"
)

var (
	decode = flag.Bool("d", false, "decode")
)

func run(decode bool) error {
	var (
		w io.Writer
		r io.Reader
	)
	codec := base32.StdEncoding
	if decode {
		w = os.Stdout
		r = base32.NewDecoder(codec, os.Stdin)

	} else {
		e := base32.NewEncoder(codec, os.Stdout)
		defer e.Close()
		w = e
		r = os.Stdin
	}

	_, err := io.Copy(w, r)
	return err
}

func main() {
	flag.Parse()
	if err := run(*decode); err != nil {
		log.Fatal(err)
	}
}
