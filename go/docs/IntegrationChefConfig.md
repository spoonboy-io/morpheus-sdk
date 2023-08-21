# IntegrationChefConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databags** | Pointer to [**[]IntegrationChefConfigDatabags**](IntegrationChefConfigDatabags.md) |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**ChefUser** | Pointer to **string** |  | [optional] 
**UserKey** | Pointer to **string** |  | [optional] 
**OrgKey** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ChefUseFqdn** | Pointer to **bool** |  | [optional] 
**WindowsVersion** | Pointer to **string** |  | [optional] 
**WindowsInstallUrl** | Pointer to **string** |  | [optional] 
**UserKeyHash** | Pointer to **string** |  | [optional] 
**OrgKeyHash** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationChefConfig

`func NewIntegrationChefConfig() *IntegrationChefConfig`

NewIntegrationChefConfig instantiates a new IntegrationChefConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationChefConfigWithDefaults

`func NewIntegrationChefConfigWithDefaults() *IntegrationChefConfig`

NewIntegrationChefConfigWithDefaults instantiates a new IntegrationChefConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabags

`func (o *IntegrationChefConfig) GetDatabags() []IntegrationChefConfigDatabags`

GetDatabags returns the Databags field if non-nil, zero value otherwise.

### GetDatabagsOk

`func (o *IntegrationChefConfig) GetDatabagsOk() (*[]IntegrationChefConfigDatabags, bool)`

GetDatabagsOk returns a tuple with the Databags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabags

`func (o *IntegrationChefConfig) SetDatabags(v []IntegrationChefConfigDatabags)`

SetDatabags sets Databags field to given value.

### HasDatabags

`func (o *IntegrationChefConfig) HasDatabags() bool`

HasDatabags returns a boolean if a field has been set.

### GetEndpoint

`func (o *IntegrationChefConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IntegrationChefConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IntegrationChefConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IntegrationChefConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetOrg

`func (o *IntegrationChefConfig) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *IntegrationChefConfig) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *IntegrationChefConfig) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *IntegrationChefConfig) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetChefUser

`func (o *IntegrationChefConfig) GetChefUser() string`

GetChefUser returns the ChefUser field if non-nil, zero value otherwise.

### GetChefUserOk

`func (o *IntegrationChefConfig) GetChefUserOk() (*string, bool)`

GetChefUserOk returns a tuple with the ChefUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefUser

`func (o *IntegrationChefConfig) SetChefUser(v string)`

SetChefUser sets ChefUser field to given value.

### HasChefUser

`func (o *IntegrationChefConfig) HasChefUser() bool`

HasChefUser returns a boolean if a field has been set.

### GetUserKey

`func (o *IntegrationChefConfig) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *IntegrationChefConfig) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *IntegrationChefConfig) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.

### HasUserKey

`func (o *IntegrationChefConfig) HasUserKey() bool`

HasUserKey returns a boolean if a field has been set.

### GetOrgKey

`func (o *IntegrationChefConfig) GetOrgKey() string`

GetOrgKey returns the OrgKey field if non-nil, zero value otherwise.

### GetOrgKeyOk

`func (o *IntegrationChefConfig) GetOrgKeyOk() (*string, bool)`

GetOrgKeyOk returns a tuple with the OrgKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgKey

`func (o *IntegrationChefConfig) SetOrgKey(v string)`

SetOrgKey sets OrgKey field to given value.

### HasOrgKey

`func (o *IntegrationChefConfig) HasOrgKey() bool`

HasOrgKey returns a boolean if a field has been set.

### GetVersion

`func (o *IntegrationChefConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IntegrationChefConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IntegrationChefConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IntegrationChefConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetChefUseFqdn

`func (o *IntegrationChefConfig) GetChefUseFqdn() bool`

GetChefUseFqdn returns the ChefUseFqdn field if non-nil, zero value otherwise.

### GetChefUseFqdnOk

`func (o *IntegrationChefConfig) GetChefUseFqdnOk() (*bool, bool)`

GetChefUseFqdnOk returns a tuple with the ChefUseFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefUseFqdn

`func (o *IntegrationChefConfig) SetChefUseFqdn(v bool)`

SetChefUseFqdn sets ChefUseFqdn field to given value.

### HasChefUseFqdn

`func (o *IntegrationChefConfig) HasChefUseFqdn() bool`

HasChefUseFqdn returns a boolean if a field has been set.

### GetWindowsVersion

`func (o *IntegrationChefConfig) GetWindowsVersion() string`

GetWindowsVersion returns the WindowsVersion field if non-nil, zero value otherwise.

### GetWindowsVersionOk

`func (o *IntegrationChefConfig) GetWindowsVersionOk() (*string, bool)`

GetWindowsVersionOk returns a tuple with the WindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsVersion

`func (o *IntegrationChefConfig) SetWindowsVersion(v string)`

SetWindowsVersion sets WindowsVersion field to given value.

### HasWindowsVersion

`func (o *IntegrationChefConfig) HasWindowsVersion() bool`

HasWindowsVersion returns a boolean if a field has been set.

### GetWindowsInstallUrl

`func (o *IntegrationChefConfig) GetWindowsInstallUrl() string`

GetWindowsInstallUrl returns the WindowsInstallUrl field if non-nil, zero value otherwise.

### GetWindowsInstallUrlOk

`func (o *IntegrationChefConfig) GetWindowsInstallUrlOk() (*string, bool)`

GetWindowsInstallUrlOk returns a tuple with the WindowsInstallUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsInstallUrl

`func (o *IntegrationChefConfig) SetWindowsInstallUrl(v string)`

SetWindowsInstallUrl sets WindowsInstallUrl field to given value.

### HasWindowsInstallUrl

`func (o *IntegrationChefConfig) HasWindowsInstallUrl() bool`

HasWindowsInstallUrl returns a boolean if a field has been set.

### GetUserKeyHash

`func (o *IntegrationChefConfig) GetUserKeyHash() string`

GetUserKeyHash returns the UserKeyHash field if non-nil, zero value otherwise.

### GetUserKeyHashOk

`func (o *IntegrationChefConfig) GetUserKeyHashOk() (*string, bool)`

GetUserKeyHashOk returns a tuple with the UserKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyHash

`func (o *IntegrationChefConfig) SetUserKeyHash(v string)`

SetUserKeyHash sets UserKeyHash field to given value.

### HasUserKeyHash

`func (o *IntegrationChefConfig) HasUserKeyHash() bool`

HasUserKeyHash returns a boolean if a field has been set.

### GetOrgKeyHash

`func (o *IntegrationChefConfig) GetOrgKeyHash() string`

GetOrgKeyHash returns the OrgKeyHash field if non-nil, zero value otherwise.

### GetOrgKeyHashOk

`func (o *IntegrationChefConfig) GetOrgKeyHashOk() (*string, bool)`

GetOrgKeyHashOk returns a tuple with the OrgKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgKeyHash

`func (o *IntegrationChefConfig) SetOrgKeyHash(v string)`

SetOrgKeyHash sets OrgKeyHash field to given value.

### HasOrgKeyHash

`func (o *IntegrationChefConfig) HasOrgKeyHash() bool`

HasOrgKeyHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


