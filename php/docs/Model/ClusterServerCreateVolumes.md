# # ClusterServerCreateVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | The id for the LV configuration being created | [optional] [default to -1]
**root_volume** | **bool** | If set to false then a non-root LV will be created | [optional] [default to true]
**name** | **string** | Name/type of the LV being created | [default to 'root']
**size** | **int** | Size of the LV to be created in GBs  Default is from the service plan | [optional]
**size_id** | **string** | Can be used to select pre-existing LV choices from Morpheus | [optional]
**storage_type** | **int** | Identifier for LV type | [optional]
**datastore_id** | **string** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
