package platform

// Functions struct to create functions
type Functions struct {
}

// Create functions
func (f Functions) Create(application Application) error {
	println("Created Function for Functions")
	return nil
}
