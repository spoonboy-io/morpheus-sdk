# ApiServicePlansServicePlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | Pointer to **string** | Specifies range min / max storage multiplier | [optional] [default to "gb"]
**MemorySizeType** | Pointer to **string** | Specifies range min / max memory multiplier | [optional] [default to "mb"]
**Ranges** | Pointer to [**ApiServicePlansServicePlanConfigRanges**](_api_service_plans_servicePlan_config_ranges.md) |  | [optional] 

## Methods

### NewApiServicePlansServicePlanConfig

`func NewApiServicePlansServicePlanConfig() *ApiServicePlansServicePlanConfig`

NewApiServicePlansServicePlanConfig instantiates a new ApiServicePlansServicePlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicePlansServicePlanConfigWithDefaults

`func NewApiServicePlansServicePlanConfigWithDefaults() *ApiServicePlansServicePlanConfig`

NewApiServicePlansServicePlanConfigWithDefaults instantiates a new ApiServicePlansServicePlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSizeType

`func (o *ApiServicePlansServicePlanConfig) GetStorageSizeType() string`

GetStorageSizeType returns the StorageSizeType field if non-nil, zero value otherwise.

### GetStorageSizeTypeOk

`func (o *ApiServicePlansServicePlanConfig) GetStorageSizeTypeOk() (*string, bool)`

GetStorageSizeTypeOk returns a tuple with the StorageSizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeType

`func (o *ApiServicePlansServicePlanConfig) SetStorageSizeType(v string)`

SetStorageSizeType sets StorageSizeType field to given value.

### HasStorageSizeType

`func (o *ApiServicePlansServicePlanConfig) HasStorageSizeType() bool`

HasStorageSizeType returns a boolean if a field has been set.

### GetMemorySizeType

`func (o *ApiServicePlansServicePlanConfig) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *ApiServicePlansServicePlanConfig) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *ApiServicePlansServicePlanConfig) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *ApiServicePlansServicePlanConfig) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.

### GetRanges

`func (o *ApiServicePlansServicePlanConfig) GetRanges() ApiServicePlansServicePlanConfigRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *ApiServicePlansServicePlanConfig) GetRangesOk() (*ApiServicePlansServicePlanConfigRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *ApiServicePlansServicePlanConfig) SetRanges(v ApiServicePlansServicePlanConfigRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *ApiServicePlansServicePlanConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


