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

.PARAMETER Group
No description available.
.PARAMETER Cloud
No description available.
.PARAMETER Type
The type of instance by code we want to fetch.
.PARAMETER Name
Name of the instance to be created.
.PARAMETER Config
No description available.
.PARAMETER Volumes
The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of
.PARAMETER HostName
Hostname of the instance to be created.  Can be the same as the instance name.
.PARAMETER Environment
Environment code
.PARAMETER Layout
No description available.
.PARAMETER Plan
No description available.
.PARAMETER Version
Version of the layout to create.
.PARAMETER Evars
Environment Variables, an array of objects that have name and value.
.PARAMETER ServicePlanOptions
No description available.
.PARAMETER SecurityGroups
Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to.
.PARAMETER NetworkInterfaces
The networkInterfaces parameter is for network configuration.  The Options API `/api/options/zoneNetworkOptions?zoneId=5&provisionTypeId=10` can be used to see which options are available. 
.PARAMETER Labels
Array of strings (keywords).
.PARAMETER Tags
Metadata tags, Array of objects having a name and value.
.PARAMETER Metadata
Alias for `tags`.
.PARAMETER Ports
The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened. 
.PARAMETER TaskSetId
The Workflow ID to execute.
.PARAMETER TaskSetName
The Workflow Name to execute.
.OUTPUTS

CatalogItemTypeInstanceScribe<PSCustomObject>
#>

