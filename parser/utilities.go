package github_trending_bot

import (
	"os"
	"strings"
)

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
		s = strings.ReplaceAll(s, "https://github.com", "")
	}
	if strings.Contains(s, "/") {
		s = strings.ReplaceAll(s, "/", "")
	}

	return s
}

func WriteProjectDetails(f *os.File, pname string, pdescript string) {
	f.WriteString("Project: ")
	f.WriteString(pname)
	f.WriteString("\n")
	f.WriteString("Description: ")
	f.WriteString(pdescript)
	f.WriteString("\n")
	f.WriteString("\n")
	f.Close()
}

func ReturnProjectNameFromURL(s string) string {
	projectName := ""
	if strings.Contains(s, "https://github.com/") {
		projectName = strings.ReplaceAll(s, "https://github.com/", "")
	}
	return projectName
}

/*
func cleanUpDescriptionLine(dirtyD string, projectURL string) string {
	cleanD := ""
	projectName := ReturnProjectNameFromURL(projectURL)
	projectString := "- " + projectName + ":"
	cleanD += strings.ReplaceAll(dirtyD, projectString, "")

	fmt.Println("Project name: " + projectName)
	fmt.Println(cleanD)
	fmt.Println(dirtyD == cleanD)
	return cleanD

}
*/
