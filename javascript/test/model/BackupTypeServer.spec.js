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
    instance = new MorpheusApi.BackupTypeServer();
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

  describe('BackupTypeServer', function() {
    it('should create an instance of BackupTypeServer', function() {
      // uncomment below and update the code to test BackupTypeServer
      //var instane = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be.a(MorpheusApi.BackupTypeServer);
    });

    it('should have the property locationType (base name: "locationType")', function() {
      // uncomment below and update the code to test the property locationType
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property serverId (base name: "serverId")', function() {
      // uncomment below and update the code to test the property serverId
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property backupType (base name: "backupType")', function() {
      // uncomment below and update the code to test the property backupType
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property jobAction (base name: "jobAction")', function() {
      // uncomment below and update the code to test the property jobAction
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property jobId (base name: "jobId")', function() {
      // uncomment below and update the code to test the property jobId
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property jobName (base name: "jobName")', function() {
      // uncomment below and update the code to test the property jobName
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property jobSchedule (base name: "jobSchedule")', function() {
      // uncomment below and update the code to test the property jobSchedule
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

    it('should have the property retentionCount (base name: "retentionCount")', function() {
      // uncomment below and update the code to test the property retentionCount
      //var instance = new MorpheusApi.BackupTypeServer();
      //expect(instance).to.be();
    });

  });

}));