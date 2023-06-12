package main

import (
	"SimpleDocker/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = utils.Name
	app.Usage = utils.Usage
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{}) // 设置日志格式为JSON格式
		logrus.SetOutput(os.Stdout)                  // 设置日志输出到标准输出
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

}
