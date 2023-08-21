# # ClusterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**OneOfStringObject**](OneOfStringObject.md) |  |
**name** | **string** | Name of the cluster to be created |
**description** | **string** | Description of the cluster to be created | [optional]
**labels** | **string[]** | Array of strings (keywords). This will override labels passed under the &#x60;server&#x60; object. | [optional]
**group** | [**\OpenAPI\Client\Model\ClusterCreateGroup**](ClusterCreateGroup.md) |  |
**cloud** | [**\OpenAPI\Client\Model\ClusterCreateCloud**](ClusterCreateCloud.md) |  |
**layout** | [**\OpenAPI\Client\Model\ClusterCreateLayout**](ClusterCreateLayout.md) |  |
**server** | [**\OpenAPI\Client\Model\ClusterServerCreate**](ClusterServerCreate.md) |  |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
