package node

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func getFallbackID() (string, bool) {
	// MachineGuid is installed by Windows and survives reboots and NIC changes.
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return "", false
	}
	defer key.Close()
	guid, _, err := key.GetStringValue("MachineGuid")
	if err != nil || strings.TrimSpace(guid) == "" {
		return "", false
	}
	return fmt.Sprintf("%x", sha256.Sum224([]byte(guid))), true
}
