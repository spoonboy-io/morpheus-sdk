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
	"time"
)

// User struct for User
type User struct {
	Id *int64 `json:"id,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	Username *string `json:"username,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ReceiveNotifications *bool `json:"receiveNotifications,omitempty"`
	IsUsing2FA *bool `json:"isUsing2FA,omitempty"`
	AccountExpired *bool `json:"accountExpired,omitempty"`
	AccountLocked *bool `json:"accountLocked,omitempty"`
	PasswordExpired *bool `json:"passwordExpired,omitempty"`
	LoginCount *int64 `json:"loginCount,omitempty"`
	LoginAttempts *int64 `json:"loginAttempts,omitempty"`
	LastLoginDate *time.Time `json:"lastLoginDate,omitempty"`
	Roles *[]UserRoles `json:"roles,omitempty"`
	Account *InlineResponse20040AppDeployInstance `json:"account,omitempty"`
	LinuxUsername NullableString `json:"linuxUsername,omitempty"`
	LinuxPassword NullableString `json:"linuxPassword,omitempty"`
	LinuxKeyPairId NullableInt64 `json:"linuxKeyPairId,omitempty"`
	WindowsUsername NullableString `json:"windowsUsername,omitempty"`
	WindowsPassword NullableString `json:"windowsPassword,omitempty"`
	DefaultPersona *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"defaultPersona,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Access NullableUserAccess `json:"access,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *User) SetId(v int64) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *User) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *User) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *User) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *User) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *User) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *User) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *User) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *User) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *User) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetReceiveNotifications returns the ReceiveNotifications field value if set, zero value otherwise.
func (o *User) GetReceiveNotifications() bool {
	if o == nil || o.ReceiveNotifications == nil {
		var ret bool
		return ret
	}
	return *o.ReceiveNotifications
}

// GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetReceiveNotificationsOk() (*bool, bool) {
	if o == nil || o.ReceiveNotifications == nil {
		return nil, false
	}
	return o.ReceiveNotifications, true
}

// HasReceiveNotifications returns a boolean if a field has been set.
func (o *User) HasReceiveNotifications() bool {
	if o != nil && o.ReceiveNotifications != nil {
		return true
	}

	return false
}

// SetReceiveNotifications gets a reference to the given bool and assigns it to the ReceiveNotifications field.
func (o *User) SetReceiveNotifications(v bool) {
	o.ReceiveNotifications = &v
}

// GetIsUsing2FA returns the IsUsing2FA field value if set, zero value otherwise.
func (o *User) GetIsUsing2FA() bool {
	if o == nil || o.IsUsing2FA == nil {
		var ret bool
		return ret
	}
	return *o.IsUsing2FA
}

// GetIsUsing2FAOk returns a tuple with the IsUsing2FA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsUsing2FAOk() (*bool, bool) {
	if o == nil || o.IsUsing2FA == nil {
		return nil, false
	}
	return o.IsUsing2FA, true
}

// HasIsUsing2FA returns a boolean if a field has been set.
func (o *User) HasIsUsing2FA() bool {
	if o != nil && o.IsUsing2FA != nil {
		return true
	}

	return false
}

// SetIsUsing2FA gets a reference to the given bool and assigns it to the IsUsing2FA field.
func (o *User) SetIsUsing2FA(v bool) {
	o.IsUsing2FA = &v
}

// GetAccountExpired returns the AccountExpired field value if set, zero value otherwise.
func (o *User) GetAccountExpired() bool {
	if o == nil || o.AccountExpired == nil {
		var ret bool
		return ret
	}
	return *o.AccountExpired
}

// GetAccountExpiredOk returns a tuple with the AccountExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountExpiredOk() (*bool, bool) {
	if o == nil || o.AccountExpired == nil {
		return nil, false
	}
	return o.AccountExpired, true
}

// HasAccountExpired returns a boolean if a field has been set.
func (o *User) HasAccountExpired() bool {
	if o != nil && o.AccountExpired != nil {
		return true
	}

	return false
}

// SetAccountExpired gets a reference to the given bool and assigns it to the AccountExpired field.
func (o *User) SetAccountExpired(v bool) {
	o.AccountExpired = &v
}

