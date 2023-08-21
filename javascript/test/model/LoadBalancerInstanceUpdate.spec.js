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
    instance = new MorpheusApi.LoadBalancerInstanceUpdate();
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

  describe('LoadBalancerInstanceUpdate', function() {
    it('should create an instance of LoadBalancerInstanceUpdate', function() {
      // uncomment below and update the code to test LoadBalancerInstanceUpdate
      //var instane = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be.a(MorpheusApi.LoadBalancerInstanceUpdate);
    });

    it('should have the property vipName (base name: "vipName")', function() {
      // uncomment below and update the code to test the property vipName
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property vipAddress (base name: "vipAddress")', function() {
      // uncomment below and update the code to test the property vipAddress
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property vipPort (base name: "vipPort")', function() {
      // uncomment below and update the code to test the property vipPort
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property vipProtocol (base name: "vipProtocol")', function() {
      // uncomment below and update the code to test the property vipProtocol
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property vipHostname (base name: "vipHostname")', function() {
      // uncomment below and update the code to test the property vipHostname
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property pool (base name: "pool")', function() {
      // uncomment below and update the code to test the property pool
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property sslCert (base name: "sslCert")', function() {
      // uncomment below and update the code to test the property sslCert
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property sslServerCert (base name: "sslServerCert")', function() {
      // uncomment below and update the code to test the property sslServerCert
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.LoadBalancerInstanceUpdate();
      //expect(instance).to.be();
    });

  });

}));