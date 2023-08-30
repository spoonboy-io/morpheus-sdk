# InstanceCreateVolume


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | The id for the LV configuration being created. | [optional]  if omitted the server will use the default value of -1
**root_volume** | **bool** | If set to false then a non-root LV will be created. | [optional]  if omitted the server will use the default value of True
**name** | **str** | Name/type of the LV being created. | [optional]  if omitted the server will use the default value of "root"
**size** | **int** | Size of the LV to be created in GBs.  Uses default from service plan. | [optional] 
**size_id** | **int, none_type** | Can be used to select pre-existing LV choices from Morpheus. | [optional] 
**storage_type** | **int, none_type** | Identifier for LV type | [optional] 
**datastore_id** | [**InstanceCreateVolumeDatastoreId**](InstanceCreateVolumeDatastoreId.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


