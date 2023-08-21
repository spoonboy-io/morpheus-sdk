# IdentitySourcesAzureADConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**DoNotIncludeSAMLRequest** | Pointer to **bool** |  | [optional] 
**SAMLSignatureMode** | Pointer to **string** |  | [optional] 
**DoNotValidateSignature** | Pointer to **bool** |  | [optional] 
**DoNotValidateStatusCode** | Pointer to **bool** |  | [optional] 
**DoNotValidateDestination** | Pointer to **bool** |  | [optional] 
**DoNotValidateIssueInstants** | Pointer to **bool** |  | [optional] 
**DoNotValidateAssertions** | Pointer to **bool** |  | [optional] 
**GivenNameAttribute** | Pointer to **string** |  | [optional] 
**SurnameAttribute** | Pointer to **string** |  | [optional] 
**EmailAttribute** | Pointer to **string** |  | [optional] 
**RequiredAttributeValue** | Pointer to **string** |  | [optional] 
**RoleAttributeName** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**AzureAppId** | Pointer to **string** |  | [optional] 
**AzureAppSecret** | Pointer to **NullableString** |  | [optional] 
**RoleLinkAttributeName** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**AzureAppSecretHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIdentitySourcesAzureADConfigConfig

`func NewIdentitySourcesAzureADConfigConfig() *IdentitySourcesAzureADConfigConfig`

NewIdentitySourcesAzureADConfigConfig instantiates a new IdentitySourcesAzureADConfigConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourcesAzureADConfigConfigWithDefaults

`func NewIdentitySourcesAzureADConfigConfigWithDefaults() *IdentitySourcesAzureADConfigConfig`

NewIdentitySourcesAzureADConfigConfigWithDefaults instantiates a new IdentitySourcesAzureADConfigConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IdentitySourcesAzureADConfigConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdentitySourcesAzureADConfigConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdentitySourcesAzureADConfigConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdentitySourcesAzureADConfigConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *IdentitySourcesAzureADConfigConfig) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *IdentitySourcesAzureADConfigConfig) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *IdentitySourcesAzureADConfigConfig) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *IdentitySourcesAzureADConfigConfig) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *IdentitySourcesAzureADConfigConfig) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *IdentitySourcesAzureADConfigConfig) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *IdentitySourcesAzureADConfigConfig) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *IdentitySourcesAzureADConfigConfig) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *IdentitySourcesAzureADConfigConfig) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *IdentitySourcesAzureADConfigConfig) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *IdentitySourcesAzureADConfigConfig) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *IdentitySourcesAzureADConfigConfig) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetDoNotValidateStatusCode

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateStatusCode() bool`

GetDoNotValidateStatusCode returns the DoNotValidateStatusCode field if non-nil, zero value otherwise.

### GetDoNotValidateStatusCodeOk

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateStatusCodeOk() (*bool, bool)`

GetDoNotValidateStatusCodeOk returns a tuple with the DoNotValidateStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateStatusCode

`func (o *IdentitySourcesAzureADConfigConfig) SetDoNotValidateStatusCode(v bool)`

SetDoNotValidateStatusCode sets DoNotValidateStatusCode field to given value.

### HasDoNotValidateStatusCode

`func (o *IdentitySourcesAzureADConfigConfig) HasDoNotValidateStatusCode() bool`

HasDoNotValidateStatusCode returns a boolean if a field has been set.

### GetDoNotValidateDestination

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateDestination() bool`

GetDoNotValidateDestination returns the DoNotValidateDestination field if non-nil, zero value otherwise.

### GetDoNotValidateDestinationOk

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateDestinationOk() (*bool, bool)`

GetDoNotValidateDestinationOk returns a tuple with the DoNotValidateDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateDestination

`func (o *IdentitySourcesAzureADConfigConfig) SetDoNotValidateDestination(v bool)`

SetDoNotValidateDestination sets DoNotValidateDestination field to given value.

### HasDoNotValidateDestination

`func (o *IdentitySourcesAzureADConfigConfig) HasDoNotValidateDestination() bool`

HasDoNotValidateDestination returns a boolean if a field has been set.

### GetDoNotValidateIssueInstants

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateIssueInstants() bool`

GetDoNotValidateIssueInstants returns the DoNotValidateIssueInstants field if non-nil, zero value otherwise.

### GetDoNotValidateIssueInstantsOk

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateIssueInstantsOk() (*bool, bool)`

GetDoNotValidateIssueInstantsOk returns a tuple with the DoNotValidateIssueInstants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateIssueInstants

`func (o *IdentitySourcesAzureADConfigConfig) SetDoNotValidateIssueInstants(v bool)`

SetDoNotValidateIssueInstants sets DoNotValidateIssueInstants field to given value.

### HasDoNotValidateIssueInstants

`func (o *IdentitySourcesAzureADConfigConfig) HasDoNotValidateIssueInstants() bool`

HasDoNotValidateIssueInstants returns a boolean if a field has been set.

