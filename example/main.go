package main

import (
	"fmt"
	"github.com/SilkageNet/wmocr"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var app = cli.NewApp()
	app.Name = "OCR Server"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "dllPath", Usage: "指定完美dll路径", Value: "WmCode.dll"},
		&cli.StringFlag{Name: "datPath", Required: true, Usage: "指定数据文件路径"},
		&cli.StringFlag{Name: "datPwd", Required: true, Usage: "指定数据文件密码"},
		&cli.StringFlag{Name: "addr", Usage: "GRPC服务地址", Value: ":6789"},
	}
	app.Action = action
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println("app run err: ", err)
	}
}

func action(c *cli.Context) error {
	var ocr, err = wmocr.NewWmOcr(c.String("dllPath"), c.String("datPath"), c.String("datPwd"))
	if err != nil {
		return fmt.Errorf("init.wm.ocr.err:%s", err.Error())
	}
	var server = NewServer(ocr)
	return server.Serve(c.String("addr"))
}
