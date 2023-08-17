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

.PARAMETER ArchiveFile
No description available.
.PARAMETER IsOwner
No description available.
.OUTPUTS

GetArchiveFileDetail200Response<PSCustomObject>
#>

function Initialize-GetArchiveFileDetail200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ArchiveFile},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IsOwner}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetArchiveFileDetail200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "archiveFile" = ${ArchiveFile}
            "isOwner" = ${IsOwner}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetArchiveFileDetail200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetArchiveFileDetail200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetArchiveFileDetail200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetArchiveFileDetail200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetArchiveFileDetail200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetArchiveFileDetail200Response
        $AllProperties = ("archiveFile", "isOwner")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "archiveFile"))) { #optional property not found
            $ArchiveFile = $null
        } else {
            $ArchiveFile = $JsonParameters.PSobject.Properties["archiveFile"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "isOwner"))) { #optional property not found
            $IsOwner = $null
        } else {
            $IsOwner = $JsonParameters.PSobject.Properties["isOwner"].value
        }

        $PSO = [PSCustomObject]@{
            "archiveFile" = ${ArchiveFile}
            "isOwner" = ${IsOwner}
        }

        return $PSO
    }

}

