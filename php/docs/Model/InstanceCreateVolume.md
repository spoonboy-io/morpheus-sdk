# # InstanceCreateVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | The id for the LV configuration being created. | [optional] [default to -1]
**root_volume** | **bool** | If set to false then a non-root LV will be created. | [optional] [default to true]
**name** | **string** | Name/type of the LV being created. | [optional] [default to 'root']
**size** | **int** | Size of the LV to be created in GBs.  Uses default from service plan. | [optional]
**size_id** | **int** | Can be used to select pre-existing LV choices from Morpheus. | [optional]
**storage_type** | **int** | Identifier for LV type | [optional]
**datastore_id** | [**AnyOfStringLong**](AnyOfStringLong.md) | The ID of the specific datastore. Auto selection can be specified as auto or autoCluster (for clusters). | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
