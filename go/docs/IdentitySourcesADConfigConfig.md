# IdentitySourcesADConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**UseSSL** | Pointer to **string** |  | [optional] 
**BindingUsername** | Pointer to **string** |  | [optional] 
**BindingPassword** | Pointer to **string** |  | [optional] 
**RequiredGroup** | Pointer to **string** |  | [optional] 
**RequiredGroupDN** | Pointer to **string** |  | [optional] 
**SearchMemberGroups** | Pointer to **bool** |  | [optional] 
**BindingPasswordHash** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentitySourcesADConfigConfig

`func NewIdentitySourcesADConfigConfig() *IdentitySourcesADConfigConfig`

NewIdentitySourcesADConfigConfig instantiates a new IdentitySourcesADConfigConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourcesADConfigConfigWithDefaults

`func NewIdentitySourcesADConfigConfigWithDefaults() *IdentitySourcesADConfigConfig`

NewIdentitySourcesADConfigConfigWithDefaults instantiates a new IdentitySourcesADConfigConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IdentitySourcesADConfigConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdentitySourcesADConfigConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdentitySourcesADConfigConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdentitySourcesADConfigConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *IdentitySourcesADConfigConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *IdentitySourcesADConfigConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *IdentitySourcesADConfigConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *IdentitySourcesADConfigConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUseSSL

`func (o *IdentitySourcesADConfigConfig) GetUseSSL() string`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *IdentitySourcesADConfigConfig) GetUseSSLOk() (*string, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *IdentitySourcesADConfigConfig) SetUseSSL(v string)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *IdentitySourcesADConfigConfig) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetBindingUsername

`func (o *IdentitySourcesADConfigConfig) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *IdentitySourcesADConfigConfig) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *IdentitySourcesADConfigConfig) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *IdentitySourcesADConfigConfig) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *IdentitySourcesADConfigConfig) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *IdentitySourcesADConfigConfig) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *IdentitySourcesADConfigConfig) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *IdentitySourcesADConfigConfig) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetRequiredGroup

`func (o *IdentitySourcesADConfigConfig) GetRequiredGroup() string`

GetRequiredGroup returns the RequiredGroup field if non-nil, zero value otherwise.

### GetRequiredGroupOk

`func (o *IdentitySourcesADConfigConfig) GetRequiredGroupOk() (*string, bool)`

GetRequiredGroupOk returns a tuple with the RequiredGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroup

`func (o *IdentitySourcesADConfigConfig) SetRequiredGroup(v string)`

SetRequiredGroup sets RequiredGroup field to given value.

### HasRequiredGroup

`func (o *IdentitySourcesADConfigConfig) HasRequiredGroup() bool`

HasRequiredGroup returns a boolean if a field has been set.

### GetRequiredGroupDN

`func (o *IdentitySourcesADConfigConfig) GetRequiredGroupDN() string`

GetRequiredGroupDN returns the RequiredGroupDN field if non-nil, zero value otherwise.

### GetRequiredGroupDNOk

`func (o *IdentitySourcesADConfigConfig) GetRequiredGroupDNOk() (*string, bool)`

GetRequiredGroupDNOk returns a tuple with the RequiredGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroupDN

`func (o *IdentitySourcesADConfigConfig) SetRequiredGroupDN(v string)`

SetRequiredGroupDN sets RequiredGroupDN field to given value.

### HasRequiredGroupDN

`func (o *IdentitySourcesADConfigConfig) HasRequiredGroupDN() bool`

HasRequiredGroupDN returns a boolean if a field has been set.

### GetSearchMemberGroups

`func (o *IdentitySourcesADConfigConfig) GetSearchMemberGroups() bool`

GetSearchMemberGroups returns the SearchMemberGroups field if non-nil, zero value otherwise.

### GetSearchMemberGroupsOk

`func (o *IdentitySourcesADConfigConfig) GetSearchMemberGroupsOk() (*bool, bool)`

GetSearchMemberGroupsOk returns a tuple with the SearchMemberGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMemberGroups

`func (o *IdentitySourcesADConfigConfig) SetSearchMemberGroups(v bool)`

SetSearchMemberGroups sets SearchMemberGroups field to given value.

### HasSearchMemberGroups

`func (o *IdentitySourcesADConfigConfig) HasSearchMemberGroups() bool`

HasSearchMemberGroups returns a boolean if a field has been set.

### GetBindingPasswordHash

`func (o *IdentitySourcesADConfigConfig) GetBindingPasswordHash() string`

GetBindingPasswordHash returns the BindingPasswordHash field if non-nil, zero value otherwise.

### GetBindingPasswordHashOk

`func (o *IdentitySourcesADConfigConfig) GetBindingPasswordHashOk() (*string, bool)`

GetBindingPasswordHashOk returns a tuple with the BindingPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPasswordHash

`func (o *IdentitySourcesADConfigConfig) SetBindingPasswordHash(v string)`

SetBindingPasswordHash sets BindingPasswordHash field to given value.

### HasBindingPasswordHash

`func (o *IdentitySourcesADConfigConfig) HasBindingPasswordHash() bool`

HasBindingPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


