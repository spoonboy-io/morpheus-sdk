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

.PARAMETER Meta
No description available.
.OUTPUTS

Meta<PSCustomObject>
#>

function Initialize-Meta {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Meta}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => Meta' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "meta" = ${Meta}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to Meta<PSCustomObject>

.DESCRIPTION

Convert from JSON to Meta<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

Meta<PSCustomObject>
#>
function ConvertFrom-JsonToMeta {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => Meta' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in Meta
        $AllProperties = ("meta")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "meta"))) { #optional property not found
            $Meta = $null
        } else {
            $Meta = $JsonParameters.PSobject.Properties["meta"].value
        }

        $PSO = [PSCustomObject]@{
            "meta" = ${Meta}
        }

        return $PSO
    }

}

