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
    instance = new MorpheusApi.ContainerType();
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

  describe('ContainerType', function() {
    it('should create an instance of ContainerType', function() {
      // uncomment below and update the code to test ContainerType
      //var instane = new MorpheusApi.ContainerType();
      //expect(instance).to.be.a(MorpheusApi.ContainerType);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property shortName (base name: "shortName")', function() {
      // uncomment below and update the code to test the property shortName
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property containerVersion (base name: "containerVersion")', function() {
      // uncomment below and update the code to test the property containerVersion
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property provisionType (base name: "provisionType")', function() {
      // uncomment below and update the code to test the property provisionType
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property virtualImage (base name: "virtualImage")', function() {
      // uncomment below and update the code to test the property virtualImage
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property containerPorts (base name: "containerPorts")', function() {
      // uncomment below and update the code to test the property containerPorts
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property containerScripts (base name: "containerScripts")', function() {
      // uncomment below and update the code to test the property containerScripts
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property containerTemplates (base name: "containerTemplates")', function() {
      // uncomment below and update the code to test the property containerTemplates
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

    it('should have the property environmentVariables (base name: "environmentVariables")', function() {
      // uncomment below and update the code to test the property environmentVariables
      //var instance = new MorpheusApi.ContainerType();
      //expect(instance).to.be();
    });

  });

}));
