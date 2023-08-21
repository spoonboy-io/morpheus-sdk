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
    instance = new MorpheusApi.HealthDatabaseInnodbStats();
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

  describe('HealthDatabaseInnodbStats', function() {
    it('should create an instance of HealthDatabaseInnodbStats', function() {
      // uncomment below and update the code to test HealthDatabaseInnodbStats
      //var instane = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be.a(MorpheusApi.HealthDatabaseInnodbStats);
    });

    it('should have the property largeMemory (base name: "largeMemory")', function() {
      // uncomment below and update the code to test the property largeMemory
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property dictionaryMemory (base name: "dictionaryMemory")', function() {
      // uncomment below and update the code to test the property dictionaryMemory
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property bufferPoolSize (base name: "bufferPoolSize")', function() {
      // uncomment below and update the code to test the property bufferPoolSize
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property freeBuffers (base name: "freeBuffers")', function() {
      // uncomment below and update the code to test the property freeBuffers
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property databasePages (base name: "databasePages")', function() {
      // uncomment below and update the code to test the property databasePages
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property oldPages (base name: "oldPages")', function() {
      // uncomment below and update the code to test the property oldPages
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property pendingReads (base name: "pendingReads")', function() {
      // uncomment below and update the code to test the property pendingReads
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property insertsPerSecond (base name: "insertsPerSecond")', function() {
      // uncomment below and update the code to test the property insertsPerSecond
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property updatesPerSecond (base name: "updatesPerSecond")', function() {
      // uncomment below and update the code to test the property updatesPerSecond
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property deletesPerSecond (base name: "deletesPerSecond")', function() {
      // uncomment below and update the code to test the property deletesPerSecond
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property readsPerSecond (base name: "readsPerSecond")', function() {
      // uncomment below and update the code to test the property readsPerSecond
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

    it('should have the property bufferHitRate (base name: "bufferHitRate")', function() {
      // uncomment below and update the code to test the property bufferHitRate
      //var instance = new MorpheusApi.HealthDatabaseInnodbStats();
      //expect(instance).to.be();
    });

  });

}));