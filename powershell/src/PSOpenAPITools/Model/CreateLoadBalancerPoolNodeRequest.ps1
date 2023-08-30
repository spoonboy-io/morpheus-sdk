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

.PARAMETER LoadBalancerNode
No description available.
.OUTPUTS

CreateLoadBalancerPoolNodeRequest<PSCustomObject>
#>

function Initialize-CreateLoadBalancerPoolNodeRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${LoadBalancerNode}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateLoadBalancerPoolNodeRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "loadBalancerNode" = ${LoadBalancerNode}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateLoadBalancerPoolNodeRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateLoadBalancerPoolNodeRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateLoadBalancerPoolNodeRequest<PSCustomObject>
#>
function ConvertFrom-JsonToCreateLoadBalancerPoolNodeRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateLoadBalancerPoolNodeRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateLoadBalancerPoolNodeRequest
        $AllProperties = ("loadBalancerNode")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loadBalancerNode"))) { #optional property not found
            $LoadBalancerNode = $null
        } else {
            $LoadBalancerNode = $JsonParameters.PSobject.Properties["loadBalancerNode"].value
        }

        $PSO = [PSCustomObject]@{
            "loadBalancerNode" = ${LoadBalancerNode}
        }

        return $PSO
    }

}

