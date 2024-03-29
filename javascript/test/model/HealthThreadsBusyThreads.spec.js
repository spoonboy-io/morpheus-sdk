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
    instance = new MorpheusApi.HealthThreadsBusyThreads();
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

  describe('HealthThreadsBusyThreads', function() {
    it('should create an instance of HealthThreadsBusyThreads', function() {
      // uncomment below and update the code to test HealthThreadsBusyThreads
      //var instane = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be.a(MorpheusApi.HealthThreadsBusyThreads);
    });

    it('should have the property threadId (base name: "threadId")', function() {
      // uncomment below and update the code to test the property threadId
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property cpuTime (base name: "cpuTime")', function() {
      // uncomment below and update the code to test the property cpuTime
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property blockedTime (base name: "blockedTime")', function() {
      // uncomment below and update the code to test the property blockedTime
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property lockName (base name: "lockName")', function() {
      // uncomment below and update the code to test the property lockName
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property lockOwnerId (base name: "lockOwnerId")', function() {
      // uncomment below and update the code to test the property lockOwnerId
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property lockOwnerName (base name: "lockOwnerName")', function() {
      // uncomment below and update the code to test the property lockOwnerName
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property state (base name: "state")', function() {
      // uncomment below and update the code to test the property state
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property waitedCount (base name: "waitedCount")', function() {
      // uncomment below and update the code to test the property waitedCount
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property waitedTime (base name: "waitedTime")', function() {
      // uncomment below and update the code to test the property waitedTime
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property isInNative (base name: "isInNative")', function() {
      // uncomment below and update the code to test the property isInNative
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property isSuspended (base name: "isSuspended")', function() {
      // uncomment below and update the code to test the property isSuspended
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property lockedMonitors (base name: "lockedMonitors")', function() {
      // uncomment below and update the code to test the property lockedMonitors
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property lockedSynchronizers (base name: "lockedSynchronizers")', function() {
      // uncomment below and update the code to test the property lockedSynchronizers
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property lockInfo (base name: "lockInfo")', function() {
      // uncomment below and update the code to test the property lockInfo
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property currentLines (base name: "currentLines")', function() {
      // uncomment below and update the code to test the property currentLines
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

    it('should have the property cpuPercent (base name: "cpuPercent")', function() {
      // uncomment below and update the code to test the property cpuPercent
      //var instance = new MorpheusApi.HealthThreadsBusyThreads();
      //expect(instance).to.be();
    });

  });

}));
