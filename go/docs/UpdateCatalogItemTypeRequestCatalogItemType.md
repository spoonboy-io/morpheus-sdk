# UpdateCatalogItemTypeRequestCatalogItemType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Catalog Item Type name | [optional] 
**Code** | Pointer to **NullableString** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | Pointer to **NullableString** | Catalog Item Type category | [optional] 
**Description** | Pointer to **string** | Catalog Item Type description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**LayoutCode** | Pointer to **NullableString** | Identifier primarily used for Plugin Catalog Item Types | [optional] 
**IconPath** | Pointer to **string** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the catalog item type. | [optional] [default to true]
**Featured** | Pointer to **bool** | Can be used to feature the catalog item type. | [optional] [default to false]
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] [default to false]
**Config** | Pointer to [**Object**](Object.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs, see Inputs. Only applies to type instance and blueprint. | [optional] 
**Blueprint** | Pointer to [**CatalogItemTypeBlueprintCreateBlueprint**](CatalogItemTypeBlueprintCreateBlueprint.md) |  | [optional] 
**AppSpec** | Pointer to **string** | The appSpec for blueprint type catalog items is a string in the Scribe YAML format with fields | [optional] 
**Workflow** | Pointer to [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**WorkflowConfig** | Pointer to **string** | Configuration object that contains settings for the workflow. | [optional] 

## Methods

### NewUpdateCatalogItemTypeRequestCatalogItemType

`func NewUpdateCatalogItemTypeRequestCatalogItemType() *UpdateCatalogItemTypeRequestCatalogItemType`

NewUpdateCatalogItemTypeRequestCatalogItemType instantiates a new UpdateCatalogItemTypeRequestCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogItemTypeRequestCatalogItemTypeWithDefaults

`func NewUpdateCatalogItemTypeRequestCatalogItemTypeWithDefaults() *UpdateCatalogItemTypeRequestCatalogItemType`

NewUpdateCatalogItemTypeRequestCatalogItemTypeWithDefaults instantiates a new UpdateCatalogItemTypeRequestCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateCatalogItemTypeRequestCatalogItemType) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateCatalogItemTypeRequestCatalogItemType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateCatalogItemTypeRequestCatalogItemType) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### SetLayoutCodeNil

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetLayoutCodeNil(b bool)`

 SetLayoutCodeNil sets the value for LayoutCode to be an explicit nil

### UnsetLayoutCode
`func (o *UpdateCatalogItemTypeRequestCatalogItemType) UnsetLayoutCode()`

UnsetLayoutCode ensures that no value is present for LayoutCode, not even an explicit nil
### GetIconPath

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetConfig() Object`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetConfigOk() (*Object, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetConfig(v Object)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOptionTypes

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetBlueprint

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetBlueprint() CatalogItemTypeBlueprintCreateBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetBlueprintOk() (*CatalogItemTypeBlueprintCreateBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetBlueprint(v CatalogItemTypeBlueprintCreateBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetAppSpec

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### GetWorkflow

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetWorkflow() UpdateBlueprintPermissionsRequestResourcePermissionSitesInner`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetWorkflowOk() (*UpdateBlueprintPermissionsRequestResourcePermissionSitesInner, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetWorkflow(v UpdateBlueprintPermissionsRequestResourcePermissionSitesInner)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetContext

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetWorkflowConfig

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetWorkflowConfig() string`

GetWorkflowConfig returns the WorkflowConfig field if non-nil, zero value otherwise.

### GetWorkflowConfigOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) GetWorkflowConfigOk() (*string, bool)`

GetWorkflowConfigOk returns a tuple with the WorkflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowConfig

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) SetWorkflowConfig(v string)`

SetWorkflowConfig sets WorkflowConfig field to given value.

### HasWorkflowConfig

`func (o *UpdateCatalogItemTypeRequestCatalogItemType) HasWorkflowConfig() bool`

HasWorkflowConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


