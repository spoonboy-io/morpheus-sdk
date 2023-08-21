# IdentitySourcesSAMLConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleAttributeName** | Pointer to **string** |  | [optional] 
**RequiredAttributeValue** | Pointer to **string** |  | [optional] 
**GivenNameAttribute** | Pointer to **string** |  | [optional] 
**SurnameAttribute** | Pointer to **string** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**DoNotIncludeSAMLRequest** | Pointer to **bool** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**EmailAttribute** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**DoNotValidateSignature** | Pointer to **bool** |  | [optional] 
**DoNotValidateStatusCode** | Pointer to **bool** |  | [optional] 
**DoNotValidateDestination** | Pointer to **bool** |  | [optional] 
**DoNotValidateIssueInstants** | Pointer to **bool** |  | [optional] 
**DoNotValidateAssertions** | Pointer to **bool** |  | [optional] 
**DoNotValidateAuthStatements** | Pointer to **bool** |  | [optional] 
**DoNotValidateSubject** | Pointer to **bool** |  | [optional] 
**DoNotValidateConditions** | Pointer to **bool** |  | [optional] 
**DoNotValidateAudiences** | Pointer to **bool** |  | [optional] 
**DoNotValidateSubjectRecipients** | Pointer to **bool** |  | [optional] 
**SAMLSignatureMode** | Pointer to **string** |  | [optional] 
**ForceAuthn** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdentitySourcesSAMLConfigConfig

`func NewIdentitySourcesSAMLConfigConfig() *IdentitySourcesSAMLConfigConfig`

NewIdentitySourcesSAMLConfigConfig instantiates a new IdentitySourcesSAMLConfigConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourcesSAMLConfigConfigWithDefaults

`func NewIdentitySourcesSAMLConfigConfigWithDefaults() *IdentitySourcesSAMLConfigConfig`

NewIdentitySourcesSAMLConfigConfigWithDefaults instantiates a new IdentitySourcesSAMLConfigConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleAttributeName

`func (o *IdentitySourcesSAMLConfigConfig) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *IdentitySourcesSAMLConfigConfig) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *IdentitySourcesSAMLConfigConfig) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *IdentitySourcesSAMLConfigConfig) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *IdentitySourcesSAMLConfigConfig) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *IdentitySourcesSAMLConfigConfig) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *IdentitySourcesSAMLConfigConfig) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *IdentitySourcesSAMLConfigConfig) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *IdentitySourcesSAMLConfigConfig) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *IdentitySourcesSAMLConfigConfig) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *IdentitySourcesSAMLConfigConfig) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *IdentitySourcesSAMLConfigConfig) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *IdentitySourcesSAMLConfigConfig) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *IdentitySourcesSAMLConfigConfig) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *IdentitySourcesSAMLConfigConfig) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *IdentitySourcesSAMLConfigConfig) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *IdentitySourcesSAMLConfigConfig) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *IdentitySourcesSAMLConfigConfig) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *IdentitySourcesSAMLConfigConfig) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *IdentitySourcesSAMLConfigConfig) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetPublicKey

`func (o *IdentitySourcesSAMLConfigConfig) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *IdentitySourcesSAMLConfigConfig) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *IdentitySourcesSAMLConfigConfig) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *IdentitySourcesSAMLConfigConfig) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *IdentitySourcesSAMLConfigConfig) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *IdentitySourcesSAMLConfigConfig) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *IdentitySourcesSAMLConfigConfig) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *IdentitySourcesSAMLConfigConfig) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetUrl

`func (o *IdentitySourcesSAMLConfigConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdentitySourcesSAMLConfigConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdentitySourcesSAMLConfigConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdentitySourcesSAMLConfigConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetDoNotValidateStatusCode

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateStatusCode() bool`

GetDoNotValidateStatusCode returns the DoNotValidateStatusCode field if non-nil, zero value otherwise.

### GetDoNotValidateStatusCodeOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateStatusCodeOk() (*bool, bool)`

GetDoNotValidateStatusCodeOk returns a tuple with the DoNotValidateStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateStatusCode

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateStatusCode(v bool)`

SetDoNotValidateStatusCode sets DoNotValidateStatusCode field to given value.

### HasDoNotValidateStatusCode

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateStatusCode() bool`

HasDoNotValidateStatusCode returns a boolean if a field has been set.

### GetDoNotValidateDestination

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateDestination() bool`

GetDoNotValidateDestination returns the DoNotValidateDestination field if non-nil, zero value otherwise.

### GetDoNotValidateDestinationOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateDestinationOk() (*bool, bool)`

GetDoNotValidateDestinationOk returns a tuple with the DoNotValidateDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateDestination

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateDestination(v bool)`

SetDoNotValidateDestination sets DoNotValidateDestination field to given value.

### HasDoNotValidateDestination

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateDestination() bool`

HasDoNotValidateDestination returns a boolean if a field has been set.

### GetDoNotValidateIssueInstants

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateIssueInstants() bool`

GetDoNotValidateIssueInstants returns the DoNotValidateIssueInstants field if non-nil, zero value otherwise.

### GetDoNotValidateIssueInstantsOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateIssueInstantsOk() (*bool, bool)`

