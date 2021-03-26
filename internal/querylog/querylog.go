package querylog

import (
	"net"
	"net/http"
	"time"

	"github.com/AdguardTeam/AdGuardHome/internal/dnsfilter"
	"github.com/miekg/dns"
)

// QueryLog - main interface
type QueryLog interface {
	Start()

	// Close query log object
	Close()

	// Add a log entry
	Add(params AddParams)

	// WriteDiskConfig - write configuration
	WriteDiskConfig(c *Config)
}

// Config - configuration object
type Config struct {
	// ConfigModified is called when the configuration is changed, for
	// example by HTTP requests.
	ConfigModified func()

	// HTTPRegister registers an HTTP handler.
	HTTPRegister func(string, string, func(http.ResponseWriter, *http.Request))

	// FindClient returns client information by their IDs.
	FindClient func(ids []string) (c *Client, err error)

	// BaseDir is the base directory for log files.
	BaseDir string

	// Interval is the interval for log rotation, in days.  After that
	// period, the old log file will be renamed, NOT deleted, so the actual
	// log retention time is twice the interval.
	Interval uint32

	// MemSize is the number of entries kept in a memory buffer before they
	// are flushed to disk.
	MemSize uint32

	// Enabled tells if the query log is enabled.
	Enabled bool

	// FileEnabled tells if the query log writes logs to files.
	FileEnabled bool

	// AnonymizeClientIP tells if the query log should anonymize clients' IP
	// addresses.
	AnonymizeClientIP bool
}

// AddParams - parameters for Add()
type AddParams struct {
	Question    *dns.Msg
	Answer      *dns.Msg          // The response we sent to the client (optional)
	OrigAnswer  *dns.Msg          // The response from an upstream server (optional)
	Result      *dnsfilter.Result // Filtering result (optional)
	Elapsed     time.Duration     // Time spent for processing the request
	ClientID    string
	ClientIP    net.IP
	Upstream    string // Upstream server URL
	ClientProto ClientProto
}

// New - create a new instance of the query log
func New(conf Config) QueryLog {
	return newQueryLog(conf)
}
