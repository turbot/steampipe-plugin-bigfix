package model

import "encoding/xml"

// BigFixPropertyListResponse represents the XML response for property list
type BigFixPropertyListResponse struct {
	XMLName    xml.Name         `xml:"BESAPI"`
	Properties []BigFixProperty `xml:"Property"`
}

// BigFixProperty represents a BigFix property (list response)
type BigFixProperty struct {
	Resource     string `xml:"Resource,attr" json:"resource"`
	LastModified string `xml:"LastModified,attr" json:"last_modified,omitempty"`
	Name         string `xml:"Name" json:"name"`
	ID           int    `xml:"ID" json:"id"`
	IsReserved   int    `xml:"IsReserved" json:"is_reserved"`
	Definition   string `json:"definition,omitempty"`
}

// BigFixPropertyDetailResponse represents the XML response for property detail
type BigFixPropertyDetailResponse struct {
	XMLName  xml.Name             `xml:"BES"`
	Property BigFixPropertyDetail `xml:"Property"`
}

// BigFixPropertyDetail represents detailed property information
type BigFixPropertyDetail struct {
	Name       string `xml:"Name,attr" json:"name"`
	Definition string `xml:",chardata" json:"definition"`
}

// ToBigFixProperty converts BigFixPropertyDetail to BigFixProperty model
func (pd *BigFixPropertyDetail) ToBigFixProperty(id int, resource, name string, isReserved int) *BigFixProperty {
	return &BigFixProperty{
		ID:         id,
		Resource:   resource,
		Name:       name,
		IsReserved: isReserved,
		Definition: pd.Definition,
	}
}

// ToBigFixPropertyFromList converts list BigFixProperty to detailed BigFixProperty model
func (p *BigFixProperty) ToBigFixPropertyFromList() *BigFixProperty {
	return &BigFixProperty{
		ID:           p.ID,
		Resource:     p.Resource,
		Name:         p.Name,
		IsReserved:   p.IsReserved,
		LastModified: p.LastModified,
	}
}
