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
    instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
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

  describe('ApiSecurityGroupsSecurityGroup', function() {
    it('should create an instance of ApiSecurityGroupsSecurityGroup', function() {
      // uncomment below and update the code to test ApiSecurityGroupsSecurityGroup
      //var instane = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be.a(MorpheusApi.ApiSecurityGroupsSecurityGroup);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

    it('should have the property zoneId (base name: "zoneId")', function() {
      // uncomment below and update the code to test the property zoneId
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

    it('should have the property customOptions (base name: "customOptions")', function() {
      // uncomment below and update the code to test the property customOptions
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

    it('should have the property tenantPermissions (base name: "tenantPermissions")', function() {
      // uncomment below and update the code to test the property tenantPermissions
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

    it('should have the property resourcePermissions (base name: "resourcePermissions")', function() {
      // uncomment below and update the code to test the property resourcePermissions
      //var instance = new MorpheusApi.ApiSecurityGroupsSecurityGroup();
      //expect(instance).to.be();
    });

  });

}));
