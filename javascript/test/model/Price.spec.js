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
    instance = new MorpheusApi.Price();
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

  describe('Price', function() {
    it('should create an instance of Price', function() {
      // uncomment below and update the code to test Price
      //var instane = new MorpheusApi.Price();
      //expect(instance).to.be.a(MorpheusApi.Price);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property priceType (base name: "priceType")', function() {
      // uncomment below and update the code to test the property priceType
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property priceUnit (base name: "priceUnit")', function() {
      // uncomment below and update the code to test the property priceUnit
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property additionalPriceUnit (base name: "additionalPriceUnit")', function() {
      // uncomment below and update the code to test the property additionalPriceUnit
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property price (base name: "price")', function() {
      // uncomment below and update the code to test the property price
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property customPrice (base name: "customPrice")', function() {
      // uncomment below and update the code to test the property customPrice
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property markupType (base name: "markupType")', function() {
      // uncomment below and update the code to test the property markupType
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property markup (base name: "markup")', function() {
      // uncomment below and update the code to test the property markup
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property markupPercent (base name: "markupPercent")', function() {
      // uncomment below and update the code to test the property markupPercent
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property cost (base name: "cost")', function() {
      // uncomment below and update the code to test the property cost
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property currency (base name: "currency")', function() {
      // uncomment below and update the code to test the property currency
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property incurCharges (base name: "incurCharges")', function() {
      // uncomment below and update the code to test the property incurCharges
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property platform (base name: "platform")', function() {
      // uncomment below and update the code to test the property platform
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property software (base name: "software")', function() {
      // uncomment below and update the code to test the property software
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property volumeType (base name: "volumeType")', function() {
      // uncomment below and update the code to test the property volumeType
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property datastore (base name: "datastore")', function() {
      // uncomment below and update the code to test the property datastore
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property crossCloudApply (base name: "crossCloudApply")', function() {
      // uncomment below and update the code to test the property crossCloudApply
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.Price();
      //expect(instance).to.be();
    });

  });

}));
