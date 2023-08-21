# # ImageBuildCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the image build | [optional]
**description** | **string** | A description for the image build | [optional]
**type** | **string** | The image builder type. | [optional]
**site** | [**\OpenAPI\Client\Model\ImageBuildCreateSite**](ImageBuildCreateSite.md) |  | [optional]
**zone** | [**\OpenAPI\Client\Model\ImageBuildCreateZone**](ImageBuildCreateZone.md) |  | [optional]
**config** | **object** | A map of config values. This is the instance config that is used for provisioning. See Provisioning Types. | [optional]
**boot_script** | [**\OpenAPI\Client\Model\ImageBuildCreateBootScript**](ImageBuildCreateBootScript.md) |  | [optional]
**preseed_script** | [**\OpenAPI\Client\Model\ImageBuildCreatePreseedScript**](ImageBuildCreatePreseedScript.md) |  | [optional]
**ssh_username** | **string** | SSH Username | [optional]
**ssh_password** | **string** | SSH Password | [optional]
**storage_provider** | **string** | Storage Provider | [optional]
**is_cloud_init** | **string** | Cloud Init | [optional]
**build_output_name** | **string** | Build Output Name | [optional]
**conversion_formats** | **string** | Conversion Formats | [optional]
**keep_results** | **int** | Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality. | [optional] [default to 0]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
