package model

import "encoding/xml"

// ActionListResponse represents the XML response for action list
type ActionListResponse struct {
	XMLName xml.Name `xml:"BESAPI"`
	Actions []Action `xml:"Action"`
}

// Action represents a BigFix action (list response)
type Action struct {
	Resource        string               `xml:"Resource,attr" json:"resource"`
	LastModified    string               `xml:"LastModified,attr" json:"last_modified"`
	Name            string               `xml:"Name" json:"name"`
	ID              int                  `xml:"ID" json:"id"`
	Title           string               `json:"title,omitempty"`
	Relevance       string               `json:"relevance,omitempty"`
	ActionScript    string               `json:"action_script,omitempty"`
	SuccessCriteria string               `json:"success_criteria,omitempty"`
	Settings        *ActionSettings      `json:"settings,omitempty"`
	SettingsLocks   *ActionSettingsLocks `json:"settings_locks,omitempty"`
	Target          *ActionTarget        `json:"target,omitempty"`
	IsUrgent        bool                 `json:"is_urgent,omitempty"`
}

// ActionDetailResponse represents the XML response for action detail
type ActionDetailResponse struct {
	XMLName xml.Name     `xml:"BES"`
	Action  ActionDetail `xml:"SingleAction"`
}

// ActionDetail represents detailed action information
type ActionDetail struct {
	Title           string               `xml:"Title" json:"title"`
	Relevance       string               `xml:"Relevance" json:"relevance"`
	ActionScript    ActionScript         `xml:"ActionScript" json:"action_script"`
	SuccessCriteria string               `xml:"SuccessCriteria" json:"success_criteria,omitempty"`
	Settings        *ActionSettings      `xml:"Settings" json:"settings,omitempty"`
	SettingsLocks   *ActionSettingsLocks `xml:"SettingsLocks" json:"settings_locks,omitempty"`
	Target          *ActionTarget        `xml:"Target" json:"target,omitempty"`
	IsUrgent        bool                 `xml:"IsUrgent" json:"is_urgent,omitempty"`
}

// ActionScript represents the action script with MIME type
type ActionScript struct {
	MIMEType string `xml:"MIMEType,attr" json:"mime_type,omitempty"`
	Content  string `xml:",chardata" json:"content"`
}

// ActionSettings represents action settings
type ActionSettings struct {
	PreActionShowUI         bool               `xml:"PreActionShowUI" json:"pre_action_show_ui,omitempty"`
	HasRunningMessage       bool               `xml:"HasRunningMessage" json:"has_running_message,omitempty"`
	HasTimeRange            bool               `xml:"HasTimeRange" json:"has_time_range,omitempty"`
	HasStartTime            bool               `xml:"HasStartTime" json:"has_start_time,omitempty"`
	HasEndTime              bool               `xml:"HasEndTime" json:"has_end_time,omitempty"`
	EndDateTimeLocalOffset  string             `xml:"EndDateTimeLocalOffset" json:"end_date_time_local_offset,omitempty"`
	HasDayOfWeekConstraint  bool               `xml:"HasDayOfWeekConstraint" json:"has_day_of_week_constraint,omitempty"`
	UseUTCTime              bool               `xml:"UseUTCTime" json:"use_utc_time,omitempty"`
	ActiveUserRequirement   string             `xml:"ActiveUserRequirement" json:"active_user_requirement,omitempty"`
	ActiveUserType          string             `xml:"ActiveUserType" json:"active_user_type,omitempty"`
	HasWhose                bool               `xml:"HasWhose" json:"has_whose,omitempty"`
	PreActionCacheDownload  bool               `xml:"PreActionCacheDownload" json:"pre_action_cache_download,omitempty"`
	Reapply                 bool               `xml:"Reapply" json:"reapply,omitempty"`
	HasReapplyLimit         bool               `xml:"HasReapplyLimit" json:"has_reapply_limit,omitempty"`
	ReapplyLimit            int                `xml:"ReapplyLimit" json:"reapply_limit,omitempty"`
	HasReapplyInterval      bool               `xml:"HasReapplyInterval" json:"has_reapply_interval,omitempty"`
	HasRetry                bool               `xml:"HasRetry" json:"has_retry,omitempty"`
	HasTemporalDistribution bool               `xml:"HasTemporalDistribution" json:"has_temporal_distribution,omitempty"`
	ContinueOnErrors        bool               `xml:"ContinueOnErrors" json:"continue_on_errors,omitempty"`
	PostActionBehavior      PostActionBehavior `xml:"PostActionBehavior" json:"post_action_behavior,omitempty"`
	IsOffer                 bool               `xml:"IsOffer" json:"is_offer,omitempty"`
}

// PostActionBehavior represents post action behavior settings
type PostActionBehavior struct {
	Behavior string `xml:"Behavior,attr" json:"behavior,omitempty"`
}

