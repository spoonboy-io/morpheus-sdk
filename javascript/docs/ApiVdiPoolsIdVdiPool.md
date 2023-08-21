# MorpheusApi.ApiVdiPoolsIdVdiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Virtual Desktop name | [optional] 
**description** | **String** | Virtual Desktop description | [optional] 
**owner** | **Number** | Owner (User) ID | [optional] 
**minIdle** | **Number** | Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.  | [optional] 
**initialPoolSize** | **Number** | The initial size of the pool to be allocated on creation | [optional] 
**maxIdle** | **Number** | Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances  | [optional] 
**maxPoolSize** | **Number** | Max limit on number of allocations and instances within the pool.  | [optional] 
**allocationTimeoutMinutes** | **Number** | Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.  | [optional] 
**persistentUser** | **Boolean** | Persistent Desktop Pool | [optional] [default to false]
**recyclable** | **Boolean** | Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) | [optional] [default to false]
**allowCopy** | **Boolean** | Allow copy from desktop | [optional] [default to false]
**allowPrinter** | **Boolean** | Allow local printers from Desktop | [optional] [default to false]
**allowFileshare** | **Boolean** | Allow File Share | [optional] [default to false]
**allowHypervisorConsole** | **Boolean** | Allow hypervisor console | [optional] [default to false]
**autoCreateLocalUserOnReservation** | **Boolean** | Auto create local user upon reservation | [optional] [default to false]
**enabled** | **Boolean** | Can be used to enable or disable the VDI pool | [optional] [default to true]
**iconPath** | **String** | The relative location of an icon image | [optional] 
**apps** | **[Number]** | Array of VDI App IDs | [optional] 
**gateway** | **Number** | VDI Gateway ID | [optional] 
**instanceConfig** | **String** | Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values. | [optional] 
**config** | [**ApiVdiPoolsIdVdiPoolConfig**](ApiVdiPoolsIdVdiPoolConfig.md) |  | [optional] 
**guestConsoleJumpHost** | **String** | Guest Console Jump Host | [optional] 
**guestConsoleJumpPort** | **Number** | Guest Console Jump Port | [optional] 
**guestConsoleJumpUsername** | **String** | Guest Console Jump Username | [optional] 
**guestConsoleJumpPassword** | **String** | Guest Console Jump Password | [optional] 
**guestConsoleJumpKeypair** | **Number** | Guest Console Jump Key Pair. see &#x60;Key Pair&#x60; | [optional] 


