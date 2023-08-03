# k3s offline downloader

## Describe:
### This binary is created for downloads k3s offline install package, include of 
* k3s exec bin
* airgap images.
### Optional, downloads some exec files like helm, k9s, etc.[TODO]

## Usage:
### help
```bash
# help
./k3s-offline-downloader -h
```
### list
```bash
# list all k3s stable version
./k3s-offline-downloader list
### download
```bash
# download k3s offline install package, include of k3s exec bin and airgap images.
./k3s-offline-downloader get -v v1.27.4+k3s1 -a amd64
```

### githubproxy
This binary will download k3s offline install package from GitHub.com, but in some cases, we can't access GitHub.com directly.
So, GitHub Proxy is added by default.The url is `https://ghproxy.com`.
If you want change it, you can use `--githubproxy` flag like this:
```bash
./k3s-offline-downloader get -v v1.27.4+k3s1 -a amd64 --githubproxy https://ghproxy.com
# or disable it
./k3s-offline-downloader get -v v1.27.4+k3s1 -a amd64 --githubproxy ""
````

### Specify the download directory
By default, the download directory is
* k3s bin : ./bin
* k3s images : ./rancher/k3s/agent/images/
If you want change it, you can use `--binpath` and `--imagepath` flag like this:
```bash
# Specify the download directory and download k3s offline install package, include of k3s exec bin and airgap images.
./k3s-offline-downloader get -v v1.27.4+k3s1 -a amd64 --binpath ./ --imagepath ./
```