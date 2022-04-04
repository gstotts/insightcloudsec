package insightcloudsec

type Badge_Creation_Request struct {
	Resource_IDs []string `json:"target_resource_ids"`
	Badges       []Badge  `json:"badges"`
}

type Badge struct {
	// The key and value of a given badge for use with filters, insights, etc.
	Key   string `json:"key"`
	Value string `json:"value"`
}
