package config

import (
	"github.com/pbcp/pbcp/common"

	"os"
	"os/user"
	"io/ioutil"
	"path"
	"net/http"
)

func Id() string {
	file := path.Join(home(), ".config", "pbcp", "id")
	err := os.MkdirAll(file)
	common.Fatal(err)

	if _, err := os.Stat(file); err == nil {
		contents, err := ioutil.ReadFile(file)
		common.Fatal(err)
		return string(contents)
	}

	res, err := http.Get(path.Join(Server(), "register"))
	common.Fatal(err)

	id, err := ioutil.ReadAll(res.Body)
	common.Fatal(err)

	err = ioutil.WriteFile(file, id, 0600)
	common.Fatal(err)

	return string(id)
}

func Server() string {
	file := path.Join(home(), ".config", "pbcp", "server")
	err := os.MkdirAll(file)
	common.Fatal(err)

	if _, err := os.Stat(file); err == nil {
		contents, err := ioutil.ReadFile(file)
		common.Fatal(err)
		return string(contents)
	}

	return DefaultServer
}

func home() string {
	u, err := user.Current()
	common.Fatal(err)

	return u.HomeDir
}