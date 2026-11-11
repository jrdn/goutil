package node

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"sort"
)

func fallbackID(systemd func() (string, bool), mac func() (string, bool)) (string, bool) {
	if id, ok := systemd(); ok {
		return id, true
	}
	if id, ok := mac(); ok {
		return id, true
	}
	return "", false
}

func getFallbackID() (string, bool) {
	return fallbackID(systemd, mac)
}

func systemd() (string, bool) {
	data, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		return "", false
	}
	return string(data), true
}

func mac() (string, bool) {
	ifs, err := net.Interfaces()
	if err != nil {
		return "", false
	}

	var data []string

	for _, iface := range ifs {
		addr := iface.HardwareAddr.String()
		if addr != "" {
			data = append(data, addr)
		}
	}
	if len(data) == 0 {
		return "", false
	}

	sort.Strings(data)
	buf := bytes.Buffer{}
	for _, x := range data {
		_, err := buf.WriteString(x)
		if err != nil {
			return "", false
		}
	}

	return fmt.Sprintf("%x", sha256.Sum224(buf.Bytes())), true
}
