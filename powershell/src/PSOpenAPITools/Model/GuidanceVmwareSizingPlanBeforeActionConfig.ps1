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

.PARAMETER StorageSizeType
No description available.
.PARAMETER MemorySizeType
No description available.
.PARAMETER Ranges
No description available.
.OUTPUTS

GuidanceVmwareSizingPlanBeforeActionConfig<PSCustomObject>
#>

function Initialize-GuidanceVmwareSizingPlanBeforeActionConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${StorageSizeType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MemorySizeType},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Ranges}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GuidanceVmwareSizingPlanBeforeActionConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "storageSizeType" = ${StorageSizeType}
            "memorySizeType" = ${MemorySizeType}
            "ranges" = ${Ranges}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GuidanceVmwareSizingPlanBeforeActionConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to GuidanceVmwareSizingPlanBeforeActionConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GuidanceVmwareSizingPlanBeforeActionConfig<PSCustomObject>
#>
function ConvertFrom-JsonToGuidanceVmwareSizingPlanBeforeActionConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GuidanceVmwareSizingPlanBeforeActionConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GuidanceVmwareSizingPlanBeforeActionConfig
        $AllProperties = ("storageSizeType", "memorySizeType", "ranges")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "storageSizeType"))) { #optional property not found
            $StorageSizeType = $null
        } else {
            $StorageSizeType = $JsonParameters.PSobject.Properties["storageSizeType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "memorySizeType"))) { #optional property not found
            $MemorySizeType = $null
        } else {
            $MemorySizeType = $JsonParameters.PSobject.Properties["memorySizeType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ranges"))) { #optional property not found
            $Ranges = $null
        } else {
            $Ranges = $JsonParameters.PSobject.Properties["ranges"].value
        }

        $PSO = [PSCustomObject]@{
            "storageSizeType" = ${StorageSizeType}
            "memorySizeType" = ${MemorySizeType}
            "ranges" = ${Ranges}
        }

        return $PSO
    }

}
