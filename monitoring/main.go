package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/otiai10/twistream"
)

type Conf struct {
	ConsumerKey       string
	ConsumerSelect    string
	AccessToken       string
	AccessTokenSelect string
	MailAddress       string
	MailPassword      string
	SMTPServer        string
	SMTPPort          string
}

func readConf() *Conf {
	js, _ := ioutil.ReadFile("../conf.json")
	var ret Conf
	json.Unmarshal(js, &ret)
	return &ret
}

func alert(upTime int64) int64 {
	c := readConf()
	t := time.Now()
	unix := t.Unix()
	if unix-upTime > 60 {
		auth := smtp.PlainAuth("", c.MailAddress, c.MailPassword, c.SMTPServer)
		err := smtp.SendMail(
			c.SMTPServer+":"+c.SMTPPort,
			auth,
			c.MailAddress,
			[]string{c.MailAddress},
			[]byte("raybanスパム感染の疑いあり"),
		)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return unix
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
	uptime := int64(0)

	// Listen timeline
	for {
		status := <-timeline.Listen()
		if strings.Contains(status.Text, "レイバンの") {
			api.DeleteTweet(status.Id, true)
			uptime = alert(uptime)
		}
	}
}
