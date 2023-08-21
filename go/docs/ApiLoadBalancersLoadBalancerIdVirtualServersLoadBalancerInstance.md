# ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VipName** | Pointer to **string** | VIP Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**VipAddress** | Pointer to **string** | VIP Address | [optional] 
**VipPort** | Pointer to **string** | VIP Port | [optional] 
**VipProtocol** | Pointer to **string** | VIP Protocol | [optional] 
**VipHostname** | Pointer to **string** | VIP Hostname | [optional] 
**SslCert** | Pointer to **int64** | SSL Client Certificate ID | [optional] 
**SslServerCert** | Pointer to **int64** | SSL Server Certificate ID | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Methods

### NewApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance

`func NewApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance() *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance`

NewApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance instantiates a new ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstanceWithDefaults

`func NewApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstanceWithDefaults() *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance`

NewApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstanceWithDefaults instantiates a new ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVipName

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipName() string`

GetVipName returns the VipName field if non-nil, zero value otherwise.

### GetVipNameOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipNameOk() (*string, bool)`

GetVipNameOk returns a tuple with the VipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipName

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetVipName(v string)`

SetVipName sets VipName field to given value.

### HasVipName

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasVipName() bool`

HasVipName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVipAddress

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipAddress() string`

GetVipAddress returns the VipAddress field if non-nil, zero value otherwise.

### GetVipAddressOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipAddressOk() (*string, bool)`

GetVipAddressOk returns a tuple with the VipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipAddress

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetVipAddress(v string)`

SetVipAddress sets VipAddress field to given value.

### HasVipAddress

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasVipAddress() bool`

HasVipAddress returns a boolean if a field has been set.

### GetVipPort

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipPort() string`

GetVipPort returns the VipPort field if non-nil, zero value otherwise.

### GetVipPortOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipPortOk() (*string, bool)`

GetVipPortOk returns a tuple with the VipPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPort

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetVipPort(v string)`

SetVipPort sets VipPort field to given value.

### HasVipPort

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasVipPort() bool`

HasVipPort returns a boolean if a field has been set.

### GetVipProtocol

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipProtocol() string`

GetVipProtocol returns the VipProtocol field if non-nil, zero value otherwise.

### GetVipProtocolOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipProtocolOk() (*string, bool)`

GetVipProtocolOk returns a tuple with the VipProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipProtocol

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetVipProtocol(v string)`

SetVipProtocol sets VipProtocol field to given value.

### HasVipProtocol

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasVipProtocol() bool`

HasVipProtocol returns a boolean if a field has been set.

### GetVipHostname

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### GetSslCert

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetSslCert() int64`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetSslCertOk() (*int64, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetSslCert(v int64)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetSslServerCert

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetSslServerCert() int64`

GetSslServerCert returns the SslServerCert field if non-nil, zero value otherwise.

### GetSslServerCertOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetSslServerCertOk() (*int64, bool)`

GetSslServerCertOk returns a tuple with the SslServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerCert

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetSslServerCert(v int64)`

SetSslServerCert sets SslServerCert field to given value.

### HasSslServerCert

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasSslServerCert() bool`

HasSslServerCert returns a boolean if a field has been set.

### GetConfig

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


