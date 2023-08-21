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

// ScriptCreate struct for ScriptCreate
type ScriptCreate struct {
	// Script name
	Name string `json:"name"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Script category
	Category *string `json:"category,omitempty"`
	// Version of the script
	ScriptVersion *string `json:"scriptVersion,omitempty"`
	// Phase for the script, provision, start, etc.
	ScriptPhase *string `json:"scriptPhase,omitempty"`
	// Type for the script
	ScriptType *string `json:"scriptType,omitempty"`
	// Script content, that is, the code itself.
	Script *string `json:"script,omitempty"`
	// Run as a specific user.
	RunAsUser *string `json:"runAsUser,omitempty"`
	// Sudo, whether or not to run with sudo.
	SudoUser *bool `json:"sudoUser,omitempty"`
}

// NewScriptCreate instantiates a new ScriptCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptCreate(name string, ) *ScriptCreate {
	this := ScriptCreate{}
	this.Name = name
	var scriptVersion string = "1"
	this.ScriptVersion = &scriptVersion
	var scriptType string = "bash"
	this.ScriptType = &scriptType
	var sudoUser bool = false
	this.SudoUser = &sudoUser
	return &this
}

// NewScriptCreateWithDefaults instantiates a new ScriptCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptCreateWithDefaults() *ScriptCreate {
	this := ScriptCreate{}
	var scriptVersion string = "1"
	this.ScriptVersion = &scriptVersion
	var scriptType string = "bash"
	this.ScriptType = &scriptType
	var sudoUser bool = false
	this.SudoUser = &sudoUser
	return &this
}

// GetName returns the Name field value
func (o *ScriptCreate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScriptCreate) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScriptCreate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScriptCreate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ScriptCreate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ScriptCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ScriptCreate) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ScriptCreate) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ScriptCreate) SetCategory(v string) {
	o.Category = &v
}

// GetScriptVersion returns the ScriptVersion field value if set, zero value otherwise.
func (o *ScriptCreate) GetScriptVersion() string {
	if o == nil || o.ScriptVersion == nil {
		var ret string
		return ret
	}
	return *o.ScriptVersion
}

// GetScriptVersionOk returns a tuple with the ScriptVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetScriptVersionOk() (*string, bool) {
	if o == nil || o.ScriptVersion == nil {
		return nil, false
	}
	return o.ScriptVersion, true
}

// HasScriptVersion returns a boolean if a field has been set.
func (o *ScriptCreate) HasScriptVersion() bool {
	if o != nil && o.ScriptVersion != nil {
		return true
	}

	return false
}

// SetScriptVersion gets a reference to the given string and assigns it to the ScriptVersion field.
func (o *ScriptCreate) SetScriptVersion(v string) {
	o.ScriptVersion = &v
}

// GetScriptPhase returns the ScriptPhase field value if set, zero value otherwise.
func (o *ScriptCreate) GetScriptPhase() string {
	if o == nil || o.ScriptPhase == nil {
		var ret string
		return ret
	}
	return *o.ScriptPhase
}

// GetScriptPhaseOk returns a tuple with the ScriptPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetScriptPhaseOk() (*string, bool) {
	if o == nil || o.ScriptPhase == nil {
		return nil, false
	}
	return o.ScriptPhase, true
}

// HasScriptPhase returns a boolean if a field has been set.
func (o *ScriptCreate) HasScriptPhase() bool {
	if o != nil && o.ScriptPhase != nil {
		return true
	}

	return false
}

// SetScriptPhase gets a reference to the given string and assigns it to the ScriptPhase field.
func (o *ScriptCreate) SetScriptPhase(v string) {
	o.ScriptPhase = &v
}

// GetScriptType returns the ScriptType field value if set, zero value otherwise.
func (o *ScriptCreate) GetScriptType() string {
	if o == nil || o.ScriptType == nil {
		var ret string
		return ret
	}
	return *o.ScriptType
}

// GetScriptTypeOk returns a tuple with the ScriptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetScriptTypeOk() (*string, bool) {
	if o == nil || o.ScriptType == nil {
		return nil, false
	}
	return o.ScriptType, true
}

// HasScriptType returns a boolean if a field has been set.
func (o *ScriptCreate) HasScriptType() bool {
	if o != nil && o.ScriptType != nil {
		return true
	}

	return false
}

// SetScriptType gets a reference to the given string and assigns it to the ScriptType field.
func (o *ScriptCreate) SetScriptType(v string) {
	o.ScriptType = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *ScriptCreate) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *ScriptCreate) HasScript() bool {
	if o != nil && o.Script != nil {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *ScriptCreate) SetScript(v string) {
	o.Script = &v
}

// GetRunAsUser returns the RunAsUser field value if set, zero value otherwise.
func (o *ScriptCreate) GetRunAsUser() string {
	if o == nil || o.RunAsUser == nil {
		var ret string
		return ret
	}
	return *o.RunAsUser
}

// GetRunAsUserOk returns a tuple with the RunAsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetRunAsUserOk() (*string, bool) {
	if o == nil || o.RunAsUser == nil {
		return nil, false
	}
	return o.RunAsUser, true
}

// HasRunAsUser returns a boolean if a field has been set.
func (o *ScriptCreate) HasRunAsUser() bool {
	if o != nil && o.RunAsUser != nil {
		return true
	}

	return false
}

// SetRunAsUser gets a reference to the given string and assigns it to the RunAsUser field.
func (o *ScriptCreate) SetRunAsUser(v string) {
	o.RunAsUser = &v
}

// GetSudoUser returns the SudoUser field value if set, zero value otherwise.
func (o *ScriptCreate) GetSudoUser() bool {
	if o == nil || o.SudoUser == nil {
		var ret bool
		return ret
	}
	return *o.SudoUser
}

// GetSudoUserOk returns a tuple with the SudoUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScriptCreate) GetSudoUserOk() (*bool, bool) {
	if o == nil || o.SudoUser == nil {
		return nil, false
	}
	return o.SudoUser, true
}

// HasSudoUser returns a boolean if a field has been set.
func (o *ScriptCreate) HasSudoUser() bool {
	if o != nil && o.SudoUser != nil {
		return true
	}

	return false
}

// SetSudoUser gets a reference to the given bool and assigns it to the SudoUser field.
func (o *ScriptCreate) SetSudoUser(v bool) {
	o.SudoUser = &v
}

func (o ScriptCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.ScriptVersion != nil {
		toSerialize["scriptVersion"] = o.ScriptVersion
	}
	if o.ScriptPhase != nil {
		toSerialize["scriptPhase"] = o.ScriptPhase
	}
	if o.ScriptType != nil {
		toSerialize["scriptType"] = o.ScriptType
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.RunAsUser != nil {
		toSerialize["runAsUser"] = o.RunAsUser
	}
	if o.SudoUser != nil {
		toSerialize["sudoUser"] = o.SudoUser
	}
	return json.Marshal(toSerialize)
}

type NullableScriptCreate struct {
	value *ScriptCreate
	isSet bool
}

func (v NullableScriptCreate) Get() *ScriptCreate {
	return v.value
}

func (v *NullableScriptCreate) Set(val *ScriptCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptCreate(val *ScriptCreate) *NullableScriptCreate {
	return &NullableScriptCreate{value: val, isSet: true}
}

func (v NullableScriptCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


