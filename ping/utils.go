package ping

import (
	"net"
	"time"
)

func timeIt(f func() interface{}) (int64, interface{}) {
	startAt := time.Now()
	res := f()
	endAt := time.Now()
	return endAt.UnixNano() - startAt.UnixNano(), res
}

// UseCustomeDNS will set the dns to default DNS resolver for global
func UseCustomeDNS(dns []string) {
	resolver := net.Resolver{
		PreferGo: true,
	}
	net.DefaultResolver = &resolver
}
