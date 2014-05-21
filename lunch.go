package lunch

import (
    "io/ioutil"
	"strings"
)

func main() {
    radius := "0.5"
    location := "10021"
    debug := false
    cuisine := "random"
    yelp_url := "http://api.yelp.com/business_review_search?"
    rating, rev := 0, 0
}

func GetYelpKey() string {
    yelp_key := ""
    b, err := ioutil.ReadFile("yelp_key")
    if err != nil { panic(err) }
    return strings.Replace(string(b), "\n", "", 1)
}
