package main

import (
	"bytes"
	"encoding/base32"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	decode = flag.Bool("d", false, "decode")

	codec = base32.StdEncoding.WithPadding(base32.NoPadding)
)

func run(decode bool) error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	if decode {
		buf := make([]byte, codec.DecodedLen(len(b)))
		n, err := codec.Decode(buf, bytes.ToUpper(b))
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(buf[:n])
		return err
	} else {

		s := codec.EncodeToString(b)
		fmt.Print(strings.ToLower(s))
	}
	return nil
}

func main() {
	flag.Parse()
	if err := run(*decode); err != nil {
		log.Fatal(err)
	}
}
