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

.PARAMETER Specs
Array of Kubernetes specs in Morpheus
.OUTPUTS

BlueprintKubernetesCreateConfig<PSCustomObject>
#>

function Initialize-BlueprintKubernetesCreateConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Specs}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintKubernetesCreateConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "specs" = ${Specs}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintKubernetesCreateConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintKubernetesCreateConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintKubernetesCreateConfig<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintKubernetesCreateConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintKubernetesCreateConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintKubernetesCreateConfig
        $AllProperties = ("specs")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "specs"))) { #optional property not found
            $Specs = $null
        } else {
            $Specs = $JsonParameters.PSobject.Properties["specs"].value
        }

        $PSO = [PSCustomObject]@{
            "specs" = ${Specs}
        }

        return $PSO
    }

}

