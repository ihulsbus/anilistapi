package anilistapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient HTTPClient

// AnilistClient provides an interface with the graphQL API of Anilist
func (h HTTPClient) httpClient(query interface{}) ([]byte, error) {
	marshaledQuery, _ := json.Marshal(query)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://graphql.anilist.co", bytes.NewBuffer(marshaledQuery))
	if err != nil {
		log.Print("Failed to create request", err)
		return nil, err
	}

	// This one line implements the authentication required for the task.
	req.Header.Add("Bearer", h.token.AccessToken)

	// Set that we accept Json as response
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Make request and show output.
	resp, err := client.Do(req)
	if err != nil {
		log.Print("Failed to send request")
		return nil, err
	}

	// This check is probably enough since the auditor is read-only
	if resp.StatusCode != 200 {
		log.Print("Encountered server error ", resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	return bodyText, nil
}

// func (c Client) AnilistClient(query string) (interface{}, error) {
// 	// create a client (safe to share across requests)
// 	graphqlClient := graphql.NewClient("https://graphql.anilist.co")

// 	// Make a request
// 	graphqlRequest := graphql.NewRequest(query)

// 	// set header fields
// 	graphqlRequest.Header.Set("Cache-Control", "no-cache")

// 	// Set authorization bearer
// 	graphqlRequest.Header.Add("Authorization", c.token.AccessToken)

// 	// define a Context for the request
// 	ctx := context.Background()

// 	// run it and capture the response
// 	var respData interface{}
// 	if err := graphqlClient.Run(ctx, graphqlRequest, &respData); err != nil {
// 		log.Fatal(err)
// 		return respData, err
// 	}
// 	return respData, nil
// }

//  Basic function to perform an http GET to Crowd and return the data.
func (c Client) httpClient(query interface{}) ([]byte, error) {
	return httpClient.httpClient(query)
}
