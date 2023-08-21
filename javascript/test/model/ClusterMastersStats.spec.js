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
    instance = new MorpheusApi.ClusterMastersStats();
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

  describe('ClusterMastersStats', function() {
    it('should create an instance of ClusterMastersStats', function() {
      // uncomment below and update the code to test ClusterMastersStats
      //var instane = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be.a(MorpheusApi.ClusterMastersStats);
    });

    it('should have the property usedStorage (base name: "usedStorage")', function() {
      // uncomment below and update the code to test the property usedStorage
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

    it('should have the property reservedStorage (base name: "reservedStorage")', function() {
      // uncomment below and update the code to test the property reservedStorage
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

    it('should have the property usedMemory (base name: "usedMemory")', function() {
      // uncomment below and update the code to test the property usedMemory
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

    it('should have the property reservedMemory (base name: "reservedMemory")', function() {
      // uncomment below and update the code to test the property reservedMemory
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

    it('should have the property maxMemory (base name: "maxMemory")', function() {
      // uncomment below and update the code to test the property maxMemory
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

    it('should have the property cpuUsage (base name: "cpuUsage")', function() {
      // uncomment below and update the code to test the property cpuUsage
      //var instance = new MorpheusApi.ClusterMastersStats();
      //expect(instance).to.be();
    });

  });

}));
