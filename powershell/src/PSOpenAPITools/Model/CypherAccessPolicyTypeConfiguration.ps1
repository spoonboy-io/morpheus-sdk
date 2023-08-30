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

Configuration settings for the following policy types: - Cypher Access 

.PARAMETER KeyPattern
No description available.
.PARAMETER Read
No description available.
.PARAMETER Write
No description available.
.PARAMETER Update
No description available.
.PARAMETER Delete
No description available.
.PARAMETER List
No description available.
.OUTPUTS

CypherAccessPolicyTypeConfiguration<PSCustomObject>
#>

function Initialize-CypherAccessPolicyTypeConfiguration {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${KeyPattern},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Read},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Write},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Update},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${Delete},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Boolean]]
        ${List}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CypherAccessPolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "keyPattern" = ${KeyPattern}
            "read" = ${Read}
            "write" = ${Write}
            "update" = ${Update}
            "delete" = ${Delete}
            "list" = ${List}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CypherAccessPolicyTypeConfiguration<PSCustomObject>

.DESCRIPTION

Convert from JSON to CypherAccessPolicyTypeConfiguration<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CypherAccessPolicyTypeConfiguration<PSCustomObject>
#>
function ConvertFrom-JsonToCypherAccessPolicyTypeConfiguration {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CypherAccessPolicyTypeConfiguration' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CypherAccessPolicyTypeConfiguration
        $AllProperties = ("keyPattern", "read", "write", "update", "delete", "list")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "keyPattern"))) { #optional property not found
            $KeyPattern = $null
        } else {
            $KeyPattern = $JsonParameters.PSobject.Properties["keyPattern"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "read"))) { #optional property not found
            $Read = $null
        } else {
            $Read = $JsonParameters.PSobject.Properties["read"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "write"))) { #optional property not found
            $Write = $null
        } else {
            $Write = $JsonParameters.PSobject.Properties["write"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "update"))) { #optional property not found
            $Update = $null
        } else {
            $Update = $JsonParameters.PSobject.Properties["update"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "delete"))) { #optional property not found
            $Delete = $null
        } else {
            $Delete = $JsonParameters.PSobject.Properties["delete"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "list"))) { #optional property not found
            $List = $null
        } else {
            $List = $JsonParameters.PSobject.Properties["list"].value
        }

        $PSO = [PSCustomObject]@{
            "keyPattern" = ${KeyPattern}
            "read" = ${Read}
            "write" = ${Write}
            "update" = ${Update}
            "delete" = ${Delete}
            "list" = ${List}
        }

        return $PSO
    }

}

