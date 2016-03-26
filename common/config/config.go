package config

import (
	"github.com/pbcp/pbcp/common"

	"os"
	"os/user"
	"io/ioutil"
	"path"
	"net/http"
	"strings"
)

func Id() string {
	dir := path.Join(home(), ".config", "pbcp")
	file := path.Join(dir, "id")
	err := os.MkdirAll(dir, 0700)
	common.Fatal(err)

	if _, err := os.Stat(file); err == nil {
		contents, err := ioutil.ReadFile(file)
		common.Fatal(err)
		return strings.TrimSpace(string(contents))
	}

	res, err := http.Get(strings.Join([]string{Server(), "register"}, "/"))
	common.Fatal(err)

	id, err := ioutil.ReadAll(res.Body)
	common.Fatal(err)

	err = ioutil.WriteFile(file, id, 0600)
	common.Fatal(err)

	return strings.TrimSpace(string(id))
}

func Server() string {
	dir := path.Join(home(), ".config", "pbcp")
	file := path.Join(dir, "server")
	err := os.MkdirAll(dir, 0700)
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