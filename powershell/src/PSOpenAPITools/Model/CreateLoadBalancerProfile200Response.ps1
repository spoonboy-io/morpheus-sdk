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

.PARAMETER LoadBalancerProfile
No description available.
.PARAMETER Success
No description available.
.PARAMETER Msg
No description available.
.OUTPUTS

CreateLoadBalancerProfile200Response<PSCustomObject>
#>

function Initialize-CreateLoadBalancerProfile200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${LoadBalancerProfile},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Msg}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CreateLoadBalancerProfile200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "loadBalancerProfile" = ${LoadBalancerProfile}
            "success" = ${Success}
            "msg" = ${Msg}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CreateLoadBalancerProfile200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to CreateLoadBalancerProfile200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CreateLoadBalancerProfile200Response<PSCustomObject>
#>
function ConvertFrom-JsonToCreateLoadBalancerProfile200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CreateLoadBalancerProfile200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CreateLoadBalancerProfile200Response
        $AllProperties = ("loadBalancerProfile", "success", "msg")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "loadBalancerProfile"))) { #optional property not found
            $LoadBalancerProfile = $null
        } else {
            $LoadBalancerProfile = $JsonParameters.PSobject.Properties["loadBalancerProfile"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "msg"))) { #optional property not found
            $Msg = $null
        } else {
            $Msg = $JsonParameters.PSobject.Properties["msg"].value
        }

        $PSO = [PSCustomObject]@{
            "loadBalancerProfile" = ${LoadBalancerProfile}
            "success" = ${Success}
            "msg" = ${Msg}
        }

        return $PSO
    }

}
