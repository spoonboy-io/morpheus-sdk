# ApiJobsIdJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the Job | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] [default to true]
**Task** | Pointer to [**NullableApiJobsIdJobTask**](_api_jobs__id__job_task.md) |  | [optional] 
**Workflow** | Pointer to [**NullableApiJobsIdJobTask**](_api_jobs__id__job_task.md) |  | [optional] 
**ScanPath** | Pointer to **NullableString** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **NullableString** | Security Profile. Only applies to type scap-package. | [optional] 
**TargetType** | Pointer to **string** | Target type where job will execute | [optional] 
**Targets** | Pointer to [**[]ApiJobsIdJobTargets**](ApiJobsIdJobTargets.md) |  | [optional] 
**InstanceLabel** | Pointer to **string** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | Pointer to **string** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**ScheduleMode** | Pointer to [**OneOfstringlong**](oneOf&lt;string,long&gt;.md) |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | Pointer to **string** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | Pointer to **time.Time** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | Pointer to **bool** | If true, executes job | [optional] 

## Methods

### NewApiJobsIdJob

`func NewApiJobsIdJob() *ApiJobsIdJob`

NewApiJobsIdJob instantiates a new ApiJobsIdJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiJobsIdJobWithDefaults

`func NewApiJobsIdJobWithDefaults() *ApiJobsIdJob`

NewApiJobsIdJobWithDefaults instantiates a new ApiJobsIdJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiJobsIdJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiJobsIdJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiJobsIdJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiJobsIdJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ApiJobsIdJob) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiJobsIdJob) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiJobsIdJob) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiJobsIdJob) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ApiJobsIdJob) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ApiJobsIdJob) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetEnabled

`func (o *ApiJobsIdJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiJobsIdJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiJobsIdJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiJobsIdJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTask

`func (o *ApiJobsIdJob) GetTask() ApiJobsIdJobTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ApiJobsIdJob) GetTaskOk() (*ApiJobsIdJobTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ApiJobsIdJob) SetTask(v ApiJobsIdJobTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *ApiJobsIdJob) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *ApiJobsIdJob) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *ApiJobsIdJob) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetWorkflow

`func (o *ApiJobsIdJob) GetWorkflow() ApiJobsIdJobTask`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ApiJobsIdJob) GetWorkflowOk() (*ApiJobsIdJobTask, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ApiJobsIdJob) SetWorkflow(v ApiJobsIdJobTask)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ApiJobsIdJob) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *ApiJobsIdJob) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *ApiJobsIdJob) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetScanPath

`func (o *ApiJobsIdJob) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *ApiJobsIdJob) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *ApiJobsIdJob) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *ApiJobsIdJob) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### SetScanPathNil

`func (o *ApiJobsIdJob) SetScanPathNil(b bool)`

 SetScanPathNil sets the value for ScanPath to be an explicit nil

### UnsetScanPath
`func (o *ApiJobsIdJob) UnsetScanPath()`

UnsetScanPath ensures that no value is present for ScanPath, not even an explicit nil
### GetSecurityProfile

`func (o *ApiJobsIdJob) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *ApiJobsIdJob) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *ApiJobsIdJob) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *ApiJobsIdJob) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### SetSecurityProfileNil

`func (o *ApiJobsIdJob) SetSecurityProfileNil(b bool)`

 SetSecurityProfileNil sets the value for SecurityProfile to be an explicit nil

### UnsetSecurityProfile
`func (o *ApiJobsIdJob) UnsetSecurityProfile()`

UnsetSecurityProfile ensures that no value is present for SecurityProfile, not even an explicit nil
### GetTargetType

`func (o *ApiJobsIdJob) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ApiJobsIdJob) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ApiJobsIdJob) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ApiJobsIdJob) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargets

`func (o *ApiJobsIdJob) GetTargets() []ApiJobsIdJobTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ApiJobsIdJob) GetTargetsOk() (*[]ApiJobsIdJobTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ApiJobsIdJob) SetTargets(v []ApiJobsIdJobTargets)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ApiJobsIdJob) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *ApiJobsIdJob) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *ApiJobsIdJob) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetInstanceLabel

`func (o *ApiJobsIdJob) GetInstanceLabel() string`

GetInstanceLabel returns the InstanceLabel field if non-nil, zero value otherwise.

### GetInstanceLabelOk

`func (o *ApiJobsIdJob) GetInstanceLabelOk() (*string, bool)`

GetInstanceLabelOk returns a tuple with the InstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceLabel

`func (o *ApiJobsIdJob) SetInstanceLabel(v string)`

SetInstanceLabel sets InstanceLabel field to given value.

### HasInstanceLabel

`func (o *ApiJobsIdJob) HasInstanceLabel() bool`

HasInstanceLabel returns a boolean if a field has been set.

### GetServerLabel

`func (o *ApiJobsIdJob) GetServerLabel() string`

GetServerLabel returns the ServerLabel field if non-nil, zero value otherwise.

### GetServerLabelOk

`func (o *ApiJobsIdJob) GetServerLabelOk() (*string, bool)`

GetServerLabelOk returns a tuple with the ServerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLabel

`func (o *ApiJobsIdJob) SetServerLabel(v string)`

SetServerLabel sets ServerLabel field to given value.

### HasServerLabel

`func (o *ApiJobsIdJob) HasServerLabel() bool`

HasServerLabel returns a boolean if a field has been set.

### GetScheduleMode

`func (o *ApiJobsIdJob) GetScheduleMode() OneOfstringlong`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *ApiJobsIdJob) GetScheduleModeOk() (*OneOfstringlong, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *ApiJobsIdJob) SetScheduleMode(v OneOfstringlong)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *ApiJobsIdJob) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ApiJobsIdJob) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ApiJobsIdJob) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ApiJobsIdJob) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ApiJobsIdJob) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ApiJobsIdJob) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ApiJobsIdJob) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ApiJobsIdJob) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ApiJobsIdJob) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDateTime

`func (o *ApiJobsIdJob) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ApiJobsIdJob) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ApiJobsIdJob) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ApiJobsIdJob) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetRun

`func (o *ApiJobsIdJob) GetRun() bool`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *ApiJobsIdJob) GetRunOk() (*bool, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *ApiJobsIdJob) SetRun(v bool)`

SetRun sets Run field to given value.

### HasRun

`func (o *ApiJobsIdJob) HasRun() bool`

HasRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


