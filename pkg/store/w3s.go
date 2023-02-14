package store

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
)

// W3Store is an interface for web3.storage
type W3Store interface {
	SaveFile(file fs.File) (cid.Cid, error)
	GetFile(cid cid.Cid) (fs.File, fs.FS, error)
}

// W3S is a web3.storage client
type W3S struct {
	client w3s.Client
}

// GetClient returns an authenticated web3.storage client
func GetClient() (W3S, error) {
	W3Client := W3S{}

	// Get the API token from the environment
	token, ok := os.LookupEnv("WEB3_STORAGE_TOKEN")
	if !ok {
		_, err := fmt.Fprintln(os.Stderr, "No API token - set the WEB3_STORAGE_TOKEN environment var and try again.")
		if err != nil {
			return W3Client, err
		}
	}

	// Create a new web3.storage client using the token
	client, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		return W3Client, err
	}

	W3Client.client = client
	return W3Client, nil
}

// SaveFile saves a file to web3.storage
func (W3Client W3S) SaveFile(file fs.File) (cid.Cid, error) {
	return W3Client.client.Put(context.Background(), file)
}

// GetFile gets a file from web3.storage
func (W3Client W3S) GetFile(cid cid.Cid) (fs.File, fs.FS, error) {
	resp, err := W3Client.client.Get(context.Background(), cid)
	if err != nil {
		return nil, nil, err
	}

	// Handle response
	if resp.StatusCode != http.StatusOK {
		return nil, nil,
			fmt.Errorf("request for %s was unsuccessful: [%d]: %s",
				cid, resp.StatusCode, resp.Status)
	}

	return resp.Files()
}
