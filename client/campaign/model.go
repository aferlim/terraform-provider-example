package campaign

// Campaign represents a single campaign
type Campaign struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ClientID       int    `json:"clientId"`
	ExternalPoints int    `json:"externalPoints"`
}
