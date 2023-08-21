# # AppCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**template_id** | **int** |  | [optional]
**blueprint_id** | [**OneOfLongString**](OneOfLongString.md) | The ID of the Blueprint. Use \&quot;existing\&quot; to create a blank app. |
**name** | **string** | A unique name for the app |
**description** | **string** | Description | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**group** | [**\OpenAPI\Client\Model\AppCreateGroup**](AppCreateGroup.md) |  | [optional]
**default_cloud** | [**\OpenAPI\Client\Model\AppCreateDefaultCloud**](AppCreateDefaultCloud.md) |  | [optional]
**environment** | **string** | Environment code (appContext) | [optional]
**tiers** | **object** | Configuration of app elements | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
