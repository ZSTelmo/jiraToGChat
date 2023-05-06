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
		//fmt.Println(responseBody)

		dataForHook := map[string]interface{}{}
		err = json.Unmarshal(bodyBytes, &dataForHook)
		if err != nil {
			panic(err)
		}

		// issue := dataForHook["issue"].(map[string]interface{})
		// user := dataForHook["user"].(map[string]interface{})
		//changelog := dataForHook["changelog"].(map[string]interface{})

		//issueId := issue["id"]
		// issueKey := issue["key"]
		// issueSelf := issue["self"]

		// userDisplayName := user["displayName"]

		//issueFields := issue["fields"].(map[string]interface{})

		//summary := issueFields["summary"]

		// templatePostBody := heredoc.Doc(`{"cardsV2":[{"cardId":"unique-card-id","card":{"header":{"title":"%v","subtitle": "%v","imageUrl":"https://logowik.com/content/uploads/images/jira3124.jpg","imageType": "CIRCLE","imageAltText": "Avatar for Jira"},"sections":[{"header": "%v","collapsible": false,"uncollapsibleWidgetsCount":1,"widgets":[{"decoratedText":{"startIcon":{"knownIcon":"STAR"},"text":"<a href=\"%v">%v/a>"}}]}]}}]}`)

		// postBody := fmt.Sprintf(templatePostBody,
		// 	issueKey,
		// 	summary,
		// 	userDisplayName,
		// 	issueSelf,
		// 	issueKey,
		// )

		postBody := heredoc.Doc(`{
    "cardsV2": [
        {
            "cardId": "unique-card-id",
            "card": {
                "header": {
                    "title": "JRA-20002",
                    "subtitle": "99291",
                    "imageUrl": "https://logowik.com/content/uploads/images/jira3124.jpg",
                    "imageType": "CIRCLE",
                    "imageAltText": "Avatar for Jira"
                },
                "sections": [
                    {
                        "header": "Bryan Rollins [Atlassian]",
                        "collapsible": false,
                        "uncollapsibleWidgetsCount": 1,
                        "widgets": [
                            {
                                "decoratedText": {
                                    "startIcon": {
                                        "knownIcon": "STAR"
                                    },
                                    "text": "<a href=\"https://your-domain.atlassian.net/rest/api/2/issue/99291\">JRA-20002</a>"
                                }
                            }
                        ]
                    }
                ]
            }
        }
    ]
}`)

		payload := strings.NewReader(postBody)
		//jsonValue, _ := json.Marshal(postBody)

		gchatWebHook := utils.GetENVasString("GCHAT_WEBHOOK")

		//	resp, err := http.Post(gchatWebHook, "application/json", bytes.NewBuffer(jsonValue))
		resp, err := http.Post(gchatWebHook, "application/json", payload)

		fmt.Println("Status:", resp.StatusCode)

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
