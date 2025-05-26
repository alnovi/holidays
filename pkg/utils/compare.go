package utils

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

var ErrHostNotCompare = errors.New("host not compare")

func CompareHosts(links ...string) error {
	if len(links) < 2 { //nolint:mnd
		return ErrHostNotCompare
	}

	firstHost := ""

	for i, item := range links {
		var host string

		if !strings.HasPrefix(item, "http") {
			item = "https://" + strings.Trim(item, "/")
		}

		link, err := url.Parse(item)
		if err != nil {
			return err
		}

		if host, _, err = net.SplitHostPort(link.Host); err != nil {
			host = link.Host
		}

		if i == 0 {
			firstHost = host
			continue
		}

		if firstHost != host {
			return ErrHostNotCompare
		}
	}

	return nil
}
