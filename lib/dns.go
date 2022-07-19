package lib

import (
	"errors"
	"fmt"
	"log"
	"net"
)

var errDNSHostRequired = errors.New("discover-dns: Host is required")

type DNSProvider struct{}

func (d *DNSProvider) Addrs(args map[string]string, l *log.Logger) ([]string, error) {
	host := args["host"]

	if host == "" {
		return nil, errDNSHostRequired
	}

	ips, err := net.LookupIP(host)

	if err != nil {
		return nil, fmt.Errorf("discover-dns: Failed to lookup IPs: %w", err)
	}

	addrs := make([]string, len(ips))

	for _, ip := range ips {
		addrs = append(addrs, ip.String())
	}

	return addrs, nil
}

func (d *DNSProvider) Help() string {
	return `DNS:

	provider: "dns"
	host:     Host to look up.
`
}
