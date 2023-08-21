# MorpheusApi.PolicyGroupCreatePolicyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **String** | The policy type | [optional] 
**config** | [**OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject**](OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject.md) | A map of config values. The expected values vary by policyType. | [optional] 
**enabled** | **Boolean** | Set to false to disable | [optional] [default to true]
**refType** | **String** | Scope object type | [optional] 
**refId** | **Number** | Scope object ID (&#x60;group&#x60;) | [optional] 
**accounts** | **[Number]** | Array of tenants to scope the policy to | [optional] 
**eachUser** | **Boolean** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 



## Enum: CodeEnum


* `Approve Delete` (value: `"Approve Delete"`)

* `Approve Provision` (value: `"Approve Provision"`)

* `Approve Reconfigure` (value: `"Approve Reconfigure"`)

* `Backup Creation` (value: `"Backup Creation"`)

* `Backup Targets` (value: `"Backup Targets"`)

* `Budget` (value: `"Budget"`)

* `Cluster Resource Name` (value: `"Cluster Resource Name"`)

* `Cypher Access` (value: `"Cypher Access"`)

* `Delayed Delete` (value: `"Delayed Delete"`)

* `Expiration` (value: `"Expiration"`)

* `File Share Storage Quota` (value: `"File Share Storage Quota"`)

* `Hostname` (value: `"Hostname"`)

* `Instance Name` (value: `"Instance Name"`)

* `Max Containers` (value: `"Max Containers"`)

* `Max Cores` (value: `"Max Cores"`)

* `Max Hosts` (value: `"Max Hosts"`)

* `Max Load Balancer Pools` (value: `"Max Load Balancer Pools"`)

* `Max Memory` (value: `"Max Memory"`)

* `Max Pool Members` (value: `"Max Pool Members"`)

* `Max Storage` (value: `"Max Storage"`)

* `Max Virtual Servers` (value: `"Max Virtual Servers"`)

* `Max VMs` (value: `"Max VMs"`)

* `Message of the Day` (value: `"Message of the Day"`)

* `Network Quota` (value: `"Network Quota"`)

* `Object Storage Quota` (value: `"Object Storage Quota"`)

* `Power Schedule` (value: `"Power Schedule"`)

* `Router Quota` (value: `"Router Quota"`)

* `Shutdown` (value: `"Shutdown"`)

* `Storage Server Storage Quota` (value: `"Storage Server Storage Quota"`)

* `Tags` (value: `"Tags"`)

* `User Creation` (value: `"User Creation"`)

* `User Group Creation` (value: `"User Group Creation"`)

* `Workflow` (value: `"Workflow"`)





## Enum: RefTypeEnum


* `ComputeSite` (value: `"ComputeSite"`)




