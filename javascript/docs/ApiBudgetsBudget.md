# MorpheusApi.ApiBudgetsBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** |  | 
**description** | **String** |  | [optional] 
**scope** | **String** |  | [optional] [default to &#39;account&#39;]
**period** | **String** |  | [optional] [default to &#39;year&#39;]
**year** | **Number** |  | [optional] 
**startDate** | **Date** |  | [optional] 
**endDate** | **Date** |  | [optional] 
**interval** | **String** |  | [optional] [default to &#39;year&#39;]
**scopeTenantId** | **Number** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;tenant  | [optional] 
**scopeGroupId** | **Number** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;group   | [optional] 
**scopeCloudId** | **Number** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;cloud  | [optional] 
**scopeUserId** | **Number** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;user  | [optional] 
**costs** | **[Number]** |  | [optional] 
**enabled** | **Boolean** |  | [optional] [default to true]



## Enum: ScopeEnum


* `account` (value: `"account"`)

* `group` (value: `"group"`)

* `cloud` (value: `"cloud"`)

* `user` (value: `"user"`)





## Enum: IntervalEnum


* `year` (value: `"year"`)

* `quarter` (value: `"quarter"`)

* `month` (value: `"month"`)




