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
    instance = new MorpheusApi.Log();
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

  describe('Log', function() {
    it('should create an instance of Log', function() {
      // uncomment below and update the code to test Log
      //var instane = new MorpheusApi.Log();
      //expect(instance).to.be.a(MorpheusApi.Log);
    });

    it('should have the property sort (base name: "sort")', function() {
      // uncomment below and update the code to test the property sort
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property offset (base name: "offset")', function() {
      // uncomment below and update the code to test the property offset
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property start (base name: "start")', function() {
      // uncomment below and update the code to test the property start
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property end (base name: "end")', function() {
      // uncomment below and update the code to test the property end
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property data (base name: "data")', function() {
      // uncomment below and update the code to test the property data
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property max (base name: "max")', function() {
      // uncomment below and update the code to test the property max
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property grandTotal (base name: "grandTotal")', function() {
      // uncomment below and update the code to test the property grandTotal
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property total (base name: "total")', function() {
      // uncomment below and update the code to test the property total
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property success (base name: "success")', function() {
      // uncomment below and update the code to test the property success
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

    it('should have the property count (base name: "count")', function() {
      // uncomment below and update the code to test the property count
      //var instance = new MorpheusApi.Log();
      //expect(instance).to.be();
    });

  });

}));
