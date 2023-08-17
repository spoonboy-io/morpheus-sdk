#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER Client
No description available.
.OUTPUTS

UpdateClientsRequest<PSCustomObject>
#>

function Initialize-UpdateClientsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Client}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => UpdateClientsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Client) {
            throw "invalid value for 'Client', 'Client' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "client" = ${Client}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to UpdateClientsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to UpdateClientsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

UpdateClientsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToUpdateClientsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => UpdateClientsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in UpdateClientsRequest
        $AllProperties = ("client")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'client' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "client"))) {
            throw "Error! JSON cannot be serialized due to the required property 'client' missing."
        } else {
            $Client = $JsonParameters.PSobject.Properties["client"].value
        }

        $PSO = [PSCustomObject]@{
            "client" = ${Client}
        }

        return $PSO
    }

}

