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
.PARAMETER Success
No description available.
.OUTPUTS

AddStorageServers200Response<PSCustomObject>
#>

function Initialize-AddStorageServers200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${StorageServer},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddStorageServers200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "storageServer" = ${StorageServer}
            "success" = ${Success}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddStorageServers200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddStorageServers200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddStorageServers200Response<PSCustomObject>
#>
function ConvertFrom-JsonToAddStorageServers200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddStorageServers200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddStorageServers200Response
        $AllProperties = ("storageServer", "success")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        $PSO = [PSCustomObject]@{
            "storageServer" = ${StorageServer}
            "success" = ${Success}
        }

        return $PSO
    }

}

