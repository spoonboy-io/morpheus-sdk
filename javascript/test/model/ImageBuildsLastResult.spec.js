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
    instance = new MorpheusApi.ImageBuildsLastResult();
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

  describe('ImageBuildsLastResult', function() {
    it('should create an instance of ImageBuildsLastResult', function() {
      // uncomment below and update the code to test ImageBuildsLastResult
      //var instane = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be.a(MorpheusApi.ImageBuildsLastResult);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property imageBuild (base name: "imageBuild")', function() {
      // uncomment below and update the code to test the property imageBuild
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property buildNumber (base name: "buildNumber")', function() {
      // uncomment below and update the code to test the property buildNumber
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property startDate (base name: "startDate")', function() {
      // uncomment below and update the code to test the property startDate
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property endDate (base name: "endDate")', function() {
      // uncomment below and update the code to test the property endDate
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property statusMessage (base name: "statusMessage")', function() {
      // uncomment below and update the code to test the property statusMessage
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property statusPercent (base name: "statusPercent")', function() {
      // uncomment below and update the code to test the property statusPercent
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property statusEta (base name: "statusEta")', function() {
      // uncomment below and update the code to test the property statusEta
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property errorMessage (base name: "errorMessage")', function() {
      // uncomment below and update the code to test the property errorMessage
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "createdBy")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property tempInstance (base name: "tempInstance")', function() {
      // uncomment below and update the code to test the property tempInstance
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

    it('should have the property virtualImages (base name: "virtualImages")', function() {
      // uncomment below and update the code to test the property virtualImages
      //var instance = new MorpheusApi.ImageBuildsLastResult();
      //expect(instance).to.be();
    });

  });

}));
