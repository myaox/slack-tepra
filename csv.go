package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func urlConv(url string) string {
	url = strings.Trim(url, "<")
	url = strings.Trim(url, ">")
	urls := strings.Split(url, "|")

	return urls[0]
}

func csvPath(name string) string {
	return path.Join(getCurPath(), name)
}

func getCurPath() string {
	dir, _ := os.Getwd()
	return dir
}

func writeCsv(args []string) string {
	csv := csvPath("value.csv")
	file, _ := os.Create(csv)
	val := strings.Join(args, ",")

	sjis, _ := utf8ToSjis(val)
	content := []byte(sjis)

	file.Write(content)
	file.Close()
	return csv
}

func utf8ToSjis(str string) (string, error) {
	iostr := strings.NewReader(str)
	rio := transform.NewReader(iostr, japanese.ShiftJIS.NewEncoder())
	ret, err := ioutil.ReadAll(rio)
	if err != nil {
		return "", err
	}
	return string(ret), err
}
