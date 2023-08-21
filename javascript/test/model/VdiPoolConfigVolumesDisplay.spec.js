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
    instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
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

  describe('VdiPoolConfigVolumesDisplay', function() {
    it('should create an instance of VdiPoolConfigVolumesDisplay', function() {
      // uncomment below and update the code to test VdiPoolConfigVolumesDisplay
      //var instane = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be.a(MorpheusApi.VdiPoolConfigVolumesDisplay);
    });

    it('should have the property storage (base name: "storage")', function() {
      // uncomment below and update the code to test the property storage
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

    it('should have the property controller (base name: "controller")', function() {
      // uncomment below and update the code to test the property controller
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

    it('should have the property datastore (base name: "datastore")', function() {
      // uncomment below and update the code to test the property datastore
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

    it('should have the property displayOrder (base name: "displayOrder")', function() {
      // uncomment below and update the code to test the property displayOrder
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

    it('should have the property size (base name: "size")', function() {
      // uncomment below and update the code to test the property size
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

    it('should have the property mountPoint (base name: "mountPoint")', function() {
      // uncomment below and update the code to test the property mountPoint
      //var instance = new MorpheusApi.VdiPoolConfigVolumesDisplay();
      //expect(instance).to.be();
    });

  });

}));
