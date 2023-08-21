# ApiAccountsAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Role** | Pointer to [**ApiAccountsAccountRole**](_api_accounts_account_role.md) |  | [optional] 
**Subdomain** | Pointer to **NullableString** | The subdomain. This will be part of the login URL and username for sub tenant users. | [optional] 
**Currency** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] [default to "USD"]

## Methods

### NewApiAccountsAccount

`func NewApiAccountsAccount(name string, ) *ApiAccountsAccount`

NewApiAccountsAccount instantiates a new ApiAccountsAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountsAccountWithDefaults

`func NewApiAccountsAccountWithDefaults() *ApiAccountsAccount`

NewApiAccountsAccountWithDefaults instantiates a new ApiAccountsAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiAccountsAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAccountsAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAccountsAccount) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiAccountsAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiAccountsAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiAccountsAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiAccountsAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiAccountsAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiAccountsAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRole

`func (o *ApiAccountsAccount) GetRole() ApiAccountsAccountRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiAccountsAccount) GetRoleOk() (*ApiAccountsAccountRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiAccountsAccount) SetRole(v ApiAccountsAccountRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApiAccountsAccount) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubdomain

`func (o *ApiAccountsAccount) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ApiAccountsAccount) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ApiAccountsAccount) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ApiAccountsAccount) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### SetSubdomainNil

`func (o *ApiAccountsAccount) SetSubdomainNil(b bool)`

 SetSubdomainNil sets the value for Subdomain to be an explicit nil

### UnsetSubdomain
`func (o *ApiAccountsAccount) UnsetSubdomain()`

UnsetSubdomain ensures that no value is present for Subdomain, not even an explicit nil
### GetCurrency

`func (o *ApiAccountsAccount) GetCurrency() CurrencyCode`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ApiAccountsAccount) GetCurrencyOk() (*CurrencyCode, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ApiAccountsAccount) SetCurrency(v CurrencyCode)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ApiAccountsAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


