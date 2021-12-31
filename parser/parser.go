package github_trending_bot

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	curler "github.com/uncleBlobby/github_trending_bot/curler"
)

func FindDailyTrendingURLS() {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")
	f, err := os.Open(todaysDate + "-dirtyhtml")
	if err != nil {
		log.Fatal(err)
	}
	of, oferr := os.Create(todaysDate + "-barelinks")
	if oferr != nil {
		log.Fatal(err)
	}

	defer f.Close()
	defer of.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0

	interestingLineFound := false
	chunkCounter := 0

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

				of.WriteString(outputLine[1])
				of.WriteString("\n")
				chunkCounter++
			}

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

func FindProjectDirtyHTMLAndWriteOutputFile() {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")
	filename := todaysDate + "-barelinks"
	// open bare links file

	bareLinks, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
	}

	outputFile, err := os.Create(todaysDate)
	if err != nil {
		log.Fatal(err)
	}
	outputFile.Close()

	scanner := bufio.NewScanner(bareLinks)

	for scanner.Scan() {
		projectNameWITHTAG := scanner.Text()
		projectName := RemoveHTTPTAG(scanner.Text())
		dirtyHTML := curler.GetHTMLFromURL(scanner.Text())
		thisFileName := projectName + "-dirtyhtml"
		of, oferr := os.Create(thisFileName)
		if oferr != nil {
			log.Fatal(oferr)
		}
		of.WriteString(dirtyHTML)
		of.Close()

		projectDirtyHTML, err := os.Open(thisFileName)
		if err != nil {
			log.Fatal(err)
		}

		secondScanner := bufio.NewScanner(projectDirtyHTML)

		for secondScanner.Scan() {
			if strings.Contains(secondScanner.Text(), "<title>") {
				f, err := os.OpenFile(todaysDate, os.O_APPEND|+os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				outputDescription := secondScanner.Text()
				outputDescription = removeHTML(outputDescription)

				WriteProjectDetails(f, projectNameWITHTAG, outputDescription)

			}
		}

		projectDirtyHTML.Close()

		err = os.Remove(projectDirtyHTML.Name())
		if err != nil {
			log.Fatal(err)
		}

	}

	bareLinks.Close()
	err = os.Remove(bareLinks.Name())
	if err != nil {
		log.Fatal(err)
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
