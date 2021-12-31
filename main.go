package main

import (
	curler "github.com/uncleBlobby/github_trending_bot/curler"
	parser "github.com/uncleBlobby/github_trending_bot/parser"
)

func main() {

	curler.WriteHTMLToFile(curler.GetHTMLFromURL("https://github.com/trending"))
	parser.FindDailyTrendingURLS()
	parser.DeleteDirtyHTMLFile()
	parser.FindProjectDirtyHTMLAndWriteOutputFile()

}
