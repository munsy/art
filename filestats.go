package main

type ArtifactoryFileStats struct {
	URI                  string `json:"uri"`
	DownloadCount        int    `json:"downloadCount"`
	LastDownloaded       int64  `json:"lastDownloaded"`
	LastDownloadedBy     string `json:"lastDownloadedBy"`
	RemoteDownloadCount  int    `json:"remoteDownloadCount"`
	RemoteLastDownloaded int    `json:"remoteLastDownloaded"`
}