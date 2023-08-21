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

.PARAMETER Success
No description available.
.PARAMETER ThreadList
No description available.
.PARAMETER BusyThreads
No description available.
.PARAMETER BlockedThreads
No description available.
.PARAMETER RunningThreads
No description available.
.PARAMETER TotalCpuTime
No description available.
.PARAMETER TotalThreads
No description available.
.PARAMETER RunningWebThreads
No description available.
.PARAMETER Status
No description available.
.OUTPUTS

HealthThreads<PSCustomObject>
#>

function Initialize-HealthThreads {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ThreadList},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${BusyThreads},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${BlockedThreads},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${RunningThreads},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TotalCpuTime},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TotalThreads},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RunningWebThreads},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HealthThreads' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "success" = ${Success}
            "threadList" = ${ThreadList}
            "busyThreads" = ${BusyThreads}
            "blockedThreads" = ${BlockedThreads}
            "runningThreads" = ${RunningThreads}
            "totalCpuTime" = ${TotalCpuTime}
            "totalThreads" = ${TotalThreads}
            "runningWebThreads" = ${RunningWebThreads}
            "status" = ${Status}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HealthThreads<PSCustomObject>

.DESCRIPTION

Convert from JSON to HealthThreads<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HealthThreads<PSCustomObject>
#>
function ConvertFrom-JsonToHealthThreads {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HealthThreads' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HealthThreads
        $AllProperties = ("success", "threadList", "busyThreads", "blockedThreads", "runningThreads", "totalCpuTime", "totalThreads", "runningWebThreads", "status")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "threadList"))) { #optional property not found
            $ThreadList = $null
        } else {
            $ThreadList = $JsonParameters.PSobject.Properties["threadList"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "busyThreads"))) { #optional property not found
            $BusyThreads = $null
        } else {
            $BusyThreads = $JsonParameters.PSobject.Properties["busyThreads"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blockedThreads"))) { #optional property not found
            $BlockedThreads = $null
        } else {
            $BlockedThreads = $JsonParameters.PSobject.Properties["blockedThreads"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "runningThreads"))) { #optional property not found
            $RunningThreads = $null
        } else {
            $RunningThreads = $JsonParameters.PSobject.Properties["runningThreads"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "totalCpuTime"))) { #optional property not found
            $TotalCpuTime = $null
        } else {
            $TotalCpuTime = $JsonParameters.PSobject.Properties["totalCpuTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "totalThreads"))) { #optional property not found
            $TotalThreads = $null
        } else {
            $TotalThreads = $JsonParameters.PSobject.Properties["totalThreads"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "runningWebThreads"))) { #optional property not found
            $RunningWebThreads = $null
        } else {
            $RunningWebThreads = $JsonParameters.PSobject.Properties["runningWebThreads"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        $PSO = [PSCustomObject]@{
            "success" = ${Success}
            "threadList" = ${ThreadList}
            "busyThreads" = ${BusyThreads}
            "blockedThreads" = ${BlockedThreads}
            "runningThreads" = ${RunningThreads}
            "totalCpuTime" = ${TotalCpuTime}
            "totalThreads" = ${TotalThreads}
            "runningWebThreads" = ${RunningWebThreads}
            "status" = ${Status}
        }

        return $PSO
    }

}
