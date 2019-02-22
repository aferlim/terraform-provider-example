package participant

// Participant represents a single campaign
type Participant struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Login                  string `json:"login"`
	Password               string `json:"password"`
	Email                  string `json:"email"`
	CustomerID             int    `json:"costumerId"`
	ProjectID              int    `json:"projectId"`
	ProjectConfigurationID int    `json:"projectConfigurationId"`
	Active                 bool   `json:"active"`
}
