package model

import "encoding/xml"

// AnalysisListResponse represents the XML response for analysis list
type AnalysisListResponse struct {
	XMLName  xml.Name   `xml:"BESAPI"`
	Analyses []Analysis `xml:"Analysis"`
}

// Analysis represents a BigFix analysis
type Analysis struct {
	Resource          string             `xml:"Resource,attr" json:"resource"`
	LastModified      string             `xml:"LastModified,attr" json:"last_modified"`
	Name              string             `xml:"Name" json:"name"`
	ID                int                `xml:"ID" json:"id"`
	SiteName          string             `json:"site_name,omitempty"`
	SiteType          string             `json:"site_type,omitempty"`
	Title             string             `json:"title,omitempty"`
	Description       string             `json:"description,omitempty"`
	Relevance         []string           `json:"relevance,omitempty"`
	Category          string             `json:"category,omitempty"`
	Source            string             `json:"source,omitempty"`
	SourceReleaseDate string             `json:"source_release_date,omitempty"`
	Delay             string             `json:"delay,omitempty"`
	MIMEFields        []MIMEField        `json:"mime_fields,omitempty"`
	Properties        []AnalysisProperty `json:"properties,omitempty"`
}

// AnalysisDetailResponse represents the XML response for analysis detail
type AnalysisDetailResponse struct {
	XMLName  xml.Name       `xml:"BES"`
	Analysis AnalysisDetail `xml:"Analysis"`
}

// AnalysisDetail represents detailed analysis information
type AnalysisDetail struct {
	Title             string             `xml:"Title" json:"title"`
	Description       string             `xml:"Description" json:"description"`
	Relevance         []string           `xml:"Relevance" json:"relevance"`
	Category          string             `xml:"Category" json:"category,omitempty"`
	Source            string             `xml:"Source" json:"source,omitempty"`
	SourceReleaseDate string             `xml:"SourceReleaseDate" json:"source_release_date,omitempty"`
	Delay             string             `xml:"Delay" json:"delay,omitempty"`
	MIMEFields        []MIMEField        `xml:"MIMEField" json:"mime_fields,omitempty"`
	Properties        []AnalysisProperty `xml:"Property" json:"properties,omitempty"`
}

// AnalysisProperty represents a property in an analysis
type AnalysisProperty struct {
	Name             string `xml:"Name,attr" json:"name"`
	ID               string `xml:"ID,attr" json:"id"`
	EvaluationPeriod string `xml:"EvaluationPeriod,attr" json:"evaluation_period,omitempty"`
	Value            string `xml:",chardata" json:"value"`
}

// MIMEField represents a MIME field in analysis/task
type MIMEField struct {
	Name  string `xml:"Name" json:"name"`
	Value string `xml:"Value" json:"value"`
}

// ToAnalysis converts AnalysisDetail to Analysis model
func (ad *AnalysisDetail) ToAnalysis(id int, resource, siteName, siteType string) *Analysis {
	return &Analysis{
		ID:                id,
		Resource:          resource,
		Name:              ad.Title,
		SiteName:          siteName,
		SiteType:          siteType,
		Title:             ad.Title,
		Description:       ad.Description,
		Relevance:         ad.Relevance,
		Category:          ad.Category,
		Source:            ad.Source,
		SourceReleaseDate: ad.SourceReleaseDate,
		Delay:             ad.Delay,
		MIMEFields:        ad.MIMEFields,
		Properties:        ad.Properties,
	}
}
