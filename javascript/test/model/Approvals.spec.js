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
    instance = new MorpheusApi.Approvals();
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

  describe('Approvals', function() {
    it('should create an instance of Approvals', function() {
      // uncomment below and update the code to test Approvals
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be.a(MorpheusApi.Approvals);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property externalName (base name: "externalName")', function() {
      // uncomment below and update the code to test the property externalName
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property requestType (base name: "requestType")', function() {
      // uncomment below and update the code to test the property requestType
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property approver (base name: "approver")', function() {
      // uncomment below and update the code to test the property approver
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property accountIntegration (base name: "accountIntegration")', function() {
      // uncomment below and update the code to test the property accountIntegration
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property errorMessage (base name: "errorMessage")', function() {
      // uncomment below and update the code to test the property errorMessage
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

    it('should have the property requestBy (base name: "requestBy")', function() {
      // uncomment below and update the code to test the property requestBy
      //var instance = new MorpheusApi.Approvals();
      //expect(instance).to.be();
    });

  });

}));
