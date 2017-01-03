package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func loadFacebook() {

	url := fmt.Sprintf("https://graph.facebook.com/v2.8/%s", config["facebook-pageid"])
	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("fields", "feed{picture,description,name,permalink_url,story,full_picture,created_time,link,status_type,shares,call_to_action,caption}")
	q.Add("access_token", fmt.Sprint(config["facebook-token"]))

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		checkWarnning(err)
	}

	content, err := ioutil.ReadAll(res.Body)
	checkWarnning(err)

	fmt.Printf("%s", content)

}
