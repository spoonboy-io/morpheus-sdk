# ApiMonitoringAppsIdMonitorApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the check app | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Severity** | Pointer to **string** | Severity level of incidents that are created when this check fails | [optional] [default to "critical"]
**Active** | Pointer to **bool** | Used to determine if check app is active | [optional] [default to true]
**Checks** | Pointer to **[]int32** |  | [optional] 
**CheckGroups** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewApiMonitoringAppsIdMonitorApp

`func NewApiMonitoringAppsIdMonitorApp() *ApiMonitoringAppsIdMonitorApp`

NewApiMonitoringAppsIdMonitorApp instantiates a new ApiMonitoringAppsIdMonitorApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringAppsIdMonitorAppWithDefaults

`func NewApiMonitoringAppsIdMonitorAppWithDefaults() *ApiMonitoringAppsIdMonitorApp`

NewApiMonitoringAppsIdMonitorAppWithDefaults instantiates a new ApiMonitoringAppsIdMonitorApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiMonitoringAppsIdMonitorApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringAppsIdMonitorApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiMonitoringAppsIdMonitorApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiMonitoringAppsIdMonitorApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiMonitoringAppsIdMonitorApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiMonitoringAppsIdMonitorApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInUptime

`func (o *ApiMonitoringAppsIdMonitorApp) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ApiMonitoringAppsIdMonitorApp) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ApiMonitoringAppsIdMonitorApp) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiMonitoringAppsIdMonitorApp) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiMonitoringAppsIdMonitorApp) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiMonitoringAppsIdMonitorApp) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetActive

`func (o *ApiMonitoringAppsIdMonitorApp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiMonitoringAppsIdMonitorApp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiMonitoringAppsIdMonitorApp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMonitoringAppsIdMonitorApp) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMonitoringAppsIdMonitorApp) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMonitoringAppsIdMonitorApp) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCheckGroups

`func (o *ApiMonitoringAppsIdMonitorApp) GetCheckGroups() []int32`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *ApiMonitoringAppsIdMonitorApp) GetCheckGroupsOk() (*[]int32, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *ApiMonitoringAppsIdMonitorApp) SetCheckGroups(v []int32)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *ApiMonitoringAppsIdMonitorApp) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


