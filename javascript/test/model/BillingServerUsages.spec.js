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
    instance = new MorpheusApi.BillingServerUsages();
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

  describe('BillingServerUsages', function() {
    it('should create an instance of BillingServerUsages', function() {
      // uncomment below and update the code to test BillingServerUsages
      //var instane = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be.a(MorpheusApi.BillingServerUsages);
    });

    it('should have the property cost (base name: "cost")', function() {
      // uncomment below and update the code to test the property cost
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property price (base name: "price")', function() {
      // uncomment below and update the code to test the property price
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property createdByUser (base name: "createdByUser")', function() {
      // uncomment below and update the code to test the property createdByUser
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property createdByUserId (base name: "createdByUserId")', function() {
      // uncomment below and update the code to test the property createdByUserId
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property siteId (base name: "siteId")', function() {
      // uncomment below and update the code to test the property siteId
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property siteName (base name: "siteName")', function() {
      // uncomment below and update the code to test the property siteName
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property siteUUID (base name: "siteUUID")', function() {
      // uncomment below and update the code to test the property siteUUID
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property siteCode (base name: "siteCode")', function() {
      // uncomment below and update the code to test the property siteCode
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property currency (base name: "currency")', function() {
      // uncomment below and update the code to test the property currency
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property applicablePrices (base name: "applicablePrices")', function() {
      // uncomment below and update the code to test the property applicablePrices
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property servicePlanId (base name: "servicePlanId")', function() {
      // uncomment below and update the code to test the property servicePlanId
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property servicePlanName (base name: "servicePlanName")', function() {
      // uncomment below and update the code to test the property servicePlanName
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property resourcePoolId (base name: "resourcePoolId")', function() {
      // uncomment below and update the code to test the property resourcePoolId
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

    it('should have the property resourcePoolName (base name: "resourcePoolName")', function() {
      // uncomment below and update the code to test the property resourcePoolName
      //var instance = new MorpheusApi.BillingServerUsages();
      //expect(instance).to.be();
    });

  });

}));
