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

.PARAMETER Name
Name
.PARAMETER Description
Description
.PARAMETER MonitorType
No description available.
.PARAMETER MonitorTimeout
No description available.
.PARAMETER Config
Configuration object with parameters that vary by type.
.OUTPUTS

ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor<PSCustomObject>
#>

function Initialize-ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Description},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MonitorType},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${MonitorTimeout},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "monitorType" = ${MonitorType}
            "monitorTimeout" = ${MonitorTimeout}
            "config" = ${Config}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor<PSCustomObject>

.DESCRIPTION

Convert from JSON to ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor<PSCustomObject>
#>
function ConvertFrom-JsonToApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor
        $AllProperties = ("name", "description", "monitorType", "monitorTimeout", "config")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "description"))) { #optional property not found
            $Description = $null
        } else {
            $Description = $JsonParameters.PSobject.Properties["description"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "monitorType"))) { #optional property not found
            $MonitorType = $null
        } else {
            $MonitorType = $JsonParameters.PSobject.Properties["monitorType"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "monitorTimeout"))) { #optional property not found
            $MonitorTimeout = $null
        } else {
            $MonitorTimeout = $JsonParameters.PSobject.Properties["monitorTimeout"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) { #optional property not found
            $Config = $null
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        $PSO = [PSCustomObject]@{
            "name" = ${Name}
            "description" = ${Description}
            "monitorType" = ${MonitorType}
            "monitorTimeout" = ${MonitorTimeout}
            "config" = ${Config}
        }

        return $PSO
    }

}
