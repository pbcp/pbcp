package main

import (
	"github.com/pbcp/pbcp/common"
	"github.com/pbcp/pbcp/common/config"

	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	id := config.Id()

	res, err := http.Get(path.Join(Server(), "board", id, "0"))
	common.Fatal(err)


}