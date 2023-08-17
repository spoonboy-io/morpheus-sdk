# CheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Instance** | Pointer to [**NullableCheckGroupInstance**](CheckGroupInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InUptime** | Pointer to **bool** |  | [optional] 
**LastCheckStatus** | Pointer to **NullableString** |  | [optional] 
**LastWarningDate** | Pointer to **NullableTime** |  | [optional] 
**LastErrorDate** | Pointer to **NullableTime** |  | [optional] 
**LastSuccessDate** | Pointer to **NullableTime** |  | [optional] 
**LastRunDate** | Pointer to **NullableTime** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**OutageTime** | Pointer to **int64** |  | [optional] 
**LastTimer** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **int64** |  | [optional] 
**History** | Pointer to **NullableString** |  | [optional] 
**MinHappy** | Pointer to **int64** |  | [optional] 
**LastMetric** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**CreateIncident** | Pointer to **bool** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**ActivityActivityInnerUser**](ActivityActivityInnerUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Availability** | Pointer to **NullableFloat32** |  | [optional] 
**CheckType** | Pointer to [**CheckCheckType**](CheckCheckType.md) |  | [optional] 
**Checks** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCheckGroup

`func NewCheckGroup() *CheckGroup`

NewCheckGroup instantiates a new CheckGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckGroupWithDefaults

`func NewCheckGroupWithDefaults() *CheckGroup`

NewCheckGroupWithDefaults instantiates a new CheckGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CheckGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *CheckGroup) GetAccount() UpdateBlueprintPermissionsRequestResourcePermissionSitesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CheckGroup) GetAccountOk() (*UpdateBlueprintPermissionsRequestResourcePermissionSitesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CheckGroup) SetAccount(v UpdateBlueprintPermissionsRequestResourcePermissionSitesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CheckGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInstance

`func (o *CheckGroup) GetInstance() CheckGroupInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CheckGroup) GetInstanceOk() (*CheckGroupInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CheckGroup) SetInstance(v CheckGroupInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CheckGroup) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *CheckGroup) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *CheckGroup) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetName

`func (o *CheckGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CheckGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CheckGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CheckGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInUptime

`func (o *CheckGroup) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *CheckGroup) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *CheckGroup) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *CheckGroup) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetLastCheckStatus

`func (o *CheckGroup) GetLastCheckStatus() string`

GetLastCheckStatus returns the LastCheckStatus field if non-nil, zero value otherwise.

### GetLastCheckStatusOk

`func (o *CheckGroup) GetLastCheckStatusOk() (*string, bool)`

GetLastCheckStatusOk returns a tuple with the LastCheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckStatus

`func (o *CheckGroup) SetLastCheckStatus(v string)`

SetLastCheckStatus sets LastCheckStatus field to given value.

### HasLastCheckStatus

`func (o *CheckGroup) HasLastCheckStatus() bool`

HasLastCheckStatus returns a boolean if a field has been set.

### SetLastCheckStatusNil

`func (o *CheckGroup) SetLastCheckStatusNil(b bool)`

 SetLastCheckStatusNil sets the value for LastCheckStatus to be an explicit nil

### UnsetLastCheckStatus
`func (o *CheckGroup) UnsetLastCheckStatus()`

UnsetLastCheckStatus ensures that no value is present for LastCheckStatus, not even an explicit nil
### GetLastWarningDate

`func (o *CheckGroup) GetLastWarningDate() time.Time`

GetLastWarningDate returns the LastWarningDate field if non-nil, zero value otherwise.

### GetLastWarningDateOk

`func (o *CheckGroup) GetLastWarningDateOk() (*time.Time, bool)`

GetLastWarningDateOk returns a tuple with the LastWarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWarningDate

`func (o *CheckGroup) SetLastWarningDate(v time.Time)`

SetLastWarningDate sets LastWarningDate field to given value.

### HasLastWarningDate

`func (o *CheckGroup) HasLastWarningDate() bool`

HasLastWarningDate returns a boolean if a field has been set.

### SetLastWarningDateNil

`func (o *CheckGroup) SetLastWarningDateNil(b bool)`

 SetLastWarningDateNil sets the value for LastWarningDate to be an explicit nil

### UnsetLastWarningDate
`func (o *CheckGroup) UnsetLastWarningDate()`

UnsetLastWarningDate ensures that no value is present for LastWarningDate, not even an explicit nil
### GetLastErrorDate

`func (o *CheckGroup) GetLastErrorDate() time.Time`

GetLastErrorDate returns the LastErrorDate field if non-nil, zero value otherwise.

### GetLastErrorDateOk

`func (o *CheckGroup) GetLastErrorDateOk() (*time.Time, bool)`

GetLastErrorDateOk returns a tuple with the LastErrorDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorDate

`func (o *CheckGroup) SetLastErrorDate(v time.Time)`

SetLastErrorDate sets LastErrorDate field to given value.

### HasLastErrorDate

`func (o *CheckGroup) HasLastErrorDate() bool`

HasLastErrorDate returns a boolean if a field has been set.

### SetLastErrorDateNil

`func (o *CheckGroup) SetLastErrorDateNil(b bool)`

 SetLastErrorDateNil sets the value for LastErrorDate to be an explicit nil

### UnsetLastErrorDate
`func (o *CheckGroup) UnsetLastErrorDate()`

UnsetLastErrorDate ensures that no value is present for LastErrorDate, not even an explicit nil
### GetLastSuccessDate

`func (o *CheckGroup) GetLastSuccessDate() time.Time`

GetLastSuccessDate returns the LastSuccessDate field if non-nil, zero value otherwise.

### GetLastSuccessDateOk

`func (o *CheckGroup) GetLastSuccessDateOk() (*time.Time, bool)`

GetLastSuccessDateOk returns a tuple with the LastSuccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessDate

`func (o *CheckGroup) SetLastSuccessDate(v time.Time)`

SetLastSuccessDate sets LastSuccessDate field to given value.

### HasLastSuccessDate

`func (o *CheckGroup) HasLastSuccessDate() bool`

HasLastSuccessDate returns a boolean if a field has been set.

### SetLastSuccessDateNil

`func (o *CheckGroup) SetLastSuccessDateNil(b bool)`

 SetLastSuccessDateNil sets the value for LastSuccessDate to be an explicit nil

### UnsetLastSuccessDate
`func (o *CheckGroup) UnsetLastSuccessDate()`

UnsetLastSuccessDate ensures that no value is present for LastSuccessDate, not even an explicit nil
### GetLastRunDate

`func (o *CheckGroup) GetLastRunDate() time.Time`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *CheckGroup) GetLastRunDateOk() (*time.Time, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *CheckGroup) SetLastRunDate(v time.Time)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *CheckGroup) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### SetLastRunDateNil

`func (o *CheckGroup) SetLastRunDateNil(b bool)`

 SetLastRunDateNil sets the value for LastRunDate to be an explicit nil

### UnsetLastRunDate
`func (o *CheckGroup) UnsetLastRunDate()`

UnsetLastRunDate ensures that no value is present for LastRunDate, not even an explicit nil
### GetLastError

`func (o *CheckGroup) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *CheckGroup) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *CheckGroup) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *CheckGroup) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *CheckGroup) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *CheckGroup) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetOutageTime

