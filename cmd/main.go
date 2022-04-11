package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const URL string = "https://health.aws.amazon.com/health/status"

var (
	regions []string = []string{"NA", "SA", "EU", "AF", "AP", "ME"}
	wg      sync.WaitGroup
	mu      sync.Mutex
	wait    int
)

func init() {
	flag.IntVar(&wait, "wait", 30, "time to wait for page to load.")
	flag.Parse()
}

func main() {
	// fmt.Printf("wait=%v\n", wait)
	doMain()
}

func doMain() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	wg.Add(len(regions))
	for _, r := range regions {
		subCtx, _ := chromedp.NewContext(ctx)
		go scrapingDashboard(subCtx, r)
	}
	wg.Wait()

}

func scrapingDashboard(ctx context.Context, region string) {
	defer wg.Done()
	element := fmt.Sprintf("[href='%s']", region)
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(URL),
		chromedp.Click(element, chromedp.NodeVisible),
		chromedp.Click(`document.querySelector("#status-history-table > div.awsui_footer_14iqq_1dn1p_75 > div > div > div > a")`, chromedp.ByJSPath),
		chromedp.Sleep(time.Duration(wait)*time.Second),
		chromedp.Nodes(`.status-history-rss-feed-button`, &nodes, chromedp.BySearch),
	)
	if err != nil {
		panic(err)
	}

	print(nodes)
}

func print(nodes []*cdp.Node) {
	mu.Lock()
	defer mu.Unlock()

	// fmt.Println("===== " + region + " =====")
	// fmt.Println(len(nodes))
	for i, _ := range nodes {
		fmt.Println(nodes[i].Attributes[3])
	}
}
