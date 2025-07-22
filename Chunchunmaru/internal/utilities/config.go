package utilities

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

// Duration Have to make a custom duration struct to allow JSON marshaling :(
type Duration time.Duration

// UnmarshalJSON Unmarshal the JSON
func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(value * float64(time.Second))
		return nil
	case string:
		parsed, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(parsed)
		return nil
	default:
		return errors.New("invalid duration format: must be a string (e.g., 5s, 5000ms, 5m) or a number of seconds (e.g., 5.0 for 5 seconds)")
	}
}

// MarshalJSON Marshal the json
func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(*d).String())
}

type Config struct {
	Port          int      `json:"port"`
	Delay         Duration `json:"delay"`
	HostName      string   `json:"hostname"`
	PathWhitelist []string `json:"path_whitelist"`
	MinSubpaths   int      `json:"min_subpaths"`
	MaxSubpaths   int      `json:"max_subpaths"`
}

type ConfigManager struct {
	config Config
	mu     sync.RWMutex
}

// NewConfigManager Returns a new config manager
func NewConfigManager(initialConfig Config) *ConfigManager {
	return &ConfigManager{
		config: initialConfig,
	}
}

var AppConfig = NewConfigManager(Config{
	Port:          8080,
	Delay:         Duration(5000 * time.Millisecond),
	HostName:      "http://localhost:8080",
	PathWhitelist: []string{},
	MinSubpaths:   1,
	MaxSubpaths:   5,
})

// GetConfig Gets the config
func (cm *ConfigManager) GetConfig() Config {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.config
}

// SetConfig Sets the config
func (cm *ConfigManager) SetConfig(config Config) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.config = config
	log.Printf("Configuration updated to: %+v\n", config)
}

// ConfigSetAPI Handler to set the config via the API
func (cm *ConfigManager) ConfigSetAPI(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON config
	var newConfig Config
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure valid config
	if newConfig.Port < 0 || newConfig.Port > 65535 {
		http.Error(w, "Port must be between 0 and 65535.", http.StatusBadRequest)
		return
	}
	if newConfig.Delay < 0 {
		http.Error(w, "Delay must be greater or equal to 0.", http.StatusBadRequest)
		return
	}

	cm.SetConfig(newConfig)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Configuration updated successfully.",
	})
}
