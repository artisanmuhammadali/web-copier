package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Website Copier Starts...")

	dir := inputProjectName()

	err := os.Mkdir(dir, 0755)
	handleErr(err)

	url := inputUrl()

	res, err := http.Get(url)
	handleErr(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	handleErr(err)

	path := dir + "/index.html"
	file := os.WriteFile(path, body, 0644)
	handleErr(file)
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
func handleErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
