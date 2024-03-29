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
.PARAMETER Description
Description
.PARAMETER IpAddress
IP Address
.PARAMETER Port
Port number
.PARAMETER Weight
Weight applied balance algoritm
.PARAMETER Config
Configuration object with parameters that vary by type.
.OUTPUTS

CreateLoadBalancerPoolNodeRequestLoadBalancerNode<PSCustomObject>
#>

function Initialize-CreateLoadBalancerPoolNodeRequestLoadBalancerNode {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${IpAddress},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${Port},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${Weight},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateLoadBalancerPoolNodeRequestLoadBalancerNode' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "ipAddress" = ${IpAddress}
            "port" = ${Port}
            "weight" = ${Weight}
            "config" = ${Config}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateLoadBalancerPoolNodeRequestLoadBalancerNode<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateLoadBalancerPoolNodeRequestLoadBalancerNode<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateLoadBalancerPoolNodeRequestLoadBalancerNode<PSCustomObject>
#>
function ConvertFrom-JsonToCreateLoadBalancerPoolNodeRequestLoadBalancerNode {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateLoadBalancerPoolNodeRequestLoadBalancerNode' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateLoadBalancerPoolNodeRequestLoadBalancerNode
        $AllProperties = ("name", "description", "ipAddress", "port", "weight", "config")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ipAddress"))) { #optional property not found
            $IpAddress = $null
        } else {
            $IpAddress = $JsonParameters.PSobject.Properties["ipAddress"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "port"))) { #optional property not found
            $Port = $null
        } else {
            $Port = $JsonParameters.PSobject.Properties["port"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "weight"))) { #optional property not found
            $Weight = $null
        } else {
            $Weight = $JsonParameters.PSobject.Properties["weight"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "ipAddress" = ${IpAddress}
            "port" = ${Port}
            "weight" = ${Weight}
            "config" = ${Config}
        }

        return $PSO
    }

}

