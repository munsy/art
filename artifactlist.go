package main

import (
	"fmt"
	"time"
)

type ArtifactList struct {
	Results [] ArtifactListResult `json:"results"`
	Range struct {
		StartPos int `json:"start_pos"`
		EndPos   int `json:"end_pos"`
		Total    int `json:"total"`
	} `json:"range"`
}

type ArtifactListResult struct {
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
}

// PathName returns the path and name combined.
func (r *ArtifactListResult) PathName() string {
	return fmt.Sprintf("%s/%s", r.Path, r.Name)
}