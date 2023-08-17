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

.PARAMETER Success
No description available.
.PARAMETER ArchiveFile
No description available.
.OUTPUTS

AddArchiveFile200Response<PSCustomObject>
#>

function Initialize-AddArchiveFile200Response {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Success},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ArchiveFile}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddArchiveFile200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "success" = ${Success}
            "archiveFile" = ${ArchiveFile}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddArchiveFile200Response<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddArchiveFile200Response<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddArchiveFile200Response<PSCustomObject>
#>
function ConvertFrom-JsonToAddArchiveFile200Response {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddArchiveFile200Response' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddArchiveFile200Response
        $AllProperties = ("success", "archiveFile")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "success"))) { #optional property not found
            $Success = $null
        } else {
            $Success = $JsonParameters.PSobject.Properties["success"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "archiveFile"))) { #optional property not found
            $ArchiveFile = $null
        } else {
            $ArchiveFile = $JsonParameters.PSobject.Properties["archiveFile"].value
        }

        $PSO = [PSCustomObject]@{
            "success" = ${Success}
            "archiveFile" = ${ArchiveFile}
        }

        return $PSO
    }

}

