# InlineResponse200167

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](user.md) |  | [optional] 
**IsMasterAccount** | Pointer to **bool** |  | [optional] 
**Permissions** | Pointer to [**[]InlineResponse200167Permissions**](InlineResponse200167Permissions.md) |  | [optional] 
**Appliance** | Pointer to [**InlineResponse200167Appliance**](inline_response_200_167_appliance.md) |  | [optional] 

## Methods

### NewInlineResponse200167

`func NewInlineResponse200167() *InlineResponse200167`

NewInlineResponse200167 instantiates a new InlineResponse200167 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200167WithDefaults

`func NewInlineResponse200167WithDefaults() *InlineResponse200167`

NewInlineResponse200167WithDefaults instantiates a new InlineResponse200167 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *InlineResponse200167) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineResponse200167) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineResponse200167) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineResponse200167) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIsMasterAccount

`func (o *InlineResponse200167) GetIsMasterAccount() bool`

GetIsMasterAccount returns the IsMasterAccount field if non-nil, zero value otherwise.

### GetIsMasterAccountOk

`func (o *InlineResponse200167) GetIsMasterAccountOk() (*bool, bool)`

GetIsMasterAccountOk returns a tuple with the IsMasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterAccount

`func (o *InlineResponse200167) SetIsMasterAccount(v bool)`

SetIsMasterAccount sets IsMasterAccount field to given value.

### HasIsMasterAccount

`func (o *InlineResponse200167) HasIsMasterAccount() bool`

HasIsMasterAccount returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse200167) GetPermissions() []InlineResponse200167Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse200167) GetPermissionsOk() (*[]InlineResponse200167Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse200167) SetPermissions(v []InlineResponse200167Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse200167) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAppliance

`func (o *InlineResponse200167) GetAppliance() InlineResponse200167Appliance`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *InlineResponse200167) GetApplianceOk() (*InlineResponse200167Appliance, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *InlineResponse200167) SetAppliance(v InlineResponse200167Appliance)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *InlineResponse200167) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


