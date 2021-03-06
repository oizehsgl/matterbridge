// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Schedule undocumented
type Schedule struct {
	// Entity is the base model of Schedule
	Entity
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
	// ProvisionStatus undocumented
	ProvisionStatus *OperationStatus `json:"provisionStatus,omitempty"`
	// ProvisionStatusCode undocumented
	ProvisionStatusCode *string `json:"provisionStatusCode,omitempty"`
	// WorkforceIntegrationIDs undocumented
	WorkforceIntegrationIDs []string `json:"workforceIntegrationIds,omitempty"`
	// TimeClockEnabled undocumented
	TimeClockEnabled *bool `json:"timeClockEnabled,omitempty"`
	// OpenShiftsEnabled undocumented
	OpenShiftsEnabled *bool `json:"openShiftsEnabled,omitempty"`
	// SwapShiftsRequestsEnabled undocumented
	SwapShiftsRequestsEnabled *bool `json:"swapShiftsRequestsEnabled,omitempty"`
	// OfferShiftRequestsEnabled undocumented
	OfferShiftRequestsEnabled *bool `json:"offerShiftRequestsEnabled,omitempty"`
	// TimeOffRequestsEnabled undocumented
	TimeOffRequestsEnabled *bool `json:"timeOffRequestsEnabled,omitempty"`
	// Shifts undocumented
	Shifts []Shift `json:"shifts,omitempty"`
	// OpenShifts undocumented
	OpenShifts []OpenShift `json:"openShifts,omitempty"`
	// TimesOff undocumented
	TimesOff []TimeOff `json:"timesOff,omitempty"`
	// TimeOffReasons undocumented
	TimeOffReasons []TimeOffReason `json:"timeOffReasons,omitempty"`
	// SchedulingGroups undocumented
	SchedulingGroups []SchedulingGroup `json:"schedulingGroups,omitempty"`
	// SwapShiftsChangeRequests undocumented
	SwapShiftsChangeRequests []SwapShiftsChangeRequestObject `json:"swapShiftsChangeRequests,omitempty"`
	// OpenShiftChangeRequests undocumented
	OpenShiftChangeRequests []OpenShiftChangeRequestObject `json:"openShiftChangeRequests,omitempty"`
	// TimeOffRequests undocumented
	TimeOffRequests []TimeOffRequestObject `json:"timeOffRequests,omitempty"`
}

// ScheduleChangeRequestObject undocumented
type ScheduleChangeRequestObject struct {
	// ChangeTrackedEntity is the base model of ScheduleChangeRequestObject
	ChangeTrackedEntity
	// AssignedTo undocumented
	AssignedTo *ScheduleChangeRequestActor `json:"assignedTo,omitempty"`
	// State undocumented
	State *ScheduleChangeState `json:"state,omitempty"`
	// SenderMessage undocumented
	SenderMessage *string `json:"senderMessage,omitempty"`
	// SenderDateTime undocumented
	SenderDateTime *time.Time `json:"senderDateTime,omitempty"`
	// ManagerActionMessage undocumented
	ManagerActionMessage *string `json:"managerActionMessage,omitempty"`
	// ManagerActionDateTime undocumented
	ManagerActionDateTime *time.Time `json:"managerActionDateTime,omitempty"`
	// SenderUserID undocumented
	SenderUserID *string `json:"senderUserId,omitempty"`
	// ManagerUserID undocumented
	ManagerUserID *string `json:"managerUserId,omitempty"`
}

// ScheduleEntity undocumented
type ScheduleEntity struct {
	// Object is the base model of ScheduleEntity
	Object
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Theme undocumented
	Theme *ScheduleEntityTheme `json:"theme,omitempty"`
}

// ScheduleInformation undocumented
type ScheduleInformation struct {
	// Object is the base model of ScheduleInformation
	Object
	// ScheduleID undocumented
	ScheduleID *string `json:"scheduleId,omitempty"`
	// ScheduleItems undocumented
	ScheduleItems []ScheduleItem `json:"scheduleItems,omitempty"`
	// AvailabilityView undocumented
	AvailabilityView *string `json:"availabilityView,omitempty"`
	// Error undocumented
	Error *FreeBusyError `json:"error,omitempty"`
	// WorkingHours undocumented
	WorkingHours *WorkingHours `json:"workingHours,omitempty"`
}

// ScheduleItem undocumented
type ScheduleItem struct {
	// Object is the base model of ScheduleItem
	Object
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// IsPrivate undocumented
	IsPrivate *bool `json:"isPrivate,omitempty"`
	// Status undocumented
	Status *FreeBusyStatus `json:"status,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// Location undocumented
	Location *string `json:"location,omitempty"`
}
