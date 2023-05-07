package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"jiraToGChat.com/utils"
)

func main() {
	utils.LoadEnvFile()

	// webHook routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		//responseBody := string(bodyBytes)

		dataForHook := map[string]interface{}{}
		err = json.Unmarshal(bodyBytes, &dataForHook)
		if err != nil {
			panic(err)
		}

		issue := dataForHook["issue"].(map[string]interface{})
		user := dataForHook["user"].(map[string]interface{})
		//changelog := dataForHook["changelog"].(map[string]interface{})

		//issueId := issue["id"]
		issueKey := issue["key"]
		issueSelf := issue["self"]

		userDisplayName := user["displayName"]

		issueFields := issue["fields"].(map[string]interface{})

		summary := issueFields["summary"]

		templatePostBody := heredoc.Doc(`{"cardsV2":[{"cardId":"unique-card-id","card":{"header":{"title":"%s","subtitle": "%s","imageUrl":"https://logowik.com/content/uploads/images/jira3124.jpg","imageType": "CIRCLE","imageAltText": "Avatar for Jira"},"sections":[{"header": "%s","collapsible": false,"uncollapsibleWidgetsCount":1,"widgets":[{"decoratedText":{"startIcon":{"knownIcon":"STAR"},"text":"<a href=\"%s\">%s</a>"}}]}]}}]}`)

		postBody := fmt.Sprintf(templatePostBody,
			issueKey,
			summary,
			userDisplayName,
			issueSelf,
			issueKey,
		)

		payload := strings.NewReader(postBody)

		gchatWebHook := utils.GetENVasString("GCHAT_WEBHOOK")

		resp, err := http.Post(gchatWebHook, "application/json", payload)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Printf("Status: %v \n", resp.StatusCode)

	})

	portFromENV := utils.GetENVasString("PORT")
	if len(portFromENV) == 0 {
		portFromENV = "9000"
	}

	port := ":" + portFromENV
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
