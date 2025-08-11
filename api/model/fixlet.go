package model

import "encoding/xml"

// FixletListResponse represents the XML response for fixlet list
type FixletListResponse struct {
	XMLName xml.Name `xml:"BESAPI"`
	Fixlets []Fixlet `xml:"Fixlet"`
}

// Fixlet represents a BigFix fixlet (list response)
type Fixlet struct {
	Resource          string         `xml:"Resource,attr" json:"resource"`
	LastModified      string         `xml:"LastModified,attr" json:"last_modified"`
	Name              string         `xml:"Name" json:"name"`
	ID                int            `xml:"ID" json:"id"`
	SiteName          string         `json:"site_name,omitempty"`
	SiteType          string         `json:"site_type,omitempty"`
	Title             string         `json:"title,omitempty"`
	Description       string         `json:"description,omitempty"`
	Relevance         []string       `json:"relevance,omitempty"`
	Category          string         `json:"category,omitempty"`
	DownloadSize      int64          `json:"download_size,omitempty"`
	Source            string         `json:"source,omitempty"`
	SourceID          string         `json:"source_id,omitempty"`
	SourceReleaseDate string         `json:"source_release_date,omitempty"`
	SourceSeverity    string         `json:"source_severity,omitempty"`
	CVENames          string         `json:"cve_names,omitempty"`
	MIMEFields        []MIMEField    `json:"mime_fields,omitempty"`
	Delay             string         `json:"delay,omitempty"`
	DefaultAction     *FixletAction  `json:"default_action,omitempty"`
	Actions           []FixletAction `json:"actions,omitempty"`
}

// FixletDetailResponse represents the XML response for fixlet detail
type FixletDetailResponse struct {
	XMLName xml.Name     `xml:"BES"`
	Fixlet  FixletDetail `xml:"Fixlet"`
}

// FixletDetail represents detailed fixlet information
type FixletDetail struct {
	Title             string         `xml:"Title" json:"title"`
	Description       string         `xml:"Description" json:"description"`
	Relevance         []string       `xml:"Relevance" json:"relevance"`
	Category          string         `xml:"Category" json:"category,omitempty"`
	DownloadSize      int64          `xml:"DownloadSize" json:"download_size,omitempty"`
	Source            string         `xml:"Source" json:"source,omitempty"`
	SourceID          string         `xml:"SourceID" json:"source_id,omitempty"`
	SourceReleaseDate string         `xml:"SourceReleaseDate" json:"source_release_date,omitempty"`
	SourceSeverity    string         `xml:"SourceSeverity" json:"source_severity,omitempty"`
	CVENames          string         `xml:"CVENames" json:"cve_names,omitempty"`
	MIMEFields        []MIMEField    `xml:"MIMEField" json:"mime_fields,omitempty"`
	Delay             string         `xml:"Delay" json:"delay,omitempty"`
	DefaultAction     *FixletAction  `xml:"DefaultAction" json:"default_action,omitempty"`
	Actions           []FixletAction `xml:"Action" json:"actions,omitempty"`
}

// FixletAction represents an action in a fixlet
type FixletAction struct {
	ID              string `xml:"ID,attr" json:"id"`
	Description     string `xml:"Description" json:"description,omitempty"`
	ActionScript    string `xml:"ActionScript" json:"action_script,omitempty"`
	SuccessCriteria string `xml:"SuccessCriteria" json:"success_criteria,omitempty"`
}

// ToFixlet converts FixletDetail to Fixlet model
func (fd *FixletDetail) ToFixlet(id int, resource, siteName, siteType string) *Fixlet {
	return &Fixlet{
		ID:                id,
		Resource:          resource,
		Name:              fd.Title,
		SiteName:          siteName,
		SiteType:          siteType,
		Title:             fd.Title,
		Description:       fd.Description,
		Relevance:         fd.Relevance,
		Category:          fd.Category,
		DownloadSize:      fd.DownloadSize,
		Source:            fd.Source,
		SourceID:          fd.SourceID,
		SourceReleaseDate: fd.SourceReleaseDate,
		SourceSeverity:    fd.SourceSeverity,
		CVENames:          fd.CVENames,
		MIMEFields:        fd.MIMEFields,
		Delay:             fd.Delay,
		DefaultAction:     fd.DefaultAction,
		Actions:           fd.Actions,
	}
}

// ToFixletFromList converts list Fixlet to detailed Fixlet model
func (f *Fixlet) ToFixletFromList() *Fixlet {
	return &Fixlet{
		ID:           f.ID,
		Resource:     f.Resource,
		Name:         f.Name,
		SiteName:     f.SiteName,
		SiteType:     f.SiteType,
		Title:        f.Name, // Use Name as Title for list items
		LastModified: f.LastModified,
	}
}
