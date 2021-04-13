package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/emulation"
)

type SpeedChecker struct{
	downloadSpeed string
	downloadSpeedUnit string
	uploadSpeed string
	uploadSpeedUnit string
	latency string
	latencyUnit string
	userIP string
	userISP string
	serverlocation string
}

func main(){
	start := time.Now()
	fmt.Println("Starting to measure, this can take upto a minute")
	speed, err := measureSpeed()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download speed:", speed.downloadSpeed, speed.downloadSpeedUnit)
	fmt.Println("Upload speed:", speed.uploadSpeed, speed.uploadSpeedUnit)
	fmt.Println("Ping:", speed.latency, speed.latencyUnit)
	fmt.Println("IP adress:", speed.userIP)
	fmt.Println("Internet provider:", speed.userISP)
	fmt.Println("Server location:", speed.serverlocation)
	fmt.Println("Duration of operation:", time.Since(start).Seconds(), "seconds")
}

func measureSpeed() (*SpeedChecker, error){
	chromeContext, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	chromeContext, cancel = context.WithTimeout(chromeContext, 60*time.Second)
	defer cancel()
	speedChecker := new(SpeedChecker)
	err  := chromedp.Run(chromeContext,
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &speedChecker.downloadSpeed, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &speedChecker.downloadSpeedUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &speedChecker.uploadSpeed, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#upload-units.succeeded`, &speedChecker.uploadSpeedUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#latency-value.succeeded`, &speedChecker.latency, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#latency-units.succeeded`, &speedChecker.latencyUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#user-ip`, &speedChecker.userIP, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#user-isp`, &speedChecker.userISP, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#server-locations`, &speedChecker.serverlocation, chromedp.NodeVisible, chromedp.ByQuery),
	)
	return speedChecker, err
}