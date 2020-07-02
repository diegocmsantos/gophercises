package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"gophercises/choose_your_adventure/src/models"
	"gophercises/choose_your_adventure/src/server"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var startHTTPServerFlag bool
var arcs map[string]models.Arc

func main() {

	parseFlags()

	byteValue, err := ioutil.ReadFile("assets/gopher.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &arcs)

	if startHTTPServerFlag {
		server.StartHTTPServer(arcs)
	} else {
		startCommandLine()
	}

}

func parseFlags() {
	flag.BoolVar(&startHTTPServerFlag, "server", false, "If it's true a HTTP server will be started. Otherwise will be a command line program.")
	flag.Parse()
}

func startCommandLine() {
	currentArc := arcs["intro"]
	for len(currentArc.Options) != 0 {
		var optionsTxt []string
		var optionsMap = make(map[int]string, len(currentArc.Options))
		for i, optTxt := range currentArc.Options {
			optionsTxt = append(optionsTxt, fmt.Sprintf("%d) %s", i+1, optTxt.Text))
			optionsMap[i] = optTxt.Arc
		}
		response := askQuestion(strings.Join(currentArc.Story, "\n") + "\n\nOptions:\n\n" + strings.Join(optionsTxt, "\n") + "\n\nType the number of your choice: ")
		intResponse, _ := strconv.Atoi(response)
		currentArc = arcs[optionsMap[intResponse-1]]
	}
	fmt.Println(strings.Join(currentArc.Story, "\n"))
}

func askQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question)

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	return strings.Replace(text, "\n", "", -1)
}

func takeNextStory(title string) models.Arc {
	return arcs[title]
}
