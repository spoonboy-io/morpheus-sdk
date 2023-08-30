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

.PARAMETER SecurityScan
No description available.
.OUTPUTS

GetSecurityScans200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetSecurityScans200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${SecurityScan}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetSecurityScans200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "securityScan" = ${SecurityScan}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetSecurityScans200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetSecurityScans200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetSecurityScans200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetSecurityScans200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetSecurityScans200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetSecurityScans200ResponseAllOf
        $AllProperties = ("securityScan")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "securityScan"))) { #optional property not found
            $SecurityScan = $null
        } else {
            $SecurityScan = $JsonParameters.PSobject.Properties["securityScan"].value
        }

        $PSO = [PSCustomObject]@{
            "securityScan" = ${SecurityScan}
        }

        return $PSO
    }

}
