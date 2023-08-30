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

.PARAMETER Code
`code` of the feature permission
.PARAMETER Access
The new access level.
.OUTPUTS

AddRolesRequestRolePermissionsInner<PSCustomObject>
#>

function Initialize-AddRolesRequestRolePermissionsInner {
    [CmdletBinding()]
    Param (
        [Parameter(Position = 0, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("account-usage", "activity", "admin-accounts", "admin-accounts-users", "admin-appliance", "admin-backupSettings", "admin-certificates", "admin-clients", "admin-cm", "admin-containers", "admin-distributed-workers", "admin-environments", "admin-global-policies", "admin-groups", "admin-guidanceSettings", "admin-health", "admin-identity-sources", "admin-keypairs", "admin-licenses", "admin-logSettings", "admin-monitorSettings", "admin-motd", "admin-packages", "admin-plugins", "admin-policies", "admin-profiles", "admin-provisioningSettings", "admin-roles", "admin-servers", "admin-servicePlans", "admin-users", "admin-whitelabel", "admin-zones", "app-templates", "apps", "arm-template", "automation-services", "backup-services", "backups", "billing", "catalog", "cloudFormation-template", "code-repositories", "credentials", "dashboard", "deployment-services", "deployments", "execution-request", "executions", "guidance", "helm-template", "infrastructure-boot", "infrastructure-cluster", "infrastructure-dhcp-pool", "infrastructure-domains", "infrastructure-ippools", "infrastructure-kube-cntl", "infrastructure-loadbalancer", "infrastructure-move-server", "infrastructure-nat", "infrastructure-network-dhcp-relay", "infrastructure-network-dhcp-routes", "infrastructure-network-dhcp-server", "infrastructure-network-firewalls", "infrastructure-network-integrations", "infrastructure-network-router-firewalls", "infrastructure-network-router-interfaces", "infrastructure-network-router-redistribution", "infrastructure-network-router-routes", "infrastructure-network-server-groups", "infrastructure-networks", "infrastructure-proxies", "infrastructure-router-dhcp-binding", "infrastructure-router-dhcp-relay", "infrastructure-routers", "infrastructure-securityGroups", "infrastructure-state", "infrastructure-storage", "infrastructure-storage-browser", "integrations-ansible", "job-executions", "job-templates", "kubernetes-template", "library-advanced-node-type-options", "library-options", "library-templates", "logs", "monitoring", "operations-alarms", "operations-approvals", "operations-budgets", "operations-invoices", "operations-wiki", "projects", "provisioning", "provisioning-add", "provisioning-admin", "provisioning-clone", "provisioning-delete", "provisioning-edit", "provisioning-environment", "provisioning-execute-script", "provisioning-execute-task", "provisioning-execute-workflow", "provisioning-force-delete", "provisioning-import-image", "provisioning-lock", "provisioning-power", "provisioning-reconfigure", "provisioning-reconfigure-add-disk", "provisioning-reconfigure-add-network", "provisioning-reconfigure-change-plan", "provisioning-reconfigure-disk-type", "provisioning-reconfigure-modify-disk", "provisioning-reconfigure-modify-network", "provisioning-reconfigure-remove-disk", "provisioning-reconfigure-remove-network", "provisioning-remove-control", "provisioning-scale", "provisioning-settings", "provisioning-state", "reports", "reports-analytics", "scheduling-execute", "scheduling-power", "security-scan", "service-catalog", "service-catalog-dashboard", "service-catalog-inventory", "services-archives", "services-cypher", "services-image-builder", "services-kubernetes", "services-network-registry", "services-vdi-copy", "services-vdi-pools", "services-vdi-printer", "snapshots", "task-scripts", "tasks", "terminal", "terminal-access", "terraform-template", "thresholds", "trust-services", "virtual-images")]
        [String]
        ${Code},
        [Parameter(Position = 1, ValueFromPipelineByPropertyName = $true)]
        [ValidateSet("full", "full_decrypted", "group", "listfiles", "managerules", "no", "none", "provision", "read", "rolemappings", "user", "view", "yes")]
        [String]
        ${Access}
    )

    Process {
        'Creating PSCustomObject: PSOpenAPITools => AddRolesRequestRolePermissionsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        if ($null -eq $Code) {
            throw "invalid value for 'Code', 'Code' cannot be null."
        }

        if ($null -eq $Access) {
            throw "invalid value for 'Access', 'Access' cannot be null."
        }


        $PSO = [PSCustomObject]@{
            "code" = ${Code}
            "access" = ${Access}
        }


        return $PSO
    }
}

<#
.SYNOPSIS

Convert from JSON to AddRolesRequestRolePermissionsInner<PSCustomObject>

.DESCRIPTION

Convert from JSON to AddRolesRequestRolePermissionsInner<PSCustomObject>

.PARAMETER Json

Json object

.OUTPUTS

AddRolesRequestRolePermissionsInner<PSCustomObject>
#>
function ConvertFrom-JsonToAddRolesRequestRolePermissionsInner {
    Param(
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        'Converting JSON to PSCustomObject: PSOpenAPITools => AddRolesRequestRolePermissionsInner' | Write-Debug
        $PSBoundParameters | Out-DebugParameter | Write-Debug

        $JsonParameters = ConvertFrom-Json -InputObject $Json

        # check if Json contains properties not defined in AddRolesRequestRolePermissionsInner
        $AllProperties = ("code", "access")
        foreach ($name in $JsonParameters.PsObject.Properties.Name) {
            if (!($AllProperties.Contains($name))) {
                throw "Error! JSON key '$name' not found in the properties: $($AllProperties)"
            }
        }

        If ([string]::IsNullOrEmpty($Json) -or $Json -eq "{}") { # empty json
            throw "Error! Empty JSON cannot be serialized due to the required property 'code' missing."
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "code"))) {
            throw "Error! JSON cannot be serialized due to the required property 'code' missing."
        } else {
            $Code = $JsonParameters.PSobject.Properties["code"].value
        }

        if (!([bool]($JsonParameters.PSobject.Properties.name -match "access"))) {
            throw "Error! JSON cannot be serialized due to the required property 'access' missing."
        } else {
            $Access = $JsonParameters.PSobject.Properties["access"].value
        }

        $PSO = [PSCustomObject]@{
            "code" = ${Code}
            "access" = ${Access}
        }

        return $PSO
    }

}

