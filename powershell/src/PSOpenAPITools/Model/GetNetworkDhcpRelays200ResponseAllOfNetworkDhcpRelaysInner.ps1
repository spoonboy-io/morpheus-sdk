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
.PARAMETER DateCreated
No description available.
.PARAMETER ProviderId
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER Name
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER ServerIpAddresses
No description available.
.PARAMETER Owner
No description available.
.PARAMETER NetworkServer
No description available.
.OUTPUTS

GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner<PSCustomObject>
#>

function Initialize-GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProviderId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${ServerIpAddresses},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Owner},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${NetworkServer}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "dateCreated" = ${DateCreated}
            "providerId" = ${ProviderId}
            "lastUpdated" = ${LastUpdated}
            "name" = ${Name}
            "externalId" = ${ExternalId}
            "serverIpAddresses" = ${ServerIpAddresses}
            "owner" = ${Owner}
            "networkServer" = ${NetworkServer}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner
        $AllProperties = ("id", "dateCreated", "providerId", "lastUpdated", "name", "externalId", "serverIpAddresses", "owner", "networkServer")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "providerId"))) { #optional property not found
            $ProviderId = $null
        } else {
            $ProviderId = $JsonParameters.PSobject.Properties["providerId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serverIpAddresses"))) { #optional property not found
            $ServerIpAddresses = $null
        } else {
            $ServerIpAddresses = $JsonParameters.PSobject.Properties["serverIpAddresses"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "owner"))) { #optional property not found
            $Owner = $null
        } else {
            $Owner = $JsonParameters.PSobject.Properties["owner"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkServer"))) { #optional property not found
            $NetworkServer = $null
        } else {
            $NetworkServer = $JsonParameters.PSobject.Properties["networkServer"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "dateCreated" = ${DateCreated}
            "providerId" = ${ProviderId}
            "lastUpdated" = ${LastUpdated}
            "name" = ${Name}
            "externalId" = ${ExternalId}
            "serverIpAddresses" = ${ServerIpAddresses}
            "owner" = ${Owner}
            "networkServer" = ${NetworkServer}
        }

        return $PSO
    }

}

