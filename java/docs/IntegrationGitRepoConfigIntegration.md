

# IntegrationGitRepoConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | [**TypeEnum**](#TypeEnum) | Integration Type Code | 
**serviceUrl** | **String** | Git URL | 
**serviceUsername** | **String** | Username | 
**servicePassword** | **String** | Password |  [optional]
**serviceToken** | **String** | Access Token |  [optional]
**serviceKey** | **Long** | Key Pair ID |  [optional]
**config** | [**IntegrationGitRepoConfigIntegrationConfig**](IntegrationGitRepoConfigIntegrationConfig.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
GIT | &quot;git&quot;



