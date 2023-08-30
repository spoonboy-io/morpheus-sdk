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

.PARAMETER BillingInfo
No description available.
.OUTPUTS

GetBillingInstancesIdentifier200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetBillingInstancesIdentifier200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${BillingInfo}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetBillingInstancesIdentifier200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "billingInfo" = ${BillingInfo}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetBillingInstancesIdentifier200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetBillingInstancesIdentifier200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetBillingInstancesIdentifier200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetBillingInstancesIdentifier200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetBillingInstancesIdentifier200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetBillingInstancesIdentifier200ResponseAllOf
        $AllProperties = ("billingInfo")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "billingInfo"))) { #optional property not found
            $BillingInfo = $null
        } else {
            $BillingInfo = $JsonParameters.PSobject.Properties["billingInfo"].value
        }

        $PSO = [PSCustomObject]@{
            "billingInfo" = ${BillingInfo}
        }

        return $PSO
    }

}

