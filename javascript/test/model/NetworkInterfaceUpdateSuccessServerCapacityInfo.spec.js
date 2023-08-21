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
    instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
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

  describe('NetworkInterfaceUpdateSuccessServerCapacityInfo', function() {
    it('should create an instance of NetworkInterfaceUpdateSuccessServerCapacityInfo', function() {
      // uncomment below and update the code to test NetworkInterfaceUpdateSuccessServerCapacityInfo
      //var instane = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be.a(MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property maxMemory (base name: "maxMemory")', function() {
      // uncomment below and update the code to test the property maxMemory
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property server (base name: "server")', function() {
      // uncomment below and update the code to test the property server
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property usedStorage (base name: "usedStorage")', function() {
      // uncomment below and update the code to test the property usedStorage
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property version (base name: "version")', function() {
      // uncomment below and update the code to test the property version
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property maxCpu (base name: "maxCpu")', function() {
      // uncomment below and update the code to test the property maxCpu
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property usedCores (base name: "usedCores")', function() {
      // uncomment below and update the code to test the property usedCores
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property usedMemory (base name: "usedMemory")', function() {
      // uncomment below and update the code to test the property usedMemory
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property maxCores (base name: "maxCores")', function() {
      // uncomment below and update the code to test the property maxCores
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.NetworkInterfaceUpdateSuccessServerCapacityInfo();
      //expect(instance).to.be();
    });

  });

}));
