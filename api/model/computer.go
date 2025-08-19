package model

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

// BESAPI represents the root element of BigFix API XML response for listing computers
type BESAPI struct {
	XMLName   xml.Name          `xml:"BESAPI"`
	Computers []ComputerListXML `xml:"Computer"`
}

// ComputerListXML represents the XML structure for computer list responses
type ComputerListXML struct {
	Resource       string `xml:"Resource,attr"`
	ID             int    `xml:"ID"`
	Name           string `xml:"Name"`
	OS             string `xml:"OS"`
	LastReportTime string `xml:"LastReportTime"`
	CPU            string `xml:"CPU"`
	IPAddress      string `xml:"IPAddress"`
}

// ComputerXMLResponse represents the XML response structure for single computer details
type ComputerXMLResponse struct {
	XMLName  xml.Name    `xml:"BESAPI"`
	Computer ComputerXML `xml:"Computer"`
}

// ComputerXML represents the XML structure of a computer with properties
type ComputerXML struct {
	Resource   string     `xml:"Resource,attr"`
	Properties []Property `xml:"Property"`
}

// Computer represents a BigFix computer/endpoint for API return
type Computer struct {
	Resource           string      `json:"resource,omitempty"`
	ID                 int         `json:"id"`
	Name               string      `json:"name,omitempty"`
	OS                 string      `json:"os,omitempty"`
	LastReportTime     *time.Time  `json:"last_report_time,omitempty"`
	CPU                string      `json:"cpu,omitempty"`
	IPAddress          string      `json:"ip_address,omitempty"`
	IPv6Address        string      `json:"ipv6_address,omitempty"`
	DNSName            string      `json:"dns_name,omitempty"`
	MACAddress         string      `json:"mac_address,omitempty"`
	OSFamily           string      `json:"os_family,omitempty"`
	OSName             string      `json:"os_name,omitempty"`
	OSVersion          string      `json:"os_version,omitempty"`
	UserName           string      `json:"user_name,omitempty"`
	RAM                string      `json:"ram,omitempty"`
	Locked             string      `json:"locked,omitempty"`
	BESRelaySelection  string      `json:"bes_relay_selection,omitempty"`
	Relay              string      `json:"relay,omitempty"`
	DistanceToBESRelay string      `json:"distance_to_bes_relay,omitempty"`
	AgentType          string      `json:"agent_type,omitempty"`
	DeviceType         string      `json:"device_type,omitempty"`
	AgentVersion       string      `json:"agent_version,omitempty"`
	ComputerType       string      `json:"computer_type,omitempty"`
	LicenseType        string      `json:"license_type,omitempty"`
	FreeSpaceOnSystem  string      `json:"free_space_on_system,omitempty"`
	TotalSizeOfSystem  string      `json:"total_size_of_system,omitempty"`
	BIOS               string      `json:"bios,omitempty"`
	SubnetAddress      string      `json:"subnet_address,omitempty"`
	ClientSettings     []NameValue `json:"client_settings,omitempty"`
	SubscribedSites    []string    `json:"subscribed_sites,omitempty"`
	Properties         []Property  `json:"properties,omitempty"`
}

