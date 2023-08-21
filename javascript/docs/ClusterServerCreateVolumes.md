# MorpheusApi.ClusterServerCreateVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** | The id for the LV configuration being created | [optional] [default to -1]
**rootVolume** | **Boolean** | If set to false then a non-root LV will be created | [optional] [default to true]
**name** | **String** | Name/type of the LV being created | [default to &#39;root&#39;]
**size** | **Number** | Size of the LV to be created in GBs  Default is from the service plan  | [optional] 
**sizeId** | **String** | Can be used to select pre-existing LV choices from Morpheus | [optional] 
**storageType** | **Number** | Identifier for LV type | [optional] 
**datastoreId** | **String** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). | 


