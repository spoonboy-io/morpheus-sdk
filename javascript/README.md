# morpheus_api

MorpheusApi - JavaScript client for morpheus_api
Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.

This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
This SDK is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: 6.1.1
- Package version: 6.1.1
- Build package: org.openapitools.codegen.languages.JavascriptClientCodegen

## Installation

### For [Node.js](https://nodejs.org/)

#### npm

To publish the library as a [npm](https://www.npmjs.com/), please follow the procedure in ["Publishing npm packages"](https://docs.npmjs.com/getting-started/publishing-npm-packages).

Then install it via:

```shell
npm install morpheus_api --save
```

Finally, you need to build the module:

```shell
npm run build
```

##### Local development

To use the library locally without publishing to a remote npm registry, first install the dependencies by changing into the directory containing `package.json` (and this README). Let's call this `JAVASCRIPT_CLIENT_DIR`. Then run:

```shell
npm install
```

Next, [link](https://docs.npmjs.com/cli/link) it globally in npm with the following, also from `JAVASCRIPT_CLIENT_DIR`:

```shell
npm link
```

To use the link you just defined in your project, switch to the directory you want to use your morpheus_api from, and run:

```shell
npm link /path/to/<JAVASCRIPT_CLIENT_DIR>
```

Finally, you need to build the module:

```shell
npm run build
```

#### git

If the library is hosted at a git repository, e.g.https://github.com/GIT_USER_ID/GIT_REPO_ID
then install it via:

```shell
    npm install GIT_USER_ID/GIT_REPO_ID --save
```

### For browser

The library also works in the browser environment via npm and [browserify](http://browserify.org/). After following
the above steps with Node.js and installing browserify with `npm install -g browserify`,
perform the following (assuming *main.js* is your entry file):

```shell
browserify main.js > bundle.js
```

Then include *bundle.js* in the HTML pages.

### Webpack Configuration

Using Webpack you may encounter the following error: "Module not found: Error:
Cannot resolve module", most certainly you should disable AMD loader. Add/merge
the following section to your webpack config:

```javascript
module: {
  rules: [
    {
      parser: {
        amd: false
      }
    }
  ]
}
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following JS code:

```javascript
var MorpheusApi = require('morpheus_api');

var defaultClient = MorpheusApi.ApiClient.instance;
// Configure Bearer access token for authorization: bearerAuth
var bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

var api = new MorpheusApi.ActivityApi()
var opts = {
  'max': 25, // {Number} Maximum number of records to return
  'offset': 0, // {Number} Offset records, the number of records to skip, for paginating requests
  'sort': "'name'", // {String} Sort order, the name of the property to sort by
  'order': "'asc'", // {String} Sort direction, use 'desc' to reverse sort
  'phrase': "phrase_example", // {String} Search phrase for partial matches on name or description
  'name': example-%, // {String} Filter by name, wildcard may be specified as %, eg. example-%
  'userId': 6, // {Number} Filter by User ID
  'tenantId': 3, // {Number} Filter by Tenant ID. Only available to the master account.
  'timeframe': "'month'", // {String} Filter by a timeframe (ex - today, yesterday, week, month, 3months)
  'start': new Date("2013-10-20T19:20:30+01:00"), // {Date} Filter by activity on or after a date(time). Default is 1 month prior
  'end': new Date("2013-10-20T19:20:30+01:00") // {Date} Filter by activity on or before a date(time). Default is current date
};
var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
api.listActivity(opts, callback);

```

## Documentation for API Endpoints

All URIs are relative to *https://CHANGEME*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MorpheusApi.ActivityApi* | [**listActivity**](docs/ActivityApi.md#listActivity) | **GET** /api/activity | Retrieves Activity
*MorpheusApi.AlertsApi* | [**addAlerts**](docs/AlertsApi.md#addAlerts) | **POST** /api/monitoring/alerts | Create a New Alert
*MorpheusApi.AlertsApi* | [**deleteAlerts**](docs/AlertsApi.md#deleteAlerts) | **DELETE** /api/monitoring/alerts/{id} | Delete a Specific Alert
*MorpheusApi.AlertsApi* | [**getAlerts**](docs/AlertsApi.md#getAlerts) | **GET** /api/monitoring/alerts/{id} | Get a Specific Alert
*MorpheusApi.AlertsApi* | [**listAlerts**](docs/AlertsApi.md#listAlerts) | **GET** /api/monitoring/alerts | List All Alerts
*MorpheusApi.AlertsApi* | [**updateAlerts**](docs/AlertsApi.md#updateAlerts) | **PUT** /api/monitoring/alerts/{id} | Update Alert
*MorpheusApi.ApplianceSettingsApi* | [**listApplianceSettings**](docs/ApplianceSettingsApi.md#listApplianceSettings) | **GET** /api/appliance-settings | Get Appliance Settings
*MorpheusApi.ApplianceSettingsApi* | [**reindex**](docs/ApplianceSettingsApi.md#reindex) | **POST** /api/appliance-settings/reindex | Reindex Search
*MorpheusApi.ApplianceSettingsApi* | [**setApplianceSettingsMaintenanceMode**](docs/ApplianceSettingsApi.md#setApplianceSettingsMaintenanceMode) | **POST** /api/appliance-settings/maintenance | Toggle Maintenance Mode
*MorpheusApi.ApplianceSettingsApi* | [**updateApplianceSettings**](docs/ApplianceSettingsApi.md#updateApplianceSettings) | **PUT** /api/appliance-settings | Update Appliance Settings
*MorpheusApi.ApprovalsApi* | [**getApprovals**](docs/ApprovalsApi.md#getApprovals) | **GET** /api/approvals/{id} | Retrieves a Specific Approval
*MorpheusApi.ApprovalsApi* | [**getApprovalsItem**](docs/ApprovalsApi.md#getApprovalsItem) | **GET** /api/approval-items/{id} | Retrieves a Specific Approval Item
*MorpheusApi.ApprovalsApi* | [**listApprovals**](docs/ApprovalsApi.md#listApprovals) | **GET** /api/approvals | Retrieves all Approvals
*MorpheusApi.ApprovalsApi* | [**updateApprovalsItem**](docs/ApprovalsApi.md#updateApprovalsItem) | **PUT** /api/approval-items/{id}/{action} | Updates a Specific Approval Item
*MorpheusApi.AppsApi* | [**addAppInstance**](docs/AppsApi.md#addAppInstance) | **POST** /api/apps/{id}/add-instance | Add Existing Instance to App
*MorpheusApi.AppsApi* | [**addAppUndoDelete**](docs/AppsApi.md#addAppUndoDelete) | **PUT** /api/apps/{id}/cancel-removal | Undo Delete of an App
*MorpheusApi.AppsApi* | [**addApps**](docs/AppsApi.md#addApps) | **POST** /api/apps | Create an App
*MorpheusApi.AppsApi* | [**applyAppState**](docs/AppsApi.md#applyAppState) | **POST** /api/apps/{id}/apply | Apply State of an App
*MorpheusApi.AppsApi* | [**deleteApp**](docs/AppsApi.md#deleteApp) | **DELETE** /api/apps/{id} | Delete an App
*MorpheusApi.AppsApi* | [**getApp**](docs/AppsApi.md#getApp) | **GET** /api/apps/{id} | Get a Specific App
*MorpheusApi.AppsApi* | [**getAppSecurityGroups**](docs/AppsApi.md#getAppSecurityGroups) | **GET** /api/apps/{id}/security-groups | Get Security Groups for an App
*MorpheusApi.AppsApi* | [**getAppState**](docs/AppsApi.md#getAppState) | **GET** /api/apps/{id}/state | Get State of an App
*MorpheusApi.AppsApi* | [**listApps**](docs/AppsApi.md#listApps) | **GET** /api/apps | Get All Apps
*MorpheusApi.AppsApi* | [**prepareAppApply**](docs/AppsApi.md#prepareAppApply) | **GET** /api/apps/{id}/prepare-apply | Prepare To Apply an App
*MorpheusApi.AppsApi* | [**refreshAppState**](docs/AppsApi.md#refreshAppState) | **POST** /api/apps/{id}/refresh | Refresh State of an App
*MorpheusApi.AppsApi* | [**removeAppInstance**](docs/AppsApi.md#removeAppInstance) | **POST** /api/apps/{id}/remove-instance | Remove Instance from App
*MorpheusApi.AppsApi* | [**setAppSecurityGroups**](docs/AppsApi.md#setAppSecurityGroups) | **POST** /api/apps/{id}/security-groups | Set Security Groups for an App
*MorpheusApi.AppsApi* | [**updateApp**](docs/AppsApi.md#updateApp) | **PUT** /api/apps/{id} | Updating an App
*MorpheusApi.AppsApi* | [**validateAppState**](docs/AppsApi.md#validateAppState) | **POST** /api/apps/{id}/validate-apply | Validate Apply State for an App
*MorpheusApi.ArchivesApi* | [**addArchiveBucket**](docs/ArchivesApi.md#addArchiveBucket) | **POST** /api/archives/buckets | Create an Archive Bucket
*MorpheusApi.ArchivesApi* | [**addArchiveFile**](docs/ArchivesApi.md#addArchiveFile) | **POST** /api/archives/buckets/{bucket}/files/{filepath} | Upload Archive File
*MorpheusApi.ArchivesApi* | [**addArchiveFileLink**](docs/ArchivesApi.md#addArchiveFileLink) | **POST** /api/archives/files/{id}/links | Create an Archive File Link
*MorpheusApi.ArchivesApi* | [**deleteArchiveBucket**](docs/ArchivesApi.md#deleteArchiveBucket) | **DELETE** /api/archives/buckets/{id} | Delete an Archive Bucket
*MorpheusApi.ArchivesApi* | [**deleteArchiveFile**](docs/ArchivesApi.md#deleteArchiveFile) | **DELETE** /api/archives/files/{id} | Delete Archive File
*MorpheusApi.ArchivesApi* | [**deleteArchiveFileLink**](docs/ArchivesApi.md#deleteArchiveFileLink) | **DELETE** /api/archives/files/{id}/links/{linkId} | Delete an Archive File Link
*MorpheusApi.ArchivesApi* | [**getArchiveBucket**](docs/ArchivesApi.md#getArchiveBucket) | **GET** /api/archives/buckets/{id} | Get a Specific Archive Bucket
*MorpheusApi.ArchivesApi* | [**getArchiveFile**](docs/ArchivesApi.md#getArchiveFile) | **GET** /api/archives/download/{bucket}/{filepath} | Download an Archive File
*MorpheusApi.ArchivesApi* | [**getArchiveFileDetail**](docs/ArchivesApi.md#getArchiveFileDetail) | **GET** /api/archives/files/{id} | Get Archive File Details
*MorpheusApi.ArchivesApi* | [**getArchiveFileLinks**](docs/ArchivesApi.md#getArchiveFileLinks) | **GET** /api/archives/files/{id}/links | Get Archive File Links
*MorpheusApi.ArchivesApi* | [**listArchiveBuckets**](docs/ArchivesApi.md#listArchiveBuckets) | **GET** /api/archives/buckets | Get All Archive Buckets
*MorpheusApi.ArchivesApi* | [**listArchiveFiles**](docs/ArchivesApi.md#listArchiveFiles) | **GET** /api/archives/buckets/{bucket}/files/{filepath} | Get All Archive Files
*MorpheusApi.ArchivesApi* | [**updateArchiveBucket**](docs/ArchivesApi.md#updateArchiveBucket) | **PUT** /api/archives/buckets/{id} | Update an Archive Bucket
*MorpheusApi.AuthenticationApi* | [**forgotPassword**](docs/AuthenticationApi.md#forgotPassword) | **POST** /api/forgot/send-email | Request a reset password email
*MorpheusApi.AuthenticationApi* | [**resetPassword**](docs/AuthenticationApi.md#resetPassword) | **POST** /api/forgot/reset-password | Reset user password
*MorpheusApi.AutomationApi* | [**addExecuteSchedules**](docs/AutomationApi.md#addExecuteSchedules) | **POST** /api/execute-schedules | Creates a Execute Schedule
*MorpheusApi.AutomationApi* | [**executeExecutionRequest**](docs/AutomationApi.md#executeExecutionRequest) | **POST** /api/execution-request/execute | Executes an Execution Request
*MorpheusApi.AutomationApi* | [**getExecuteSchedules**](docs/AutomationApi.md#getExecuteSchedules) | **GET** /api/execute-schedules/{id} | Retrieves a Specific Execute Schedule
*MorpheusApi.AutomationApi* | [**getExecutionRequest**](docs/AutomationApi.md#getExecutionRequest) | **GET** /api/execution-request/{uniqueId} | Retrieves a Specific Execution Request
*MorpheusApi.AutomationApi* | [**listExecuteSchedules**](docs/AutomationApi.md#listExecuteSchedules) | **GET** /api/execute-schedules | Retrieves all Execute Schedules
*MorpheusApi.AutomationApi* | [**removeExecuteSchedules**](docs/AutomationApi.md#removeExecuteSchedules) | **DELETE** /api/execute-schedules/{id} | Deletes a Execute Schedule
*MorpheusApi.AutomationApi* | [**updateExecuteSchedules**](docs/AutomationApi.md#updateExecuteSchedules) | **PUT** /api/execute-schedules/{id} | Updates a Execute Schedule
*MorpheusApi.BackupSettingsApi* | [**listBackupSettings**](docs/BackupSettingsApi.md#listBackupSettings) | **GET** /api/backup-settings | Get Backup Settings
*MorpheusApi.BackupSettingsApi* | [**updateBackupSettings**](docs/BackupSettingsApi.md#updateBackupSettings) | **PUT** /api/backup-settings | Update Backup Settings
*MorpheusApi.BillingApi* | [**getBillingAccount**](docs/BillingApi.md#getBillingAccount) | **GET** /api/billing/account/{id} | This endpoint will retrieve a specific account by id if the user has permission to access it
*MorpheusApi.BillingApi* | [**getBillingInstancesIdentifier**](docs/BillingApi.md#getBillingInstancesIdentifier) | **GET** /api/billing/instances/{identifier} | Retrieves billing information for an instance in the requestor&#39;s account. Use instanceUUID whenever possible.
*MorpheusApi.BillingApi* | [**getBillingServersIdentifier**](docs/BillingApi.md#getBillingServersIdentifier) | **GET** /api/billing/servers/{identifier} | Retrieves billing information for a specific server (container host) in the requestor&#39;s account. Use refUUID whenever possible.
*MorpheusApi.BillingApi* | [**getBillingZoneIdentifier**](docs/BillingApi.md#getBillingZoneIdentifier) | **GET** /api/billing/zones/{identifier} | Retrieves billing information for a specific zone in the requestor&#39;s account. Use zoneUUID whenever possible.
*MorpheusApi.BillingApi* | [**listBillingAccount**](docs/BillingApi.md#listBillingAccount) | **GET** /api/billing/account | Retrieves billing information for the requesting user&#39;s account.
*MorpheusApi.BillingApi* | [**listBillingInstances**](docs/BillingApi.md#listBillingInstances) | **GET** /api/billing/instances | Retrieves billing information for all instances on the requestor&#39;s account.
*MorpheusApi.BillingApi* | [**listBillingServers**](docs/BillingApi.md#listBillingServers) | **GET** /api/billing/servers | Retrieves billing information for all servers (container hosts) on the requestor&#39;s account.
*MorpheusApi.BillingApi* | [**listBillingZone**](docs/BillingApi.md#listBillingZone) | **GET** /api/billing/zones | Retrieves billing information for all zones on the requestor&#39;s account.
*MorpheusApi.BlueprintsApi* | [**addBlueprint**](docs/BlueprintsApi.md#addBlueprint) | **POST** /api/blueprints | Create a Blueprint
*MorpheusApi.BlueprintsApi* | [**deleteBlueprint**](docs/BlueprintsApi.md#deleteBlueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint
*MorpheusApi.BlueprintsApi* | [**getBlueprint**](docs/BlueprintsApi.md#getBlueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint
*MorpheusApi.BlueprintsApi* | [**listBlueprints**](docs/BlueprintsApi.md#listBlueprints) | **GET** /api/blueprints | Get All Blueprints
*MorpheusApi.BlueprintsApi* | [**updateBlueprint**](docs/BlueprintsApi.md#updateBlueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint
*MorpheusApi.BlueprintsApi* | [**updateBlueprintImage**](docs/BlueprintsApi.md#updateBlueprintImage) | **POST** /api/blueprints/{id}/image | Update Blueprint Image
*MorpheusApi.BlueprintsApi* | [**updateBlueprintPermissions**](docs/BlueprintsApi.md#updateBlueprintPermissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions
*MorpheusApi.BudgetsApi* | [**addBudgets**](docs/BudgetsApi.md#addBudgets) | **POST** /api/budgets | Creates a Budget
*MorpheusApi.BudgetsApi* | [**getBudgets**](docs/BudgetsApi.md#getBudgets) | **GET** /api/budgets/{id} | Retrieves a Specific Budget
*MorpheusApi.BudgetsApi* | [**listBudgets**](docs/BudgetsApi.md#listBudgets) | **GET** /api/budgets | Retrieves all Budgets
*MorpheusApi.BudgetsApi* | [**removeBudgets**](docs/BudgetsApi.md#removeBudgets) | **DELETE** /api/budgets/{id} | Deletes a Budget
*MorpheusApi.BudgetsApi* | [**updateBudgets**](docs/BudgetsApi.md#updateBudgets) | **PUT** /api/budgets/{id} | Updates a Budget
*MorpheusApi.CatalogItemsApi* | [**addCatalogItemType**](docs/CatalogItemsApi.md#addCatalogItemType) | **POST** /api/catalog-item-types | Create a Catalog Item Type
*MorpheusApi.CatalogItemsApi* | [**getCatalogItemType**](docs/CatalogItemsApi.md#getCatalogItemType) | **GET** /api/catalog-item-types/{id} | Get a Specific Catalog Item Type
*MorpheusApi.CatalogItemsApi* | [**listCatalogItemTypes**](docs/CatalogItemsApi.md#listCatalogItemTypes) | **GET** /api/catalog-item-types | Get All Catalog Item Types
*MorpheusApi.CatalogItemsApi* | [**removeCatalogItemType**](docs/CatalogItemsApi.md#removeCatalogItemType) | **DELETE** /api/catalog-item-types/{id} | Delete a Catalog Item Type
*MorpheusApi.CatalogItemsApi* | [**updateCatalogItemType**](docs/CatalogItemsApi.md#updateCatalogItemType) | **PUT** /api/catalog-item-types/{id} | Update a Catalog Item Type
*MorpheusApi.CatalogItemsApi* | [**updateCatalogItemTypeLogo**](docs/CatalogItemsApi.md#updateCatalogItemTypeLogo) | **PUT** /api/catalog-item-types/{id}/update-logo | Update Logo For Catalog Item Type
*MorpheusApi.ChecksApi* | [**addCheckApps**](docs/ChecksApi.md#addCheckApps) | **POST** /api/monitoring/apps | Create a New Check App
*MorpheusApi.ChecksApi* | [**listCheckApps**](docs/ChecksApi.md#listCheckApps) | **GET** /api/monitoring/apps | List All Check Apps
*MorpheusApi.ClientsApi* | [**addClient**](docs/ClientsApi.md#addClient) | **POST** /api/clients | Create an Oauth Client
*MorpheusApi.ClientsApi* | [**getClients**](docs/ClientsApi.md#getClients) | **GET** /api/clients/{id} | Retrieves a Specific Oauth Client
*MorpheusApi.ClientsApi* | [**listClients**](docs/ClientsApi.md#listClients) | **GET** /api/clients | Get All Oauth Clients
*MorpheusApi.ClientsApi* | [**removeClients**](docs/ClientsApi.md#removeClients) | **DELETE** /api/clients/{id} | Deletes an Oauth Client
*MorpheusApi.ClientsApi* | [**updateClients**](docs/ClientsApi.md#updateClients) | **PUT** /api/clients/{id} | Updates an Oauth Client


## Documentation for Models

 - [MorpheusApi.Activity](docs/Activity.md)
 - [MorpheusApi.ActivityActivityInner](docs/ActivityActivityInner.md)
 - [MorpheusApi.ActivityActivityInnerUser](docs/ActivityActivityInnerUser.md)
 - [MorpheusApi.AddAlerts200Response](docs/AddAlerts200Response.md)
 - [MorpheusApi.AddAlertsRequest](docs/AddAlertsRequest.md)
 - [MorpheusApi.AddAlertsRequestAlert](docs/AddAlertsRequestAlert.md)
 - [MorpheusApi.AddAlertsRequestAlertContactsInner](docs/AddAlertsRequestAlertContactsInner.md)
 - [MorpheusApi.AddAppInstanceRequest](docs/AddAppInstanceRequest.md)
 - [MorpheusApi.AddApps200Response](docs/AddApps200Response.md)
 - [MorpheusApi.AddArchiveBucket200Response](docs/AddArchiveBucket200Response.md)
 - [MorpheusApi.AddArchiveBucketRequest](docs/AddArchiveBucketRequest.md)
 - [MorpheusApi.AddArchiveFile200Response](docs/AddArchiveFile200Response.md)
 - [MorpheusApi.AddArchiveFileLink200Response](docs/AddArchiveFileLink200Response.md)
 - [MorpheusApi.AddBlueprint200Response](docs/AddBlueprint200Response.md)
 - [MorpheusApi.AddBlueprintRequest](docs/AddBlueprintRequest.md)
 - [MorpheusApi.AddBudgets200Response](docs/AddBudgets200Response.md)
 - [MorpheusApi.AddBudgetsRequest](docs/AddBudgetsRequest.md)
 - [MorpheusApi.AddBudgetsRequestBudget](docs/AddBudgetsRequestBudget.md)
 - [MorpheusApi.AddCatalogItemType200Response](docs/AddCatalogItemType200Response.md)
 - [MorpheusApi.AddCatalogItemTypeRequest](docs/AddCatalogItemTypeRequest.md)
 - [MorpheusApi.AddCatalogItemTypeRequestCatalogItemType](docs/AddCatalogItemTypeRequestCatalogItemType.md)
 - [MorpheusApi.AddCheckApps200Response](docs/AddCheckApps200Response.md)
 - [MorpheusApi.AddCheckAppsRequest](docs/AddCheckAppsRequest.md)
 - [MorpheusApi.AddCheckAppsRequestMonitorApp](docs/AddCheckAppsRequestMonitorApp.md)
 - [MorpheusApi.AddClient200Response](docs/AddClient200Response.md)
 - [MorpheusApi.AddClientRequest](docs/AddClientRequest.md)
 - [MorpheusApi.AddClientRequestClient](docs/AddClientRequestClient.md)
 - [MorpheusApi.AddExecuteSchedules200Response](docs/AddExecuteSchedules200Response.md)
 - [MorpheusApi.AddExecuteSchedulesRequest](docs/AddExecuteSchedulesRequest.md)
 - [MorpheusApi.AddExecuteSchedulesRequestSchedule](docs/AddExecuteSchedulesRequestSchedule.md)
 - [MorpheusApi.Alert](docs/Alert.md)
 - [MorpheusApi.App](docs/App.md)
 - [MorpheusApi.AppBlueprint](docs/AppBlueprint.md)
 - [MorpheusApi.AppCreate](docs/AppCreate.md)
 - [MorpheusApi.AppCreateBlueprintId](docs/AppCreateBlueprintId.md)
 - [MorpheusApi.AppCreateDefaultCloud](docs/AppCreateDefaultCloud.md)
 - [MorpheusApi.AppCreateGroup](docs/AppCreateGroup.md)
 - [MorpheusApi.AppCreateResponse](docs/AppCreateResponse.md)
 - [MorpheusApi.AppPrepareApply](docs/AppPrepareApply.md)
 - [MorpheusApi.AppPrepareApplyData](docs/AppPrepareApplyData.md)
 - [MorpheusApi.AppPrepareApplyDataGroup](docs/AppPrepareApplyDataGroup.md)
 - [MorpheusApi.AppPrepareApplyDataTerraform](docs/AppPrepareApplyDataTerraform.md)
 - [MorpheusApi.AppSecurityGroups](docs/AppSecurityGroups.md)
 - [MorpheusApi.AppState](docs/AppState.md)
 - [MorpheusApi.AppStateInput](docs/AppStateInput.md)
 - [MorpheusApi.AppStateInputDataInner](docs/AppStateInputDataInner.md)
 - [MorpheusApi.AppStateInputDataInnerName](docs/AppStateInputDataInnerName.md)
 - [MorpheusApi.AppStateInputProvidersInner](docs/AppStateInputProvidersInner.md)
 - [MorpheusApi.AppStateInputVariablesInner](docs/AppStateInputVariablesInner.md)
 - [MorpheusApi.AppStateOutput](docs/AppStateOutput.md)
 - [MorpheusApi.AppStateSpecsInner](docs/AppStateSpecsInner.md)
 - [MorpheusApi.AppStateSpecsInnerTemplate](docs/AppStateSpecsInnerTemplate.md)
 - [MorpheusApi.AppStateWorkloadsInner](docs/AppStateWorkloadsInner.md)
 - [MorpheusApi.AppStats](docs/AppStats.md)
 - [MorpheusApi.AppUpdate](docs/AppUpdate.md)
 - [MorpheusApi.ApplianceSettings](docs/ApplianceSettings.md)
 - [MorpheusApi.ApplianceSettingsEnabledZoneTypesInner](docs/ApplianceSettingsEnabledZoneTypesInner.md)
 - [MorpheusApi.ApplianceSettingsUpdate](docs/ApplianceSettingsUpdate.md)
 - [MorpheusApi.ApplyAppStateRequest](docs/ApplyAppStateRequest.md)
 - [MorpheusApi.Approval](docs/Approval.md)
 - [MorpheusApi.ApprovalItem](docs/ApprovalItem.md)
 - [MorpheusApi.ApprovalItemApprovalItem](docs/ApprovalItemApprovalItem.md)
 - [MorpheusApi.ApprovalItemApprovalItemReference](docs/ApprovalItemApprovalItemReference.md)
 - [MorpheusApi.Approvals](docs/Approvals.md)
 - [MorpheusApi.ArchiveBucket](docs/ArchiveBucket.md)
 - [MorpheusApi.ArchiveBucketCreate](docs/ArchiveBucketCreate.md)
 - [MorpheusApi.ArchiveBucketCreateStorageProvider](docs/ArchiveBucketCreateStorageProvider.md)
 - [MorpheusApi.ArchiveBucketCreatedBy](docs/ArchiveBucketCreatedBy.md)
 - [MorpheusApi.ArchiveBucketFile](docs/ArchiveBucketFile.md)
 - [MorpheusApi.ArchiveBucketFileArchiveBucket](docs/ArchiveBucketFileArchiveBucket.md)
 - [MorpheusApi.ArchiveBucketFileCreatedBy](docs/ArchiveBucketFileCreatedBy.md)
 - [MorpheusApi.ArchiveBucketUpdate](docs/ArchiveBucketUpdate.md)
 - [MorpheusApi.ArchiveFileLinks](docs/ArchiveFileLinks.md)
 - [MorpheusApi.ArchiveFileLinksArchiveFile](docs/ArchiveFileLinksArchiveFile.md)
 - [MorpheusApi.BackupSettings](docs/BackupSettings.md)
 - [MorpheusApi.BackupSettingsDefaultSchedule](docs/BackupSettingsDefaultSchedule.md)
 - [MorpheusApi.BackupSettingsUpdate](docs/BackupSettingsUpdate.md)
 - [MorpheusApi.BackupSettingsUpdateDefaultSchedule](docs/BackupSettingsUpdateDefaultSchedule.md)
 - [MorpheusApi.BackupSettingsUpdateDefaultStorageBucket](docs/BackupSettingsUpdateDefaultStorageBucket.md)
 - [MorpheusApi.Billing](docs/Billing.md)
 - [MorpheusApi.BillingInstance](docs/BillingInstance.md)
 - [MorpheusApi.BillingInstanceContainersInner](docs/BillingInstanceContainersInner.md)
 - [MorpheusApi.BillingInstanceContainersInnerUsagesInner](docs/BillingInstanceContainersInnerUsagesInner.md)
 - [MorpheusApi.BillingInstanceContainersInnerUsagesInnerApplicablePricesInner](docs/BillingInstanceContainersInnerUsagesInnerApplicablePricesInner.md)
 - [MorpheusApi.BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner](docs/BillingInstanceContainersInnerUsagesInnerApplicablePricesInnerPricesInner.md)
 - [MorpheusApi.BillingInstances](docs/BillingInstances.md)
 - [MorpheusApi.BillingInstancesInstancesInner](docs/BillingInstancesInstancesInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInner](docs/BillingInstancesInstancesInnerContainersInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInnerUsagesInner](docs/BillingInstancesInstancesInnerContainersInnerUsagesInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner](docs/BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner](docs/BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner](docs/BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner](docs/BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.md)
 - [MorpheusApi.BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore](docs/BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.md)
 - [MorpheusApi.BillingServer](docs/BillingServer.md)
 - [MorpheusApi.BillingServerUsagesInner](docs/BillingServerUsagesInner.md)
 - [MorpheusApi.BillingServerUsagesInnerApplicablePricesInner](docs/BillingServerUsagesInnerApplicablePricesInner.md)
 - [MorpheusApi.BillingServerUsagesInnerApplicablePricesInnerPricesInner](docs/BillingServerUsagesInnerApplicablePricesInnerPricesInner.md)
 - [MorpheusApi.BillingServers](docs/BillingServers.md)
 - [MorpheusApi.BillingServersServersInner](docs/BillingServersServersInner.md)
 - [MorpheusApi.BillingServersServersInnerUsagesInner](docs/BillingServersServersInnerUsagesInner.md)
 - [MorpheusApi.BillingServersServersInnerUsagesInnerApplicablePricesInner](docs/BillingServersServersInnerUsagesInnerApplicablePricesInner.md)
 - [MorpheusApi.BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner](docs/BillingServersServersInnerUsagesInnerApplicablePricesInnerPricesInner.md)
 - [MorpheusApi.BillingServersServersInnerUsagesInnerPricesUsedInner](docs/BillingServersServersInnerUsagesInnerPricesUsedInner.md)
 - [MorpheusApi.BillingServersServersInnerUsagesInnerVolumesInner](docs/BillingServersServersInnerUsagesInnerVolumesInner.md)
 - [MorpheusApi.BillingZone](docs/BillingZone.md)
 - [MorpheusApi.BillingZones](docs/BillingZones.md)
 - [MorpheusApi.BillingZonesInner](docs/BillingZonesInner.md)
 - [MorpheusApi.BillingZonesInnerComputeServers](docs/BillingZonesInnerComputeServers.md)
 - [MorpheusApi.BillingZonesInnerInstances](docs/BillingZonesInnerInstances.md)
 - [MorpheusApi.BillingZonesInnerLoadBalancers](docs/BillingZonesInnerLoadBalancers.md)
 - [MorpheusApi.BillingZonesInnerSnapshots](docs/BillingZonesInnerSnapshots.md)
 - [MorpheusApi.BillingZonesInnerVirtualImages](docs/BillingZonesInnerVirtualImages.md)
 - [MorpheusApi.Blueprint](docs/Blueprint.md)
 - [MorpheusApi.BlueprintARMCreate](docs/BlueprintARMCreate.md)
 - [MorpheusApi.BlueprintARMCreateArm](docs/BlueprintARMCreateArm.md)
 - [MorpheusApi.BlueprintARMCreateArmCloudInitEnabled](docs/BlueprintARMCreateArmCloudInitEnabled.md)
 - [MorpheusApi.BlueprintARMCreateArmGit](docs/BlueprintARMCreateArmGit.md)
 - [MorpheusApi.BlueprintARMCreateArmInstallAgent](docs/BlueprintARMCreateArmInstallAgent.md)
 - [MorpheusApi.BlueprintCFTCreate](docs/BlueprintCFTCreate.md)
 - [MorpheusApi.BlueprintCFTCreateCloudFormation](docs/BlueprintCFTCreateCloudFormation.md)
 - [MorpheusApi.BlueprintCFTCreateCloudFormationGit](docs/BlueprintCFTCreateCloudFormationGit.md)
 - [MorpheusApi.BlueprintCreateSuccess](docs/BlueprintCreateSuccess.md)
 - [MorpheusApi.BlueprintCreateSuccessConfig](docs/BlueprintCreateSuccessConfig.md)
 - [MorpheusApi.BlueprintHelmCreate](docs/BlueprintHelmCreate.md)
 - [MorpheusApi.BlueprintHelmCreateHelm](docs/BlueprintHelmCreateHelm.md)
 - [MorpheusApi.BlueprintHelmCreateHelmGit](docs/BlueprintHelmCreateHelmGit.md)
 - [MorpheusApi.BlueprintKubernetesCreate](docs/BlueprintKubernetesCreate.md)
 - [MorpheusApi.BlueprintKubernetesCreateConfig](docs/BlueprintKubernetesCreateConfig.md)
 - [MorpheusApi.BlueprintKubernetesCreateConfigSpecsInner](docs/BlueprintKubernetesCreateConfigSpecsInner.md)
 - [MorpheusApi.BlueprintKubernetesCreateKubernetes](docs/BlueprintKubernetesCreateKubernetes.md)
 - [MorpheusApi.BlueprintKubernetesCreateKubernetesGit](docs/BlueprintKubernetesCreateKubernetesGit.md)
 - [MorpheusApi.BlueprintMorpheusCreate](docs/BlueprintMorpheusCreate.md)
 - [MorpheusApi.BlueprintTerraformCreate](docs/BlueprintTerraformCreate.md)
 - [MorpheusApi.BlueprintTerraformCreateConfig](docs/BlueprintTerraformCreateConfig.md)
 - [MorpheusApi.BlueprintTerraformCreateTerraform](docs/BlueprintTerraformCreateTerraform.md)
 - [MorpheusApi.BlueprintTerraformCreateTerraformGit](docs/BlueprintTerraformCreateTerraformGit.md)
 - [MorpheusApi.Budget](docs/Budget.md)
 - [MorpheusApi.BudgetStats](docs/BudgetStats.md)
 - [MorpheusApi.BudgetStatsCurrent](docs/BudgetStatsCurrent.md)
 - [MorpheusApi.BudgetStatsIntervalsInner](docs/BudgetStatsIntervalsInner.md)
 - [MorpheusApi.Budgets](docs/Budgets.md)
 - [MorpheusApi.CatalogItemType](docs/CatalogItemType.md)
 - [MorpheusApi.CatalogItemTypeBlueprintCreate](docs/CatalogItemTypeBlueprintCreate.md)
 - [MorpheusApi.CatalogItemTypeBlueprintCreateBlueprint](docs/CatalogItemTypeBlueprintCreateBlueprint.md)
 - [MorpheusApi.CatalogItemTypeBlueprintUpdate](docs/CatalogItemTypeBlueprintUpdate.md)
 - [MorpheusApi.CatalogItemTypeInstanceCreate](docs/CatalogItemTypeInstanceCreate.md)
 - [MorpheusApi.CatalogItemTypeInstanceUpdate](docs/CatalogItemTypeInstanceUpdate.md)
 - [MorpheusApi.CatalogItemTypeWorkflowCreate](docs/CatalogItemTypeWorkflowCreate.md)
 - [MorpheusApi.CatalogItemTypeWorkflowUpdate](docs/CatalogItemTypeWorkflowUpdate.md)
 - [MorpheusApi.Check](docs/Check.md)
 - [MorpheusApi.CheckApp](docs/CheckApp.md)
 - [MorpheusApi.CheckCheckType](docs/CheckCheckType.md)
 - [MorpheusApi.CheckConfig](docs/CheckConfig.md)
 - [MorpheusApi.CheckContainer](docs/CheckContainer.md)
 - [MorpheusApi.CheckCreatedBy](docs/CheckCreatedBy.md)
 - [MorpheusApi.CheckGroup](docs/CheckGroup.md)
 - [MorpheusApi.CheckGroupInstance](docs/CheckGroupInstance.md)
 - [MorpheusApi.Client](docs/Client.md)
 - [MorpheusApi.ClientUpdate](docs/ClientUpdate.md)
 - [MorpheusApi.DefaultError](docs/DefaultError.md)
 - [MorpheusApi.ExecuteExecutionRequest200Response](docs/ExecuteExecutionRequest200Response.md)
 - [MorpheusApi.ExecuteExecutionRequestRequest](docs/ExecuteExecutionRequestRequest.md)
 - [MorpheusApi.ExecuteSchedule](docs/ExecuteSchedule.md)
 - [MorpheusApi.ExecutionRequest](docs/ExecutionRequest.md)
 - [MorpheusApi.ForgotPassword200Response](docs/ForgotPassword200Response.md)
 - [MorpheusApi.ForgotPasswordRequest](docs/ForgotPasswordRequest.md)
 - [MorpheusApi.GetAlerts200Response](docs/GetAlerts200Response.md)
 - [MorpheusApi.GetApp200Response](docs/GetApp200Response.md)
 - [MorpheusApi.GetAppSecurityGroups200Response](docs/GetAppSecurityGroups200Response.md)
 - [MorpheusApi.GetAppState200Response](docs/GetAppState200Response.md)
 - [MorpheusApi.GetApprovals200Response](docs/GetApprovals200Response.md)
 - [MorpheusApi.GetArchiveBucket200Response](docs/GetArchiveBucket200Response.md)
 - [MorpheusApi.GetArchiveFileDetail200Response](docs/GetArchiveFileDetail200Response.md)
 - [MorpheusApi.GetArchiveFileLinks200Response](docs/GetArchiveFileLinks200Response.md)
 - [MorpheusApi.GetBillingInstancesIdentifier200Response](docs/GetBillingInstancesIdentifier200Response.md)
 - [MorpheusApi.GetBillingServersIdentifier200Response](docs/GetBillingServersIdentifier200Response.md)
 - [MorpheusApi.GetBillingZoneIdentifier200Response](docs/GetBillingZoneIdentifier200Response.md)
 - [MorpheusApi.GetBlueprint200Response](docs/GetBlueprint200Response.md)
 - [MorpheusApi.GetBudgets200Response](docs/GetBudgets200Response.md)
 - [MorpheusApi.GetCatalogItemType200Response](docs/GetCatalogItemType200Response.md)
 - [MorpheusApi.GetClients200Response](docs/GetClients200Response.md)
 - [MorpheusApi.GetExecuteSchedules200Response](docs/GetExecuteSchedules200Response.md)
 - [MorpheusApi.GetExecutionRequest200Response](docs/GetExecutionRequest200Response.md)
 - [MorpheusApi.ListActivity200Response](docs/ListActivity200Response.md)
 - [MorpheusApi.ListAlerts200Response](docs/ListAlerts200Response.md)
 - [MorpheusApi.ListApplianceSettings200Response](docs/ListApplianceSettings200Response.md)
 - [MorpheusApi.ListApprovals200Response](docs/ListApprovals200Response.md)
 - [MorpheusApi.ListApps200Response](docs/ListApps200Response.md)
 - [MorpheusApi.ListArchiveBuckets200Response](docs/ListArchiveBuckets200Response.md)
 - [MorpheusApi.ListArchiveFiles200Response](docs/ListArchiveFiles200Response.md)
 - [MorpheusApi.ListBackupSettings200Response](docs/ListBackupSettings200Response.md)
 - [MorpheusApi.ListBillingAccount200Response](docs/ListBillingAccount200Response.md)
 - [MorpheusApi.ListBillingInstances200Response](docs/ListBillingInstances200Response.md)
 - [MorpheusApi.ListBillingServers200Response](docs/ListBillingServers200Response.md)
 - [MorpheusApi.ListBillingZone200Response](docs/ListBillingZone200Response.md)
 - [MorpheusApi.ListBlueprints200Response](docs/ListBlueprints200Response.md)
 - [MorpheusApi.ListBudgets200Response](docs/ListBudgets200Response.md)
 - [MorpheusApi.ListCatalogItemTypes200Response](docs/ListCatalogItemTypes200Response.md)
 - [MorpheusApi.ListCheckApps200Response](docs/ListCheckApps200Response.md)
 - [MorpheusApi.ListClients200Response](docs/ListClients200Response.md)
 - [MorpheusApi.ListExecuteSchedules200Response](docs/ListExecuteSchedules200Response.md)
 - [MorpheusApi.Meta](docs/Meta.md)
 - [MorpheusApi.MetaMeta](docs/MetaMeta.md)
 - [MorpheusApi.Model200Success](docs/Model200Success.md)
 - [MorpheusApi.PrepareAppApply200Response](docs/PrepareAppApply200Response.md)
 - [MorpheusApi.RemoveAppInstanceRequest](docs/RemoveAppInstanceRequest.md)
 - [MorpheusApi.ResetPassword200Response](docs/ResetPassword200Response.md)
 - [MorpheusApi.ResetPasswordRequest](docs/ResetPasswordRequest.md)
 - [MorpheusApi.SetAppSecurityGroups200Response](docs/SetAppSecurityGroups200Response.md)
 - [MorpheusApi.SetAppSecurityGroupsRequest](docs/SetAppSecurityGroupsRequest.md)
 - [MorpheusApi.UpdateAlerts200Response](docs/UpdateAlerts200Response.md)
 - [MorpheusApi.UpdateAlertsRequest](docs/UpdateAlertsRequest.md)
 - [MorpheusApi.UpdateAlertsRequestAlert](docs/UpdateAlertsRequestAlert.md)
 - [MorpheusApi.UpdateApplianceSettingsRequest](docs/UpdateApplianceSettingsRequest.md)
 - [MorpheusApi.UpdateArchiveBucketRequest](docs/UpdateArchiveBucketRequest.md)
 - [MorpheusApi.UpdateBackupSettingsRequest](docs/UpdateBackupSettingsRequest.md)
 - [MorpheusApi.UpdateBlueprintPermissionsRequest](docs/UpdateBlueprintPermissionsRequest.md)
 - [MorpheusApi.UpdateBlueprintPermissionsRequestResourcePermission](docs/UpdateBlueprintPermissionsRequestResourcePermission.md)
 - [MorpheusApi.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner](docs/UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md)
 - [MorpheusApi.UpdateCatalogItemType200Response](docs/UpdateCatalogItemType200Response.md)
 - [MorpheusApi.UpdateCatalogItemTypeRequest](docs/UpdateCatalogItemTypeRequest.md)
 - [MorpheusApi.UpdateCatalogItemTypeRequestCatalogItemType](docs/UpdateCatalogItemTypeRequestCatalogItemType.md)
 - [MorpheusApi.UpdateClientsRequest](docs/UpdateClientsRequest.md)
 - [MorpheusApi.UpdateExecuteSchedulesRequest](docs/UpdateExecuteSchedulesRequest.md)
 - [MorpheusApi.UpdateExecuteSchedulesRequestSchedule](docs/UpdateExecuteSchedulesRequestSchedule.md)
 - [MorpheusApi.ValidateAppState200Response](docs/ValidateAppState200Response.md)


## Documentation for Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: Bearer authentication

### cypherAuth-XCToken


- **Type**: API key
- **API key parameter name**: X-Cypher-Token
- **Location**: HTTP header

### cypherAuth-XVToken


- **Type**: API key
- **API key parameter name**: X-Vault-Token
- **Location**: HTTP header

### cypherAuth-XMLease


- **Type**: API key
- **API key parameter name**: X-Morpheus-Lease
- **Location**: HTTP header
