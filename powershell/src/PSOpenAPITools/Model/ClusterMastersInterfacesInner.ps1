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
.PARAMETER RefType
No description available.
.PARAMETER RefId
No description available.
.PARAMETER Name
No description available.
.PARAMETER InternalId
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER UniqueId
No description available.
.PARAMETER PublicIpAddress
No description available.
.PARAMETER PublicIpv6Address
No description available.
.PARAMETER IpAddress
No description available.
.PARAMETER Ipv6Address
No description available.
.PARAMETER IpSubnet
No description available.
.PARAMETER Ipv6Subnet
No description available.
.PARAMETER Description
No description available.
.PARAMETER Dhcp
No description available.
.PARAMETER Active
No description available.
.PARAMETER PoolAssigned
No description available.
.PARAMETER PrimaryInterface
No description available.
.PARAMETER Network
No description available.
.PARAMETER Subnet
No description available.
.PARAMETER NetworkGroup
No description available.
.PARAMETER NetworkPosition
No description available.
.PARAMETER NetworkPool
No description available.
.PARAMETER NetworkDomain
No description available.
.PARAMETER Type
No description available.
.PARAMETER IpMode
No description available.
.PARAMETER MacAddress
No description available.
.OUTPUTS

ClusterMastersInterfacesInner<PSCustomObject>
#>

