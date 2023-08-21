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
    instance = new MorpheusApi.InstanceTypeLayoutCreate();
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

  describe('InstanceTypeLayoutCreate', function() {
    it('should create an instance of InstanceTypeLayoutCreate', function() {
      // uncomment below and update the code to test InstanceTypeLayoutCreate
      //var instane = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be.a(MorpheusApi.InstanceTypeLayoutCreate);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property instanceVersion (base name: "instanceVersion")', function() {
      // uncomment below and update the code to test the property instanceVersion
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property creatable (base name: "creatable")', function() {
      // uncomment below and update the code to test the property creatable
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property provisionTypeCode (base name: "provisionTypeCode")', function() {
      // uncomment below and update the code to test the property provisionTypeCode
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property memoryRequirement (base name: "memoryRequirement")', function() {
      // uncomment below and update the code to test the property memoryRequirement
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property hasAutoScale (base name: "hasAutoScale")', function() {
      // uncomment below and update the code to test the property hasAutoScale
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property supportsConvertToManaged (base name: "supportsConvertToManaged")', function() {
      // uncomment below and update the code to test the property supportsConvertToManaged
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property containerTypes (base name: "containerTypes")', function() {
      // uncomment below and update the code to test the property containerTypes
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property optionTypes (base name: "optionTypes")', function() {
      // uncomment below and update the code to test the property optionTypes
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property specTemplates (base name: "specTemplates")', function() {
      // uncomment below and update the code to test the property specTemplates
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property environmentVariables (base name: "environmentVariables")', function() {
      // uncomment below and update the code to test the property environmentVariables
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property priceSets (base name: "priceSets")', function() {
      // uncomment below and update the code to test the property priceSets
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

    it('should have the property permissions (base name: "permissions")', function() {
      // uncomment below and update the code to test the property permissions
      //var instance = new MorpheusApi.InstanceTypeLayoutCreate();
      //expect(instance).to.be();
    });

  });

}));
