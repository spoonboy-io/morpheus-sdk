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
    instance = new MorpheusApi.ProcessEvents();
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

  describe('ProcessEvents', function() {
    it('should create an instance of ProcessEvents', function() {
      // uncomment below and update the code to test ProcessEvents
      //var instane = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be.a(MorpheusApi.ProcessEvents);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property processId (base name: "processId")', function() {
      // uncomment below and update the code to test the property processId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property accountId (base name: "accountId")', function() {
      // uncomment below and update the code to test the property accountId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property uniqueId (base name: "uniqueId")', function() {
      // uncomment below and update the code to test the property uniqueId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property processType (base name: "processType")', function() {
      // uncomment below and update the code to test the property processType
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property subType (base name: "subType")', function() {
      // uncomment below and update the code to test the property subType
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property subId (base name: "subId")', function() {
      // uncomment below and update the code to test the property subId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property zoneId (base name: "zoneId")', function() {
      // uncomment below and update the code to test the property zoneId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property integrationId (base name: "integrationId")', function() {
      // uncomment below and update the code to test the property integrationId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property instanceId (base name: "instanceId")', function() {
      // uncomment below and update the code to test the property instanceId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property containerId (base name: "containerId")', function() {
      // uncomment below and update the code to test the property containerId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property serverId (base name: "serverId")', function() {
      // uncomment below and update the code to test the property serverId
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property containerName (base name: "containerName")', function() {
      // uncomment below and update the code to test the property containerName
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property displayName (base name: "displayName")', function() {
      // uncomment below and update the code to test the property displayName
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property reason (base name: "reason")', function() {
      // uncomment below and update the code to test the property reason
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property percent (base name: "percent")', function() {
      // uncomment below and update the code to test the property percent
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property statusEta (base name: "statusEta")', function() {
      // uncomment below and update the code to test the property statusEta
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property message (base name: "message")', function() {
      // uncomment below and update the code to test the property message
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property output (base name: "output")', function() {
      // uncomment below and update the code to test the property output
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property error (base name: "error")', function() {
      // uncomment below and update the code to test the property error
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property duration (base name: "duration")', function() {
      // uncomment below and update the code to test the property duration
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "createdBy")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

    it('should have the property updatedBy (base name: "updatedBy")', function() {
      // uncomment below and update the code to test the property updatedBy
      //var instance = new MorpheusApi.ProcessEvents();
      //expect(instance).to.be();
    });

  });

}));
