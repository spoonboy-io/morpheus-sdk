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

.PARAMETER SecurityGroups
No description available.
.PARAMETER FirewallEnabled
No description available.
.PARAMETER Success
No description available.
.OUTPUTS

GetAppSecurityGroups200Response<PSCustomObject>
#>

function Initialize-GetAppSecurityGroups200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${SecurityGroups},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${FirewallEnabled},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetAppSecurityGroups200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "securityGroups" = ${SecurityGroups}
            "firewallEnabled" = ${FirewallEnabled}
            "success" = ${Success}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetAppSecurityGroups200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetAppSecurityGroups200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetAppSecurityGroups200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetAppSecurityGroups200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetAppSecurityGroups200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetAppSecurityGroups200Response
        $AllProperties = ("securityGroups", "firewallEnabled", "success")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "securityGroups"))) { #optional property not found
            $SecurityGroups = $null
        } else {
            $SecurityGroups = $JsonParameters.PSobject.Properties["securityGroups"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "firewallEnabled"))) { #optional property not found
            $FirewallEnabled = $null
        } else {
            $FirewallEnabled = $JsonParameters.PSobject.Properties["firewallEnabled"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        $PSO = [PSCustomObject]@{
            "securityGroups" = ${SecurityGroups}
            "firewallEnabled" = ${FirewallEnabled}
            "success" = ${Success}
        }

        return $PSO
    }

}

