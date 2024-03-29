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
    instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
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

  describe('GuidanceAzureReservationsConfigDetailList', function() {
    it('should create an instance of GuidanceAzureReservationsConfigDetailList', function() {
      // uncomment below and update the code to test GuidanceAzureReservationsConfigDetailList
      //var instane = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be.a(MorpheusApi.GuidanceAzureReservationsConfigDetailList);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property apiName (base name: "apiName")', function() {
      // uncomment below and update the code to test the property apiName
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property apiType (base name: "apiType")', function() {
      // uncomment below and update the code to test the property apiType
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property period (base name: "period")', function() {
      // uncomment below and update the code to test the property period
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property size (base name: "size")', function() {
      // uncomment below and update the code to test the property size
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property region (base name: "region")', function() {
      // uncomment below and update the code to test the property region
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property term (base name: "term")', function() {
      // uncomment below and update the code to test the property term
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property meterId (base name: "meterId")', function() {
      // uncomment below and update the code to test the property meterId
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property onDemandCount (base name: "onDemandCount")', function() {
      // uncomment below and update the code to test the property onDemandCount
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property onDemandCost (base name: "onDemandCost")', function() {
      // uncomment below and update the code to test the property onDemandCost
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property reservedCount (base name: "reservedCount")', function() {
      // uncomment below and update the code to test the property reservedCount
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property reservedCost (base name: "reservedCost")', function() {
      // uncomment below and update the code to test the property reservedCost
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property recommendedCount (base name: "recommendedCount")', function() {
      // uncomment below and update the code to test the property recommendedCount
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property recommendedCost (base name: "recommendedCost")', function() {
      // uncomment below and update the code to test the property recommendedCost
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property totalSavings (base name: "totalSavings")', function() {
      // uncomment below and update the code to test the property totalSavings
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

    it('should have the property totalSavingsPercent (base name: "totalSavingsPercent")', function() {
      // uncomment below and update the code to test the property totalSavingsPercent
      //var instance = new MorpheusApi.GuidanceAzureReservationsConfigDetailList();
      //expect(instance).to.be();
    });

  });

}));
