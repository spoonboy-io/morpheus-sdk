

# ApiInstancesIdDeploysAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deploymentId** | **Long** | Deployment ID. |  [optional]
**version** | **String** | Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version |  [optional]
**versionId** | **Long** | Deployment Version ID. This can be passed instead of deploymentId and version. |  [optional]
**config** | **Object** | Map of configuration properties that vary by instance type. |  [optional]
**stageOnly** | **Boolean** | Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. |  [optional]



