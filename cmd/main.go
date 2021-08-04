package main

import (
	"fmt"
	_ "os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// "github.com/kr/pretty"
)

var regionCode = []string{"us-east-1",
	"us-east-2",
	"us-west-2",
	"us-west-1",
	"ca-central-1",
	"sa-east-1",
	"eu-west-1",
	"eu-central-1",
	"eu-west-2",
	"eu-west-3",
	"eu-north-1",
	"me-south-1",
	"ap-southeast-1",
	"ap-northeast-1",
	"ap-southeast-2",
	"ap-northeast-2",
	"ap-south-1",
	"ap-east-1"}

var regionRssMap = map[string][]string{}

const serviceUrl string = "https://status.aws.amazon.com/"
const rssUrlPrefix string = "https://status.aws.amazon.com"

func main() {
	for _, value := range regionCode {
		regionRssMap[value] = []string{}
	}
	// DEBUG
	// pretty.Printf("--- m:\n%# v\n\n", regionRssMap)

	parseAmazonStatusPage()

	printRssOnly()
}

func parseAmazonStatusPage() {
	doc, _ := goquery.NewDocument(serviceUrl)
	doc.Find("a").Each(func(_ int, selection *goquery.Selection) {
		hrefVal, _ := selection.Attr("href")
		url := rssUrlPrefix + hrefVal
		switch {
		case strings.HasSuffix(url, "us-east-1.rss"):
			regionRssMap["us-east-1"] = append(regionRssMap["us-east-1"], url)
		case strings.HasSuffix(url, "us-east-2.rss"):
			regionRssMap["us-east-2"] = append(regionRssMap["us-east-2"], url)
		case strings.HasSuffix(url, "us-west-2.rss"):
			regionRssMap["us-west-2"] = append(regionRssMap["us-west-1"], url)
		case strings.HasSuffix(url, "us-west-1.rss"):
			regionRssMap["us-west-1"] = append(regionRssMap["us-west-1"], url)
		case strings.HasSuffix(url, "ca-central-1.rss"):
			regionRssMap["ca-central-1"] = append(regionRssMap["ca-central-1"], url)
		case strings.HasSuffix(url, "sa-east-1.rss"):
			regionRssMap["sa-east-1"] = append(regionRssMap["sa-east-1"], url)
		case strings.HasSuffix(url, "eu-west-1.rss"):
			regionRssMap["eu-west-1"] = append(regionRssMap["eu-west-1"], url)
		case strings.HasSuffix(url, "eu-central-1.rss"):
			regionRssMap["eu-central-1"] = append(regionRssMap["eu-central-1"], url)
		case strings.HasSuffix(url, "eu-west-2.rss"):
			regionRssMap["eu-west-2"] = append(regionRssMap["eu-west-2"], url)
		case strings.HasSuffix(url, "eu-west-3.rss"):
			regionRssMap["eu-west-3"] = append(regionRssMap["eu-west-3"], url)
		case strings.HasSuffix(url, "eu-north-1.rss"):
			regionRssMap["eu-north-1"] = append(regionRssMap["eu-north-1"], url)
		case strings.HasSuffix(url, "me-south-1.rss"):
			regionRssMap["me-south-1"] = append(regionRssMap["me-south-1"], url)
		case strings.HasSuffix(url, "ap-southeast-1.rss"):
			regionRssMap["ap-southeast-1"] = append(regionRssMap["ap-southeast-1"], url)
		case strings.HasSuffix(url, "ap-northeast-1.rss"):
			regionRssMap["ap-northeast-1"] = append(regionRssMap["ap-northeast-1"], url)
		case strings.HasSuffix(url, "ap-southeast-2.rss"):
			regionRssMap["ap-southeast-2"] = append(regionRssMap["ap-southeast-2"], url)
		case strings.HasSuffix(url, "ap-northeast-2.rss"):
			regionRssMap["ap-northeast-2"] = append(regionRssMap["ap-northeast-2"], url)
		case strings.HasSuffix(url, "ap-south-1.rss"):
			regionRssMap["ap-south-1"] = append(regionRssMap["ap-south-1"], url)
		case strings.HasSuffix(url, "ap-east-1.rss"):
			regionRssMap["ap-east-1"] = append(regionRssMap["ap-east-1"], url)
		}
	})
}

func print() {
    for key, value := range regionRssMap {
	fmt.Println(key)
	for _, value1 := range value {
		fmt.Println("     " + value1)
	}
    }
}

func printRssOnly() {
    for _, value := range regionRssMap {
	for _, value1 := range value {
		fmt.Println(value1)
	}
    }
}