`func (o *CheckGroup) GetOutageTime() int64`

GetOutageTime returns the OutageTime field if non-nil, zero value otherwise.

### GetOutageTimeOk

`func (o *CheckGroup) GetOutageTimeOk() (*int64, bool)`

GetOutageTimeOk returns a tuple with the OutageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutageTime

`func (o *CheckGroup) SetOutageTime(v int64)`

SetOutageTime sets OutageTime field to given value.

### HasOutageTime

`func (o *CheckGroup) HasOutageTime() bool`

HasOutageTime returns a boolean if a field has been set.

### GetLastTimer

`func (o *CheckGroup) GetLastTimer() int64`

GetLastTimer returns the LastTimer field if non-nil, zero value otherwise.

### GetLastTimerOk

`func (o *CheckGroup) GetLastTimerOk() (*int64, bool)`

GetLastTimerOk returns a tuple with the LastTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimer

`func (o *CheckGroup) SetLastTimer(v int64)`

SetLastTimer sets LastTimer field to given value.

### HasLastTimer

`func (o *CheckGroup) HasLastTimer() bool`

HasLastTimer returns a boolean if a field has been set.

### GetHealth

`func (o *CheckGroup) GetHealth() int64`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CheckGroup) GetHealthOk() (*int64, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CheckGroup) SetHealth(v int64)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CheckGroup) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHistory

