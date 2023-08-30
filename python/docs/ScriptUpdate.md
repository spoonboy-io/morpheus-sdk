# ScriptUpdate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Script name | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**category** | **str** | Script category | [optional] 
**script_version** | **str** | Version of the script | [optional]  if omitted the server will use the default value of "1"
**script_phase** | **str** | Phase for the script, provision, start, etc. | [optional] 
**script_type** | **str** | Type for the script | [optional]  if omitted the server will use the default value of "bash"
**script** | **str** | Script content, that is, the code itself. | [optional] 
**run_as_user** | **str** | Run as a specific user. | [optional] 
**sudo_user** | **bool** | Sudo, whether or not to run with sudo. | [optional]  if omitted the server will use the default value of False
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


