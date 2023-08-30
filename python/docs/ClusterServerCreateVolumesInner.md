# ClusterServerCreateVolumesInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**datastore_id** | **str, none_type** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). | 
**name** | **str** | Name/type of the LV being created | defaults to "root"
**id** | **int** | The id for the LV configuration being created | [optional]  if omitted the server will use the default value of -1
**root_volume** | **bool** | If set to false then a non-root LV will be created | [optional]  if omitted the server will use the default value of True
**size** | **int** | Size of the LV to be created in GBs  Default is from the service plan  | [optional] 
**size_id** | **str, none_type** | Can be used to select pre-existing LV choices from Morpheus | [optional] 
**storage_type** | **int** | Identifier for LV type | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


