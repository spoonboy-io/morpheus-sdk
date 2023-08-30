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

.PARAMETER LineItem
No description available.
.PARAMETER MasterAccount
No description available.
.OUTPUTS

GetInvoiceLineItems200ResponseAllOf<PSCustomObject>
#>

function Initialize-GetInvoiceLineItems200ResponseAllOf {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${LineItem},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${MasterAccount}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => GetInvoiceLineItems200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "lineItem" = ${LineItem}
            "masterAccount" = ${MasterAccount}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to GetInvoiceLineItems200ResponseAllOf<PSCustomObject>

.DESCRIPTION

Convert from JSON to GetInvoiceLineItems200ResponseAllOf<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

GetInvoiceLineItems200ResponseAllOf<PSCustomObject>
#>
function ConvertFrom-JsonToGetInvoiceLineItems200ResponseAllOf {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => GetInvoiceLineItems200ResponseAllOf' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in GetInvoiceLineItems200ResponseAllOf
        $AllProperties = ("lineItem", "masterAccount")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lineItem"))) { #optional property not found
            $LineItem = $null
        } else {
            $LineItem = $JsonParameters.PSobject.Properties["lineItem"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "masterAccount"))) { #optional property not found
            $MasterAccount = $null
        } else {
            $MasterAccount = $JsonParameters.PSobject.Properties["masterAccount"].value
        }

        $PSO = [PSCustomObject]@{
            "lineItem" = ${LineItem}
            "masterAccount" = ${MasterAccount}
        }

        return $PSO
    }

}

