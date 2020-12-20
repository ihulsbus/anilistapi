package anilistapi

import "golang.org/x/oauth2"

/*
* Structs to initialize the client
 */

// Client declares the content of the Client
type Client struct {
	token *oauth2.Token
}

// HTTPClient provides an initializer for the HTTP Client
type HTTPClient struct {
	token *oauth2.Token
}

/*
* Function related structs
 */

//UsersAnimeLists defines the response structure when retrieving the users's anilist lists
type UsersAnimeLists struct {
	Data struct {
		MediaListCollection struct {
			Lists []struct {
				Name string `json:"name"`
			} `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}

//AnimeListContent defines the response structure when retrieving the content of a list
type AnimeListContent struct {
	Data struct {
		MediaListCollection struct {
			Lists []struct {
				Name    string `json:"name"`
				Entries []struct {
					MediaID  int    `json:"mediaId"`
					Status   string `json:"status"`
					Progress int    `json:"progress"`
				} `json:"entries"`
			} `json:"lists"`
		} `json:"MediaListCollection"`
	} `json:"data"`
}
