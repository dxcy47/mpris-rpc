package main

import (
	"fmt"
	"os"
	"time"
	"github.com/Pauloo27/go-mpris"
	"github.com/godbus/dbus/v5"
	"github.com/hugolgst/rich-go/client"
	"github.com/sblinch/kdl-go"
)

func main() {

	type Conf struct {
		Appid   string `kdl:"appid"`
		Updtime int `kdl:"updtime"`
	}
	var conf Conf
	
	confdir, err := os.UserConfigDir()
	kdlconf := confdir + "/mpris-rpc/mpris-rpc.kdl"
	file, err := os.ReadFile(kdlconf)
	if err != nil {
		panic(err)
	}
 	err = kdl.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}


	err = client.Login(conf.Appid)
	if err != nil {
		panic(err)
	}
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	names, err := mpris.List(conn)
	if err != nil {
		panic(err)
	}
	if len(names) == 0 {
		panic("No player found")
	}
	name := names[0]
	player := mpris.New(conn, name)
	for {
		status, err := player.GetMetadata()
		if err != nil {
			panic(err)
		}

		err = client.SetActivity(client.Activity{
			State:      fmt.Sprint("by", status["xesam:artist"]),
			Details:    fmt.Sprint(status["xesam:title"]),
			LargeImage: "BINGLE.jpg",
			LargeText:  "This is the large image :D",
			SmallImage: "smallimageid",
			SmallText:  "And this is the small image",
		})
		if err != nil {
			panic(err)
		}
 	
		time.Sleep(time.Duration(conf.Updtime) * time.Second)

	}
}
