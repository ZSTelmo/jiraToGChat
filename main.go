package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

		responseBody := string(bodyBytes)
		fmt.Println(responseBody)

		var dataForHook map[string]interface{}
		err = json.Unmarshal([]byte(responseBody), &dataForHook)
		if err != nil {
			panic(err)
		}

		fmt.Println(dataForHook["issue"])

		postBody := `{
    "cardsV2": [
        {
            "cardId": "unique-card-id",
            "card": {
                "header": {
                    "title": "Jira",
                    "subtitle": "Jira BOT",
                    "imageUrl": "https://developers.google.com/chat/images/quickstart-app-avatar.png",
                    "imageType": "CIRCLE",
                    "imageAltText": "Avatar for Jira"
                },
                "sections": [
                    {
                        "header": "Contact Info",
                        "collapsible": true,
                        "uncollapsibleWidgetsCount": 1,
                        "widgets": [
                            {
                                "decoratedText": {
                                    "startIcon": {
                                        "knownIcon": "EMAIL"
                                    },
                                    "text": "sasha@example.com"
                                }
                            },
                            {
                                "decoratedText": {
                                    "startIcon": {
                                        "knownIcon": "PERSON"
                                    },
                                    "text": "<font color=\"#80e27e\">Online</font>"
                                }
                            },
                            {
                                "decoratedText": {
                                    "startIcon": {
                                        "knownIcon": "PHONE"
                                    },
                                    "text": "+1 (555) 555-1234"
                                }
                            },
                            {
                                "buttonList": {
                                    "buttons": [
                                        {
                                            "text": "Share",
                                            "onClick": {
                                                "openLink": {
                                                    "url": "https://example.com/share"
                                                }
                                            }
                                        },
                                        {
                                            "text": "Edit",
                                            "onClick": {
                                                "action": {
                                                    "function": "goToView",
                                                    "parameters": [
                                                        {
                                                            "key": "viewType",
                                                            "value": "EDIT"
                                                        }
                                                    ]
                                                }
                                            }
                                        }
                                    ]
                                }
                            }
                        ]
                    }
                ]
            }
        }
    ]
}`

		var dataToPost map[string]interface{}
		err = json.Unmarshal([]byte(postBody), &dataToPost)

		// 	values := map[string]string{

		// 		"username": username, "password": password

		// }

		jsonValue, _ := json.Marshal(dataToPost)

		gchatWebHook := utils.GetENVasString("GCHAT_WEBHOOK")

		resp, err := http.Post(gchatWebHook, "application/json", bytes.NewBuffer(jsonValue))

		fmt.Println(w, resp)

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
