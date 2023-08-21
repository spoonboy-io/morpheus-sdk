# InlineResponse20081LoadBalancerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancer**](inline_response_200_79_loadBalancerMonitor_loadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**ServiceTypeDisplay** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProxyType** | Pointer to **NullableString** |  | [optional] 
**RedirectRewrite** | Pointer to **NullableString** |  | [optional] 
**PersistenceType** | Pointer to **NullableString** |  | [optional] 
**SslEnabled** | Pointer to **NullableString** |  | [optional] 
**SslCert** | Pointer to **NullableString** |  | [optional] 
**AccountCertificate** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**InsertXforwardedFor** | Pointer to **bool** |  | [optional] 
**PersistenceCookieName** | Pointer to **NullableString** |  | [optional] 
**PersistenceExpiresIn** | Pointer to **NullableString** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse20081LoadBalancerProfile

`func NewInlineResponse20081LoadBalancerProfile() *InlineResponse20081LoadBalancerProfile`

NewInlineResponse20081LoadBalancerProfile instantiates a new InlineResponse20081LoadBalancerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20081LoadBalancerProfileWithDefaults

`func NewInlineResponse20081LoadBalancerProfileWithDefaults() *InlineResponse20081LoadBalancerProfile`

NewInlineResponse20081LoadBalancerProfileWithDefaults instantiates a new InlineResponse20081LoadBalancerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20081LoadBalancerProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20081LoadBalancerProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20081LoadBalancerProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20081LoadBalancerProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *InlineResponse20081LoadBalancerProfile) GetLoadBalancer() InlineResponse20079LoadBalancerMonitorLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *InlineResponse20081LoadBalancerProfile) GetLoadBalancerOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *InlineResponse20081LoadBalancerProfile) SetLoadBalancer(v InlineResponse20079LoadBalancerMonitorLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *InlineResponse20081LoadBalancerProfile) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20081LoadBalancerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20081LoadBalancerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20081LoadBalancerProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20081LoadBalancerProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse20081LoadBalancerProfile) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20081LoadBalancerProfile) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20081LoadBalancerProfile) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20081LoadBalancerProfile) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetServiceType

`func (o *InlineResponse20081LoadBalancerProfile) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InlineResponse20081LoadBalancerProfile) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InlineResponse20081LoadBalancerProfile) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *InlineResponse20081LoadBalancerProfile) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceTypeDisplay

`func (o *InlineResponse20081LoadBalancerProfile) GetServiceTypeDisplay() string`

GetServiceTypeDisplay returns the ServiceTypeDisplay field if non-nil, zero value otherwise.

### GetServiceTypeDisplayOk

`func (o *InlineResponse20081LoadBalancerProfile) GetServiceTypeDisplayOk() (*string, bool)`

GetServiceTypeDisplayOk returns a tuple with the ServiceTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeDisplay

`func (o *InlineResponse20081LoadBalancerProfile) SetServiceTypeDisplay(v string)`

SetServiceTypeDisplay sets ServiceTypeDisplay field to given value.

### HasServiceTypeDisplay

`func (o *InlineResponse20081LoadBalancerProfile) HasServiceTypeDisplay() bool`

HasServiceTypeDisplay returns a boolean if a field has been set.

### GetVisibility

`func (o *InlineResponse20081LoadBalancerProfile) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse20081LoadBalancerProfile) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse20081LoadBalancerProfile) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse20081LoadBalancerProfile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20081LoadBalancerProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20081LoadBalancerProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20081LoadBalancerProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20081LoadBalancerProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *InlineResponse20081LoadBalancerProfile) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse20081LoadBalancerProfile) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse20081LoadBalancerProfile) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse20081LoadBalancerProfile) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *InlineResponse20081LoadBalancerProfile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20081LoadBalancerProfile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20081LoadBalancerProfile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20081LoadBalancerProfile) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProxyType

