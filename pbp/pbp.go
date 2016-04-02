package main

import (
	"github.com/pbcp/pbcp/common"
	"github.com/pbcp/pbcp/common/config"

	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	id := config.Id()
	index := index()

	res, err := http.Get(strings.Join([]string{config.Server(), "board", id, index}, "/"))
	common.Fatal(err)

	switch res.StatusCode {
	case 404:
		common.Fatal(errors.New("Clip not found!"))
	case 500:
		common.Fatal(errors.New("Internal server error!"))
	}
	io.Copy(os.Stdout, res.Body)
}

func index() string {
	if len(os.Args) > 2 {
		usage()
	}
	if len(os.Args) == 2 {
		i, err := strconv.Atoi(os.Args[1])
		if err == nil && i >= 1 && i <= 10 {
			return strconv.Itoa(i - 1)
		}
		usage()
	}
	return "0"
}

func usage() {
	common.Fatal(errors.New(
		"Usage: pbp [index]\n  index: nth last clip (1 through 10); default 1",
	))
}
