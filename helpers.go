package insightcloudsec

// HELPER FUNCTIONS
///////////////////////////////////////////

func isInSlice(v string, s []string) bool {
	isInSlice := false
	for _, item := range s {
		if item == v {
			isInSlice = true
		}
	}
	return isInSlice
}
