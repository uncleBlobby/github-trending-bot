package github_trending_bot

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	curler "github.com/uncleBlobby/github_trending_bot/curler"
)

// interesting text we want to save seems to always start with "article class="Box-row"

func FindDailyTrendingURLS() {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")
	// open the file
	f, err := os.Open(todaysDate + "-dirtyhtml")
	if err != nil {
		log.Fatal(err)
	}
	of, oferr := os.Create(todaysDate + "-barelinks")
	if oferr != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	defer of.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	//interestingLineIDs := []int{}
	interestingLineFound := false
	chunkCounter := 0
	//saveNextLineAsDescription := false
	for scanner.Scan() {

		// do something with a line
		// program finds line of interest and logs it ...
		// need to log the next ten lines after each line of interest and then skip

		if strings.Contains(scanner.Text(), "article class=\"Box-row") {
			//fmt.Printf("line:%d, %s\n", lineNumber, scanner.Text())
			//interestingLineIDs = append(interestingLineIDs, lineNumber)
			interestingLineFound = true
		}
		if interestingLineFound && chunkCounter < 50 {

			if strings.Contains(scanner.Text(), "a href=") &&
				!strings.Contains(scanner.Text(), "a href=\"/login?") &&
				!strings.Contains(scanner.Text(), "network/members") &&
				!strings.Contains(scanner.Text(), "docs.github.com/en/") &&
				!strings.Contains(scanner.Text(), "g-emoji") &&
				!strings.Contains(scanner.Text(), "span") &&
				!strings.Contains(scanner.Text(), "data-hydro-click") {
				outputLine := strings.Fields(scanner.Text())
				outputLine[1] = strings.ReplaceAll(outputLine[1], "href=\"", "https://github.com")
				outputLine[1] = strings.ReplaceAll(outputLine[1], "stargazers\"", "")
				//fmt.Printf("line:%d, %s\n", lineNumber, scanner.Text())
				of.WriteString(outputLine[1])
				of.WriteString("\n")
				chunkCounter++
			}
			/*
				if saveNextLineAsDescription {
					outputLine := strings.Fields(scanner.Text())
					fmt.Printf("desctipion: %v", outputLine[0])
					//of.WriteString(outputDescription[0])
					of.WriteString("\n")
					saveNextLineAsDescription = false
				}
				if strings.Contains(scanner.Text(), "p class=\"col-9") {
					saveNextLineAsDescription = true

				}
			*/

		}
		if chunkCounter >= 50 {
			interestingLineFound = false
			chunkCounter = 0
		}
		lineNumber++
	}
	//fmt.Printf("interesting line IDs: %v", interestingLineIDs)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func FindProjectDirtyHTMLAndWriteToFile() {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")
	filename := todaysDate + "-barelinks"
	// open bare links file

	bareLinks, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
	}

	defer bareLinks.Close()
	scanner := bufio.NewScanner(bareLinks)

	//tempCounter := 0

	for scanner.Scan() {
		projectName := RemoveHTTPTAG(scanner.Text())
		//fileIDString := string(tempCounter)
		dirtyHTML := curler.GetHTMLFromURL(scanner.Text())

		of, oferr := os.Create(projectName + "-dirtyhtml")
		if oferr != nil {
			log.Fatal(oferr)
		}
		of.WriteString(dirtyHTML)
	}

}

func DeleteDirtyHTMLFile() {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")
	filename := todaysDate + "-dirtyhtml"
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}
