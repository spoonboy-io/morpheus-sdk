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
    instance = new MorpheusApi.BillingZone();
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

  describe('BillingZone', function() {
    it('should create an instance of BillingZone', function() {
      // uncomment below and update the code to test BillingZone
      //var instane = new MorpheusApi.BillingZone();
      //expect(instance).to.be.a(MorpheusApi.BillingZone);
    });

    it('should have the property zoneName (base name: "zoneName")', function() {
      // uncomment below and update the code to test the property zoneName
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property zoneId (base name: "zoneId")', function() {
      // uncomment below and update the code to test the property zoneId
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property zoneUUID (base name: "zoneUUID")', function() {
      // uncomment below and update the code to test the property zoneUUID
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property zoneCode (base name: "zoneCode")', function() {
      // uncomment below and update the code to test the property zoneCode
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property priceUnit (base name: "priceUnit")', function() {
      // uncomment below and update the code to test the property priceUnit
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property computeServers (base name: "computeServers")', function() {
      // uncomment below and update the code to test the property computeServers
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property instances (base name: "instances")', function() {
      // uncomment below and update the code to test the property instances
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property discoveredServers (base name: "discoveredServers")', function() {
      // uncomment below and update the code to test the property discoveredServers
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property loadBalancers (base name: "loadBalancers")', function() {
      // uncomment below and update the code to test the property loadBalancers
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property virtualImages (base name: "virtualImages")', function() {
      // uncomment below and update the code to test the property virtualImages
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property snapshots (base name: "snapshots")', function() {
      // uncomment below and update the code to test the property snapshots
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property price (base name: "price")', function() {
      // uncomment below and update the code to test the property price
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

    it('should have the property cost (base name: "cost")', function() {
      // uncomment below and update the code to test the property cost
      //var instance = new MorpheusApi.BillingZone();
      //expect(instance).to.be();
    });

  });

}));
