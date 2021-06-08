package valueobject

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token,omitempty"`
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
