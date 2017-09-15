package main

import (
	"github.com/ottojo/toolboxPSDVoteMonitor/psdSiteParser"
	"net/http"
	"log"
	"io/ioutil"
	"sort"
	"fmt"
)

func main() {
	response, err := http.Get("https://www.psd-miteinander-leben.de/profile/")
	panicErr(err)
	body, err := ioutil.ReadAll(response.Body)
	panicErr(err)
	profiles := psdSiteParser.Parse(body)
	sort.Sort(profiles)

	var toolboxIndex int
	for i, profile := range profiles {
		if profile.Id == "toolbox-bodensee-e-v" {
			toolboxIndex = i
			break
		}
	}

	if toolboxIndex != len(profiles)-1 {
		fmt.Println("Toolbox needs ", profiles[len(profiles)-1].Votes-profiles[toolboxIndex].Votes, " more votes to be Nr. 1 in the voting!")
		fmt.Println("Toolbox needs ", profiles[toolboxIndex+1].Votes-profiles[toolboxIndex].Votes, "more votes to climb one rank!")
	} else {
		fmt.Println("Toolbox is Nr. 1")
	}
}

func panicErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
