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
    instance = new MorpheusApi.NetworkDomainCreate();
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

  describe('NetworkDomainCreate', function() {
    it('should create an instance of NetworkDomainCreate', function() {
      // uncomment below and update the code to test NetworkDomainCreate
      //var instane = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be.a(MorpheusApi.NetworkDomainCreate);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property displayName (base name: "displayName")', function() {
      // uncomment below and update the code to test the property displayName
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property publicZone (base name: "publicZone")', function() {
      // uncomment below and update the code to test the property publicZone
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property taskSetId (base name: "taskSetId")', function() {
      // uncomment below and update the code to test the property taskSetId
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property active (base name: "active")', function() {
      // uncomment below and update the code to test the property active
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property domainController (base name: "domainController")', function() {
      // uncomment below and update the code to test the property domainController
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property domainUsername (base name: "domainUsername")', function() {
      // uncomment below and update the code to test the property domainUsername
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property domainPassword (base name: "domainPassword")', function() {
      // uncomment below and update the code to test the property domainPassword
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property dcServer (base name: "dcServer")', function() {
      // uncomment below and update the code to test the property dcServer
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property ouPath (base name: "ouPath")', function() {
      // uncomment below and update the code to test the property ouPath
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property guestUsername (base name: "guestUsername")', function() {
      // uncomment below and update the code to test the property guestUsername
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

    it('should have the property guestPassword (base name: "guestPassword")', function() {
      // uncomment below and update the code to test the property guestPassword
      //var instance = new MorpheusApi.NetworkDomainCreate();
      //expect(instance).to.be();
    });

  });

}));
