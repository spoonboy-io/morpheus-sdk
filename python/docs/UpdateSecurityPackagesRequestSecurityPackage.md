# UpdateSecurityPackagesRequestSecurityPackage


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the security package | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**type** | **str** | Security Package Type Code | [optional]  if omitted the server will use the default value of "scap-package"
**description** | **str** | A description for the security package | [optional] 
**url** | **str** | URL to download the security package zip file from | [optional] 
**enabled** | **bool** | Can be used to disable the security package | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


