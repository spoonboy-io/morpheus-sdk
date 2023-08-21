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
    instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
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

  describe('GuidanceAzureReservationsConfigServicesAzureVmsSummary', function() {
    it('should create an instance of GuidanceAzureReservationsConfigServicesAzureVmsSummary', function() {
      // uncomment below and update the code to test GuidanceAzureReservationsConfigServicesAzureVmsSummary
      //var instane = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be.a(MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary);
    });

    it('should have the property totalSavings (base name: "totalSavings")', function() {
      // uncomment below and update the code to test the property totalSavings
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property currencyCode (base name: "currencyCode")', function() {
      // uncomment below and update the code to test the property currencyCode
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property totalSavingsPercent (base name: "totalSavingsPercent")', function() {
      // uncomment below and update the code to test the property totalSavingsPercent
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property term (base name: "term")', function() {
      // uncomment below and update the code to test the property term
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property paymentOption (base name: "paymentOption")', function() {
      // uncomment below and update the code to test the property paymentOption
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property service (base name: "service")', function() {
      // uncomment below and update the code to test the property service
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property onDemandCount (base name: "onDemandCount")', function() {
      // uncomment below and update the code to test the property onDemandCount
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property onDemandCost (base name: "onDemandCost")', function() {
      // uncomment below and update the code to test the property onDemandCost
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property reservedCount (base name: "reservedCount")', function() {
      // uncomment below and update the code to test the property reservedCount
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property reservedCost (base name: "reservedCost")', function() {
      // uncomment below and update the code to test the property reservedCost
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property recommendedCount (base name: "recommendedCount")', function() {
      // uncomment below and update the code to test the property recommendedCount
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

    it('should have the property recommendedCost (base name: "recommendedCost")', function() {
      // uncomment below and update the code to test the property recommendedCost
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigServicesAzureVmsSummary();
      //expect(instance).to.be();
    });

  });

}));
