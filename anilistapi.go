package anilistapi

import (
	"context"
	"fmt"
	"log"

	"github.com/animenotifier/anilist"
	"github.com/machinebox/graphql"
	"golang.org/x/oauth2"
)

// Anilistwrapper interface that exposes the public functions in a orderly fashion
type Anilistwrapper interface {
	AnilistClient(query string) (interface{}, error)
	GetUserID(userName string) int
	GetUserInformation() (interface{}, error)
	GetUsersAnimeLists(userID int) (interface{}, error)
	GetUsersAnimeListContent(userID int) (interface{}, error)
}

// Client declares the content of the Client
type Client struct {
	token *oauth2.Token
}

// InitClient provides a mechanism to
func InitClient(token *oauth2.Token) Anilistwrapper {
	return Client{
		token,
	}
}

// AnilistClient provides an interface with the graphQL API of Anilist
func (c Client) AnilistClient(query string) (interface{}, error) {
	// create a client (safe to share across requests)
	graphqlClient := graphql.NewClient("https://graphql.anilist.co")

	// Make a request
	graphqlRequest := graphql.NewRequest(query)

	// set header fields
	graphqlRequest.Header.Set("Cache-Control", "no-cache")

	// Set authorization bearer
	graphqlRequest.Header.Add("Authorization", c.token.AccessToken)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData interface{}
	if err := graphqlClient.Run(ctx, graphqlRequest, &respData); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return respData, nil
}

// GetUserID retrieves the anilist userID belonging to an anilist username
func (c Client) GetUserID(userName string) int {
	anilistUser, _ := anilist.GetUser(userName)

	return anilistUser.ID
}

// // GetUsersAnimeLists retrieves the lists with content stored in the anilist profile
// func GetUsersAnimeLists(userID int) []struct {
// 	Name    string                   `json:"name"`
// 	Entries []*anilist.AnimeListItem `json:"entries"`
// } {
// 	anilistAnimeList, _ := anilist.GetAnimeList(userID)

// 	return anilistAnimeList.Lists
// }

// GetUsersAnimeLists retrieves information about the current user from the API
func (c Client) GetUsersAnimeLists(userID int) (interface{}, error) {
	graphqlQuery := `query {
		MediaListCollection(userId:` + string(userID) + `, type: ANIME) {
			lists {
				name
				isCustomList
				isSplitCompletedList
				status
			}
		}
	}`
	response, err := c.AnilistClient(graphqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}

// GetUsersAnimeListContent retrieves the content of an list in anilist
func (c Client) GetUsersAnimeListContent(userID int) (interface{}, error) {
	graphqlQuery := `query {
		MediaListCollection(userId:` + string(userID) + `, type:ANIME) {
			lists {
				name
				entries {
					mediaId
					status
					progress
				}
		  	}
		}
	}`
	response, err := c.AnilistClient(graphqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}

// GetMediaDetails retrieves details about media from the anilist API
func (c Client) GetMediaDetails(mediaID int) (interface{}, error) {
	graphqlQuery := `media {
		id
		idMal
		title {
			romaji
			english
			native
			userPreferred
		}
	}`
	response, err := c.AnilistClient(graphqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}

// GetUserInformation retrieves information about the current authenticated user from the API
func (c Client) GetUserInformation() (interface{}, error) {
	grapqlQuery := `query {
		Viewer {
			id
			name
			avatar {
				large
			}
			options {
				titleLanguage
				displayAdultContent
			}
			mediaListOptions {
			scoreFormat
			animeList {
				splitCompletedSectionByFormat
				customLists
				advancedScoringEnabled
			}
			}
		}
	}`
	response, err := c.AnilistClient(grapqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}
