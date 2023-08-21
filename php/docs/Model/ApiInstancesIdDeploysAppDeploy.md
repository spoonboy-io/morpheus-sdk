# # ApiInstancesIdDeploysAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_id** | **int** | Deployment ID. | [optional]
**version** | **string** | Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version | [optional]
**version_id** | **int** | Deployment Version ID. This can be passed instead of deploymentId and version. | [optional]
**config** | **object** | Map of configuration properties that vary by instance type. | [optional]
**stage_only** | **bool** | Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. | [optional] [default to false]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
