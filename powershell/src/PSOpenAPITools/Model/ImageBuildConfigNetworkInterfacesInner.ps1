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

.PARAMETER PrimaryInterface
No description available.
.PARAMETER Network
No description available.
.OUTPUTS

ImageBuildConfigNetworkInterfacesInner<PSCustomObject>
#>

function Initialize-ImageBuildConfigNetworkInterfacesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${PrimaryInterface},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Network}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ImageBuildConfigNetworkInterfacesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "primaryInterface" = ${PrimaryInterface}
            "network" = ${Network}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ImageBuildConfigNetworkInterfacesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to ImageBuildConfigNetworkInterfacesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ImageBuildConfigNetworkInterfacesInner<PSCustomObject>
#>
function ConvertFrom-JsonToImageBuildConfigNetworkInterfacesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ImageBuildConfigNetworkInterfacesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ImageBuildConfigNetworkInterfacesInner
        $AllProperties = ("primaryInterface", "network")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "primaryInterface"))) { #optional property not found
            $PrimaryInterface = $null
        } else {
            $PrimaryInterface = $JsonParameters.PSobject.Properties["primaryInterface"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "network"))) { #optional property not found
            $Network = $null
        } else {
            $Network = $JsonParameters.PSobject.Properties["network"].value
        }

        $PSO = [PSCustomObject]@{
            "primaryInterface" = ${PrimaryInterface}
            "network" = ${Network}
        }

        return $PSO
    }

}

