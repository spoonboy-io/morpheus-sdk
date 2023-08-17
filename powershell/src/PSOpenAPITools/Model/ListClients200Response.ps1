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

.PARAMETER Meta
No description available.
.PARAMETER Clients
No description available.
.OUTPUTS

ListClients200Response<PSCustomObject>
#>

function Initialize-ListClients200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Meta},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Clients}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ListClients200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "meta" = ${Meta}
            "clients" = ${Clients}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ListClients200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to ListClients200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ListClients200Response<PSCustomObject>
#>
function ConvertFrom-JsonToListClients200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ListClients200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ListClients200Response
        $AllProperties = ("meta", "clients")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "meta"))) { #optional property not found
            $Meta = $null
        } else {
            $Meta = $JsonParameters.PSobject.Properties["meta"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "clients"))) { #optional property not found
            $Clients = $null
        } else {
            $Clients = $JsonParameters.PSobject.Properties["clients"].value
        }

        $PSO = [PSCustomObject]@{
            "meta" = ${Meta}
            "clients" = ${Clients}
        }

        return $PSO
    }

}

