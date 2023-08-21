# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**Workflow** | Pointer to [**NullableJobWorkflow**](job_workflow.md) |  | [optional] 
**Task** | Pointer to [**NullableJobTask**](job_task.md) |  | [optional] 
**SecurityPackage** | Pointer to [**NullableJobSecurityPackage**](job_securityPackage.md) |  | [optional] 
**JobSummary** | Pointer to **NullableString** |  | [optional] 
**ScheduleMode** | Pointer to [**OneOfstringlong**](oneOf&lt;string,long&gt;.md) |  | [optional] 
**DateTime** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastRun** | Pointer to **time.Time** |  | [optional] 
**LastResult** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**NullableJobCreatedBy**](job_createdBy.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**Targets** | Pointer to [**[]JobTargets**](JobTargets.md) |  | [optional] 
**ScanPath** | Pointer to **NullableString** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | Pointer to **NullableString** | Security Profile. Only applies to type scap-package. | [optional] 
**CustomConfig** | Pointer to **NullableString** |  | [optional] 
**CustomOptions** | Pointer to [**NullableJobCustomOptions**](job_customOptions.md) |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Job) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *Job) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Job) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Job) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Job) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *Job) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Job) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *Job) GetType() InlineResponse20094Network`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*InlineResponse20094Network, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v InlineResponse20094Network)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkflow

`func (o *Job) GetWorkflow() JobWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *Job) GetWorkflowOk() (*JobWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *Job) SetWorkflow(v JobWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *Job) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *Job) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *Job) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetTask

`func (o *Job) GetTask() JobTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Job) GetTaskOk() (*JobTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Job) SetTask(v JobTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *Job) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *Job) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *Job) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetSecurityPackage

`func (o *Job) GetSecurityPackage() JobSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *Job) GetSecurityPackageOk() (*JobSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *Job) SetSecurityPackage(v JobSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *Job) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### SetSecurityPackageNil

`func (o *Job) SetSecurityPackageNil(b bool)`

 SetSecurityPackageNil sets the value for SecurityPackage to be an explicit nil

### UnsetSecurityPackage
`func (o *Job) UnsetSecurityPackage()`

UnsetSecurityPackage ensures that no value is present for SecurityPackage, not even an explicit nil
### GetJobSummary

`func (o *Job) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *Job) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *Job) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *Job) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### SetJobSummaryNil

`func (o *Job) SetJobSummaryNil(b bool)`

 SetJobSummaryNil sets the value for JobSummary to be an explicit nil

### UnsetJobSummary
`func (o *Job) UnsetJobSummary()`

UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
### GetScheduleMode

`func (o *Job) GetScheduleMode() OneOfstringlong`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *Job) GetScheduleModeOk() (*OneOfstringlong, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *Job) SetScheduleMode(v OneOfstringlong)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *Job) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### GetDateTime

`func (o *Job) GetDateTime() string`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *Job) GetDateTimeOk() (*string, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *Job) SetDateTime(v string)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *Job) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *Job) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *Job) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Job) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Job) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *Job) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Job) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Job) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Job) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *Job) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Job) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCategory

`func (o *Job) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Job) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Job) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Job) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Job) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Job) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *Job) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Job) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Job) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Job) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Job) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Job) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *Job) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Job) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Job) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Job) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *Job) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Job) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Job) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Job) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Job) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Job) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Job) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Job) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *Job) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *Job) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *Job) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *Job) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *Job) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *Job) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *Job) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *Job) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *Job) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *Job) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetCreatedBy

`func (o *Job) GetCreatedBy() JobCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Job) GetCreatedByOk() (*JobCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Job) SetCreatedBy(v JobCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Job) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Job) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Job) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetTargetType

`func (o *Job) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Job) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Job) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Job) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *Job) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *Job) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargets

`func (o *Job) GetTargets() []JobTargets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *Job) GetTargetsOk() (*[]JobTargets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *Job) SetTargets(v []JobTargets)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *Job) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *Job) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *Job) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil
### GetScanPath

`func (o *Job) GetScanPath() string`

GetScanPath returns the ScanPath field if non-nil, zero value otherwise.

### GetScanPathOk

`func (o *Job) GetScanPathOk() (*string, bool)`

GetScanPathOk returns a tuple with the ScanPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPath

`func (o *Job) SetScanPath(v string)`

SetScanPath sets ScanPath field to given value.

### HasScanPath

`func (o *Job) HasScanPath() bool`

HasScanPath returns a boolean if a field has been set.

### SetScanPathNil

`func (o *Job) SetScanPathNil(b bool)`

 SetScanPathNil sets the value for ScanPath to be an explicit nil

### UnsetScanPath
`func (o *Job) UnsetScanPath()`

UnsetScanPath ensures that no value is present for ScanPath, not even an explicit nil
### GetSecurityProfile

`func (o *Job) GetSecurityProfile() string`

GetSecurityProfile returns the SecurityProfile field if non-nil, zero value otherwise.

### GetSecurityProfileOk

`func (o *Job) GetSecurityProfileOk() (*string, bool)`

GetSecurityProfileOk returns a tuple with the SecurityProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProfile

`func (o *Job) SetSecurityProfile(v string)`

SetSecurityProfile sets SecurityProfile field to given value.

### HasSecurityProfile

`func (o *Job) HasSecurityProfile() bool`

HasSecurityProfile returns a boolean if a field has been set.

### SetSecurityProfileNil

`func (o *Job) SetSecurityProfileNil(b bool)`

 SetSecurityProfileNil sets the value for SecurityProfile to be an explicit nil

### UnsetSecurityProfile
`func (o *Job) UnsetSecurityProfile()`

UnsetSecurityProfile ensures that no value is present for SecurityProfile, not even an explicit nil
### GetCustomConfig

`func (o *Job) GetCustomConfig() string`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *Job) GetCustomConfigOk() (*string, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *Job) SetCustomConfig(v string)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *Job) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *Job) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *Job) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
### GetCustomOptions

`func (o *Job) GetCustomOptions() JobCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *Job) GetCustomOptionsOk() (*JobCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *Job) SetCustomOptions(v JobCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *Job) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### SetCustomOptionsNil

`func (o *Job) SetCustomOptionsNil(b bool)`

 SetCustomOptionsNil sets the value for CustomOptions to be an explicit nil

### UnsetCustomOptions
`func (o *Job) UnsetCustomOptions()`

UnsetCustomOptions ensures that no value is present for CustomOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


