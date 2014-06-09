package main

import (
	"github.com/mguindin/goLunch/lunchLib"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lunch"
	app.Version = "0.1"
	app.Usage = "CLI lunch selector, written in Go."
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Search for lunch",
			Flags: []cli.Flag{
				cli.BoolFlag{"debug", "Output request URL"},
				cli.StringFlag{"cuisine", "random", "cuisine to search for"},
				cli.StringFlag{"radius", "0.5", "Radius for search"},
				cli.StringFlag{"location", "10021", "Location to search from"},
				cli.StringFlag{"choice", "1", "Choice in selection"},
			},
			Action: func(c *cli.Context) {
				run(c)
			},
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	lunch := lunchLib.Lunch{
		Radius:   c.String("radius"),
		Location: "&location=" + c.String("location"),
		Debug:    c.Bool("debug"),
		Cuisine:  c.String("cuisine"),
		Yelp_url: "http://api.yelp.com/business_review_search?",
		Rating:   0,
		Rev:      make(map[string]interface{}),
		Choice:   c.Int("choice")}
	yelp_key := lunchLib.GetYelpKey()
	if lunch.Debug {
		fmt.Println(lunch.BuildYelpUrl(yelp_key))
		fmt.Printf("%+v\n", lunch)
	}
	fmt.Print(lunch.ProcessYelpReturn(lunch.MakeRequest(yelp_key)))
}
