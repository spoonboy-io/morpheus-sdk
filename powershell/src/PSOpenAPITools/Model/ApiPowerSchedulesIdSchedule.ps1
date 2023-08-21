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
A name for the power schedule
.PARAMETER Description
A description for the power schedule
.PARAMETER ScheduleType
Type of schedule `power` on or `power off`
.PARAMETER ScheduleTimezone
Time Zone eg. America/New_York, Europe/Amsterdam, etc.
.PARAMETER Enabled
Is the power schedule enabled
.PARAMETER MondayOnTime
Monday Start time of the day in 24-hour format
.PARAMETER MondayOffTime
Monday Off time of the day in 24-hour format
.PARAMETER TuesdayOnTime
Tuesday Start time of the day in 24-hour format
.PARAMETER TuesdayOffTime
Tuesday Off time of the day in 24-hour format
.PARAMETER WednesdayOnTime
Wednesday Start time of the day in 24-hour format
.PARAMETER WednesdayOffTime
Wednesday Off time of the day in 24-hour format
.PARAMETER ThursdayOnTime
Thursday Start time of the day in 24-hour format
.PARAMETER ThursdayOffTime
Thursday Off time of the day in 24-hour format
.PARAMETER FridayOnTime
Friday Start time of the day in 24-hour format
.PARAMETER FridayOffTime
Friday Off time of the day in 24-hour format
.PARAMETER SaturdayOnTime
Saturday Start time of the day in 24-hour format
.PARAMETER SaturdayOffTime
Saturday Off time of the day in 24-hour format
.PARAMETER SundayOnTime
Sunday Start time of the day in 24-hour format
.PARAMETER SundayOffTime
Sunday Off time of the day in 24-hour format
.OUTPUTS

ApiPowerSchedulesIdSchedule<PSCustomObject>
#>

function Initialize-ApiPowerSchedulesIdSchedule {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("power", "power off")]
        [String]
        ${ScheduleType},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ScheduleTimezone} = "UTC",
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true,
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MondayOnTime} = "00:00",
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MondayOffTime} = "24:00",
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TuesdayOnTime} = "00:00",
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TuesdayOffTime} = "24:00",
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${WednesdayOnTime} = "00:00",
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${WednesdayOffTime} = "24:00",
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ThursdayOnTime} = "00:00",
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ThursdayOffTime} = "24:00",
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${FridayOnTime} = "00:00",
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${FridayOffTime} = "24:00",
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SaturdayOnTime} = "00:00",
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SaturdayOffTime} = "24:00",
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SundayOnTime} = "00:00",
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SundayOffTime} = "24:00"
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiPowerSchedulesIdSchedule' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "scheduleType" = ${ScheduleType}
            "scheduleTimezone" = ${ScheduleTimezone}
            "enabled" = ${Enabled}
            "mondayOnTime" = ${MondayOnTime}
            "mondayOffTime" = ${MondayOffTime}
            "tuesdayOnTime" = ${TuesdayOnTime}
            "tuesdayOffTime" = ${TuesdayOffTime}
            "wednesdayOnTime" = ${WednesdayOnTime}
            "wednesdayOffTime" = ${WednesdayOffTime}
            "thursdayOnTime" = ${ThursdayOnTime}
            "thursdayOffTime" = ${ThursdayOffTime}
            "fridayOnTime" = ${FridayOnTime}
            "fridayOffTime" = ${FridayOffTime}
            "saturdayOnTime" = ${SaturdayOnTime}
            "saturdayOffTime" = ${SaturdayOffTime}
            "sundayOnTime" = ${SundayOnTime}
            "sundayOffTime" = ${SundayOffTime}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiPowerSchedulesIdSchedule<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiPowerSchedulesIdSchedule<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiPowerSchedulesIdSchedule<PSCustomObject>
