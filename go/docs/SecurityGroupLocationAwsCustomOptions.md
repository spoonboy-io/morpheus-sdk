# SecurityGroupLocationAwsCustomOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpc** | Pointer to **string** | External ID of Amazon VPC | [optional] 

## Methods

### NewSecurityGroupLocationAwsCustomOptions

`func NewSecurityGroupLocationAwsCustomOptions() *SecurityGroupLocationAwsCustomOptions`

NewSecurityGroupLocationAwsCustomOptions instantiates a new SecurityGroupLocationAwsCustomOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupLocationAwsCustomOptionsWithDefaults

`func NewSecurityGroupLocationAwsCustomOptionsWithDefaults() *SecurityGroupLocationAwsCustomOptions`

NewSecurityGroupLocationAwsCustomOptionsWithDefaults instantiates a new SecurityGroupLocationAwsCustomOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpc

`func (o *SecurityGroupLocationAwsCustomOptions) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *SecurityGroupLocationAwsCustomOptions) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *SecurityGroupLocationAwsCustomOptions) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *SecurityGroupLocationAwsCustomOptions) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


