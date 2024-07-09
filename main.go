package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("Website Copier Starts...")

	dir := inputProjectName()

	err := os.Mkdir(dir, 0755)
	handleErr(err)

	url := inputUrl()

	body, _ := getPageBody(url)

	parseResBodyAsHtml(string(body), url, dir)

	fmt.Println("Website Copy successfully...")
}

func inputProjectName() string {
	fmt.Println("Enter Project Name :")

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	handleErr(err)
	name = strings.TrimSpace(name)

	return name

}
func inputUrl() string {
	fmt.Println("Enter Website Url :")

	reader := bufio.NewReader(os.Stdin)

	url, err := reader.ReadString('\n')
	handleErr(err)

	url = strings.TrimSpace(url)
	return url
}

func parseResBodyAsHtml(body string, url string, dir string) {
	doc, err := html.Parse(strings.NewReader(body))
	handleErr(err)
	extractAnchorTags(doc, url, dir)
}
func extractAnchorTags(doc *html.Node, url string, dir string) {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, attr := range doc.Attr {
			if attr.Key == "href" {
				result := strings.Contains(attr.Val, url)
				if result {
					body, _ := getPageBody(attr.Val)
					createFile(attr.Val, url, dir, body)
				}
			}
		}
	}
	for i := doc.FirstChild; i != nil; i = i.NextSibling {
		extractAnchorTags(i, url, dir)
	}
}
func handleErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
func createFile(url string, baseUrl string, dir string, body []byte) {
	fileName := strings.Replace(url, baseUrl, "", -1)
	if fileName == "" {
		fileName = "index"
	}
	path := dir + "/" + fileName + ".html"

	file := os.WriteFile(path, body, 0644)
	if file != nil {
		parts := strings.Split(path, "/")

		lastPart := parts[len(parts)-1]
		if filepath.Ext(lastPart) != "" {
			parts = parts[:len(parts)-1]
		}
		checkdir := filepath.Join(parts...)
		if _, err := os.Stat(checkdir); os.IsNotExist(err) {
			err := os.MkdirAll(checkdir, 0755)
			handleErr(err)
			createFile(url, baseUrl, dir, body)
		} else {
			fmt.Println("The provided directory named", dir, "exists")
		}
	}
}
func getPageBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	handleErr(err)

	body, err := io.ReadAll(res.Body)
	handleErr(err)
	res.Body.Close()
	return body, nil
}
