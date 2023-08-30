# FileTemplate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Code** | **String** |  | [optional] 
**Account** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**FileName** | **String** |  | [optional] 
**FilePath** | **String** |  | [optional] 
**TemplateType** | **String** |  | [optional] 
**TemplatePhase** | **String** |  | [optional] 
**Template** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**SettingCategory** | **String** |  | [optional] 
**SettingName** | **String** |  | [optional] 
**AutoRun** | **Boolean** |  | [optional] 
**RunOnScale** | **Boolean** |  | [optional] 
**RunOnDeploy** | **Boolean** |  | [optional] 
**FileOwner** | **String** |  | [optional] 
**FileGroup** | **String** |  | [optional] 
**Permissions** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$FileTemplate = Initialize-PSOpenAPIToolsFileTemplate  -Id null `
 -Code null `
 -Account null `
 -Name null `
 -Labels null `
 -FileName null `
 -FilePath null `
 -TemplateType null `
 -TemplatePhase null `
 -Template null `
 -Category null `
 -SettingCategory null `
 -SettingName null `
 -AutoRun null `
 -RunOnScale null `
 -RunOnDeploy null `
 -FileOwner null `
 -FileGroup null `
 -Permissions null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$FileTemplate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

