package tests

import (
	"os"
	"path"
	"runtime"

	"github.com/ericklima-ca/formx/router"
)

func Init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func setupRouter() router.Router {
	r := router.NewRouter()
	return r
}
