package main

import (
	curler "github.com/uncleBlobby/github_trending_bot/curler"
	parser "github.com/uncleBlobby/github_trending_bot/parser"
)

func main() {
	curler.WriteHTMLToFile(curler.GetHTMLFromURL("https://github.com/trending"))
	parser.FindDailyTrendingURLS()
	parser.DeleteDirtyHTMLFile()

	// at this point:
	// 1. we have curled the entire trending page dirtyhtml into a file
	// 2. we parsed that file for the links to the trending repos
	// 3. we wrote those links to a new file (todaysDate-barelinks)
	// 4. we deleted the dirtyhtml file

	// next steps:
	// 1. curl each page from the barelinks file
	// 2. parse each page for the project title/description
	// 3. write the project title/description AND barelink to new file
	// 4. delete barelinks file

}
