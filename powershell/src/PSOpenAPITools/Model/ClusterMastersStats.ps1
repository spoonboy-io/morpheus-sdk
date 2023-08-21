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

.PARAMETER UsedStorage
No description available.
.PARAMETER ReservedStorage
No description available.
.PARAMETER MaxStorage
No description available.
.PARAMETER UsedMemory
No description available.
.PARAMETER ReservedMemory
No description available.
.PARAMETER MaxMemory
No description available.
.PARAMETER CpuUsage
No description available.
.OUTPUTS

ClusterMastersStats<PSCustomObject>
#>

function Initialize-ClusterMastersStats {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${UsedStorage},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ReservedStorage},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxStorage},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${UsedMemory},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ReservedMemory},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MaxMemory},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsage}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterMastersStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "usedStorage" = ${UsedStorage}
            "reservedStorage" = ${ReservedStorage}
            "maxStorage" = ${MaxStorage}
            "usedMemory" = ${UsedMemory}
            "reservedMemory" = ${ReservedMemory}
            "maxMemory" = ${MaxMemory}
            "cpuUsage" = ${CpuUsage}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterMastersStats<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterMastersStats<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterMastersStats<PSCustomObject>
#>
function ConvertFrom-JsonToClusterMastersStats {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterMastersStats' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterMastersStats
        $AllProperties = ("usedStorage", "reservedStorage", "maxStorage", "usedMemory", "reservedMemory", "maxMemory", "cpuUsage")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedStorage"))) { #optional property not found
            $UsedStorage = $null
        } else {
            $UsedStorage = $JsonParameters.PSobject.Properties["usedStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "reservedStorage"))) { #optional property not found
            $ReservedStorage = $null
        } else {
            $ReservedStorage = $JsonParameters.PSobject.Properties["reservedStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxStorage"))) { #optional property not found
            $MaxStorage = $null
        } else {
            $MaxStorage = $JsonParameters.PSobject.Properties["maxStorage"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemory"))) { #optional property not found
            $UsedMemory = $null
        } else {
            $UsedMemory = $JsonParameters.PSobject.Properties["usedMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "reservedMemory"))) { #optional property not found
            $ReservedMemory = $null
        } else {
            $ReservedMemory = $JsonParameters.PSobject.Properties["reservedMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemory"))) { #optional property not found
            $MaxMemory = $null
        } else {
            $MaxMemory = $JsonParameters.PSobject.Properties["maxMemory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsage"))) { #optional property not found
            $CpuUsage = $null
        } else {
            $CpuUsage = $JsonParameters.PSobject.Properties["cpuUsage"].value
        }

        $PSO = [PSCustomObject]@{
            "usedStorage" = ${UsedStorage}
            "reservedStorage" = ${ReservedStorage}
            "maxStorage" = ${MaxStorage}
            "usedMemory" = ${UsedMemory}
            "reservedMemory" = ${ReservedMemory}
            "maxMemory" = ${MaxMemory}
            "cpuUsage" = ${CpuUsage}
        }

        return $PSO
    }

}
