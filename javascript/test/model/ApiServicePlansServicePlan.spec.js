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
    instance = new MorpheusApi.ApiServicePlansServicePlan();
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

  describe('ApiServicePlansServicePlan', function() {
    it('should create an instance of ApiServicePlansServicePlan', function() {
      // uncomment below and update the code to test ApiServicePlansServicePlan
      //var instane = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be.a(MorpheusApi.ApiServicePlansServicePlan);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property editable (base name: "editable")', function() {
      // uncomment below and update the code to test the property editable
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property maxStorage (base name: "maxStorage")', function() {
      // uncomment below and update the code to test the property maxStorage
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property maxMemory (base name: "maxMemory")', function() {
      // uncomment below and update the code to test the property maxMemory
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property maxCores (base name: "maxCores")', function() {
      // uncomment below and update the code to test the property maxCores
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property maxDisks (base name: "maxDisks")', function() {
      // uncomment below and update the code to test the property maxDisks
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property provisionType (base name: "provisionType")', function() {
      // uncomment below and update the code to test the property provisionType
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property customCores (base name: "customCores")', function() {
      // uncomment below and update the code to test the property customCores
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property customMaxStorage (base name: "customMaxStorage")', function() {
      // uncomment below and update the code to test the property customMaxStorage
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property customMaxDataStorage (base name: "customMaxDataStorage")', function() {
      // uncomment below and update the code to test the property customMaxDataStorage
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property customMaxMemory (base name: "customMaxMemory")', function() {
      // uncomment below and update the code to test the property customMaxMemory
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property addVolumes (base name: "addVolumes")', function() {
      // uncomment below and update the code to test the property addVolumes
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property sortOrder (base name: "sortOrder")', function() {
      // uncomment below and update the code to test the property sortOrder
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property priceSets (base name: "priceSets")', function() {
      // uncomment below and update the code to test the property priceSets
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.ApiServicePlansServicePlan();
      //expect(instance).to.be();
    });

  });

}));
