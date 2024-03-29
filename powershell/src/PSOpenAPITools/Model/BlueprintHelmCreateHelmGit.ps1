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

.PARAMETER RepoId
Morpheus SCM Repository ID
.PARAMETER Path
Path to Helm Files in the Repository
.PARAMETER IntegrationId
Morpheus SCM Integration ID
.PARAMETER Branch
Branch Name
.OUTPUTS

BlueprintHelmCreateHelmGit<PSCustomObject>
#>

function Initialize-BlueprintHelmCreateHelmGit {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [Int64]
        ${RepoId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Path},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [Int64]
        ${IntegrationId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Branch}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BlueprintHelmCreateHelmGit' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $RepoId) {
            throw "invalid value for 'RepoId', 'RepoId' cannot be null."
        }

        if ($null -eq $Path) {
            throw "invalid value for 'Path', 'Path' cannot be null."
        }

        if ($null -eq $IntegrationId) {
            throw "invalid value for 'IntegrationId', 'IntegrationId' cannot be null."
        }

        if ($null -eq $Branch) {
            throw "invalid value for 'Branch', 'Branch' cannot be null."
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

<#
.SYNOPSIS

Convert from JSON to BlueprintHelmCreateHelmGit<PSCustomObject>

.DESCRIPTION

Convert from JSON to BlueprintHelmCreateHelmGit<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BlueprintHelmCreateHelmGit<PSCustomObject>
#>
function ConvertFrom-JsonToBlueprintHelmCreateHelmGit {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BlueprintHelmCreateHelmGit' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BlueprintHelmCreateHelmGit
        $AllProperties = ("repoId", "path", "integrationId", "branch")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'repoId' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "repoId"))) {
            throw "Error! JSON cannot be serialized due to the required property 'repoId' missing."
        } else {
            $RepoId = $JsonParameters.PSobject.Properties["repoId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "path"))) {
            throw "Error! JSON cannot be serialized due to the required property 'path' missing."
        } else {
            $Path = $JsonParameters.PSobject.Properties["path"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "integrationId"))) {
            throw "Error! JSON cannot be serialized due to the required property 'integrationId' missing."
        } else {
            $IntegrationId = $JsonParameters.PSobject.Properties["integrationId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "branch"))) {
            throw "Error! JSON cannot be serialized due to the required property 'branch' missing."
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

