# MorpheusApi.RolePermissionFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**permissionCode** | **String** | The code of the feature permission being changed | 
**access** | **String** | The new access level. | 



## Enum: PermissionCodeEnum


* `account-usage` (value: `"account-usage"`)

* `activity` (value: `"activity"`)

* `admin-accounts` (value: `"admin-accounts"`)

* `admin-accounts-users` (value: `"admin-accounts-users"`)

* `admin-appliance` (value: `"admin-appliance"`)

* `admin-backupSettings` (value: `"admin-backupSettings"`)

* `admin-certificates` (value: `"admin-certificates"`)

* `admin-clients` (value: `"admin-clients"`)

* `admin-cm` (value: `"admin-cm"`)

* `admin-containers` (value: `"admin-containers"`)

* `admin-distributed-workers` (value: `"admin-distributed-workers"`)

* `admin-environments` (value: `"admin-environments"`)

* `admin-global-policies` (value: `"admin-global-policies"`)

* `admin-groups` (value: `"admin-groups"`)

* `admin-guidanceSettings` (value: `"admin-guidanceSettings"`)

* `admin-health` (value: `"admin-health"`)

* `admin-identity-sources` (value: `"admin-identity-sources"`)

* `admin-keypairs` (value: `"admin-keypairs"`)

* `admin-licenses` (value: `"admin-licenses"`)

* `admin-logSettings` (value: `"admin-logSettings"`)

* `admin-monitorSettings` (value: `"admin-monitorSettings"`)

* `admin-motd` (value: `"admin-motd"`)

* `admin-packages` (value: `"admin-packages"`)

* `admin-plugins` (value: `"admin-plugins"`)

* `admin-policies` (value: `"admin-policies"`)

* `admin-profiles` (value: `"admin-profiles"`)

* `admin-provisioningSettings` (value: `"admin-provisioningSettings"`)

* `admin-roles` (value: `"admin-roles"`)

* `admin-servers` (value: `"admin-servers"`)

* `admin-servicePlans` (value: `"admin-servicePlans"`)

* `admin-users` (value: `"admin-users"`)

* `admin-whitelabel` (value: `"admin-whitelabel"`)

* `admin-zones` (value: `"admin-zones"`)

* `app-templates` (value: `"app-templates"`)

* `apps` (value: `"apps"`)

* `arm-template` (value: `"arm-template"`)

* `automation-services` (value: `"automation-services"`)

* `backup-services` (value: `"backup-services"`)

* `backups` (value: `"backups"`)

* `billing` (value: `"billing"`)

* `catalog` (value: `"catalog"`)

* `cloudFormation-template` (value: `"cloudFormation-template"`)

* `code-repositories` (value: `"code-repositories"`)

* `credentials` (value: `"credentials"`)

* `dashboard` (value: `"dashboard"`)

* `deployment-services` (value: `"deployment-services"`)

* `deployments` (value: `"deployments"`)

* `execution-request` (value: `"execution-request"`)

* `executions` (value: `"executions"`)

* `guidance` (value: `"guidance"`)

* `helm-template` (value: `"helm-template"`)

* `infrastructure-boot` (value: `"infrastructure-boot"`)

* `infrastructure-cluster` (value: `"infrastructure-cluster"`)

* `infrastructure-dhcp-pool` (value: `"infrastructure-dhcp-pool"`)

* `infrastructure-domains` (value: `"infrastructure-domains"`)

* `infrastructure-ippools` (value: `"infrastructure-ippools"`)

* `infrastructure-kube-cntl` (value: `"infrastructure-kube-cntl"`)

* `infrastructure-loadbalancer` (value: `"infrastructure-loadbalancer"`)

* `infrastructure-move-server` (value: `"infrastructure-move-server"`)

* `infrastructure-nat` (value: `"infrastructure-nat"`)

* `infrastructure-network-dhcp-relay` (value: `"infrastructure-network-dhcp-relay"`)

* `infrastructure-network-dhcp-routes` (value: `"infrastructure-network-dhcp-routes"`)

* `infrastructure-network-dhcp-server` (value: `"infrastructure-network-dhcp-server"`)

* `infrastructure-network-firewalls` (value: `"infrastructure-network-firewalls"`)

* `infrastructure-network-integrations` (value: `"infrastructure-network-integrations"`)

* `infrastructure-network-router-firewalls` (value: `"infrastructure-network-router-firewalls"`)

* `infrastructure-network-router-interfaces` (value: `"infrastructure-network-router-interfaces"`)

* `infrastructure-network-router-redistribution` (value: `"infrastructure-network-router-redistribution"`)

* `infrastructure-network-router-routes` (value: `"infrastructure-network-router-routes"`)

* `infrastructure-network-server-groups` (value: `"infrastructure-network-server-groups"`)

* `infrastructure-networks` (value: `"infrastructure-networks"`)

* `infrastructure-proxies` (value: `"infrastructure-proxies"`)

* `infrastructure-router-dhcp-binding` (value: `"infrastructure-router-dhcp-binding"`)

* `infrastructure-router-dhcp-relay` (value: `"infrastructure-router-dhcp-relay"`)

* `infrastructure-routers` (value: `"infrastructure-routers"`)

* `infrastructure-securityGroups` (value: `"infrastructure-securityGroups"`)

