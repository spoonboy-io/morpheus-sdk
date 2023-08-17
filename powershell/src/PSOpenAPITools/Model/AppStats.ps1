#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER UsedMemory
No description available.
.PARAMETER MaxMemory
No description available.
.PARAMETER UsedStorage
No description available.
.PARAMETER MaxStorage
No description available.
.PARAMETER Running
No description available.
.PARAMETER Total
No description available.
.PARAMETER CpuUsage
No description available.
.PARAMETER InstanceCount
No description available.
.PARAMETER InstanceDayCount
No description available.
.PARAMETER InstanceDayCountTotal
No description available.
.OUTPUTS

AppStats<PSCustomObject>
#>

function Initialize-AppStats {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${UsedMemory},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxMemory},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${UsedStorage},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxStorage},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Running},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Total},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsage},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${InstanceCount},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [Int64[]]
        ${InstanceDayCount},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${InstanceDayCountTotal}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "usedMemory" = ${UsedMemory}
            "maxMemory" = ${MaxMemory}
            "usedStorage" = ${UsedStorage}
            "maxStorage" = ${MaxStorage}
            "running" = ${Running}
            "total" = ${Total}
            "cpuUsage" = ${CpuUsage}
            "instanceCount" = ${InstanceCount}
            "instanceDayCount" = ${InstanceDayCount}
            "instanceDayCountTotal" = ${InstanceDayCountTotal}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppStats<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppStats<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppStats<PSCustomObject>
#>
function ConvertFrom-JsonToAppStats {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppStats
        $AllProperties = ("usedMemory", "maxMemory", "usedStorage", "maxStorage", "running", "total", "cpuUsage", "instanceCount", "instanceDayCount", "instanceDayCountTotal")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemory"))) { #optional property not found
            $UsedMemory = $null
        } else {
            $UsedMemory = $JsonParameters.PSobject.Properties["usedMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemory"))) { #optional property not found
            $MaxMemory = $null
        } else {
            $MaxMemory = $JsonParameters.PSobject.Properties["maxMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedStorage"))) { #optional property not found
            $UsedStorage = $null
        } else {
            $UsedStorage = $JsonParameters.PSobject.Properties["usedStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxStorage"))) { #optional property not found
            $MaxStorage = $null
        } else {
            $MaxStorage = $JsonParameters.PSobject.Properties["maxStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "running"))) { #optional property not found
            $Running = $null
        } else {
            $Running = $JsonParameters.PSobject.Properties["running"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "total"))) { #optional property not found
            $Total = $null
        } else {
            $Total = $JsonParameters.PSobject.Properties["total"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsage"))) { #optional property not found
            $CpuUsage = $null
        } else {
            $CpuUsage = $JsonParameters.PSobject.Properties["cpuUsage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceCount"))) { #optional property not found
            $InstanceCount = $null
        } else {
            $InstanceCount = $JsonParameters.PSobject.Properties["instanceCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceDayCount"))) { #optional property not found
            $InstanceDayCount = $null
        } else {
            $InstanceDayCount = $JsonParameters.PSobject.Properties["instanceDayCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceDayCountTotal"))) { #optional property not found
            $InstanceDayCountTotal = $null
        } else {
            $InstanceDayCountTotal = $JsonParameters.PSobject.Properties["instanceDayCountTotal"].value
        }

        $PSO = [PSCustomObject]@{
            "usedMemory" = ${UsedMemory}
            "maxMemory" = ${MaxMemory}
            "usedStorage" = ${UsedStorage}
            "maxStorage" = ${MaxStorage}
            "running" = ${Running}
            "total" = ${Total}
            "cpuUsage" = ${CpuUsage}
            "instanceCount" = ${InstanceCount}
            "instanceDayCount" = ${InstanceDayCount}
            "instanceDayCountTotal" = ${InstanceDayCountTotal}
        }

        return $PSO
    }

}

