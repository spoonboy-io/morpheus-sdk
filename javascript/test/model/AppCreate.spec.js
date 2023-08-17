/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
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
    instance = new MorpheusApi.AppCreate();
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

  describe('AppCreate', function() {
    it('should create an instance of AppCreate', function() {
      // uncomment below and update the code to test AppCreate
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be.a(MorpheusApi.AppCreate);
    });

    it('should have the property templateId (base name: "templateId")', function() {
      // uncomment below and update the code to test the property templateId
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property blueprintId (base name: "blueprintId")', function() {
      // uncomment below and update the code to test the property blueprintId
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property group (base name: "group")', function() {
      // uncomment below and update the code to test the property group
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property defaultCloud (base name: "defaultCloud")', function() {
      // uncomment below and update the code to test the property defaultCloud
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property environment (base name: "environment")', function() {
      // uncomment below and update the code to test the property environment
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

    it('should have the property tiers (base name: "tiers")', function() {
      // uncomment below and update the code to test the property tiers
      //var instance = new MorpheusApi.AppCreate();
      //expect(instance).to.be();
    });

  });

}));
