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

.PARAMETER StorageVolume
No description available.
.OUTPUTS

AddStorageVolumesRequest<PSCustomObject>
#>

function Initialize-AddStorageVolumesRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${StorageVolume}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddStorageVolumesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $StorageVolume) {
            throw "invalid value for 'StorageVolume', 'StorageVolume' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "storageVolume" = ${StorageVolume}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddStorageVolumesRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddStorageVolumesRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddStorageVolumesRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddStorageVolumesRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddStorageVolumesRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddStorageVolumesRequest
        $AllProperties = ("storageVolume")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'storageVolume' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "storageVolume"))) {
            throw "Error! JSON cannot be serialized due to the required property 'storageVolume' missing."
        } else {
            $StorageVolume = $JsonParameters.PSobject.Properties["storageVolume"].value
        }

        $PSO = [PSCustomObject]@{
            "storageVolume" = ${StorageVolume}
        }

        return $PSO
    }

}
