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

// CheckSqlConfig struct for CheckSqlConfig
type CheckSqlConfig struct {
	// Hostname or IP address of the database
	DbHost string `json:"dbHost"`
	// Database Port (defaults to default port of DB type selected)
	DbPort string `json:"dbPort"`
	// Database username
	DbUser string `json:"dbUser"`
	// Database password, (all check data is encrypted inside the database)
	DbPassword string `json:"dbPassword"`
	// Database password hash
	DbPasswordHash *string `json:"dbPasswordHash,omitempty"`
	// Database name you would like to connect to
	DbName string `json:"dbName"`
	// Query to test
	DbQuery string `json:"dbQuery"`
	// Can be set to `lt` (less than), `gt` (greater than), `equal` (Equal to) for comparison
	CheckOperator *string `json:"checkOperator,omitempty"`
	CheckResult *float32 `json:"checkResult,omitempty"`
	CheckUser *string `json:"checkUser,omitempty"`
	TextCheckOn *string `json:"textCheckOn,omitempty"`
	CheckPassword *string `json:"checkPassword,omitempty"`
	WebTextMatch *string `json:"webTextMatch,omitempty"`
	CheckPasswordHash *string `json:"checkPasswordHash,omitempty"`
	// Set to on to turn on tunneling
	TunnelOn *string `json:"tunnelOn,omitempty"`
	// Hostname or IP address of the proxy host
	SshHost *string `json:"sshHost,omitempty"`
	// Port for SSH on the proxy host, defaults to 22
	SshPort *int64 `json:"sshPort,omitempty"`
	// SSH user on the proxy host to login as
	SshUser *string `json:"sshUser,omitempty"`
	// Password for user, if not using key based authentication
	SshPassword *string `json:"sshPassword,omitempty"`
}

// NewCheckSqlConfig instantiates a new CheckSqlConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckSqlConfig(dbHost string, dbPort string, dbUser string, dbPassword string, dbName string, dbQuery string, ) *CheckSqlConfig {
	this := CheckSqlConfig{}
	this.DbHost = dbHost
	this.DbPort = dbPort
	this.DbUser = dbUser
	this.DbPassword = dbPassword
	this.DbName = dbName
	this.DbQuery = dbQuery
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// NewCheckSqlConfigWithDefaults instantiates a new CheckSqlConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckSqlConfigWithDefaults() *CheckSqlConfig {
	this := CheckSqlConfig{}
	var tunnelOn string = "off"
	this.TunnelOn = &tunnelOn
	return &this
}

// GetDbHost returns the DbHost field value
func (o *CheckSqlConfig) GetDbHost() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbHost
}

// GetDbHostOk returns a tuple with the DbHost field value
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbHost, true
}

// SetDbHost sets field value
func (o *CheckSqlConfig) SetDbHost(v string) {
	o.DbHost = v
}

// GetDbPort returns the DbPort field value
func (o *CheckSqlConfig) GetDbPort() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbPort
}

// GetDbPortOk returns a tuple with the DbPort field value
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbPort, true
}

// SetDbPort sets field value
func (o *CheckSqlConfig) SetDbPort(v string) {
	o.DbPort = v
}

// GetDbUser returns the DbUser field value
func (o *CheckSqlConfig) GetDbUser() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbUser
}

// GetDbUserOk returns a tuple with the DbUser field value
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbUser, true
}

// SetDbUser sets field value
func (o *CheckSqlConfig) SetDbUser(v string) {
	o.DbUser = v
}

// GetDbPassword returns the DbPassword field value
func (o *CheckSqlConfig) GetDbPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbPassword
}

// GetDbPasswordOk returns a tuple with the DbPassword field value
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbPassword, true
}

// SetDbPassword sets field value
func (o *CheckSqlConfig) SetDbPassword(v string) {
	o.DbPassword = v
}

// GetDbPasswordHash returns the DbPasswordHash field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetDbPasswordHash() string {
	if o == nil || o.DbPasswordHash == nil {
		var ret string
		return ret
	}
	return *o.DbPasswordHash
}

// GetDbPasswordHashOk returns a tuple with the DbPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbPasswordHashOk() (*string, bool) {
	if o == nil || o.DbPasswordHash == nil {
		return nil, false
	}
	return o.DbPasswordHash, true
}

// HasDbPasswordHash returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasDbPasswordHash() bool {
	if o != nil && o.DbPasswordHash != nil {
		return true
	}

	return false
}

// SetDbPasswordHash gets a reference to the given string and assigns it to the DbPasswordHash field.
func (o *CheckSqlConfig) SetDbPasswordHash(v string) {
	o.DbPasswordHash = &v
}

// GetDbName returns the DbName field value
func (o *CheckSqlConfig) GetDbName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *CheckSqlConfig) SetDbName(v string) {
	o.DbName = v
}

// GetDbQuery returns the DbQuery field value
func (o *CheckSqlConfig) GetDbQuery() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbQuery
}

