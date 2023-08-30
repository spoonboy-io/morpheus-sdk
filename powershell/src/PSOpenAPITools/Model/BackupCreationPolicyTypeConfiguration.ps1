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

Configuration settings for the following policy types: - Backup Creation 

.PARAMETER CreateBackupType
No description available.
.PARAMETER CreateBackup
No description available.
.OUTPUTS

BackupCreationPolicyTypeConfiguration<PSCustomObject>
#>

function Initialize-BackupCreationPolicyTypeConfiguration {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CreateBackupType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${CreateBackup}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BackupCreationPolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "createBackupType" = ${CreateBackupType}
            "createBackup" = ${CreateBackup}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BackupCreationPolicyTypeConfiguration<PSCustomObject>

.DESCRIPTION

Convert from JSON to BackupCreationPolicyTypeConfiguration<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BackupCreationPolicyTypeConfiguration<PSCustomObject>
#>
function ConvertFrom-JsonToBackupCreationPolicyTypeConfiguration {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BackupCreationPolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BackupCreationPolicyTypeConfiguration
        $AllProperties = ("createBackupType", "createBackup")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createBackupType"))) { #optional property not found
            $CreateBackupType = $null
        } else {
            $CreateBackupType = $JsonParameters.PSobject.Properties["createBackupType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createBackup"))) { #optional property not found
            $CreateBackup = $null
        } else {
            $CreateBackup = $JsonParameters.PSobject.Properties["createBackup"].value
        }

        $PSO = [PSCustomObject]@{
            "createBackupType" = ${CreateBackupType}
            "createBackup" = ${CreateBackup}
        }

        return $PSO
    }

}

