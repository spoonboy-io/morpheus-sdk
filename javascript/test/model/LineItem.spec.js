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
    instance = new MorpheusApi.LineItem();
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

  describe('LineItem', function() {
    it('should create an instance of LineItem', function() {
      // uncomment below and update the code to test LineItem
      //var instane = new MorpheusApi.LineItem();
      //expect(instance).to.be.a(MorpheusApi.LineItem);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property invoiceId (base name: "invoiceId")', function() {
      // uncomment below and update the code to test the property invoiceId
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property refName (base name: "refName")', function() {
      // uncomment below and update the code to test the property refName
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemId (base name: "itemId")', function() {
      // uncomment below and update the code to test the property itemId
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemType (base name: "itemType")', function() {
      // uncomment below and update the code to test the property itemType
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemName (base name: "itemName")', function() {
      // uncomment below and update the code to test the property itemName
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemDescription (base name: "itemDescription")', function() {
      // uncomment below and update the code to test the property itemDescription
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property productId (base name: "productId")', function() {
      // uncomment below and update the code to test the property productId
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property productCode (base name: "productCode")', function() {
      // uncomment below and update the code to test the property productCode
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property productName (base name: "productName")', function() {
      // uncomment below and update the code to test the property productName
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemSeller (base name: "itemSeller")', function() {
      // uncomment below and update the code to test the property itemSeller
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemAction (base name: "itemAction")', function() {
      // uncomment below and update the code to test the property itemAction
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property rateId (base name: "rateId")', function() {
      // uncomment below and update the code to test the property rateId
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property rateClass (base name: "rateClass")', function() {
      // uncomment below and update the code to test the property rateClass
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property rateUnit (base name: "rateUnit")', function() {
      // uncomment below and update the code to test the property rateUnit
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property rateTerm (base name: "rateTerm")', function() {
      // uncomment below and update the code to test the property rateTerm
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property usageType (base name: "usageType")', function() {
      // uncomment below and update the code to test the property usageType
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property usageCategory (base name: "usageCategory")', function() {
      // uncomment below and update the code to test the property usageCategory
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property usageService (base name: "usageService")', function() {
      // uncomment below and update the code to test the property usageService
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemUsage (base name: "itemUsage")', function() {
      // uncomment below and update the code to test the property itemUsage
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemRate (base name: "itemRate")', function() {
      // uncomment below and update the code to test the property itemRate
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemCost (base name: "itemCost")', function() {
      // uncomment below and update the code to test the property itemCost
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemPriceRate (base name: "itemPriceRate")', function() {
      // uncomment below and update the code to test the property itemPriceRate
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemPrice (base name: "itemPrice")', function() {
      // uncomment below and update the code to test the property itemPrice
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemTax (base name: "itemTax")', function() {
      // uncomment below and update the code to test the property itemTax
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property itemTerm (base name: "itemTerm")', function() {
      // uncomment below and update the code to test the property itemTerm
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property taxType (base name: "taxType")', function() {
      // uncomment below and update the code to test the property taxType
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property regionCode (base name: "regionCode")', function() {
      // uncomment below and update the code to test the property regionCode
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property currency (base name: "currency")', function() {
      // uncomment below and update the code to test the property currency
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property conversionRate (base name: "conversionRate")', function() {
      // uncomment below and update the code to test the property conversionRate
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.LineItem();
      //expect(instance).to.be();
    });

  });

}));
