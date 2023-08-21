# ApiMonitoringGroupsIdCheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the check group | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**MinHappy** | Pointer to **int32** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. | [optional] [default to 1]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Severity** | Pointer to **string** | Determines the maximum severity level this group can incur on an incident when failing | [optional] [default to "critical"]
**Active** | Pointer to **bool** | Used to determine if check group is active | [optional] [default to true]
**Checks** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewApiMonitoringGroupsIdCheckGroup

`func NewApiMonitoringGroupsIdCheckGroup() *ApiMonitoringGroupsIdCheckGroup`

NewApiMonitoringGroupsIdCheckGroup instantiates a new ApiMonitoringGroupsIdCheckGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringGroupsIdCheckGroupWithDefaults

`func NewApiMonitoringGroupsIdCheckGroupWithDefaults() *ApiMonitoringGroupsIdCheckGroup`

NewApiMonitoringGroupsIdCheckGroupWithDefaults instantiates a new ApiMonitoringGroupsIdCheckGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiMonitoringGroupsIdCheckGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringGroupsIdCheckGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiMonitoringGroupsIdCheckGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiMonitoringGroupsIdCheckGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiMonitoringGroupsIdCheckGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiMonitoringGroupsIdCheckGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMinHappy

`func (o *ApiMonitoringGroupsIdCheckGroup) GetMinHappy() int32`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetMinHappyOk() (*int32, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *ApiMonitoringGroupsIdCheckGroup) SetMinHappy(v int32)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *ApiMonitoringGroupsIdCheckGroup) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetInUptime

`func (o *ApiMonitoringGroupsIdCheckGroup) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ApiMonitoringGroupsIdCheckGroup) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ApiMonitoringGroupsIdCheckGroup) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiMonitoringGroupsIdCheckGroup) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiMonitoringGroupsIdCheckGroup) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiMonitoringGroupsIdCheckGroup) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetActive

`func (o *ApiMonitoringGroupsIdCheckGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiMonitoringGroupsIdCheckGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiMonitoringGroupsIdCheckGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMonitoringGroupsIdCheckGroup) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMonitoringGroupsIdCheckGroup) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMonitoringGroupsIdCheckGroup) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMonitoringGroupsIdCheckGroup) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


