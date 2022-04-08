package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const URL string = "https://health.aws.amazon.com/health/status"

var regions []string = []string{"NA", "SA", "EU", "AF", "AP", "ME"}

// var wg sync.WaitGroup

func main() {
	doMain()
}

func doMain() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// wg.Add(2)
	for _, r := range regions {
		scrapingDashboard(ctx, r)
	}
	// wg.Wait()

}

func scrapingDashboard(ctx context.Context, region string) {
	// defer wg.Done()
	element := fmt.Sprintf("[href='%s']", region)
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		chromedp.Click(element, chromedp.NodeVisible),
		chromedp.Click(`document.querySelector("#status-history-table > div.awsui_footer_14iqq_1dn1p_75 > div > div > div > a")`, chromedp.ByJSPath),
		chromedp.Sleep(5*time.Second),
		chromedp.Nodes(`.status-history-rss-feed-button`, &nodes, chromedp.BySearch),
	)
	if err != nil {
		panic(err)
	}
	// fmt.Println("===== " + region + " =====")
	// fmt.Println(len(nodes))
	for i, _ := range nodes {
		fmt.Println(nodes[i].Attributes[3])
	}
}
