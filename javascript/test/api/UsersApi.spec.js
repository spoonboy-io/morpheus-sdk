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
    instance = new MorpheusApi.UsersApi();
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

  describe('UsersApi', function() {
    describe('addUser', function() {
      it('should call addUser successfully', function(done) {
        //uncomment below and update the code to test addUser
        //instance.addUser(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteUser', function() {
      it('should call deleteUser successfully', function(done) {
        //uncomment below and update the code to test deleteUser
        //instance.deleteUser(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteUserSettingsAccessToken', function() {
      it('should call deleteUserSettingsAccessToken successfully', function(done) {
        //uncomment below and update the code to test deleteUserSettingsAccessToken
        //instance.deleteUserSettingsAccessToken(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteUserSettingsAvatar', function() {
      it('should call deleteUserSettingsAvatar successfully', function(done) {
        //uncomment below and update the code to test deleteUserSettingsAvatar
        //instance.deleteUserSettingsAvatar(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteUserSettingsDesktopBackground', function() {
      it('should call deleteUserSettingsDesktopBackground successfully', function(done) {
        //uncomment below and update the code to test deleteUserSettingsDesktopBackground
        //instance.deleteUserSettingsDesktopBackground(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getUser', function() {
      it('should call getUser successfully', function(done) {
        //uncomment below and update the code to test getUser
        //instance.getUser(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getUserPermissions', function() {
      it('should call getUserPermissions successfully', function(done) {
        //uncomment below and update the code to test getUserPermissions
        //instance.getUserPermissions(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getUserSettingsApiClients', function() {
      it('should call getUserSettingsApiClients successfully', function(done) {
        //uncomment below and update the code to test getUserSettingsApiClients
        //instance.getUserSettingsApiClients(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listUserSettings', function() {
      it('should call listUserSettings successfully', function(done) {
        //uncomment below and update the code to test listUserSettings
        //instance.listUserSettings(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listUsers', function() {
      it('should call listUsers successfully', function(done) {
        //uncomment below and update the code to test listUsers
        //instance.listUsers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listUsersAvailableRoles', function() {
      it('should call listUsersAvailableRoles successfully', function(done) {
        //uncomment below and update the code to test listUsersAvailableRoles
        //instance.listUsersAvailableRoles(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateUserSettings', function() {
      it('should call updateUserSettings successfully', function(done) {
        //uncomment below and update the code to test updateUserSettings
        //instance.updateUserSettings(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateUserSettingsAccessToken', function() {
      it('should call updateUserSettingsAccessToken successfully', function(done) {
        //uncomment below and update the code to test updateUserSettingsAccessToken
        //instance.updateUserSettingsAccessToken(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateUserSettingsAvatar', function() {
      it('should call updateUserSettingsAvatar successfully', function(done) {
        //uncomment below and update the code to test updateUserSettingsAvatar
        //instance.updateUserSettingsAvatar(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateUserSettingsDesktopBackground', function() {
      it('should call updateUserSettingsDesktopBackground successfully', function(done) {
        //uncomment below and update the code to test updateUserSettingsDesktopBackground
        //instance.updateUserSettingsDesktopBackground(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateUsers', function() {
      it('should call updateUsers successfully', function(done) {
        //uncomment below and update the code to test updateUsers
        //instance.updateUsers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('whoami', function() {
      it('should call whoami successfully', function(done) {
        //uncomment below and update the code to test whoami
        //instance.whoami(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
