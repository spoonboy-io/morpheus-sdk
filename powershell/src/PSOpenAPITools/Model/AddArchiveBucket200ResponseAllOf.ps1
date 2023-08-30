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

.PARAMETER ArchiveBucket
No description available.
.OUTPUTS

AddArchiveBucket200ResponseAllOf<PSCustomObject>
#>

function Initialize-AddArchiveBucket200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ArchiveBucket}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddArchiveBucket200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "archiveBucket" = ${ArchiveBucket}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddArchiveBucket200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddArchiveBucket200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddArchiveBucket200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToAddArchiveBucket200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddArchiveBucket200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddArchiveBucket200ResponseAllOf
        $AllProperties = ("archiveBucket")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "archiveBucket"))) { #optional property not found
            $ArchiveBucket = $null
        } else {
            $ArchiveBucket = $JsonParameters.PSobject.Properties["archiveBucket"].value
        }

        $PSO = [PSCustomObject]@{
            "archiveBucket" = ${ArchiveBucket}
        }

        return $PSO
    }

}

