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

.PARAMETER Exists
No description available.
.PARAMETER ObjectId
No description available.
.PARAMETER CpuTotalTimeCount
No description available.
.PARAMETER CpuTotalTimeMin
No description available.
.PARAMETER CpuTotalTimeMax
No description available.
.PARAMETER CpuTotalTimeAvg
No description available.
.PARAMETER CpuTotalTimeSum
No description available.
.PARAMETER CpuIdleTimeCount
No description available.
.PARAMETER CpuIdleTimeMin
No description available.
.PARAMETER CpuIdleTimeMax
No description available.
.PARAMETER CpuIdleTimeAvg
No description available.
.PARAMETER CpuIdleTimeSum
No description available.
.PARAMETER CpuUsageCount
No description available.
.PARAMETER CpuUsageMin
No description available.
.PARAMETER CpuUsageMax
No description available.
.PARAMETER CpuUsageAvg
No description available.
.PARAMETER CpuUsageSum
No description available.
.PARAMETER MaxMemoryCount
No description available.
.PARAMETER MaxMemoryMin
No description available.
.PARAMETER MaxMemoryMax
No description available.
.PARAMETER MaxMemoryAvg
No description available.
.PARAMETER MaxMemorySum
No description available.
.PARAMETER CpuUserTimeCount
No description available.
.PARAMETER CpuUserTimeMin
No description available.
.PARAMETER CpuUserTimeMax
No description available.
.PARAMETER CpuUserTimeAvg
No description available.
.PARAMETER CpuUserTimeSum
No description available.
.PARAMETER CpuSystemTimeCount
No description available.
.PARAMETER CpuSystemTimeMin
No description available.
.PARAMETER CpuSystemTimeMax
No description available.
.PARAMETER CpuSystemTimeAvg
No description available.
.PARAMETER CpuSystemTimeSum
No description available.
.PARAMETER UsedMemoryCount
No description available.
.PARAMETER UsedMemoryMin
No description available.
.PARAMETER UsedMemoryMax
No description available.
.PARAMETER UsedMemoryAvg
No description available.
.PARAMETER UsedMemorySum
No description available.
.PARAMETER FreeMemoryCount
No description available.
.PARAMETER FreeMemoryMin
No description available.
.PARAMETER FreeMemoryMax
No description available.
.PARAMETER FreeMemoryAvg
No description available.
.PARAMETER FreeMemorySum
No description available.
.OUTPUTS

GuidanceVmwareSizingConfig<PSCustomObject>
#>

function Initialize-GuidanceVmwareSizingConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Exists},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ObjectId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuTotalTimeCount},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuTotalTimeMin},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuTotalTimeMax},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuTotalTimeAvg},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuTotalTimeSum},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuIdleTimeCount},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuIdleTimeMin},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuIdleTimeMax},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuIdleTimeAvg},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuIdleTimeSum},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsageCount},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsageMin},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsageMax},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsageAvg},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUsageSum},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxMemoryCount},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxMemoryMin},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxMemoryMax},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxMemoryAvg},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxMemorySum},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUserTimeCount},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUserTimeMin},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUserTimeMax},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUserTimeAvg},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuUserTimeSum},
        [Parameter(Position = 27, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuSystemTimeCount},
        [Parameter(Position = 28, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuSystemTimeMin},
        [Parameter(Position = 29, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuSystemTimeMax},
        [Parameter(Position = 30, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuSystemTimeAvg},
        [Parameter(Position = 31, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${CpuSystemTimeSum},
        [Parameter(Position = 32, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${UsedMemoryCount},
        [Parameter(Position = 33, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${UsedMemoryMin},
        [Parameter(Position = 34, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${UsedMemoryMax},
        [Parameter(Position = 35, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${UsedMemoryAvg},
        [Parameter(Position = 36, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${UsedMemorySum},
        [Parameter(Position = 37, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${FreeMemoryCount},
        [Parameter(Position = 38, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${FreeMemoryMin},
        [Parameter(Position = 39, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${FreeMemoryMax},
        [Parameter(Position = 40, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${FreeMemoryAvg},
        [Parameter(Position = 41, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${FreeMemorySum}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GuidanceVmwareSizingConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "exists" = ${Exists}
            "objectId" = ${ObjectId}
            "cpuTotalTimeCount" = ${CpuTotalTimeCount}
            "cpuTotalTimeMin" = ${CpuTotalTimeMin}
            "cpuTotalTimeMax" = ${CpuTotalTimeMax}
            "cpuTotalTimeAvg" = ${CpuTotalTimeAvg}
            "cpuTotalTimeSum" = ${CpuTotalTimeSum}
            "cpuIdleTimeCount" = ${CpuIdleTimeCount}
            "cpuIdleTimeMin" = ${CpuIdleTimeMin}
            "cpuIdleTimeMax" = ${CpuIdleTimeMax}
            "cpuIdleTimeAvg" = ${CpuIdleTimeAvg}
            "cpuIdleTimeSum" = ${CpuIdleTimeSum}
            "cpuUsageCount" = ${CpuUsageCount}
            "cpuUsageMin" = ${CpuUsageMin}
            "cpuUsageMax" = ${CpuUsageMax}
            "cpuUsageAvg" = ${CpuUsageAvg}
            "cpuUsageSum" = ${CpuUsageSum}
            "maxMemoryCount" = ${MaxMemoryCount}
            "maxMemoryMin" = ${MaxMemoryMin}
            "maxMemoryMax" = ${MaxMemoryMax}
            "maxMemoryAvg" = ${MaxMemoryAvg}
            "maxMemorySum" = ${MaxMemorySum}
            "cpuUserTimeCount" = ${CpuUserTimeCount}
            "cpuUserTimeMin" = ${CpuUserTimeMin}
            "cpuUserTimeMax" = ${CpuUserTimeMax}
            "cpuUserTimeAvg" = ${CpuUserTimeAvg}
            "cpuUserTimeSum" = ${CpuUserTimeSum}
            "cpuSystemTimeCount" = ${CpuSystemTimeCount}
            "cpuSystemTimeMin" = ${CpuSystemTimeMin}
            "cpuSystemTimeMax" = ${CpuSystemTimeMax}
            "cpuSystemTimeAvg" = ${CpuSystemTimeAvg}
            "cpuSystemTimeSum" = ${CpuSystemTimeSum}
            "usedMemoryCount" = ${UsedMemoryCount}
            "usedMemoryMin" = ${UsedMemoryMin}
            "usedMemoryMax" = ${UsedMemoryMax}
            "usedMemoryAvg" = ${UsedMemoryAvg}
            "usedMemorySum" = ${UsedMemorySum}
            "freeMemoryCount" = ${FreeMemoryCount}
            "freeMemoryMin" = ${FreeMemoryMin}
            "freeMemoryMax" = ${FreeMemoryMax}
            "freeMemoryAvg" = ${FreeMemoryAvg}
            "freeMemorySum" = ${FreeMemorySum}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GuidanceVmwareSizingConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to GuidanceVmwareSizingConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GuidanceVmwareSizingConfig<PSCustomObject>
#>
function ConvertFrom-JsonToGuidanceVmwareSizingConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GuidanceVmwareSizingConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GuidanceVmwareSizingConfig
        $AllProperties = ("exists", "objectId", "cpuTotalTimeCount", "cpuTotalTimeMin", "cpuTotalTimeMax", "cpuTotalTimeAvg", "cpuTotalTimeSum", "cpuIdleTimeCount", "cpuIdleTimeMin", "cpuIdleTimeMax", "cpuIdleTimeAvg", "cpuIdleTimeSum", "cpuUsageCount", "cpuUsageMin", "cpuUsageMax", "cpuUsageAvg", "cpuUsageSum", "maxMemoryCount", "maxMemoryMin", "maxMemoryMax", "maxMemoryAvg", "maxMemorySum", "cpuUserTimeCount", "cpuUserTimeMin", "cpuUserTimeMax", "cpuUserTimeAvg", "cpuUserTimeSum", "cpuSystemTimeCount", "cpuSystemTimeMin", "cpuSystemTimeMax", "cpuSystemTimeAvg", "cpuSystemTimeSum", "usedMemoryCount", "usedMemoryMin", "usedMemoryMax", "usedMemoryAvg", "usedMemorySum", "freeMemoryCount", "freeMemoryMin", "freeMemoryMax", "freeMemoryAvg", "freeMemorySum")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "exists"))) { #optional property not found
            $Exists = $null
        } else {
            $Exists = $JsonParameters.PSobject.Properties["exists"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "objectId"))) { #optional property not found
            $ObjectId = $null
        } else {
            $ObjectId = $JsonParameters.PSobject.Properties["objectId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuTotalTimeCount"))) { #optional property not found
            $CpuTotalTimeCount = $null
        } else {
            $CpuTotalTimeCount = $JsonParameters.PSobject.Properties["cpuTotalTimeCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuTotalTimeMin"))) { #optional property not found
            $CpuTotalTimeMin = $null
        } else {
            $CpuTotalTimeMin = $JsonParameters.PSobject.Properties["cpuTotalTimeMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuTotalTimeMax"))) { #optional property not found
            $CpuTotalTimeMax = $null
        } else {
            $CpuTotalTimeMax = $JsonParameters.PSobject.Properties["cpuTotalTimeMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuTotalTimeAvg"))) { #optional property not found
            $CpuTotalTimeAvg = $null
        } else {
            $CpuTotalTimeAvg = $JsonParameters.PSobject.Properties["cpuTotalTimeAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuTotalTimeSum"))) { #optional property not found
            $CpuTotalTimeSum = $null
        } else {
            $CpuTotalTimeSum = $JsonParameters.PSobject.Properties["cpuTotalTimeSum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuIdleTimeCount"))) { #optional property not found
            $CpuIdleTimeCount = $null
        } else {
            $CpuIdleTimeCount = $JsonParameters.PSobject.Properties["cpuIdleTimeCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuIdleTimeMin"))) { #optional property not found
            $CpuIdleTimeMin = $null
        } else {
            $CpuIdleTimeMin = $JsonParameters.PSobject.Properties["cpuIdleTimeMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuIdleTimeMax"))) { #optional property not found
            $CpuIdleTimeMax = $null
        } else {
            $CpuIdleTimeMax = $JsonParameters.PSobject.Properties["cpuIdleTimeMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuIdleTimeAvg"))) { #optional property not found
            $CpuIdleTimeAvg = $null
        } else {
            $CpuIdleTimeAvg = $JsonParameters.PSobject.Properties["cpuIdleTimeAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuIdleTimeSum"))) { #optional property not found
            $CpuIdleTimeSum = $null
        } else {
            $CpuIdleTimeSum = $JsonParameters.PSobject.Properties["cpuIdleTimeSum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsageCount"))) { #optional property not found
            $CpuUsageCount = $null
        } else {
            $CpuUsageCount = $JsonParameters.PSobject.Properties["cpuUsageCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsageMin"))) { #optional property not found
            $CpuUsageMin = $null
        } else {
            $CpuUsageMin = $JsonParameters.PSobject.Properties["cpuUsageMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsageMax"))) { #optional property not found
            $CpuUsageMax = $null
        } else {
            $CpuUsageMax = $JsonParameters.PSobject.Properties["cpuUsageMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsageAvg"))) { #optional property not found
            $CpuUsageAvg = $null
        } else {
            $CpuUsageAvg = $JsonParameters.PSobject.Properties["cpuUsageAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUsageSum"))) { #optional property not found
            $CpuUsageSum = $null
        } else {
            $CpuUsageSum = $JsonParameters.PSobject.Properties["cpuUsageSum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemoryCount"))) { #optional property not found
            $MaxMemoryCount = $null
        } else {
            $MaxMemoryCount = $JsonParameters.PSobject.Properties["maxMemoryCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemoryMin"))) { #optional property not found
            $MaxMemoryMin = $null
        } else {
            $MaxMemoryMin = $JsonParameters.PSobject.Properties["maxMemoryMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemoryMax"))) { #optional property not found
            $MaxMemoryMax = $null
        } else {
            $MaxMemoryMax = $JsonParameters.PSobject.Properties["maxMemoryMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemoryAvg"))) { #optional property not found
            $MaxMemoryAvg = $null
        } else {
            $MaxMemoryAvg = $JsonParameters.PSobject.Properties["maxMemoryAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxMemorySum"))) { #optional property not found
            $MaxMemorySum = $null
        } else {
            $MaxMemorySum = $JsonParameters.PSobject.Properties["maxMemorySum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUserTimeCount"))) { #optional property not found
            $CpuUserTimeCount = $null
        } else {
            $CpuUserTimeCount = $JsonParameters.PSobject.Properties["cpuUserTimeCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUserTimeMin"))) { #optional property not found
            $CpuUserTimeMin = $null
        } else {
            $CpuUserTimeMin = $JsonParameters.PSobject.Properties["cpuUserTimeMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUserTimeMax"))) { #optional property not found
            $CpuUserTimeMax = $null
        } else {
            $CpuUserTimeMax = $JsonParameters.PSobject.Properties["cpuUserTimeMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUserTimeAvg"))) { #optional property not found
            $CpuUserTimeAvg = $null
        } else {
            $CpuUserTimeAvg = $JsonParameters.PSobject.Properties["cpuUserTimeAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuUserTimeSum"))) { #optional property not found
            $CpuUserTimeSum = $null
        } else {
            $CpuUserTimeSum = $JsonParameters.PSobject.Properties["cpuUserTimeSum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuSystemTimeCount"))) { #optional property not found
            $CpuSystemTimeCount = $null
        } else {
            $CpuSystemTimeCount = $JsonParameters.PSobject.Properties["cpuSystemTimeCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuSystemTimeMin"))) { #optional property not found
            $CpuSystemTimeMin = $null
        } else {
            $CpuSystemTimeMin = $JsonParameters.PSobject.Properties["cpuSystemTimeMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuSystemTimeMax"))) { #optional property not found
            $CpuSystemTimeMax = $null
        } else {
            $CpuSystemTimeMax = $JsonParameters.PSobject.Properties["cpuSystemTimeMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuSystemTimeAvg"))) { #optional property not found
            $CpuSystemTimeAvg = $null
        } else {
            $CpuSystemTimeAvg = $JsonParameters.PSobject.Properties["cpuSystemTimeAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cpuSystemTimeSum"))) { #optional property not found
            $CpuSystemTimeSum = $null
        } else {
            $CpuSystemTimeSum = $JsonParameters.PSobject.Properties["cpuSystemTimeSum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemoryCount"))) { #optional property not found
            $UsedMemoryCount = $null
        } else {
            $UsedMemoryCount = $JsonParameters.PSobject.Properties["usedMemoryCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemoryMin"))) { #optional property not found
            $UsedMemoryMin = $null
        } else {
            $UsedMemoryMin = $JsonParameters.PSobject.Properties["usedMemoryMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemoryMax"))) { #optional property not found
            $UsedMemoryMax = $null
        } else {
            $UsedMemoryMax = $JsonParameters.PSobject.Properties["usedMemoryMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemoryAvg"))) { #optional property not found
            $UsedMemoryAvg = $null
        } else {
            $UsedMemoryAvg = $JsonParameters.PSobject.Properties["usedMemoryAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "usedMemorySum"))) { #optional property not found
            $UsedMemorySum = $null
        } else {
            $UsedMemorySum = $JsonParameters.PSobject.Properties["usedMemorySum"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "freeMemoryCount"))) { #optional property not found
            $FreeMemoryCount = $null
        } else {
            $FreeMemoryCount = $JsonParameters.PSobject.Properties["freeMemoryCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "freeMemoryMin"))) { #optional property not found
            $FreeMemoryMin = $null
        } else {
            $FreeMemoryMin = $JsonParameters.PSobject.Properties["freeMemoryMin"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "freeMemoryMax"))) { #optional property not found
            $FreeMemoryMax = $null
        } else {
            $FreeMemoryMax = $JsonParameters.PSobject.Properties["freeMemoryMax"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "freeMemoryAvg"))) { #optional property not found
            $FreeMemoryAvg = $null
        } else {
            $FreeMemoryAvg = $JsonParameters.PSobject.Properties["freeMemoryAvg"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "freeMemorySum"))) { #optional property not found
            $FreeMemorySum = $null
        } else {
            $FreeMemorySum = $JsonParameters.PSobject.Properties["freeMemorySum"].value
        }

        $PSO = [PSCustomObject]@{
            "exists" = ${Exists}
            "objectId" = ${ObjectId}
            "cpuTotalTimeCount" = ${CpuTotalTimeCount}
            "cpuTotalTimeMin" = ${CpuTotalTimeMin}
            "cpuTotalTimeMax" = ${CpuTotalTimeMax}
            "cpuTotalTimeAvg" = ${CpuTotalTimeAvg}
            "cpuTotalTimeSum" = ${CpuTotalTimeSum}
            "cpuIdleTimeCount" = ${CpuIdleTimeCount}
            "cpuIdleTimeMin" = ${CpuIdleTimeMin}
            "cpuIdleTimeMax" = ${CpuIdleTimeMax}
            "cpuIdleTimeAvg" = ${CpuIdleTimeAvg}
            "cpuIdleTimeSum" = ${CpuIdleTimeSum}
            "cpuUsageCount" = ${CpuUsageCount}
            "cpuUsageMin" = ${CpuUsageMin}
            "cpuUsageMax" = ${CpuUsageMax}
            "cpuUsageAvg" = ${CpuUsageAvg}
            "cpuUsageSum" = ${CpuUsageSum}
            "maxMemoryCount" = ${MaxMemoryCount}
            "maxMemoryMin" = ${MaxMemoryMin}
            "maxMemoryMax" = ${MaxMemoryMax}
            "maxMemoryAvg" = ${MaxMemoryAvg}
            "maxMemorySum" = ${MaxMemorySum}
            "cpuUserTimeCount" = ${CpuUserTimeCount}
            "cpuUserTimeMin" = ${CpuUserTimeMin}
            "cpuUserTimeMax" = ${CpuUserTimeMax}
            "cpuUserTimeAvg" = ${CpuUserTimeAvg}
            "cpuUserTimeSum" = ${CpuUserTimeSum}
            "cpuSystemTimeCount" = ${CpuSystemTimeCount}
            "cpuSystemTimeMin" = ${CpuSystemTimeMin}
            "cpuSystemTimeMax" = ${CpuSystemTimeMax}
            "cpuSystemTimeAvg" = ${CpuSystemTimeAvg}
            "cpuSystemTimeSum" = ${CpuSystemTimeSum}
            "usedMemoryCount" = ${UsedMemoryCount}
            "usedMemoryMin" = ${UsedMemoryMin}
            "usedMemoryMax" = ${UsedMemoryMax}
            "usedMemoryAvg" = ${UsedMemoryAvg}
            "usedMemorySum" = ${UsedMemorySum}
            "freeMemoryCount" = ${FreeMemoryCount}
            "freeMemoryMin" = ${FreeMemoryMin}
            "freeMemoryMax" = ${FreeMemoryMax}
            "freeMemoryAvg" = ${FreeMemoryAvg}
            "freeMemorySum" = ${FreeMemorySum}
        }

        return $PSO
    }

}