// GetAccountLocked returns the AccountLocked field value if set, zero value otherwise.
func (o *User) GetAccountLocked() bool {
	if o == nil || o.AccountLocked == nil {
		var ret bool
		return ret
	}
	return *o.AccountLocked
}

// GetAccountLockedOk returns a tuple with the AccountLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountLockedOk() (*bool, bool) {
	if o == nil || o.AccountLocked == nil {
		return nil, false
	}
	return o.AccountLocked, true
}

// HasAccountLocked returns a boolean if a field has been set.
func (o *User) HasAccountLocked() bool {
	if o != nil && o.AccountLocked != nil {
		return true
	}

	return false
}

// SetAccountLocked gets a reference to the given bool and assigns it to the AccountLocked field.
func (o *User) SetAccountLocked(v bool) {
	o.AccountLocked = &v
}

// GetPasswordExpired returns the PasswordExpired field value if set, zero value otherwise.
func (o *User) GetPasswordExpired() bool {
	if o == nil || o.PasswordExpired == nil {
		var ret bool
		return ret
	}
	return *o.PasswordExpired
}

// GetPasswordExpiredOk returns a tuple with the PasswordExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordExpiredOk() (*bool, bool) {
	if o == nil || o.PasswordExpired == nil {
		return nil, false
	}
	return o.PasswordExpired, true
}

// HasPasswordExpired returns a boolean if a field has been set.
func (o *User) HasPasswordExpired() bool {
	if o != nil && o.PasswordExpired != nil {
		return true
	}

	return false
}

// SetPasswordExpired gets a reference to the given bool and assigns it to the PasswordExpired field.
func (o *User) SetPasswordExpired(v bool) {
	o.PasswordExpired = &v
}

// GetLoginCount returns the LoginCount field value if set, zero value otherwise.
func (o *User) GetLoginCount() int64 {
	if o == nil || o.LoginCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginCount
}

// GetLoginCountOk returns a tuple with the LoginCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLoginCountOk() (*int64, bool) {
	if o == nil || o.LoginCount == nil {
		return nil, false
	}
	return o.LoginCount, true
}

// HasLoginCount returns a boolean if a field has been set.
func (o *User) HasLoginCount() bool {
	if o != nil && o.LoginCount != nil {
		return true
	}

	return false
}

// SetLoginCount gets a reference to the given int64 and assigns it to the LoginCount field.
func (o *User) SetLoginCount(v int64) {
	o.LoginCount = &v
}

// GetLoginAttempts returns the LoginAttempts field value if set, zero value otherwise.
func (o *User) GetLoginAttempts() int64 {
	if o == nil || o.LoginAttempts == nil {
		var ret int64
		return ret
	}
	return *o.LoginAttempts
}

// GetLoginAttemptsOk returns a tuple with the LoginAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLoginAttemptsOk() (*int64, bool) {
	if o == nil || o.LoginAttempts == nil {
		return nil, false
	}
	return o.LoginAttempts, true
}

// HasLoginAttempts returns a boolean if a field has been set.
func (o *User) HasLoginAttempts() bool {
	if o != nil && o.LoginAttempts != nil {
		return true
	}

	return false
}

