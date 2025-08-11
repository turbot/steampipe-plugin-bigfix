package model

import "encoding/xml"

// RoleListResponse represents the XML response for role list
type RoleListResponse struct {
	XMLName xml.Name `xml:"BESAPI"`
	Roles   []Role   `xml:"Role"`
}

// Role represents a BigFix role
type Role struct {
	Resource                      string          `xml:"Resource,attr" json:"resource"`
	LastModified                  string          `xml:"LastModified,attr" json:"last_modified,omitempty"`
	Name                          string          `xml:"Name" json:"name"`
	ID                            int             `xml:"ID" json:"id"`
	MasterOperator                int             `xml:"MasterOperator" json:"master_operator"`
	CustomContent                 int             `xml:"CustomContent" json:"custom_content"`
	ShowOtherActions              int             `xml:"ShowOtherActions" json:"show_other_actions"`
	StopOtherActions              int             `xml:"StopOtherActions" json:"stop_other_actions"`
	CanCreateActions              int             `xml:"CanCreateActions" json:"can_create_actions"`
	PostActionBehaviorPrivilege   string          `xml:"PostActionBehaviorPrivilege" json:"post_action_behavior_privilege"`
	ActionScriptCommandsPrivilege string          `xml:"ActionScriptCommandsPrivilege" json:"action_script_commands_privilege"`
	CanSendMultipleRefresh        int             `xml:"CanSendMultipleRefresh" json:"can_send_multiple_refresh"`
	CanSubmitQueries              int             `xml:"CanSubmitQueries" json:"can_submit_queries"`
	CanLock                       int             `xml:"CanLock" json:"can_lock"`
	UnmanagedAssetPrivilege       string          `xml:"UnmanagedAssetPrivilege" json:"unmanaged_asset_privilege"`
	InterfaceLogins               InterfaceLogins `xml:"InterfaceLogins" json:"interface_logins"`
}

// InterfaceLogins represents the interface login permissions for a role
type InterfaceLogins struct {
	Console bool `xml:"Console" json:"console"`
	WebUI   bool `xml:"WebUI" json:"webui"`
	API     bool `xml:"API" json:"api"`
}

// RoleDetailResponse represents the XML response for role detail
type RoleDetailResponse struct {
	XMLName xml.Name `xml:"BESAPI"`
	Role    Role     `xml:"Role"`
}

// ToRole converts Role to itself (for consistency with other models)
func (r *Role) ToRole() *Role {
	return r
}

// ToRoleFromList converts list Role to detailed Role model
func (r *Role) ToRoleFromList() *Role {
	return &Role{
		ID:                            r.ID,
		Resource:                      r.Resource,
		Name:                          r.Name,
		MasterOperator:                r.MasterOperator,
		CustomContent:                 r.CustomContent,
		ShowOtherActions:              r.ShowOtherActions,
		StopOtherActions:              r.StopOtherActions,
		CanCreateActions:              r.CanCreateActions,
		PostActionBehaviorPrivilege:   r.PostActionBehaviorPrivilege,
		ActionScriptCommandsPrivilege: r.ActionScriptCommandsPrivilege,
		CanSendMultipleRefresh:        r.CanSendMultipleRefresh,
		CanSubmitQueries:              r.CanSubmitQueries,
		CanLock:                       r.CanLock,
		UnmanagedAssetPrivilege:       r.UnmanagedAssetPrivilege,
		InterfaceLogins:               r.InterfaceLogins,
		LastModified:                  r.LastModified,
	}
}
