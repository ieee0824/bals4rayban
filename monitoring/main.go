package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/otiai10/twistream"
)

type Conf struct {
	ConsumerKey       string
	ConsumerSelect    string
	AccessToken       string
	AccessTokenSelect string
}

func readConf() *Conf {
	js, _ := ioutil.ReadFile("../conf.json")
	var ret Conf
	json.Unmarshal(js, &ret)
	return &ret
}

func main() {
	c := readConf()
	timeline, _ := twistream.New(
		"https://userstream.twitter.com/1.1/user.json",
		c.ConsumerKey,
		c.ConsumerSelect,
		c.AccessToken,
		c.AccessTokenSelect,
	)

	anaconda.SetConsumerKey(c.ConsumerKey)
	anaconda.SetConsumerSecret(c.ConsumerSelect)
	api := anaconda.NewTwitterApi(c.AccessToken, c.AccessTokenSelect)
	fmt.Println(api)

	// Listen timeline
	for {
		status := <-timeline.Listen()
		//fmt.Println(status.User.Name)
		//fmt.Println(status.User.Id)
		//fmt.Println(status.Text)
		//if strings.Contains(status.Text, "レイバンの") {
		if strings.Contains(status.Text, "レイバンの") {
			api.DeleteTweet(status.Id, true)
		}
	}
}
