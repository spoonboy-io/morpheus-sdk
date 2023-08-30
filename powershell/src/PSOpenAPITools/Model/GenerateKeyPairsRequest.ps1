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

.PARAMETER KeyPair
No description available.
.OUTPUTS

GenerateKeyPairsRequest<PSCustomObject>
#>

function Initialize-GenerateKeyPairsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${KeyPair}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GenerateKeyPairsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $KeyPair) {
            throw "invalid value for 'KeyPair', 'KeyPair' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "keyPair" = ${KeyPair}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GenerateKeyPairsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to GenerateKeyPairsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GenerateKeyPairsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToGenerateKeyPairsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GenerateKeyPairsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GenerateKeyPairsRequest
        $AllProperties = ("keyPair")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'keyPair' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "keyPair"))) {
            throw "Error! JSON cannot be serialized due to the required property 'keyPair' missing."
        } else {
            $KeyPair = $JsonParameters.PSobject.Properties["keyPair"].value
        }

        $PSO = [PSCustomObject]@{
            "keyPair" = ${KeyPair}
        }

        return $PSO
    }

}
