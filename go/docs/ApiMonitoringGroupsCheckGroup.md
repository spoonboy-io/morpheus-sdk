# ApiMonitoringGroupsCheckGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name scoped to your account for the check group | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**MinHappy** | Pointer to **int32** | This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. | [optional] [default to 1]
**InUptime** | Pointer to **bool** | Used to determine if check should affect account wide availability calculations | [optional] [default to true]
**Severity** | Pointer to **string** | Determines the maximum severity level this group can incur on an incident when failing | [optional] [default to "critical"]
**Active** | Pointer to **bool** | Used to determine if check group is active | [optional] [default to true]
**Checks** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewApiMonitoringGroupsCheckGroup

`func NewApiMonitoringGroupsCheckGroup(name string, ) *ApiMonitoringGroupsCheckGroup`

NewApiMonitoringGroupsCheckGroup instantiates a new ApiMonitoringGroupsCheckGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMonitoringGroupsCheckGroupWithDefaults

`func NewApiMonitoringGroupsCheckGroupWithDefaults() *ApiMonitoringGroupsCheckGroup`

NewApiMonitoringGroupsCheckGroupWithDefaults instantiates a new ApiMonitoringGroupsCheckGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiMonitoringGroupsCheckGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMonitoringGroupsCheckGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMonitoringGroupsCheckGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiMonitoringGroupsCheckGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiMonitoringGroupsCheckGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiMonitoringGroupsCheckGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiMonitoringGroupsCheckGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMinHappy

`func (o *ApiMonitoringGroupsCheckGroup) GetMinHappy() int32`

GetMinHappy returns the MinHappy field if non-nil, zero value otherwise.

### GetMinHappyOk

`func (o *ApiMonitoringGroupsCheckGroup) GetMinHappyOk() (*int32, bool)`

GetMinHappyOk returns a tuple with the MinHappy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHappy

`func (o *ApiMonitoringGroupsCheckGroup) SetMinHappy(v int32)`

SetMinHappy sets MinHappy field to given value.

### HasMinHappy

`func (o *ApiMonitoringGroupsCheckGroup) HasMinHappy() bool`

HasMinHappy returns a boolean if a field has been set.

### GetInUptime

`func (o *ApiMonitoringGroupsCheckGroup) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *ApiMonitoringGroupsCheckGroup) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *ApiMonitoringGroupsCheckGroup) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *ApiMonitoringGroupsCheckGroup) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiMonitoringGroupsCheckGroup) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiMonitoringGroupsCheckGroup) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiMonitoringGroupsCheckGroup) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiMonitoringGroupsCheckGroup) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetActive

`func (o *ApiMonitoringGroupsCheckGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiMonitoringGroupsCheckGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiMonitoringGroupsCheckGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiMonitoringGroupsCheckGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetChecks

`func (o *ApiMonitoringGroupsCheckGroup) GetChecks() []int32`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ApiMonitoringGroupsCheckGroup) GetChecksOk() (*[]int32, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ApiMonitoringGroupsCheckGroup) SetChecks(v []int32)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ApiMonitoringGroupsCheckGroup) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


