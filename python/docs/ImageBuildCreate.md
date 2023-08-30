# ImageBuildCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the image build | [optional] 
**description** | **str, none_type** | A description for the image build | [optional] 
**type** | **str** | The image builder type. | [optional]  if omitted the server will use the default value of "vmware"
**site** | [**ImageBuildCreateSite**](ImageBuildCreateSite.md) |  | [optional] 
**zone** | [**ImageBuildCreateZone**](ImageBuildCreateZone.md) |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | A map of config values. This is the instance config that is used for provisioning. See Provisioning Types. | [optional] 
**boot_script** | [**ImageBuildCreateBootScript**](ImageBuildCreateBootScript.md) |  | [optional] 
**preseed_script** | [**ImageBuildCreatePreseedScript**](ImageBuildCreatePreseedScript.md) |  | [optional] 
**ssh_username** | **str** | SSH Username | [optional] 
**ssh_password** | **str** | SSH Password | [optional] 
**storage_provider** | **str, none_type** | Storage Provider | [optional] 
**is_cloud_init** | **str** | Cloud Init | [optional] 
**build_output_name** | **str, none_type** | Build Output Name | [optional] 
**conversion_formats** | **str, none_type** | Conversion Formats | [optional] 
**keep_results** | **int** | Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality. | [optional]  if omitted the server will use the default value of 0
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


