# ApiServicePlansIdServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Service plan name | [optional] 
**Code** | Pointer to **string** | Service plan code, must be unique | [optional] 
**Description** | Pointer to **string** | Service plan description | [optional] 
**Editable** | Pointer to **bool** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**MaxStorage** | Pointer to **float32** | Max storage size in bytes | [optional] 
**MaxMemory** | Pointer to **float32** | Max memory size in bytes | [optional] 
**MaxCores** | Pointer to **float32** | Max cores | [optional] 
**MaxDisks** | Pointer to **float32** | Max disks allowed | [optional] 
**ProvisionType** | Pointer to [**[]ApiServicePlansIdServicePlanProvisionType**](ApiServicePlansIdServicePlanProvisionType.md) |  | [optional] 
**CustomCores** | Pointer to **bool** | Can be used to enable / disable customizable cores | [optional] [default to false]
**CustomMaxStorage** | Pointer to **bool** | Can be used to enable / disable customizable storage | [optional] [default to false]
**CustomMaxDataStorage** | Pointer to **bool** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**CustomMaxMemory** | Pointer to **bool** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**AddVolumes** | Pointer to **bool** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**SortOrder** | Pointer to **float32** | Sort order | [optional] 
**PriceSets** | Pointer to **[]int64** | List of price sets to include in service plan | [optional] 
**Config** | Pointer to [**ApiServicePlansIdServicePlanConfig**](_api_service_plans__id__servicePlan_config.md) |  | [optional] 

## Methods

### NewApiServicePlansIdServicePlan

`func NewApiServicePlansIdServicePlan() *ApiServicePlansIdServicePlan`

NewApiServicePlansIdServicePlan instantiates a new ApiServicePlansIdServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicePlansIdServicePlanWithDefaults

`func NewApiServicePlansIdServicePlanWithDefaults() *ApiServicePlansIdServicePlan`

NewApiServicePlansIdServicePlanWithDefaults instantiates a new ApiServicePlansIdServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiServicePlansIdServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiServicePlansIdServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiServicePlansIdServicePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiServicePlansIdServicePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ApiServicePlansIdServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiServicePlansIdServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiServicePlansIdServicePlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiServicePlansIdServicePlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ApiServicePlansIdServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiServicePlansIdServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiServicePlansIdServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiServicePlansIdServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditable

`func (o *ApiServicePlansIdServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ApiServicePlansIdServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ApiServicePlansIdServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ApiServicePlansIdServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ApiServicePlansIdServicePlan) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ApiServicePlansIdServicePlan) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ApiServicePlansIdServicePlan) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ApiServicePlansIdServicePlan) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ApiServicePlansIdServicePlan) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ApiServicePlansIdServicePlan) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ApiServicePlansIdServicePlan) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ApiServicePlansIdServicePlan) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCores

`func (o *ApiServicePlansIdServicePlan) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ApiServicePlansIdServicePlan) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ApiServicePlansIdServicePlan) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ApiServicePlansIdServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ApiServicePlansIdServicePlan) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ApiServicePlansIdServicePlan) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ApiServicePlansIdServicePlan) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ApiServicePlansIdServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetProvisionType

`func (o *ApiServicePlansIdServicePlan) GetProvisionType() []ApiServicePlansIdServicePlanProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ApiServicePlansIdServicePlan) GetProvisionTypeOk() (*[]ApiServicePlansIdServicePlanProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ApiServicePlansIdServicePlan) SetProvisionType(v []ApiServicePlansIdServicePlanProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ApiServicePlansIdServicePlan) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetCustomCores

`func (o *ApiServicePlansIdServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ApiServicePlansIdServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ApiServicePlansIdServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ApiServicePlansIdServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ApiServicePlansIdServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ApiServicePlansIdServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ApiServicePlansIdServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ApiServicePlansIdServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ApiServicePlansIdServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ApiServicePlansIdServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ApiServicePlansIdServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ApiServicePlansIdServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ApiServicePlansIdServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ApiServicePlansIdServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ApiServicePlansIdServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ApiServicePlansIdServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ApiServicePlansIdServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ApiServicePlansIdServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ApiServicePlansIdServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ApiServicePlansIdServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetSortOrder

`func (o *ApiServicePlansIdServicePlan) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ApiServicePlansIdServicePlan) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ApiServicePlansIdServicePlan) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ApiServicePlansIdServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetPriceSets

`func (o *ApiServicePlansIdServicePlan) GetPriceSets() []int64`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ApiServicePlansIdServicePlan) GetPriceSetsOk() (*[]int64, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ApiServicePlansIdServicePlan) SetPriceSets(v []int64)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ApiServicePlansIdServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *ApiServicePlansIdServicePlan) GetConfig() ApiServicePlansIdServicePlanConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiServicePlansIdServicePlan) GetConfigOk() (*ApiServicePlansIdServicePlanConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiServicePlansIdServicePlan) SetConfig(v ApiServicePlansIdServicePlanConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiServicePlansIdServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