`func (o *InlineResponse20081LoadBalancerProfile) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *InlineResponse20081LoadBalancerProfile) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *InlineResponse20081LoadBalancerProfile) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *InlineResponse20081LoadBalancerProfile) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### SetProxyTypeNil

`func (o *InlineResponse20081LoadBalancerProfile) SetProxyTypeNil(b bool)`

 SetProxyTypeNil sets the value for ProxyType to be an explicit nil

### UnsetProxyType
`func (o *InlineResponse20081LoadBalancerProfile) UnsetProxyType()`

UnsetProxyType ensures that no value is present for ProxyType, not even an explicit nil
### GetRedirectRewrite

`func (o *InlineResponse20081LoadBalancerProfile) GetRedirectRewrite() string`

GetRedirectRewrite returns the RedirectRewrite field if non-nil, zero value otherwise.

### GetRedirectRewriteOk

`func (o *InlineResponse20081LoadBalancerProfile) GetRedirectRewriteOk() (*string, bool)`

GetRedirectRewriteOk returns a tuple with the RedirectRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRewrite

`func (o *InlineResponse20081LoadBalancerProfile) SetRedirectRewrite(v string)`

SetRedirectRewrite sets RedirectRewrite field to given value.

### HasRedirectRewrite

`func (o *InlineResponse20081LoadBalancerProfile) HasRedirectRewrite() bool`

HasRedirectRewrite returns a boolean if a field has been set.

### SetRedirectRewriteNil

`func (o *InlineResponse20081LoadBalancerProfile) SetRedirectRewriteNil(b bool)`

 SetRedirectRewriteNil sets the value for RedirectRewrite to be an explicit nil

### UnsetRedirectRewrite
`func (o *InlineResponse20081LoadBalancerProfile) UnsetRedirectRewrite()`

UnsetRedirectRewrite ensures that no value is present for RedirectRewrite, not even an explicit nil
### GetPersistenceType

`func (o *InlineResponse20081LoadBalancerProfile) GetPersistenceType() string`

GetPersistenceType returns the PersistenceType field if non-nil, zero value otherwise.

### GetPersistenceTypeOk

`func (o *InlineResponse20081LoadBalancerProfile) GetPersistenceTypeOk() (*string, bool)`

GetPersistenceTypeOk returns a tuple with the PersistenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceType

`func (o *InlineResponse20081LoadBalancerProfile) SetPersistenceType(v string)`

SetPersistenceType sets PersistenceType field to given value.

### HasPersistenceType

`func (o *InlineResponse20081LoadBalancerProfile) HasPersistenceType() bool`

HasPersistenceType returns a boolean if a field has been set.

### SetPersistenceTypeNil

`func (o *InlineResponse20081LoadBalancerProfile) SetPersistenceTypeNil(b bool)`

 SetPersistenceTypeNil sets the value for PersistenceType to be an explicit nil

### UnsetPersistenceType
`func (o *InlineResponse20081LoadBalancerProfile) UnsetPersistenceType()`

UnsetPersistenceType ensures that no value is present for PersistenceType, not even an explicit nil
### GetSslEnabled

`func (o *InlineResponse20081LoadBalancerProfile) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *InlineResponse20081LoadBalancerProfile) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *InlineResponse20081LoadBalancerProfile) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *InlineResponse20081LoadBalancerProfile) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *InlineResponse20081LoadBalancerProfile) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *InlineResponse20081LoadBalancerProfile) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *InlineResponse20081LoadBalancerProfile) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *InlineResponse20081LoadBalancerProfile) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *InlineResponse20081LoadBalancerProfile) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *InlineResponse20081LoadBalancerProfile) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *InlineResponse20081LoadBalancerProfile) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *InlineResponse20081LoadBalancerProfile) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetAccountCertificate

`func (o *InlineResponse20081LoadBalancerProfile) GetAccountCertificate() string`

GetAccountCertificate returns the AccountCertificate field if non-nil, zero value otherwise.

### GetAccountCertificateOk

`func (o *InlineResponse20081LoadBalancerProfile) GetAccountCertificateOk() (*string, bool)`

GetAccountCertificateOk returns a tuple with the AccountCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCertificate

`func (o *InlineResponse20081LoadBalancerProfile) SetAccountCertificate(v string)`

SetAccountCertificate sets AccountCertificate field to given value.

### HasAccountCertificate

`func (o *InlineResponse20081LoadBalancerProfile) HasAccountCertificate() bool`

HasAccountCertificate returns a boolean if a field has been set.

### SetAccountCertificateNil

`func (o *InlineResponse20081LoadBalancerProfile) SetAccountCertificateNil(b bool)`

 SetAccountCertificateNil sets the value for AccountCertificate to be an explicit nil

### UnsetAccountCertificate
`func (o *InlineResponse20081LoadBalancerProfile) UnsetAccountCertificate()`

UnsetAccountCertificate ensures that no value is present for AccountCertificate, not even an explicit nil
### GetEnabled

