package oeis

import (
	"kasperaldrin/oeis/pkg/models"
	"kasperaldrin/oeis/pkg/services"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// OEISClientMode is the mode of the OEIS client
type OEISClientMode string

var (
	ModeOnline  OEISClientMode = "online"
	ModeOffline OEISClientMode = "offline"
)

// OEISClientConfig is the configuration for the OEIS client
type OEISClientConfig struct {
	Mode        OEISClientMode // Wether the client should use the online or offline mode
	OfflinePath string         // The path to the offline dataset
}

// OEISClient is the interface for the OEIS client
type OEISClient interface {
	Search(query string) ([]*models.OEISSequence, error)
	Get(id string) (*models.OEISSequence, error)
}

// OnlineClient is a client that uses the online database
type OnlineClient struct {
	baseURL string
	http    *http.Client
}

// Search searches for sequences by name or description
func (c *OnlineClient) Search(query string) ([]*models.OEISSequence, error) {
	return nil, nil
}

// Get retrieves a sequence by its ID
func (c *OnlineClient) Get(id string) (*models.OEISSequence, error) {
	return nil, nil
}

// OfflineClient is a client that uses a local dataset
type OfflineClient struct {
	OfflinePath string // Path to the offline dataset
}

// Search searches for sequences by name or description
func (c *OfflineClient) Search(query string) ([]*models.OEISSequence, error) {
	return nil, nil
}

// Get retrieves a sequence by its ID
func (c *OfflineClient) Get(id string) (*models.OEISSequence, error) {

	// The folder is the first 4 characters of the ID
	folder := id[:4]

	filePath := filepath.Join(c.OfflinePath, "seq", folder, id+".seq")

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return services.ParseOEIS(f)
}

// NewClient creates a new OEIS client based on the configuration
// Either the online or offline mode can be used, but not both at the same time.
// If using offline mode, the OfflinePath must be set: https://oeis.org/wiki/Compressed_Versions_of_the_Database
// OBS: The online mode is not yet implemented.
func NewClient(config *OEISClientConfig) OEISClient {
	switch config.Mode {
	case ModeOnline:
		return &OnlineClient{baseURL: "https://oeis.org", http: &http.Client{Timeout: 10 * time.Second}}
	case ModeOffline:
		return &OfflineClient{OfflinePath: config.OfflinePath}
	default:
		panic("invalid mode")
	}
}
