package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/emulation"
)

type speedCheck struct{
	DownloadSpeed string
	DownloadSpeedUnit string
	UploadSpeed string
	UploadSpeedUnit string
	Latency string
	LatencyUnit string
	UserIP string
	UserISP string
	Serverlocation string
}

func main(){
	start := time.Now()
	fmt.Println("Starting to measure, this can take upto a minute")
	speed, err := measureSpeed()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download speed:", speed.DownloadSpeed, speed.DownloadSpeedUnit)
	fmt.Println("Upload speed:", speed.UploadSpeed, speed.UploadSpeedUnit)
	fmt.Println("Ping:", speed.Latency, speed.LatencyUnit)
	fmt.Println("IP adress:", speed.UserIP)
	fmt.Println("Internet provider:", speed.UserISP)
	fmt.Println("Server location:", speed.Serverlocation)
	fmt.Println("Duration of operation:", time.Since(start).Seconds(), "seconds")
}

func measureSpeed() (*speedCheck, error){
	chromeContext, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	chromeContext, cancel = context.WithTimeout(chromeContext, 60*time.Second)
	defer cancel()
	speedCheck := new(speedCheck)
	err  := chromedp.Run(chromeContext,
		emulation.SetUserAgentOverride(`chromedp/chromedp v0.6.10`),
		chromedp.Navigate(`https://fast.com`),
		chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Text(`#speed-value.succeeded`, &speedCheck.DownloadSpeed, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#speed-units.succeeded`, &speedCheck.DownloadSpeedUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#upload-value.succeeded`, &speedCheck.UploadSpeed, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#upload-units.succeeded`, &speedCheck.UploadSpeedUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#latency-value.succeeded`, &speedCheck.Latency, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#latency-units.succeeded`, &speedCheck.LatencyUnit, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#user-ip`, &speedCheck.UserIP, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#user-isp`, &speedCheck.UserISP, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#server-locations`, &speedCheck.Serverlocation, chromedp.NodeVisible, chromedp.ByQuery),
	)
	return speedCheck, err
}