# SecurityScanSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 

## Methods

### NewSecurityScanSecurityPackage

`func NewSecurityScanSecurityPackage() *SecurityScanSecurityPackage`

NewSecurityScanSecurityPackage instantiates a new SecurityScanSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanSecurityPackageWithDefaults

`func NewSecurityScanSecurityPackageWithDefaults() *SecurityScanSecurityPackage`

NewSecurityScanSecurityPackageWithDefaults instantiates a new SecurityScanSecurityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityScanSecurityPackage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityScanSecurityPackage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityScanSecurityPackage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityScanSecurityPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityScanSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityScanSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityScanSecurityPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityScanSecurityPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SecurityScanSecurityPackage) GetType() InlineResponse20094Network`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityScanSecurityPackage) GetTypeOk() (*InlineResponse20094Network, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityScanSecurityPackage) SetType(v InlineResponse20094Network)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityScanSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


