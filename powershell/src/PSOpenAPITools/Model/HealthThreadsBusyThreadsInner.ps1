#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER ThreadId
No description available.
.PARAMETER Name
No description available.
.PARAMETER CpuTime
No description available.
.PARAMETER BlockedTime
No description available.
.PARAMETER LockName
No description available.
.PARAMETER LockOwnerId
No description available.
.PARAMETER LockOwnerName
No description available.
.PARAMETER State
No description available.
.PARAMETER WaitedCount
No description available.
.PARAMETER WaitedTime
No description available.
.PARAMETER IsInNative
No description available.
.PARAMETER IsSuspended
No description available.
.PARAMETER LockedMonitors
No description available.
.PARAMETER LockedSynchronizers
No description available.
.PARAMETER LockInfo
No description available.
.PARAMETER CurrentLines
No description available.
.PARAMETER CpuPercent
No description available.
.OUTPUTS

HealthThreadsBusyThreadsInner<PSCustomObject>
#>

function Initialize-HealthThreadsBusyThreadsInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ThreadId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${CpuTime},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${BlockedTime},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LockName},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${LockOwnerId},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LockOwnerName},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${State},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${WaitedCount},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${WaitedTime},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IsInNative},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IsSuspended},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${LockedMonitors},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${LockedSynchronizers},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${LockInfo},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CurrentLines},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuPercent}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HealthThreadsBusyThreadsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "threadId" = ${ThreadId}
            "name" = ${Name}
            "cpuTime" = ${CpuTime}
            "blockedTime" = ${BlockedTime}
            "lockName" = ${LockName}
            "lockOwnerId" = ${LockOwnerId}
            "lockOwnerName" = ${LockOwnerName}
            "state" = ${State}
            "waitedCount" = ${WaitedCount}
            "waitedTime" = ${WaitedTime}
            "isInNative" = ${IsInNative}
            "isSuspended" = ${IsSuspended}
            "lockedMonitors" = ${LockedMonitors}
            "lockedSynchronizers" = ${LockedSynchronizers}
            "lockInfo" = ${LockInfo}
            "currentLines" = ${CurrentLines}
            "cpuPercent" = ${CpuPercent}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HealthThreadsBusyThreadsInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to HealthThreadsBusyThreadsInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HealthThreadsBusyThreadsInner<PSCustomObject>
#>
function ConvertFrom-JsonToHealthThreadsBusyThreadsInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HealthThreadsBusyThreadsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HealthThreadsBusyThreadsInner
        $AllProperties = ("threadId", "name", "cpuTime", "blockedTime", "lockName", "lockOwnerId", "lockOwnerName", "state", "waitedCount", "waitedTime", "isInNative", "isSuspended", "lockedMonitors", "lockedSynchronizers", "lockInfo", "currentLines", "cpuPercent")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "threadId"))) { #optional property not found
            $ThreadId = $null
        } else {
            $ThreadId = $JsonParameters.PSobject.Properties["threadId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuTime"))) { #optional property not found
            $CpuTime = $null
        } else {
            $CpuTime = $JsonParameters.PSobject.Properties["cpuTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blockedTime"))) { #optional property not found
            $BlockedTime = $null
        } else {
            $BlockedTime = $JsonParameters.PSobject.Properties["blockedTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lockName"))) { #optional property not found
            $LockName = $null
        } else {
            $LockName = $JsonParameters.PSobject.Properties["lockName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lockOwnerId"))) { #optional property not found
            $LockOwnerId = $null
        } else {
            $LockOwnerId = $JsonParameters.PSobject.Properties["lockOwnerId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lockOwnerName"))) { #optional property not found
            $LockOwnerName = $null
        } else {
            $LockOwnerName = $JsonParameters.PSobject.Properties["lockOwnerName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "state"))) { #optional property not found
            $State = $null
        } else {
            $State = $JsonParameters.PSobject.Properties["state"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "waitedCount"))) { #optional property not found
            $WaitedCount = $null
        } else {
            $WaitedCount = $JsonParameters.PSobject.Properties["waitedCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "waitedTime"))) { #optional property not found
            $WaitedTime = $null
        } else {
            $WaitedTime = $JsonParameters.PSobject.Properties["waitedTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "isInNative"))) { #optional property not found
            $IsInNative = $null
        } else {
            $IsInNative = $JsonParameters.PSobject.Properties["isInNative"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "isSuspended"))) { #optional property not found
            $IsSuspended = $null
        } else {
            $IsSuspended = $JsonParameters.PSobject.Properties["isSuspended"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lockedMonitors"))) { #optional property not found
            $LockedMonitors = $null
        } else {
            $LockedMonitors = $JsonParameters.PSobject.Properties["lockedMonitors"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lockedSynchronizers"))) { #optional property not found
            $LockedSynchronizers = $null
        } else {
            $LockedSynchronizers = $JsonParameters.PSobject.Properties["lockedSynchronizers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lockInfo"))) { #optional property not found
            $LockInfo = $null
        } else {
            $LockInfo = $JsonParameters.PSobject.Properties["lockInfo"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "currentLines"))) { #optional property not found
            $CurrentLines = $null
        } else {
            $CurrentLines = $JsonParameters.PSobject.Properties["currentLines"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuPercent"))) { #optional property not found
            $CpuPercent = $null
        } else {
            $CpuPercent = $JsonParameters.PSobject.Properties["cpuPercent"].value
        }

        $PSO = [PSCustomObject]@{
            "threadId" = ${ThreadId}
            "name" = ${Name}
            "cpuTime" = ${CpuTime}
            "blockedTime" = ${BlockedTime}
            "lockName" = ${LockName}
            "lockOwnerId" = ${LockOwnerId}
            "lockOwnerName" = ${LockOwnerName}
            "state" = ${State}
            "waitedCount" = ${WaitedCount}
            "waitedTime" = ${WaitedTime}
            "isInNative" = ${IsInNative}
            "isSuspended" = ${IsSuspended}
            "lockedMonitors" = ${LockedMonitors}
            "lockedSynchronizers" = ${LockedSynchronizers}
            "lockInfo" = ${LockInfo}
            "currentLines" = ${CurrentLines}
            "cpuPercent" = ${CpuPercent}
        }

        return $PSO
    }

}

