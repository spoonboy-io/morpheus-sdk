

# ClusterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**OneOfstringobject**](OneOfstringobject.md) |  | 
**name** | **String** | Name of the cluster to be created | 
**description** | **String** | Description of the cluster to be created |  [optional]
**labels** | **List&lt;String&gt;** | Array of strings (keywords). This will override labels passed under the &#x60;server&#x60; object. |  [optional]
**group** | [**ClusterCreateGroup**](ClusterCreateGroup.md) |  | 
**cloud** | [**ClusterCreateCloud**](ClusterCreateCloud.md) |  | 
**layout** | [**ClusterCreateLayout**](ClusterCreateLayout.md) |  | 
**server** | [**ClusterServerCreate**](ClusterServerCreate.md) |  | 



