// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LocateDeviceActionResult undocumented
type LocateDeviceActionResult struct {
	// DeviceActionResult is the base model of LocateDeviceActionResult
	DeviceActionResult
	// DeviceLocation device location
	DeviceLocation *DeviceGeoLocation `json:"deviceLocation,omitempty"`
}