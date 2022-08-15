package handle

// helper method
func newServices() *Services {
	return &Services{
		User: &userServer{},
	}
}
