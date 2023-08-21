# ApiServicePlansIdServicePlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | Pointer to **string** | Specifies range min / max storage multiplier | [optional] [default to "gb"]
**MemorySizeType** | Pointer to **string** | Specifies range min / max memory multiplier | [optional] [default to "mb"]
**Ranges** | Pointer to [**ApiServicePlansIdServicePlanConfigRanges**](_api_service_plans__id__servicePlan_config_ranges.md) |  | [optional] 

## Methods

### NewApiServicePlansIdServicePlanConfig

`func NewApiServicePlansIdServicePlanConfig() *ApiServicePlansIdServicePlanConfig`

NewApiServicePlansIdServicePlanConfig instantiates a new ApiServicePlansIdServicePlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicePlansIdServicePlanConfigWithDefaults

`func NewApiServicePlansIdServicePlanConfigWithDefaults() *ApiServicePlansIdServicePlanConfig`

NewApiServicePlansIdServicePlanConfigWithDefaults instantiates a new ApiServicePlansIdServicePlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSizeType

`func (o *ApiServicePlansIdServicePlanConfig) GetStorageSizeType() string`

GetStorageSizeType returns the StorageSizeType field if non-nil, zero value otherwise.

### GetStorageSizeTypeOk

`func (o *ApiServicePlansIdServicePlanConfig) GetStorageSizeTypeOk() (*string, bool)`

GetStorageSizeTypeOk returns a tuple with the StorageSizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeType

`func (o *ApiServicePlansIdServicePlanConfig) SetStorageSizeType(v string)`

SetStorageSizeType sets StorageSizeType field to given value.

### HasStorageSizeType

`func (o *ApiServicePlansIdServicePlanConfig) HasStorageSizeType() bool`

HasStorageSizeType returns a boolean if a field has been set.

### GetMemorySizeType

`func (o *ApiServicePlansIdServicePlanConfig) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *ApiServicePlansIdServicePlanConfig) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *ApiServicePlansIdServicePlanConfig) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *ApiServicePlansIdServicePlanConfig) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.

### GetRanges

`func (o *ApiServicePlansIdServicePlanConfig) GetRanges() ApiServicePlansIdServicePlanConfigRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *ApiServicePlansIdServicePlanConfig) GetRangesOk() (*ApiServicePlansIdServicePlanConfigRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *ApiServicePlansIdServicePlanConfig) SetRanges(v ApiServicePlansIdServicePlanConfigRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *ApiServicePlansIdServicePlanConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


