# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**Owner** | Pointer to [**NullableOwner**](owner.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**ApiPort** | Pointer to **NullableString** |  | [optional] 
**AdminPort** | Pointer to **NullableString** |  | [optional] 
**SslEnabled** | Pointer to **NullableBool** |  | [optional] 
**SslCert** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Credential** | Pointer to [**Creds**](creds.md) |  | [optional] 
**Tenants** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ResourcePermission** | Pointer to [**ResourcePermissions**](resourcePermissions.md) |  | [optional] 

## Methods

### NewLoadBalancer

`func NewLoadBalancer() *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *LoadBalancer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LoadBalancer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LoadBalancer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LoadBalancer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *LoadBalancer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LoadBalancer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LoadBalancer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LoadBalancer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCloud

`func (o *LoadBalancer) GetCloud() InlineResponse20040AppDeployInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *LoadBalancer) GetCloudOk() (*InlineResponse20040AppDeployInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *LoadBalancer) SetCloud(v InlineResponse20040AppDeployInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *LoadBalancer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *LoadBalancer) GetType() InlineResponse20094Network`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancer) GetTypeOk() (*InlineResponse20094Network, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancer) SetType(v InlineResponse20094Network)`

SetType sets Type field to given value.

### HasType

`func (o *LoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *LoadBalancer) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *LoadBalancer) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *LoadBalancer) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *LoadBalancer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *LoadBalancer) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *LoadBalancer) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetVisibility

`func (o *LoadBalancer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *LoadBalancer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *LoadBalancer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *LoadBalancer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *LoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *LoadBalancer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LoadBalancer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LoadBalancer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LoadBalancer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *LoadBalancer) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancer) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancer) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoadBalancer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *LoadBalancer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LoadBalancer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LoadBalancer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LoadBalancer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *LoadBalancer) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *LoadBalancer) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetIp

`func (o *LoadBalancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LoadBalancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LoadBalancer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LoadBalancer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *LoadBalancer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *LoadBalancer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *LoadBalancer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *LoadBalancer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *LoadBalancer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *LoadBalancer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *LoadBalancer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *LoadBalancer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *LoadBalancer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *LoadBalancer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *LoadBalancer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *LoadBalancer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetApiPort

`func (o *LoadBalancer) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *LoadBalancer) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *LoadBalancer) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *LoadBalancer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *LoadBalancer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *LoadBalancer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *LoadBalancer) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *LoadBalancer) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *LoadBalancer) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *LoadBalancer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *LoadBalancer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *LoadBalancer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetSslEnabled

`func (o *LoadBalancer) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *LoadBalancer) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *LoadBalancer) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *LoadBalancer) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *LoadBalancer) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *LoadBalancer) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *LoadBalancer) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *LoadBalancer) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *LoadBalancer) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *LoadBalancer) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *LoadBalancer) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *LoadBalancer) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetConfig

`func (o *LoadBalancer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoadBalancer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoadBalancer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoadBalancer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *LoadBalancer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *LoadBalancer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *LoadBalancer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *LoadBalancer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *LoadBalancer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LoadBalancer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LoadBalancer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *LoadBalancer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCredential

`func (o *LoadBalancer) GetCredential() Creds`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *LoadBalancer) GetCredentialOk() (*Creds, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *LoadBalancer) SetCredential(v Creds)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *LoadBalancer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *LoadBalancer) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *LoadBalancer) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *LoadBalancer) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *LoadBalancer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *LoadBalancer) GetResourcePermission() ResourcePermissions`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *LoadBalancer) GetResourcePermissionOk() (*ResourcePermissions, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *LoadBalancer) SetResourcePermission(v ResourcePermissions)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *LoadBalancer) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


