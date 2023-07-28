package main

import (
	"fmt"
	"log"
	"os"

	"github.com/g-code99/learning-path/week1/contador-palabras/pkg/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "contador",
		Usage:   "contador de palabras de un fichero txt",
		Version: "v0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Value:   "all",
				Usage:   "Count Lines and Words",
			},
		},
		Action: func(ctx *cli.Context) error {
			// var file string = ctx.Args().First()
			file := "test.txt"
			if ctx.NArg() == 0 || !utils.IsVerifyDir(file) {
				return nil
			}

			switch ctx.String("mode") {
			case "all":
				words, lines := utils.ReadAndCountWordAndLines(file)
				fmt.Printf("El archivo %s tiene %d palabras y %d saltos de linea", file, words, lines)
				// case "words":
				// 	words := utils.ReadAndCountWords(file)
				// 	fmt.Printf("El archivo %s tiene %d palabras", file, words)
				// case "lines":
				// 	lines := utils.ReadAndCountLines(file)
				// 	fmt.Printf("El archivo %s tiene %d saltos de linea", file, lines)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
