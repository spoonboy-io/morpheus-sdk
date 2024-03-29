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
    instance = new MorpheusApi.User();
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

  describe('User', function() {
    it('should create an instance of User', function() {
      // uncomment below and update the code to test User
      //var instane = new MorpheusApi.User();
      //expect(instance).to.be.a(MorpheusApi.User);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property accountId (base name: "accountId")', function() {
      // uncomment below and update the code to test the property accountId
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property username (base name: "username")', function() {
      // uncomment below and update the code to test the property username
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property displayName (base name: "displayName")', function() {
      // uncomment below and update the code to test the property displayName
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property email (base name: "email")', function() {
      // uncomment below and update the code to test the property email
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property firstName (base name: "firstName")', function() {
      // uncomment below and update the code to test the property firstName
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property lastName (base name: "lastName")', function() {
      // uncomment below and update the code to test the property lastName
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property receiveNotifications (base name: "receiveNotifications")', function() {
      // uncomment below and update the code to test the property receiveNotifications
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property isUsing2FA (base name: "isUsing2FA")', function() {
      // uncomment below and update the code to test the property isUsing2FA
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property accountExpired (base name: "accountExpired")', function() {
      // uncomment below and update the code to test the property accountExpired
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property accountLocked (base name: "accountLocked")', function() {
      // uncomment below and update the code to test the property accountLocked
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property passwordExpired (base name: "passwordExpired")', function() {
      // uncomment below and update the code to test the property passwordExpired
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property loginCount (base name: "loginCount")', function() {
      // uncomment below and update the code to test the property loginCount
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property loginAttempts (base name: "loginAttempts")', function() {
      // uncomment below and update the code to test the property loginAttempts
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property lastLoginDate (base name: "lastLoginDate")', function() {
      // uncomment below and update the code to test the property lastLoginDate
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property roles (base name: "roles")', function() {
      // uncomment below and update the code to test the property roles
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property account (base name: "account")', function() {
      // uncomment below and update the code to test the property account
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property linuxUsername (base name: "linuxUsername")', function() {
      // uncomment below and update the code to test the property linuxUsername
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property linuxPassword (base name: "linuxPassword")', function() {
      // uncomment below and update the code to test the property linuxPassword
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property linuxKeyPairId (base name: "linuxKeyPairId")', function() {
      // uncomment below and update the code to test the property linuxKeyPairId
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property windowsUsername (base name: "windowsUsername")', function() {
      // uncomment below and update the code to test the property windowsUsername
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property windowsPassword (base name: "windowsPassword")', function() {
      // uncomment below and update the code to test the property windowsPassword
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property defaultPersona (base name: "defaultPersona")', function() {
      // uncomment below and update the code to test the property defaultPersona
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property dateCreated (base name: "dateCreated")', function() {
      // uncomment below and update the code to test the property dateCreated
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property lastUpdated (base name: "lastUpdated")', function() {
      // uncomment below and update the code to test the property lastUpdated
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

    it('should have the property access (base name: "access")', function() {
      // uncomment below and update the code to test the property access
      //var instance = new MorpheusApi.User();
      //expect(instance).to.be();
    });

  });

}));
