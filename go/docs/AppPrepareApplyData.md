# AppPrepareApplyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AutoValidate** | Pointer to **bool** |  | [optional] 
**Terraform** | Pointer to [**AppPrepareApplyDataTerraform**](appPrepareApply_data_terraform.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**BlueprintName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | Pointer to **int64** |  | [optional] 
**BlueprintId** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**AppPrepareApplyDataGroup**](appPrepareApply_data_group.md) |  | [optional] 

## Methods

### NewAppPrepareApplyData

`func NewAppPrepareApplyData() *AppPrepareApplyData`

NewAppPrepareApplyData instantiates a new AppPrepareApplyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPrepareApplyDataWithDefaults

`func NewAppPrepareApplyDataWithDefaults() *AppPrepareApplyData`

NewAppPrepareApplyDataWithDefaults instantiates a new AppPrepareApplyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *AppPrepareApplyData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AppPrepareApplyData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AppPrepareApplyData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AppPrepareApplyData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *AppPrepareApplyData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPrepareApplyData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPrepareApplyData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppPrepareApplyData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutoValidate

`func (o *AppPrepareApplyData) GetAutoValidate() bool`

GetAutoValidate returns the AutoValidate field if non-nil, zero value otherwise.

### GetAutoValidateOk

`func (o *AppPrepareApplyData) GetAutoValidateOk() (*bool, bool)`

GetAutoValidateOk returns a tuple with the AutoValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoValidate

`func (o *AppPrepareApplyData) SetAutoValidate(v bool)`

SetAutoValidate sets AutoValidate field to given value.

### HasAutoValidate

`func (o *AppPrepareApplyData) HasAutoValidate() bool`

HasAutoValidate returns a boolean if a field has been set.

### GetTerraform

`func (o *AppPrepareApplyData) GetTerraform() AppPrepareApplyDataTerraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *AppPrepareApplyData) GetTerraformOk() (*AppPrepareApplyDataTerraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *AppPrepareApplyData) SetTerraform(v AppPrepareApplyDataTerraform)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *AppPrepareApplyData) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.

### GetType

`func (o *AppPrepareApplyData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPrepareApplyData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPrepareApplyData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppPrepareApplyData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *AppPrepareApplyData) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AppPrepareApplyData) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AppPrepareApplyData) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AppPrepareApplyData) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBlueprintName

`func (o *AppPrepareApplyData) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *AppPrepareApplyData) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *AppPrepareApplyData) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *AppPrepareApplyData) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetDescription

`func (o *AppPrepareApplyData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPrepareApplyData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPrepareApplyData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPrepareApplyData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppPrepareApplyData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppPrepareApplyData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateId

`func (o *AppPrepareApplyData) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *AppPrepareApplyData) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *AppPrepareApplyData) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *AppPrepareApplyData) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetBlueprintId

`func (o *AppPrepareApplyData) GetBlueprintId() int64`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *AppPrepareApplyData) GetBlueprintIdOk() (*int64, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *AppPrepareApplyData) SetBlueprintId(v int64)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *AppPrepareApplyData) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetGroup

`func (o *AppPrepareApplyData) GetGroup() AppPrepareApplyDataGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AppPrepareApplyData) GetGroupOk() (*AppPrepareApplyDataGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AppPrepareApplyData) SetGroup(v AppPrepareApplyDataGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AppPrepareApplyData) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


