package env

import (
	"fmt"

	"github.com/google/uuid"
)

func NodeID() string {
	return fmt.Sprintf("%x", uuid.NodeID())
}
