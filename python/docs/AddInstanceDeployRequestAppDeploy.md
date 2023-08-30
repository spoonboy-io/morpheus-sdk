# AddInstanceDeployRequestAppDeploy


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_id** | **int** | Deployment ID. | [optional] 
**version** | **str** | Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version | [optional] 
**version_id** | **int** | Deployment Version ID. This can be passed instead of deploymentId and version. | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Map of configuration properties that vary by instance type. | [optional] 
**stage_only** | **bool** | Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. | [optional]  if omitted the server will use the default value of False
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


