package main

import (
    "io/ioutil"
	"strings"
	"fmt"
)

type Lunch struct {
	yelp_url, radius, location, cuisine string
	debug bool
	rating, rev, choice int8
}

func main() {
	lunch := Lunch{
		radius : "0.5",
		location : "10021",
		debug : false,
		cuisine : "random",
		yelp_url : "http://api.yelp.com/business_review_search?",
		rating : 0,
		rev : 0,
		choice : 1}
	fmt.Print(BuildYelpURL(lunch))
}

func GetYelpKey() string {
    b, err := ioutil.ReadFile("yelp_key")
    if err != nil { panic(err) }
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
