# AddIntegrations200ResponseAllOfIntegration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Type** | **String** |  | [optional] 
**IntegrationType** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Url** | **String** |  | [optional] 
**ServiceKey** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**IsPlugin** | **Boolean** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Status** | **String** |  | [optional] 
**StatusDate** | **System.DateTime** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**LastSync** | **String** |  | [optional] 
**LastSyncDuration** | **String** |  | [optional] 
**Credential** | [**Creds**](Creds.md) |  | [optional] 
**Version** | **String** |  | [optional] 
**VarHost** | **String** |  | [optional] 
**Username** | **String** |  | [optional] 
**Password** | **String** |  | [optional] 
**PasswordHash** | **String** |  | [optional] 
**Token** | **String** |  | [optional] 
**TokenHash** | **String** |  | [optional] 
**ServiceFlag** | **Boolean** |  | [optional] 
**Port** | **String** |  | [optional] 
**Path** | **String** |  | [optional] 
**WindowsVersion** | **String** |  | [optional] 
**RepoUrl** | **String** |  | [optional] 
**ServiceMode** | **String** |  | [optional] 
**AuthType** | **String** |  | [optional] 
**AuthId** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddIntegrations200ResponseAllOfIntegration = Initialize-PSOpenAPIToolsAddIntegrations200ResponseAllOfIntegration  -Id null `
 -Name null `
 -Enabled null `
 -Type null `
 -IntegrationType null `
 -Url null `
 -ServiceKey null `
 -IsPlugin null `
 -Config null `
 -Status null `
 -StatusDate null `
 -StatusMessage null `
 -LastSync null `
 -LastSyncDuration null `
 -Credential null `
 -Version null `
 -VarHost null `
 -Username null `
 -Password null `
 -PasswordHash null `
 -Token null `
 -TokenHash null `
 -ServiceFlag null `
 -Port null `
 -Path null `
 -WindowsVersion null `
 -RepoUrl null `
 -ServiceMode null `
 -AuthType null `
 -AuthId null
```

- Convert the resource to JSON
```powershell
$AddIntegrations200ResponseAllOfIntegration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

