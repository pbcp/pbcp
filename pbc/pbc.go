package main

import (
	"github.com/pbcp/pbcp/common"
	"github.com/pbcp/pbcp/common/config"

	"errors"
	"net/http"
	"os"
	_ "io"
	"strings"
)

const MIMEType = "application/octet-stream"

func main() {
	id := config.Id()

	_ = strings.Join([]string{config.Server(), "board", id}, "/")

	res, err := http.Post(
		"http://httpbin.org/post",
		MIMEType,
		os.Stdin,
	)
	common.Fatal(err)

	// io.Copy(os.Stderr, res.Body)

	switch res.StatusCode {
	case 400:
		usage()
	case 500:
		common.Fatal(errors.New("Internal server error!"))
	}
}

func usage() {
	common.Fatal(errors.New(
		"Usage: pbc\n  Copies data to clipboard from stdin.\n  Check server for maximum file size.",
	))
}
