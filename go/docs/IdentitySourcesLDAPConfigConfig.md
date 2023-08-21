# IdentitySourcesLDAPConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**BindingUsername** | Pointer to **string** |  | [optional] 
**BindingPassword** | Pointer to **string** |  | [optional] 
**UserFqnExpression** | Pointer to **string** |  | [optional] 
**RequiredRoleFqn** | Pointer to **string** |  | [optional] 
**UsernameAttribute** | Pointer to **string** |  | [optional] 
**CommonNameAttribute** | Pointer to **string** |  | [optional] 
**FirstNameAttribute** | Pointer to **string** |  | [optional] 
**LastNameAttribute** | Pointer to **string** |  | [optional] 
**EmailAttribute** | Pointer to **string** |  | [optional] 
**UniqueMemberAttribute** | Pointer to **string** |  | [optional] 
**MemberOfAttribute** | Pointer to **string** |  | [optional] 
**BindingPasswordHash** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentitySourcesLDAPConfigConfig

`func NewIdentitySourcesLDAPConfigConfig() *IdentitySourcesLDAPConfigConfig`

NewIdentitySourcesLDAPConfigConfig instantiates a new IdentitySourcesLDAPConfigConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourcesLDAPConfigConfigWithDefaults

`func NewIdentitySourcesLDAPConfigConfigWithDefaults() *IdentitySourcesLDAPConfigConfig`

NewIdentitySourcesLDAPConfigConfigWithDefaults instantiates a new IdentitySourcesLDAPConfigConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *IdentitySourcesLDAPConfigConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IdentitySourcesLDAPConfigConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IdentitySourcesLDAPConfigConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IdentitySourcesLDAPConfigConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBindingUsername

`func (o *IdentitySourcesLDAPConfigConfig) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *IdentitySourcesLDAPConfigConfig) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *IdentitySourcesLDAPConfigConfig) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *IdentitySourcesLDAPConfigConfig) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *IdentitySourcesLDAPConfigConfig) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *IdentitySourcesLDAPConfigConfig) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *IdentitySourcesLDAPConfigConfig) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *IdentitySourcesLDAPConfigConfig) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetUserFqnExpression

`func (o *IdentitySourcesLDAPConfigConfig) GetUserFqnExpression() string`

GetUserFqnExpression returns the UserFqnExpression field if non-nil, zero value otherwise.

### GetUserFqnExpressionOk

`func (o *IdentitySourcesLDAPConfigConfig) GetUserFqnExpressionOk() (*string, bool)`

GetUserFqnExpressionOk returns a tuple with the UserFqnExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFqnExpression

`func (o *IdentitySourcesLDAPConfigConfig) SetUserFqnExpression(v string)`

SetUserFqnExpression sets UserFqnExpression field to given value.

### HasUserFqnExpression

`func (o *IdentitySourcesLDAPConfigConfig) HasUserFqnExpression() bool`

HasUserFqnExpression returns a boolean if a field has been set.

### GetRequiredRoleFqn

`func (o *IdentitySourcesLDAPConfigConfig) GetRequiredRoleFqn() string`

GetRequiredRoleFqn returns the RequiredRoleFqn field if non-nil, zero value otherwise.

### GetRequiredRoleFqnOk

`func (o *IdentitySourcesLDAPConfigConfig) GetRequiredRoleFqnOk() (*string, bool)`

GetRequiredRoleFqnOk returns a tuple with the RequiredRoleFqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRoleFqn

`func (o *IdentitySourcesLDAPConfigConfig) SetRequiredRoleFqn(v string)`

SetRequiredRoleFqn sets RequiredRoleFqn field to given value.

### HasRequiredRoleFqn

`func (o *IdentitySourcesLDAPConfigConfig) HasRequiredRoleFqn() bool`

HasRequiredRoleFqn returns a boolean if a field has been set.

### GetUsernameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetUsernameAttribute() string`

GetUsernameAttribute returns the UsernameAttribute field if non-nil, zero value otherwise.

### GetUsernameAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetUsernameAttributeOk() (*string, bool)`

