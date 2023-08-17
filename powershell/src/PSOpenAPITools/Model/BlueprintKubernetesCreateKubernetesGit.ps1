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

.PARAMETER RepoId
Morpheus SCM Repository ID
.PARAMETER Path
Path to kubernetes Files in the Repository
.PARAMETER IntegrationId
Morpheus SCM Integration ID
.PARAMETER Branch
Branch Name
.OUTPUTS

BlueprintKubernetesCreateKubernetesGit<PSCustomObject>
#>

function Initialize-BlueprintKubernetesCreateKubernetesGit {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${RepoId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Path},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${IntegrationId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Branch}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintKubernetesCreateKubernetesGit' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "repoId" = ${RepoId}
            "path" = ${Path}
            "integrationId" = ${IntegrationId}
            "branch" = ${Branch}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BlueprintKubernetesCreateKubernetesGit<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintKubernetesCreateKubernetesGit<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintKubernetesCreateKubernetesGit<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintKubernetesCreateKubernetesGit {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintKubernetesCreateKubernetesGit' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintKubernetesCreateKubernetesGit
        $AllProperties = ("repoId", "path", "integrationId", "branch")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "repoId"))) { #optional property not found
            $RepoId = $null
        } else {
            $RepoId = $JsonParameters.PSobject.Properties["repoId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "path"))) { #optional property not found
            $Path = $null
        } else {
            $Path = $JsonParameters.PSobject.Properties["path"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "integrationId"))) { #optional property not found
            $IntegrationId = $null
        } else {
            $IntegrationId = $JsonParameters.PSobject.Properties["integrationId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "branch"))) { #optional property not found
            $Branch = $null
        } else {
            $Branch = $JsonParameters.PSobject.Properties["branch"].value
        }

        $PSO = [PSCustomObject]@{
            "repoId" = ${RepoId}
            "path" = ${Path}
            "integrationId" = ${IntegrationId}
            "branch" = ${Branch}
        }

        return $PSO
    }

}
