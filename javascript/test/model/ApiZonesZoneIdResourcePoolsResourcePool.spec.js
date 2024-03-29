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
    instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
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

  describe('ApiZonesZoneIdResourcePoolsResourcePool', function() {
    it('should create an instance of ApiZonesZoneIdResourcePoolsResourcePool', function() {
      // uncomment below and update the code to test ApiZonesZoneIdResourcePoolsResourcePool
      //var instane = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be.a(MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property defaultPool (base name: "defaultPool")', function() {
      // uncomment below and update the code to test the property defaultPool
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property defaultImage (base name: "defaultImage")', function() {
      // uncomment below and update the code to test the property defaultImage
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property displayName (base name: "displayName")', function() {
      // uncomment below and update the code to test the property displayName
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property inventory (base name: "inventory")', function() {
      // uncomment below and update the code to test the property inventory
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property tenantPermissions (base name: "tenantPermissions")', function() {
      // uncomment below and update the code to test the property tenantPermissions
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

    it('should have the property resourcePermissions (base name: "resourcePermissions")', function() {
      // uncomment below and update the code to test the property resourcePermissions
      //var instance = new MorpheusApi.ApiZonesZoneIdResourcePoolsResourcePool();
      //expect(instance).to.be();
    });

  });

}));
