package statute

import (
	"github.com/bepass-org/ipscanner/internal/dialer"
	"net/http"
	"time"
)

var HTTPPing = 1 << 0
var TLSPing = 1 << 1
var TCPPing = 1 << 2
var QUICPing = 1 << 3

type ScannerOptions struct {
	UseIPv4            bool
	UseIPv6            bool
	CidrList           []string // CIDR ranges to scan
	SelectedOps        int
	Logger             Logger
	Timeout            time.Duration
	InsecureSkipVerify bool
	Dialer             *dialer.AppDialer
	TLSDialer          *dialer.AppTLSDialer
	RawDialerFunc      dialer.TDialerFunc
	TLSDialerFunc      dialer.TDialerFunc
	HttpClient         *http.Client
	HTTPPath           string
	Hostname           string
	Port               uint16
	IPBasketSize       int
}