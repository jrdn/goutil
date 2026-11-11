package node

import (
	"crypto/sha256"
	"fmt"
	"os/exec"
	"strings"
)

func getFallbackID() (string, bool) {
	// IOPlatformUUID is a stable hardware UUID exposed by macOS. Hash it so
	// callers do not receive the raw platform identifier.
	out, err := exec.Command("/usr/sbin/ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
	if err != nil {
		return "", false
	}
	const key = "IOPlatformUUID"
	for _, line := range strings.Split(string(out), "\n") {
		if !strings.Contains(line, key) {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		uuid := strings.Trim(strings.TrimSpace(parts[1]), `"`)
		if uuid == "" {
			continue
		}
		return fmt.Sprintf("%x", sha256.Sum224([]byte(uuid))), true
	}
	return "", false
}
