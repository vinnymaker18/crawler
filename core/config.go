package core 

// Crawler program configuration
type Config struct {

	// System directory to store crawled data.
	dataDir string
}

// Error in crawler configuration file.
type ConfigError struct {
	reason string
}

func (err *ConfigError) Error() string {
	return err.reason
}

// Config file parser.
func ParseConfig(configFile string) (*Config, *ConfigError) {
	// TODO - Change this to a proper implementation.
	err := &ConfigError{
		reason: "Bad config file",
	}

	return nil, err
}
