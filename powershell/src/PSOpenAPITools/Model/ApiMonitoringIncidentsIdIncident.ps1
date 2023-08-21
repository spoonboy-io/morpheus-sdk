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

Payload for update an incident

.PARAMETER Resolution
Description of the resolution to this incident
.PARAMETER Comment
Comment on this incident, updates summary field
.PARAMETER Status
Set status
.PARAMETER Severity
Set severity
.PARAMETER Name
Set display name
.PARAMETER StartDate
Set start time
.PARAMETER EndDate
Set start time
.PARAMETER InUptime
Set 'In Availability'
.OUTPUTS

ApiMonitoringIncidentsIdIncident<PSCustomObject>
#>

function Initialize-ApiMonitoringIncidentsIdIncident {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Resolution},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Comment},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("open", "closed")]
        [String]
        ${Status},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("info", "warning", "critical")]
        [String]
        ${Severity},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${InUptime}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiMonitoringIncidentsIdIncident' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "resolution" = ${Resolution}
            "comment" = ${Comment}
            "status" = ${Status}
            "severity" = ${Severity}
            "name" = ${Name}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "inUptime" = ${InUptime}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiMonitoringIncidentsIdIncident<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiMonitoringIncidentsIdIncident<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiMonitoringIncidentsIdIncident<PSCustomObject>
#>
function ConvertFrom-JsonToApiMonitoringIncidentsIdIncident {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiMonitoringIncidentsIdIncident' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiMonitoringIncidentsIdIncident
        $AllProperties = ("resolution", "comment", "status", "severity", "name", "startDate", "endDate", "inUptime")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "resolution"))) { #optional property not found
            $Resolution = $null
        } else {
            $Resolution = $JsonParameters.PSobject.Properties["resolution"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "comment"))) { #optional property not found
            $Comment = $null
        } else {
            $Comment = $JsonParameters.PSobject.Properties["comment"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "severity"))) { #optional property not found
            $Severity = $null
        } else {
            $Severity = $JsonParameters.PSobject.Properties["severity"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "startDate"))) { #optional property not found
            $StartDate = $null
        } else {
            $StartDate = $JsonParameters.PSobject.Properties["startDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "endDate"))) { #optional property not found
            $EndDate = $null
        } else {
            $EndDate = $JsonParameters.PSobject.Properties["endDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "inUptime"))) { #optional property not found
            $InUptime = $null
        } else {
            $InUptime = $JsonParameters.PSobject.Properties["inUptime"].value
        }

        $PSO = [PSCustomObject]@{
            "resolution" = ${Resolution}
            "comment" = ${Comment}
            "status" = ${Status}
            "severity" = ${Severity}
            "name" = ${Name}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "inUptime" = ${InUptime}
        }

        return $PSO
    }

}
