package store

// Store represents a single campaign
type Store struct {
	ID                     int      `json:"id"`
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	VendorID               int      `json:"vendorId"`
	ProjectConfigurationID int      `json:"projectConfigurationId"`
	Visible                int      `json:"visible"`
	Parameters             []string `json:"parameters"`
}
