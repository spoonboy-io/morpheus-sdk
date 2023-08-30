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

.PARAMETER Policy
No description available.
.OUTPUTS

AddPoliciesCloudRequest<PSCustomObject>
#>

function Initialize-AddPoliciesCloudRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Policy}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddPoliciesCloudRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Policy) {
            throw "invalid value for 'Policy', 'Policy' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "policy" = ${Policy}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddPoliciesCloudRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddPoliciesCloudRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddPoliciesCloudRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddPoliciesCloudRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddPoliciesCloudRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddPoliciesCloudRequest
        $AllProperties = ("policy")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'policy' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "policy"))) {
            throw "Error! JSON cannot be serialized due to the required property 'policy' missing."
        } else {
            $Policy = $JsonParameters.PSobject.Properties["policy"].value
        }

        $PSO = [PSCustomObject]@{
            "policy" = ${Policy}
        }

        return $PSO
    }

}

