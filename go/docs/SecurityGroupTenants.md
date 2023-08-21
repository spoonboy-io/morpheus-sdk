# SecurityGroupTenants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityGroupTenants

`func NewSecurityGroupTenants() *SecurityGroupTenants`

NewSecurityGroupTenants instantiates a new SecurityGroupTenants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupTenantsWithDefaults

`func NewSecurityGroupTenantsWithDefaults() *SecurityGroupTenants`

NewSecurityGroupTenantsWithDefaults instantiates a new SecurityGroupTenants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupTenants) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupTenants) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupTenants) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupTenants) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupTenants) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupTenants) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupTenants) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupTenants) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCanManage

`func (o *SecurityGroupTenants) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *SecurityGroupTenants) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *SecurityGroupTenants) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *SecurityGroupTenants) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


