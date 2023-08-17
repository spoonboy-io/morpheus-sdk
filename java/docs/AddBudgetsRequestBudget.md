

# AddBudgetsRequestBudget


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** |  |  |
|**description** | **String** |  |  [optional] |
|**scope** | [**ScopeEnum**](#ScopeEnum) |  |  [optional] |
|**period** | **String** |  |  [optional] |
|**year** | **Long** |  |  [optional] |
|**startDate** | **OffsetDateTime** |  |  [optional] |
|**endDate** | **OffsetDateTime** |  |  [optional] |
|**interval** | [**IntervalEnum**](#IntervalEnum) |  |  [optional] |
|**scopeTenantId** | **Long** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;tenant  |  [optional] |
|**scopeGroupId** | **Long** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;group   |  [optional] |
|**scopeCloudId** | **Long** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;cloud  |  [optional] |
|**scopeUserId** | **Long** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;user  |  [optional] |
|**costs** | **List&lt;Long&gt;** |  |  [optional] |
|**enabled** | **Boolean** |  |  [optional] |



## Enum: ScopeEnum

| Name | Value |
|---- | -----|
| ACCOUNT | &quot;account&quot; |
| GROUP | &quot;group&quot; |
| CLOUD | &quot;cloud&quot; |
| USER | &quot;user&quot; |



## Enum: IntervalEnum

| Name | Value |
|---- | -----|
| YEAR | &quot;year&quot; |
| QUARTER | &quot;quarter&quot; |
| MONTH | &quot;month&quot; |