#>
function ConvertFrom-JsonToApiPowerSchedulesIdSchedule {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiPowerSchedulesIdSchedule' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiPowerSchedulesIdSchedule
        $AllProperties = ("name", "description", "scheduleType", "scheduleTimezone", "enabled", "mondayOnTime", "mondayOffTime", "tuesdayOnTime", "tuesdayOffTime", "wednesdayOnTime", "wednesdayOffTime", "thursdayOnTime", "thursdayOffTime", "fridayOnTime", "fridayOffTime", "saturdayOnTime", "saturdayOffTime", "sundayOnTime", "sundayOffTime")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "mondayOnTime"))) { #optional property not found
            $MondayOnTime = $null
        } else {
            $MondayOnTime = $JsonParameters.PSobject.Properties["mondayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "mondayOffTime"))) { #optional property not found
            $MondayOffTime = $null
        } else {
            $MondayOffTime = $JsonParameters.PSobject.Properties["mondayOffTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tuesdayOnTime"))) { #optional property not found
            $TuesdayOnTime = $null
        } else {
            $TuesdayOnTime = $JsonParameters.PSobject.Properties["tuesdayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tuesdayOffTime"))) { #optional property not found
            $TuesdayOffTime = $null
        } else {
            $TuesdayOffTime = $JsonParameters.PSobject.Properties["tuesdayOffTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "wednesdayOnTime"))) { #optional property not found
            $WednesdayOnTime = $null
        } else {
            $WednesdayOnTime = $JsonParameters.PSobject.Properties["wednesdayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "wednesdayOffTime"))) { #optional property not found
            $WednesdayOffTime = $null
        } else {
            $WednesdayOffTime = $JsonParameters.PSobject.Properties["wednesdayOffTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "thursdayOnTime"))) { #optional property not found
            $ThursdayOnTime = $null
        } else {
            $ThursdayOnTime = $JsonParameters.PSobject.Properties["thursdayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "thursdayOffTime"))) { #optional property not found
            $ThursdayOffTime = $null
        } else {
            $ThursdayOffTime = $JsonParameters.PSobject.Properties["thursdayOffTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "fridayOnTime"))) { #optional property not found
            $FridayOnTime = $null
        } else {
            $FridayOnTime = $JsonParameters.PSobject.Properties["fridayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "fridayOffTime"))) { #optional property not found
            $FridayOffTime = $null
        } else {
            $FridayOffTime = $JsonParameters.PSobject.Properties["fridayOffTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "saturdayOnTime"))) { #optional property not found
            $SaturdayOnTime = $null
        } else {
            $SaturdayOnTime = $JsonParameters.PSobject.Properties["saturdayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "saturdayOffTime"))) { #optional property not found
            $SaturdayOffTime = $null
        } else {
            $SaturdayOffTime = $JsonParameters.PSobject.Properties["saturdayOffTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sundayOnTime"))) { #optional property not found
            $SundayOnTime = $null
        } else {
            $SundayOnTime = $JsonParameters.PSobject.Properties["sundayOnTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "sundayOffTime"))) { #optional property not found
            $SundayOffTime = $null
        } else {
            $SundayOffTime = $JsonParameters.PSobject.Properties["sundayOffTime"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "scheduleType" = ${ScheduleType}
            "scheduleTimezone" = ${ScheduleTimezone}
            "enabled" = ${Enabled}
            "mondayOnTime" = ${MondayOnTime}
            "mondayOffTime" = ${MondayOffTime}
            "tuesdayOnTime" = ${TuesdayOnTime}
            "tuesdayOffTime" = ${TuesdayOffTime}
            "wednesdayOnTime" = ${WednesdayOnTime}
            "wednesdayOffTime" = ${WednesdayOffTime}
            "thursdayOnTime" = ${ThursdayOnTime}
            "thursdayOffTime" = ${ThursdayOffTime}
            "fridayOnTime" = ${FridayOnTime}
            "fridayOffTime" = ${FridayOffTime}
            "saturdayOnTime" = ${SaturdayOnTime}
            "saturdayOffTime" = ${SaturdayOffTime}
            "sundayOnTime" = ${SundayOnTime}
            "sundayOffTime" = ${SundayOffTime}
        }

        return $PSO
    }

}

