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
    instance = new MorpheusApi.ApiReportsReport();
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

  describe('ApiReportsReport', function() {
    it('should create an instance of ApiReportsReport', function() {
      // uncomment below and update the code to test ApiReportsReport
      //var instane = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be.a(MorpheusApi.ApiReportsReport);
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

    it('should have the property startMonth (base name: "startMonth")', function() {
      // uncomment below and update the code to test the property startMonth
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

    it('should have the property endMonth (base name: "endMonth")', function() {
      // uncomment below and update the code to test the property endMonth
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

    it('should have the property groupId (base name: "groupId")', function() {
      // uncomment below and update the code to test the property groupId
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

    it('should have the property cloudId (base name: "cloudId")', function() {
      // uncomment below and update the code to test the property cloudId
      //var instance = new MorpheusApi.ApiReportsReport();
      //expect(instance).to.be();
    });

  });

}));