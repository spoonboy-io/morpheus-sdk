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

// Job struct for Job
type Job struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Type *InlineResponse20094Network `json:"type,omitempty"`
	Workflow NullableJobWorkflow `json:"workflow,omitempty"`
	Task NullableJobTask `json:"task,omitempty"`
	SecurityPackage NullableJobSecurityPackage `json:"securityPackage,omitempty"`
	JobSummary NullableString `json:"jobSummary,omitempty"`
	ScheduleMode *OneOfstringlong `json:"scheduleMode,omitempty"`
	DateTime NullableString `json:"dateTime,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Category NullableString `json:"category,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	LastRun *time.Time `json:"lastRun,omitempty"`
	LastResult NullableString `json:"lastResult,omitempty"`
	CreatedBy NullableJobCreatedBy `json:"createdBy,omitempty"`
	TargetType NullableString `json:"targetType,omitempty"`
	Targets []JobTargets `json:"targets,omitempty"`
	// Scan Checklist. Only applies to type scap-package.
	ScanPath NullableString `json:"scanPath,omitempty"`
	// Security Profile. Only applies to type scap-package.
	SecurityProfile NullableString `json:"securityProfile,omitempty"`
	CustomConfig NullableString `json:"customConfig,omitempty"`
	CustomOptions NullableJobCustomOptions `json:"customOptions,omitempty"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob() *Job {
	this := Job{}
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Job) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Job) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Job) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Job) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Job) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Job) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Job) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Job) SetLabels(v []string) {
	o.Labels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Job) GetType() InlineResponse20094Network {
	if o == nil || o.Type == nil {
		var ret InlineResponse20094Network
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTypeOk() (*InlineResponse20094Network, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Job) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given InlineResponse20094Network and assigns it to the Type field.
func (o *Job) SetType(v InlineResponse20094Network) {
	o.Type = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetWorkflow() JobWorkflow {
	if o == nil || o.Workflow.Get() == nil {
		var ret JobWorkflow
		return ret
	}
	return *o.Workflow.Get()
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetWorkflowOk() (*JobWorkflow, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Workflow.Get(), o.Workflow.IsSet()
}

// HasWorkflow returns a boolean if a field has been set.
func (o *Job) HasWorkflow() bool {
	if o != nil && o.Workflow.IsSet() {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given NullableJobWorkflow and assigns it to the Workflow field.
func (o *Job) SetWorkflow(v JobWorkflow) {
	o.Workflow.Set(&v)
}
// SetWorkflowNil sets the value for Workflow to be an explicit nil
func (o *Job) SetWorkflowNil() {
	o.Workflow.Set(nil)
}

// UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
func (o *Job) UnsetWorkflow() {
	o.Workflow.Unset()
}

// GetTask returns the Task field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetTask() JobTask {
	if o == nil || o.Task.Get() == nil {
		var ret JobTask
		return ret
	}
	return *o.Task.Get()
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetTaskOk() (*JobTask, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Task.Get(), o.Task.IsSet()
}

// HasTask returns a boolean if a field has been set.
func (o *Job) HasTask() bool {
	if o != nil && o.Task.IsSet() {
		return true
	}

	return false
}

// SetTask gets a reference to the given NullableJobTask and assigns it to the Task field.
func (o *Job) SetTask(v JobTask) {
	o.Task.Set(&v)
}
// SetTaskNil sets the value for Task to be an explicit nil
func (o *Job) SetTaskNil() {
	o.Task.Set(nil)
}

// UnsetTask ensures that no value is present for Task, not even an explicit nil
func (o *Job) UnsetTask() {
	o.Task.Unset()
}

// GetSecurityPackage returns the SecurityPackage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetSecurityPackage() JobSecurityPackage {
	if o == nil || o.SecurityPackage.Get() == nil {
		var ret JobSecurityPackage
		return ret
	}
	return *o.SecurityPackage.Get()
}

// GetSecurityPackageOk returns a tuple with the SecurityPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetSecurityPackageOk() (*JobSecurityPackage, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityPackage.Get(), o.SecurityPackage.IsSet()
}

// HasSecurityPackage returns a boolean if a field has been set.
func (o *Job) HasSecurityPackage() bool {
	if o != nil && o.SecurityPackage.IsSet() {
		return true
	}

	return false
}

// SetSecurityPackage gets a reference to the given NullableJobSecurityPackage and assigns it to the SecurityPackage field.
func (o *Job) SetSecurityPackage(v JobSecurityPackage) {
	o.SecurityPackage.Set(&v)
}
// SetSecurityPackageNil sets the value for SecurityPackage to be an explicit nil
func (o *Job) SetSecurityPackageNil() {
	o.SecurityPackage.Set(nil)
}

// UnsetSecurityPackage ensures that no value is present for SecurityPackage, not even an explicit nil
func (o *Job) UnsetSecurityPackage() {
	o.SecurityPackage.Unset()
}

// GetJobSummary returns the JobSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetJobSummary() string {
	if o == nil || o.JobSummary.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobSummary.Get()
}

// GetJobSummaryOk returns a tuple with the JobSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetJobSummaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobSummary.Get(), o.JobSummary.IsSet()
}

// HasJobSummary returns a boolean if a field has been set.
func (o *Job) HasJobSummary() bool {
	if o != nil && o.JobSummary.IsSet() {
		return true
	}

	return false
}

// SetJobSummary gets a reference to the given NullableString and assigns it to the JobSummary field.
func (o *Job) SetJobSummary(v string) {
	o.JobSummary.Set(&v)
}
// SetJobSummaryNil sets the value for JobSummary to be an explicit nil
func (o *Job) SetJobSummaryNil() {
	o.JobSummary.Set(nil)
}

// UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
func (o *Job) UnsetJobSummary() {
	o.JobSummary.Unset()
}

// GetScheduleMode returns the ScheduleMode field value if set, zero value otherwise.
func (o *Job) GetScheduleMode() OneOfstringlong {
	if o == nil || o.ScheduleMode == nil {
		var ret OneOfstringlong
		return ret
	}
	return *o.ScheduleMode
}

// GetScheduleModeOk returns a tuple with the ScheduleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetScheduleModeOk() (*OneOfstringlong, bool) {
	if o == nil || o.ScheduleMode == nil {
		return nil, false
	}
	return o.ScheduleMode, true
}

// HasScheduleMode returns a boolean if a field has been set.
func (o *Job) HasScheduleMode() bool {
	if o != nil && o.ScheduleMode != nil {
		return true
	}

	return false
}

// SetScheduleMode gets a reference to the given OneOfstringlong and assigns it to the ScheduleMode field.
func (o *Job) SetScheduleMode(v OneOfstringlong) {
	o.ScheduleMode = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetDateTime() string {
	if o == nil || o.DateTime.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateTime.Get()
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetDateTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateTime.Get(), o.DateTime.IsSet()
}

// HasDateTime returns a boolean if a field has been set.
func (o *Job) HasDateTime() bool {
	if o != nil && o.DateTime.IsSet() {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given NullableString and assigns it to the DateTime field.
func (o *Job) SetDateTime(v string) {
	o.DateTime.Set(&v)
}
// SetDateTimeNil sets the value for DateTime to be an explicit nil
func (o *Job) SetDateTimeNil() {
	o.DateTime.Set(nil)
}

// UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
func (o *Job) UnsetDateTime() {
	o.DateTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Job) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Job) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Job) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Job) UnsetStatus() {
	o.Status.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetNamespace() string {
	if o == nil || o.Namespace.Get() == nil {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetNamespaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *Job) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *Job) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *Job) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *Job) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *Job) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *Job) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *Job) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *Job) UnsetCategory() {
	o.Category.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Job) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Job) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Job) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Job) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Job) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Job) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Job) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Job) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Job) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Job) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Job) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Job) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Job) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *Job) GetLastRun() time.Time {
	if o == nil || o.LastRun == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetLastRunOk() (*time.Time, bool) {
	if o == nil || o.LastRun == nil {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *Job) HasLastRun() bool {
	if o != nil && o.LastRun != nil {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given time.Time and assigns it to the LastRun field.
func (o *Job) SetLastRun(v time.Time) {
	o.LastRun = &v
}

// GetLastResult returns the LastResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetLastResult() string {
	if o == nil || o.LastResult.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastResult.Get()
}

// GetLastResultOk returns a tuple with the LastResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetLastResultOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastResult.Get(), o.LastResult.IsSet()
}

// HasLastResult returns a boolean if a field has been set.
func (o *Job) HasLastResult() bool {
	if o != nil && o.LastResult.IsSet() {
		return true
	}

	return false
}

// SetLastResult gets a reference to the given NullableString and assigns it to the LastResult field.
func (o *Job) SetLastResult(v string) {
	o.LastResult.Set(&v)
}
// SetLastResultNil sets the value for LastResult to be an explicit nil
func (o *Job) SetLastResultNil() {
	o.LastResult.Set(nil)
}

// UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
func (o *Job) UnsetLastResult() {
	o.LastResult.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetCreatedBy() JobCreatedBy {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret JobCreatedBy
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetCreatedByOk() (*JobCreatedBy, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Job) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableJobCreatedBy and assigns it to the CreatedBy field.
func (o *Job) SetCreatedBy(v JobCreatedBy) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *Job) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *Job) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetTargetType returns the TargetType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetTargetType() string {
	if o == nil || o.TargetType.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetType.Get()
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetType.Get(), o.TargetType.IsSet()
}

// HasTargetType returns a boolean if a field has been set.
func (o *Job) HasTargetType() bool {
	if o != nil && o.TargetType.IsSet() {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given NullableString and assigns it to the TargetType field.
func (o *Job) SetTargetType(v string) {
	o.TargetType.Set(&v)
}
// SetTargetTypeNil sets the value for TargetType to be an explicit nil
func (o *Job) SetTargetTypeNil() {
	o.TargetType.Set(nil)
}

// UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
func (o *Job) UnsetTargetType() {
	o.TargetType.Unset()
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetTargets() []JobTargets {
	if o == nil  {
		var ret []JobTargets
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetTargetsOk() (*[]JobTargets, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *Job) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []JobTargets and assigns it to the Targets field.
func (o *Job) SetTargets(v []JobTargets) {
	o.Targets = v
}

// GetScanPath returns the ScanPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetScanPath() string {
	if o == nil || o.ScanPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ScanPath.Get()
}

// GetScanPathOk returns a tuple with the ScanPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetScanPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScanPath.Get(), o.ScanPath.IsSet()
}

// HasScanPath returns a boolean if a field has been set.
func (o *Job) HasScanPath() bool {
	if o != nil && o.ScanPath.IsSet() {
		return true
	}

	return false
}

// SetScanPath gets a reference to the given NullableString and assigns it to the ScanPath field.
func (o *Job) SetScanPath(v string) {
	o.ScanPath.Set(&v)
}
// SetScanPathNil sets the value for ScanPath to be an explicit nil
func (o *Job) SetScanPathNil() {
	o.ScanPath.Set(nil)
}

// UnsetScanPath ensures that no value is present for ScanPath, not even an explicit nil
func (o *Job) UnsetScanPath() {
	o.ScanPath.Unset()
}

// GetSecurityProfile returns the SecurityProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetSecurityProfile() string {
	if o == nil || o.SecurityProfile.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecurityProfile.Get()
}

// GetSecurityProfileOk returns a tuple with the SecurityProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetSecurityProfileOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecurityProfile.Get(), o.SecurityProfile.IsSet()
}

// HasSecurityProfile returns a boolean if a field has been set.
func (o *Job) HasSecurityProfile() bool {
	if o != nil && o.SecurityProfile.IsSet() {
		return true
	}

	return false
}

// SetSecurityProfile gets a reference to the given NullableString and assigns it to the SecurityProfile field.
func (o *Job) SetSecurityProfile(v string) {
	o.SecurityProfile.Set(&v)
}
// SetSecurityProfileNil sets the value for SecurityProfile to be an explicit nil
func (o *Job) SetSecurityProfileNil() {
	o.SecurityProfile.Set(nil)
}

// UnsetSecurityProfile ensures that no value is present for SecurityProfile, not even an explicit nil
func (o *Job) UnsetSecurityProfile() {
	o.SecurityProfile.Unset()
}

// GetCustomConfig returns the CustomConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetCustomConfig() string {
	if o == nil || o.CustomConfig.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomConfig.Get()
}

// GetCustomConfigOk returns a tuple with the CustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetCustomConfigOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomConfig.Get(), o.CustomConfig.IsSet()
}

// HasCustomConfig returns a boolean if a field has been set.
func (o *Job) HasCustomConfig() bool {
	if o != nil && o.CustomConfig.IsSet() {
		return true
	}

	return false
}

// SetCustomConfig gets a reference to the given NullableString and assigns it to the CustomConfig field.
func (o *Job) SetCustomConfig(v string) {
	o.CustomConfig.Set(&v)
}
// SetCustomConfigNil sets the value for CustomConfig to be an explicit nil
func (o *Job) SetCustomConfigNil() {
	o.CustomConfig.Set(nil)
}

// UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
func (o *Job) UnsetCustomConfig() {
	o.CustomConfig.Unset()
}

// GetCustomOptions returns the CustomOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Job) GetCustomOptions() JobCustomOptions {
	if o == nil || o.CustomOptions.Get() == nil {
		var ret JobCustomOptions
		return ret
	}
	return *o.CustomOptions.Get()
}

// GetCustomOptionsOk returns a tuple with the CustomOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetCustomOptionsOk() (*JobCustomOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomOptions.Get(), o.CustomOptions.IsSet()
}

// HasCustomOptions returns a boolean if a field has been set.
func (o *Job) HasCustomOptions() bool {
	if o != nil && o.CustomOptions.IsSet() {
		return true
	}

	return false
}

// SetCustomOptions gets a reference to the given NullableJobCustomOptions and assigns it to the CustomOptions field.
func (o *Job) SetCustomOptions(v JobCustomOptions) {
	o.CustomOptions.Set(&v)
}
// SetCustomOptionsNil sets the value for CustomOptions to be an explicit nil
func (o *Job) SetCustomOptionsNil() {
	o.CustomOptions.Set(nil)
}

// UnsetCustomOptions ensures that no value is present for CustomOptions, not even an explicit nil
func (o *Job) UnsetCustomOptions() {
	o.CustomOptions.Unset()
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Workflow.IsSet() {
		toSerialize["workflow"] = o.Workflow.Get()
	}
	if o.Task.IsSet() {
		toSerialize["task"] = o.Task.Get()
	}
	if o.SecurityPackage.IsSet() {
		toSerialize["securityPackage"] = o.SecurityPackage.Get()
	}
	if o.JobSummary.IsSet() {
		toSerialize["jobSummary"] = o.JobSummary.Get()
	}
	if o.ScheduleMode != nil {
		toSerialize["scheduleMode"] = o.ScheduleMode
	}
	if o.DateTime.IsSet() {
		toSerialize["dateTime"] = o.DateTime.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.LastRun != nil {
		toSerialize["lastRun"] = o.LastRun
	}
	if o.LastResult.IsSet() {
		toSerialize["lastResult"] = o.LastResult.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.TargetType.IsSet() {
		toSerialize["targetType"] = o.TargetType.Get()
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.ScanPath.IsSet() {
		toSerialize["scanPath"] = o.ScanPath.Get()
	}
	if o.SecurityProfile.IsSet() {
		toSerialize["securityProfile"] = o.SecurityProfile.Get()
	}
	if o.CustomConfig.IsSet() {
		toSerialize["customConfig"] = o.CustomConfig.Get()
	}
	if o.CustomOptions.IsSet() {
		toSerialize["customOptions"] = o.CustomOptions.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


