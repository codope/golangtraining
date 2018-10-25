package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var imgDb = []string{"sagar", "sagar2", "sagars", "sumit", "sagars3", "hellosagar"}
	var newsDb = []string{"sagarnews", "sagar2News", "sagarsnews", "newssumit", "newssagars3", "newshellosagar"}
	var webDb = []string{"websagar", "sagar2web", "sagarsweb", "sumitweb", "sagars3web", "hellosagarweb"}
	var vidDb = []string{"vidsagar", "sagar2vid", "vidsagars", "sumitvid", "vidsagars3", "hellosagarvid"}

	searchStr := "sagar"

	// For our example we'll select across two channels.
	imgChan := make(chan string, 3)
	newsChan := make(chan string, 3)
	webChan := make(chan string, 3)
	vidChan := make(chan string, 3)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		for _, v := range imgDb {
			if strings.Contains(v, searchStr) {
				time.Sleep(time.Second * 1)
				imgChan <- v
			}

		}
		close(imgChan)
	}()

	go func() {
		for _, v := range newsDb {
			if strings.Contains(v, searchStr) {
				time.Sleep(time.Second * 1)
				newsChan <- v
			}

		}
		close(newsChan)
	}()

	go func() {
		for _, v := range webDb {
			if strings.Contains(v, searchStr) {
				time.Sleep(time.Second * 1)
				webChan <- v
			}

		}
		close(webChan)
	}()

	go func() {
		for _, v := range vidDb {
			if strings.Contains(v, searchStr) {
				time.Sleep(time.Second * 1)
				vidChan <- v
			}

		}
		close(vidChan)
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for ; imgChan != nil && newsChan != nil && webChan != nil && vidChan != nil;  {
		select {
		case x, ok := <-imgChan:
			fmt.Println("image", x, ok)
			if !ok {
				imgChan = nil
			}
		case x, ok := <-newsChan:
			fmt.Println("news", x, ok)
			if !ok {
				newsChan = nil
			}
		case x, ok := <-webChan:
			fmt.Println("web", x, ok)
			if !ok {
				webChan = nil
			}
		case x, ok := <-vidChan:
			fmt.Println("video", x, ok)
			if !ok {
				vidChan = nil
			}
		}
	}
}
