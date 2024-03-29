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

Object for logging in to the [Morpheus Hub](https://morpheushub.com) with existing credentials. This is only required for `hubmode=login`.

.PARAMETER Hub
No description available.
.OUTPUTS

HubLoginObject<PSCustomObject>
#>

function Initialize-HubLoginObject {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Hub}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HubLoginObject' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Hub) {
            throw "invalid value for 'Hub', 'Hub' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "hub" = ${Hub}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HubLoginObject<PSCustomObject>

.DESCRIPTION

Convert from JSON to HubLoginObject<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HubLoginObject<PSCustomObject>
#>
function ConvertFrom-JsonToHubLoginObject {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HubLoginObject' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HubLoginObject
        $AllProperties = ("hub")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'hub' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hub"))) {
            throw "Error! JSON cannot be serialized due to the required property 'hub' missing."
        } else {
            $Hub = $JsonParameters.PSobject.Properties["hub"].value
        }

        $PSO = [PSCustomObject]@{
            "hub" = ${Hub}
        }

        return $PSO
    }

}

