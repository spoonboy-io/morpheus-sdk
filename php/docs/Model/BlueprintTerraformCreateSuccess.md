# # BlueprintTerraformCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the blueprint | [optional]
**image** | **string** | Path to display image. Defaults to an internal Morpheus image. | [optional]
**type** | **string** | Blueprint Type | [optional]
**terraform** | [**\OpenAPI\Client\Model\BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | [optional]
**config** | [**\OpenAPI\Client\Model\BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional]
**visibility** | **string** | Private or Public Access | [optional] [default to 'private']
**resource_permission** | **object** | Resource Permission Block | [optional]
**owner** | **object** | Owner | [optional]
**tenant** | **object** | Tenant | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
