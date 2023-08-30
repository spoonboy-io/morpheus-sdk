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

.PARAMETER SecurityGroup
No description available.
.OUTPUTS

GetSecurityGroups200Response<PSCustomObject>
#>

function Initialize-GetSecurityGroups200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${SecurityGroup}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetSecurityGroups200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "securityGroup" = ${SecurityGroup}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetSecurityGroups200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetSecurityGroups200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetSecurityGroups200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetSecurityGroups200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetSecurityGroups200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetSecurityGroups200Response
        $AllProperties = ("securityGroup")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "securityGroup"))) { #optional property not found
            $SecurityGroup = $null
        } else {
            $SecurityGroup = $JsonParameters.PSobject.Properties["securityGroup"].value
        }

        $PSO = [PSCustomObject]@{
            "securityGroup" = ${SecurityGroup}
        }

        return $PSO
    }

}