`func (o *InlineResponse20081LoadBalancerProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20081LoadBalancerProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20081LoadBalancerProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20081LoadBalancerProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineResponse20081LoadBalancerProfile) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineResponse20081LoadBalancerProfile) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineResponse20081LoadBalancerProfile) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineResponse20081LoadBalancerProfile) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *InlineResponse20081LoadBalancerProfile) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *InlineResponse20081LoadBalancerProfile) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetInsertXforwardedFor

`func (o *InlineResponse20081LoadBalancerProfile) GetInsertXforwardedFor() bool`

GetInsertXforwardedFor returns the InsertXforwardedFor field if non-nil, zero value otherwise.

### GetInsertXforwardedForOk

`func (o *InlineResponse20081LoadBalancerProfile) GetInsertXforwardedForOk() (*bool, bool)`

GetInsertXforwardedForOk returns a tuple with the InsertXforwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertXforwardedFor

`func (o *InlineResponse20081LoadBalancerProfile) SetInsertXforwardedFor(v bool)`

SetInsertXforwardedFor sets InsertXforwardedFor field to given value.

### HasInsertXforwardedFor

`func (o *InlineResponse20081LoadBalancerProfile) HasInsertXforwardedFor() bool`

HasInsertXforwardedFor returns a boolean if a field has been set.

### GetPersistenceCookieName

`func (o *InlineResponse20081LoadBalancerProfile) GetPersistenceCookieName() string`

GetPersistenceCookieName returns the PersistenceCookieName field if non-nil, zero value otherwise.

### GetPersistenceCookieNameOk

`func (o *InlineResponse20081LoadBalancerProfile) GetPersistenceCookieNameOk() (*string, bool)`

GetPersistenceCookieNameOk returns a tuple with the PersistenceCookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceCookieName

`func (o *InlineResponse20081LoadBalancerProfile) SetPersistenceCookieName(v string)`

SetPersistenceCookieName sets PersistenceCookieName field to given value.

### HasPersistenceCookieName

`func (o *InlineResponse20081LoadBalancerProfile) HasPersistenceCookieName() bool`

HasPersistenceCookieName returns a boolean if a field has been set.

### SetPersistenceCookieNameNil

`func (o *InlineResponse20081LoadBalancerProfile) SetPersistenceCookieNameNil(b bool)`

 SetPersistenceCookieNameNil sets the value for PersistenceCookieName to be an explicit nil

### UnsetPersistenceCookieName
`func (o *InlineResponse20081LoadBalancerProfile) UnsetPersistenceCookieName()`

UnsetPersistenceCookieName ensures that no value is present for PersistenceCookieName, not even an explicit nil
### GetPersistenceExpiresIn

`func (o *InlineResponse20081LoadBalancerProfile) GetPersistenceExpiresIn() string`

GetPersistenceExpiresIn returns the PersistenceExpiresIn field if non-nil, zero value otherwise.

### GetPersistenceExpiresInOk

`func (o *InlineResponse20081LoadBalancerProfile) GetPersistenceExpiresInOk() (*string, bool)`

GetPersistenceExpiresInOk returns a tuple with the PersistenceExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceExpiresIn

`func (o *InlineResponse20081LoadBalancerProfile) SetPersistenceExpiresIn(v string)`

SetPersistenceExpiresIn sets PersistenceExpiresIn field to given value.

### HasPersistenceExpiresIn

`func (o *InlineResponse20081LoadBalancerProfile) HasPersistenceExpiresIn() bool`

HasPersistenceExpiresIn returns a boolean if a field has been set.

### SetPersistenceExpiresInNil

`func (o *InlineResponse20081LoadBalancerProfile) SetPersistenceExpiresInNil(b bool)`

 SetPersistenceExpiresInNil sets the value for PersistenceExpiresIn to be an explicit nil

### UnsetPersistenceExpiresIn
`func (o *InlineResponse20081LoadBalancerProfile) UnsetPersistenceExpiresIn()`

UnsetPersistenceExpiresIn ensures that no value is present for PersistenceExpiresIn, not even an explicit nil
### GetEditable

`func (o *InlineResponse20081LoadBalancerProfile) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *InlineResponse20081LoadBalancerProfile) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *InlineResponse20081LoadBalancerProfile) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *InlineResponse20081LoadBalancerProfile) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetConfig

`func (o *InlineResponse20081LoadBalancerProfile) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineResponse20081LoadBalancerProfile) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineResponse20081LoadBalancerProfile) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineResponse20081LoadBalancerProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse20081LoadBalancerProfile) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse20081LoadBalancerProfile) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse20081LoadBalancerProfile) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse20081LoadBalancerProfile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *InlineResponse20081LoadBalancerProfile) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InlineResponse20081LoadBalancerProfile) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *InlineResponse20081LoadBalancerProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20081LoadBalancerProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20081LoadBalancerProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20081LoadBalancerProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20081LoadBalancerProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20081LoadBalancerProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20081LoadBalancerProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20081LoadBalancerProfile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


