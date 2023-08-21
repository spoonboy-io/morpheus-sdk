

# IntegrationAnsibleConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | [**TypeEnum**](#TypeEnum) | Integration Type Code | 
**enabled** | **Boolean** | Set &#x60;true&#x60; to enable integration |  [optional]
**refresh** | **Boolean** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  |  [optional]
**serviceUrl** | **String** | Ansible Git URL | 
**serviceUsername** | **String** | Git Username |  [optional]
**servicePassword** | **String** | Git Password or Token depending on the Git host |  [optional]
**serviceToken** | **String** | Git Token |  [optional]
**serviceKey** | **Long** | Keypair ID |  [optional]
**config** | [**IntegrationAnsibleConfigIntegrationConfig**](IntegrationAnsibleConfigIntegrationConfig.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
ANSIBLE | &quot;ansible&quot;



