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

.PARAMETER Name
Name
.PARAMETER ProxyHost
Proxy Host
.PARAMETER ProxyPort
Proxy Port
.PARAMETER ProxyUser
Proxy Username
.PARAMETER ProxyPassword
Proxy Password
.PARAMETER ProxyDomain
Proxy Domain
.PARAMETER ProxyWorkstation
Proxy Workstation
.PARAMETER Visibility
Visibility
.PARAMETER Account
No description available.
.OUTPUTS

ApiNetworksProxiesNetworkProxy<PSCustomObject>
#>

function Initialize-ApiNetworksProxiesNetworkProxy {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProxyHost},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProxyPort},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProxyUser},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProxyPassword},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProxyDomain},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ProxyWorkstation},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Visibility} = "private",
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Account}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiNetworksProxiesNetworkProxy' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "proxyHost" = ${ProxyHost}
            "proxyPort" = ${ProxyPort}
            "proxyUser" = ${ProxyUser}
            "proxyPassword" = ${ProxyPassword}
            "proxyDomain" = ${ProxyDomain}
            "proxyWorkstation" = ${ProxyWorkstation}
            "visibility" = ${Visibility}
            "account" = ${Account}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiNetworksProxiesNetworkProxy<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiNetworksProxiesNetworkProxy<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiNetworksProxiesNetworkProxy<PSCustomObject>
#>
function ConvertFrom-JsonToApiNetworksProxiesNetworkProxy {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiNetworksProxiesNetworkProxy' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiNetworksProxiesNetworkProxy
        $AllProperties = ("name", "proxyHost", "proxyPort", "proxyUser", "proxyPassword", "proxyDomain", "proxyWorkstation", "visibility", "account")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "proxyHost"))) { #optional property not found
            $ProxyHost = $null
        } else {
            $ProxyHost = $JsonParameters.PSobject.Properties["proxyHost"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "proxyPort"))) { #optional property not found
            $ProxyPort = $null
        } else {
            $ProxyPort = $JsonParameters.PSobject.Properties["proxyPort"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "proxyUser"))) { #optional property not found
            $ProxyUser = $null
        } else {
            $ProxyUser = $JsonParameters.PSobject.Properties["proxyUser"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "proxyPassword"))) { #optional property not found
            $ProxyPassword = $null
        } else {
            $ProxyPassword = $JsonParameters.PSobject.Properties["proxyPassword"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "proxyDomain"))) { #optional property not found
            $ProxyDomain = $null
        } else {
            $ProxyDomain = $JsonParameters.PSobject.Properties["proxyDomain"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "proxyWorkstation"))) { #optional property not found
            $ProxyWorkstation = $null
        } else {
            $ProxyWorkstation = $JsonParameters.PSobject.Properties["proxyWorkstation"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "visibility"))) { #optional property not found
            $Visibility = $null
        } else {
            $Visibility = $JsonParameters.PSobject.Properties["visibility"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "account"))) { #optional property not found
            $Account = $null
        } else {
            $Account = $JsonParameters.PSobject.Properties["account"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "proxyHost" = ${ProxyHost}
            "proxyPort" = ${ProxyPort}
            "proxyUser" = ${ProxyUser}
            "proxyPassword" = ${ProxyPassword}
            "proxyDomain" = ${ProxyDomain}
            "proxyWorkstation" = ${ProxyWorkstation}
            "visibility" = ${Visibility}
            "account" = ${Account}
        }

        return $PSO
    }

}

