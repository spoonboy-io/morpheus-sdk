# ApiUsersIdUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Username** | Pointer to **string** | Username (unique per tenant). | [optional] 
**Email** | Pointer to **string** | Email address | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Roles** | Pointer to [**[]ReferenceObject**](ReferenceObject.md) | List of Roles | [optional] 

## Methods

### NewApiUsersIdUser

`func NewApiUsersIdUser() *ApiUsersIdUser`

NewApiUsersIdUser instantiates a new ApiUsersIdUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsersIdUserWithDefaults

`func NewApiUsersIdUserWithDefaults() *ApiUsersIdUser`

NewApiUsersIdUserWithDefaults instantiates a new ApiUsersIdUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ApiUsersIdUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApiUsersIdUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApiUsersIdUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApiUsersIdUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ApiUsersIdUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApiUsersIdUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApiUsersIdUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApiUsersIdUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *ApiUsersIdUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiUsersIdUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiUsersIdUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiUsersIdUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *ApiUsersIdUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiUsersIdUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiUsersIdUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiUsersIdUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *ApiUsersIdUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiUsersIdUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiUsersIdUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiUsersIdUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *ApiUsersIdUser) GetRoles() []ReferenceObject`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUsersIdUser) GetRolesOk() (*[]ReferenceObject, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUsersIdUser) SetRoles(v []ReferenceObject)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiUsersIdUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


