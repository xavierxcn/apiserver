package podman

import (
	"testing"
)

func TestClient_ImageList(t *testing.T) {
	client := NewClient("/var/run/podman/podman.sock", 1)
	req, err := client.ImageList()
	if err != nil {
		t.Error(err)
	}

	t.Log(req)
}

func TestClient_ImagePull(t *testing.T) {
	client := NewClient("/var/run/podman/podman.sock", 1)
	err := client.ImagePull("docker.io/library/alpine")
	if err != nil {
		t.Error(err)
	}
}

func TestClient_ImagePush(t *testing.T) {
	client := NewClient("/var/run/podman/podman.sock", 1)
	err := client.ImagePush("172.16.8.10:5000/alpine:latest")
	if err != nil {
		t.Error(err)
	}
}

func TestClient_ImageBuild(t *testing.T) {
	client := NewClient("/var/run/podman/podman.sock", 1)
	err := client.ImageBuild("/root/code/api/sources-6SRUsZEorc/", "172.16.8.10:5000/api:latest")
	if err != nil {
		t.Error(err)
	}
}
