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

.PARAMETER Id
No description available.
.PARAMETER ItemKey
No description available.
.PARAMETER LeaseTimeout
No description available.
.PARAMETER ExpireDate
No description available.
.PARAMETER DateCreated
No description available.
.PARAMETER LastUpdated
No description available.
.PARAMETER LastAccessed
No description available.
.PARAMETER CreatedBy
No description available.
.OUTPUTS

Cypher<PSCustomObject>
#>

function Initialize-Cypher {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int32]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ItemKey},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${LeaseTimeout},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${ExpireDate},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${DateCreated},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastUpdated},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[System.DateTime]]
        ${LastAccessed},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${CreatedBy}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => Cypher' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "itemKey" = ${ItemKey}
            "leaseTimeout" = ${LeaseTimeout}
            "expireDate" = ${ExpireDate}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "lastAccessed" = ${LastAccessed}
            "createdBy" = ${CreatedBy}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to Cypher<PSCustomObject>

.DESCRIPTION

Convert from JSON to Cypher<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

Cypher<PSCustomObject>
#>
function ConvertFrom-JsonToCypher {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => Cypher' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in Cypher
        $AllProperties = ("id", "itemKey", "leaseTimeout", "expireDate", "dateCreated", "lastUpdated", "lastAccessed", "createdBy")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "id"))) { #optional property not found
            $Id = $null
        } else {
            $Id = $JsonParameters.PSobject.Properties["id"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "itemKey"))) { #optional property not found
            $ItemKey = $null
        } else {
            $ItemKey = $JsonParameters.PSobject.Properties["itemKey"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "leaseTimeout"))) { #optional property not found
            $LeaseTimeout = $null
        } else {
            $LeaseTimeout = $JsonParameters.PSobject.Properties["leaseTimeout"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "expireDate"))) { #optional property not found
            $ExpireDate = $null
        } else {
            $ExpireDate = $JsonParameters.PSobject.Properties["expireDate"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "dateCreated"))) { #optional property not found
            $DateCreated = $null
        } else {
            $DateCreated = $JsonParameters.PSobject.Properties["dateCreated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastUpdated"))) { #optional property not found
            $LastUpdated = $null
        } else {
            $LastUpdated = $JsonParameters.PSobject.Properties["lastUpdated"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "lastAccessed"))) { #optional property not found
            $LastAccessed = $null
        } else {
            $LastAccessed = $JsonParameters.PSobject.Properties["lastAccessed"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "createdBy"))) { #optional property not found
            $CreatedBy = $null
        } else {
            $CreatedBy = $JsonParameters.PSobject.Properties["createdBy"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "itemKey" = ${ItemKey}
            "leaseTimeout" = ${LeaseTimeout}
            "expireDate" = ${ExpireDate}
            "dateCreated" = ${DateCreated}
            "lastUpdated" = ${LastUpdated}
            "lastAccessed" = ${LastAccessed}
            "createdBy" = ${CreatedBy}
        }

        return $PSO
    }

}

