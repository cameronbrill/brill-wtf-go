package main

import (
	"github.com/SuperPaintman/nice/cli"
	"github.com/cameronbrill/brill-wtf-go/link"
)

func main() {
	app := cli.App{
		Name:  "link",
		Usage: cli.Usage("Generate a short url!"),
		Action: cli.ActionFunc(func(cmd *cli.Command) cli.ActionRunner {

			url := cli.StringArg(cmd, "url", cli.Usage("The URL you are shortening"))

			return func(cmd *cli.Command) error {
				l, err := link.New(link.URL(*url))
				if err != nil {
					return err
				}
				cmd.Println("Visit ", (&l).String())
				return nil
			}
		}),
		CommandFlags: []cli.CommandFlag{
			cli.HelpCommandFlag(),
			cli.VersionCommandFlag("0.0.0"),
		},
	}

	app.HandleError(app.Run())
}
