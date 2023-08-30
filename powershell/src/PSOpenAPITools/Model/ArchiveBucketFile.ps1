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

.PARAMETER Id
No description available.
.PARAMETER Name
No description available.
.PARAMETER FilePath
No description available.
.PARAMETER ArchiveBucket
No description available.
.PARAMETER CreatedBy
No description available.
.PARAMETER IsDirectory
No description available.
.PARAMETER Status
No description available.
.PARAMETER RawSize
No description available.
.PARAMETER ContentType
No description available.
.PARAMETER DownloadCount
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.OUTPUTS

ArchiveBucketFile<PSCustomObject>
#>

function Initialize-ArchiveBucketFile {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${FilePath},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ArchiveBucket},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${CreatedBy},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${IsDirectory},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RawSize},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ContentType},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${DownloadCount},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ArchiveBucketFile' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "filePath" = ${FilePath}
            "archiveBucket" = ${ArchiveBucket}
            "createdBy" = ${CreatedBy}
            "isDirectory" = ${IsDirectory}
            "status" = ${Status}
            "rawSize" = ${RawSize}
            "contentType" = ${ContentType}
            "downloadCount" = ${DownloadCount}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ArchiveBucketFile<PSCustomObject>

.DESCRIPTION

Convert from JSON to ArchiveBucketFile<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ArchiveBucketFile<PSCustomObject>
#>
function ConvertFrom-JsonToArchiveBucketFile {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ArchiveBucketFile' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ArchiveBucketFile
        $AllProperties = ("id", "name", "filePath", "archiveBucket", "createdBy", "isDirectory", "status", "rawSize", "contentType", "downloadCount", "dateCreated", "lastUpdated")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "filePath"))) { #optional property not found
            $FilePath = $null
        } else {
            $FilePath = $JsonParameters.PSobject.Properties["filePath"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "archiveBucket"))) { #optional property not found
            $ArchiveBucket = $null
        } else {
            $ArchiveBucket = $JsonParameters.PSobject.Properties["archiveBucket"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdBy"))) { #optional property not found
            $CreatedBy = $null
        } else {
            $CreatedBy = $JsonParameters.PSobject.Properties["createdBy"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "isDirectory"))) { #optional property not found
            $IsDirectory = $null
        } else {
            $IsDirectory = $JsonParameters.PSobject.Properties["isDirectory"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "rawSize"))) { #optional property not found
            $RawSize = $null
        } else {
            $RawSize = $JsonParameters.PSobject.Properties["rawSize"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "contentType"))) { #optional property not found
            $ContentType = $null
        } else {
            $ContentType = $JsonParameters.PSobject.Properties["contentType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "downloadCount"))) { #optional property not found
            $DownloadCount = $null
        } else {
            $DownloadCount = $JsonParameters.PSobject.Properties["downloadCount"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "filePath" = ${FilePath}
            "archiveBucket" = ${ArchiveBucket}
            "createdBy" = ${CreatedBy}
            "isDirectory" = ${IsDirectory}
            "status" = ${Status}
            "rawSize" = ${RawSize}
            "contentType" = ${ContentType}
            "downloadCount" = ${DownloadCount}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }

        return $PSO
    }

}

