package catalog

// Catalog represents a single campaign
type Catalog struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ProjectID      int    `json:"projectId"`
	ConversionRate int    `json:"conversionRate"`
}
