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
