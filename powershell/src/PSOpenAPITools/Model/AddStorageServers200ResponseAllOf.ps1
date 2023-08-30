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

.PARAMETER StorageServer
No description available.
.OUTPUTS

AddStorageServers200ResponseAllOf<PSCustomObject>
#>

function Initialize-AddStorageServers200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${StorageServer}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddStorageServers200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "storageServer" = ${StorageServer}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddStorageServers200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddStorageServers200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddStorageServers200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddStorageServers200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddStorageServers200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddStorageServers200ResponseAllOf
        $AllProperties = ("storageServer")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "storageServer"))) { #optional property not found
            $StorageServer = $null
        } else {
            $StorageServer = $JsonParameters.PSobject.Properties["storageServer"].value
        }

        $PSO = [PSCustomObject]@{
            "storageServer" = ${StorageServer}
        }

        return $PSO
    }

}

