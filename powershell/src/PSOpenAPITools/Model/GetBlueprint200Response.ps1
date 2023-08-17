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

.PARAMETER Blueprint
No description available.
.OUTPUTS

GetBlueprint200Response<PSCustomObject>
#>

function Initialize-GetBlueprint200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Blueprint}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetBlueprint200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "blueprint" = ${Blueprint}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetBlueprint200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetBlueprint200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetBlueprint200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetBlueprint200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetBlueprint200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetBlueprint200Response
        $AllProperties = ("blueprint")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "blueprint"))) { #optional property not found
            $Blueprint = $null
        } else {
            $Blueprint = $JsonParameters.PSobject.Properties["blueprint"].value
        }

        $PSO = [PSCustomObject]@{
            "blueprint" = ${Blueprint}
        }

        return $PSO
    }

}

