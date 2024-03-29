# ClusterHistoryEventsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**ProcessId** | **Int64** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**UniqueId** | **String** |  | [optional] 
**ProcessType** | [**ClusterContainersAvailableActionsInner**](ClusterContainersAvailableActionsInner.md) |  | [optional] 
**Description** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**SubType** | **String** |  | [optional] 
**SubId** | **Int64** |  | [optional] 
**ZoneId** | **Int64** |  | [optional] 
**IntegrationId** | **Int64** |  | [optional] 
**InstanceId** | **Int64** |  | [optional] 
**ContainerId** | **Int64** |  | [optional] 
**ServerId** | **Int64** |  | [optional] 
**ContainerName** | **String** |  | [optional] 
**DisplayName** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**Reason** | **String** |  | [optional] 
**Percent** | **Int64** |  | [optional] 
**StatusEta** | **Int64** |  | [optional] 
**Message** | **String** |  | [optional] 
**Output** | **String** |  | [optional] 
**VarError** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Duration** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**CreatedBy** | [**ClusterHistoryCreatedBy**](ClusterHistoryCreatedBy.md) |  | [optional] 
**UpdatedBy** | [**ClusterHistoryCreatedBy**](ClusterHistoryCreatedBy.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterHistoryEventsInner = Initialize-PSOpenAPIToolsClusterHistoryEventsInner  -Id null `
 -ProcessId null `
 -AccountId null `
 -UniqueId null `
 -ProcessType null `
 -Description null `
 -RefType null `
 -RefId null `
 -SubType null `
 -SubId null `
 -ZoneId null `
 -IntegrationId null `
 -InstanceId null `
 -ContainerId null `
 -ServerId null `
 -ContainerName null `
 -DisplayName null `
 -Status null `
 -Reason null `
 -Percent null `
 -StatusEta null `
 -Message null `
 -Output null `
 -VarError null `
 -StartDate null `
 -EndDate null `
 -Duration null `
 -DateCreated null `
 -LastUpdated null `
 -CreatedBy null `
 -UpdatedBy null
```

- Convert the resource to JSON
```powershell
$ClusterHistoryEventsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

