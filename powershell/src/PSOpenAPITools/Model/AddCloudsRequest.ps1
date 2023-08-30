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

.PARAMETER Zone
No description available.
.OUTPUTS

AddCloudsRequest<PSCustomObject>
#>

function Initialize-AddCloudsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Zone}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddCloudsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Zone) {
            throw "invalid value for 'Zone', 'Zone' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "zone" = ${Zone}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddCloudsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddCloudsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddCloudsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddCloudsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddCloudsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddCloudsRequest
        $AllProperties = ("zone")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'zone' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "zone"))) {
            throw "Error! JSON cannot be serialized due to the required property 'zone' missing."
        } else {
            $Zone = $JsonParameters.PSobject.Properties["zone"].value
        }

        $PSO = [PSCustomObject]@{
            "zone" = ${Zone}
        }

        return $PSO
    }

}

