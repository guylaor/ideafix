package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var youtubeFeed map[string]interface{}

// type videos struct {
// 	Kind  string                 `json:"kind"`
// 	Items map[string]interface{} `json:"items"`
// }

func loadYoutube() {

	req, _ := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/playlistItems", nil)

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("maxResults", "50")
	q.Add("playlistId", fmt.Sprint(config["youtube-playlistid"]))
	q.Add("key", fmt.Sprint(config["youtube-key"]))
	req.URL.RawQuery = q.Encode()

	res, err := http.Get(req.URL.String())
	checkError(err)

	content, err := ioutil.ReadAll(res.Body)
	checkError(err)

	// vids := videos{}
	if err := json.Unmarshal(content, &youtubeFeed); err != nil {
		checkError(err)
	}

	//fmt.Printf("%s \n\n\n", vids)
	fmt.Printf("%s \n", reflect.TypeOf(youtubeFeed["items"]))

}
