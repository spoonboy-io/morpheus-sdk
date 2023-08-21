# VdiPoolConfigStorageControllers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 
**UnitNumber** | Pointer to **string** |  | [optional] 
**BusNumber** | Pointer to **string** |  | [optional] 
**MaxDevices** | Pointer to **float32** |  | [optional] 
**Removable** | Pointer to **NullableBool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ReservedUnitNumber** | Pointer to **float32** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **float32** |  | [optional] 

## Methods

### NewVdiPoolConfigStorageControllers

`func NewVdiPoolConfigStorageControllers() *VdiPoolConfigStorageControllers`

NewVdiPoolConfigStorageControllers instantiates a new VdiPoolConfigStorageControllers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiPoolConfigStorageControllersWithDefaults

`func NewVdiPoolConfigStorageControllersWithDefaults() *VdiPoolConfigStorageControllers`

NewVdiPoolConfigStorageControllersWithDefaults instantiates a new VdiPoolConfigStorageControllers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VdiPoolConfigStorageControllers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VdiPoolConfigStorageControllers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VdiPoolConfigStorageControllers) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VdiPoolConfigStorageControllers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VdiPoolConfigStorageControllers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdiPoolConfigStorageControllers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdiPoolConfigStorageControllers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VdiPoolConfigStorageControllers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *VdiPoolConfigStorageControllers) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VdiPoolConfigStorageControllers) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VdiPoolConfigStorageControllers) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *VdiPoolConfigStorageControllers) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *VdiPoolConfigStorageControllers) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *VdiPoolConfigStorageControllers) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetTypeId

`func (o *VdiPoolConfigStorageControllers) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VdiPoolConfigStorageControllers) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VdiPoolConfigStorageControllers) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *VdiPoolConfigStorageControllers) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetTypeName

`func (o *VdiPoolConfigStorageControllers) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *VdiPoolConfigStorageControllers) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *VdiPoolConfigStorageControllers) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *VdiPoolConfigStorageControllers) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetUnitNumber

`func (o *VdiPoolConfigStorageControllers) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *VdiPoolConfigStorageControllers) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *VdiPoolConfigStorageControllers) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *VdiPoolConfigStorageControllers) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetBusNumber

`func (o *VdiPoolConfigStorageControllers) GetBusNumber() string`

GetBusNumber returns the BusNumber field if non-nil, zero value otherwise.

### GetBusNumberOk

`func (o *VdiPoolConfigStorageControllers) GetBusNumberOk() (*string, bool)`

GetBusNumberOk returns a tuple with the BusNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusNumber

`func (o *VdiPoolConfigStorageControllers) SetBusNumber(v string)`

SetBusNumber sets BusNumber field to given value.

### HasBusNumber

`func (o *VdiPoolConfigStorageControllers) HasBusNumber() bool`

HasBusNumber returns a boolean if a field has been set.

### GetMaxDevices

`func (o *VdiPoolConfigStorageControllers) GetMaxDevices() float32`

GetMaxDevices returns the MaxDevices field if non-nil, zero value otherwise.

### GetMaxDevicesOk

`func (o *VdiPoolConfigStorageControllers) GetMaxDevicesOk() (*float32, bool)`

GetMaxDevicesOk returns a tuple with the MaxDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDevices

`func (o *VdiPoolConfigStorageControllers) SetMaxDevices(v float32)`

SetMaxDevices sets MaxDevices field to given value.

### HasMaxDevices

`func (o *VdiPoolConfigStorageControllers) HasMaxDevices() bool`

HasMaxDevices returns a boolean if a field has been set.

### GetRemovable

`func (o *VdiPoolConfigStorageControllers) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *VdiPoolConfigStorageControllers) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *VdiPoolConfigStorageControllers) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.

### HasRemovable

`func (o *VdiPoolConfigStorageControllers) HasRemovable() bool`

HasRemovable returns a boolean if a field has been set.

### SetRemovableNil

`func (o *VdiPoolConfigStorageControllers) SetRemovableNil(b bool)`

 SetRemovableNil sets the value for Removable to be an explicit nil

### UnsetRemovable
`func (o *VdiPoolConfigStorageControllers) UnsetRemovable()`

UnsetRemovable ensures that no value is present for Removable, not even an explicit nil
### GetEditable

`func (o *VdiPoolConfigStorageControllers) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *VdiPoolConfigStorageControllers) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *VdiPoolConfigStorageControllers) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *VdiPoolConfigStorageControllers) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetReservedUnitNumber

`func (o *VdiPoolConfigStorageControllers) GetReservedUnitNumber() float32`

GetReservedUnitNumber returns the ReservedUnitNumber field if non-nil, zero value otherwise.

### GetReservedUnitNumberOk

`func (o *VdiPoolConfigStorageControllers) GetReservedUnitNumberOk() (*float32, bool)`

GetReservedUnitNumberOk returns a tuple with the ReservedUnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUnitNumber

`func (o *VdiPoolConfigStorageControllers) SetReservedUnitNumber(v float32)`

SetReservedUnitNumber sets ReservedUnitNumber field to given value.

### HasReservedUnitNumber

`func (o *VdiPoolConfigStorageControllers) HasReservedUnitNumber() bool`

HasReservedUnitNumber returns a boolean if a field has been set.

### GetCategory

`func (o *VdiPoolConfigStorageControllers) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VdiPoolConfigStorageControllers) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VdiPoolConfigStorageControllers) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *VdiPoolConfigStorageControllers) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *VdiPoolConfigStorageControllers) GetDisplayOrder() float32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *VdiPoolConfigStorageControllers) GetDisplayOrderOk() (*float32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *VdiPoolConfigStorageControllers) SetDisplayOrder(v float32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *VdiPoolConfigStorageControllers) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


