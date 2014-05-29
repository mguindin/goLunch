package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"strconv"
	"math"
	"github.com/mguindin/lunch/third_party/github.com/codegangsta/cli"
)

type Lunch struct {
	yelp_url, radius, location, cuisine string
	debug                               bool
	rating                              int8
	rev									map[string] interface{}
	choice                              int
}

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
			Flags : []cli.Flag{
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
	lunch := Lunch{
		radius:   c.String("radius"),
		location: c.String("location"),
		debug:    c.Bool("debug"),
		cuisine:  c.String("cuisine"),
		yelp_url: "http://api.yelp.com/business_review_search?",
		rating:   0,
		rev:      make(map[string]interface{}),
		choice:   c.Int("choice")}
	if lunch.debug {
		fmt.Println(BuildYelpURL(lunch))
		fmt.Printf("%+v\n", lunch)
	}
	ProcessYelpReturn(MakeRequest(BuildYelpURL(lunch)), lunch)
}

func ProcessYelpReturn(ret []byte, lunch Lunch) {
	var dat map[string]interface{}
	err := json.Unmarshal(ret, &dat)
	if err != nil {
		fmt.Println("error:", err)
	}
	if (lunch.debug) {
		fmt.Printf("businesses %+v\n", dat["businesses"])
		fmt.Printf("message %+v\n", dat["message"])
	}
	businesses := dat["businesses"].([]interface{})
	restaurant := businesses[0].(map[string]interface{})
	fmt.Println("You will be having " + restaurant["name"].(string) +
			", which is located at " + restaurant["address1"].(string) + ".")
	rating := restaurant["avg_rating"].(float64)
	fmt.Println(restaurant["name"].(string) + " has a rating of " + strconv.FormatFloat(rating, 'f', -1, 64))
	reviews := restaurant["reviews"].([]interface{})
	for _, val := range reviews {
		m := val.(map[string] interface{})
		if (math.Abs(rating - m["rating"].(float64)) < 1.0) {
			lunch.rev = m
		}
	}
	review_text := strings.Replace(lunch.rev["text_excerpt"].(string), "\n", " ", -1)
	fmt.Println("People are saying: " + review_text)

}

func GetYelpKey() string {
	b, err := ioutil.ReadFile("yelp_key")
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(b), "\n", "", 1)
}

func BuildYelpURL(lunch Lunch) string {
	return lunch.yelp_url +
		"term=" + lunch.cuisine +
		"&location=" + lunch.location +
		"&radius=" + lunch.radius +
		"&limit=20" +
		"&ywsid=" + GetYelpKey() +
		"&category=restaurants"
}

func MakeRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(body))
	return body
}
