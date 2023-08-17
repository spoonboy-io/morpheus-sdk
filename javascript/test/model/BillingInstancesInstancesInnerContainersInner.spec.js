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
    instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
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

  describe('BillingInstancesInstancesInnerContainersInner', function() {
    it('should create an instance of BillingInstancesInstancesInnerContainersInner', function() {
      // uncomment below and update the code to test BillingInstancesInstancesInnerContainersInner
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be.a(MorpheusApi.BillingInstancesInstancesInnerContainersInner);
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property refUUID (base name: "refUUID")', function() {
      // uncomment below and update the code to test the property refUUID
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property cost (base name: "cost")', function() {
      // uncomment below and update the code to test the property cost
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property price (base name: "price")', function() {
      // uncomment below and update the code to test the property price
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property numUnits (base name: "numUnits")', function() {
      // uncomment below and update the code to test the property numUnits
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property unit (base name: "unit")', function() {
      // uncomment below and update the code to test the property unit
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property currency (base name: "currency")', function() {
      // uncomment below and update the code to test the property currency
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property usages (base name: "usages")', function() {
      // uncomment below and update the code to test the property usages
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property numUsages (base name: "numUsages")', function() {
      // uncomment below and update the code to test the property numUsages
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property totalUsages (base name: "totalUsages")', function() {
      // uncomment below and update the code to test the property totalUsages
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property hasMoreUsages (base name: "hasMoreUsages")', function() {
      // uncomment below and update the code to test the property hasMoreUsages
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property foundPricing (base name: "foundPricing")', function() {
      // uncomment below and update the code to test the property foundPricing
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property serverId (base name: "serverId")', function() {
      // uncomment below and update the code to test the property serverId
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property serverUUID (base name: "serverUUID")', function() {
      // uncomment below and update the code to test the property serverUUID
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property serverUniqueId (base name: "serverUniqueId")', function() {
      // uncomment below and update the code to test the property serverUniqueId
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property serverExternalId (base name: "serverExternalId")', function() {
      // uncomment below and update the code to test the property serverExternalId
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property serverInternalId (base name: "serverInternalId")', function() {
      // uncomment below and update the code to test the property serverInternalId
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property resourcePoolId (base name: "resourcePoolId")', function() {
      // uncomment below and update the code to test the property resourcePoolId
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

    it('should have the property resourcePoolName (base name: "resourcePoolName")', function() {
      // uncomment below and update the code to test the property resourcePoolName
      //var instance = new MorpheusApi.BillingInstancesInstancesInnerContainersInner();
      //expect(instance).to.be();
    });

  });

}));
