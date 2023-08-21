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
    instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
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

  describe('ApiPowerSchedulesIdSchedule', function() {
    it('should create an instance of ApiPowerSchedulesIdSchedule', function() {
      // uncomment below and update the code to test ApiPowerSchedulesIdSchedule
      //var instane = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be.a(MorpheusApi.ApiPowerSchedulesIdSchedule);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property scheduleType (base name: "scheduleType")', function() {
      // uncomment below and update the code to test the property scheduleType
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property scheduleTimezone (base name: "scheduleTimezone")', function() {
      // uncomment below and update the code to test the property scheduleTimezone
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property mondayOnTime (base name: "mondayOnTime")', function() {
      // uncomment below and update the code to test the property mondayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property mondayOffTime (base name: "mondayOffTime")', function() {
      // uncomment below and update the code to test the property mondayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property tuesdayOnTime (base name: "tuesdayOnTime")', function() {
      // uncomment below and update the code to test the property tuesdayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property tuesdayOffTime (base name: "tuesdayOffTime")', function() {
      // uncomment below and update the code to test the property tuesdayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property wednesdayOnTime (base name: "wednesdayOnTime")', function() {
      // uncomment below and update the code to test the property wednesdayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property wednesdayOffTime (base name: "wednesdayOffTime")', function() {
      // uncomment below and update the code to test the property wednesdayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property thursdayOnTime (base name: "thursdayOnTime")', function() {
      // uncomment below and update the code to test the property thursdayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property thursdayOffTime (base name: "thursdayOffTime")', function() {
      // uncomment below and update the code to test the property thursdayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property fridayOnTime (base name: "fridayOnTime")', function() {
      // uncomment below and update the code to test the property fridayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property fridayOffTime (base name: "fridayOffTime")', function() {
      // uncomment below and update the code to test the property fridayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property saturdayOnTime (base name: "saturdayOnTime")', function() {
      // uncomment below and update the code to test the property saturdayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property saturdayOffTime (base name: "saturdayOffTime")', function() {
      // uncomment below and update the code to test the property saturdayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property sundayOnTime (base name: "sundayOnTime")', function() {
      // uncomment below and update the code to test the property sundayOnTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

    it('should have the property sundayOffTime (base name: "sundayOffTime")', function() {
      // uncomment below and update the code to test the property sundayOffTime
      //var instance = new MorpheusApi.ApiPowerSchedulesIdSchedule();
      //expect(instance).to.be();
    });

  });

}));
