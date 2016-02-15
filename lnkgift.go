/************************************************
File:			lnkgift.go
Version:		0.0.1
Description:	main
Author:			Qianno.Xie
Email:			qianlnk@163.com
************************************************/
package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/golang/glog"
	"lnkgift/config"
	"lnkgift/router"
	"os"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			glog.V(1).Infof("ERROR: recover err:%+v", err)
		}
	}()

	defer glog.Flush()
	ticker := time.NewTicker(1000 * 1000 * 100)
	go func() {
		for _ = range ticker.C {
			glog.Flush()
		}
	}()
	defer ticker.Stop()
	flag.Parse()
	conf := config.GetConfig()
	fmt.Println(conf)
	app := cli.NewApp()
	app.Name = "lnkgift server"
	app.Author = "Qianno.Xie"
	app.Version = "0.0.1"
	app.Usage = "lnkgift [port] default 9000"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) {
		port := ":9000"
		m := router.NewRouter()
		fmt.Printf("c is %+v, m is %+v\n", c.Args().First(), m)
		if c.Args().First() != "" {
			port = ":" + c.Args().First()
		}
		m.RunOnAddr(port)
	}
	app.Run(os.Args)

}
