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
    instance = new MorpheusApi.InlineResponse200143Certificates();
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

  describe('InlineResponse200143Certificates', function() {
    it('should create an instance of InlineResponse200143Certificates', function() {
      // uncomment below and update the code to test InlineResponse200143Certificates
      //var instane = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be.a(MorpheusApi.InlineResponse200143Certificates);
    });

    it('should have the property accountId (base name: "accountId")', function() {
      // uncomment below and update the code to test the property accountId
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property certFile (base name: "certFile")', function() {
      // uncomment below and update the code to test the property certFile
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property domainName (base name: "domainName")', function() {
      // uncomment below and update the code to test the property domainName
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property generated (base name: "generated")', function() {
      // uncomment below and update the code to test the property generated
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property keyFile (base name: "keyFile")', function() {
      // uncomment below and update the code to test the property keyFile
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

    it('should have the property wildcard (base name: "wildcard")', function() {
      // uncomment below and update the code to test the property wildcard
      //var instance = new MorpheusApi.InlineResponse200143Certificates();
      //expect(instance).to.be();
    });

  });

}));
