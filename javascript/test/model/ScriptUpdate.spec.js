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
    instance = new MorpheusApi.ScriptUpdate();
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

  describe('ScriptUpdate', function() {
    it('should create an instance of ScriptUpdate', function() {
      // uncomment below and update the code to test ScriptUpdate
      //var instane = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be.a(MorpheusApi.ScriptUpdate);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property scriptVersion (base name: "scriptVersion")', function() {
      // uncomment below and update the code to test the property scriptVersion
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property scriptPhase (base name: "scriptPhase")', function() {
      // uncomment below and update the code to test the property scriptPhase
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property scriptType (base name: "scriptType")', function() {
      // uncomment below and update the code to test the property scriptType
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property script (base name: "script")', function() {
      // uncomment below and update the code to test the property script
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property runAsUser (base name: "runAsUser")', function() {
      // uncomment below and update the code to test the property runAsUser
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

    it('should have the property sudoUser (base name: "sudoUser")', function() {
      // uncomment below and update the code to test the property sudoUser
      //var instance = new MorpheusApi.ScriptUpdate();
      //expect(instance).to.be();
    });

  });

}));
