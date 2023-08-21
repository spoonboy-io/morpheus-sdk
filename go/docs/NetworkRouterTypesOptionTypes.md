# NetworkRouterTypesOptionTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**FieldLabel** | Pointer to **string** |  | [optional] 
**FieldContext** | Pointer to **string** |  | [optional] 
**FieldGroup** | Pointer to **NullableString** |  | [optional] 
**FieldClass** | Pointer to **NullableString** |  | [optional] 
**FieldAddOn** | Pointer to **NullableString** |  | [optional] 
**FieldComponent** | Pointer to **NullableString** |  | [optional] 
**FieldInput** | Pointer to **string** |  | [optional] 
**PlaceHolder** | Pointer to **NullableString** |  | [optional] 
**HelpBlock** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**OptionSource** | Pointer to **string** |  | [optional] 
**OptionList** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Advanced** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**WrapperClass** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**NoBlank** | Pointer to **bool** |  | [optional] 
**DependsOnCode** | Pointer to **string** |  | [optional] 
**ContextualDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkRouterTypesOptionTypes

`func NewNetworkRouterTypesOptionTypes() *NetworkRouterTypesOptionTypes`

NewNetworkRouterTypesOptionTypes instantiates a new NetworkRouterTypesOptionTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterTypesOptionTypesWithDefaults

`func NewNetworkRouterTypesOptionTypesWithDefaults() *NetworkRouterTypesOptionTypes`

NewNetworkRouterTypesOptionTypesWithDefaults instantiates a new NetworkRouterTypesOptionTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkRouterTypesOptionTypes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkRouterTypesOptionTypes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkRouterTypesOptionTypes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkRouterTypesOptionTypes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkRouterTypesOptionTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkRouterTypesOptionTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkRouterTypesOptionTypes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkRouterTypesOptionTypes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkRouterTypesOptionTypes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkRouterTypesOptionTypes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkRouterTypesOptionTypes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkRouterTypesOptionTypes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkRouterTypesOptionTypes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkRouterTypesOptionTypes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *NetworkRouterTypesOptionTypes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkRouterTypesOptionTypes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkRouterTypesOptionTypes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkRouterTypesOptionTypes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *NetworkRouterTypesOptionTypes) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *NetworkRouterTypesOptionTypes) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *NetworkRouterTypesOptionTypes) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *NetworkRouterTypesOptionTypes) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *NetworkRouterTypesOptionTypes) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *NetworkRouterTypesOptionTypes) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldContext

`func (o *NetworkRouterTypesOptionTypes) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *NetworkRouterTypesOptionTypes) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *NetworkRouterTypesOptionTypes) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *NetworkRouterTypesOptionTypes) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *NetworkRouterTypesOptionTypes) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *NetworkRouterTypesOptionTypes) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *NetworkRouterTypesOptionTypes) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *NetworkRouterTypesOptionTypes) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *NetworkRouterTypesOptionTypes) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *NetworkRouterTypesOptionTypes) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *NetworkRouterTypesOptionTypes) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *NetworkRouterTypesOptionTypes) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *NetworkRouterTypesOptionTypes) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *NetworkRouterTypesOptionTypes) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *NetworkRouterTypesOptionTypes) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *NetworkRouterTypesOptionTypes) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *NetworkRouterTypesOptionTypes) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *NetworkRouterTypesOptionTypes) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *NetworkRouterTypesOptionTypes) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *NetworkRouterTypesOptionTypes) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *NetworkRouterTypesOptionTypes) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *NetworkRouterTypesOptionTypes) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *NetworkRouterTypesOptionTypes) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *NetworkRouterTypesOptionTypes) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *NetworkRouterTypesOptionTypes) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *NetworkRouterTypesOptionTypes) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *NetworkRouterTypesOptionTypes) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### GetPlaceHolder

`func (o *NetworkRouterTypesOptionTypes) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *NetworkRouterTypesOptionTypes) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *NetworkRouterTypesOptionTypes) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *NetworkRouterTypesOptionTypes) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *NetworkRouterTypesOptionTypes) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *NetworkRouterTypesOptionTypes) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetHelpBlock

`func (o *NetworkRouterTypesOptionTypes) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *NetworkRouterTypesOptionTypes) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *NetworkRouterTypesOptionTypes) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *NetworkRouterTypesOptionTypes) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### GetDefaultValue

`func (o *NetworkRouterTypesOptionTypes) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NetworkRouterTypesOptionTypes) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NetworkRouterTypesOptionTypes) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NetworkRouterTypesOptionTypes) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NetworkRouterTypesOptionTypes) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NetworkRouterTypesOptionTypes) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *NetworkRouterTypesOptionTypes) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *NetworkRouterTypesOptionTypes) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *NetworkRouterTypesOptionTypes) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *NetworkRouterTypesOptionTypes) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### GetOptionList

`func (o *NetworkRouterTypesOptionTypes) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *NetworkRouterTypesOptionTypes) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *NetworkRouterTypesOptionTypes) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *NetworkRouterTypesOptionTypes) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### SetOptionListNil

`func (o *NetworkRouterTypesOptionTypes) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *NetworkRouterTypesOptionTypes) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetType

`func (o *NetworkRouterTypesOptionTypes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkRouterTypesOptionTypes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkRouterTypesOptionTypes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkRouterTypesOptionTypes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *NetworkRouterTypesOptionTypes) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *NetworkRouterTypesOptionTypes) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *NetworkRouterTypesOptionTypes) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *NetworkRouterTypesOptionTypes) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *NetworkRouterTypesOptionTypes) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NetworkRouterTypesOptionTypes) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NetworkRouterTypesOptionTypes) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *NetworkRouterTypesOptionTypes) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetEditable

`func (o *NetworkRouterTypesOptionTypes) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NetworkRouterTypesOptionTypes) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NetworkRouterTypesOptionTypes) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *NetworkRouterTypesOptionTypes) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *NetworkRouterTypesOptionTypes) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *NetworkRouterTypesOptionTypes) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *NetworkRouterTypesOptionTypes) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *NetworkRouterTypesOptionTypes) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkRouterTypesOptionTypes) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkRouterTypesOptionTypes) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkRouterTypesOptionTypes) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkRouterTypesOptionTypes) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *NetworkRouterTypesOptionTypes) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *NetworkRouterTypesOptionTypes) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *NetworkRouterTypesOptionTypes) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *NetworkRouterTypesOptionTypes) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *NetworkRouterTypesOptionTypes) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *NetworkRouterTypesOptionTypes) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *NetworkRouterTypesOptionTypes) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *NetworkRouterTypesOptionTypes) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *NetworkRouterTypesOptionTypes) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *NetworkRouterTypesOptionTypes) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *NetworkRouterTypesOptionTypes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkRouterTypesOptionTypes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkRouterTypesOptionTypes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkRouterTypesOptionTypes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *NetworkRouterTypesOptionTypes) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *NetworkRouterTypesOptionTypes) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *NetworkRouterTypesOptionTypes) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *NetworkRouterTypesOptionTypes) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *NetworkRouterTypesOptionTypes) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *NetworkRouterTypesOptionTypes) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *NetworkRouterTypesOptionTypes) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *NetworkRouterTypesOptionTypes) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### GetContextualDefault

`func (o *NetworkRouterTypesOptionTypes) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *NetworkRouterTypesOptionTypes) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *NetworkRouterTypesOptionTypes) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *NetworkRouterTypesOptionTypes) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


