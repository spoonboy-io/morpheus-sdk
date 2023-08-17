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

.PARAMETER ApplianceSettings
No description available.
.OUTPUTS

UpdateApplianceSettingsRequest<PSCustomObject>
#>

function Initialize-UpdateApplianceSettingsRequest {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ApplianceSettings}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => UpdateApplianceSettingsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "applianceSettings" = ${ApplianceSettings}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to UpdateApplianceSettingsRequest<PSCustomObject>

.DESCRIPTION

Convert from JSON to UpdateApplianceSettingsRequest<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

UpdateApplianceSettingsRequest<PSCustomObject>
#>
function ConvertFrom-JsonToUpdateApplianceSettingsRequest {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => UpdateApplianceSettingsRequest' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in UpdateApplianceSettingsRequest
        $AllProperties = ("applianceSettings")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "applianceSettings"))) { #optional property not found
            $ApplianceSettings = $null
        } else {
            $ApplianceSettings = $JsonParameters.PSobject.Properties["applianceSettings"].value
        }

        $PSO = [PSCustomObject]@{
            "applianceSettings" = ${ApplianceSettings}
        }

        return $PSO
    }

}
