# InlineResponse200110NetworkProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProxyHost** | Pointer to **string** |  | [optional] 
**ProxyPort** | Pointer to **int64** |  | [optional] 
**ProxyUser** | Pointer to **NullableString** |  | [optional] 
**ProxyPassword** | Pointer to **NullableString** |  | [optional] 
**ProxyDomain** | Pointer to **string** |  | [optional] 
**ProxyWorkstation** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Owner** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 

## Methods

### NewInlineResponse200110NetworkProxy

`func NewInlineResponse200110NetworkProxy() *InlineResponse200110NetworkProxy`

NewInlineResponse200110NetworkProxy instantiates a new InlineResponse200110NetworkProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200110NetworkProxyWithDefaults

`func NewInlineResponse200110NetworkProxyWithDefaults() *InlineResponse200110NetworkProxy`

NewInlineResponse200110NetworkProxyWithDefaults instantiates a new InlineResponse200110NetworkProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200110NetworkProxy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200110NetworkProxy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200110NetworkProxy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200110NetworkProxy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200110NetworkProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200110NetworkProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200110NetworkProxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200110NetworkProxy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxyHost

`func (o *InlineResponse200110NetworkProxy) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *InlineResponse200110NetworkProxy) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *InlineResponse200110NetworkProxy) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *InlineResponse200110NetworkProxy) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *InlineResponse200110NetworkProxy) GetProxyPort() int64`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *InlineResponse200110NetworkProxy) GetProxyPortOk() (*int64, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *InlineResponse200110NetworkProxy) SetProxyPort(v int64)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *InlineResponse200110NetworkProxy) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *InlineResponse200110NetworkProxy) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *InlineResponse200110NetworkProxy) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *InlineResponse200110NetworkProxy) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *InlineResponse200110NetworkProxy) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### SetProxyUserNil

`func (o *InlineResponse200110NetworkProxy) SetProxyUserNil(b bool)`

 SetProxyUserNil sets the value for ProxyUser to be an explicit nil

### UnsetProxyUser
`func (o *InlineResponse200110NetworkProxy) UnsetProxyUser()`

UnsetProxyUser ensures that no value is present for ProxyUser, not even an explicit nil
### GetProxyPassword

`func (o *InlineResponse200110NetworkProxy) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *InlineResponse200110NetworkProxy) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *InlineResponse200110NetworkProxy) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *InlineResponse200110NetworkProxy) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### SetProxyPasswordNil

`func (o *InlineResponse200110NetworkProxy) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *InlineResponse200110NetworkProxy) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
### GetProxyDomain

`func (o *InlineResponse200110NetworkProxy) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *InlineResponse200110NetworkProxy) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *InlineResponse200110NetworkProxy) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *InlineResponse200110NetworkProxy) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *InlineResponse200110NetworkProxy) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *InlineResponse200110NetworkProxy) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *InlineResponse200110NetworkProxy) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *InlineResponse200110NetworkProxy) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### SetProxyWorkstationNil

`func (o *InlineResponse200110NetworkProxy) SetProxyWorkstationNil(b bool)`

 SetProxyWorkstationNil sets the value for ProxyWorkstation to be an explicit nil

### UnsetProxyWorkstation
`func (o *InlineResponse200110NetworkProxy) UnsetProxyWorkstation()`

UnsetProxyWorkstation ensures that no value is present for ProxyWorkstation, not even an explicit nil
### GetVisibility

`func (o *InlineResponse200110NetworkProxy) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse200110NetworkProxy) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse200110NetworkProxy) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse200110NetworkProxy) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *InlineResponse200110NetworkProxy) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InlineResponse200110NetworkProxy) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InlineResponse200110NetworkProxy) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InlineResponse200110NetworkProxy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *InlineResponse200110NetworkProxy) GetOwner() InlineResponse20040AppDeployInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InlineResponse200110NetworkProxy) GetOwnerOk() (*InlineResponse20040AppDeployInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InlineResponse200110NetworkProxy) SetOwner(v InlineResponse20040AppDeployInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InlineResponse200110NetworkProxy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


