

# IntegrationGitHub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**name** | **String** |  |  [optional]
**enabled** | **Boolean** |  |  [optional]
**type** | [**TypeEnum**](#TypeEnum) |  |  [optional]
**integrationType** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  |  [optional]
**username** | **String** |  |  [optional]
**token** | **String** |  |  [optional]
**tokenHash** | **String** |  |  [optional]
**serviceKey** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  |  [optional]
**isPlugin** | **Boolean** |  |  [optional]
**config** | [**IntegrationGitRepoConfig**](IntegrationGitRepoConfig.md) |  |  [optional]
**status** | **String** |  |  [optional]
**statusDate** | **OffsetDateTime** |  |  [optional]
**statusMessage** | **String** |  |  [optional]
**lastSync** | **String** |  |  [optional]
**lastSyncDuration** | **String** |  |  [optional]
**credential** | [**Creds**](Creds.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
GITHUB | &quot;github&quot;



