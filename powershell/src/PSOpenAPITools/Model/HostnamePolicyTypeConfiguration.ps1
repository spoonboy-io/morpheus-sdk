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

Configuration settings for the following policy types: - Hostname 

.PARAMETER HostNamingType
No description available.
.PARAMETER HostNamingPattern
No description available.
.OUTPUTS

HostnamePolicyTypeConfiguration<PSCustomObject>
#>

function Initialize-HostnamePolicyTypeConfiguration {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${HostNamingType},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${HostNamingPattern}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => HostnamePolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "hostNamingType" = ${HostNamingType}
            "hostNamingPattern" = ${HostNamingPattern}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to HostnamePolicyTypeConfiguration<PSCustomObject>

.DESCRIPTION

Convert from JSON to HostnamePolicyTypeConfiguration<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

HostnamePolicyTypeConfiguration<PSCustomObject>
#>
function ConvertFrom-JsonToHostnamePolicyTypeConfiguration {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => HostnamePolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in HostnamePolicyTypeConfiguration
        $AllProperties = ("hostNamingType", "hostNamingPattern")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostNamingType"))) { #optional property not found
            $HostNamingType = $null
        } else {
            $HostNamingType = $JsonParameters.PSobject.Properties["hostNamingType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostNamingPattern"))) { #optional property not found
            $HostNamingPattern = $null
        } else {
            $HostNamingPattern = $JsonParameters.PSobject.Properties["hostNamingPattern"].value
        }

        $PSO = [PSCustomObject]@{
            "hostNamingType" = ${HostNamingType}
            "hostNamingPattern" = ${HostNamingPattern}
        }

        return $PSO
    }

}

