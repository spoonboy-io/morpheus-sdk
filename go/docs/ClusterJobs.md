# ClusterJobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**JobSummary** | Pointer to **NullableString** |  | [optional] 
**ScheduleMode** | Pointer to **NullableString** |  | [optional] 
**DateTime** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastRun** | Pointer to **NullableTime** |  | [optional] 
**LastResult** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**NullableInlineResponse20083LoadBalancerNodeCreatedBy**](inline_response_200_83_loadBalancerNode_createdBy.md) |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**Targets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CustomConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomOptions** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewClusterJobs

`func NewClusterJobs() *ClusterJobs`

NewClusterJobs instantiates a new ClusterJobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterJobsWithDefaults

`func NewClusterJobsWithDefaults() *ClusterJobs`

NewClusterJobsWithDefaults instantiates a new ClusterJobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterJobs) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterJobs) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterJobs) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterJobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterJobs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterJobs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterJobs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterJobs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterJobs) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterJobs) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterJobs) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterJobs) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ClusterJobs) GetType() InlineResponse20094Network`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterJobs) GetTypeOk() (*InlineResponse20094Network, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterJobs) SetType(v InlineResponse20094Network)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterJobs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetJobSummary

`func (o *ClusterJobs) GetJobSummary() string`

GetJobSummary returns the JobSummary field if non-nil, zero value otherwise.

### GetJobSummaryOk

`func (o *ClusterJobs) GetJobSummaryOk() (*string, bool)`

GetJobSummaryOk returns a tuple with the JobSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSummary

`func (o *ClusterJobs) SetJobSummary(v string)`

SetJobSummary sets JobSummary field to given value.

### HasJobSummary

`func (o *ClusterJobs) HasJobSummary() bool`

HasJobSummary returns a boolean if a field has been set.

### SetJobSummaryNil

`func (o *ClusterJobs) SetJobSummaryNil(b bool)`

 SetJobSummaryNil sets the value for JobSummary to be an explicit nil

### UnsetJobSummary
`func (o *ClusterJobs) UnsetJobSummary()`

UnsetJobSummary ensures that no value is present for JobSummary, not even an explicit nil
### GetScheduleMode

`func (o *ClusterJobs) GetScheduleMode() string`

GetScheduleMode returns the ScheduleMode field if non-nil, zero value otherwise.

### GetScheduleModeOk

`func (o *ClusterJobs) GetScheduleModeOk() (*string, bool)`

GetScheduleModeOk returns a tuple with the ScheduleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMode

`func (o *ClusterJobs) SetScheduleMode(v string)`

SetScheduleMode sets ScheduleMode field to given value.

### HasScheduleMode

`func (o *ClusterJobs) HasScheduleMode() bool`

HasScheduleMode returns a boolean if a field has been set.

### SetScheduleModeNil

`func (o *ClusterJobs) SetScheduleModeNil(b bool)`

 SetScheduleModeNil sets the value for ScheduleMode to be an explicit nil

### UnsetScheduleMode
`func (o *ClusterJobs) UnsetScheduleMode()`

UnsetScheduleMode ensures that no value is present for ScheduleMode, not even an explicit nil
### GetDateTime

`func (o *ClusterJobs) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ClusterJobs) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *ClusterJobs) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *ClusterJobs) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTimeNil

`func (o *ClusterJobs) SetDateTimeNil(b bool)`

 SetDateTimeNil sets the value for DateTime to be an explicit nil

### UnsetDateTime
`func (o *ClusterJobs) UnsetDateTime()`

UnsetDateTime ensures that no value is present for DateTime, not even an explicit nil
### GetStatus

`func (o *ClusterJobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterJobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterJobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterJobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ClusterJobs) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ClusterJobs) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNamespace

`func (o *ClusterJobs) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ClusterJobs) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ClusterJobs) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ClusterJobs) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ClusterJobs) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ClusterJobs) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCategory

`func (o *ClusterJobs) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterJobs) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterJobs) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterJobs) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterJobs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterJobs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterJobs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterJobs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterJobs) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterJobs) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ClusterJobs) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterJobs) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterJobs) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterJobs) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDateCreated

`func (o *ClusterJobs) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterJobs) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterJobs) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterJobs) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterJobs) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterJobs) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterJobs) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterJobs) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastRun

`func (o *ClusterJobs) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *ClusterJobs) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *ClusterJobs) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *ClusterJobs) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### SetLastRunNil

`func (o *ClusterJobs) SetLastRunNil(b bool)`

 SetLastRunNil sets the value for LastRun to be an explicit nil

### UnsetLastRun
`func (o *ClusterJobs) UnsetLastRun()`

UnsetLastRun ensures that no value is present for LastRun, not even an explicit nil
### GetLastResult

`func (o *ClusterJobs) GetLastResult() string`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ClusterJobs) GetLastResultOk() (*string, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ClusterJobs) SetLastResult(v string)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ClusterJobs) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *ClusterJobs) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *ClusterJobs) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetCreatedBy

`func (o *ClusterJobs) GetCreatedBy() InlineResponse20083LoadBalancerNodeCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClusterJobs) GetCreatedByOk() (*InlineResponse20083LoadBalancerNodeCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClusterJobs) SetCreatedBy(v InlineResponse20083LoadBalancerNodeCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClusterJobs) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ClusterJobs) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ClusterJobs) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetTargetType

`func (o *ClusterJobs) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ClusterJobs) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ClusterJobs) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ClusterJobs) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ClusterJobs) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ClusterJobs) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargets

`func (o *ClusterJobs) GetTargets() []map[string]interface{}`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ClusterJobs) GetTargetsOk() (*[]map[string]interface{}, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ClusterJobs) SetTargets(v []map[string]interface{})`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ClusterJobs) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetCustomConfig

`func (o *ClusterJobs) GetCustomConfig() map[string]interface{}`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *ClusterJobs) GetCustomConfigOk() (*map[string]interface{}, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *ClusterJobs) SetCustomConfig(v map[string]interface{})`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *ClusterJobs) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *ClusterJobs) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *ClusterJobs) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil
### GetCustomOptions

`func (o *ClusterJobs) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ClusterJobs) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ClusterJobs) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ClusterJobs) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### SetCustomOptionsNil

`func (o *ClusterJobs) SetCustomOptionsNil(b bool)`

 SetCustomOptionsNil sets the value for CustomOptions to be an explicit nil

### UnsetCustomOptions
`func (o *ClusterJobs) UnsetCustomOptions()`

UnsetCustomOptions ensures that no value is present for CustomOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