* `infrastructure-state` (value: `"infrastructure-state"`)

* `infrastructure-storage` (value: `"infrastructure-storage"`)

* `infrastructure-storage-browser` (value: `"infrastructure-storage-browser"`)

* `integrations-ansible` (value: `"integrations-ansible"`)

* `job-executions` (value: `"job-executions"`)

* `job-templates` (value: `"job-templates"`)

* `kubernetes-template` (value: `"kubernetes-template"`)

* `library-advanced-node-type-options` (value: `"library-advanced-node-type-options"`)

* `library-options` (value: `"library-options"`)

* `library-templates` (value: `"library-templates"`)

* `logs` (value: `"logs"`)

* `monitoring` (value: `"monitoring"`)

* `operations-alarms` (value: `"operations-alarms"`)

* `operations-approvals` (value: `"operations-approvals"`)

* `operations-budgets` (value: `"operations-budgets"`)

* `operations-invoices` (value: `"operations-invoices"`)

* `operations-wiki` (value: `"operations-wiki"`)

* `projects` (value: `"projects"`)

* `provisioning` (value: `"provisioning"`)

* `provisioning-add` (value: `"provisioning-add"`)

* `provisioning-admin` (value: `"provisioning-admin"`)

* `provisioning-clone` (value: `"provisioning-clone"`)

* `provisioning-delete` (value: `"provisioning-delete"`)

* `provisioning-edit` (value: `"provisioning-edit"`)

* `provisioning-environment` (value: `"provisioning-environment"`)

* `provisioning-execute-script` (value: `"provisioning-execute-script"`)

* `provisioning-execute-task` (value: `"provisioning-execute-task"`)

* `provisioning-execute-workflow` (value: `"provisioning-execute-workflow"`)

* `provisioning-force-delete` (value: `"provisioning-force-delete"`)

* `provisioning-import-image` (value: `"provisioning-import-image"`)

* `provisioning-lock` (value: `"provisioning-lock"`)

* `provisioning-power` (value: `"provisioning-power"`)

* `provisioning-reconfigure` (value: `"provisioning-reconfigure"`)

* `provisioning-reconfigure-add-disk` (value: `"provisioning-reconfigure-add-disk"`)

* `provisioning-reconfigure-add-network` (value: `"provisioning-reconfigure-add-network"`)

* `provisioning-reconfigure-change-plan` (value: `"provisioning-reconfigure-change-plan"`)

* `provisioning-reconfigure-disk-type` (value: `"provisioning-reconfigure-disk-type"`)

* `provisioning-reconfigure-modify-disk` (value: `"provisioning-reconfigure-modify-disk"`)

* `provisioning-reconfigure-modify-network` (value: `"provisioning-reconfigure-modify-network"`)

* `provisioning-reconfigure-remove-disk` (value: `"provisioning-reconfigure-remove-disk"`)

* `provisioning-reconfigure-remove-network` (value: `"provisioning-reconfigure-remove-network"`)

* `provisioning-remove-control` (value: `"provisioning-remove-control"`)

* `provisioning-scale` (value: `"provisioning-scale"`)

* `provisioning-settings` (value: `"provisioning-settings"`)

* `provisioning-state` (value: `"provisioning-state"`)

* `reports` (value: `"reports"`)

* `reports-analytics` (value: `"reports-analytics"`)

* `scheduling-execute` (value: `"scheduling-execute"`)

* `scheduling-power` (value: `"scheduling-power"`)

* `security-scan` (value: `"security-scan"`)

* `service-catalog` (value: `"service-catalog"`)

* `service-catalog-dashboard` (value: `"service-catalog-dashboard"`)

* `service-catalog-inventory` (value: `"service-catalog-inventory"`)

* `services-archives` (value: `"services-archives"`)

* `services-cypher` (value: `"services-cypher"`)

* `services-image-builder` (value: `"services-image-builder"`)

* `services-kubernetes` (value: `"services-kubernetes"`)

* `services-network-registry` (value: `"services-network-registry"`)

* `services-vdi-copy` (value: `"services-vdi-copy"`)

* `services-vdi-pools` (value: `"services-vdi-pools"`)

* `services-vdi-printer` (value: `"services-vdi-printer"`)

* `snapshots` (value: `"snapshots"`)

* `task-scripts` (value: `"task-scripts"`)

* `tasks` (value: `"tasks"`)

* `terminal` (value: `"terminal"`)

* `terminal-access` (value: `"terminal-access"`)

* `terraform-template` (value: `"terraform-template"`)

* `thresholds` (value: `"thresholds"`)

* `trust-services` (value: `"trust-services"`)

* `virtual-images` (value: `"virtual-images"`)





## Enum: AccessEnum


* `full` (value: `"full"`)

* `full_decrypted` (value: `"full_decrypted"`)

* `group` (value: `"group"`)

* `listfiles` (value: `"listfiles"`)

* `managerules` (value: `"managerules"`)

* `no` (value: `"no"`)

* `none` (value: `"none"`)

* `provision` (value: `"provision"`)

* `read` (value: `"read"`)

* `rolemappings` (value: `"rolemappings"`)

* `user` (value: `"user"`)

* `view` (value: `"view"`)

* `yes` (value: `"yes"`)