GetDoNotValidateIssueInstantsOk returns a tuple with the DoNotValidateIssueInstants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateIssueInstants

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateIssueInstants(v bool)`

SetDoNotValidateIssueInstants sets DoNotValidateIssueInstants field to given value.

### HasDoNotValidateIssueInstants

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateIssueInstants() bool`

HasDoNotValidateIssueInstants returns a boolean if a field has been set.

### GetDoNotValidateAssertions

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateAssertions() bool`

GetDoNotValidateAssertions returns the DoNotValidateAssertions field if non-nil, zero value otherwise.

### GetDoNotValidateAssertionsOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateAssertionsOk() (*bool, bool)`

GetDoNotValidateAssertionsOk returns a tuple with the DoNotValidateAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateAssertions

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateAssertions(v bool)`

SetDoNotValidateAssertions sets DoNotValidateAssertions field to given value.

### HasDoNotValidateAssertions

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateAssertions() bool`

HasDoNotValidateAssertions returns a boolean if a field has been set.

### GetDoNotValidateAuthStatements

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateAuthStatements() bool`

GetDoNotValidateAuthStatements returns the DoNotValidateAuthStatements field if non-nil, zero value otherwise.

### GetDoNotValidateAuthStatementsOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateAuthStatementsOk() (*bool, bool)`

GetDoNotValidateAuthStatementsOk returns a tuple with the DoNotValidateAuthStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateAuthStatements

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateAuthStatements(v bool)`

SetDoNotValidateAuthStatements sets DoNotValidateAuthStatements field to given value.

### HasDoNotValidateAuthStatements

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateAuthStatements() bool`

HasDoNotValidateAuthStatements returns a boolean if a field has been set.

### GetDoNotValidateSubject

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateSubject() bool`

GetDoNotValidateSubject returns the DoNotValidateSubject field if non-nil, zero value otherwise.

### GetDoNotValidateSubjectOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateSubjectOk() (*bool, bool)`

GetDoNotValidateSubjectOk returns a tuple with the DoNotValidateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSubject

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateSubject(v bool)`

SetDoNotValidateSubject sets DoNotValidateSubject field to given value.

### HasDoNotValidateSubject

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateSubject() bool`

HasDoNotValidateSubject returns a boolean if a field has been set.

### GetDoNotValidateConditions

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateConditions() bool`

GetDoNotValidateConditions returns the DoNotValidateConditions field if non-nil, zero value otherwise.

### GetDoNotValidateConditionsOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateConditionsOk() (*bool, bool)`

GetDoNotValidateConditionsOk returns a tuple with the DoNotValidateConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateConditions

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateConditions(v bool)`

SetDoNotValidateConditions sets DoNotValidateConditions field to given value.

### HasDoNotValidateConditions

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateConditions() bool`

HasDoNotValidateConditions returns a boolean if a field has been set.

### GetDoNotValidateAudiences

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateAudiences() bool`

GetDoNotValidateAudiences returns the DoNotValidateAudiences field if non-nil, zero value otherwise.

### GetDoNotValidateAudiencesOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateAudiencesOk() (*bool, bool)`

GetDoNotValidateAudiencesOk returns a tuple with the DoNotValidateAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateAudiences

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateAudiences(v bool)`

SetDoNotValidateAudiences sets DoNotValidateAudiences field to given value.

### HasDoNotValidateAudiences

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateAudiences() bool`

HasDoNotValidateAudiences returns a boolean if a field has been set.

### GetDoNotValidateSubjectRecipients

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateSubjectRecipients() bool`

GetDoNotValidateSubjectRecipients returns the DoNotValidateSubjectRecipients field if non-nil, zero value otherwise.

### GetDoNotValidateSubjectRecipientsOk

`func (o *IdentitySourcesSAMLConfigConfig) GetDoNotValidateSubjectRecipientsOk() (*bool, bool)`

GetDoNotValidateSubjectRecipientsOk returns a tuple with the DoNotValidateSubjectRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSubjectRecipients

`func (o *IdentitySourcesSAMLConfigConfig) SetDoNotValidateSubjectRecipients(v bool)`

SetDoNotValidateSubjectRecipients sets DoNotValidateSubjectRecipients field to given value.

### HasDoNotValidateSubjectRecipients

`func (o *IdentitySourcesSAMLConfigConfig) HasDoNotValidateSubjectRecipients() bool`

HasDoNotValidateSubjectRecipients returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *IdentitySourcesSAMLConfigConfig) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *IdentitySourcesSAMLConfigConfig) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *IdentitySourcesSAMLConfigConfig) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *IdentitySourcesSAMLConfigConfig) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetForceAuthn

`func (o *IdentitySourcesSAMLConfigConfig) GetForceAuthn() bool`

GetForceAuthn returns the ForceAuthn field if non-nil, zero value otherwise.

### GetForceAuthnOk

`func (o *IdentitySourcesSAMLConfigConfig) GetForceAuthnOk() (*bool, bool)`

GetForceAuthnOk returns a tuple with the ForceAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAuthn

`func (o *IdentitySourcesSAMLConfigConfig) SetForceAuthn(v bool)`

SetForceAuthn sets ForceAuthn field to given value.

### HasForceAuthn

`func (o *IdentitySourcesSAMLConfigConfig) HasForceAuthn() bool`

HasForceAuthn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


