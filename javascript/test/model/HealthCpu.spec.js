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
    instance = new MorpheusApi.HealthCpu();
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

  describe('HealthCpu', function() {
    it('should create an instance of HealthCpu', function() {
      // uncomment below and update the code to test HealthCpu
      //var instane = new MorpheusApi.HealthCpu();
      //expect(instance).to.be.a(MorpheusApi.HealthCpu);
    });

    it('should have the property success (base name: "success")', function() {
      // uncomment below and update the code to test the property success
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

    it('should have the property cpuLoad (base name: "cpuLoad")', function() {
      // uncomment below and update the code to test the property cpuLoad
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

    it('should have the property cpuTotalLoad (base name: "cpuTotalLoad")', function() {
      // uncomment below and update the code to test the property cpuTotalLoad
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

    it('should have the property processorCount (base name: "processorCount")', function() {
      // uncomment below and update the code to test the property processorCount
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

    it('should have the property processTime (base name: "processTime")', function() {
      // uncomment below and update the code to test the property processTime
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

    it('should have the property systemLoad (base name: "systemLoad")', function() {
      // uncomment below and update the code to test the property systemLoad
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.HealthCpu();
      //expect(instance).to.be();
    });

  });

}));
