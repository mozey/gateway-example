package hutil

// Config for shared handler
type Config struct {
	Debug          string
	Region         string
	AwsProfile     string
}

// Handler with shared fields.
// Embed this type in internal handlers
type Handler struct {
	Config       *Config
}

// NewHandler creates a new shared handler instance
func NewHandler(config *Config) (h *Handler) {
	h = &Handler{}
	h.Config = config
	//h.SetupLogging(e)
	return h
}

// Cleanup function must be called before the application exits
func (h *Handler) Cleanup() {
	// Close db connections etc
}
