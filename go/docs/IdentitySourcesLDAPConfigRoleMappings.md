# IdentitySourcesLDAPConfigRoleMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRoleName** | Pointer to **NullableString** |  | [optional] 
**SourceRoleFqn** | Pointer to **string** |  | [optional] 
**MappedRole** | Pointer to [**IdentitySourcesLDAPConfigDefaultAccountRole**](identitySourcesLDAPConfig_defaultAccountRole.md) |  | [optional] 

## Methods

### NewIdentitySourcesLDAPConfigRoleMappings

`func NewIdentitySourcesLDAPConfigRoleMappings() *IdentitySourcesLDAPConfigRoleMappings`

NewIdentitySourcesLDAPConfigRoleMappings instantiates a new IdentitySourcesLDAPConfigRoleMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourcesLDAPConfigRoleMappingsWithDefaults

`func NewIdentitySourcesLDAPConfigRoleMappingsWithDefaults() *IdentitySourcesLDAPConfigRoleMappings`

NewIdentitySourcesLDAPConfigRoleMappingsWithDefaults instantiates a new IdentitySourcesLDAPConfigRoleMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRoleName

`func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleName() string`

GetSourceRoleName returns the SourceRoleName field if non-nil, zero value otherwise.

### GetSourceRoleNameOk

`func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleNameOk() (*string, bool)`

GetSourceRoleNameOk returns a tuple with the SourceRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoleName

`func (o *IdentitySourcesLDAPConfigRoleMappings) SetSourceRoleName(v string)`

SetSourceRoleName sets SourceRoleName field to given value.

### HasSourceRoleName

`func (o *IdentitySourcesLDAPConfigRoleMappings) HasSourceRoleName() bool`

HasSourceRoleName returns a boolean if a field has been set.

### SetSourceRoleNameNil

`func (o *IdentitySourcesLDAPConfigRoleMappings) SetSourceRoleNameNil(b bool)`

 SetSourceRoleNameNil sets the value for SourceRoleName to be an explicit nil

### UnsetSourceRoleName
`func (o *IdentitySourcesLDAPConfigRoleMappings) UnsetSourceRoleName()`

UnsetSourceRoleName ensures that no value is present for SourceRoleName, not even an explicit nil
### GetSourceRoleFqn

`func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleFqn() string`

GetSourceRoleFqn returns the SourceRoleFqn field if non-nil, zero value otherwise.

### GetSourceRoleFqnOk

`func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleFqnOk() (*string, bool)`

GetSourceRoleFqnOk returns a tuple with the SourceRoleFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoleFqn

`func (o *IdentitySourcesLDAPConfigRoleMappings) SetSourceRoleFqn(v string)`

SetSourceRoleFqn sets SourceRoleFqn field to given value.

### HasSourceRoleFqn

`func (o *IdentitySourcesLDAPConfigRoleMappings) HasSourceRoleFqn() bool`

HasSourceRoleFqn returns a boolean if a field has been set.

### GetMappedRole

`func (o *IdentitySourcesLDAPConfigRoleMappings) GetMappedRole() IdentitySourcesLDAPConfigDefaultAccountRole`

GetMappedRole returns the MappedRole field if non-nil, zero value otherwise.

### GetMappedRoleOk

`func (o *IdentitySourcesLDAPConfigRoleMappings) GetMappedRoleOk() (*IdentitySourcesLDAPConfigDefaultAccountRole, bool)`

GetMappedRoleOk returns a tuple with the MappedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedRole

`func (o *IdentitySourcesLDAPConfigRoleMappings) SetMappedRole(v IdentitySourcesLDAPConfigDefaultAccountRole)`

SetMappedRole sets MappedRole field to given value.

### HasMappedRole

`func (o *IdentitySourcesLDAPConfigRoleMappings) HasMappedRole() bool`

HasMappedRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


