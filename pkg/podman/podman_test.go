package podman

import (
	"testing"
)

func TestClient_Ping(t *testing.T) {
	client := NewClient("/var/run/podman/podman.sock", 2)
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}
