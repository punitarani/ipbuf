package test

import (
	"io"
	"io/fs"
	"os"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/punitarani/ipbuf/pkg/store"
)

// ipbuf.txt CID from web3.storage
const ipbufTxt = "bafybeic7kvcqfr5bn3h5ozs4ulixju3m7trbyvgvrtrbprqewjrvry72ze"

// TestGetClient tests the GetClient() function
func TestGetClient(t *testing.T) {
	_, err := store.GetClient()
	if err != nil {
		t.Fatalf("GetClient() returned an error: %v", err)
	}
}

// TestSaveFile tests the W3S.GetFile() function
func TestGetFile(t *testing.T) {
	// Initialize a web3.storage client
	W3SClient, err := store.GetClient()
	if err != nil {
		t.Fatalf("GetClient() returned an error: %v", err)
	}

	// Build a CID from ipbuf.txt
	ipbufTextCid, err := cid.Parse(ipbufTxt)
	if err != nil {
		t.Fatalf("cid.Parse() returned an error for %s: %v", ipbufTxt, err)
	}

	// Get the file contents from web3.storage
	_, gotFS, err := W3SClient.GetFile(ipbufTextCid)
	if err != nil {
		t.Fatalf("GetFile() returned an error while reading %s: %v", ipbufTxt, err)
	}

	// Ensure "ipbuf.txt" is in the returned fs.FS
	ipbufFile, err := gotFS.Open("/ipbuf.txt")
	if err != nil {
		t.Fatalf("Error opening w3s /ipbuf.txt: %v", err)
	}
	defer func(ipbufFile fs.File) {
		err := ipbufFile.Close()
		if err != nil {
			t.Errorf("Error closing w3s /ipbuf.txt: %v", err)
		}
	}(ipbufFile)

	// Read and validate the file contents
	got, err := io.ReadAll(ipbufFile)
	if err != nil {
		t.Fatalf("Error reading w3s /ipbuf.txt: %v", err)
	}

	// Compare to local ipbuf.txt file
	wantFile, err := os.Open("ipbuf.txt")
	if err != nil {
		t.Fatalf("Error opening local ipbuf.txt: %v", err)
	}
	want, err := io.ReadAll(wantFile)
	if err != nil {
		t.Fatalf("Error reading local ipbuf.txt: %v", err)
	}

	if string(got) != string(want) {
		t.Fatalf("Got '%s', want '%s'", string(got), string(want))
	}
}

// TestSaveFile tests the W3S.SaveFile() function
func TestSaveFile(t *testing.T) {
	// Initialize a web3.storage client
	W3SClient, err := store.GetClient()
	if err != nil {
		t.Fatalf("GetClient() returned an error: %v", err)
	}

	// Open the test file
	testFile, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Error opening test.txt: %v", err)
	}

	// Save the file to web3.storage
	testCID, err := W3SClient.SaveFile(testFile)
	if err != nil {
		t.Fatalf("SaveFile() returned an error while saving test.txt: %v", err)
	}

	// Get the file contents from web3.storage
	_, gotFS, err := W3SClient.GetFile(testCID)
	if err != nil {
		t.Fatalf("GetFile() returned an error while reading %s: %v", testCID, err)
	}

	// Ensure "test.txt" is in the returned fs.FS
	gotTestFile, err := gotFS.Open("/test.txt")
	if err != nil {
		t.Fatalf("Error opening /test.txt: %v", err)
	}
	defer func(ipbufFile fs.File) {
		err := ipbufFile.Close()
		if err != nil {
			t.Errorf("Error closing /test.txt: %v", err)
		}
	}(gotTestFile)

	// Read and validate the contents
	got, err := io.ReadAll(gotTestFile)
	if err != nil {
		t.Fatalf("Error reading /test.txt: %v", err)
	}

	want := "test"
	if string(got) != want {
		t.Fatalf("Got '%s', want '%s'", string(got), want)
	}
}
