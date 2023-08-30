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
.PARAMETER NetworkPoolId
No description available.
.PARAMETER IpType
No description available.
.PARAMETER IpAddress
No description available.
.PARAMETER GatewayAddress
No description available.
.PARAMETER SubnetMask
No description available.
.PARAMETER DnsServer
No description available.
.PARAMETER InterfaceName
No description available.
.PARAMETER Description
No description available.
.PARAMETER Active
No description available.
.PARAMETER StaticIp
No description available.
.PARAMETER Fqdn
No description available.
.PARAMETER DomainName
No description available.
.PARAMETER Hostname
No description available.
.PARAMETER InternalId
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER PtrId
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER StartDate
No description available.
.PARAMETER EndDate
No description available.
.PARAMETER RefType
No description available.
.PARAMETER RefId
No description available.
.PARAMETER SubRefId
No description available.
.PARAMETER NetworkDomain
No description available.
.PARAMETER CreatedBy
No description available.
.OUTPUTS

GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner<PSCustomObject>
#>

function Initialize-GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${NetworkPoolId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IpType},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IpAddress},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GatewayAddress},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SubnetMask},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DnsServer},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InterfaceName},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Active},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${StaticIp},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Fqdn},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DomainName},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Hostname},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalId},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${PtrId},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${StartDate},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${EndDate},
        [Parameter(Position = 21, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefType},
        [Parameter(Position = 22, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${RefId},
        [Parameter(Position = 23, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${SubRefId},
        [Parameter(Position = 24, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${NetworkDomain},
        [Parameter(Position = 25, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CreatedBy}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "networkPoolId" = ${NetworkPoolId}
            "ipType" = ${IpType}
            "ipAddress" = ${IpAddress}
            "gatewayAddress" = ${GatewayAddress}
            "subnetMask" = ${SubnetMask}
            "dnsServer" = ${DnsServer}
            "interfaceName" = ${InterfaceName}
            "description" = ${Description}
            "active" = ${Active}
            "staticIp" = ${StaticIp}
            "fqdn" = ${Fqdn}
            "domainName" = ${DomainName}
            "hostname" = ${Hostname}
            "internalId" = ${InternalId}
            "externalId" = ${ExternalId}
            "ptrId" = ${PtrId}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "refType" = ${RefType}
            "refId" = ${RefId}
            "subRefId" = ${SubRefId}
            "networkDomain" = ${NetworkDomain}
            "createdBy" = ${CreatedBy}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner<PSCustomObject>
#>
function ConvertFrom-JsonToGetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner
        $AllProperties = ("id", "networkPoolId", "ipType", "ipAddress", "gatewayAddress", "subnetMask", "dnsServer", "interfaceName", "description", "active", "staticIp", "fqdn", "domainName", "hostname", "internalId", "externalId", "ptrId", "dateCreated", "lastUpdated", "startDate", "endDate", "refType", "refId", "subRefId", "networkDomain", "createdBy")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkPoolId"))) { #optional property not found
            $NetworkPoolId = $null
        } else {
            $NetworkPoolId = $JsonParameters.PSobject.Properties["networkPoolId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipType"))) { #optional property not found
            $IpType = $null
        } else {
            $IpType = $JsonParameters.PSobject.Properties["ipType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipAddress"))) { #optional property not found
            $IpAddress = $null
        } else {
            $IpAddress = $JsonParameters.PSobject.Properties["ipAddress"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "gatewayAddress"))) { #optional property not found
            $GatewayAddress = $null
        } else {
            $GatewayAddress = $JsonParameters.PSobject.Properties["gatewayAddress"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subnetMask"))) { #optional property not found
            $SubnetMask = $null
        } else {
            $SubnetMask = $JsonParameters.PSobject.Properties["subnetMask"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dnsServer"))) { #optional property not found
            $DnsServer = $null
        } else {
            $DnsServer = $JsonParameters.PSobject.Properties["dnsServer"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "interfaceName"))) { #optional property not found
            $InterfaceName = $null
        } else {
            $InterfaceName = $JsonParameters.PSobject.Properties["interfaceName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "active"))) { #optional property not found
            $Active = $null
        } else {
            $Active = $JsonParameters.PSobject.Properties["active"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "staticIp"))) { #optional property not found
            $StaticIp = $null
        } else {
            $StaticIp = $JsonParameters.PSobject.Properties["staticIp"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "fqdn"))) { #optional property not found
            $Fqdn = $null
        } else {
            $Fqdn = $JsonParameters.PSobject.Properties["fqdn"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "domainName"))) { #optional property not found
            $DomainName = $null
        } else {
            $DomainName = $JsonParameters.PSobject.Properties["domainName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostname"))) { #optional property not found
            $Hostname = $null
        } else {
            $Hostname = $JsonParameters.PSobject.Properties["hostname"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ptrId"))) { #optional property not found
            $PtrId = $null
        } else {
            $PtrId = $JsonParameters.PSobject.Properties["ptrId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "startDate"))) { #optional property not found
            $StartDate = $null
        } else {
            $StartDate = $JsonParameters.PSobject.Properties["startDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "endDate"))) { #optional property not found
            $EndDate = $null
        } else {
            $EndDate = $JsonParameters.PSobject.Properties["endDate"].value
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "subRefId"))) { #optional property not found
            $SubRefId = $null
        } else {
            $SubRefId = $JsonParameters.PSobject.Properties["subRefId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkDomain"))) { #optional property not found
            $NetworkDomain = $null
        } else {
            $NetworkDomain = $JsonParameters.PSobject.Properties["networkDomain"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdBy"))) { #optional property not found
            $CreatedBy = $null
        } else {
            $CreatedBy = $JsonParameters.PSobject.Properties["createdBy"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "networkPoolId" = ${NetworkPoolId}
            "ipType" = ${IpType}
            "ipAddress" = ${IpAddress}
            "gatewayAddress" = ${GatewayAddress}
            "subnetMask" = ${SubnetMask}
            "dnsServer" = ${DnsServer}
            "interfaceName" = ${InterfaceName}
            "description" = ${Description}
            "active" = ${Active}
            "staticIp" = ${StaticIp}
            "fqdn" = ${Fqdn}
            "domainName" = ${DomainName}
            "hostname" = ${Hostname}
            "internalId" = ${InternalId}
            "externalId" = ${ExternalId}
            "ptrId" = ${PtrId}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "startDate" = ${StartDate}
            "endDate" = ${EndDate}
            "refType" = ${RefType}
            "refId" = ${RefId}
            "subRefId" = ${SubRefId}
            "networkDomain" = ${NetworkDomain}
            "createdBy" = ${CreatedBy}
        }

        return $PSO
    }

}

