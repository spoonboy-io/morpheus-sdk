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
    instance = new MorpheusApi.InlineObject223();
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

  describe('InlineObject223', function() {
    it('should create an instance of InlineObject223', function() {
      // uncomment below and update the code to test InlineObject223
      //var instane = new MorpheusApi.InlineObject223();
      //expect(instance).to.be.a(MorpheusApi.InlineObject223);
    });

    it('should have the property server (base name: "server")', function() {
      // uncomment below and update the code to test the property server
      //var instance = new MorpheusApi.InlineObject223();
      //expect(instance).to.be();
    });

    it('should have the property installAgent (base name: "installAgent")', function() {
      // uncomment below and update the code to test the property installAgent
      //var instance = new MorpheusApi.InlineObject223();
      //expect(instance).to.be();
    });

    it('should have the property instanceTypeId (base name: "instanceTypeId")', function() {
      // uncomment below and update the code to test the property instanceTypeId
      //var instance = new MorpheusApi.InlineObject223();
      //expect(instance).to.be();
    });

    it('should have the property layout (base name: "layout")', function() {
      // uncomment below and update the code to test the property layout
      //var instance = new MorpheusApi.InlineObject223();
      //expect(instance).to.be();
    });

  });

}));
