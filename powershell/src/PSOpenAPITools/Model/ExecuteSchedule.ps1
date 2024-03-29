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

.PARAMETER Id
No description available.
.PARAMETER Name
No description available.
.PARAMETER Description
No description available.
.PARAMETER Visibility
No description available.
.PARAMETER Enabled
No description available.
.PARAMETER ScheduleType
No description available.
.PARAMETER ScheduleTimezone
No description available.
.PARAMETER Cron
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.OUTPUTS

ExecuteSchedule<PSCustomObject>
#>

function Initialize-ExecuteSchedule {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Visibility},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ScheduleType},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ScheduleTimezone},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Cron},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ExecuteSchedule' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "visibility" = ${Visibility}
            "enabled" = ${Enabled}
            "scheduleType" = ${ScheduleType}
            "scheduleTimezone" = ${ScheduleTimezone}
            "cron" = ${Cron}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ExecuteSchedule<PSCustomObject>

.DESCRIPTION

Convert from JSON to ExecuteSchedule<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ExecuteSchedule<PSCustomObject>
#>
function ConvertFrom-JsonToExecuteSchedule {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ExecuteSchedule' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ExecuteSchedule
        $AllProperties = ("id", "name", "description", "visibility", "enabled", "scheduleType", "scheduleTimezone", "cron", "dateCreated", "lastUpdated")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scheduleType"))) { #optional property not found
            $ScheduleType = $null
        } else {
            $ScheduleType = $JsonParameters.PSobject.Properties["scheduleType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scheduleTimezone"))) { #optional property not found
            $ScheduleTimezone = $null
        } else {
            $ScheduleTimezone = $JsonParameters.PSobject.Properties["scheduleTimezone"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cron"))) { #optional property not found
            $Cron = $null
        } else {
            $Cron = $JsonParameters.PSobject.Properties["cron"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "description" = ${Description}
            "visibility" = ${Visibility}
            "enabled" = ${Enabled}
            "scheduleType" = ${ScheduleType}
            "scheduleTimezone" = ${ScheduleTimezone}
            "cron" = ${Cron}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }

        return $PSO
    }

}

