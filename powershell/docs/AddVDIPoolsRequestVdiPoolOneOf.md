# AddVDIPoolsRequestVdiPoolOneOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Virtual Desktop name | 
**Description** | **String** | Virtual Desktop description | [optional] 
**Owner** | **Int64** | Owner (User) ID | [optional] 
**MinIdle** | **Decimal** | Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.  | [optional] 
**InitialPoolSize** | **Decimal** | The initial size of the pool to be allocated on creation | [optional] 
**MaxIdle** | **Decimal** | Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances  | [optional] 
**MaxPoolSize** | **Decimal** | Max limit on number of allocations and instances within the pool.  | 
**AllocationTimeoutMinutes** | **Decimal** | Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.  | [optional] 
**PersistentUser** | **Boolean** | Persistent Desktop Pool | [optional] [default to $false]
**Recyclable** | **Boolean** | Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) | [optional] [default to $false]
**AllowCopy** | **Boolean** | Allow copy from desktop | [optional] [default to $false]
**AllowPrinter** | **Boolean** | Allow local printers from Desktop | [optional] [default to $false]
**AllowFileshare** | **Boolean** | Allow File Share | [optional] [default to $false]
**AllowHypervisorConsole** | **Boolean** | Allow hypervisor console | [optional] [default to $false]
**AutoCreateLocalUserOnReservation** | **Boolean** | Auto create local user upon reservation | [optional] [default to $false]
**Enabled** | **Boolean** | Can be used to enable or disable the VDI pool | [optional] [default to $true]
**IconPath** | **String** | The relative location of an icon image | [optional] 
**Apps** | **Int64[]** | Array of VDI App IDs | [optional] 
**Gateway** | **Int64** | VDI Gateway ID | [optional] 
**InstanceConfig** | **String** | Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values. | 
**Config** | [**AddVDIPoolsRequestVdiPoolOneOfConfig**](AddVDIPoolsRequestVdiPoolOneOfConfig.md) |  | [optional] 
**GuestConsoleJumpHost** | **String** | Guest Console Jump Host | [optional] 
**GuestConsoleJumpPort** | **Int64** | Guest Console Jump Port | [optional] 
**GuestConsoleJumpUsername** | **String** | Guest Console Jump Username | [optional] 
**GuestConsoleJumpPassword** | **String** | Guest Console Jump Password | [optional] 
**GuestConsoleJumpKeypair** | **Int64** | Guest Console Jump Key Pair. see &#x60;Key Pair&#x60; | [optional] 

## Examples

- Prepare the resource
```powershell
$AddVDIPoolsRequestVdiPoolOneOf = Initialize-PSOpenAPIToolsAddVDIPoolsRequestVdiPoolOneOf  -Name null `
 -Description null `
 -Owner 21 `
 -MinIdle 5 `
 -InitialPoolSize 10 `
 -MaxIdle 2 `
 -MaxPoolSize 50 `
 -AllocationTimeoutMinutes 20 `
 -PersistentUser null `
 -Recyclable null `
 -AllowCopy null `
 -AllowPrinter null `
 -AllowFileshare null `
 -AllowHypervisorConsole null `
 -AutoCreateLocalUserOnReservation null `
 -Enabled null `
 -IconPath /assets/containers-png/windows.png `
 -Apps null `
 -Gateway null `
 -InstanceConfig null `
 -Config null `
 -GuestConsoleJumpHost null `
 -GuestConsoleJumpPort null `
 -GuestConsoleJumpUsername null `
 -GuestConsoleJumpPassword null `
 -GuestConsoleJumpKeypair null
```

- Convert the resource to JSON
```powershell
$AddVDIPoolsRequestVdiPoolOneOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

