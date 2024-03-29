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
    instance = new MorpheusApi.ServicePlanConfigRanges();
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

  describe('ServicePlanConfigRanges', function() {
    it('should create an instance of ServicePlanConfigRanges', function() {
      // uncomment below and update the code to test ServicePlanConfigRanges
      //var instane = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be.a(MorpheusApi.ServicePlanConfigRanges);
    });

    it('should have the property minStorage (base name: "minStorage")', function() {
      // uncomment below and update the code to test the property minStorage
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property minPerDiskSize (base name: "minPerDiskSize")', function() {
      // uncomment below and update the code to test the property minPerDiskSize
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property maxPerDiskSize (base name: "maxPerDiskSize")', function() {
      // uncomment below and update the code to test the property maxPerDiskSize
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property minMemory (base name: "minMemory")', function() {
      // uncomment below and update the code to test the property minMemory
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property maxMemory (base name: "maxMemory")', function() {
      // uncomment below and update the code to test the property maxMemory
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property minCores (base name: "minCores")', function() {
      // uncomment below and update the code to test the property minCores
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property maxCores (base name: "maxCores")', function() {
      // uncomment below and update the code to test the property maxCores
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property minSockets (base name: "minSockets")', function() {
      // uncomment below and update the code to test the property minSockets
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property maxSockets (base name: "maxSockets")', function() {
      // uncomment below and update the code to test the property maxSockets
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property minCoresPerSocket (base name: "minCoresPerSocket")', function() {
      // uncomment below and update the code to test the property minCoresPerSocket
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

    it('should have the property maxCoresPerSocket (base name: "maxCoresPerSocket")', function() {
      // uncomment below and update the code to test the property maxCoresPerSocket
      //var instance = new MorpheusApi.ServicePlanConfigRanges();
      //expect(instance).to.be();
    });

  });

}));
