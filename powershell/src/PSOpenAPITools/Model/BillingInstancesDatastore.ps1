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
.PARAMETER Name
No description available.
.PARAMETER Type
No description available.
.PARAMETER ExternalId
No description available.
.PARAMETER InternalId
No description available.
.PARAMETER ExternalPath
No description available.
.OUTPUTS

BillingInstancesDatastore<PSCustomObject>
#>

function Initialize-BillingInstancesDatastore {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Id},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalId},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${InternalId},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${ExternalPath}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingInstancesDatastore' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "type" = ${Type}
            "externalId" = ${ExternalId}
            "internalId" = ${InternalId}
            "externalPath" = ${ExternalPath}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingInstancesDatastore<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingInstancesDatastore<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingInstancesDatastore<PSCustomObject>
#>
function ConvertFrom-JsonToBillingInstancesDatastore {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingInstancesDatastore' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingInstancesDatastore
        $AllProperties = ("id", "name", "type", "externalId", "internalId", "externalPath")
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

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) { #optional property not found
            $Name = $null
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) { #optional property not found
            $Type = $null
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalId"))) { #optional property not found
            $ExternalId = $null
        } else {
            $ExternalId = $JsonParameters.PSobject.Properties["externalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "internalId"))) { #optional property not found
            $InternalId = $null
        } else {
            $InternalId = $JsonParameters.PSobject.Properties["internalId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "externalPath"))) { #optional property not found
            $ExternalPath = $null
        } else {
            $ExternalPath = $JsonParameters.PSobject.Properties["externalPath"].value
        }

        $PSO = [PSCustomObject]@{
            "id" = ${Id}
            "name" = ${Name}
            "type" = ${Type}
            "externalId" = ${ExternalId}
            "internalId" = ${InternalId}
            "externalPath" = ${ExternalPath}
        }

        return $PSO
    }

}
