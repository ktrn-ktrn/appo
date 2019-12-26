package controllers

import (
	"appo/app/models/auth"

	"log"

	"github.com/revel/revel"
)

type CAuth struct {
	*revel.Controller
	provider *auth.AuthProvider
}

func (c *CAuth) Init() {

	c.provider = new(auth.AuthProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CAuth) Login() revel.Result {
	c.Init()

	return nil
}

func (c *CAuth) Logout() revel.Result {
	c.Init()

	return nil
}
