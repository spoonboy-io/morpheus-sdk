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

.PARAMETER Environment
No description available.
.OUTPUTS

AddEnvironments200ResponseAllOf<PSCustomObject>
#>

function Initialize-AddEnvironments200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Environment}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddEnvironments200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "environment" = ${Environment}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddEnvironments200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddEnvironments200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddEnvironments200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddEnvironments200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddEnvironments200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddEnvironments200ResponseAllOf
        $AllProperties = ("environment")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "environment"))) { #optional property not found
            $Environment = $null
        } else {
            $Environment = $JsonParameters.PSobject.Properties["environment"].value
        }

        $PSO = [PSCustomObject]@{
            "environment" = ${Environment}
        }

        return $PSO
    }

}

