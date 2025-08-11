package model

import "encoding/xml"

// TaskListResponse represents the XML response for task list
type TaskListResponse struct {
	XMLName xml.Name `xml:"BESAPI"`
	Tasks   []Task   `xml:"Task"`
}

// Task represents a BigFix task
type Task struct {
	Resource          string       `xml:"Resource,attr" json:"resource"`
	LastModified      string       `xml:"LastModified,attr" json:"last_modified"`
	Name              string       `xml:"Name" json:"name"`
	ID                int          `xml:"ID" json:"id"`
	SiteName          string       `json:"site_name,omitempty"`
	SiteType          string       `json:"site_type,omitempty"`
	Title             string       `json:"title,omitempty"`
	Description       string       `json:"description,omitempty"`
	Relevance         []string     `json:"relevance,omitempty"`
	Category          string       `json:"category,omitempty"`
	DownloadSize      int64        `json:"download_size,omitempty"`
	Source            string       `json:"source,omitempty"`
	SourceID          string       `json:"source_id,omitempty"`
	SourceReleaseDate string       `json:"source_release_date,omitempty"`
	SourceSeverity    string       `json:"source_severity,omitempty"`
	Delay             string       `json:"delay,omitempty"`
	MIMEFields        []MIMEField  `json:"mime_fields,omitempty"`
	DefaultAction     *TaskAction  `json:"default_action,omitempty"`
	Actions           []TaskAction `json:"actions,omitempty"`
}

// TaskDetailResponse represents the XML response for task detail
type TaskDetailResponse struct {
	XMLName xml.Name   `xml:"BES"`
	Task    TaskDetail `xml:"Task"`
}

// TaskDetail represents detailed task information
type TaskDetail struct {
	Title             string       `xml:"Title" json:"title"`
	Description       string       `xml:"Description" json:"description"`
	Relevance         []string     `xml:"Relevance" json:"relevance"`
	Category          string       `xml:"Category" json:"category,omitempty"`
	DownloadSize      int64        `xml:"DownloadSize" json:"download_size,omitempty"`
	Source            string       `xml:"Source" json:"source,omitempty"`
	SourceID          string       `xml:"SourceID" json:"source_id,omitempty"`
	SourceReleaseDate string       `xml:"SourceReleaseDate" json:"source_release_date,omitempty"`
	SourceSeverity    string       `xml:"SourceSeverity" json:"source_severity,omitempty"`
	Delay             string       `xml:"Delay" json:"delay,omitempty"`
	MIMEFields        []MIMEField  `xml:"MIMEField" json:"mime_fields,omitempty"`
	DefaultAction     *TaskAction  `xml:"DefaultAction" json:"default_action,omitempty"`
	Actions           []TaskAction `xml:"Action" json:"actions,omitempty"`
}

// TaskAction represents an action in a task
type TaskAction struct {
	ID              string `xml:"ID,attr" json:"id"`
	Description     string `xml:"Description" json:"description,omitempty"`
	ActionScript    string `xml:"ActionScript" json:"action_script,omitempty"`
	SuccessCriteria string `xml:"SuccessCriteria" json:"success_criteria,omitempty"`
}

// ToTask converts TaskDetail to Task model
func (td *TaskDetail) ToTask(id int, resource, siteName, siteType string) *Task {
	return &Task{
		ID:                id,
		Resource:          resource,
		Name:              td.Title,
		SiteName:          siteName,
		SiteType:          siteType,
		Title:             td.Title,
		Description:       td.Description,
		Relevance:         td.Relevance,
		Category:          td.Category,
		DownloadSize:      td.DownloadSize,
		Source:            td.Source,
		SourceID:          td.SourceID,
		SourceReleaseDate: td.SourceReleaseDate,
		SourceSeverity:    td.SourceSeverity,
		Delay:             td.Delay,
		MIMEFields:        td.MIMEFields,
		DefaultAction:     td.DefaultAction,
		Actions:           td.Actions,
	}
}
