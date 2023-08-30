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
A name for the execute schedule
.PARAMETER Description
A description for the execute schedule
.PARAMETER ScheduleType
Type of schedule
.PARAMETER ScheduleTimezone
Time Zone eg. America/New_York, Europe/Amsterdam, etc.
.PARAMETER Cron
Cron Expression. The default is daily at midnight
.PARAMETER Enabled
Is enabled
.OUTPUTS

AddExecuteSchedulesRequestSchedule<PSCustomObject>
#>

function Initialize-AddExecuteSchedulesRequestSchedule {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("execute")]
        [String]
        ${ScheduleType},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ScheduleTimezone} = "UTC",
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Cron} = "0 0 * * *",
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddExecuteSchedulesRequestSchedule' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "scheduleType" = ${ScheduleType}
            "scheduleTimezone" = ${ScheduleTimezone}
            "cron" = ${Cron}
            "enabled" = ${Enabled}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddExecuteSchedulesRequestSchedule<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddExecuteSchedulesRequestSchedule<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddExecuteSchedulesRequestSchedule<PSCustomObject>
#>
function ConvertFrom-JsonToAddExecuteSchedulesRequestSchedule {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddExecuteSchedulesRequestSchedule' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddExecuteSchedulesRequestSchedule
        $AllProperties = ("name", "description", "scheduleType", "scheduleTimezone", "cron", "enabled")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'name' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property 'name' missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "scheduleType" = ${ScheduleType}
            "scheduleTimezone" = ${ScheduleTimezone}
            "cron" = ${Cron}
            "enabled" = ${Enabled}
        }

        return $PSO
    }

}