// NameValue represents a name-value pair for client settings
type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ToComputer converts ComputerXML response to Computer model
func (cx *ComputerXML) ToComputer() (*Computer, error) {
	computer := &Computer{
		Resource:   cx.Resource,
		Properties: cx.Properties,
	}

	// Time parsing layouts
	layouts := []string{
		"Mon, 2 Jan 2006 15:04:05 -0700",  // RFC1123Z format with numeric timezone
		"Mon, 02 Jan 2006 15:04:05 -0700", // RFC1123Z with zero-padded day
		time.RFC1123Z,
		time.RFC1123,
		"Mon, 2 Jan 2006 15:04:05 MST",
		time.RFC3339,
		time.RFC3339Nano,
		time.ANSIC,
	}

	// Parse properties into structured fields
	for _, prop := range cx.Properties {
		switch prop.Name {
		case "ID":
			if id := parseIntFromString(prop.Value); id != 0 {
				computer.ID = id
			}
		case "Computer Name":
			computer.Name = prop.Value
		case "OS":
			computer.OS = prop.Value
		case "Last Report Time":
			for _, layout := range layouts {
				if parsed, err := time.Parse(layout, prop.Value); err == nil {
					computer.LastReportTime = &parsed
					break
				}
			}
		case "CPU":
			computer.CPU = prop.Value
		case "IP Address":
			computer.IPAddress = prop.Value
		case "IPv6 Address":
			computer.IPv6Address = prop.Value
		case "DNS Name":
			computer.DNSName = prop.Value
		case "MAC Address":
			computer.MACAddress = prop.Value
		case "OS Family":
			computer.OSFamily = prop.Value
		case "OS Name":
			computer.OSName = prop.Value
		case "OS Version":
			computer.OSVersion = prop.Value
		case "User Name":
			computer.UserName = prop.Value
		case "RAM":
			computer.RAM = prop.Value
		case "Locked":
			computer.Locked = prop.Value
		case "BES Relay Selection Method":
			computer.BESRelaySelection = prop.Value
		case "Relay":
			computer.Relay = prop.Value
		case "Distance to BES Relay":
			computer.DistanceToBESRelay = prop.Value
		case "Agent Type":
			computer.AgentType = prop.Value
		case "Device Type":
			computer.DeviceType = prop.Value
		case "Agent Version":
			computer.AgentVersion = prop.Value
		case "Computer Type":
			computer.ComputerType = prop.Value
		case "License Type":
			computer.LicenseType = prop.Value
		case "Free Space on System Drive":
			computer.FreeSpaceOnSystem = prop.Value
		case "Total Size of System Drive":
			computer.TotalSizeOfSystem = prop.Value
		case "BIOS":
			computer.BIOS = prop.Value
		case "Subnet Address":
			computer.SubnetAddress = prop.Value
		case "Client Settings":
			// Parse client setting value to extract name=value pairs
			if name, value := parseClientSetting(prop.Value); name != "" {
				computer.ClientSettings = append(computer.ClientSettings, NameValue{
					Name:  name,
					Value: value,
				})
			}
		case "Subscribed Sites":
			computer.SubscribedSites = append(computer.SubscribedSites, prop.Value)
		}
	}

	return computer, nil
}

// Helper function to parse integer from string
func parseIntFromString(s string) int {
	if result, err := strconv.Atoi(s); err == nil {
		return result
	}
	return 0
}

// Helper function to parse client settings in format "name=value"
func parseClientSetting(setting string) (name, value string) {
	parts := strings.SplitN(setting, "=", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	// If no equals sign, treat the whole string as the name with empty value
	return setting, ""
}

// ToComputer converts ComputerListXML to Computer model for list responses
func (cl *ComputerListXML) ToComputer() (*Computer, error) {
	computer := &Computer{
		Resource:  cl.Resource,
		ID:        cl.ID,
		Name:      cl.Name,
		OS:        cl.OS,
		CPU:       cl.CPU,
		IPAddress: cl.IPAddress,
	}

	// Parse LastReportTime with multiple layouts
	layouts := []string{
		"Mon, 2 Jan 2006 15:04:05 -0700",  // RFC1123Z format with numeric timezone
		"Mon, 02 Jan 2006 15:04:05 -0700", // RFC1123Z with zero-padded day
		time.RFC1123Z,
		time.RFC1123,
		"Mon, 2 Jan 2006 15:04:05 MST",
		time.RFC3339,
		time.RFC3339Nano,
		time.ANSIC,
	}

	for _, layout := range layouts {
		if parsed, err := time.Parse(layout, cl.LastReportTime); err == nil {
			computer.LastReportTime = &parsed
			break
		}
	}

	return computer, nil
}

// Property represents a generic BigFix Property element
type Property struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:",chardata"`
}
