# Credential
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Integration** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Username** | **String** |  | [optional] 
**Password** | **String** |  | [optional] 
**PasswordHash** | **String** |  | [optional] 
**AuthKey** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**AuthPath** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Scope** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**StatusDate** | **System.DateTime** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**User** | [**CredentialUser**](CredentialUser.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Config** | [**CredentialConfig**](CredentialConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Credential = Initialize-PSOpenAPIToolsCredential  -Id null `
 -Name null `
 -Type null `
 -Integration null `
 -Description null `
 -Username null `
 -Password null `
 -PasswordHash null `
 -AuthKey null `
 -AuthPath null `
 -ExternalId null `
 -RefType null `
 -RefId null `
 -Category null `
 -Scope null `
 -Status null `
 -StatusMessage null `
 -StatusDate null `
 -Enabled null `
 -Account null `
 -User null `
 -DateCreated null `
 -LastUpdated null `
 -Config null
```

- Convert the resource to JSON
```powershell
$Credential | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

