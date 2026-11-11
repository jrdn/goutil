package node

import (
	"errors"
	"fmt"

	"github.com/denisbrodbeck/machineid"
)

var errNoFallbackID = errors.New("unable to determine node ID: no machine identity source available")

// ID returns a stable ID for this machine. It preserves the historical API by
// returning an empty string if neither the protected machine ID nor the
// platform fallback is available. Use IDWithError when the caller needs to
// distinguish that failure from a valid ID.
func ID() string {
	id, _ := IDWithError()
	return id
}

// IDWithError returns a stable, application-scoped ID for this machine.
// Platform-specific fallbacks are used when the OS machine ID cannot be read.
func IDWithError() (string, error) {
	id, err := machineid.ProtectedID("plexus")
	if err == nil {
		return id, nil
	}

	fallback, ok := getFallbackID()
	if !ok {
		return "", fmt.Errorf("machineid: %v; fallback: %w", err, errNoFallbackID)
	}
	return fallback, nil
}
