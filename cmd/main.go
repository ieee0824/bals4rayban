package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

type Conf struct {
	ConsumerKey       string
	ConsumerSelect    string
	AccessToken       string
	AccessTokenSelect string
}

func createSettingTemp() {
	c := Conf{ConsumerKey: "", ConsumerSelect: "", AccessToken: "", AccessTokenSelect: ""}
	js, _ := json.Marshal(c)
	ioutil.WriteFile("../conf-example.json", js, 0644)
}

func readConf() *Conf {
	js, _ := ioutil.ReadFile("../conf.json")
	var ret Conf
	json.Unmarshal(js, &ret)
	return &ret
}

func bals() bool {
	ret := false
	c := readConf()
	anaconda.SetConsumerKey(c.ConsumerKey)
	anaconda.SetConsumerSecret(c.ConsumerSelect)
	api := anaconda.NewTwitterApi(c.AccessToken, c.AccessTokenSelect)

	results, _ := api.GetUserTimeline(nil)
	for _, r := range results {
		if strings.Contains(r.Text, "レイバンの") {
			ret = true
			fmt.Println(r.Text)
			api.DeleteTweet(r.Id, true)
		}
	}
	return ret
}

func main() {
	for bals() {
	}
}
