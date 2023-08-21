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

.PARAMETER Name
A name for the backup job
.PARAMETER Code
A code for the backup job
.PARAMETER ScheduleId
Execute Schedule ID to use for the backup job
.PARAMETER RetentionCount
Retention Count
.OUTPUTS

ApiBackupsJobsIdJob<PSCustomObject>
#>

function Initialize-ApiBackupsJobsIdJob {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ScheduleId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RetentionCount}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiBackupsJobsIdJob' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "scheduleId" = ${ScheduleId}
            "retentionCount" = ${RetentionCount}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiBackupsJobsIdJob<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiBackupsJobsIdJob<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiBackupsJobsIdJob<PSCustomObject>
#>
function ConvertFrom-JsonToApiBackupsJobsIdJob {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiBackupsJobsIdJob' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiBackupsJobsIdJob
        $AllProperties = ("name", "code", "scheduleId", "retentionCount")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scheduleId"))) { #optional property not found
            $ScheduleId = $null
        } else {
            $ScheduleId = $JsonParameters.PSobject.Properties["scheduleId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "retentionCount"))) { #optional property not found
            $RetentionCount = $null
        } else {
            $RetentionCount = $JsonParameters.PSobject.Properties["retentionCount"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "code" = ${Code}
            "scheduleId" = ${ScheduleId}
            "retentionCount" = ${RetentionCount}
        }

        return $PSO
    }

}

