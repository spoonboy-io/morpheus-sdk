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

.PARAMETER App
No description available.
.OUTPUTS

GetApp200Response<PSCustomObject>
#>

function Initialize-GetApp200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${App}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetApp200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "app" = ${App}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetApp200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetApp200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetApp200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetApp200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetApp200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetApp200Response
        $AllProperties = ("app")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "app"))) { #optional property not found
            $App = $null
        } else {
            $App = $JsonParameters.PSobject.Properties["app"].value
        }

        $PSO = [PSCustomObject]@{
            "app" = ${App}
        }

        return $PSO
    }

}

