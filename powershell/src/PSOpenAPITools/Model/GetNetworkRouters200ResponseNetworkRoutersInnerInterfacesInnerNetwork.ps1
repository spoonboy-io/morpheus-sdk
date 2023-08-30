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

.PARAMETER Id
No description available.
.PARAMETER Name
No description available.
.PARAMETER Code
No description available.
.OUTPUTS

GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork<PSCustomObject>
#>

function Initialize-GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Code}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "code" = ${Code}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork
        $AllProperties = ("id", "name", "code")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) { #optional property not found
            $Code = $null
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "code" = ${Code}
        }

        return $PSO
    }

}