function Initialize-CatalogItemTypeInstanceScribe {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Group},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Cloud},
        [Parameter(Position = 2, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Type},
        [Parameter(Position = 3, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Name},
        [Parameter(Position = 4, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Config},
        [Parameter(Position = 5, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Volumes},
        [Parameter(Position = 6, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${HostName},
        [Parameter(Position = 7, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Environment},
        [Parameter(Position = 8, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Layout},
        [Parameter(Position = 9, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${Plan},
        [Parameter(Position = 10, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${Version},
        [Parameter(Position = 11, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Evars},
        [Parameter(Position = 12, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject]
        ${ServicePlanOptions},
        [Parameter(Position = 13, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${SecurityGroups},
        [Parameter(Position = 14, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${NetworkInterfaces},
        [Parameter(Position = 15, ValueFromPipelineByPropertyName = $true)]
        [String[]]
        ${Labels},
        [Parameter(Position = 16, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Tags},
        [Parameter(Position = 17, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Metadata},
        [Parameter(Position = 18, ValueFromPipelineByPropertyName = $true)]
        [PSCustomObject[]]
        ${Ports},
        [Parameter(Position = 19, ValueFromPipelineByPropertyName = $true)]
        [System.Nullable[Int64]]
        ${TaskSetId},
        [Parameter(Position = 20, ValueFromPipelineByPropertyName = $true)]
        [String]
        ${TaskSetName}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => CatalogItemTypeInstanceScribe' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if (!$Group) {
            throw "invalid value for 'Group', 'Group' cannot be null."
        }

        if (!$Cloud) {
            throw "invalid value for 'Cloud', 'Cloud' cannot be null."
        }

        if (!$Type) {
            throw "invalid value for 'Type', 'Type' cannot be null."
        }

        if (!$Name) {
            throw "invalid value for 'Name', 'Name' cannot be null."
        }

        if (!$Config) {
            throw "invalid value for 'Config', 'Config' cannot be null."
        }

        if (!$Volumes) {
            throw "invalid value for 'Volumes', 'Volumes' cannot be null."
        }

        if (!$Layout) {
            throw "invalid value for 'Layout', 'Layout' cannot be null."
        }

        if (!$Plan) {
            throw "invalid value for 'Plan', 'Plan' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "group" = ${Group}
            "cloud" = ${Cloud}
            "type" = ${Type}
            "name" = ${Name}
            "config" = ${Config}
            "volumes" = ${Volumes}
            "hostName" = ${HostName}
            "environment" = ${Environment}
            "layout" = ${Layout}
            "plan" = ${Plan}
            "version" = ${Version}
            "evars" = ${Evars}
            "servicePlanOptions" = ${ServicePlanOptions}
            "securityGroups" = ${SecurityGroups}
            "networkInterfaces" = ${NetworkInterfaces}
            "labels" = ${Labels}
            "tags" = ${Tags}
            "metadata" = ${Metadata}
            "ports" = ${Ports}
            "taskSetId" = ${TaskSetId}
            "taskSetName" = ${TaskSetName}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to CatalogItemTypeInstanceScribe<PSCustomObject>

.DESCRIPTION

Convert from JSON to CatalogItemTypeInstanceScribe<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

CatalogItemTypeInstanceScribe<PSCustomObject>
#>
function ConvertFrom-JsonToCatalogItemTypeInstanceScribe {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => CatalogItemTypeInstanceScribe' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in CatalogItemTypeInstanceScribe
        $AllProperties = ("group", "cloud", "type", "name", "config", "volumes", "hostName", "environment", "layout", "plan", "version", "evars", "servicePlanOptions", "securityGroups", "networkInterfaces", "labels", "tags", "metadata", "ports", "taskSetId", "taskSetName")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property `group` missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "group"))) {
            throw "Error! JSON cannot be serialized due to the required property `group` missing."
        } else {
            $Group = $JsonParameters.PSobject.Properties["group"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "cloud"))) {
            throw "Error! JSON cannot be serialized due to the required property `cloud` missing."
        } else {
            $Cloud = $JsonParameters.PSobject.Properties["cloud"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "type"))) {
            throw "Error! JSON cannot be serialized due to the required property `type` missing."
        } else {
            $Type = $JsonParameters.PSobject.Properties["type"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "name"))) {
            throw "Error! JSON cannot be serialized due to the required property `name` missing."
        } else {
            $Name = $JsonParameters.PSobject.Properties["name"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "config"))) {
            throw "Error! JSON cannot be serialized due to the required property `config` missing."
        } else {
            $Config = $JsonParameters.PSobject.Properties["config"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "volumes"))) {
            throw "Error! JSON cannot be serialized due to the required property `volumes` missing."
        } else {
            $Volumes = $JsonParameters.PSobject.Properties["volumes"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "layout"))) {
            throw "Error! JSON cannot be serialized due to the required property `layout` missing."
        } else {
            $Layout = $JsonParameters.PSobject.Properties["layout"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "plan"))) {
            throw "Error! JSON cannot be serialized due to the required property `plan` missing."
        } else {
            $Plan = $JsonParameters.PSobject.Properties["plan"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "hostName"))) { #optional property not found
            $HostName = $null
        } else {
            $HostName = $JsonParameters.PSobject.Properties["hostName"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "environment"))) { #optional property not found
            $Environment = $null
        } else {
            $Environment = $JsonParameters.PSobject.Properties["environment"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "version"))) { #optional property not found
            $Version = $null
        } else {
            $Version = $JsonParameters.PSobject.Properties["version"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "evars"))) { #optional property not found
            $Evars = $null
        } else {
            $Evars = $JsonParameters.PSobject.Properties["evars"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "servicePlanOptions"))) { #optional property not found
            $ServicePlanOptions = $null
        } else {
            $ServicePlanOptions = $JsonParameters.PSobject.Properties["servicePlanOptions"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "securityGroups"))) { #optional property not found
            $SecurityGroups = $null
        } else {
            $SecurityGroups = $JsonParameters.PSobject.Properties["securityGroups"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "networkInterfaces"))) { #optional property not found
            $NetworkInterfaces = $null
        } else {
            $NetworkInterfaces = $JsonParameters.PSobject.Properties["networkInterfaces"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "labels"))) { #optional property not found
            $Labels = $null
        } else {
            $Labels = $JsonParameters.PSobject.Properties["labels"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "tags"))) { #optional property not found
            $Tags = $null
        } else {
            $Tags = $JsonParameters.PSobject.Properties["tags"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "metadata"))) { #optional property not found
            $Metadata = $null
        } else {
            $Metadata = $JsonParameters.PSobject.Properties["metadata"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "ports"))) { #optional property not found
            $Ports = $null
        } else {
            $Ports = $JsonParameters.PSobject.Properties["ports"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskSetId"))) { #optional property not found
            $TaskSetId = $null
        } else {
            $TaskSetId = $JsonParameters.PSobject.Properties["taskSetId"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "taskSetName"))) { #optional property not found
            $TaskSetName = $null
        } else {
            $TaskSetName = $JsonParameters.PSobject.Properties["taskSetName"].value
        }

        $PSO = [PSCustomObject]@{
            "group" = ${Group}
            "cloud" = ${Cloud}
            "type" = ${Type}
            "name" = ${Name}
            "config" = ${Config}
            "volumes" = ${Volumes}
            "hostName" = ${HostName}
            "environment" = ${Environment}
            "layout" = ${Layout}
            "plan" = ${Plan}
            "version" = ${Version}
            "evars" = ${Evars}
            "servicePlanOptions" = ${ServicePlanOptions}
            "securityGroups" = ${SecurityGroups}
            "networkInterfaces" = ${NetworkInterfaces}
            "labels" = ${Labels}
            "tags" = ${Tags}
            "metadata" = ${Metadata}
            "ports" = ${Ports}
            "taskSetId" = ${TaskSetId}
            "taskSetName" = ${TaskSetName}
        }

        return $PSO
    }

}
