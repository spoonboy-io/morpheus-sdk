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

AddStorageBucketsRequest<PSCustomObject>
#>

function Initialize-AddStorageBucketsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${StorageBucket}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddStorageBucketsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $StorageBucket) {
            throw "invalid value for 'StorageBucket', 'StorageBucket' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "storageBucket" = ${StorageBucket}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddStorageBucketsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddStorageBucketsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddStorageBucketsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToAddStorageBucketsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddStorageBucketsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddStorageBucketsRequest
        $AllProperties = ("storageBucket")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'storageBucket' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "storageBucket"))) {
            throw "Error! JSON cannot be serialized due to the required property 'storageBucket' missing."
        } else {
            $StorageBucket = $JsonParameters.PSobject.Properties["storageBucket"].value
        }

        $PSO = [PSCustomObject]@{
            "storageBucket" = ${StorageBucket}
        }

        return $PSO
    }

}
