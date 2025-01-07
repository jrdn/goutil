package node

import (
	"github.com/denisbrodbeck/machineid"
)

func ID() string {
	id, err := machineid.ProtectedID("plexus")
	if err == nil {
		return id
	}

	// fallback behavior, machineid failed me
	return getFallbackID()
}
