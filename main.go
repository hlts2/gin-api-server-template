package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/config"
	"github.com/hlts2/gin-server-template/interfaces"
	"github.com/kpango/glg"
	"github.com/pkg/errors"
)

type params struct {
	configFilePath string
}

func parseParams() (*params, error) {
	p := new(params)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&p.configFilePath,
		"f",
		"/etc/server/config.yaml",
		"config file path")

	err := f.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	return p, nil
}

func init() {
	p, err := parseParams()
	if err != nil {
		glg.Fatal(errors.Wrap(err, "parse params faild"))
	}

	err = config.Init(p.configFilePath)
	if err != nil {
		glg.Fatal(errors.Wrap(err, "faild to initialize config file"))
	}
}

func main() {
	g := gin.Default()

	g.GET("/Users/", interfaces.GetUsers)
	g.GET("/Users/:id", interfaces.GetUser)

	if err := g.Run(":" + config.GetConfig().Server.Port); err != nil {
		glg.Fatal(err)
	}
}
