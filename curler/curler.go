package curler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetHTMLFromURL(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func WriteHTMLToFile(html string) {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")

	of, oferr := os.Create(todaysDate)
	if oferr != nil {
		log.Fatal(oferr)
	}

	defer of.Close()

	of.WriteString(html)

}
