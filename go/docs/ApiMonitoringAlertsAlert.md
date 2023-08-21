# ApiMonitoringAlertsAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the alert | 
**MinDuration** | Pointer to **int32** | Duration in minutes of the delay before sending notification(s) | [optional] [default to 0]
**MinSeverity** | Pointer to **string** | Severity level threshold for sending notifications. | [optional] [default to "critical"]
**Active** | Pointer to **bool** | Set to false to disable notifications | [optional] [default to true]
**AllChecks** | Pointer to **bool** | Trigger for all checks | [optional] [default to false]
**AllGroups** | Pointer to **bool** | Trigger for all check groups | [optional] [default to false]
**AllApps** | Pointer to **bool** | Trigger for all monitor apps | [optional] [default to false]
**Checks** | Pointer to **[]int32** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Apps** | Pointer to **[]int32** |  | [optional] 
**Contacts** | Pointer to [**[]ApiMonitoringAlertsAlertContacts**](ApiMonitoringAlertsAlertContacts.md) |  | [optional] 

## Methods

### NewApiMonitoringAlertsAlert

`func NewApiMonitoringAlertsAlert(name string, ) *ApiMonitoringAlertsAlert`

NewApiMonitoringAlertsAlert instantiates a new ApiMonitoringAlertsAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringAlertsAlertWithDefaults

`func NewApiMonitoringAlertsAlertWithDefaults() *ApiMonitoringAlertsAlert`

NewApiMonitoringAlertsAlertWithDefaults instantiates a new ApiMonitoringAlertsAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiMonitoringAlertsAlert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringAlertsAlert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringAlertsAlert) SetName(v string)`

SetName sets Name field to given value.


### GetMinDuration

`func (o *ApiMonitoringAlertsAlert) GetMinDuration() int32`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *ApiMonitoringAlertsAlert) GetMinDurationOk() (*int32, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *ApiMonitoringAlertsAlert) SetMinDuration(v int32)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *ApiMonitoringAlertsAlert) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetMinSeverity

`func (o *ApiMonitoringAlertsAlert) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *ApiMonitoringAlertsAlert) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *ApiMonitoringAlertsAlert) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *ApiMonitoringAlertsAlert) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetActive

`func (o *ApiMonitoringAlertsAlert) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiMonitoringAlertsAlert) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiMonitoringAlertsAlert) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiMonitoringAlertsAlert) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAllChecks

`func (o *ApiMonitoringAlertsAlert) GetAllChecks() bool`

GetAllChecks returns the AllChecks field if non-nil, zero value otherwise.

### GetAllChecksOk

`func (o *ApiMonitoringAlertsAlert) GetAllChecksOk() (*bool, bool)`

GetAllChecksOk returns a tuple with the AllChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllChecks

`func (o *ApiMonitoringAlertsAlert) SetAllChecks(v bool)`

SetAllChecks sets AllChecks field to given value.

### HasAllChecks

`func (o *ApiMonitoringAlertsAlert) HasAllChecks() bool`

HasAllChecks returns a boolean if a field has been set.

### GetAllGroups

`func (o *ApiMonitoringAlertsAlert) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *ApiMonitoringAlertsAlert) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *ApiMonitoringAlertsAlert) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *ApiMonitoringAlertsAlert) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetAllApps

`func (o *ApiMonitoringAlertsAlert) GetAllApps() bool`

GetAllApps returns the AllApps field if non-nil, zero value otherwise.

### GetAllAppsOk

`func (o *ApiMonitoringAlertsAlert) GetAllAppsOk() (*bool, bool)`

GetAllAppsOk returns a tuple with the AllApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllApps

`func (o *ApiMonitoringAlertsAlert) SetAllApps(v bool)`

SetAllApps sets AllApps field to given value.

### HasAllApps

`func (o *ApiMonitoringAlertsAlert) HasAllApps() bool`

HasAllApps returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMonitoringAlertsAlert) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMonitoringAlertsAlert) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMonitoringAlertsAlert) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMonitoringAlertsAlert) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetGroups

`func (o *ApiMonitoringAlertsAlert) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ApiMonitoringAlertsAlert) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ApiMonitoringAlertsAlert) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ApiMonitoringAlertsAlert) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetApps

`func (o *ApiMonitoringAlertsAlert) GetApps() []int32`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ApiMonitoringAlertsAlert) GetAppsOk() (*[]int32, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ApiMonitoringAlertsAlert) SetApps(v []int32)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ApiMonitoringAlertsAlert) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetContacts

`func (o *ApiMonitoringAlertsAlert) GetContacts() []ApiMonitoringAlertsAlertContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ApiMonitoringAlertsAlert) GetContactsOk() (*[]ApiMonitoringAlertsAlertContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ApiMonitoringAlertsAlert) SetContacts(v []ApiMonitoringAlertsAlertContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ApiMonitoringAlertsAlert) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


