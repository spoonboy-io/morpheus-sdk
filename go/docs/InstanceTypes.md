# InstanceTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisionTypeCode** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 
**InstanceTypeLayouts** | Pointer to [**[]InstanceTypesInstanceTypeLayouts**](InstanceTypesInstanceTypeLayouts.md) |  | [optional] 
**ImagePath** | Pointer to **NullableString** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **NullableString** | Dark logo image URL | [optional] 

## Methods

### NewInstanceTypes

`func NewInstanceTypes() *InstanceTypes`

NewInstanceTypes instantiates a new InstanceTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypesWithDefaults

`func NewInstanceTypesWithDefaults() *InstanceTypes`

NewInstanceTypesWithDefaults instantiates a new InstanceTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *InstanceTypes) GetAccount() InlineResponse20082LoadBalancerInstanceSslCert`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InstanceTypes) GetAccountOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InstanceTypes) SetAccount(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InstanceTypes) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *InstanceTypes) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *InstanceTypes) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetName

`func (o *InstanceTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceTypes) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceTypes) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceTypes) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceTypes) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *InstanceTypes) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *InstanceTypes) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetCode

`func (o *InstanceTypes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceTypes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceTypes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceTypes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceTypes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceTypes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceTypes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisionTypeCode

`func (o *InstanceTypes) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *InstanceTypes) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *InstanceTypes) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *InstanceTypes) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### SetProvisionTypeCodeNil

`func (o *InstanceTypes) SetProvisionTypeCodeNil(b bool)`

 SetProvisionTypeCodeNil sets the value for ProvisionTypeCode to be an explicit nil

### UnsetProvisionTypeCode
`func (o *InstanceTypes) UnsetProvisionTypeCode()`

UnsetProvisionTypeCode ensures that no value is present for ProvisionTypeCode, not even an explicit nil
### GetCategory

`func (o *InstanceTypes) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceTypes) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceTypes) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceTypes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetActive

`func (o *InstanceTypes) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InstanceTypes) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InstanceTypes) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InstanceTypes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *InstanceTypes) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *InstanceTypes) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *InstanceTypes) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *InstanceTypes) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetVisibility

`func (o *InstanceTypes) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InstanceTypes) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InstanceTypes) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InstanceTypes) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetFeatured

`func (o *InstanceTypes) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *InstanceTypes) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *InstanceTypes) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *InstanceTypes) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVersions

`func (o *InstanceTypes) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *InstanceTypes) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *InstanceTypes) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *InstanceTypes) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetInstanceTypeLayouts

`func (o *InstanceTypes) GetInstanceTypeLayouts() []InstanceTypesInstanceTypeLayouts`

GetInstanceTypeLayouts returns the InstanceTypeLayouts field if non-nil, zero value otherwise.

### GetInstanceTypeLayoutsOk

`func (o *InstanceTypes) GetInstanceTypeLayoutsOk() (*[]InstanceTypesInstanceTypeLayouts, bool)`

GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeLayouts

`func (o *InstanceTypes) SetInstanceTypeLayouts(v []InstanceTypesInstanceTypeLayouts)`

SetInstanceTypeLayouts sets InstanceTypeLayouts field to given value.

### HasInstanceTypeLayouts

`func (o *InstanceTypes) HasInstanceTypeLayouts() bool`

HasInstanceTypeLayouts returns a boolean if a field has been set.

### GetImagePath

`func (o *InstanceTypes) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *InstanceTypes) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *InstanceTypes) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *InstanceTypes) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *InstanceTypes) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *InstanceTypes) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetDarkImagePath

`func (o *InstanceTypes) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *InstanceTypes) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *InstanceTypes) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *InstanceTypes) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### SetDarkImagePathNil

`func (o *InstanceTypes) SetDarkImagePathNil(b bool)`

 SetDarkImagePathNil sets the value for DarkImagePath to be an explicit nil

### UnsetDarkImagePath
`func (o *InstanceTypes) UnsetDarkImagePath()`

UnsetDarkImagePath ensures that no value is present for DarkImagePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


