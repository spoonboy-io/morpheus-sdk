# IdentitySourcesSAMLConfigRoleMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRoleName** | Pointer to **string** |  | [optional] 
**SourceRoleFqn** | Pointer to **NullableString** |  | [optional] 
**MappedRole** | Pointer to [**IdentitySourcesLDAPConfigDefaultAccountRole**](identitySourcesLDAPConfig_defaultAccountRole.md) |  | [optional] 

## Methods

### NewIdentitySourcesSAMLConfigRoleMappings

`func NewIdentitySourcesSAMLConfigRoleMappings() *IdentitySourcesSAMLConfigRoleMappings`

NewIdentitySourcesSAMLConfigRoleMappings instantiates a new IdentitySourcesSAMLConfigRoleMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourcesSAMLConfigRoleMappingsWithDefaults

`func NewIdentitySourcesSAMLConfigRoleMappingsWithDefaults() *IdentitySourcesSAMLConfigRoleMappings`

NewIdentitySourcesSAMLConfigRoleMappingsWithDefaults instantiates a new IdentitySourcesSAMLConfigRoleMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRoleName

`func (o *IdentitySourcesSAMLConfigRoleMappings) GetSourceRoleName() string`

GetSourceRoleName returns the SourceRoleName field if non-nil, zero value otherwise.

### GetSourceRoleNameOk

`func (o *IdentitySourcesSAMLConfigRoleMappings) GetSourceRoleNameOk() (*string, bool)`

GetSourceRoleNameOk returns a tuple with the SourceRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoleName

`func (o *IdentitySourcesSAMLConfigRoleMappings) SetSourceRoleName(v string)`

SetSourceRoleName sets SourceRoleName field to given value.

### HasSourceRoleName

`func (o *IdentitySourcesSAMLConfigRoleMappings) HasSourceRoleName() bool`

HasSourceRoleName returns a boolean if a field has been set.

### GetSourceRoleFqn

`func (o *IdentitySourcesSAMLConfigRoleMappings) GetSourceRoleFqn() string`

GetSourceRoleFqn returns the SourceRoleFqn field if non-nil, zero value otherwise.

### GetSourceRoleFqnOk

`func (o *IdentitySourcesSAMLConfigRoleMappings) GetSourceRoleFqnOk() (*string, bool)`

GetSourceRoleFqnOk returns a tuple with the SourceRoleFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoleFqn

`func (o *IdentitySourcesSAMLConfigRoleMappings) SetSourceRoleFqn(v string)`

SetSourceRoleFqn sets SourceRoleFqn field to given value.

### HasSourceRoleFqn

`func (o *IdentitySourcesSAMLConfigRoleMappings) HasSourceRoleFqn() bool`

HasSourceRoleFqn returns a boolean if a field has been set.

### SetSourceRoleFqnNil

`func (o *IdentitySourcesSAMLConfigRoleMappings) SetSourceRoleFqnNil(b bool)`

 SetSourceRoleFqnNil sets the value for SourceRoleFqn to be an explicit nil

### UnsetSourceRoleFqn
`func (o *IdentitySourcesSAMLConfigRoleMappings) UnsetSourceRoleFqn()`

UnsetSourceRoleFqn ensures that no value is present for SourceRoleFqn, not even an explicit nil
### GetMappedRole

`func (o *IdentitySourcesSAMLConfigRoleMappings) GetMappedRole() IdentitySourcesLDAPConfigDefaultAccountRole`

GetMappedRole returns the MappedRole field if non-nil, zero value otherwise.

### GetMappedRoleOk

`func (o *IdentitySourcesSAMLConfigRoleMappings) GetMappedRoleOk() (*IdentitySourcesLDAPConfigDefaultAccountRole, bool)`

GetMappedRoleOk returns a tuple with the MappedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedRole

`func (o *IdentitySourcesSAMLConfigRoleMappings) SetMappedRole(v IdentitySourcesLDAPConfigDefaultAccountRole)`

SetMappedRole sets MappedRole field to given value.

### HasMappedRole

`func (o *IdentitySourcesSAMLConfigRoleMappings) HasMappedRole() bool`

HasMappedRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


