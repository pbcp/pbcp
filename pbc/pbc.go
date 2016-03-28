package main

import (
	"github.com/pbcp/pbcp/common"
	"github.com/pbcp/pbcp/common/config"

	"errors"
	"net/http"
	"os"
	"strings"
)

const MIMEType = "application/octet-stream"

func main() {
	id := config.Id()

	res, err := http.Post(
		strings.Join([]string{config.Server(), "board", id}, "/"),
		MIMEType,
		os.Stdin,
	)
	common.Fatal(err)

	switch res.StatusCode {
	case 400:
		usage()
	case 413:
		common.Fatal(errors.New("Maximum file size exceeded!"))
	case 500:
		common.Fatal(errors.New("Internal server error!"))
	}
}

func usage() {
	common.Fatal(errors.New(
		"Usage: pbc\n  Copies data to clipboard from stdin.\n  Check server for maximum file size.",
	))
}
