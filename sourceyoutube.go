package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

var youtubeFeed map[string]interface{}

//var youtubeItems []interface{}

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
	//q.Add("pageToken", "[nextpagetoker]")
	req.URL.RawQuery = q.Encode()

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		checkWarnning(err)
		return
	}

	content, err := ioutil.ReadAll(res.Body)
	checkWarnning(err)

	a1 := content
	b1 := content

	// vids := videos{}
	if err := json.Unmarshal(content, &youtubeFeed); err != nil {
		checkWarnning(err)
	}

	//fmt.Printf("%s \n\n\n", vids)
	fmt.Printf("%s \n", youtubeFeed["items"])
	fmt.Printf("a1 %s, b1 %s \n", len(a1), len(b1))

	fmt.Printf("%s \n", reflect.TypeOf(youtubeFeed["items"]))
	// pasing the youtube feed
	youtubeItems := youtubeFeed["items"].([]interface{})
	for _, item := range youtubeItems {
		viditem := item.(map[string]interface{})
		snippet := viditem["snippet"].(map[string]interface{})
		//thumbs := snippet["thumbnails"].(map[string]interface{})
		//kind := viditem["kind"].(string)
		fmt.Printf("\n============\n %s, %s", snippet["title"], snippet["description"])
	}

	fmt.Printf("\n %s", req.URL.String())
}
