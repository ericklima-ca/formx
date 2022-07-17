package tests

import (
	"github.com/ericklima-ca/formx/router"
)

func setupRouter() router.Router {
	r := router.NewRouter("../static")
	return r
}
