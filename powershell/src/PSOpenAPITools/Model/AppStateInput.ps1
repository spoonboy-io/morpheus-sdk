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

.PARAMETER Variables
No description available.
.PARAMETER Providers
No description available.
.PARAMETER VarData
No description available.
.OUTPUTS

AppStateInput<PSCustomObject>
#>

function Initialize-AppStateInput {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Variables},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Providers},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${VarData}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AppStateInput' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "variables" = ${Variables}
            "providers" = ${Providers}
            "data" = ${VarData}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AppStateInput<PSCustomObject>

.DESCRIPTION

Convert from JSON to AppStateInput<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AppStateInput<PSCustomObject>
#>
function ConvertFrom-JsonToAppStateInput {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AppStateInput' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AppStateInput
        $AllProperties = ("variables", "providers", "data")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "variables"))) { #optional property not found
            $Variables = $null
        } else {
            $Variables = $JsonParameters.PSobject.Properties["variables"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "providers"))) { #optional property not found
            $Providers = $null
        } else {
            $Providers = $JsonParameters.PSobject.Properties["providers"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "data"))) { #optional property not found
            $VarData = $null
        } else {
            $VarData = $JsonParameters.PSobject.Properties["data"].value
        }

        $PSO = [PSCustomObject]@{
            "variables" = ${Variables}
            "providers" = ${Providers}
            "data" = ${VarData}
        }

        return $PSO
    }

}

