# ClusterCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**ClusterCreateType**](ClusterCreateType.md) |  | 
**name** | **str** | Name of the cluster to be created | 
**group** | [**ClusterCreateGroup**](ClusterCreateGroup.md) |  | 
**cloud** | [**ClusterCreateCloud**](ClusterCreateCloud.md) |  | 
**layout** | [**ClusterCreateLayout**](ClusterCreateLayout.md) |  | 
**server** | [**ClusterServerCreate**](ClusterServerCreate.md) |  | 
**description** | **str** | Description of the cluster to be created | [optional] 
**labels** | **[str]** | Array of strings (keywords). This will override labels passed under the &#x60;server&#x60; object. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


