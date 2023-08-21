

# ClusterServerCreateVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** | The id for the LV configuration being created |  [optional]
**rootVolume** | **Boolean** | If set to false then a non-root LV will be created |  [optional]
**name** | **String** | Name/type of the LV being created | 
**size** | **Long** | Size of the LV to be created in GBs  Default is from the service plan  |  [optional]
**sizeId** | **String** | Can be used to select pre-existing LV choices from Morpheus |  [optional]
**storageType** | **Long** | Identifier for LV type |  [optional]
**datastoreId** | **String** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). | 



