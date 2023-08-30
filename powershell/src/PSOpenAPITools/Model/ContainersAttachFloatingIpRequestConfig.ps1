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

.PARAMETER OsExternalNetworkId
The Floating IP identifier in the format: ""ip-ID"" or ""pool-ID"".  The Options API /api/options/openStack/openstackFloatingIpOptions?containerId={{containerId}} can be used to see which options are available. 
.PARAMETER FloatingIpBandwidth
Bandwidth (Mbit/s) Only the following cloud types support this parameter: Huawei, OpenTelekom 
.OUTPUTS

ContainersAttachFloatingIpRequestConfig<PSCustomObject>
#>

function Initialize-ContainersAttachFloatingIpRequestConfig {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${OsExternalNetworkId},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${FloatingIpBandwidth}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ContainersAttachFloatingIpRequestConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $OsExternalNetworkId) {
            throw "invalid value for 'OsExternalNetworkId', 'OsExternalNetworkId' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "osExternalNetworkId" = ${OsExternalNetworkId}
            "floatingIpBandwidth" = ${FloatingIpBandwidth}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ContainersAttachFloatingIpRequestConfig<PSCustomObject>

.DESCRIPTION

Convert from JSON to ContainersAttachFloatingIpRequestConfig<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ContainersAttachFloatingIpRequestConfig<PSCustomObject>
#>
function ConvertFrom-JsonToContainersAttachFloatingIpRequestConfig {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ContainersAttachFloatingIpRequestConfig' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ContainersAttachFloatingIpRequestConfig
        $AllProperties = ("osExternalNetworkId", "floatingIpBandwidth")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'osExternalNetworkId' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "osExternalNetworkId"))) {
            throw "Error! JSON cannot be serialized due to the required property 'osExternalNetworkId' missing."
        } else {
            $OsExternalNetworkId = $JsonParameters.PSobject.Properties["osExternalNetworkId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "floatingIpBandwidth"))) { #optional property not found
            $FloatingIpBandwidth = $null
        } else {
            $FloatingIpBandwidth = $JsonParameters.PSobject.Properties["floatingIpBandwidth"].value
        }

        $PSO = [PSCustomObject]@{
            "osExternalNetworkId" = ${OsExternalNetworkId}
            "floatingIpBandwidth" = ${FloatingIpBandwidth}
        }

        return $PSO
    }

}

