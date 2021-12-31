package github_trending_bot

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseProjectForDescription() {
	// open the input file
	f, err := os.Open("particularProjectPage")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "<title>") {
			outputDescription := scanner.Text()
			outputDescription = removeHTML(outputDescription)
			appendDescriptionToOutputFile()
			fmt.Println(outputDescription)
		}

	}

}

func appendDescriptionToOutputFile() {
	f2, err := os.OpenFile("2021-12-31", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f2.WriteString("appended to file\n"); err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

}

func removeHTML(s string) string {
	if strings.Contains(s, "<title>") {
		s = strings.ReplaceAll(s, "<title>", "")
	}
	if strings.Contains(s, "</title>") {
		s = strings.ReplaceAll(s, "</title>", "")
	}
	if strings.Contains(s, "GitHub") {
		s = strings.ReplaceAll(s, "GitHub", "")
	}
	return s
}

func RemoveHTTPTAG(s string) string {
	if strings.Contains(s, "https://") {
		s = strings.ReplaceAll(s, "https://", "")
	}
	if strings.Contains(s, "/") {
		s = strings.ReplaceAll(s, "/", "")
	}
	return s
}
