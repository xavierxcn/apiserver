package podman

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// Client is the main struct of podman
type Client struct {
	podmanPath string // path to podman
	sockPath   string // socket path
	preCmd     string // pre-command
	url        string
	httpClient *http.Client
}

var (
	preCmd = "podman system service -t %d &"
	url    = "http://d/v4.0.0"
)

// NewClient creates a new Client
func NewClient(sock string, timeout int64) *Client {
	if sock == "" {
		sock = "/var/run/podman/podman.sock"
	}
	if timeout == 0 {
		timeout = 10
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", sock)
			},
		},
	}
	return &Client{
		sockPath:   sock,
		preCmd:     fmt.Sprintf(preCmd, timeout),
		httpClient: httpClient,
		url:        url,
	}
}

// Ping performs a ping request
func (c *Client) Ping() error {
	path := "/_ping"
	request, err := http.NewRequest("HEAD", c.url+path, nil)
	if err != nil {
		return err
	}

	response, err := c.DoHttp(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("ping failed: %s", response.Status)
	}

	return nil
}

func (c *Client) systemService() error {
	_, err := os.Stat(c.sockPath)
	if err != nil {
		cmd := exec.Command("/bin/sh", "-c", c.preCmd)
		err := cmd.Run()
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 500)

		return nil
	}

	return nil
}

// DoHttp performs an HTTP request
func (c *Client) DoHttp(request *http.Request) (*http.Response, error) {
	err := c.systemService()
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(request)
}

// =================== 结构体 ===================

// Image is the struct of container image
type Image struct {
	ID          string      `json:"Id"`
	ParentID    string      `json:"ParentId"`
	RepoTags    []string    `json:"RepoTags"`
	RepoDigests []string    `json:"RepoDigests"`
	Created     int         `json:"Created"`
	Size        int         `json:"Size"`
	SharedSize  int         `json:"SharedSize"`
	VirtualSize int         `json:"VirtualSize"`
	Labels      interface{} `json:"Labels"`
	Containers  int         `json:"Containers"`
	Names       []string    `json:"Names"`
	Digest      string      `json:"Digest"`
	History     []string    `json:"History"`
}
