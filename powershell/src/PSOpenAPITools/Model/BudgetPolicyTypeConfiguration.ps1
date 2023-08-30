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

Configuration settings for the following policy types: - Budget 

.PARAMETER MaxPrice
No description available.
.PARAMETER MaxPriceCurrency
No description available.
.PARAMETER MaxPriceUnit
No description available.
.OUTPUTS

BudgetPolicyTypeConfiguration<PSCustomObject>
#>

function Initialize-BudgetPolicyTypeConfiguration {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Decimal]]
        ${MaxPrice},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MaxPriceCurrency},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${MaxPriceUnit}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BudgetPolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "maxPrice" = ${MaxPrice}
            "maxPriceCurrency" = ${MaxPriceCurrency}
            "maxPriceUnit" = ${MaxPriceUnit}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BudgetPolicyTypeConfiguration<PSCustomObject>

.DESCRIPTION

Convert from JSON to BudgetPolicyTypeConfiguration<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BudgetPolicyTypeConfiguration<PSCustomObject>
#>
function ConvertFrom-JsonToBudgetPolicyTypeConfiguration {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BudgetPolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BudgetPolicyTypeConfiguration
        $AllProperties = ("maxPrice", "maxPriceCurrency", "maxPriceUnit")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxPrice"))) { #optional property not found
            $MaxPrice = $null
        } else {
            $MaxPrice = $JsonParameters.PSobject.Properties["maxPrice"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxPriceCurrency"))) { #optional property not found
            $MaxPriceCurrency = $null
        } else {
            $MaxPriceCurrency = $JsonParameters.PSobject.Properties["maxPriceCurrency"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "maxPriceUnit"))) { #optional property not found
            $MaxPriceUnit = $null
        } else {
            $MaxPriceUnit = $JsonParameters.PSobject.Properties["maxPriceUnit"].value
        }

        $PSO = [PSCustomObject]@{
            "maxPrice" = ${MaxPrice}
            "maxPriceCurrency" = ${MaxPriceCurrency}
            "maxPriceUnit" = ${MaxPriceUnit}
        }

        return $PSO
    }

}

