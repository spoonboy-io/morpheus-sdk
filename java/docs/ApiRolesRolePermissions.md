

# ApiRolesRolePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | [**CodeEnum**](#CodeEnum) | &#x60;code&#x60; of the feature permission | 
**access** | [**AccessEnum**](#AccessEnum) | The new access level. | 



## Enum: CodeEnum

Name | Value
---- | -----
ACCOUNT_USAGE | &quot;account-usage&quot;
ACTIVITY | &quot;activity&quot;
ADMIN_ACCOUNTS | &quot;admin-accounts&quot;
ADMIN_ACCOUNTS_USERS | &quot;admin-accounts-users&quot;
ADMIN_APPLIANCE | &quot;admin-appliance&quot;
ADMIN_BACKUPSETTINGS | &quot;admin-backupSettings&quot;
ADMIN_CERTIFICATES | &quot;admin-certificates&quot;
ADMIN_CLIENTS | &quot;admin-clients&quot;
ADMIN_CM | &quot;admin-cm&quot;
ADMIN_CONTAINERS | &quot;admin-containers&quot;
ADMIN_DISTRIBUTED_WORKERS | &quot;admin-distributed-workers&quot;
ADMIN_ENVIRONMENTS | &quot;admin-environments&quot;
ADMIN_GLOBAL_POLICIES | &quot;admin-global-policies&quot;
ADMIN_GROUPS | &quot;admin-groups&quot;
ADMIN_GUIDANCESETTINGS | &quot;admin-guidanceSettings&quot;
ADMIN_HEALTH | &quot;admin-health&quot;
ADMIN_IDENTITY_SOURCES | &quot;admin-identity-sources&quot;
ADMIN_KEYPAIRS | &quot;admin-keypairs&quot;
ADMIN_LICENSES | &quot;admin-licenses&quot;
ADMIN_LOGSETTINGS | &quot;admin-logSettings&quot;
ADMIN_MONITORSETTINGS | &quot;admin-monitorSettings&quot;
ADMIN_MOTD | &quot;admin-motd&quot;
ADMIN_PACKAGES | &quot;admin-packages&quot;
ADMIN_PLUGINS | &quot;admin-plugins&quot;
ADMIN_POLICIES | &quot;admin-policies&quot;
ADMIN_PROFILES | &quot;admin-profiles&quot;
ADMIN_PROVISIONINGSETTINGS | &quot;admin-provisioningSettings&quot;
ADMIN_ROLES | &quot;admin-roles&quot;
ADMIN_SERVERS | &quot;admin-servers&quot;
ADMIN_SERVICEPLANS | &quot;admin-servicePlans&quot;
ADMIN_USERS | &quot;admin-users&quot;
ADMIN_WHITELABEL | &quot;admin-whitelabel&quot;
ADMIN_ZONES | &quot;admin-zones&quot;
APP_TEMPLATES | &quot;app-templates&quot;
APPS | &quot;apps&quot;
ARM_TEMPLATE | &quot;arm-template&quot;
AUTOMATION_SERVICES | &quot;automation-services&quot;
BACKUP_SERVICES | &quot;backup-services&quot;
BACKUPS | &quot;backups&quot;
BILLING | &quot;billing&quot;
CATALOG | &quot;catalog&quot;
CLOUDFORMATION_TEMPLATE | &quot;cloudFormation-template&quot;
CODE_REPOSITORIES | &quot;code-repositories&quot;
CREDENTIALS | &quot;credentials&quot;
DASHBOARD | &quot;dashboard&quot;
DEPLOYMENT_SERVICES | &quot;deployment-services&quot;
DEPLOYMENTS | &quot;deployments&quot;
EXECUTION_REQUEST | &quot;execution-request&quot;
EXECUTIONS | &quot;executions&quot;
GUIDANCE | &quot;guidance&quot;
HELM_TEMPLATE | &quot;helm-template&quot;
INFRASTRUCTURE_BOOT | &quot;infrastructure-boot&quot;
INFRASTRUCTURE_CLUSTER | &quot;infrastructure-cluster&quot;
INFRASTRUCTURE_DHCP_POOL | &quot;infrastructure-dhcp-pool&quot;
INFRASTRUCTURE_DOMAINS | &quot;infrastructure-domains&quot;
INFRASTRUCTURE_IPPOOLS | &quot;infrastructure-ippools&quot;
INFRASTRUCTURE_KUBE_CNTL | &quot;infrastructure-kube-cntl&quot;
INFRASTRUCTURE_LOADBALANCER | &quot;infrastructure-loadbalancer&quot;
INFRASTRUCTURE_MOVE_SERVER | &quot;infrastructure-move-server&quot;
INFRASTRUCTURE_NAT | &quot;infrastructure-nat&quot;
INFRASTRUCTURE_NETWORK_DHCP_RELAY | &quot;infrastructure-network-dhcp-relay&quot;
INFRASTRUCTURE_NETWORK_DHCP_ROUTES | &quot;infrastructure-network-dhcp-routes&quot;
INFRASTRUCTURE_NETWORK_DHCP_SERVER | &quot;infrastructure-network-dhcp-server&quot;
INFRASTRUCTURE_NETWORK_FIREWALLS | &quot;infrastructure-network-firewalls&quot;
INFRASTRUCTURE_NETWORK_INTEGRATIONS | &quot;infrastructure-network-integrations&quot;
INFRASTRUCTURE_NETWORK_ROUTER_FIREWALLS | &quot;infrastructure-network-router-firewalls&quot;
INFRASTRUCTURE_NETWORK_ROUTER_INTERFACES | &quot;infrastructure-network-router-interfaces&quot;
INFRASTRUCTURE_NETWORK_ROUTER_REDISTRIBUTION | &quot;infrastructure-network-router-redistribution&quot;
INFRASTRUCTURE_NETWORK_ROUTER_ROUTES | &quot;infrastructure-network-router-routes&quot;
INFRASTRUCTURE_NETWORK_SERVER_GROUPS | &quot;infrastructure-network-server-groups&quot;
INFRASTRUCTURE_NETWORKS | &quot;infrastructure-networks&quot;
INFRASTRUCTURE_PROXIES | &quot;infrastructure-proxies&quot;
INFRASTRUCTURE_ROUTER_DHCP_BINDING | &quot;infrastructure-router-dhcp-binding&quot;
INFRASTRUCTURE_ROUTER_DHCP_RELAY | &quot;infrastructure-router-dhcp-relay&quot;
INFRASTRUCTURE_ROUTERS | &quot;infrastructure-routers&quot;
INFRASTRUCTURE_SECURITYGROUPS | &quot;infrastructure-securityGroups&quot;
INFRASTRUCTURE_STATE | &quot;infrastructure-state&quot;
INFRASTRUCTURE_STORAGE | &quot;infrastructure-storage&quot;
INFRASTRUCTURE_STORAGE_BROWSER | &quot;infrastructure-storage-browser&quot;
INTEGRATIONS_ANSIBLE | &quot;integrations-ansible&quot;
JOB_EXECUTIONS | &quot;job-executions&quot;
JOB_TEMPLATES | &quot;job-templates&quot;
KUBERNETES_TEMPLATE | &quot;kubernetes-template&quot;
LIBRARY_ADVANCED_NODE_TYPE_OPTIONS | &quot;library-advanced-node-type-options&quot;
LIBRARY_OPTIONS | &quot;library-options&quot;
LIBRARY_TEMPLATES | &quot;library-templates&quot;
LOGS | &quot;logs&quot;
MONITORING | &quot;monitoring&quot;
OPERATIONS_ALARMS | &quot;operations-alarms&quot;
OPERATIONS_APPROVALS | &quot;operations-approvals&quot;
OPERATIONS_BUDGETS | &quot;operations-budgets&quot;
OPERATIONS_INVOICES | &quot;operations-invoices&quot;
OPERATIONS_WIKI | &quot;operations-wiki&quot;
PROJECTS | &quot;projects&quot;
PROVISIONING | &quot;provisioning&quot;
PROVISIONING_ADD | &quot;provisioning-add&quot;
PROVISIONING_ADMIN | &quot;provisioning-admin&quot;
PROVISIONING_CLONE | &quot;provisioning-clone&quot;
PROVISIONING_DELETE | &quot;provisioning-delete&quot;
PROVISIONING_EDIT | &quot;provisioning-edit&quot;
PROVISIONING_ENVIRONMENT | &quot;provisioning-environment&quot;
PROVISIONING_EXECUTE_SCRIPT | &quot;provisioning-execute-script&quot;
PROVISIONING_EXECUTE_TASK | &quot;provisioning-execute-task&quot;
PROVISIONING_EXECUTE_WORKFLOW | &quot;provisioning-execute-workflow&quot;
PROVISIONING_FORCE_DELETE | &quot;provisioning-force-delete&quot;
PROVISIONING_IMPORT_IMAGE | &quot;provisioning-import-image&quot;
PROVISIONING_LOCK | &quot;provisioning-lock&quot;
PROVISIONING_POWER | &quot;provisioning-power&quot;
PROVISIONING_RECONFIGURE | &quot;provisioning-reconfigure&quot;
PROVISIONING_RECONFIGURE_ADD_DISK | &quot;provisioning-reconfigure-add-disk&quot;
PROVISIONING_RECONFIGURE_ADD_NETWORK | &quot;provisioning-reconfigure-add-network&quot;
PROVISIONING_RECONFIGURE_CHANGE_PLAN | &quot;provisioning-reconfigure-change-plan&quot;
PROVISIONING_RECONFIGURE_DISK_TYPE | &quot;provisioning-reconfigure-disk-type&quot;
PROVISIONING_RECONFIGURE_MODIFY_DISK | &quot;provisioning-reconfigure-modify-disk&quot;
PROVISIONING_RECONFIGURE_MODIFY_NETWORK | &quot;provisioning-reconfigure-modify-network&quot;
PROVISIONING_RECONFIGURE_REMOVE_DISK | &quot;provisioning-reconfigure-remove-disk&quot;
PROVISIONING_RECONFIGURE_REMOVE_NETWORK | &quot;provisioning-reconfigure-remove-network&quot;
PROVISIONING_REMOVE_CONTROL | &quot;provisioning-remove-control&quot;
PROVISIONING_SCALE | &quot;provisioning-scale&quot;
PROVISIONING_SETTINGS | &quot;provisioning-settings&quot;
PROVISIONING_STATE | &quot;provisioning-state&quot;
REPORTS | &quot;reports&quot;
REPORTS_ANALYTICS | &quot;reports-analytics&quot;
SCHEDULING_EXECUTE | &quot;scheduling-execute&quot;
SCHEDULING_POWER | &quot;scheduling-power&quot;
SECURITY_SCAN | &quot;security-scan&quot;
SERVICE_CATALOG | &quot;service-catalog&quot;
SERVICE_CATALOG_DASHBOARD | &quot;service-catalog-dashboard&quot;
SERVICE_CATALOG_INVENTORY | &quot;service-catalog-inventory&quot;
SERVICES_ARCHIVES | &quot;services-archives&quot;
SERVICES_CYPHER | &quot;services-cypher&quot;
SERVICES_IMAGE_BUILDER | &quot;services-image-builder&quot;
SERVICES_KUBERNETES | &quot;services-kubernetes&quot;
SERVICES_NETWORK_REGISTRY | &quot;services-network-registry&quot;
SERVICES_VDI_COPY | &quot;services-vdi-copy&quot;
SERVICES_VDI_POOLS | &quot;services-vdi-pools&quot;
SERVICES_VDI_PRINTER | &quot;services-vdi-printer&quot;
SNAPSHOTS | &quot;snapshots&quot;
TASK_SCRIPTS | &quot;task-scripts&quot;
TASKS | &quot;tasks&quot;
TERMINAL | &quot;terminal&quot;
TERMINAL_ACCESS | &quot;terminal-access&quot;
TERRAFORM_TEMPLATE | &quot;terraform-template&quot;
THRESHOLDS | &quot;thresholds&quot;
TRUST_SERVICES | &quot;trust-services&quot;
VIRTUAL_IMAGES | &quot;virtual-images&quot;



## Enum: AccessEnum

Name | Value
---- | -----
FULL | &quot;full&quot;
FULL_DECRYPTED | &quot;full_decrypted&quot;
GROUP | &quot;group&quot;
LISTFILES | &quot;listfiles&quot;
MANAGERULES | &quot;managerules&quot;
NO | &quot;no&quot;
NONE | &quot;none&quot;
PROVISION | &quot;provision&quot;
READ | &quot;read&quot;
ROLEMAPPINGS | &quot;rolemappings&quot;
USER | &quot;user&quot;
VIEW | &quot;view&quot;
YES | &quot;yes&quot;



