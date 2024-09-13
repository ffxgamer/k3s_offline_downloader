package cmd

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"log"
	"os"
	"path/filepath"
	"time"
)

func downloadFile(fileUrl string, dstPath string) {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", fileUrl)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	// copy file from tmp to destination path
	_, filename := filepath.Split(resp.Filename)
	destFileName := filepath.Join(dstPath, filename)
	if err := os.Rename(resp.Filename, destFileName); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}
