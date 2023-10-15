package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"time"
)

func main() {
	u := launcher.New().
		Headless(false).
		MustLaunch()
	browser := rod.New().
		ControlURL(u).
		MustConnect().
		// disable default viewport
		NoDefaultDevice()
	browser.MustPage("chrome://new-tab-page").MustWaitLoad()
	time.Sleep(10 * time.Minute)
}
