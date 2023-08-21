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
    instance = new MorpheusApi.CheckGroup();
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

  describe('CheckGroup', function() {
    it('should create an instance of CheckGroup', function() {
      // uncomment below and update the code to test CheckGroup
      //var instane = new MorpheusApi.CheckGroup();
      //expect(instance).to.be.a(MorpheusApi.CheckGroup);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property instance (base name: "instance")', function() {
      // uncomment below and update the code to test the property instance
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property inUptime (base name: "inUptime")', function() {
      // uncomment below and update the code to test the property inUptime
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastCheckStatus (base name: "lastCheckStatus")', function() {
      // uncomment below and update the code to test the property lastCheckStatus
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastWarningDate (base name: "lastWarningDate")', function() {
      // uncomment below and update the code to test the property lastWarningDate
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastErrorDate (base name: "lastErrorDate")', function() {
      // uncomment below and update the code to test the property lastErrorDate
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastSuccessDate (base name: "lastSuccessDate")', function() {
      // uncomment below and update the code to test the property lastSuccessDate
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastRunDate (base name: "lastRunDate")', function() {
      // uncomment below and update the code to test the property lastRunDate
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastError (base name: "lastError")', function() {
      // uncomment below and update the code to test the property lastError
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property outageTime (base name: "outageTime")', function() {
      // uncomment below and update the code to test the property outageTime
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastTimer (base name: "lastTimer")', function() {
      // uncomment below and update the code to test the property lastTimer
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property health (base name: "health")', function() {
      // uncomment below and update the code to test the property health
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property history (base name: "history")', function() {
      // uncomment below and update the code to test the property history
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property minHappy (base name: "minHappy")', function() {
      // uncomment below and update the code to test the property minHappy
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastMetric (base name: "lastMetric")', function() {
      // uncomment below and update the code to test the property lastMetric
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property severity (base name: "severity")', function() {
      // uncomment below and update the code to test the property severity
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property createIncident (base name: "createIncident")', function() {
      // uncomment below and update the code to test the property createIncident
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property muted (base name: "muted")', function() {
      // uncomment below and update the code to test the property muted
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "createdBy")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property availability (base name: "availability")', function() {
      // uncomment below and update the code to test the property availability
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property checkType (base name: "checkType")', function() {
      // uncomment below and update the code to test the property checkType
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

    it('should have the property checks (base name: "checks")', function() {
      // uncomment below and update the code to test the property checks
      //var instance = new MorpheusApi.CheckGroup();
      //expect(instance).to.be();
    });

  });

}));
