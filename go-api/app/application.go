package app

type StatsResponse struct {
	Status string `json:"status"`
}

// TODO: Please define business logic in this folder or create a new folder for it.
// The following codes are just an example.

// CheckSystemStatus is used to check the system status
func CheckSystemStatus() StatsResponse {
	// TODO: Please define how to check the system status
	return StatsResponse{Status: "OK"}
}


