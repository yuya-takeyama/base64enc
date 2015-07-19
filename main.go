package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/yuya-takeyama/argf"
)

const AppName = "base64enc"

type Options struct {
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = AppName
	parser.Usage = "[OPTIONS] FILES..."

	args, err := parser.Parse()
	if err != nil {
		fmt.Print(err)
		return
	}

	r, err := argf.From(args)
	if err != nil {
		panic(err)
	}

	err = base64enc(r, os.Stdout, os.Stderr, opts)
	if err != nil {
		panic(err)
	}
}

func base64enc(r io.Reader, stdout io.Writer, stderr io.Writer, opts Options) error {
	if opts.ShowVersion {
		io.WriteString(stdout, fmt.Sprintf("%s v%s, build %s\n", AppName, Version, GitCommit))
		return nil
	}

	encoder := base64.NewEncoder(base64.StdEncoding, stdout)
	_, err := io.Copy(encoder, r)
	if err != nil {
		return err
	}

	err = encoder.Close()
	if err != nil {
		return err
	}

	io.WriteString(stdout, "\n")
	return nil
}
