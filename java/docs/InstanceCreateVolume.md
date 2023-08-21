

# InstanceCreateVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** | The id for the LV configuration being created. |  [optional]
**rootVolume** | **Boolean** | If set to false then a non-root LV will be created. |  [optional]
**name** | **String** | Name/type of the LV being created. |  [optional]
**size** | **Long** | Size of the LV to be created in GBs.  Uses default from service plan. |  [optional]
**sizeId** | **Long** | Can be used to select pre-existing LV choices from Morpheus. |  [optional]
**storageType** | **Long** | Identifier for LV type |  [optional]
**datastoreId** | [**AnyOfstringlong**](AnyOfstringlong.md) | The ID of the specific datastore. Auto selection can be specified as auto or autoCluster (for clusters). |  [optional]



