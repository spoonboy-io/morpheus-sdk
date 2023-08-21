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
A name for the Job
.PARAMETER Labels
Array of label strings, can be used for filtering.
.PARAMETER Enabled
Use this to set enabled state
.PARAMETER Task
No description available.
.PARAMETER Workflow
No description available.
.PARAMETER ScanPath
Scan Checklist. Only applies to type scap-package.
.PARAMETER SecurityProfile
Security Profile. Only applies to type scap-package.
.PARAMETER TargetType
Target type where job will execute
.PARAMETER Targets
No description available.
.PARAMETER InstanceLabel
Instance Label. Only applicable if `targetType` is `instance-label`.
.PARAMETER ServerLabel
Server Label. Only applicable if `targetType` is `server-label`.
.PARAMETER ScheduleMode
No description available.
.PARAMETER CustomOptions
Map of options to be used as values in the workflow tasks. These correspond to option types.
.PARAMETER CustomConfig
Job custom configuration (String in JSON format)
.PARAMETER DateTime
Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is 'dateTime'.
.PARAMETER Run
If true, executes job
.OUTPUTS

ApiJobsIdJob<PSCustomObject>
#>

function Initialize-ApiJobsIdJob {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Enabled} = $true,
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Task},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Workflow},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ScanPath},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SecurityProfile},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("appliance", "instance", "instance-label", "server", "server-label")]
        [String]
        ${TargetType},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Targets},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InstanceLabel},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServerLabel},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ScheduleMode},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CustomOptions},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CustomConfig},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateTime},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Run}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiJobsIdJob' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "labels" = ${Labels}
            "enabled" = ${Enabled}
            "task" = ${Task}
            "workflow" = ${Workflow}
            "scanPath" = ${ScanPath}
            "securityProfile" = ${SecurityProfile}
            "targetType" = ${TargetType}
            "targets" = ${Targets}
            "instanceLabel" = ${InstanceLabel}
            "serverLabel" = ${ServerLabel}
            "scheduleMode" = ${ScheduleMode}
            "customOptions" = ${CustomOptions}
            "customConfig" = ${CustomConfig}
            "dateTime" = ${DateTime}
            "run" = ${Run}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiJobsIdJob<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiJobsIdJob<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiJobsIdJob<PSCustomObject>
#>
function ConvertFrom-JsonToApiJobsIdJob {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiJobsIdJob' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiJobsIdJob
        $AllProperties = ("name", "labels", "enabled", "task", "workflow", "scanPath", "securityProfile", "targetType", "targets", "instanceLabel", "serverLabel", "scheduleMode", "customOptions", "customConfig", "dateTime", "run")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "enabled"))) { #optional property not found
            $Enabled = $null
        } else {
            $Enabled = $JsonParameters.PSobject.Properties["enabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "task"))) { #optional property not found
            $Task = $null
        } else {
            $Task = $JsonParameters.PSobject.Properties["task"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "workflow"))) { #optional property not found
            $Workflow = $null
        } else {
            $Workflow = $JsonParameters.PSobject.Properties["workflow"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scanPath"))) { #optional property not found
            $ScanPath = $null
        } else {
            $ScanPath = $JsonParameters.PSobject.Properties["scanPath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "securityProfile"))) { #optional property not found
            $SecurityProfile = $null
        } else {
            $SecurityProfile = $JsonParameters.PSobject.Properties["securityProfile"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "targetType"))) { #optional property not found
            $TargetType = $null
        } else {
            $TargetType = $JsonParameters.PSobject.Properties["targetType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "targets"))) { #optional property not found
            $Targets = $null
        } else {
            $Targets = $JsonParameters.PSobject.Properties["targets"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceLabel"))) { #optional property not found
            $InstanceLabel = $null
        } else {
            $InstanceLabel = $JsonParameters.PSobject.Properties["instanceLabel"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverLabel"))) { #optional property not found
            $ServerLabel = $null
        } else {
            $ServerLabel = $JsonParameters.PSobject.Properties["serverLabel"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "scheduleMode"))) { #optional property not found
            $ScheduleMode = $null
        } else {
            $ScheduleMode = $JsonParameters.PSobject.Properties["scheduleMode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "customOptions"))) { #optional property not found
            $CustomOptions = $null
        } else {
            $CustomOptions = $JsonParameters.PSobject.Properties["customOptions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "customConfig"))) { #optional property not found
            $CustomConfig = $null
        } else {
            $CustomConfig = $JsonParameters.PSobject.Properties["customConfig"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateTime"))) { #optional property not found
            $DateTime = $null
        } else {
            $DateTime = $JsonParameters.PSobject.Properties["dateTime"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "run"))) { #optional property not found
            $Run = $null
        } else {
            $Run = $JsonParameters.PSobject.Properties["run"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "labels" = ${Labels}
            "enabled" = ${Enabled}
            "task" = ${Task}
            "workflow" = ${Workflow}
            "scanPath" = ${ScanPath}
            "securityProfile" = ${SecurityProfile}
            "targetType" = ${TargetType}
            "targets" = ${Targets}
            "instanceLabel" = ${InstanceLabel}
            "serverLabel" = ${ServerLabel}
            "scheduleMode" = ${ScheduleMode}
            "customOptions" = ${CustomOptions}
            "customConfig" = ${CustomConfig}
            "dateTime" = ${DateTime}
            "run" = ${Run}
        }

        return $PSO
    }

}
