/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserCreate Payload for creating a new user
type UserCreate struct {
	// The user's first name (optional)
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name (optional)
	LastName *string `json:"lastName,omitempty"`
	// Username (unique per tenant).
	Username string `json:"username"`
	// Email address
	Email string `json:"email"`
	// Password to apply to the user
	Password string `json:"password"`
	// Array of objects with id of the role(s) to assign to the user.
	Roles []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"roles"`
	// Receive Notifications?
	ReceiveNotifications *bool `json:"receiveNotifications,omitempty"`
	// Linux Username, user settings for provisioning
	LinuxUsername *string `json:"linuxUsername,omitempty"`
	// Linux Password, user settings for provisioning
	LinuxPassword *string `json:"linuxPassword,omitempty"`
	// Linux SSH Key, user settings for provisioning
	LinuxKeyPairId *int64 `json:"linuxKeyPairId,omitempty"`
	// Windows Username, user settings for provisioning
	WindowsUsername *string `json:"windowsUsername,omitempty"`
	// Windows Password, user settings for provisioning
	WindowsPassword *string `json:"windowsPassword,omitempty"`
}

// NewUserCreate instantiates a new UserCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreate(username string, email string, password string, roles []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, ) *UserCreate {
	this := UserCreate{}
	this.Username = username
	this.Email = email
	this.Password = password
	this.Roles = roles
	var receiveNotifications bool = true
	this.ReceiveNotifications = &receiveNotifications
	return &this
}

// NewUserCreateWithDefaults instantiates a new UserCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateWithDefaults() *UserCreate {
	this := UserCreate{}
	var receiveNotifications bool = true
	this.ReceiveNotifications = &receiveNotifications
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserCreate) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserCreate) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserCreate) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserCreate) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserCreate) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserCreate) SetLastName(v string) {
	o.LastName = &v
}

// GetUsername returns the Username field value
func (o *UserCreate) GetUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserCreate) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *UserCreate) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserCreate) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *UserCreate) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserCreate) SetPassword(v string) {
	o.Password = v
}

// GetRoles returns the Roles field value
func (o *UserCreate) GetRoles() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil  {
		var ret []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetRolesOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *UserCreate) SetRoles(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Roles = v
}

// GetReceiveNotifications returns the ReceiveNotifications field value if set, zero value otherwise.
func (o *UserCreate) GetReceiveNotifications() bool {
	if o == nil || o.ReceiveNotifications == nil {
		var ret bool
		return ret
	}
	return *o.ReceiveNotifications
}

// GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetReceiveNotificationsOk() (*bool, bool) {
	if o == nil || o.ReceiveNotifications == nil {
		return nil, false
	}
	return o.ReceiveNotifications, true
}

// HasReceiveNotifications returns a boolean if a field has been set.
func (o *UserCreate) HasReceiveNotifications() bool {
	if o != nil && o.ReceiveNotifications != nil {
		return true
	}

	return false
}

// SetReceiveNotifications gets a reference to the given bool and assigns it to the ReceiveNotifications field.
func (o *UserCreate) SetReceiveNotifications(v bool) {
	o.ReceiveNotifications = &v
}

// GetLinuxUsername returns the LinuxUsername field value if set, zero value otherwise.
func (o *UserCreate) GetLinuxUsername() string {
	if o == nil || o.LinuxUsername == nil {
		var ret string
		return ret
	}
	return *o.LinuxUsername
}

// GetLinuxUsernameOk returns a tuple with the LinuxUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetLinuxUsernameOk() (*string, bool) {
	if o == nil || o.LinuxUsername == nil {
		return nil, false
	}
	return o.LinuxUsername, true
}

// HasLinuxUsername returns a boolean if a field has been set.
func (o *UserCreate) HasLinuxUsername() bool {
	if o != nil && o.LinuxUsername != nil {
		return true
	}

	return false
}

// SetLinuxUsername gets a reference to the given string and assigns it to the LinuxUsername field.
func (o *UserCreate) SetLinuxUsername(v string) {
	o.LinuxUsername = &v
}

// GetLinuxPassword returns the LinuxPassword field value if set, zero value otherwise.
func (o *UserCreate) GetLinuxPassword() string {
	if o == nil || o.LinuxPassword == nil {
		var ret string
		return ret
	}
	return *o.LinuxPassword
}

// GetLinuxPasswordOk returns a tuple with the LinuxPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetLinuxPasswordOk() (*string, bool) {
	if o == nil || o.LinuxPassword == nil {
		return nil, false
	}
	return o.LinuxPassword, true
}

// HasLinuxPassword returns a boolean if a field has been set.
func (o *UserCreate) HasLinuxPassword() bool {
	if o != nil && o.LinuxPassword != nil {
		return true
	}

	return false
}

// SetLinuxPassword gets a reference to the given string and assigns it to the LinuxPassword field.
func (o *UserCreate) SetLinuxPassword(v string) {
	o.LinuxPassword = &v
}

// GetLinuxKeyPairId returns the LinuxKeyPairId field value if set, zero value otherwise.
func (o *UserCreate) GetLinuxKeyPairId() int64 {
	if o == nil || o.LinuxKeyPairId == nil {
		var ret int64
		return ret
	}
	return *o.LinuxKeyPairId
}

// GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetLinuxKeyPairIdOk() (*int64, bool) {
	if o == nil || o.LinuxKeyPairId == nil {
		return nil, false
	}
	return o.LinuxKeyPairId, true
}

// HasLinuxKeyPairId returns a boolean if a field has been set.
func (o *UserCreate) HasLinuxKeyPairId() bool {
	if o != nil && o.LinuxKeyPairId != nil {
		return true
	}

	return false
}

// SetLinuxKeyPairId gets a reference to the given int64 and assigns it to the LinuxKeyPairId field.
func (o *UserCreate) SetLinuxKeyPairId(v int64) {
	o.LinuxKeyPairId = &v
}

// GetWindowsUsername returns the WindowsUsername field value if set, zero value otherwise.
func (o *UserCreate) GetWindowsUsername() string {
	if o == nil || o.WindowsUsername == nil {
		var ret string
		return ret
	}
	return *o.WindowsUsername
}

// GetWindowsUsernameOk returns a tuple with the WindowsUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetWindowsUsernameOk() (*string, bool) {
	if o == nil || o.WindowsUsername == nil {
		return nil, false
	}
	return o.WindowsUsername, true
}

// HasWindowsUsername returns a boolean if a field has been set.
func (o *UserCreate) HasWindowsUsername() bool {
	if o != nil && o.WindowsUsername != nil {
		return true
	}

	return false
}

// SetWindowsUsername gets a reference to the given string and assigns it to the WindowsUsername field.
func (o *UserCreate) SetWindowsUsername(v string) {
	o.WindowsUsername = &v
}

// GetWindowsPassword returns the WindowsPassword field value if set, zero value otherwise.
func (o *UserCreate) GetWindowsPassword() string {
	if o == nil || o.WindowsPassword == nil {
		var ret string
		return ret
	}
	return *o.WindowsPassword
}

// GetWindowsPasswordOk returns a tuple with the WindowsPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetWindowsPasswordOk() (*string, bool) {
	if o == nil || o.WindowsPassword == nil {
		return nil, false
	}
	return o.WindowsPassword, true
}

// HasWindowsPassword returns a boolean if a field has been set.
func (o *UserCreate) HasWindowsPassword() bool {
	if o != nil && o.WindowsPassword != nil {
		return true
	}

	return false
}

// SetWindowsPassword gets a reference to the given string and assigns it to the WindowsPassword field.
func (o *UserCreate) SetWindowsPassword(v string) {
	o.WindowsPassword = &v
}

func (o UserCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	if o.ReceiveNotifications != nil {
		toSerialize["receiveNotifications"] = o.ReceiveNotifications
	}
	if o.LinuxUsername != nil {
		toSerialize["linuxUsername"] = o.LinuxUsername
	}
	if o.LinuxPassword != nil {
		toSerialize["linuxPassword"] = o.LinuxPassword
	}
	if o.LinuxKeyPairId != nil {
		toSerialize["linuxKeyPairId"] = o.LinuxKeyPairId
	}
	if o.WindowsUsername != nil {
		toSerialize["windowsUsername"] = o.WindowsUsername
	}
	if o.WindowsPassword != nil {
		toSerialize["windowsPassword"] = o.WindowsPassword
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreate struct {
	value *UserCreate
	isSet bool
}

func (v NullableUserCreate) Get() *UserCreate {
	return v.value
}

func (v *NullableUserCreate) Set(val *UserCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreate(val *UserCreate) *NullableUserCreate {
	return &NullableUserCreate{value: val, isSet: true}
}

func (v NullableUserCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


