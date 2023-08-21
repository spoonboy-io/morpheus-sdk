# ApiMonitoringAppsMonitorApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check app | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level of incidents that are created when this check fails | [optional] [default to "critical"]
**Active** | Pointer to **bool** | Used to determine if check app is active | [optional] [default to true]
**Checks** | Pointer to **[]int32** |  | [optional] 
**CheckGroups** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewApiMonitoringAppsMonitorApp

`func NewApiMonitoringAppsMonitorApp(name string, ) *ApiMonitoringAppsMonitorApp`

NewApiMonitoringAppsMonitorApp instantiates a new ApiMonitoringAppsMonitorApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringAppsMonitorAppWithDefaults

`func NewApiMonitoringAppsMonitorAppWithDefaults() *ApiMonitoringAppsMonitorApp`

NewApiMonitoringAppsMonitorAppWithDefaults instantiates a new ApiMonitoringAppsMonitorApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiMonitoringAppsMonitorApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringAppsMonitorApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringAppsMonitorApp) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiMonitoringAppsMonitorApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiMonitoringAppsMonitorApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiMonitoringAppsMonitorApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiMonitoringAppsMonitorApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInUptime

`func (o *ApiMonitoringAppsMonitorApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ApiMonitoringAppsMonitorApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ApiMonitoringAppsMonitorApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ApiMonitoringAppsMonitorApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiMonitoringAppsMonitorApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiMonitoringAppsMonitorApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiMonitoringAppsMonitorApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiMonitoringAppsMonitorApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetActive

`func (o *ApiMonitoringAppsMonitorApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiMonitoringAppsMonitorApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiMonitoringAppsMonitorApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiMonitoringAppsMonitorApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMonitoringAppsMonitorApp) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMonitoringAppsMonitorApp) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMonitoringAppsMonitorApp) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMonitoringAppsMonitorApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *ApiMonitoringAppsMonitorApp) GetCheckGroups() []int32`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *ApiMonitoringAppsMonitorApp) GetCheckGroupsOk() (*[]int32, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *ApiMonitoringAppsMonitorApp) SetCheckGroups(v []int32)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *ApiMonitoringAppsMonitorApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


