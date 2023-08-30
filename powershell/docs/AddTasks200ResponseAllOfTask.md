# AddTasks200ResponseAllOfTask
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**TaskType** | [**TaskWriteAttributesConfigTaskType**](TaskWriteAttributesConfigTaskType.md) |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**TaskOptions** | [**TaskWriteAttributesConfigTaskOptions**](TaskWriteAttributesConfigTaskOptions.md) |  | [optional] 
**File** | [**TaskAnsiblePlaybookConfigFile**](TaskAnsiblePlaybookConfigFile.md) |  | [optional] 
**ResultType** | **String** |  | [optional] 
**ExecuteTarget** | **String** |  | [optional] 
**Retryable** | **Boolean** |  | [optional] 
**RetryCount** | **Int64** |  | [optional] 
**RetryDelaySeconds** | **Int64** |  | [optional] 
**AllowCustomConfig** | **Boolean** |  | [optional] 
**Credential** | [**ZoneCredentialAnyOf**](ZoneCredentialAnyOf.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddTasks200ResponseAllOfTask = Initialize-PSOpenAPIToolsAddTasks200ResponseAllOfTask  -Id null `
 -AccountId null `
 -Name null `
 -Code null `
 -TaskType null `
 -Labels null `
 -Visibility null `
 -TaskOptions null `
 -File null `
 -ResultType null `
 -ExecuteTarget null `
 -Retryable null `
 -RetryCount null `
 -RetryDelaySeconds null `
 -AllowCustomConfig null `
 -Credential null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$AddTasks200ResponseAllOfTask | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

