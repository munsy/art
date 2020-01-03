package main

import (
	"time"
)

type ArtifactList struct {
	Results []ArtifactResult `json:"results"`
	Range struct {
		StartPos int `json:"start_pos"`
		EndPos   int `json:"end_pos"`
		Total    int `json:"total"`
	} `json:"range"`
}

type ArtifactResult struct {
	Repo       string    `json:"repo"`
	Path       string    `json:"path"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Size       int       `json:"size"`
	Created    time.Time `json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `json:"modified"`
	ModifiedBy string    `json:"modified_by"`
	Updated    time.Time `json:"updated"`
	Stats      []ArtifactStats `json:"stats"`
}

type ArtifactStats struct {
	Downloaded      time.Time `json:"downloaded"`
	DownloadedBy    string    `json:"downloaded_by"`
	Downloads       int       `json:"downloads"`
	RemoteDownloads int       `json:"remote_downloads"`
}

func (r ArtifactResult) TotalDownloads() int {
	count := 0
	for i := 0; i < len(r.Stats); i++ {
		count += r.Stats[i].Downloads
	}
	return count
}

// Len is the number of elements in the collection.
func (list *ArtifactList) Len() int {
	return len(list.Results)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (list *ArtifactList) Less(i, j int) bool {
	return list.Results[i].TotalDownloads() < list.Results[j].TotalDownloads()
}

// Swap swaps the elements with indexes i and j.
func (list *ArtifactList) Swap(i, j int) {
	var temp1 = list.Results[i]
	list.Results[i] = list.Results[j]
	list.Results[j] = temp1 
}
