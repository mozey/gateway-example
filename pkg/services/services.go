package services

// Options for services
type Options struct {
	Debug      string
	Region     string
	AwsProfile string
}

// Services to embed in internal handler
type Services struct {
	Options *Options
}

// NewServices initialises services
func NewServices(o *Options) (s *Services) {
	s = &Services{}
	s.Options = o
	return s
}

// Cleanup function must be called before the application exits
func (s *Services) Cleanup() {
	// Close db connections etc
}
