# ApiSecurityPackagesIdSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the security package | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Security Package Type Code | [optional] [default to "scap-package"]
**Description** | Pointer to **string** | A description for the security package | [optional] 
**Url** | Pointer to **string** | URL to download the security package zip file from | [optional] 
**Enabled** | Pointer to **bool** | Can be used to disable the security package | [optional] 

## Methods

### NewApiSecurityPackagesIdSecurityPackage

`func NewApiSecurityPackagesIdSecurityPackage() *ApiSecurityPackagesIdSecurityPackage`

NewApiSecurityPackagesIdSecurityPackage instantiates a new ApiSecurityPackagesIdSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityPackagesIdSecurityPackageWithDefaults

`func NewApiSecurityPackagesIdSecurityPackageWithDefaults() *ApiSecurityPackagesIdSecurityPackage`

NewApiSecurityPackagesIdSecurityPackageWithDefaults instantiates a new ApiSecurityPackagesIdSecurityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiSecurityPackagesIdSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSecurityPackagesIdSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSecurityPackagesIdSecurityPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiSecurityPackagesIdSecurityPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ApiSecurityPackagesIdSecurityPackage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiSecurityPackagesIdSecurityPackage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiSecurityPackagesIdSecurityPackage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiSecurityPackagesIdSecurityPackage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ApiSecurityPackagesIdSecurityPackage) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ApiSecurityPackagesIdSecurityPackage) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *ApiSecurityPackagesIdSecurityPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiSecurityPackagesIdSecurityPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiSecurityPackagesIdSecurityPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiSecurityPackagesIdSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ApiSecurityPackagesIdSecurityPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiSecurityPackagesIdSecurityPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiSecurityPackagesIdSecurityPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiSecurityPackagesIdSecurityPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ApiSecurityPackagesIdSecurityPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiSecurityPackagesIdSecurityPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiSecurityPackagesIdSecurityPackage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiSecurityPackagesIdSecurityPackage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiSecurityPackagesIdSecurityPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiSecurityPackagesIdSecurityPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiSecurityPackagesIdSecurityPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiSecurityPackagesIdSecurityPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


