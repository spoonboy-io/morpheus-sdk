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
    instance = new MorpheusApi.BackupRestore();
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

  describe('BackupRestore', function() {
    it('should create an instance of BackupRestore', function() {
      // uncomment below and update the code to test BackupRestore
      //var instane = new MorpheusApi.BackupRestore();
      //expect(instance).to.be.a(MorpheusApi.BackupRestore);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property backupResultId (base name: "backupResultId")', function() {
      // uncomment below and update the code to test the property backupResultId
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property backupId (base name: "backupId")', function() {
      // uncomment below and update the code to test the property backupId
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property backup (base name: "backup")', function() {
      // uncomment below and update the code to test the property backup
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property containerId (base name: "containerId")', function() {
      // uncomment below and update the code to test the property containerId
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property container (base name: "container")', function() {
      // uncomment below and update the code to test the property container
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property instance (base name: "instance")', function() {
      // uncomment below and update the code to test the property instance
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property restoreToNew (base name: "restoreToNew")', function() {
      // uncomment below and update the code to test the property restoreToNew
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property errorMessage (base name: "errorMessage")', function() {
      // uncomment below and update the code to test the property errorMessage
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property durationMillis (base name: "durationMillis")', function() {
      // uncomment below and update the code to test the property durationMillis
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property externalStatusRef (base name: "externalStatusRef")', function() {
      // uncomment below and update the code to test the property externalStatusRef
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.BackupRestore();
      //expect(instance).to.be();
    });

  });

}));
