package model

import (
	"encoding/xml"
)

// SiteListResponse represents the root element of BigFix API XML response for listing sites
type SiteListResponse struct {
	XMLName       xml.Name       `xml:"BESAPI"`
	ExternalSites []ExternalSite `xml:"ExternalSite"`
	OperatorSites []OperatorSite `xml:"OperatorSite"`
	ActionSites   []ActionSite   `xml:"ActionSite"`
}

// ExternalSite represents an external site in the list response
type ExternalSite struct {
	Resource    string `xml:"Resource,attr"`
	Name        string `xml:"Name"`
	DisplayName string `xml:"DisplayName"`
	GatherURL   string `xml:"GatherURL"`
}

// OperatorSite represents an operator site in the list response
type OperatorSite struct {
	Resource    string `xml:"Resource,attr"`
	Name        string `xml:"Name"`
	DisplayName string `xml:"DisplayName"`
	GatherURL   string `xml:"GatherURL"`
}

// ActionSite represents an action site in the list response
type ActionSite struct {
	Resource    string `xml:"Resource,attr"`
	Name        string `xml:"Name"`
	DisplayName string `xml:"DisplayName"`
	GatherURL   string `xml:"GatherURL"`
}

// SiteDetailResponse represents the XML response structure for single site details
type SiteDetailResponse struct {
	XMLName      xml.Name            `xml:"BES"`
	ActionSite   *ActionSiteDetail   `xml:"ActionSite,omitempty"`
	ExternalSite *ExternalSiteDetail `xml:"ExternalSite,omitempty"`
	OperatorSite *OperatorSiteDetail `xml:"OperatorSite,omitempty"`
}

// Subscription represents the subscription structure in BigFix XML
type Subscription struct {
	Mode string `xml:"Mode"`
}

// ActionSiteDetail represents detailed action site information
type ActionSiteDetail struct {
	Name                 string       `xml:"Name"`
	DisplayName          string       `xml:"DisplayName"`
	Description          string       `xml:"Description"`
	GlobalReadPermission string       `xml:"GlobalReadPermission"`
	Subscription         Subscription `xml:"Subscription"`
	GatherURL            string       `xml:"GatherURL"`
}

// ExternalSiteDetail represents detailed external site information
type ExternalSiteDetail struct {
	Name                 string       `xml:"Name"`
	DisplayName          string       `xml:"DisplayName"`
	Description          string       `xml:"Description"`
	GlobalReadPermission string       `xml:"GlobalReadPermission"`
	Subscription         Subscription `xml:"Subscription"`
	GatherURL            string       `xml:"GatherURL"`
}

// OperatorSiteDetail represents detailed operator site information
type OperatorSiteDetail struct {
	Name                 string       `xml:"Name"`
	DisplayName          string       `xml:"DisplayName"`
	Description          string       `xml:"Description"`
	GlobalReadPermission string       `xml:"GlobalReadPermission"`
	Subscription         Subscription `xml:"Subscription"`
	GatherURL            string       `xml:"GatherURL"`
}

// Site represents a BigFix site for API return
type Site struct {
	Resource             string `json:"resource,omitempty"`
	Name                 string `json:"name"`
	DisplayName          string `json:"display_name,omitempty"`
	Description          string `json:"description,omitempty"`
	Type                 string `json:"type"` // "action", "external", "operator"
	GlobalReadPermission *bool  `json:"global_read_permission,omitempty"`
	SubscriptionMode     string `json:"subscription_mode,omitempty"`
	GatherURL            string `json:"gather_url,omitempty"`
}

// ToSite converts different site types to unified Site model
func (es *ExternalSite) ToSite() *Site {
	return &Site{
		Resource:    es.Resource,
		Name:        es.Name,
		DisplayName: es.DisplayName,
		Type:        "external",
		GatherURL:   es.GatherURL,
	}
}

func (os *OperatorSite) ToSite() *Site {
	return &Site{
		Resource:    os.Resource,
		Name:        os.Name,
		DisplayName: os.DisplayName,
		Type:        "operator",
		GatherURL:   os.GatherURL,
	}
}

func (as *ActionSite) ToSite() *Site {
	return &Site{
		Resource:    as.Resource,
		Name:        as.Name,
		DisplayName: as.DisplayName,
		Type:        "action",
		GatherURL:   as.GatherURL,
	}
}

// ToSite converts detailed site information to Site model
func (asd *ActionSiteDetail) ToSite() *Site {
	globalRead := asd.GlobalReadPermission == "true"
	return &Site{
		Resource:             "", // Set by calling function
		Name:                 asd.Name,
		DisplayName:          asd.DisplayName,
		Description:          asd.Description,
		Type:                 "action",
		GlobalReadPermission: &globalRead,
		SubscriptionMode:     asd.Subscription.Mode,
		GatherURL:            asd.GatherURL,
	}
}

func (esd *ExternalSiteDetail) ToSite() *Site {
	globalRead := esd.GlobalReadPermission == "true"
	return &Site{
		Resource:             "", // Set by calling function
		Name:                 esd.Name,
		DisplayName:          esd.DisplayName,
		Description:          esd.Description,
		Type:                 "external",
		GlobalReadPermission: &globalRead,
		SubscriptionMode:     esd.Subscription.Mode,
		GatherURL:            esd.GatherURL,
	}
}

func (osd *OperatorSiteDetail) ToSite() *Site {
	globalRead := osd.GlobalReadPermission == "true"
	return &Site{
		Resource:             "", // Set by calling function
		Name:                 osd.Name,
		DisplayName:          osd.DisplayName,
		Description:          osd.Description,
		Type:                 "operator",
		GlobalReadPermission: &globalRead,
		SubscriptionMode:     osd.Subscription.Mode,
		GatherURL:            osd.GatherURL,
	}
}

// SitePermission represents a site permission
type SitePermission struct {
	Resource   string             `xml:"Resource,attr" json:"resource"`
	Permission string             `xml:"Permission" json:"permission"`
	Operator   SitePermissionUser `xml:"Operator" json:"operator"`
}

// SitePermissionUser represents the operator in a site permission
type SitePermissionUser struct {
	Resource string `xml:"Resource,attr" json:"resource"`
	Name     string `xml:",chardata" json:"name"`
}

// SitePermissionsResponse represents the XML response for site permissions
type SitePermissionsResponse struct {
	XMLName     xml.Name         `xml:"BESAPI"`
	Permissions []SitePermission `xml:"SitePermission"`
}

// SiteFile represents a file in a site
type SiteFile struct {
	Resource     string `xml:"Resource,attr" json:"resource,omitempty"`
	Name         string `xml:"Name" json:"name"`
	ID           int    `xml:"ID" json:"id"`
	LastModified string `xml:"LastModified" json:"last_modified"`
	FileSize     string `xml:"FileSize" json:"file_size"`
	IsClientFile int    `xml:"IsClientFile" json:"is_client_file"`
	Size         int64  `xml:"Size,attr" json:"size,omitempty"`
	SHA1         string `xml:"SHA1,attr" json:"sha1,omitempty"`
	SHA256       string `xml:"SHA256,attr" json:"sha256,omitempty"`
	DownloadURL  string `xml:",chardata" json:"download_url,omitempty"`
}

// SiteFilesResponse represents the XML response for site files
type SiteFilesResponse struct {
	XMLName xml.Name   `xml:"BESAPI"`
	Files   []SiteFile `xml:"SiteFile"`
}