// SetLoginAttempts gets a reference to the given int64 and assigns it to the LoginAttempts field.
func (o *User) SetLoginAttempts(v int64) {
	o.LoginAttempts = &v
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise.
func (o *User) GetLastLoginDate() time.Time {
	if o == nil || o.LastLoginDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil || o.LastLoginDate == nil {
		return nil, false
	}
	return o.LastLoginDate, true
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *User) HasLastLoginDate() bool {
	if o != nil && o.LastLoginDate != nil {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given time.Time and assigns it to the LastLoginDate field.
func (o *User) SetLastLoginDate(v time.Time) {
	o.LastLoginDate = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *User) GetRoles() []UserRoles {
	if o == nil || o.Roles == nil {
		var ret []UserRoles
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRolesOk() (*[]UserRoles, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *User) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []UserRoles and assigns it to the Roles field.
func (o *User) SetRoles(v []UserRoles) {
	o.Roles = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *User) GetAccount() InlineResponse20040AppDeployInstance {
	if o == nil || o.Account == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *User) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Account field.
func (o *User) SetAccount(v InlineResponse20040AppDeployInstance) {
	o.Account = &v
}

// GetLinuxUsername returns the LinuxUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLinuxUsername() string {
	if o == nil || o.LinuxUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.LinuxUsername.Get()
}

// GetLinuxUsernameOk returns a tuple with the LinuxUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLinuxUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinuxUsername.Get(), o.LinuxUsername.IsSet()
}

// HasLinuxUsername returns a boolean if a field has been set.
func (o *User) HasLinuxUsername() bool {
	if o != nil && o.LinuxUsername.IsSet() {
		return true
	}

	return false
}

// SetLinuxUsername gets a reference to the given NullableString and assigns it to the LinuxUsername field.
func (o *User) SetLinuxUsername(v string) {
	o.LinuxUsername.Set(&v)
}
// SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil
func (o *User) SetLinuxUsernameNil() {
	o.LinuxUsername.Set(nil)
}

// UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
func (o *User) UnsetLinuxUsername() {
	o.LinuxUsername.Unset()
}

// GetLinuxPassword returns the LinuxPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLinuxPassword() string {
	if o == nil || o.LinuxPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.LinuxPassword.Get()
}

// GetLinuxPasswordOk returns a tuple with the LinuxPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLinuxPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinuxPassword.Get(), o.LinuxPassword.IsSet()
}

// HasLinuxPassword returns a boolean if a field has been set.
func (o *User) HasLinuxPassword() bool {
	if o != nil && o.LinuxPassword.IsSet() {
		return true
	}

	return false
}

// SetLinuxPassword gets a reference to the given NullableString and assigns it to the LinuxPassword field.
func (o *User) SetLinuxPassword(v string) {
	o.LinuxPassword.Set(&v)
}
// SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil
func (o *User) SetLinuxPasswordNil() {
	o.LinuxPassword.Set(nil)
}

// UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
func (o *User) UnsetLinuxPassword() {
	o.LinuxPassword.Unset()
}

// GetLinuxKeyPairId returns the LinuxKeyPairId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLinuxKeyPairId() int64 {
	if o == nil || o.LinuxKeyPairId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LinuxKeyPairId.Get()
}

// GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLinuxKeyPairIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinuxKeyPairId.Get(), o.LinuxKeyPairId.IsSet()
}

// HasLinuxKeyPairId returns a boolean if a field has been set.
func (o *User) HasLinuxKeyPairId() bool {
	if o != nil && o.LinuxKeyPairId.IsSet() {
		return true
	}

	return false
}

// SetLinuxKeyPairId gets a reference to the given NullableInt64 and assigns it to the LinuxKeyPairId field.
func (o *User) SetLinuxKeyPairId(v int64) {
	o.LinuxKeyPairId.Set(&v)
}
// SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil
func (o *User) SetLinuxKeyPairIdNil() {
	o.LinuxKeyPairId.Set(nil)
}

// UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
func (o *User) UnsetLinuxKeyPairId() {
	o.LinuxKeyPairId.Unset()
}

// GetWindowsUsername returns the WindowsUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetWindowsUsername() string {
	if o == nil || o.WindowsUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.WindowsUsername.Get()
}

// GetWindowsUsernameOk returns a tuple with the WindowsUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetWindowsUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WindowsUsername.Get(), o.WindowsUsername.IsSet()
}

// HasWindowsUsername returns a boolean if a field has been set.
func (o *User) HasWindowsUsername() bool {
	if o != nil && o.WindowsUsername.IsSet() {
		return true
	}

	return false
}

// SetWindowsUsername gets a reference to the given NullableString and assigns it to the WindowsUsername field.
func (o *User) SetWindowsUsername(v string) {
	o.WindowsUsername.Set(&v)
}
// SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil
func (o *User) SetWindowsUsernameNil() {
	o.WindowsUsername.Set(nil)
}

// UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
func (o *User) UnsetWindowsUsername() {
	o.WindowsUsername.Unset()
}

// GetWindowsPassword returns the WindowsPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetWindowsPassword() string {
	if o == nil || o.WindowsPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.WindowsPassword.Get()
}

