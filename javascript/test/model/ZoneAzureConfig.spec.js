/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.MorpheusApi);
  }
}(this, function(expect, MorpheusApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MorpheusApi.ZoneAzureConfig();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ZoneAzureConfig', function() {
    it('should create an instance of ZoneAzureConfig', function() {
      // uncomment below and update the code to test ZoneAzureConfig
      //var instane = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be.a(MorpheusApi.ZoneAzureConfig);
    });

    it('should have the property subscriberId (base name: "subscriberId")', function() {
      // uncomment below and update the code to test the property subscriberId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property tenantId (base name: "tenantId")', function() {
      // uncomment below and update the code to test the property tenantId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property clientId (base name: "clientId")', function() {
      // uncomment below and update the code to test the property clientId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property clientSecret (base name: "clientSecret")', function() {
      // uncomment below and update the code to test the property clientSecret
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property resourceGroup (base name: "resourceGroup")', function() {
      // uncomment below and update the code to test the property resourceGroup
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property importExisting (base name: "_importExisting")', function() {
      // uncomment below and update the code to test the property importExisting
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property importExisting (base name: "importExisting")', function() {
      // uncomment below and update the code to test the property importExisting
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property inventoryLevel (base name: "inventoryLevel")', function() {
      // uncomment below and update the code to test the property inventoryLevel
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property applianceUrl (base name: "applianceUrl")', function() {
      // uncomment below and update the code to test the property applianceUrl
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property datacenterName (base name: "datacenterName")', function() {
      // uncomment below and update the code to test the property datacenterName
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property networkServerId (base name: "networkServer.id")', function() {
      // uncomment below and update the code to test the property networkServerId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property networkServer (base name: "networkServer")', function() {
      // uncomment below and update the code to test the property networkServer
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property securityMode (base name: "securityMode")', function() {
      // uncomment below and update the code to test the property securityMode
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property certificateProvider (base name: "certificateProvider")', function() {
      // uncomment below and update the code to test the property certificateProvider
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property backupMode (base name: "backupMode")', function() {
      // uncomment below and update the code to test the property backupMode
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property replicationMode (base name: "replicationMode")', function() {
      // uncomment below and update the code to test the property replicationMode
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property dnsIntegrationId (base name: "dnsIntegrationId")', function() {
      // uncomment below and update the code to test the property dnsIntegrationId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property configManagementId (base name: "configManagementId")', function() {
      // uncomment below and update the code to test the property configManagementId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property configCmdbId (base name: "configCmdbId")', function() {
      // uncomment below and update the code to test the property configCmdbId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property securityServer (base name: "securityServer")', function() {
      // uncomment below and update the code to test the property securityServer
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property accountType (base name: "accountType")', function() {
      // uncomment below and update the code to test the property accountType
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceRegistryId (base name: "serviceRegistryId")', function() {
      // uncomment below and update the code to test the property serviceRegistryId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property cloudType (base name: "cloudType")', function() {
      // uncomment below and update the code to test the property cloudType
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property rpcMode (base name: "rpcMode")', function() {
      // uncomment below and update the code to test the property rpcMode
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property diskEncryption (base name: "diskEncryption")', function() {
      // uncomment below and update the code to test the property diskEncryption
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property encryptionSet (base name: "encryptionSet")', function() {
      // uncomment below and update the code to test the property encryptionSet
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property cspTenantId (base name: "cspTenantId")', function() {
      // uncomment below and update the code to test the property cspTenantId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property cspClientId (base name: "cspClientId")', function() {
      // uncomment below and update the code to test the property cspClientId
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property cspClientSecret (base name: "cspClientSecret")', function() {
      // uncomment below and update the code to test the property cspClientSecret
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property cspCustomer (base name: "cspCustomer")', function() {
      // uncomment below and update the code to test the property cspCustomer
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property configCmdbDiscovery (base name: "configCmdbDiscovery")', function() {
      // uncomment below and update the code to test the property configCmdbDiscovery
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property azureCostingMode (base name: "azureCostingMode")', function() {
      // uncomment below and update the code to test the property azureCostingMode
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property clientSecretHash (base name: "clientSecretHash")', function() {
      // uncomment below and update the code to test the property clientSecretHash
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

    it('should have the property cspClientSecretHash (base name: "cspClientSecretHash")', function() {
      // uncomment below and update the code to test the property cspClientSecretHash
      //var instance = new MorpheusApi.ZoneAzureConfig();
      //expect(instance).to.be();
    });

  });

}));
