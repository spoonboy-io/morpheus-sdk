# ApiServicePlansServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Service plan name | 
**Code** | **string** | Service plan code, must be unique | 
**Description** | Pointer to **string** | Service plan description | [optional] 
**Editable** | Pointer to **bool** | Can be used to enable / disable the editability of the service plan. | [optional] [default to true]
**MaxStorage** | **float32** | Max storage size in bytes | 
**MaxMemory** | **float32** | Max memory size in bytes | 
**MaxCores** | Pointer to **float32** | Max cores | [optional] 
**MaxDisks** | Pointer to **float32** | Max disks allowed | [optional] 
**ProvisionType** | [**[]ApiServicePlansServicePlanProvisionType**](ApiServicePlansServicePlanProvisionType.md) |  | 
**CustomCores** | Pointer to **bool** | Can be used to enable / disable customizable cores | [optional] [default to false]
**CustomMaxStorage** | Pointer to **bool** | Can be used to enable / disable customizable storage | [optional] [default to false]
**CustomMaxDataStorage** | Pointer to **bool** | Can be used to enable / disable customizable extra volumes. | [optional] [default to false]
**CustomMaxMemory** | Pointer to **bool** | Can be used to enable / disable customizable memory. | [optional] [default to false]
**AddVolumes** | Pointer to **bool** | Can be used to enable / disable ability to add volumes | [optional] [default to false]
**SortOrder** | Pointer to **float32** | Sort order | [optional] 
**PriceSets** | Pointer to **[]int64** | List of price sets to include in service plan | [optional] 
**Config** | Pointer to [**ApiServicePlansServicePlanConfig**](_api_service_plans_servicePlan_config.md) |  | [optional] 

## Methods

### NewApiServicePlansServicePlan

`func NewApiServicePlansServicePlan(name string, code string, maxStorage float32, maxMemory float32, provisionType []ApiServicePlansServicePlanProvisionType, ) *ApiServicePlansServicePlan`

NewApiServicePlansServicePlan instantiates a new ApiServicePlansServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicePlansServicePlanWithDefaults

`func NewApiServicePlansServicePlanWithDefaults() *ApiServicePlansServicePlan`

NewApiServicePlansServicePlanWithDefaults instantiates a new ApiServicePlansServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiServicePlansServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiServicePlansServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiServicePlansServicePlan) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ApiServicePlansServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiServicePlansServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiServicePlansServicePlan) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *ApiServicePlansServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiServicePlansServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiServicePlansServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiServicePlansServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditable

`func (o *ApiServicePlansServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ApiServicePlansServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ApiServicePlansServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ApiServicePlansServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ApiServicePlansServicePlan) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ApiServicePlansServicePlan) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ApiServicePlansServicePlan) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxMemory

`func (o *ApiServicePlansServicePlan) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ApiServicePlansServicePlan) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ApiServicePlansServicePlan) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxCores

`func (o *ApiServicePlansServicePlan) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ApiServicePlansServicePlan) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ApiServicePlansServicePlan) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ApiServicePlansServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ApiServicePlansServicePlan) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ApiServicePlansServicePlan) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ApiServicePlansServicePlan) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ApiServicePlansServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetProvisionType

`func (o *ApiServicePlansServicePlan) GetProvisionType() []ApiServicePlansServicePlanProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ApiServicePlansServicePlan) GetProvisionTypeOk() (*[]ApiServicePlansServicePlanProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ApiServicePlansServicePlan) SetProvisionType(v []ApiServicePlansServicePlanProvisionType)`

SetProvisionType sets ProvisionType field to given value.


### GetCustomCores

`func (o *ApiServicePlansServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ApiServicePlansServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ApiServicePlansServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ApiServicePlansServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ApiServicePlansServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ApiServicePlansServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ApiServicePlansServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ApiServicePlansServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ApiServicePlansServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ApiServicePlansServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ApiServicePlansServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ApiServicePlansServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ApiServicePlansServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ApiServicePlansServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ApiServicePlansServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ApiServicePlansServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ApiServicePlansServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ApiServicePlansServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ApiServicePlansServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ApiServicePlansServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetSortOrder

`func (o *ApiServicePlansServicePlan) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ApiServicePlansServicePlan) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ApiServicePlansServicePlan) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ApiServicePlansServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetPriceSets

`func (o *ApiServicePlansServicePlan) GetPriceSets() []int64`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ApiServicePlansServicePlan) GetPriceSetsOk() (*[]int64, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ApiServicePlansServicePlan) SetPriceSets(v []int64)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ApiServicePlansServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *ApiServicePlansServicePlan) GetConfig() ApiServicePlansServicePlanConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiServicePlansServicePlan) GetConfigOk() (*ApiServicePlansServicePlanConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiServicePlansServicePlan) SetConfig(v ApiServicePlansServicePlanConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiServicePlansServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


