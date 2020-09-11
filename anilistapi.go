package anilistapi

import (
	"fmt"

	"github.com/animenotifier/anilist"
	"golang.org/x/oauth2"
)

/*
* Setup the client
 */

// Anilistwrapper interface that exposes the public functions in a orderly fashion
type Anilistapi interface {
	GetUserID(userName string) int
	GetUserInformation() (interface{}, error)
	GetUsersAnimeLists(userID int) (*anilist.AnimeList, error)
	GetUsersAnimeListContent(userID int) (interface{}, error)
}

// InitClient provides a mechanism to
func InitClient(token *oauth2.Token) Anilistapi {
	httpClient = HTTPClient{token}
	return Client{
		token,
	}
}

// GetUserID retrieves the anilist userID belonging to an anilist username
func (c Client) GetUserID(userName string) int {
	anilistUser, _ := anilist.GetUser(userName)

	return anilistUser.ID
}

// GetUsersAnimeLists retrieves the lists with content stored in the anilist profile
func (c Client) GetUsersAnimeLists(userID int) (*anilist.AnimeList, error) {
	anilistAnimeList, err := anilist.GetAnimeList(userID)
	if err != nil {
		return nil, err
	}

	return anilistAnimeList, nil
}

// GetUsersAnimeListContent retrieves the content of an list in anilist
func (c Client) GetUsersAnimeListContent(userID int) (interface{}, error) {
	graphqlQuery := map[string]string{"query": `{
		MediaListCollection(userId:433795, type:ANIME) {
			lists {
				name
				entries {
					mediaId
					status
					progress
				}
		  	}
		}
	}`,
	}
	response, err := c.httpClient(graphqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}

// GetMediaDetails retrieves details about media from the anilist API
func (c Client) GetMediaDetails(mediaID int) (interface{}, error) {
	graphqlQuery := map[string]string{"query": `{
		media {
			id
			idMal
			title {
				romaji
				english
				native
				userPreferred
			}
		}
	}`,
	}
	response, err := c.httpClient(graphqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}

// GetUserInformation retrieves information about the current authenticated user from the API
func (c Client) GetUserInformation() (interface{}, error) {
	graphqlQuery := map[string]string{"query": `{
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
	}`,
	}
	response, err := c.httpClient(graphqlQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}
