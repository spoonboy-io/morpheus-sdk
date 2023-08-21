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
.PARAMETER Method
No description available.
.PARAMETER Notify
No description available.
.PARAMETER Close
No description available.
.OUTPUTS

ApiMonitoringAlertsAlertContacts<PSCustomObject>
#>

function Initialize-ApiMonitoringAlertsAlertContacts {
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
        ${Method},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Notify},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Close}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiMonitoringAlertsAlertContacts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "method" = ${Method}
            "notify" = ${Notify}
            "close" = ${Close}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiMonitoringAlertsAlertContacts<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiMonitoringAlertsAlertContacts<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiMonitoringAlertsAlertContacts<PSCustomObject>
#>
function ConvertFrom-JsonToApiMonitoringAlertsAlertContacts {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiMonitoringAlertsAlertContacts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiMonitoringAlertsAlertContacts
        $AllProperties = ("id", "name", "method", "notify", "close")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "method"))) { #optional property not found
            $Method = $null
        } else {
            $Method = $JsonParameters.PSobject.Properties["method"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "notify"))) { #optional property not found
            $Notify = $null
        } else {
            $Notify = $JsonParameters.PSobject.Properties["notify"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "close"))) { #optional property not found
            $Close = $null
        } else {
            $Close = $JsonParameters.PSobject.Properties["close"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "method" = ${Method}
            "notify" = ${Notify}
            "close" = ${Close}
        }

        return $PSO
    }

}

