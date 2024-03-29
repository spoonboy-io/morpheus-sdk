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

.PARAMETER Stats
No description available.
.OUTPUTS

GetGuidanceStats200Response<PSCustomObject>
#>

function Initialize-GetGuidanceStats200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Stats}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetGuidanceStats200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "stats" = ${Stats}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetGuidanceStats200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetGuidanceStats200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetGuidanceStats200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetGuidanceStats200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetGuidanceStats200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetGuidanceStats200Response
        $AllProperties = ("stats")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "stats"))) { #optional property not found
            $Stats = $null
        } else {
            $Stats = $JsonParameters.PSobject.Properties["stats"].value
        }

        $PSO = [PSCustomObject]@{
            "stats" = ${Stats}
        }

        return $PSO
    }

}