// ActionSettingsLocks represents action settings locks
type ActionSettingsLocks struct {
	ActionUITitle          bool                    `xml:"ActionUITitle" json:"action_ui_title,omitempty"`
	PreActionShowUI        bool                    `xml:"PreActionShowUI" json:"pre_action_show_ui,omitempty"`
	PreAction              PreActionLocks          `xml:"PreAction" json:"pre_action,omitempty"`
	HasRunningMessage      bool                    `xml:"HasRunningMessage" json:"has_running_message,omitempty"`
	RunningMessage         RunningMessageLocks     `xml:"RunningMessage" json:"running_message,omitempty"`
	TimeRange              bool                    `xml:"TimeRange" json:"time_range,omitempty"`
	StartDateTimeOffset    bool                    `xml:"StartDateTimeOffset" json:"start_date_time_offset,omitempty"`
	EndDateTimeOffset      bool                    `xml:"EndDateTimeOffset" json:"end_date_time_offset,omitempty"`
	DayOfWeekConstraint    bool                    `xml:"DayOfWeekConstraint" json:"day_of_week_constraint,omitempty"`
	ActiveUserRequirement  bool                    `xml:"ActiveUserRequirement" json:"active_user_requirement,omitempty"`
	ActiveUserType         bool                    `xml:"ActiveUserType" json:"active_user_type,omitempty"`
	Whose                  bool                    `xml:"Whose" json:"whose,omitempty"`
	PreActionCacheDownload bool                    `xml:"PreActionCacheDownload" json:"pre_action_cache_download,omitempty"`
	Reapply                bool                    `xml:"Reapply" json:"reapply,omitempty"`
	ReapplyLimit           bool                    `xml:"ReapplyLimit" json:"reapply_limit,omitempty"`
	RetryCount             bool                    `xml:"RetryCount" json:"retry_count,omitempty"`
	RetryWait              bool                    `xml:"RetryWait" json:"retry_wait,omitempty"`
	TemporalDistribution   bool                    `xml:"TemporalDistribution" json:"temporal_distribution,omitempty"`
	ContinueOnErrors       bool                    `xml:"ContinueOnErrors" json:"continue_on_errors,omitempty"`
	PostActionBehavior     PostActionBehaviorLocks `xml:"PostActionBehavior" json:"post_action_behavior,omitempty"`
	IsOffer                bool                    `xml:"IsOffer" json:"is_offer,omitempty"`
	AnnounceOffer          bool                    `xml:"AnnounceOffer" json:"announce_offer,omitempty"`
	OfferCategory          bool                    `xml:"OfferCategory" json:"offer_category,omitempty"`
	OfferDescriptionHTML   bool                    `xml:"OfferDescriptionHTML" json:"offer_description_html,omitempty"`
}

// PreActionLocks represents pre-action locks
type PreActionLocks struct {
	Text             bool `xml:"Text" json:"text,omitempty"`
	AskToSaveWork    bool `xml:"AskToSaveWork" json:"ask_to_save_work,omitempty"`
	ShowActionButton bool `xml:"ShowActionButton" json:"show_action_button,omitempty"`
	ShowCancelButton bool `xml:"ShowCancelButton" json:"show_cancel_button,omitempty"`
	DeadlineBehavior bool `xml:"DeadlineBehavior" json:"deadline_behavior,omitempty"`
	ShowConfirmation bool `xml:"ShowConfirmation" json:"show_confirmation,omitempty"`
}

// RunningMessageLocks represents running message locks
type RunningMessageLocks struct {
	Text bool `xml:"Text" json:"text,omitempty"`
}

// PostActionBehaviorLocks represents post action behavior locks
type PostActionBehaviorLocks struct {
	Behavior    bool `xml:"Behavior" json:"behavior,omitempty"`
	AllowCancel bool `xml:"AllowCancel" json:"allow_cancel,omitempty"`
	Deadline    bool `xml:"Deadline" json:"deadline,omitempty"`
	Title       bool `xml:"Title" json:"title,omitempty"`
	Text        bool `xml:"Text" json:"text,omitempty"`
}

// ActionTarget represents action target
type ActionTarget struct {
	AllComputers bool `xml:"AllComputers" json:"all_computers,omitempty"`
	ComputerID   int  `xml:"ComputerID" json:"computer_id,omitempty"`
}

// ToAction converts ActionDetail to Action model
func (ad *ActionDetail) ToAction(id int, resource string) *Action {
	return &Action{
		ID:              id,
		Resource:        resource,
		Name:            ad.Title,
		Title:           ad.Title,
		Relevance:       ad.Relevance,
		ActionScript:    ad.ActionScript.Content,
		SuccessCriteria: ad.SuccessCriteria,
		Settings:        ad.Settings,
		SettingsLocks:   ad.SettingsLocks,
		Target:          ad.Target,
		IsUrgent:        ad.IsUrgent,
	}
}

// ToActionFromList converts list Action to detailed Action model
func (a *Action) ToActionFromList() *Action {
	return &Action{
		ID:           a.ID,
		Resource:     a.Resource,
		Name:         a.Name,
		Title:        a.Name, // Use Name as Title for list items
		LastModified: a.LastModified,
	}
}
