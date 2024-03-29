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
.PARAMETER IsOwner
No description available.
.PARAMETER ParentDirectory
No description available.
.PARAMETER ArchiveFiles
No description available.
.PARAMETER ArchiveFileCount
No description available.
.OUTPUTS

GetArchiveBucket200Response<PSCustomObject>
#>

function Initialize-GetArchiveBucket200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ArchiveBucket},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IsOwner},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ParentDirectory},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${ArchiveFiles},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${ArchiveFileCount}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetArchiveBucket200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "archiveBucket" = ${ArchiveBucket}
            "isOwner" = ${IsOwner}
            "parentDirectory" = ${ParentDirectory}
            "archiveFiles" = ${ArchiveFiles}
            "archiveFileCount" = ${ArchiveFileCount}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetArchiveBucket200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetArchiveBucket200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetArchiveBucket200Response<PSCustomObject>
#>
function ConvertFrom-JsonToGetArchiveBucket200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetArchiveBucket200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetArchiveBucket200Response
        $AllProperties = ("archiveBucket", "isOwner", "parentDirectory", "archiveFiles", "archiveFileCount")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "isOwner"))) { #optional property not found
            $IsOwner = $null
        } else {
            $IsOwner = $JsonParameters.PSobject.Properties["isOwner"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "parentDirectory"))) { #optional property not found
            $ParentDirectory = $null
        } else {
            $ParentDirectory = $JsonParameters.PSobject.Properties["parentDirectory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "archiveFiles"))) { #optional property not found
            $ArchiveFiles = $null
        } else {
            $ArchiveFiles = $JsonParameters.PSobject.Properties["archiveFiles"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "archiveFileCount"))) { #optional property not found
            $ArchiveFileCount = $null
        } else {
            $ArchiveFileCount = $JsonParameters.PSobject.Properties["archiveFileCount"].value
        }

        $PSO = [PSCustomObject]@{
            "archiveBucket" = ${ArchiveBucket}
            "isOwner" = ${IsOwner}
            "parentDirectory" = ${ParentDirectory}
            "archiveFiles" = ${ArchiveFiles}
            "archiveFileCount" = ${ArchiveFileCount}
        }

        return $PSO
    }

}

