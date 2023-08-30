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

Configuration settings for the following policy types: - Approve Delete - Approve Provisiong - Approve Reconfigure 

.PARAMETER AccountIntegrationId
No description available.
.OUTPUTS

ApprovePolicyTypeConfiguration<PSCustomObject>
#>

function Initialize-ApprovePolicyTypeConfiguration {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${AccountIntegrationId}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApprovePolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "accountIntegrationId" = ${AccountIntegrationId}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApprovePolicyTypeConfiguration<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApprovePolicyTypeConfiguration<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApprovePolicyTypeConfiguration<PSCustomObject>
#>
function ConvertFrom-JsonToApprovePolicyTypeConfiguration {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApprovePolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApprovePolicyTypeConfiguration
        $AllProperties = ("accountIntegrationId")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "accountIntegrationId"))) { #optional property not found
            $AccountIntegrationId = $null
        } else {
            $AccountIntegrationId = $JsonParameters.PSobject.Properties["accountIntegrationId"].value
        }

        $PSO = [PSCustomObject]@{
            "accountIntegrationId" = ${AccountIntegrationId}
        }

        return $PSO
    }

}
