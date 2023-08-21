# ApiSecurityPackagesSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the security package | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Security Package Type Code | [optional] [default to "scap-package"]
**Description** | Pointer to **string** | A description for the security package | [optional] 
**Url** | **string** | URL to download the security package zip file from | 
**Enabled** | Pointer to **bool** | Can be used to disable the security package | [optional] [default to true]

## Methods

### NewApiSecurityPackagesSecurityPackage

`func NewApiSecurityPackagesSecurityPackage(name string, url string, ) *ApiSecurityPackagesSecurityPackage`

NewApiSecurityPackagesSecurityPackage instantiates a new ApiSecurityPackagesSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityPackagesSecurityPackageWithDefaults

`func NewApiSecurityPackagesSecurityPackageWithDefaults() *ApiSecurityPackagesSecurityPackage`

NewApiSecurityPackagesSecurityPackageWithDefaults instantiates a new ApiSecurityPackagesSecurityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiSecurityPackagesSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSecurityPackagesSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSecurityPackagesSecurityPackage) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *ApiSecurityPackagesSecurityPackage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiSecurityPackagesSecurityPackage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiSecurityPackagesSecurityPackage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiSecurityPackagesSecurityPackage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ApiSecurityPackagesSecurityPackage) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ApiSecurityPackagesSecurityPackage) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *ApiSecurityPackagesSecurityPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiSecurityPackagesSecurityPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiSecurityPackagesSecurityPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiSecurityPackagesSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ApiSecurityPackagesSecurityPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiSecurityPackagesSecurityPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiSecurityPackagesSecurityPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiSecurityPackagesSecurityPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ApiSecurityPackagesSecurityPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiSecurityPackagesSecurityPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiSecurityPackagesSecurityPackage) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *ApiSecurityPackagesSecurityPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiSecurityPackagesSecurityPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiSecurityPackagesSecurityPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiSecurityPackagesSecurityPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


