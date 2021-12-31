# GITHUB-TRENDING-BOT

Simple application to archive daily github trending repos, written in Go.

## Project Description

Uses Go to parse raw html from the github website to build a list of the daily trending repo names and descriptions, then writes that information to a file named with today's date [YYYY-MM-DD] in the archives folder.

## Purpose

Saves you time by automatically scraping the trending repos -- imagine all the extra productive tasks you can accomplish in a given day when you don't have to spend 3 minutes surfing the web!  Seriously.

Actually just a simple exercise used as an excuse to noodle around with the Go language.  The previous version of this project relied on shell scripting (curl) to grab the html content, but in this implementation I have removed the shell scripting entirely and used Go to make the http request, streamlining the process.

---
## Instructions

### Clone Repo

1. [Install Go Language](https://go.dev/doc/install)


2. From a Linux Terminal:

    ```git clone uncleBlobby@github.com/github-trending-bot```

3. ### Run 

    ```go run .```

4. *Go* be productive doing something else.

---


## TODO:

Prettify output in a markdown file

Catalogue output and make it searchable
