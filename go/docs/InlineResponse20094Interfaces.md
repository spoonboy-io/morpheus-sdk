# InlineResponse20094Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**InterfaceType** | Pointer to **NullableString** |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**ExternalLink** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 

## Methods

### NewInlineResponse20094Interfaces

`func NewInlineResponse20094Interfaces() *InlineResponse20094Interfaces`

NewInlineResponse20094Interfaces instantiates a new InlineResponse20094Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20094InterfacesWithDefaults

`func NewInlineResponse20094InterfacesWithDefaults() *InlineResponse20094Interfaces`

NewInlineResponse20094InterfacesWithDefaults instantiates a new InlineResponse20094Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20094Interfaces) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20094Interfaces) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20094Interfaces) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20094Interfaces) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20094Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20094Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20094Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20094Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20094Interfaces) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20094Interfaces) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20094Interfaces) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20094Interfaces) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *InlineResponse20094Interfaces) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InlineResponse20094Interfaces) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetInterfaceType

`func (o *InlineResponse20094Interfaces) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *InlineResponse20094Interfaces) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *InlineResponse20094Interfaces) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *InlineResponse20094Interfaces) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### SetInterfaceTypeNil

`func (o *InlineResponse20094Interfaces) SetInterfaceTypeNil(b bool)`

 SetInterfaceTypeNil sets the value for InterfaceType to be an explicit nil

### UnsetInterfaceType
`func (o *InlineResponse20094Interfaces) UnsetInterfaceType()`

UnsetInterfaceType ensures that no value is present for InterfaceType, not even an explicit nil
### GetNetworkPosition

`func (o *InlineResponse20094Interfaces) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *InlineResponse20094Interfaces) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *InlineResponse20094Interfaces) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *InlineResponse20094Interfaces) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *InlineResponse20094Interfaces) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *InlineResponse20094Interfaces) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetIpAddress

`func (o *InlineResponse20094Interfaces) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineResponse20094Interfaces) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineResponse20094Interfaces) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineResponse20094Interfaces) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20094Interfaces) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20094Interfaces) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20094Interfaces) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20094Interfaces) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetExternalLink

`func (o *InlineResponse20094Interfaces) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *InlineResponse20094Interfaces) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *InlineResponse20094Interfaces) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *InlineResponse20094Interfaces) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### SetExternalLinkNil

`func (o *InlineResponse20094Interfaces) SetExternalLinkNil(b bool)`

 SetExternalLinkNil sets the value for ExternalLink to be an explicit nil

### UnsetExternalLink
`func (o *InlineResponse20094Interfaces) UnsetExternalLink()`

UnsetExternalLink ensures that no value is present for ExternalLink, not even an explicit nil
### GetEnabled

`func (o *InlineResponse20094Interfaces) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20094Interfaces) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20094Interfaces) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20094Interfaces) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse20094Interfaces) GetNetwork() InlineResponse20094Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse20094Interfaces) GetNetworkOk() (*InlineResponse20094Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse20094Interfaces) SetNetwork(v InlineResponse20094Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse20094Interfaces) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


