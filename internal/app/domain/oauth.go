package domain

type OauthConfig struct {
	ID          uint
	Provider    string
	ClientID    string
	CientSecret string
	RedirectURL string
	State       string
}
