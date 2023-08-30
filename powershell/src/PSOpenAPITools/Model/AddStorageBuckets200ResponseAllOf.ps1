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

.PARAMETER StorageBucket
No description available.
.OUTPUTS

AddStorageBuckets200ResponseAllOf<PSCustomObject>
#>

function Initialize-AddStorageBuckets200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${StorageBucket}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddStorageBuckets200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "storageBucket" = ${StorageBucket}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddStorageBuckets200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddStorageBuckets200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddStorageBuckets200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddStorageBuckets200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddStorageBuckets200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddStorageBuckets200ResponseAllOf
        $AllProperties = ("storageBucket")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "storageBucket"))) { #optional property not found
            $StorageBucket = $null
        } else {
            $StorageBucket = $JsonParameters.PSobject.Properties["storageBucket"].value
        }

        $PSO = [PSCustomObject]@{
            "storageBucket" = ${StorageBucket}
        }

        return $PSO
    }

}
