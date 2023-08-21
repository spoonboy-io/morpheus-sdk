

# DeploymentVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**version** | **String** | Version number (userVersion), a unique version identifier for the deployment version. |  [optional]
**userVersion** | **String** | Alias for version |  [optional]
**deployType** | [**DeployTypeEnum**](#DeployTypeEnum) | Deploy Type, eg. file, git, fetch |  [optional]
**gitUrl** | **String** |  |  [optional]
**gitRef** | **String** |  |  [optional]
**fetchUrl** | **String** |  |  [optional]



## Enum: DeployTypeEnum

Name | Value
---- | -----
FILE | &quot;file&quot;
GIT | &quot;git&quot;
FETCH | &quot;fetch&quot;