// GetWindowsPasswordOk returns a tuple with the WindowsPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetWindowsPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WindowsPassword.Get(), o.WindowsPassword.IsSet()
}

// HasWindowsPassword returns a boolean if a field has been set.
func (o *User) HasWindowsPassword() bool {
	if o != nil && o.WindowsPassword.IsSet() {
		return true
	}

	return false
}

// SetWindowsPassword gets a reference to the given NullableString and assigns it to the WindowsPassword field.
func (o *User) SetWindowsPassword(v string) {
	o.WindowsPassword.Set(&v)
}
// SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil
func (o *User) SetWindowsPasswordNil() {
	o.WindowsPassword.Set(nil)
}

// UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
func (o *User) UnsetWindowsPassword() {
	o.WindowsPassword.Unset()
}

// GetDefaultPersona returns the DefaultPersona field value if set, zero value otherwise.
func (o *User) GetDefaultPersona() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.DefaultPersona == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.DefaultPersona
}

// GetDefaultPersonaOk returns a tuple with the DefaultPersona field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDefaultPersonaOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.DefaultPersona == nil {
		return nil, false
	}
	return o.DefaultPersona, true
}

// HasDefaultPersona returns a boolean if a field has been set.
func (o *User) HasDefaultPersona() bool {
	if o != nil && o.DefaultPersona != nil {
		return true
	}

	return false
}

// SetDefaultPersona gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the DefaultPersona field.
func (o *User) SetDefaultPersona(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.DefaultPersona = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *User) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *User) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *User) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *User) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *User) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *User) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAccess returns the Access field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetAccess() UserAccess {
	if o == nil || o.Access.Get() == nil {
		var ret UserAccess
		return ret
	}
	return *o.Access.Get()
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetAccessOk() (*UserAccess, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Access.Get(), o.Access.IsSet()
}

// HasAccess returns a boolean if a field has been set.
func (o *User) HasAccess() bool {
	if o != nil && o.Access.IsSet() {
		return true
	}

	return false
}

// SetAccess gets a reference to the given NullableUserAccess and assigns it to the Access field.
func (o *User) SetAccess(v UserAccess) {
	o.Access.Set(&v)
}
// SetAccessNil sets the value for Access to be an explicit nil
func (o *User) SetAccessNil() {
	o.Access.Set(nil)
}

// UnsetAccess ensures that no value is present for Access, not even an explicit nil
func (o *User) UnsetAccess() {
	o.Access.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ReceiveNotifications != nil {
		toSerialize["receiveNotifications"] = o.ReceiveNotifications
	}
	if o.IsUsing2FA != nil {
		toSerialize["isUsing2FA"] = o.IsUsing2FA
	}
	if o.AccountExpired != nil {
		toSerialize["accountExpired"] = o.AccountExpired
	}
	if o.AccountLocked != nil {
		toSerialize["accountLocked"] = o.AccountLocked
	}
	if o.PasswordExpired != nil {
		toSerialize["passwordExpired"] = o.PasswordExpired
	}
	if o.LoginCount != nil {
		toSerialize["loginCount"] = o.LoginCount
	}
	if o.LoginAttempts != nil {
		toSerialize["loginAttempts"] = o.LoginAttempts
	}
	if o.LastLoginDate != nil {
		toSerialize["lastLoginDate"] = o.LastLoginDate
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.LinuxUsername.IsSet() {
		toSerialize["linuxUsername"] = o.LinuxUsername.Get()
	}
	if o.LinuxPassword.IsSet() {
		toSerialize["linuxPassword"] = o.LinuxPassword.Get()
	}
	if o.LinuxKeyPairId.IsSet() {
		toSerialize["linuxKeyPairId"] = o.LinuxKeyPairId.Get()
	}
	if o.WindowsUsername.IsSet() {
		toSerialize["windowsUsername"] = o.WindowsUsername.Get()
	}
	if o.WindowsPassword.IsSet() {
		toSerialize["windowsPassword"] = o.WindowsPassword.Get()
	}
	if o.DefaultPersona != nil {
		toSerialize["defaultPersona"] = o.DefaultPersona
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Access.IsSet() {
		toSerialize["access"] = o.Access.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

