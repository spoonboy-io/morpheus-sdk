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
    instance = new MorpheusApi.UserSourceCreateJumpCloud();
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

  describe('UserSourceCreateJumpCloud', function() {
    it('should create an instance of UserSourceCreateJumpCloud', function() {
      // uncomment below and update the code to test UserSourceCreateJumpCloud
      //var instane = new MorpheusApi.UserSourceCreateJumpCloud();
      //expect(instance).to.be.a(MorpheusApi.UserSourceCreateJumpCloud);
    });

    it('should have the property organizationId (base name: "organizationId")', function() {
      // uncomment below and update the code to test the property organizationId
      //var instance = new MorpheusApi.UserSourceCreateJumpCloud();
      //expect(instance).to.be();
    });

    it('should have the property bindingUsername (base name: "bindingUsername")', function() {
      // uncomment below and update the code to test the property bindingUsername
      //var instance = new MorpheusApi.UserSourceCreateJumpCloud();
      //expect(instance).to.be();
    });

    it('should have the property bindingPassword (base name: "bindingPassword")', function() {
      // uncomment below and update the code to test the property bindingPassword
      //var instance = new MorpheusApi.UserSourceCreateJumpCloud();
      //expect(instance).to.be();
    });

    it('should have the property requiredRole (base name: "requiredRole")', function() {
      // uncomment below and update the code to test the property requiredRole
      //var instance = new MorpheusApi.UserSourceCreateJumpCloud();
      //expect(instance).to.be();
    });

  });

}));
