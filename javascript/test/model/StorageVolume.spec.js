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
    instance = new MorpheusApi.StorageVolume();
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

  describe('StorageVolume', function() {
    it('should create an instance of StorageVolume', function() {
      // uncomment below and update the code to test StorageVolume
      //var instane = new MorpheusApi.StorageVolume();
      //expect(instance).to.be.a(MorpheusApi.StorageVolume);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property controllerId (base name: "controllerId")', function() {
      // uncomment below and update the code to test the property controllerId
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property controllerMountPoint (base name: "controllerMountPoint")', function() {
      // uncomment below and update the code to test the property controllerMountPoint
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property resizeable (base name: "resizeable")', function() {
      // uncomment below and update the code to test the property resizeable
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property rootVolume (base name: "rootVolume")', function() {
      // uncomment below and update the code to test the property rootVolume
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property unitNumber (base name: "unitNumber")', function() {
      // uncomment below and update the code to test the property unitNumber
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property deviceName (base name: "deviceName")', function() {
      // uncomment below and update the code to test the property deviceName
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property deviceDisplayName (base name: "deviceDisplayName")', function() {
      // uncomment below and update the code to test the property deviceDisplayName
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property typeId (base name: "typeId")', function() {
      // uncomment below and update the code to test the property typeId
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property statusMessage (base name: "statusMessage")', function() {
      // uncomment below and update the code to test the property statusMessage
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property configurableIOPS (base name: "configurableIOPS")', function() {
      // uncomment below and update the code to test the property configurableIOPS
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property displayOrder (base name: "displayOrder")', function() {
      // uncomment below and update the code to test the property displayOrder
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property maxIOPS (base name: "maxIOPS")', function() {
      // uncomment below and update the code to test the property maxIOPS
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property uuid (base name: "uuid")', function() {
      // uncomment below and update the code to test the property uuid
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property zone (base name: "zone")', function() {
      // uncomment below and update the code to test the property zone
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property zoneId (base name: "zoneId")', function() {
      // uncomment below and update the code to test the property zoneId
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property datastore (base name: "datastore")', function() {
      // uncomment below and update the code to test the property datastore
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property datastoreId (base name: "datastoreId")', function() {
      // uncomment below and update the code to test the property datastoreId
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property storageGroup (base name: "storageGroup")', function() {
      // uncomment below and update the code to test the property storageGroup
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property namespace (base name: "namespace")', function() {
      // uncomment below and update the code to test the property namespace
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property storageServer (base name: "storageServer")', function() {
      // uncomment below and update the code to test the property storageServer
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property source (base name: "source")', function() {
      // uncomment below and update the code to test the property source
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

    it('should have the property owner (base name: "owner")', function() {
      // uncomment below and update the code to test the property owner
      //var instance = new MorpheusApi.StorageVolume();
      //expect(instance).to.be();
    });

  });

}));
