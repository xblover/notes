package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/otiai10/opengraph"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Usage = "Fetch URL and extract OpenGraph meta informations."
	app.UsageText = "ogp [-A] {URL}"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "absolute,A",
			Usage: "populate relative URLs to absolute URLs",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		rawurl := ctx.Args().First()
		if rawurl == "" {
			return fmt.Errorf("URL must be specified")
		}
		og, err := opengraph.Fetch(rawurl)
		if err != nil {
			return err
		}
		if ctx.Bool("absolute") {
			og = og.ToAbsURL()
		}
		b, err := json.MarshalIndent(og, "", "\t")
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", string(b))
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
