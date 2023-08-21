# MorpheusApi.IntegrationSaltMasterConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | **String** | Integration Type Code | 
**serviceMode** | **String** | Topology | [optional] [default to &#39;single&#39;]
**serviceUrl** | **String** | Salt Master | 
**secondary** | **String** | Salt Syndic | [optional] 
**servicePort** | **Number** | SSH Port | [optional] [default to 22]
**serviceUsername** | **String** | Username | 
**servicePassword** | **String** | Password | [optional] 
**serviceKey** | **String** | Master Key Pair | [optional] 
**authKey** | **String** | Signing Key | [optional] 
**servicePath** | **String** | Working Directory | [optional] 
**serviceVersion** | **String** | Salt Version | [optional] 
**serviceWindowsVersion** | **String** | Salt Version (Windows) | [optional] 
**repoUrl** | **String** | Repo URL | [optional] 
**serviceConfig** | **String** | Minion Config | [optional] 
**serviceCommand** | **String** | Post Provision Commands | [optional] 
**config** | [**IntegrationSaltMasterConfigIntegrationConfig**](IntegrationSaltMasterConfigIntegrationConfig.md) |  | [optional] 



## Enum: TypeEnum


* `saltMaster` (value: `"saltMaster"`)




