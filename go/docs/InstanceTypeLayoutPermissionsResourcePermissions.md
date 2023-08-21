# InstanceTypeLayoutPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**NullableCheckContainer**](check_container.md) |  | [optional] 
**Sites** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInstanceTypeLayoutPermissionsResourcePermissions

`func NewInstanceTypeLayoutPermissionsResourcePermissions() *InstanceTypeLayoutPermissionsResourcePermissions`

NewInstanceTypeLayoutPermissionsResourcePermissions instantiates a new InstanceTypeLayoutPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults

`func NewInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults() *InstanceTypeLayoutPermissionsResourcePermissions`

NewInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults instantiates a new InstanceTypeLayoutPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAccount() CheckContainer`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAccountOk() (*CheckContainer, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAccount(v CheckContainer)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *InstanceTypeLayoutPermissionsResourcePermissions) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetSites

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetSites() []map[string]interface{}`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetSitesOk() (*[]map[string]interface{}, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetSites(v []map[string]interface{})`

SetSites sets Sites field to given value.

### HasSites

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *InstanceTypeLayoutPermissionsResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetPlans

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *InstanceTypeLayoutPermissionsResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


