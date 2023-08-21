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
    instance = new MorpheusApi.TaskHttpConfig();
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

  describe('TaskHttpConfig', function() {
    it('should create an instance of TaskHttpConfig', function() {
      // uncomment below and update the code to test TaskHttpConfig
      //var instane = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be.a(MorpheusApi.TaskHttpConfig);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property accountId (base name: "accountId")', function() {
      // uncomment below and update the code to test the property accountId
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property taskType (base name: "taskType")', function() {
      // uncomment below and update the code to test the property taskType
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property taskOptions (base name: "taskOptions")', function() {
      // uncomment below and update the code to test the property taskOptions
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property file (base name: "file")', function() {
      // uncomment below and update the code to test the property file
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property resultType (base name: "resultType")', function() {
      // uncomment below and update the code to test the property resultType
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property executeTarget (base name: "executeTarget")', function() {
      // uncomment below and update the code to test the property executeTarget
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property retryable (base name: "retryable")', function() {
      // uncomment below and update the code to test the property retryable
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property retryCount (base name: "retryCount")', function() {
      // uncomment below and update the code to test the property retryCount
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property retryDelaySeconds (base name: "retryDelaySeconds")', function() {
      // uncomment below and update the code to test the property retryDelaySeconds
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property allowCustomConfig (base name: "allowCustomConfig")', function() {
      // uncomment below and update the code to test the property allowCustomConfig
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property credential (base name: "credential")', function() {
      // uncomment below and update the code to test the property credential
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.TaskHttpConfig();
      //expect(instance).to.be();
    });

  });

}));
