/************************************************
File: 		router.go
Version:		0.0.1
Decription:	all Api router
Author:		Qianno.Xie
Email:		qianno.xie@appcaochs.com
History:		2015.12.07 created by Qianno.Xie
************************************************/
package router

import (
	"github.com/go-martini/martini"
	//"haostudent/config"
	"lnkgift/handlers/user"
	//"html/template"
	//"net/http"
	//"fmt"
)

func NewRouter() *martini.ClassicMartini {
	m := martini.Classic()
	m.Use(martini.Static("template"))
	m.Get("/Login", user.Login)
	return m
}
