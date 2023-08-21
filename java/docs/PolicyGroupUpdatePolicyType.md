

# PolicyGroupUpdatePolicyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | [**CodeEnum**](#CodeEnum) | The policy type |  [optional]
**config** | [**OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject**](OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject.md) | A map of config values. The expected values vary by policyType. |  [optional]
**enabled** | **Boolean** | Set to false to disable |  [optional]
**refType** | [**RefTypeEnum**](#RefTypeEnum) | Scope object type |  [optional]
**refId** | **Long** | Scope object ID (&#x60;group&#x60;) |  [optional]
**accounts** | **List&lt;Long&gt;** | Array of tenants to scope the policy to |  [optional]
**eachUser** | **Boolean** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; |  [optional]



## Enum: CodeEnum

Name | Value
---- | -----
APPROVE_DELETE | &quot;Approve Delete&quot;
APPROVE_PROVISION | &quot;Approve Provision&quot;
APPROVE_RECONFIGURE | &quot;Approve Reconfigure&quot;
BACKUP_CREATION | &quot;Backup Creation&quot;
BACKUP_TARGETS | &quot;Backup Targets&quot;
BUDGET | &quot;Budget&quot;
CLUSTER_RESOURCE_NAME | &quot;Cluster Resource Name&quot;
CYPHER_ACCESS | &quot;Cypher Access&quot;
DELAYED_DELETE | &quot;Delayed Delete&quot;
EXPIRATION | &quot;Expiration&quot;
FILE_SHARE_STORAGE_QUOTA | &quot;File Share Storage Quota&quot;
HOSTNAME | &quot;Hostname&quot;
INSTANCE_NAME | &quot;Instance Name&quot;
MAX_CONTAINERS | &quot;Max Containers&quot;
MAX_CORES | &quot;Max Cores&quot;
MAX_HOSTS | &quot;Max Hosts&quot;
MAX_LOAD_BALANCER_POOLS | &quot;Max Load Balancer Pools&quot;
MAX_MEMORY | &quot;Max Memory&quot;
MAX_POOL_MEMBERS | &quot;Max Pool Members&quot;
MAX_STORAGE | &quot;Max Storage&quot;
MAX_VIRTUAL_SERVERS | &quot;Max Virtual Servers&quot;
MAX_VMS | &quot;Max VMs&quot;
MESSAGE_OF_THE_DAY | &quot;Message of the Day&quot;
NETWORK_QUOTA | &quot;Network Quota&quot;
OBJECT_STORAGE_QUOTA | &quot;Object Storage Quota&quot;
POWER_SCHEDULE | &quot;Power Schedule&quot;
ROUTER_QUOTA | &quot;Router Quota&quot;
SHUTDOWN | &quot;Shutdown&quot;
STORAGE_SERVER_STORAGE_QUOTA | &quot;Storage Server Storage Quota&quot;
TAGS | &quot;Tags&quot;
USER_CREATION | &quot;User Creation&quot;
USER_GROUP_CREATION | &quot;User Group Creation&quot;
WORKFLOW | &quot;Workflow&quot;



## Enum: RefTypeEnum

Name | Value
---- | -----
COMPUTESITE | &quot;ComputeSite&quot;



