

# IntegrationSaltMaster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**name** | **String** |  |  [optional]
**enabled** | **Boolean** |  |  [optional]
**type** | [**TypeEnum**](#TypeEnum) |  |  [optional]
**integrationType** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  |  [optional]
**url** | **String** |  |  [optional]
**port** | **String** |  |  [optional]
**username** | **String** |  |  [optional]
**password** | **String** |  |  [optional]
**passwordHash** | **String** |  |  [optional]
**path** | **String** |  |  [optional]
**version** | **String** |  |  [optional]
**windowsVersion** | **String** |  |  [optional]
**repoUrl** | **String** |  |  [optional]
**serviceMode** | **String** |  |  [optional]
**isPlugin** | **Boolean** |  |  [optional]
**config** | [**IntegrationSaltMasterConfig**](IntegrationSaltMasterConfig.md) |  |  [optional]
**status** | **String** |  |  [optional]
**statusDate** | **OffsetDateTime** |  |  [optional]
**statusMessage** | **String** |  |  [optional]
**lastSync** | **String** |  |  [optional]
**lastSyncDuration** | **String** |  |  [optional]
**credential** | [**Creds**](Creds.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
SALTMASTER | &quot;saltMaster&quot;



