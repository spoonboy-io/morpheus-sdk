#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.1.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER Size
No description available.
.PARAMETER TypeCode
No description available.
.PARAMETER Datastore
No description available.
.OUTPUTS

BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner<PSCustomObject>
#>

function Initialize-BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${Size},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TypeCode},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Datastore}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug


        $PSO = [PSCustomObject]@{
            "size" = ${Size}
            "typeCode" = ${TypeCode}
            "datastore" = ${Datastore}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner<PSCustomObject>
#>
function ConvertFrom-JsonToBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner
        $AllProperties = ("size", "typeCode", "datastore")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "size"))) { #optional property not found
            $Size = $null
        } else {
            $Size = $JsonParameters.PSobject.Properties["size"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "typeCode"))) { #optional property not found
            $TypeCode = $null
        } else {
            $TypeCode = $JsonParameters.PSobject.Properties["typeCode"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "datastore"))) { #optional property not found
            $Datastore = $null
        } else {
            $Datastore = $JsonParameters.PSobject.Properties["datastore"].value
        }

        $PSO = [PSCustomObject]@{
            "size" = ${Size}
            "typeCode" = ${TypeCode}
            "datastore" = ${Datastore}
        }

        return $PSO
    }

}