GetUsernameAttributeOk returns a tuple with the UsernameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetUsernameAttribute(v string)`

SetUsernameAttribute sets UsernameAttribute field to given value.

### HasUsernameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasUsernameAttribute() bool`

HasUsernameAttribute returns a boolean if a field has been set.

### GetCommonNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetCommonNameAttribute() string`

GetCommonNameAttribute returns the CommonNameAttribute field if non-nil, zero value otherwise.

### GetCommonNameAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetCommonNameAttributeOk() (*string, bool)`

GetCommonNameAttributeOk returns a tuple with the CommonNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetCommonNameAttribute(v string)`

SetCommonNameAttribute sets CommonNameAttribute field to given value.

### HasCommonNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasCommonNameAttribute() bool`

HasCommonNameAttribute returns a boolean if a field has been set.

### GetFirstNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetFirstNameAttribute() string`

GetFirstNameAttribute returns the FirstNameAttribute field if non-nil, zero value otherwise.

### GetFirstNameAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetFirstNameAttributeOk() (*string, bool)`

GetFirstNameAttributeOk returns a tuple with the FirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetFirstNameAttribute(v string)`

SetFirstNameAttribute sets FirstNameAttribute field to given value.

### HasFirstNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasFirstNameAttribute() bool`

HasFirstNameAttribute returns a boolean if a field has been set.

### GetLastNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetLastNameAttribute() string`

GetLastNameAttribute returns the LastNameAttribute field if non-nil, zero value otherwise.

### GetLastNameAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetLastNameAttributeOk() (*string, bool)`

GetLastNameAttributeOk returns a tuple with the LastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetLastNameAttribute(v string)`

SetLastNameAttribute sets LastNameAttribute field to given value.

### HasLastNameAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasLastNameAttribute() bool`

HasLastNameAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetUniqueMemberAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetUniqueMemberAttribute() string`

GetUniqueMemberAttribute returns the UniqueMemberAttribute field if non-nil, zero value otherwise.

### GetUniqueMemberAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetUniqueMemberAttributeOk() (*string, bool)`

GetUniqueMemberAttributeOk returns a tuple with the UniqueMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueMemberAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetUniqueMemberAttribute(v string)`

SetUniqueMemberAttribute sets UniqueMemberAttribute field to given value.

### HasUniqueMemberAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasUniqueMemberAttribute() bool`

HasUniqueMemberAttribute returns a boolean if a field has been set.

### GetMemberOfAttribute

`func (o *IdentitySourcesLDAPConfigConfig) GetMemberOfAttribute() string`

GetMemberOfAttribute returns the MemberOfAttribute field if non-nil, zero value otherwise.

### GetMemberOfAttributeOk

`func (o *IdentitySourcesLDAPConfigConfig) GetMemberOfAttributeOk() (*string, bool)`

GetMemberOfAttributeOk returns a tuple with the MemberOfAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfAttribute

`func (o *IdentitySourcesLDAPConfigConfig) SetMemberOfAttribute(v string)`

SetMemberOfAttribute sets MemberOfAttribute field to given value.

### HasMemberOfAttribute

`func (o *IdentitySourcesLDAPConfigConfig) HasMemberOfAttribute() bool`

HasMemberOfAttribute returns a boolean if a field has been set.

### GetBindingPasswordHash

`func (o *IdentitySourcesLDAPConfigConfig) GetBindingPasswordHash() string`

GetBindingPasswordHash returns the BindingPasswordHash field if non-nil, zero value otherwise.

### GetBindingPasswordHashOk

`func (o *IdentitySourcesLDAPConfigConfig) GetBindingPasswordHashOk() (*string, bool)`

GetBindingPasswordHashOk returns a tuple with the BindingPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPasswordHash

`func (o *IdentitySourcesLDAPConfigConfig) SetBindingPasswordHash(v string)`

SetBindingPasswordHash sets BindingPasswordHash field to given value.

### HasBindingPasswordHash

`func (o *IdentitySourcesLDAPConfigConfig) HasBindingPasswordHash() bool`

HasBindingPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


