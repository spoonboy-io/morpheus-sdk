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
    instance = new MorpheusApi.Issue();
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

  describe('Issue', function() {
    it('should create an instance of Issue', function() {
      // uncomment below and update the code to test Issue
      //var instane = new MorpheusApi.Issue();
      //expect(instance).to.be.a(MorpheusApi.Issue);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property attachmentType (base name: "attachmentType")', function() {
      // uncomment below and update the code to test the property attachmentType
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property app (base name: "app")', function() {
      // uncomment below and update the code to test the property app
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property available (base name: "available")', function() {
      // uncomment below and update the code to test the property available
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property check (base name: "check")', function() {
      // uncomment below and update the code to test the property check
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property checkGroup (base name: "checkGroup")', function() {
      // uncomment below and update the code to test the property checkGroup
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property checkStatus (base name: "checkStatus")', function() {
      // uncomment below and update the code to test the property checkStatus
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property health (base name: "health")', function() {
      // uncomment below and update the code to test the property health
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property inUptime (base name: "inUptime")', function() {
      // uncomment below and update the code to test the property inUptime
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property incident (base name: "incident")', function() {
      // uncomment below and update the code to test the property incident
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property lastCheckTime (base name: "lastCheckTime")', function() {
      // uncomment below and update the code to test the property lastCheckTime
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property lastError (base name: "lastError")', function() {
      // uncomment below and update the code to test the property lastError
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property lastMessage (base name: "lastMessage")', function() {
      // uncomment below and update the code to test the property lastMessage
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property severity (base name: "severity")', function() {
      // uncomment below and update the code to test the property severity
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property severityId (base name: "severityId")', function() {
      // uncomment below and update the code to test the property severityId
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.Issue();
      //expect(instance).to.be();
    });

  });

}));
