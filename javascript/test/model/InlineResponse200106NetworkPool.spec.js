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
    instance = new MorpheusApi.InlineResponse200106NetworkPool();
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

  describe('InlineResponse200106NetworkPool', function() {
    it('should create an instance of InlineResponse200106NetworkPool', function() {
      // uncomment below and update the code to test InlineResponse200106NetworkPool
      //var instane = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be.a(MorpheusApi.InlineResponse200106NetworkPool);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property displayName (base name: "displayName")', function() {
      // uncomment below and update the code to test the property displayName
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property internalId (base name: "internalId")', function() {
      // uncomment below and update the code to test the property internalId
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property externalId (base name: "externalId")', function() {
      // uncomment below and update the code to test the property externalId
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property dnsDomain (base name: "dnsDomain")', function() {
      // uncomment below and update the code to test the property dnsDomain
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property dnsSearchPath (base name: "dnsSearchPath")', function() {
      // uncomment below and update the code to test the property dnsSearchPath
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property hostPrefix (base name: "hostPrefix")', function() {
      // uncomment below and update the code to test the property hostPrefix
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property httpProxy (base name: "httpProxy")', function() {
      // uncomment below and update the code to test the property httpProxy
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property dnsServers (base name: "dnsServers")', function() {
      // uncomment below and update the code to test the property dnsServers
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property dnsSuffixList (base name: "dnsSuffixList")', function() {
      // uncomment below and update the code to test the property dnsSuffixList
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property dhcpServer (base name: "dhcpServer")', function() {
      // uncomment below and update the code to test the property dhcpServer
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property dhcpIp (base name: "dhcpIp")', function() {
      // uncomment below and update the code to test the property dhcpIp
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property gateway (base name: "gateway")', function() {
      // uncomment below and update the code to test the property gateway
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property netmask (base name: "netmask")', function() {
      // uncomment below and update the code to test the property netmask
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property subnetAddress (base name: "subnetAddress")', function() {
      // uncomment below and update the code to test the property subnetAddress
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property ipCount (base name: "ipCount")', function() {
      // uncomment below and update the code to test the property ipCount
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property freeCount (base name: "freeCount")', function() {
      // uncomment below and update the code to test the property freeCount
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property poolEnabled (base name: "poolEnabled")', function() {
      // uncomment below and update the code to test the property poolEnabled
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property tftpServer (base name: "tftpServer")', function() {
      // uncomment below and update the code to test the property tftpServer
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property bootFile (base name: "bootFile")', function() {
      // uncomment below and update the code to test the property bootFile
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property refType (base name: "refType")', function() {
      // uncomment below and update the code to test the property refType
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property refId (base name: "refId")', function() {
      // uncomment below and update the code to test the property refId
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property parentType (base name: "parentType")', function() {
      // uncomment below and update the code to test the property parentType
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property parentId (base name: "parentId")', function() {
      // uncomment below and update the code to test the property parentId
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property poolGroup (base name: "poolGroup")', function() {
      // uncomment below and update the code to test the property poolGroup
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

    it('should have the property ipRanges (base name: "ipRanges")', function() {
      // uncomment below and update the code to test the property ipRanges
      //var instance = new MorpheusApi.InlineResponse200106NetworkPool();
      //expect(instance).to.be();
    });

  });

}));
