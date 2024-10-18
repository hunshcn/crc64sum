package main

import (
	"fmt"
	"hash/crc64"
	"io"
	"os"
)

const name = "crc64sum"

func sum(name string, reader io.Reader) error {
	c := crc64.New(crc64.MakeTable(crc64.ECMA))
	_, err := io.Copy(c, reader)
	if err != nil {
		return err
	}
	fmt.Printf("%x\t%s\n", c.Sum64(), name)
	return nil
}

func sumFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return sum(name, f)
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if err := sumFile(arg); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
				os.Exit(1)
			}
		}
	} else {
		if err := sum("-", os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
			os.Exit(1)
		}
	}
}
