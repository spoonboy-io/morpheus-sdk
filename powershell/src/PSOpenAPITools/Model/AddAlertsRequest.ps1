#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER Alert
No description available.
.OUTPUTS

AddAlertsRequest<PSCustomObject>
#>

function Initialize-AddAlertsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Alert}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddAlertsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Alert) {
            throw "invalid value for 'Alert', 'Alert' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "alert" = ${Alert}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddAlertsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddAlertsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddAlertsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddAlertsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddAlertsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddAlertsRequest
        $AllProperties = ("alert")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'alert' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "alert"))) {
            throw "Error! JSON cannot be serialized due to the required property 'alert' missing."
        } else {
            $Alert = $JsonParameters.PSobject.Properties["alert"].value
        }

        $PSO = [PSCustomObject]@{
            "alert" = ${Alert}
        }

        return $PSO
    }

}

