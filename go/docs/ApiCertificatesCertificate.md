# ApiCertificatesCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name scoped to your account for the key | [optional] 
**CertFile** | Pointer to **string** | The contents of the certificate file | [optional] 
**KeyFile** | Pointer to **string** | The contents of the key file | [optional] 
**DomainName** | Pointer to **string** | The domain name this certificate is tied to | [optional] 
**Wildcard** | Pointer to **bool** | Wether or not this certificate is a wildcard cert | [optional] [default to false]

## Methods

### NewApiCertificatesCertificate

`func NewApiCertificatesCertificate() *ApiCertificatesCertificate`

NewApiCertificatesCertificate instantiates a new ApiCertificatesCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCertificatesCertificateWithDefaults

`func NewApiCertificatesCertificateWithDefaults() *ApiCertificatesCertificate`

NewApiCertificatesCertificateWithDefaults instantiates a new ApiCertificatesCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiCertificatesCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiCertificatesCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiCertificatesCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiCertificatesCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertFile

`func (o *ApiCertificatesCertificate) GetCertFile() string`

GetCertFile returns the CertFile field if non-nil, zero value otherwise.

### GetCertFileOk

`func (o *ApiCertificatesCertificate) GetCertFileOk() (*string, bool)`

GetCertFileOk returns a tuple with the CertFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFile

`func (o *ApiCertificatesCertificate) SetCertFile(v string)`

SetCertFile sets CertFile field to given value.

### HasCertFile

`func (o *ApiCertificatesCertificate) HasCertFile() bool`

HasCertFile returns a boolean if a field has been set.

### GetKeyFile

`func (o *ApiCertificatesCertificate) GetKeyFile() string`

GetKeyFile returns the KeyFile field if non-nil, zero value otherwise.

### GetKeyFileOk

`func (o *ApiCertificatesCertificate) GetKeyFileOk() (*string, bool)`

GetKeyFileOk returns a tuple with the KeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFile

`func (o *ApiCertificatesCertificate) SetKeyFile(v string)`

SetKeyFile sets KeyFile field to given value.

### HasKeyFile

`func (o *ApiCertificatesCertificate) HasKeyFile() bool`

HasKeyFile returns a boolean if a field has been set.

### GetDomainName

`func (o *ApiCertificatesCertificate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ApiCertificatesCertificate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ApiCertificatesCertificate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ApiCertificatesCertificate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetWildcard

`func (o *ApiCertificatesCertificate) GetWildcard() bool`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *ApiCertificatesCertificate) GetWildcardOk() (*bool, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *ApiCertificatesCertificate) SetWildcard(v bool)`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *ApiCertificatesCertificate) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