`func (o *CheckGroup) GetHistory() string`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CheckGroup) GetHistoryOk() (*string, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CheckGroup) SetHistory(v string)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CheckGroup) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *CheckGroup) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *CheckGroup) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetMinHappy

`func (o *CheckGroup) GetMinHappy() int64`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *CheckGroup) GetMinHappyOk() (*int64, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *CheckGroup) SetMinHappy(v int64)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *CheckGroup) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetLastMetric

`func (o *CheckGroup) GetLastMetric() string`

GetLastMetric returns the LastMetric field if non-nil, zero value otherwise.

### GetLastMetricOk

`func (o *CheckGroup) GetLastMetricOk() (*string, bool)`

GetLastMetricOk returns a tuple with the LastMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMetric

`func (o *CheckGroup) SetLastMetric(v string)`

SetLastMetric sets LastMetric field to given value.

### HasLastMetric

`func (o *CheckGroup) HasLastMetric() bool`

HasLastMetric returns a boolean if a field has been set.

### SetLastMetricNil

`func (o *CheckGroup) SetLastMetricNil(b bool)`

 SetLastMetricNil sets the value for LastMetric to be an explicit nil

### UnsetLastMetric
`func (o *CheckGroup) UnsetLastMetric()`

UnsetLastMetric ensures that no value is present for LastMetric, not even an explicit nil
### GetSeverity

`func (o *CheckGroup) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CheckGroup) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CheckGroup) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CheckGroup) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCreateIncident

`func (o *CheckGroup) GetCreateIncident() bool`

GetCreateIncident returns the CreateIncident field if non-nil, zero value otherwise.

### GetCreateIncidentOk

`func (o *CheckGroup) GetCreateIncidentOk() (*bool, bool)`

GetCreateIncidentOk returns a tuple with the CreateIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIncident

`func (o *CheckGroup) SetCreateIncident(v bool)`

SetCreateIncident sets CreateIncident field to given value.

### HasCreateIncident

`func (o *CheckGroup) HasCreateIncident() bool`

HasCreateIncident returns a boolean if a field has been set.

### GetMuted

`func (o *CheckGroup) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *CheckGroup) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *CheckGroup) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *CheckGroup) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CheckGroup) GetCreatedBy() ActivityActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CheckGroup) GetCreatedByOk() (*ActivityActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CheckGroup) SetCreatedBy(v ActivityActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CheckGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *CheckGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CheckGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CheckGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CheckGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CheckGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CheckGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CheckGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CheckGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAvailability

`func (o *CheckGroup) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CheckGroup) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CheckGroup) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *CheckGroup) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *CheckGroup) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *CheckGroup) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetCheckType

`func (o *CheckGroup) GetCheckType() CheckCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *CheckGroup) GetCheckTypeOk() (*CheckCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *CheckGroup) SetCheckType(v CheckCheckType)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *CheckGroup) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### GetChecks

`func (o *CheckGroup) GetChecks() []int64`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *CheckGroup) GetChecksOk() (*[]int64, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *CheckGroup) SetChecks(v []int64)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *CheckGroup) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


