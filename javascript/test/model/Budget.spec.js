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
    instance = new MorpheusApi.Budget();
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

  describe('Budget', function() {
    it('should create an instance of Budget', function() {
      // uncomment below and update the code to test Budget
      //var instane = new MorpheusApi.Budget();
      //expect(instance).to.be.a(MorpheusApi.Budget);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property refScope (base name: "refScope")', function() {
      // uncomment below and update the code to test the property refScope
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property refName (base name: "refName")', function() {
      // uncomment below and update the code to test the property refName
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property period (base name: "period")', function() {
      // uncomment below and update the code to test the property period
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property year (base name: "year")', function() {
      // uncomment below and update the code to test the property year
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property resourceType (base name: "resourceType")', function() {
      // uncomment below and update the code to test the property resourceType
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property timezone (base name: "timezone")', function() {
      // uncomment below and update the code to test the property timezone
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property interval (base name: "interval")', function() {
      // uncomment below and update the code to test the property interval
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property costs (base name: "costs")', function() {
      // uncomment below and update the code to test the property costs
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property isFiscal (base name: "isFiscal")', function() {
      // uncomment below and update the code to test the property isFiscal
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property averageCost (base name: "averageCost")', function() {
      // uncomment below and update the code to test the property averageCost
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property totalCost (base name: "totalCost")', function() {
      // uncomment below and update the code to test the property totalCost
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property currency (base name: "currency")', function() {
      // uncomment below and update the code to test the property currency
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property rollover (base name: "rollover")', function() {
      // uncomment below and update the code to test the property rollover
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property warningLimit (base name: "warningLimit")', function() {
      // uncomment below and update the code to test the property warningLimit
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property overLimit (base name: "overLimit")', function() {
      // uncomment below and update the code to test the property overLimit
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property createdById (base name: "createdById")', function() {
      // uncomment below and update the code to test the property createdById
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property createdByName (base name: "createdByName")', function() {
      // uncomment below and update the code to test the property createdByName
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property updatedById (base name: "updatedById")', function() {
      // uncomment below and update the code to test the property updatedById
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property updatedByName (base name: "updatedByName")', function() {
      // uncomment below and update the code to test the property updatedByName
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

    it('should have the property stats (base name: "stats")', function() {
      // uncomment below and update the code to test the property stats
      //var instance = new MorpheusApi.Budget();
      //expect(instance).to.be();
    });

  });

}));
