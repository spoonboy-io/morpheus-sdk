# MorpheusApi.IntegrationAnsibleConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | **String** | Integration Type Code | 
**enabled** | **Boolean** | Set &#x60;true&#x60; to enable integration | [optional] 
**refresh** | **Boolean** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  | [optional] [default to true]
**serviceUrl** | **String** | Ansible Git URL | 
**serviceUsername** | **String** | Git Username | [optional] 
**servicePassword** | **String** | Git Password or Token depending on the Git host | [optional] 
**serviceToken** | **String** | Git Token | [optional] 
**serviceKey** | **Number** | Keypair ID | [optional] 
**config** | [**IntegrationAnsibleConfigIntegrationConfig**](IntegrationAnsibleConfigIntegrationConfig.md) |  | [optional] 



## Enum: TypeEnum


* `ansible` (value: `"ansible"`)




