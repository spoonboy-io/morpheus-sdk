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
.OUTPUTS

GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner<PSCustomObject>
#>

function Initialize-GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner
        $AllProperties = ("id", "name")
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

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
        }

        return $PSO
    }

}

