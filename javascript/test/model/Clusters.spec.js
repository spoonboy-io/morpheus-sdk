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
    instance = new MorpheusApi.Clusters();
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

  describe('Clusters', function() {
    it('should create an instance of Clusters', function() {
      // uncomment below and update the code to test Clusters
      //var instane = new MorpheusApi.Clusters();
      //expect(instance).to.be.a(MorpheusApi.Clusters);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property location (base name: "location")', function() {
      // uncomment below and update the code to test the property location
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceUrl (base name: "serviceUrl")', function() {
      // uncomment below and update the code to test the property serviceUrl
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceHost (base name: "serviceHost")', function() {
      // uncomment below and update the code to test the property serviceHost
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property servicePath (base name: "servicePath")', function() {
      // uncomment below and update the code to test the property servicePath
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceHostname (base name: "serviceHostname")', function() {
      // uncomment below and update the code to test the property serviceHostname
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property servicePort (base name: "servicePort")', function() {
      // uncomment below and update the code to test the property servicePort
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceUsername (base name: "serviceUsername")', function() {
      // uncomment below and update the code to test the property serviceUsername
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property servicePassword (base name: "servicePassword")', function() {
      // uncomment below and update the code to test the property servicePassword
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property servicePasswordHash (base name: "servicePasswordHash")', function() {
      // uncomment below and update the code to test the property servicePasswordHash
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceToken (base name: "serviceToken")', function() {
      // uncomment below and update the code to test the property serviceToken
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceTokenHash (base name: "serviceTokenHash")', function() {
      // uncomment below and update the code to test the property serviceTokenHash
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceAccess (base name: "serviceAccess")', function() {
      // uncomment below and update the code to test the property serviceAccess
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceAccessHash (base name: "serviceAccessHash")', function() {
      // uncomment below and update the code to test the property serviceAccessHash
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceCert (base name: "serviceCert")', function() {
      // uncomment below and update the code to test the property serviceCert
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceCertHash (base name: "serviceCertHash")', function() {
      // uncomment below and update the code to test the property serviceCertHash
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceVersion (base name: "serviceVersion")', function() {
      // uncomment below and update the code to test the property serviceVersion
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property searchDomains (base name: "searchDomains")', function() {
      // uncomment below and update the code to test the property searchDomains
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property enableInternalDns (base name: "enableInternalDns")', function() {
      // uncomment below and update the code to test the property enableInternalDns
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property datacenterId (base name: "datacenterId")', function() {
      // uncomment below and update the code to test the property datacenterId
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property statusDate (base name: "statusDate")', function() {
      // uncomment below and update the code to test the property statusDate
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property statusMessage (base name: "statusMessage")', function() {
      // uncomment below and update the code to test the property statusMessage
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property inventoryLevel (base name: "inventoryLevel")', function() {
      // uncomment below and update the code to test the property inventoryLevel
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property lastSync (base name: "lastSync")', function() {
      // uncomment below and update the code to test the property lastSync
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property nextRunDate (base name: "nextRunDate")', function() {
      // uncomment below and update the code to test the property nextRunDate
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property lastSyncDuration (base name: "lastSyncDuration")', function() {
      // uncomment below and update the code to test the property lastSyncDuration
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property managed (base name: "managed")', function() {
      // uncomment below and update the code to test the property managed
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property serviceEntry (base name: "serviceEntry")', function() {
      // uncomment below and update the code to test the property serviceEntry
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "createdBy")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property userGroup (base name: "userGroup")', function() {
      // uncomment below and update the code to test the property userGroup
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property layout (base name: "layout")', function() {
      // uncomment below and update the code to test the property layout
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property owner (base name: "owner")', function() {
      // uncomment below and update the code to test the property owner
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property servers (base name: "servers")', function() {
      // uncomment below and update the code to test the property servers
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property accounts (base name: "accounts")', function() {
      // uncomment below and update the code to test the property accounts
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property integrations (base name: "integrations")', function() {
      // uncomment below and update the code to test the property integrations
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property site (base name: "site")', function() {
      // uncomment below and update the code to test the property site
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property zone (base name: "zone")', function() {
      // uncomment below and update the code to test the property zone
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property workerStats (base name: "workerStats")', function() {
      // uncomment below and update the code to test the property workerStats
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.Clusters();
      //expect(instance).to.be();
    });

  });

}));