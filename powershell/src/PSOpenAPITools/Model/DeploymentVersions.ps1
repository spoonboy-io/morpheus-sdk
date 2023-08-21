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
.PARAMETER DeployType
No description available.
.PARAMETER FetchUrl
No description available.
.PARAMETER GitUrl
No description available.
.PARAMETER GitRef
No description available.
.PARAMETER UserVersion
No description available.
.PARAMETER Version
No description available.
.PARAMETER Status
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.OUTPUTS

DeploymentVersions<PSCustomObject>
#>

function Initialize-DeploymentVersions {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DeployType},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${FetchUrl},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GitUrl},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${GitRef},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${UserVersion},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Version},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Status},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => DeploymentVersions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "deployType" = ${DeployType}
            "fetchUrl" = ${FetchUrl}
            "gitUrl" = ${GitUrl}
            "gitRef" = ${GitRef}
            "userVersion" = ${UserVersion}
            "version" = ${Version}
            "status" = ${Status}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to DeploymentVersions<PSCustomObject>

.DESCRIPTION

Convert from JSON to DeploymentVersions<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

DeploymentVersions<PSCustomObject>
#>
function ConvertFrom-JsonToDeploymentVersions {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => DeploymentVersions' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in DeploymentVersions
        $AllProperties = ("id", "deployType", "fetchUrl", "gitUrl", "gitRef", "userVersion", "version", "status", "dateCreated", "lastUpdated")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "deployType"))) { #optional property not found
            $DeployType = $null
        } else {
            $DeployType = $JsonParameters.PSobject.Properties["deployType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "fetchUrl"))) { #optional property not found
            $FetchUrl = $null
        } else {
            $FetchUrl = $JsonParameters.PSobject.Properties["fetchUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "gitUrl"))) { #optional property not found
            $GitUrl = $null
        } else {
            $GitUrl = $JsonParameters.PSobject.Properties["gitUrl"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "gitRef"))) { #optional property not found
            $GitRef = $null
        } else {
            $GitRef = $JsonParameters.PSobject.Properties["gitRef"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "userVersion"))) { #optional property not found
            $UserVersion = $null
        } else {
            $UserVersion = $JsonParameters.PSobject.Properties["userVersion"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "version"))) { #optional property not found
            $Version = $null
        } else {
            $Version = $JsonParameters.PSobject.Properties["version"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "status"))) { #optional property not found
            $Status = $null
        } else {
            $Status = $JsonParameters.PSobject.Properties["status"].value
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
            "deployType" = ${DeployType}
            "fetchUrl" = ${FetchUrl}
            "gitUrl" = ${GitUrl}
            "gitRef" = ${GitRef}
            "userVersion" = ${UserVersion}
            "version" = ${Version}
            "status" = ${Status}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
        }

        return $PSO
    }

}

