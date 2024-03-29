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
    instance = new MorpheusApi.TaskType();
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

  describe('TaskType', function() {
    it('should create an instance of TaskType', function() {
      // uncomment below and update the code to test TaskType
      //var instane = new MorpheusApi.TaskType();
      //expect(instance).to.be.a(MorpheusApi.TaskType);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property scriptable (base name: "scriptable")', function() {
      // uncomment below and update the code to test the property scriptable
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property hasResults (base name: "hasResults")', function() {
      // uncomment below and update the code to test the property hasResults
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property allowExecuteLocal (base name: "allowExecuteLocal")', function() {
      // uncomment below and update the code to test the property allowExecuteLocal
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property allowExecuteRemote (base name: "allowExecuteRemote")', function() {
      // uncomment below and update the code to test the property allowExecuteRemote
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property allowExecuteResource (base name: "allowExecuteResource")', function() {
      // uncomment below and update the code to test the property allowExecuteResource
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property allowLocalRepo (base name: "allowLocalRepo")', function() {
      // uncomment below and update the code to test the property allowLocalRepo
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property allowRemoteKeyAuth (base name: "allowRemoteKeyAuth")', function() {
      // uncomment below and update the code to test the property allowRemoteKeyAuth
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

    it('should have the property optionTypes (base name: "optionTypes")', function() {
      // uncomment below and update the code to test the property optionTypes
      //var instance = new MorpheusApi.TaskType();
      //expect(instance).to.be();
    });

  });

}));
