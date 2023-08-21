# MorpheusApi.NetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** | Network Pool Server ID | [optional] 
**type** | [**NetworkPoolServerType**](NetworkPoolServerType.md) |  | [optional] 
**name** | **String** | Name | [optional] 
**enabled** | **Boolean** |  | [optional] 
**serviceUrl** | **String** | Service URL | [optional] 
**serviceHost** | **String** | Service Host | [optional] 
**servicePort** | **Number** | Service Port | [optional] 
**serviceMode** | **String** | Service Mode | [optional] 
**serviceUsername** | **String** | Service Username | [optional] 
**servicePassword** | **String** | Service Password | [optional] 
**servicePasswordHash** | **String** |  | [optional] 
**serviceThrottleRate** | **Number** | Throttle Rate | [optional] [default to 0]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification | [optional] [default to true]
**status** | **String** | Status | [optional] 
**statusMessage** | **String** | Status Message | [optional] 
**statusDate** | **Date** |  | [optional] 
**config** | **Object** | Config object varies with pool server type. | [optional] 
**networkFilter** | **String** | Network Filter | [optional] 
**zoneFilter** | **String** | Zone Filter | [optional] 
**tenantMatch** | **String** | Tenant Match | [optional] 
**dateCreated** | **Date** |  | [optional] 
**lastUpdated** | **Date** |  | [optional] 
**account** | [**NetworkPoolServerAccount**](NetworkPoolServerAccount.md) |  | [optional] 
**integration** | [**NetworkPoolServerIntegration**](NetworkPoolServerIntegration.md) |  | [optional] 
**pools** | [**[InlineResponse20040AppDeployInstance]**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**credential** | [**Creds2**](Creds2.md) |  | [optional] 


