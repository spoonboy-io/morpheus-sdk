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

.PARAMETER InstanceTypeLayout
No description available.
.OUTPUTS

AddLayoutRequest<PSCustomObject>
#>

function Initialize-AddLayoutRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${InstanceTypeLayout}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddLayoutRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "instanceTypeLayout" = ${InstanceTypeLayout}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddLayoutRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddLayoutRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddLayoutRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddLayoutRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddLayoutRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddLayoutRequest
        $AllProperties = ("instanceTypeLayout")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "instanceTypeLayout"))) { #optional property not found
            $InstanceTypeLayout = $null
        } else {
            $InstanceTypeLayout = $JsonParameters.PSobject.Properties["instanceTypeLayout"].value
        }

        $PSO = [PSCustomObject]@{
            "instanceTypeLayout" = ${InstanceTypeLayout}
        }

        return $PSO
    }

}

