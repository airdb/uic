package valueobject

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Username string `json:"userName,omitempty"`
	Token    string `json:"token,omitempty"`
	Redirect string `json:"redirect,omitempty"`
}

type OauthConfig struct {
	ID          uint   `json:"id"`
	Provider    string `json:"provider,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	CientSecret string `json:"cient_secret,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	State       string `json:"state,omitempty"`
}

// Ticker is basic data type of trade
type Ticker struct {
	Sell      string
	Buy       string
	High      string
	Low       string
	Last      string
	Vol       string
	Timestamp string
}

// Payment is payment info of domain.Order
type Payment struct {
}
