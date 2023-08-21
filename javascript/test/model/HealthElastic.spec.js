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
    instance = new MorpheusApi.HealthElastic();
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

  describe('HealthElastic', function() {
    it('should create an instance of HealthElastic', function() {
      // uncomment below and update the code to test HealthElastic
      //var instane = new MorpheusApi.HealthElastic();
      //expect(instance).to.be.a(MorpheusApi.HealthElastic);
    });

    it('should have the property success (base name: "success")', function() {
      // uncomment below and update the code to test the property success
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

    it('should have the property master (base name: "master")', function() {
      // uncomment below and update the code to test the property master
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

    it('should have the property nodes (base name: "nodes")', function() {
      // uncomment below and update the code to test the property nodes
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

    it('should have the property stats (base name: "stats")', function() {
      // uncomment below and update the code to test the property stats
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

    it('should have the property indices (base name: "indices")', function() {
      // uncomment below and update the code to test the property indices
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

    it('should have the property badIndices (base name: "badIndices")', function() {
      // uncomment below and update the code to test the property badIndices
      //var instance = new MorpheusApi.HealthElastic();
      //expect(instance).to.be();
    });

  });

}));