function Initialize-ClusterMastersInterfacesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefType},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalId},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UniqueId},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PublicIpAddress},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PublicIpv6Address},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IpAddress},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Ipv6Address},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IpSubnet},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Ipv6Subnet},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Dhcp},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Active},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${PoolAssigned},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${PrimaryInterface},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Network},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Subnet},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NetworkGroup},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NetworkPosition},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NetworkPool},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NetworkDomain},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IpMode},
        [Parameter(Position = 26, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MacAddress}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ClusterMastersInterfacesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "refType" = ${RefType}
            "refId" = ${RefId}
            "name" = ${Name}
            "internalId" = ${InternalId}
            "externalId" = ${ExternalId}
            "uniqueId" = ${UniqueId}
            "publicIpAddress" = ${PublicIpAddress}
            "publicIpv6Address" = ${PublicIpv6Address}
            "ipAddress" = ${IpAddress}
            "ipv6Address" = ${Ipv6Address}
            "ipSubnet" = ${IpSubnet}
            "ipv6Subnet" = ${Ipv6Subnet}
            "description" = ${Description}
            "dhcp" = ${Dhcp}
            "active" = ${Active}
            "poolAssigned" = ${PoolAssigned}
            "primaryInterface" = ${PrimaryInterface}
            "network" = ${Network}
            "subnet" = ${Subnet}
            "networkGroup" = ${NetworkGroup}
            "networkPosition" = ${NetworkPosition}
            "networkPool" = ${NetworkPool}
            "networkDomain" = ${NetworkDomain}
            "type" = ${Type}
            "ipMode" = ${IpMode}
            "macAddress" = ${MacAddress}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ClusterMastersInterfacesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to ClusterMastersInterfacesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ClusterMastersInterfacesInner<PSCustomObject>
#>
function ConvertFrom-JsonToClusterMastersInterfacesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ClusterMastersInterfacesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ClusterMastersInterfacesInner
        $AllProperties = ("id", "refType", "refId", "name", "internalId", "externalId", "uniqueId", "publicIpAddress", "publicIpv6Address", "ipAddress", "ipv6Address", "ipSubnet", "ipv6Subnet", "description", "dhcp", "active", "poolAssigned", "primaryInterface", "network", "subnet", "networkGroup", "networkPosition", "networkPool", "networkDomain", "type", "ipMode", "macAddress")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "refType"))) { #optional property not found
            $RefType = $null
        } else {
            $RefType = $JsonParameters.PSobject.Properties["refType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "refId"))) { #optional property not found
            $RefId = $null
        } else {
            $RefId = $JsonParameters.PSobject.Properties["refId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalId"))) { #optional property not found
            $InternalId = $null
        } else {
            $InternalId = $JsonParameters.PSobject.Properties["internalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "uniqueId"))) { #optional property not found
            $UniqueId = $null
        } else {
            $UniqueId = $JsonParameters.PSobject.Properties["uniqueId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "publicIpAddress"))) { #optional property not found
            $PublicIpAddress = $null
        } else {
            $PublicIpAddress = $JsonParameters.PSobject.Properties["publicIpAddress"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "publicIpv6Address"))) { #optional property not found
            $PublicIpv6Address = $null
        } else {
            $PublicIpv6Address = $JsonParameters.PSobject.Properties["publicIpv6Address"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipAddress"))) { #optional property not found
            $IpAddress = $null
        } else {
            $IpAddress = $JsonParameters.PSobject.Properties["ipAddress"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipv6Address"))) { #optional property not found
            $Ipv6Address = $null
        } else {
            $Ipv6Address = $JsonParameters.PSobject.Properties["ipv6Address"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipSubnet"))) { #optional property not found
            $IpSubnet = $null
        } else {
            $IpSubnet = $JsonParameters.PSobject.Properties["ipSubnet"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipv6Subnet"))) { #optional property not found
            $Ipv6Subnet = $null
        } else {
            $Ipv6Subnet = $JsonParameters.PSobject.Properties["ipv6Subnet"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dhcp"))) { #optional property not found
            $Dhcp = $null
        } else {
            $Dhcp = $JsonParameters.PSobject.Properties["dhcp"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "active"))) { #optional property not found
            $Active = $null
        } else {
            $Active = $JsonParameters.PSobject.Properties["active"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "poolAssigned"))) { #optional property not found
            $PoolAssigned = $null
        } else {
            $PoolAssigned = $JsonParameters.PSobject.Properties["poolAssigned"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "primaryInterface"))) { #optional property not found
            $PrimaryInterface = $null
        } else {
            $PrimaryInterface = $JsonParameters.PSobject.Properties["primaryInterface"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "network"))) { #optional property not found
            $Network = $null
        } else {
            $Network = $JsonParameters.PSobject.Properties["network"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subnet"))) { #optional property not found
            $Subnet = $null
        } else {
            $Subnet = $JsonParameters.PSobject.Properties["subnet"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkGroup"))) { #optional property not found
            $NetworkGroup = $null
        } else {
            $NetworkGroup = $JsonParameters.PSobject.Properties["networkGroup"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkPosition"))) { #optional property not found
            $NetworkPosition = $null
        } else {
            $NetworkPosition = $JsonParameters.PSobject.Properties["networkPosition"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkPool"))) { #optional property not found
            $NetworkPool = $null
        } else {
            $NetworkPool = $JsonParameters.PSobject.Properties["networkPool"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkDomain"))) { #optional property not found
            $NetworkDomain = $null
        } else {
            $NetworkDomain = $JsonParameters.PSobject.Properties["networkDomain"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipMode"))) { #optional property not found
            $IpMode = $null
        } else {
            $IpMode = $JsonParameters.PSobject.Properties["ipMode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "macAddress"))) { #optional property not found
            $MacAddress = $null
        } else {
            $MacAddress = $JsonParameters.PSobject.Properties["macAddress"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "refType" = ${RefType}
            "refId" = ${RefId}
            "name" = ${Name}
            "internalId" = ${InternalId}
            "externalId" = ${ExternalId}
            "uniqueId" = ${UniqueId}
            "publicIpAddress" = ${PublicIpAddress}
            "publicIpv6Address" = ${PublicIpv6Address}
            "ipAddress" = ${IpAddress}
            "ipv6Address" = ${Ipv6Address}
            "ipSubnet" = ${IpSubnet}
            "ipv6Subnet" = ${Ipv6Subnet}
            "description" = ${Description}
            "dhcp" = ${Dhcp}
            "active" = ${Active}
            "poolAssigned" = ${PoolAssigned}
            "primaryInterface" = ${PrimaryInterface}
            "network" = ${Network}
            "subnet" = ${Subnet}
            "networkGroup" = ${NetworkGroup}
            "networkPosition" = ${NetworkPosition}
            "networkPool" = ${NetworkPool}
            "networkDomain" = ${NetworkDomain}
            "type" = ${Type}
            "ipMode" = ${IpMode}
            "macAddress" = ${MacAddress}
        }

        return $PSO
    }

}

