# InlineResponse200117Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 
**Owner** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 
**NetworkServer** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 
**Permissions** | Pointer to [**Permissions**](permissions.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Members** | Pointer to [**[]NetworkServerGroupMember**](NetworkServerGroupMember.md) |  | [optional] 

## Methods

### NewInlineResponse200117Group

`func NewInlineResponse200117Group() *InlineResponse200117Group`

NewInlineResponse200117Group instantiates a new InlineResponse200117Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200117GroupWithDefaults

`func NewInlineResponse200117GroupWithDefaults() *InlineResponse200117Group`

NewInlineResponse200117GroupWithDefaults instantiates a new InlineResponse200117Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200117Group) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200117Group) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200117Group) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200117Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200117Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200117Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200117Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200117Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200117Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200117Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200117Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200117Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse200117Group) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse200117Group) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *InlineResponse200117Group) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse200117Group) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse200117Group) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse200117Group) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *InlineResponse200117Group) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse200117Group) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse200117Group) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse200117Group) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *InlineResponse200117Group) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse200117Group) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse200117Group) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse200117Group) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *InlineResponse200117Group) GetAccount() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InlineResponse200117Group) GetAccountOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InlineResponse200117Group) SetAccount(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InlineResponse200117Group) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *InlineResponse200117Group) GetOwner() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InlineResponse200117Group) GetOwnerOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InlineResponse200117Group) SetOwner(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InlineResponse200117Group) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *InlineResponse200117Group) GetNetworkServer() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *InlineResponse200117Group) GetNetworkServerOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *InlineResponse200117Group) SetNetworkServer(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *InlineResponse200117Group) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse200117Group) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse200117Group) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse200117Group) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse200117Group) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200117Group) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200117Group) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200117Group) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200117Group) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMembers

`func (o *InlineResponse200117Group) GetMembers() []NetworkServerGroupMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *InlineResponse200117Group) GetMembersOk() (*[]NetworkServerGroupMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *InlineResponse200117Group) SetMembers(v []NetworkServerGroupMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *InlineResponse200117Group) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


