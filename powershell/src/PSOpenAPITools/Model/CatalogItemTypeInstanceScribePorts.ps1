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

.PARAMETER Port
Port number.
.PARAMETER Name
A name for the port.
.PARAMETER Lb
The load balancer protocol. HTTP, HTTPS, or TCP.
.OUTPUTS

CatalogItemTypeInstanceScribePorts<PSCustomObject>
#>

function Initialize-CatalogItemTypeInstanceScribePorts {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Port},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Lb}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CatalogItemTypeInstanceScribePorts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "port" = ${Port}
            "name" = ${Name}
            "lb" = ${Lb}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CatalogItemTypeInstanceScribePorts<PSCustomObject>

.DESCRIPTION

Convert from JSON to CatalogItemTypeInstanceScribePorts<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CatalogItemTypeInstanceScribePorts<PSCustomObject>
#>
function ConvertFrom-JsonToCatalogItemTypeInstanceScribePorts {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CatalogItemTypeInstanceScribePorts' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CatalogItemTypeInstanceScribePorts
        $AllProperties = ("port", "name", "lb")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "port"))) { #optional property not found
            $Port = $null
        } else {
            $Port = $JsonParameters.PSobject.Properties["port"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lb"))) { #optional property not found
            $Lb = $null
        } else {
            $Lb = $JsonParameters.PSobject.Properties["lb"].value
        }

        $PSO = [PSCustomObject]@{
            "port" = ${Port}
            "name" = ${Name}
            "lb" = ${Lb}
        }

        return $PSO
    }

}
