package lunchLib

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"strconv"
)


type Lunch struct {
	Yelp_url, Radius, Location, Cuisine string
	Debug                               bool
	Rating                              int8
	Rev                                 map[string]interface{}
	Choice                              int
}

func (lunch *Lunch) BuildYelpUrl() string {
	return lunch.Yelp_url +
			"term=" + lunch.Cuisine +
			"&location=" + lunch.Location +
			"&radius=" + lunch.Radius +
			"&limit=20" +
			"&ywsid=" + GetYelpKey() +
			"&category=restaurants"
}

func (lunch *Lunch) MakeRequest() []byte {
	resp, err := http.Get(lunch.BuildYelpUrl())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(body))
	return body
}

func (lunch *Lunch) ProcessYelpReturn(ret []byte) string {
	out := ""
	var dat map[string]interface{}
	err := json.Unmarshal(ret, &dat)
	if err != nil {
		fmt.Println("error:", err)
	}
	if lunch.Debug {
		fmt.Printf("businesses %+v\n", dat["businesses"])
		fmt.Printf("message %+v\n", dat["message"])
	}
	businesses := dat["businesses"].([]interface{})
	restaurant := businesses[0].(map[string]interface{})
	out += "You will be having " + restaurant["name"].(string) +
			", which is located at " + restaurant["address1"].(string) + ".\n"
	rating := restaurant["avg_rating"].(float64)
	out += restaurant["name"].(string) + " has a rating of " + strconv.FormatFloat(rating, 'f', -1, 64)
	reviews := restaurant["reviews"].([]interface{})
	for _, val := range reviews {
		m := val.(map[string]interface{})
		if math.Abs(rating-m["rating"].(float64)) < 1.0 {
			lunch.Rev = m
		}
	}
	review_text := strings.Replace(lunch.Rev["text_excerpt"].(string), "\n", " ", -1)
	return out + "People are saying: " + review_text
}

func GetYelpKey() string {
	b, err := ioutil.ReadFile("yelp_key")
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(b), "\n", "", 1)
}
