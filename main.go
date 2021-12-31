package main

import GTB "github.com/uncleBlobby/github_trending_bot/parser"

func main() {
	GTB.FindDailyTrendingURLS()

	GTB.ParseProjectForDescription()
}
