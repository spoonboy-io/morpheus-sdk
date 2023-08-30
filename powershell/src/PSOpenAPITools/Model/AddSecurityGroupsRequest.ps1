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

AddSecurityGroupsRequest<PSCustomObject>
#>

function Initialize-AddSecurityGroupsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${SecurityGroup}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddSecurityGroupsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $SecurityGroup) {
            throw "invalid value for 'SecurityGroup', 'SecurityGroup' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "securityGroup" = ${SecurityGroup}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddSecurityGroupsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddSecurityGroupsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddSecurityGroupsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddSecurityGroupsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddSecurityGroupsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddSecurityGroupsRequest
        $AllProperties = ("securityGroup")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'securityGroup' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "securityGroup"))) {
            throw "Error! JSON cannot be serialized due to the required property 'securityGroup' missing."
        } else {
            $SecurityGroup = $JsonParameters.PSobject.Properties["securityGroup"].value
        }

        $PSO = [PSCustomObject]@{
            "securityGroup" = ${SecurityGroup}
        }

        return $PSO
    }

}

