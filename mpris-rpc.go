package main

import (
	"fmt"
	"os"
	"strings"
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
        ConfPlayer string `kdl:"player"`
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
    var name string // create an empty string variable
    for _, e := range names { // iterate over all available names and check if any of them match our configured name
        if (strings.Contains(e, conf.ConfPlayer)) {
            // I used contains since the player will be named something like 'org.freedesktop.elisa' which is a mouthfull
            // if we match a player, define the name as that player and break the iterator loopja
            // this means that we match our first found match
            name = e
            break
        }
    }
	player := mpris.New(conn, name)
    // --NOTE this for now just crashes if the media player find nothing, we should prob change that -w-
	for {
        fmt.Println(name);
		status, err := player.GetMetadata()
		if err != nil {
			panic(err)
		}
		err = client.SetActivity(client.Activity{
            State:      fmt.Sprint("by ", status["xesam:artist"].String()[2 : len(status["xesam:artist"].String()) -2]),
            // .String returns the String vallue, for artist it's encased in a [""] so I just cut those of with a slice
            Details:    fmt.Sprint(status["xesam:title"].String()[1 : len(status["xesam:title"].String()) -1]),
            // same here but with title it just encases it in ""
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
