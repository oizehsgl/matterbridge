// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RequiredResourceAccess undocumented
type RequiredResourceAccess struct {
	// Object is the base model of RequiredResourceAccess
	Object
	// ResourceAppID undocumented
	ResourceAppID *string `json:"resourceAppId,omitempty"`
	// ResourceAccess undocumented
	ResourceAccess []ResourceAccess `json:"resourceAccess,omitempty"`
}
