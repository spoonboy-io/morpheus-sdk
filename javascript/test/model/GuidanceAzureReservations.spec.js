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
    instance = new MorpheusApi.GuidanceAzureReservations();
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

  describe('GuidanceAzureReservations', function() {
    it('should create an instance of GuidanceAzureReservations', function() {
      // uncomment below and update the code to test GuidanceAzureReservations
      //var instane = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be.a(MorpheusApi.GuidanceAzureReservations);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionCategory (base name: "actionCategory")', function() {
      // uncomment below and update the code to test the property actionCategory
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionMessage (base name: "actionMessage")', function() {
      // uncomment below and update the code to test the property actionMessage
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionTitle (base name: "actionTitle")', function() {
      // uncomment below and update the code to test the property actionTitle
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionType (base name: "actionType")', function() {
      // uncomment below and update the code to test the property actionType
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionValue (base name: "actionValue")', function() {
      // uncomment below and update the code to test the property actionValue
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionValueType (base name: "actionValueType")', function() {
      // uncomment below and update the code to test the property actionValueType
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property actionPlanId (base name: "actionPlanId")', function() {
      // uncomment below and update the code to test the property actionPlanId
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property statusMessage (base name: "statusMessage")', function() {
      // uncomment below and update the code to test the property statusMessage
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property accountId (base name: "accountId")', function() {
      // uncomment below and update the code to test the property accountId
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property userId (base name: "userId")', function() {
      // uncomment below and update the code to test the property userId
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property siteId (base name: "siteId")', function() {
      // uncomment below and update the code to test the property siteId
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property zone (base name: "zone")', function() {
      // uncomment below and update the code to test the property zone
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property state (base name: "state")', function() {
      // uncomment below and update the code to test the property state
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property stateMessage (base name: "stateMessage")', function() {
      // uncomment below and update the code to test the property stateMessage
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property severity (base name: "severity")', function() {
      // uncomment below and update the code to test the property severity
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property resolved (base name: "resolved")', function() {
      // uncomment below and update the code to test the property resolved
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property resolvedMessage (base name: "resolvedMessage")', function() {
      // uncomment below and update the code to test the property resolvedMessage
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property refName (base name: "refName")', function() {
      // uncomment below and update the code to test the property refName
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property savings (base name: "savings")', function() {
      // uncomment below and update the code to test the property savings
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.GuidanceAzureReservations();
      //expect(instance).to.be();
    });

  });

}));
