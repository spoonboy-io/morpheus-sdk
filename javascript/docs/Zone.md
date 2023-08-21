# MorpheusApi.Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** |  | [optional] 
**uuid** | **String** |  | [optional] 
**externalId** | **String** |  | [optional] 
**name** | **String** |  | [optional] 
**code** | **String** |  | [optional] 
**location** | **String** |  | [optional] 
**owner** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**accountId** | **Number** |  | [optional] 
**account** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**visibility** | **String** |  | [optional] 
**enabled** | **Boolean** |  | [optional] 
**status** | **String** |  | [optional] 
**statusMessage** | **String** |  | [optional] 
**statusDate** | **Date** |  | [optional] 
**costStatus** | **String** |  | [optional] 
**costStatusMessage** | **String** |  | [optional] 
**costStatusDate** | **Date** |  | [optional] 
**costLastSyncDuration** | **Number** |  | [optional] 
**costLastSync** | **Date** |  | [optional] 
**zoneType** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  | [optional] 
**zoneTypeId** | **Number** |  | [optional] 
**guidanceMode** | **String** |  | [optional] 
**storageMode** | **String** |  | [optional] 
**agentMode** | **String** |  | [optional] 
**userDataLinux** | **String** |  | [optional] 
**userDataWindows** | **String** |  | [optional] 
**consoleKeymap** | **String** |  | [optional] 
**containerMode** | **String** |  | [optional] 
**costingMode** | **String** |  | [optional] 
**serviceVersion** | **String** |  | [optional] 
**securityMode** | **String** |  | [optional] 
**inventoryLevel** | **String** |  | [optional] 
**timezone** | **String** |  | [optional] 
**apiProxy** | **String** |  | [optional] 
**provisioningProxy** | **String** |  | [optional] 
**networkDomain** | [**InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional] 
**domainName** | **String** |  | [optional] 
**regionCode** | **String** |  | [optional] 
**autoRecoverPowerState** | **Boolean** |  | [optional] 
**scalePriority** | **Number** |  | [optional] 
**config** | [**AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig**](AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig.md) |  | [optional] 
**credential** | [**AnyOfobjectobject**](AnyOfobjectobject.md) |  | [optional] 
**imagePath** | **String** | Logo image URL | [optional] 
**darkImagePath** | **String** | Dark logo image URL | [optional] 
**dateCreated** | **Date** |  | [optional] 
**lastUpdated** | **Date** |  | [optional] 
**lastSync** | **Date** |  | [optional] 
**lastSyncDuration** | **Number** |  | [optional] 
**nextRunDate** | **Date** |  | [optional] 
**groups** | [**[ZoneGroups]**](ZoneGroups.md) |  | [optional] 
**securityServer** | [**InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional] 
**networkServer** | [**InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional] 
**stats** | [**ZoneStats**](ZoneStats.md) |  | [optional] 
**serverCount** | **Number** |  | [optional] 

