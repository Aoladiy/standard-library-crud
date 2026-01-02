package item

type Item struct {
	Message *string `json:"message,omitempty"`
}

var items []Item