### GetDoNotValidateAssertions

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateAssertions() bool`

GetDoNotValidateAssertions returns the DoNotValidateAssertions field if non-nil, zero value otherwise.

### GetDoNotValidateAssertionsOk

`func (o *IdentitySourcesAzureADConfigConfig) GetDoNotValidateAssertionsOk() (*bool, bool)`

GetDoNotValidateAssertionsOk returns a tuple with the DoNotValidateAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateAssertions

`func (o *IdentitySourcesAzureADConfigConfig) SetDoNotValidateAssertions(v bool)`

SetDoNotValidateAssertions sets DoNotValidateAssertions field to given value.

### HasDoNotValidateAssertions

`func (o *IdentitySourcesAzureADConfigConfig) HasDoNotValidateAssertions() bool`

HasDoNotValidateAssertions returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *IdentitySourcesAzureADConfigConfig) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *IdentitySourcesAzureADConfigConfig) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *IdentitySourcesAzureADConfigConfig) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *IdentitySourcesAzureADConfigConfig) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *IdentitySourcesAzureADConfigConfig) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *IdentitySourcesAzureADConfigConfig) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *IdentitySourcesAzureADConfigConfig) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *IdentitySourcesAzureADConfigConfig) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *IdentitySourcesAzureADConfigConfig) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *IdentitySourcesAzureADConfigConfig) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *IdentitySourcesAzureADConfigConfig) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *IdentitySourcesAzureADConfigConfig) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *IdentitySourcesAzureADConfigConfig) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *IdentitySourcesAzureADConfigConfig) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *IdentitySourcesAzureADConfigConfig) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *IdentitySourcesAzureADConfigConfig) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.

### GetRoleAttributeName

`func (o *IdentitySourcesAzureADConfigConfig) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *IdentitySourcesAzureADConfigConfig) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *IdentitySourcesAzureADConfigConfig) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *IdentitySourcesAzureADConfigConfig) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *IdentitySourcesAzureADConfigConfig) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *IdentitySourcesAzureADConfigConfig) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetAzureAppId

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureAppId() string`

GetAzureAppId returns the AzureAppId field if non-nil, zero value otherwise.

### GetAzureAppIdOk

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureAppIdOk() (*string, bool)`

GetAzureAppIdOk returns a tuple with the AzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppId

`func (o *IdentitySourcesAzureADConfigConfig) SetAzureAppId(v string)`

SetAzureAppId sets AzureAppId field to given value.

### HasAzureAppId

`func (o *IdentitySourcesAzureADConfigConfig) HasAzureAppId() bool`

HasAzureAppId returns a boolean if a field has been set.

### GetAzureAppSecret

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureAppSecret() string`

GetAzureAppSecret returns the AzureAppSecret field if non-nil, zero value otherwise.

### GetAzureAppSecretOk

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureAppSecretOk() (*string, bool)`

GetAzureAppSecretOk returns a tuple with the AzureAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppSecret

`func (o *IdentitySourcesAzureADConfigConfig) SetAzureAppSecret(v string)`

SetAzureAppSecret sets AzureAppSecret field to given value.

### HasAzureAppSecret

`func (o *IdentitySourcesAzureADConfigConfig) HasAzureAppSecret() bool`

HasAzureAppSecret returns a boolean if a field has been set.

### SetAzureAppSecretNil

`func (o *IdentitySourcesAzureADConfigConfig) SetAzureAppSecretNil(b bool)`

 SetAzureAppSecretNil sets the value for AzureAppSecret to be an explicit nil

### UnsetAzureAppSecret
`func (o *IdentitySourcesAzureADConfigConfig) UnsetAzureAppSecret()`

UnsetAzureAppSecret ensures that no value is present for AzureAppSecret, not even an explicit nil
### GetRoleLinkAttributeName

`func (o *IdentitySourcesAzureADConfigConfig) GetRoleLinkAttributeName() string`

GetRoleLinkAttributeName returns the RoleLinkAttributeName field if non-nil, zero value otherwise.

### GetRoleLinkAttributeNameOk

`func (o *IdentitySourcesAzureADConfigConfig) GetRoleLinkAttributeNameOk() (*string, bool)`

GetRoleLinkAttributeNameOk returns a tuple with the RoleLinkAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleLinkAttributeName

`func (o *IdentitySourcesAzureADConfigConfig) SetRoleLinkAttributeName(v string)`

SetRoleLinkAttributeName sets RoleLinkAttributeName field to given value.

### HasRoleLinkAttributeName

`func (o *IdentitySourcesAzureADConfigConfig) HasRoleLinkAttributeName() bool`

HasRoleLinkAttributeName returns a boolean if a field has been set.

### GetPublicKey

`func (o *IdentitySourcesAzureADConfigConfig) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *IdentitySourcesAzureADConfigConfig) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *IdentitySourcesAzureADConfigConfig) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *IdentitySourcesAzureADConfigConfig) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetAzureAppSecretHash

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureAppSecretHash() string`

GetAzureAppSecretHash returns the AzureAppSecretHash field if non-nil, zero value otherwise.

### GetAzureAppSecretHashOk

`func (o *IdentitySourcesAzureADConfigConfig) GetAzureAppSecretHashOk() (*string, bool)`

GetAzureAppSecretHashOk returns a tuple with the AzureAppSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppSecretHash

`func (o *IdentitySourcesAzureADConfigConfig) SetAzureAppSecretHash(v string)`

SetAzureAppSecretHash sets AzureAppSecretHash field to given value.

### HasAzureAppSecretHash

`func (o *IdentitySourcesAzureADConfigConfig) HasAzureAppSecretHash() bool`

HasAzureAppSecretHash returns a boolean if a field has been set.

### SetAzureAppSecretHashNil

`func (o *IdentitySourcesAzureADConfigConfig) SetAzureAppSecretHashNil(b bool)`

 SetAzureAppSecretHashNil sets the value for AzureAppSecretHash to be an explicit nil

### UnsetAzureAppSecretHash
`func (o *IdentitySourcesAzureADConfigConfig) UnsetAzureAppSecretHash()`

UnsetAzureAppSecretHash ensures that no value is present for AzureAppSecretHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


