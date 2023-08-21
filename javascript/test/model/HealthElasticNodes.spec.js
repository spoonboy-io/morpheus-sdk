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
    instance = new MorpheusApi.HealthElasticNodes();
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

  describe('HealthElasticNodes', function() {
    it('should create an instance of HealthElasticNodes', function() {
      // uncomment below and update the code to test HealthElasticNodes
      //var instane = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be.a(MorpheusApi.HealthElasticNodes);
    });

    it('should have the property ip (base name: "ip")', function() {
      // uncomment below and update the code to test the property ip
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property heapPercent (base name: "heapPercent")', function() {
      // uncomment below and update the code to test the property heapPercent
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property ramPercent (base name: "ramPercent")', function() {
      // uncomment below and update the code to test the property ramPercent
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property cpuCount (base name: "cpuCount")', function() {
      // uncomment below and update the code to test the property cpuCount
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property loadOne (base name: "loadOne")', function() {
      // uncomment below and update the code to test the property loadOne
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property loadFive (base name: "loadFive")', function() {
      // uncomment below and update the code to test the property loadFive
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property loadFifteen (base name: "loadFifteen")', function() {
      // uncomment below and update the code to test the property loadFifteen
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property role (base name: "role")', function() {
      // uncomment below and update the code to test the property role
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property master (base name: "master")', function() {
      // uncomment below and update the code to test the property master
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.HealthElasticNodes();
      //expect(instance).to.be();
    });

  });

}));