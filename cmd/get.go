/*
Copyright © 2023 Peng Wenming <ffxgamer@163.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"

	"github.com/spf13/cobra"
)

var Version string
var Arch string
var GithubProxy string
var BinPath string
var ImagePath string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download k3s binary and images",
	Long: `Download k3s binary and images, you can specify the version and arch.
For example:
  k3s_offline_downloader get // downloads the latest version of k3s binary and images with amd64 arch
    or
  k3s_offline_downloader get -v v1.21.4+k3s1 -a arm64 // downloads the specified version of k3s binary and images with arm64 arch`,
	Run: func(cmd *cobra.Command, args []string) {
		client := github.NewClient(nil)
		ctx := context.Background()
		owner := "k3s-io"
		repo := "k3s"
		var binNo int
		var imageNo int
		if Arch == "amd64" {
			binNo = 0
			imageNo = 2
		} else if Arch == "arm64" {
			binNo = 13
			imageNo = 8
		} else {
			fmt.Println("arch error")
		}
		var release *github.RepositoryRelease
		if Version == "latest" {
			var err error
			release, _, err = client.Repositories.GetLatestRelease(ctx, owner, repo)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			var err error
			release, _, err = client.Repositories.GetReleaseByTag(ctx, owner, repo, Version)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Printf("Starting download k3s binary and images\n  version: %s\n  arch: %s\n", Version, Arch)
		if GithubProxy != "" {
			*release.Assets[binNo].BrowserDownloadURL = GithubProxy + "/" + *release.Assets[binNo].BrowserDownloadURL
			*release.Assets[imageNo].BrowserDownloadURL = GithubProxy + "/" + *release.Assets[imageNo].BrowserDownloadURL
		}
		fmt.Println("--------------------------------------")
		downloadFile(*release.Assets[binNo].BrowserDownloadURL, BinPath)
		fmt.Println("--------------------------------------")
		downloadFile(*release.Assets[imageNo].BrowserDownloadURL, ImagePath)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")
	getCmd.PersistentFlags().StringVarP(&Version, "version", "v", "latest", "the version of k3s, like v1.27.4+k3s1.")
	getCmd.PersistentFlags().StringVarP(&Arch, "arch", "a", "amd64", "the architecture of k3s, amd64 or arm64.")
	getCmd.PersistentFlags().StringVarP(&GithubProxy, "proxy", "p", "https://gh-proxy.com", "the proxy of github, like https://gh-proxy.com")
	//getCmd.MarkPersistentFlagRequired("proxy")
	getCmd.PersistentFlags().StringVarP(&BinPath, "binpath", "b", "./bin/", "the path to save k3s binary.")
	getCmd.PersistentFlags().StringVarP(&ImagePath, "imagepath", "i", "./rancher/k3s/agent/images/", "the path to save k3s images.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
