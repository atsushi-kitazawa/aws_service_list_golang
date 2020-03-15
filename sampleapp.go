package sampleapp

import (
	"fmt"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

// TODO
// リージョンの日本語表記と英語表記のマッピング
// サービスの日本語表記と英語表記のマッピング
// スクレイピング
// - CSV形式にする
// dat ファイルを作成する

var resionToCode = map[string]string{"Northern Virginia": "us-east-1",
	"Ohio":                "us-east-2",
	"Oregon":              "us-west-2",
	"Northern California": "us-west-1",
	"Montreal":            "ca-central-1",
	"São Paulo":           "sa-east-1",
	"Ireland":             "eu-west-1",
	"Frankfurt":           "eu-central-1",
	"London":              "eu-west-2",
	"Paris":               "eu-west-3",
	"Stockholm":           "eu-north-1",
	"Bahrain":             "me-south-1",
	"Singapore":           "ap-southeast-1",
	"Tokyo":               "ap-northeast-1",
	"Sydney":              "ap-southeast-2",
	"Seoul":               "ap-northeast-2",
	"Mumbai":              "ap-south-1",
	"Hong Kong":           "ap-east-1"}

var serviceToCode = map[string]string{"aaa": "AAA"}

// GetPage return void
func GetPage(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
	})
}

func helloworld() {
	msg1 := "hello"
	msg2 := "world."
	fmt.Printf("%v\n", msg1+" "+msg2)
}

func add(a int, b int) int {
	return a + b
}

// func readFile(file string) string {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		fmt.Printf("error")
// 	}
// 	defer f.Close()
// 	var n int
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := f.Read(buf)
// 		if n == 0 {
// 			break
// 		}
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

func readFile(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error")
	}
	return string(bytes)
}