// GetDbQueryOk returns a tuple with the DbQuery field value
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetDbQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbQuery, true
}

// SetDbQuery sets field value
func (o *CheckSqlConfig) SetDbQuery(v string) {
	o.DbQuery = v
}

// GetCheckOperator returns the CheckOperator field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetCheckOperator() string {
	if o == nil || o.CheckOperator == nil {
		var ret string
		return ret
	}
	return *o.CheckOperator
}

// GetCheckOperatorOk returns a tuple with the CheckOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetCheckOperatorOk() (*string, bool) {
	if o == nil || o.CheckOperator == nil {
		return nil, false
	}
	return o.CheckOperator, true
}

// HasCheckOperator returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasCheckOperator() bool {
	if o != nil && o.CheckOperator != nil {
		return true
	}

	return false
}

// SetCheckOperator gets a reference to the given string and assigns it to the CheckOperator field.
func (o *CheckSqlConfig) SetCheckOperator(v string) {
	o.CheckOperator = &v
}

// GetCheckResult returns the CheckResult field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetCheckResult() float32 {
	if o == nil || o.CheckResult == nil {
		var ret float32
		return ret
	}
	return *o.CheckResult
}

// GetCheckResultOk returns a tuple with the CheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetCheckResultOk() (*float32, bool) {
	if o == nil || o.CheckResult == nil {
		return nil, false
	}
	return o.CheckResult, true
}

// HasCheckResult returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasCheckResult() bool {
	if o != nil && o.CheckResult != nil {
		return true
	}

	return false
}

// SetCheckResult gets a reference to the given float32 and assigns it to the CheckResult field.
func (o *CheckSqlConfig) SetCheckResult(v float32) {
	o.CheckResult = &v
}

// GetCheckUser returns the CheckUser field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetCheckUser() string {
	if o == nil || o.CheckUser == nil {
		var ret string
		return ret
	}
	return *o.CheckUser
}

// GetCheckUserOk returns a tuple with the CheckUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetCheckUserOk() (*string, bool) {
	if o == nil || o.CheckUser == nil {
		return nil, false
	}
	return o.CheckUser, true
}

// HasCheckUser returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasCheckUser() bool {
	if o != nil && o.CheckUser != nil {
		return true
	}

	return false
}

// SetCheckUser gets a reference to the given string and assigns it to the CheckUser field.
func (o *CheckSqlConfig) SetCheckUser(v string) {
	o.CheckUser = &v
}

// GetTextCheckOn returns the TextCheckOn field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetTextCheckOn() string {
	if o == nil || o.TextCheckOn == nil {
		var ret string
		return ret
	}
	return *o.TextCheckOn
}

// GetTextCheckOnOk returns a tuple with the TextCheckOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetTextCheckOnOk() (*string, bool) {
	if o == nil || o.TextCheckOn == nil {
		return nil, false
	}
	return o.TextCheckOn, true
}

// HasTextCheckOn returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasTextCheckOn() bool {
	if o != nil && o.TextCheckOn != nil {
		return true
	}

	return false
}

// SetTextCheckOn gets a reference to the given string and assigns it to the TextCheckOn field.
func (o *CheckSqlConfig) SetTextCheckOn(v string) {
	o.TextCheckOn = &v
}

// GetCheckPassword returns the CheckPassword field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetCheckPassword() string {
	if o == nil || o.CheckPassword == nil {
		var ret string
		return ret
	}
	return *o.CheckPassword
}

// GetCheckPasswordOk returns a tuple with the CheckPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetCheckPasswordOk() (*string, bool) {
	if o == nil || o.CheckPassword == nil {
		return nil, false
	}
	return o.CheckPassword, true
}

// HasCheckPassword returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasCheckPassword() bool {
	if o != nil && o.CheckPassword != nil {
		return true
	}

	return false
}

// SetCheckPassword gets a reference to the given string and assigns it to the CheckPassword field.
func (o *CheckSqlConfig) SetCheckPassword(v string) {
	o.CheckPassword = &v
}

// GetWebTextMatch returns the WebTextMatch field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetWebTextMatch() string {
	if o == nil || o.WebTextMatch == nil {
		var ret string
		return ret
	}
	return *o.WebTextMatch
}

// GetWebTextMatchOk returns a tuple with the WebTextMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetWebTextMatchOk() (*string, bool) {
	if o == nil || o.WebTextMatch == nil {
		return nil, false
	}
	return o.WebTextMatch, true
}

// HasWebTextMatch returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasWebTextMatch() bool {
	if o != nil && o.WebTextMatch != nil {
		return true
	}

	return false
}

// SetWebTextMatch gets a reference to the given string and assigns it to the WebTextMatch field.
func (o *CheckSqlConfig) SetWebTextMatch(v string) {
	o.WebTextMatch = &v
}

// GetCheckPasswordHash returns the CheckPasswordHash field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetCheckPasswordHash() string {
	if o == nil || o.CheckPasswordHash == nil {
		var ret string
		return ret
	}
	return *o.CheckPasswordHash
}

// GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetCheckPasswordHashOk() (*string, bool) {
	if o == nil || o.CheckPasswordHash == nil {
		return nil, false
	}
	return o.CheckPasswordHash, true
}

// HasCheckPasswordHash returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasCheckPasswordHash() bool {
	if o != nil && o.CheckPasswordHash != nil {
		return true
	}

	return false
}

// SetCheckPasswordHash gets a reference to the given string and assigns it to the CheckPasswordHash field.
func (o *CheckSqlConfig) SetCheckPasswordHash(v string) {
	o.CheckPasswordHash = &v
}

// GetTunnelOn returns the TunnelOn field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetTunnelOn() string {
	if o == nil || o.TunnelOn == nil {
		var ret string
		return ret
	}
	return *o.TunnelOn
}

// GetTunnelOnOk returns a tuple with the TunnelOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetTunnelOnOk() (*string, bool) {
	if o == nil || o.TunnelOn == nil {
		return nil, false
	}
	return o.TunnelOn, true
}

// HasTunnelOn returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasTunnelOn() bool {
	if o != nil && o.TunnelOn != nil {
		return true
	}

	return false
}

// SetTunnelOn gets a reference to the given string and assigns it to the TunnelOn field.
func (o *CheckSqlConfig) SetTunnelOn(v string) {
	o.TunnelOn = &v
}

// GetSshHost returns the SshHost field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetSshHost() string {
	if o == nil || o.SshHost == nil {
		var ret string
		return ret
	}
	return *o.SshHost
}

// GetSshHostOk returns a tuple with the SshHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetSshHostOk() (*string, bool) {
	if o == nil || o.SshHost == nil {
		return nil, false
	}
	return o.SshHost, true
}

// HasSshHost returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasSshHost() bool {
	if o != nil && o.SshHost != nil {
		return true
	}

	return false
}

// SetSshHost gets a reference to the given string and assigns it to the SshHost field.
func (o *CheckSqlConfig) SetSshHost(v string) {
	o.SshHost = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetSshPort() int64 {
	if o == nil || o.SshPort == nil {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetSshPortOk() (*int64, bool) {
	if o == nil || o.SshPort == nil {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasSshPort() bool {
	if o != nil && o.SshPort != nil {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *CheckSqlConfig) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *CheckSqlConfig) SetSshUser(v string) {
	o.SshUser = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *CheckSqlConfig) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckSqlConfig) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *CheckSqlConfig) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *CheckSqlConfig) SetSshPassword(v string) {
	o.SshPassword = &v
}

func (o CheckSqlConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dbHost"] = o.DbHost
	}
	if true {
		toSerialize["dbPort"] = o.DbPort
	}
	if true {
		toSerialize["dbUser"] = o.DbUser
	}
	if true {
		toSerialize["dbPassword"] = o.DbPassword
	}
	if o.DbPasswordHash != nil {
		toSerialize["dbPasswordHash"] = o.DbPasswordHash
	}
	if true {
		toSerialize["dbName"] = o.DbName
	}
	if true {
		toSerialize["dbQuery"] = o.DbQuery
	}
	if o.CheckOperator != nil {
		toSerialize["checkOperator"] = o.CheckOperator
	}
	if o.CheckResult != nil {
		toSerialize["checkResult"] = o.CheckResult
	}
	if o.CheckUser != nil {
		toSerialize["checkUser"] = o.CheckUser
	}
	if o.TextCheckOn != nil {
		toSerialize["textCheckOn"] = o.TextCheckOn
	}
	if o.CheckPassword != nil {
		toSerialize["checkPassword"] = o.CheckPassword
	}
	if o.WebTextMatch != nil {
		toSerialize["webTextMatch"] = o.WebTextMatch
	}
	if o.CheckPasswordHash != nil {
		toSerialize["checkPasswordHash"] = o.CheckPasswordHash
	}
	if o.TunnelOn != nil {
		toSerialize["tunnelOn"] = o.TunnelOn
	}
	if o.SshHost != nil {
		toSerialize["sshHost"] = o.SshHost
	}
	if o.SshPort != nil {
		toSerialize["sshPort"] = o.SshPort
	}
	if o.SshUser != nil {
		toSerialize["sshUser"] = o.SshUser
	}
	if o.SshPassword != nil {
		toSerialize["sshPassword"] = o.SshPassword
	}
	return json.Marshal(toSerialize)
}

type NullableCheckSqlConfig struct {
	value *CheckSqlConfig
	isSet bool
}

func (v NullableCheckSqlConfig) Get() *CheckSqlConfig {
	return v.value
}

func (v *NullableCheckSqlConfig) Set(val *CheckSqlConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckSqlConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckSqlConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckSqlConfig(val *CheckSqlConfig) *NullableCheckSqlConfig {
	return &NullableCheckSqlConfig{value: val, isSet: true}
}

func (v NullableCheckSqlConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckSqlConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


