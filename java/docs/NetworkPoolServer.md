

# NetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** | Network Pool Server ID |  [optional]
**type** | [**NetworkPoolServerType**](NetworkPoolServerType.md) |  |  [optional]
**name** | **String** | Name |  [optional]
**enabled** | **Boolean** |  |  [optional]
**serviceUrl** | **String** | Service URL |  [optional]
**serviceHost** | **String** | Service Host |  [optional]
**servicePort** | **Integer** | Service Port |  [optional]
**serviceMode** | **String** | Service Mode |  [optional]
**serviceUsername** | **String** | Service Username |  [optional]
**servicePassword** | **String** | Service Password |  [optional]
**servicePasswordHash** | **String** |  |  [optional]
**serviceThrottleRate** | **Long** | Throttle Rate |  [optional]
**ignoreSsl** | **Boolean** | Disable SSL SNI Verification |  [optional]
**status** | **String** | Status |  [optional]
**statusMessage** | **String** | Status Message |  [optional]
**statusDate** | **OffsetDateTime** |  |  [optional]
**config** | **Object** | Config object varies with pool server type. |  [optional]
**networkFilter** | **String** | Network Filter |  [optional]
**zoneFilter** | **String** | Zone Filter |  [optional]
**tenantMatch** | **String** | Tenant Match |  [optional]
**dateCreated** | **OffsetDateTime** |  |  [optional]
**lastUpdated** | **OffsetDateTime** |  |  [optional]
**account** | [**NetworkPoolServerAccount**](NetworkPoolServerAccount.md) |  |  [optional]
**integration** | [**NetworkPoolServerIntegration**](NetworkPoolServerIntegration.md) |  |  [optional]
**pools** | [**List&lt;InlineResponse20040AppDeployInstance&gt;**](InlineResponse20040AppDeployInstance.md) |  |  [optional]
**credential** | [**Creds2**](Creds2.md) |  |  [optional]



