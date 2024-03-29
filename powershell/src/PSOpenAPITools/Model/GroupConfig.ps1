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

.PARAMETER DnsIntegrationId
No description available.
.PARAMETER ConfigCmdbId
No description available.
.PARAMETER ConfigCmId
No description available.
.PARAMETER ServiceRegistryId
No description available.
.PARAMETER ConfigManagementId
No description available.
.PARAMETER ConfigCmdbDiscovery
No description available.
.OUTPUTS

GroupConfig<PSCustomObject>
#>

function Initialize-GroupConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${DnsIntegrationId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ConfigCmdbId},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ConfigCmId},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ServiceRegistryId},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ConfigManagementId},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${ConfigCmdbDiscovery}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GroupConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "dnsIntegrationId" = ${DnsIntegrationId}
            "configCmdbId" = ${ConfigCmdbId}
            "configCmId" = ${ConfigCmId}
            "serviceRegistryId" = ${ServiceRegistryId}
            "configManagementId" = ${ConfigManagementId}
            "configCmdbDiscovery" = ${ConfigCmdbDiscovery}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GroupConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to GroupConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GroupConfig<PSCustomObject>
#>
function ConvertFrom-JsonToGroupConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GroupConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GroupConfig
        $AllProperties = ("dnsIntegrationId", "configCmdbId", "configCmId", "serviceRegistryId", "configManagementId", "configCmdbDiscovery")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dnsIntegrationId"))) { #optional property not found
            $DnsIntegrationId = $null
        } else {
            $DnsIntegrationId = $JsonParameters.PSobject.Properties["dnsIntegrationId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configCmdbId"))) { #optional property not found
            $ConfigCmdbId = $null
        } else {
            $ConfigCmdbId = $JsonParameters.PSobject.Properties["configCmdbId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configCmId"))) { #optional property not found
            $ConfigCmId = $null
        } else {
            $ConfigCmId = $JsonParameters.PSobject.Properties["configCmId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "serviceRegistryId"))) { #optional property not found
            $ServiceRegistryId = $null
        } else {
            $ServiceRegistryId = $JsonParameters.PSobject.Properties["serviceRegistryId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configManagementId"))) { #optional property not found
            $ConfigManagementId = $null
        } else {
            $ConfigManagementId = $JsonParameters.PSobject.Properties["configManagementId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "configCmdbDiscovery"))) { #optional property not found
            $ConfigCmdbDiscovery = $null
        } else {
            $ConfigCmdbDiscovery = $JsonParameters.PSobject.Properties["configCmdbDiscovery"].value
        }

        $PSO = [PSCustomObject]@{
            "dnsIntegrationId" = ${DnsIntegrationId}
            "configCmdbId" = ${ConfigCmdbId}
            "configCmId" = ${ConfigCmId}
            "serviceRegistryId" = ${ServiceRegistryId}
            "configManagementId" = ${ConfigManagementId}
            "configCmdbDiscovery" = ${ConfigCmdbDiscovery}
        }

        return $PSO
    }

}

