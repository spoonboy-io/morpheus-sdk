<?php
/**
 * PowerSchedule
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * PowerSchedule Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class PowerSchedule implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'powerSchedule';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'description' => 'string',
        'visibility' => 'string',
        'enabled' => 'bool',
        'schedule_type' => 'string',
        'schedule_timezone' => 'string',
        'monday_on' => 'int',
        'monday_on_time' => 'string',
        'monday_off' => 'int',
        'monday_off_time' => 'string',
        'tuesday_on' => 'int',
        'tuesday_on_time' => 'string',
        'tuesday_off' => 'int',
        'tuesday_off_time' => 'string',
        'wednesday_on' => 'int',
        'wednesday_on_time' => 'string',
        'wednesday_off' => 'int',
        'wednesday_off_time' => 'string',
        'thursday_on' => 'int',
        'thursday_on_time' => 'string',
        'thursday_off' => 'int',
        'thursday_off_time' => 'string',
        'friday_on' => 'int',
        'friday_on_time' => 'string',
        'friday_off' => 'int',
        'friday_off_time' => 'string',
        'saturday_on' => 'int',
        'saturday_on_time' => 'string',
        'saturday_off' => 'int',
        'saturday_off_time' => 'string',
        'sunday_on' => 'int',
        'sunday_on_time' => 'string',
        'sunday_off' => 'int',
        'sunday_off_time' => 'string',
        'total_monthly_hours_saved' => 'float',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'name' => null,
        'description' => null,
        'visibility' => null,
        'enabled' => null,
        'schedule_type' => null,
        'schedule_timezone' => null,
        'monday_on' => 'int64',
        'monday_on_time' => null,
        'monday_off' => 'int64',
        'monday_off_time' => null,
        'tuesday_on' => 'int64',
        'tuesday_on_time' => null,
        'tuesday_off' => 'int64',
        'tuesday_off_time' => null,
        'wednesday_on' => 'int64',
        'wednesday_on_time' => null,
        'wednesday_off' => 'int64',
        'wednesday_off_time' => null,
        'thursday_on' => 'int64',
        'thursday_on_time' => null,
        'thursday_off' => 'int64',
        'thursday_off_time' => null,
        'friday_on' => 'int64',
        'friday_on_time' => null,
        'friday_off' => 'int64',
        'friday_off_time' => null,
        'saturday_on' => 'int64',
        'saturday_on_time' => null,
        'saturday_off' => 'int64',
        'saturday_off_time' => null,
        'sunday_on' => 'int64',
        'sunday_on_time' => null,
        'sunday_off' => 'int64',
        'sunday_off_time' => null,
        'total_monthly_hours_saved' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time'
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'id' => 'id',
        'name' => 'name',
        'description' => 'description',
        'visibility' => 'visibility',
        'enabled' => 'enabled',
        'schedule_type' => 'scheduleType',
        'schedule_timezone' => 'scheduleTimezone',
        'monday_on' => 'mondayOn',
        'monday_on_time' => 'mondayOnTime',
        'monday_off' => 'mondayOff',
        'monday_off_time' => 'mondayOffTime',
        'tuesday_on' => 'tuesdayOn',
        'tuesday_on_time' => 'tuesdayOnTime',
        'tuesday_off' => 'tuesdayOff',
        'tuesday_off_time' => 'tuesdayOffTime',
        'wednesday_on' => 'wednesdayOn',
        'wednesday_on_time' => 'wednesdayOnTime',
        'wednesday_off' => 'wednesdayOff',
        'wednesday_off_time' => 'wednesdayOffTime',
        'thursday_on' => 'thursdayOn',
        'thursday_on_time' => 'thursdayOnTime',
        'thursday_off' => 'thursdayOff',
        'thursday_off_time' => 'thursdayOffTime',
        'friday_on' => 'fridayOn',
        'friday_on_time' => 'fridayOnTime',
        'friday_off' => 'fridayOff',
        'friday_off_time' => 'fridayOffTime',
        'saturday_on' => 'saturdayOn',
        'saturday_on_time' => 'saturdayOnTime',
        'saturday_off' => 'saturdayOff',
        'saturday_off_time' => 'saturdayOffTime',
        'sunday_on' => 'sundayOn',
        'sunday_on_time' => 'sundayOnTime',
        'sunday_off' => 'sundayOff',
        'sunday_off_time' => 'sundayOffTime',
        'total_monthly_hours_saved' => 'totalMonthlyHoursSaved',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'description' => 'setDescription',
        'visibility' => 'setVisibility',
        'enabled' => 'setEnabled',
        'schedule_type' => 'setScheduleType',
        'schedule_timezone' => 'setScheduleTimezone',
        'monday_on' => 'setMondayOn',
        'monday_on_time' => 'setMondayOnTime',
        'monday_off' => 'setMondayOff',
        'monday_off_time' => 'setMondayOffTime',
        'tuesday_on' => 'setTuesdayOn',
        'tuesday_on_time' => 'setTuesdayOnTime',
        'tuesday_off' => 'setTuesdayOff',
        'tuesday_off_time' => 'setTuesdayOffTime',
        'wednesday_on' => 'setWednesdayOn',
        'wednesday_on_time' => 'setWednesdayOnTime',
        'wednesday_off' => 'setWednesdayOff',
        'wednesday_off_time' => 'setWednesdayOffTime',
        'thursday_on' => 'setThursdayOn',
        'thursday_on_time' => 'setThursdayOnTime',
        'thursday_off' => 'setThursdayOff',
        'thursday_off_time' => 'setThursdayOffTime',
        'friday_on' => 'setFridayOn',
        'friday_on_time' => 'setFridayOnTime',
        'friday_off' => 'setFridayOff',
        'friday_off_time' => 'setFridayOffTime',
        'saturday_on' => 'setSaturdayOn',
        'saturday_on_time' => 'setSaturdayOnTime',
        'saturday_off' => 'setSaturdayOff',
        'saturday_off_time' => 'setSaturdayOffTime',
        'sunday_on' => 'setSundayOn',
        'sunday_on_time' => 'setSundayOnTime',
        'sunday_off' => 'setSundayOff',
        'sunday_off_time' => 'setSundayOffTime',
        'total_monthly_hours_saved' => 'setTotalMonthlyHoursSaved',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'description' => 'getDescription',
        'visibility' => 'getVisibility',
        'enabled' => 'getEnabled',
        'schedule_type' => 'getScheduleType',
        'schedule_timezone' => 'getScheduleTimezone',
        'monday_on' => 'getMondayOn',
        'monday_on_time' => 'getMondayOnTime',
        'monday_off' => 'getMondayOff',
        'monday_off_time' => 'getMondayOffTime',
        'tuesday_on' => 'getTuesdayOn',
        'tuesday_on_time' => 'getTuesdayOnTime',
        'tuesday_off' => 'getTuesdayOff',
        'tuesday_off_time' => 'getTuesdayOffTime',
        'wednesday_on' => 'getWednesdayOn',
        'wednesday_on_time' => 'getWednesdayOnTime',
        'wednesday_off' => 'getWednesdayOff',
        'wednesday_off_time' => 'getWednesdayOffTime',
        'thursday_on' => 'getThursdayOn',
        'thursday_on_time' => 'getThursdayOnTime',
        'thursday_off' => 'getThursdayOff',
        'thursday_off_time' => 'getThursdayOffTime',
        'friday_on' => 'getFridayOn',
        'friday_on_time' => 'getFridayOnTime',
        'friday_off' => 'getFridayOff',
        'friday_off_time' => 'getFridayOffTime',
        'saturday_on' => 'getSaturdayOn',
        'saturday_on_time' => 'getSaturdayOnTime',
        'saturday_off' => 'getSaturdayOff',
        'saturday_off_time' => 'getSaturdayOffTime',
        'sunday_on' => 'getSundayOn',
        'sunday_on_time' => 'getSundayOnTime',
        'sunday_off' => 'getSundayOff',
        'sunday_off_time' => 'getSundayOffTime',
        'total_monthly_hours_saved' => 'getTotalMonthlyHoursSaved',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['id'] = $data['id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['schedule_type'] = $data['schedule_type'] ?? null;
        $this->container['schedule_timezone'] = $data['schedule_timezone'] ?? null;
        $this->container['monday_on'] = $data['monday_on'] ?? null;
        $this->container['monday_on_time'] = $data['monday_on_time'] ?? null;
        $this->container['monday_off'] = $data['monday_off'] ?? null;
        $this->container['monday_off_time'] = $data['monday_off_time'] ?? null;
        $this->container['tuesday_on'] = $data['tuesday_on'] ?? null;
        $this->container['tuesday_on_time'] = $data['tuesday_on_time'] ?? null;
        $this->container['tuesday_off'] = $data['tuesday_off'] ?? null;
        $this->container['tuesday_off_time'] = $data['tuesday_off_time'] ?? null;
        $this->container['wednesday_on'] = $data['wednesday_on'] ?? null;
        $this->container['wednesday_on_time'] = $data['wednesday_on_time'] ?? null;
        $this->container['wednesday_off'] = $data['wednesday_off'] ?? null;
        $this->container['wednesday_off_time'] = $data['wednesday_off_time'] ?? null;
        $this->container['thursday_on'] = $data['thursday_on'] ?? null;
        $this->container['thursday_on_time'] = $data['thursday_on_time'] ?? null;
        $this->container['thursday_off'] = $data['thursday_off'] ?? null;
        $this->container['thursday_off_time'] = $data['thursday_off_time'] ?? null;
        $this->container['friday_on'] = $data['friday_on'] ?? null;
        $this->container['friday_on_time'] = $data['friday_on_time'] ?? null;
        $this->container['friday_off'] = $data['friday_off'] ?? null;
        $this->container['friday_off_time'] = $data['friday_off_time'] ?? null;
        $this->container['saturday_on'] = $data['saturday_on'] ?? null;
        $this->container['saturday_on_time'] = $data['saturday_on_time'] ?? null;
        $this->container['saturday_off'] = $data['saturday_off'] ?? null;
        $this->container['saturday_off_time'] = $data['saturday_off_time'] ?? null;
        $this->container['sunday_on'] = $data['sunday_on'] ?? null;
        $this->container['sunday_on_time'] = $data['sunday_on_time'] ?? null;
        $this->container['sunday_off'] = $data['sunday_off'] ?? null;
        $this->container['sunday_off_time'] = $data['sunday_off_time'] ?? null;
        $this->container['total_monthly_hours_saved'] = $data['total_monthly_hours_saved'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets schedule_type
     *
     * @return string|null
     */
    public function getScheduleType()
    {
        return $this->container['schedule_type'];
    }

    /**
     * Sets schedule_type
     *
     * @param string|null $schedule_type schedule_type
     *
     * @return self
     */
    public function setScheduleType($schedule_type)
    {
        $this->container['schedule_type'] = $schedule_type;

        return $this;
    }

    /**
     * Gets schedule_timezone
     *
     * @return string|null
     */
    public function getScheduleTimezone()
    {
        return $this->container['schedule_timezone'];
    }

    /**
     * Sets schedule_timezone
     *
     * @param string|null $schedule_timezone schedule_timezone
     *
     * @return self
     */
    public function setScheduleTimezone($schedule_timezone)
    {
        $this->container['schedule_timezone'] = $schedule_timezone;

        return $this;
    }

    /**
     * Gets monday_on
     *
     * @return int|null
     */
    public function getMondayOn()
    {
        return $this->container['monday_on'];
    }

    /**
     * Sets monday_on
     *
     * @param int|null $monday_on monday_on
     *
     * @return self
     */
    public function setMondayOn($monday_on)
    {
        $this->container['monday_on'] = $monday_on;

        return $this;
    }

    /**
     * Gets monday_on_time
     *
     * @return string|null
     */
    public function getMondayOnTime()
    {
        return $this->container['monday_on_time'];
    }

    /**
     * Sets monday_on_time
     *
     * @param string|null $monday_on_time monday_on_time
     *
     * @return self
     */
    public function setMondayOnTime($monday_on_time)
    {
        $this->container['monday_on_time'] = $monday_on_time;

        return $this;
    }

    /**
     * Gets monday_off
     *
     * @return int|null
     */
    public function getMondayOff()
    {
        return $this->container['monday_off'];
    }

    /**
     * Sets monday_off
     *
     * @param int|null $monday_off monday_off
     *
     * @return self
     */
    public function setMondayOff($monday_off)
    {
        $this->container['monday_off'] = $monday_off;

        return $this;
    }

    /**
     * Gets monday_off_time
     *
     * @return string|null
     */
    public function getMondayOffTime()
    {
        return $this->container['monday_off_time'];
    }

    /**
     * Sets monday_off_time
     *
     * @param string|null $monday_off_time monday_off_time
     *
     * @return self
     */
    public function setMondayOffTime($monday_off_time)
    {
        $this->container['monday_off_time'] = $monday_off_time;

        return $this;
    }

    /**
     * Gets tuesday_on
     *
     * @return int|null
     */
    public function getTuesdayOn()
    {
        return $this->container['tuesday_on'];
    }

    /**
     * Sets tuesday_on
     *
     * @param int|null $tuesday_on tuesday_on
     *
     * @return self
     */
    public function setTuesdayOn($tuesday_on)
    {
        $this->container['tuesday_on'] = $tuesday_on;

        return $this;
    }

    /**
     * Gets tuesday_on_time
     *
     * @return string|null
     */
    public function getTuesdayOnTime()
    {
        return $this->container['tuesday_on_time'];
    }

    /**
     * Sets tuesday_on_time
     *
     * @param string|null $tuesday_on_time tuesday_on_time
     *
     * @return self
     */
    public function setTuesdayOnTime($tuesday_on_time)
    {
        $this->container['tuesday_on_time'] = $tuesday_on_time;

        return $this;
    }

    /**
     * Gets tuesday_off
     *
     * @return int|null
     */
    public function getTuesdayOff()
    {
        return $this->container['tuesday_off'];
    }

    /**
     * Sets tuesday_off
     *
     * @param int|null $tuesday_off tuesday_off
     *
     * @return self
     */
    public function setTuesdayOff($tuesday_off)
    {
        $this->container['tuesday_off'] = $tuesday_off;

        return $this;
    }

    /**
     * Gets tuesday_off_time
     *
     * @return string|null
     */
    public function getTuesdayOffTime()
    {
        return $this->container['tuesday_off_time'];
    }

    /**
     * Sets tuesday_off_time
     *
     * @param string|null $tuesday_off_time tuesday_off_time
     *
     * @return self
     */
    public function setTuesdayOffTime($tuesday_off_time)
    {
        $this->container['tuesday_off_time'] = $tuesday_off_time;

        return $this;
    }

    /**
     * Gets wednesday_on
     *
     * @return int|null
     */
    public function getWednesdayOn()
    {
        return $this->container['wednesday_on'];
    }

    /**
     * Sets wednesday_on
     *
     * @param int|null $wednesday_on wednesday_on
     *
     * @return self
     */
    public function setWednesdayOn($wednesday_on)
    {
        $this->container['wednesday_on'] = $wednesday_on;

        return $this;
    }

    /**
     * Gets wednesday_on_time
     *
     * @return string|null
     */
    public function getWednesdayOnTime()
    {
        return $this->container['wednesday_on_time'];
    }

    /**
     * Sets wednesday_on_time
     *
     * @param string|null $wednesday_on_time wednesday_on_time
     *
     * @return self
     */
    public function setWednesdayOnTime($wednesday_on_time)
    {
        $this->container['wednesday_on_time'] = $wednesday_on_time;

        return $this;
    }

    /**
     * Gets wednesday_off
     *
     * @return int|null
     */
    public function getWednesdayOff()
    {
        return $this->container['wednesday_off'];
    }

    /**
     * Sets wednesday_off
     *
     * @param int|null $wednesday_off wednesday_off
     *
     * @return self
     */
    public function setWednesdayOff($wednesday_off)
    {
        $this->container['wednesday_off'] = $wednesday_off;

        return $this;
    }

    /**
     * Gets wednesday_off_time
     *
     * @return string|null
     */
    public function getWednesdayOffTime()
    {
        return $this->container['wednesday_off_time'];
    }

    /**
     * Sets wednesday_off_time
     *
     * @param string|null $wednesday_off_time wednesday_off_time
     *
     * @return self
     */
    public function setWednesdayOffTime($wednesday_off_time)
    {
        $this->container['wednesday_off_time'] = $wednesday_off_time;

        return $this;
    }

    /**
     * Gets thursday_on
     *
     * @return int|null
     */
    public function getThursdayOn()
    {
        return $this->container['thursday_on'];
    }

    /**
     * Sets thursday_on
     *
     * @param int|null $thursday_on thursday_on
     *
     * @return self
     */
    public function setThursdayOn($thursday_on)
    {
        $this->container['thursday_on'] = $thursday_on;

        return $this;
    }

    /**
     * Gets thursday_on_time
     *
     * @return string|null
     */
    public function getThursdayOnTime()
    {
        return $this->container['thursday_on_time'];
    }

    /**
     * Sets thursday_on_time
     *
     * @param string|null $thursday_on_time thursday_on_time
     *
     * @return self
     */
    public function setThursdayOnTime($thursday_on_time)
    {
        $this->container['thursday_on_time'] = $thursday_on_time;

        return $this;
    }

    /**
     * Gets thursday_off
     *
     * @return int|null
     */
    public function getThursdayOff()
    {
        return $this->container['thursday_off'];
    }

    /**
     * Sets thursday_off
     *
     * @param int|null $thursday_off thursday_off
     *
     * @return self
     */
    public function setThursdayOff($thursday_off)
    {
        $this->container['thursday_off'] = $thursday_off;

        return $this;
    }

    /**
     * Gets thursday_off_time
     *
     * @return string|null
     */
    public function getThursdayOffTime()
    {
        return $this->container['thursday_off_time'];
    }

    /**
     * Sets thursday_off_time
     *
     * @param string|null $thursday_off_time thursday_off_time
     *
     * @return self
     */
    public function setThursdayOffTime($thursday_off_time)
    {
        $this->container['thursday_off_time'] = $thursday_off_time;

        return $this;
    }

    /**
     * Gets friday_on
     *
     * @return int|null
     */
    public function getFridayOn()
    {
        return $this->container['friday_on'];
    }

    /**
     * Sets friday_on
     *
     * @param int|null $friday_on friday_on
     *
     * @return self
     */
    public function setFridayOn($friday_on)
    {
        $this->container['friday_on'] = $friday_on;

        return $this;
    }

    /**
     * Gets friday_on_time
     *
     * @return string|null
     */
    public function getFridayOnTime()
    {
        return $this->container['friday_on_time'];
    }

    /**
     * Sets friday_on_time
     *
     * @param string|null $friday_on_time friday_on_time
     *
     * @return self
     */
    public function setFridayOnTime($friday_on_time)
    {
        $this->container['friday_on_time'] = $friday_on_time;

        return $this;
    }

    /**
     * Gets friday_off
     *
     * @return int|null
     */
    public function getFridayOff()
    {
        return $this->container['friday_off'];
    }

    /**
     * Sets friday_off
     *
     * @param int|null $friday_off friday_off
     *
     * @return self
     */
    public function setFridayOff($friday_off)
    {
        $this->container['friday_off'] = $friday_off;

        return $this;
    }

    /**
     * Gets friday_off_time
     *
     * @return string|null
     */
    public function getFridayOffTime()
    {
        return $this->container['friday_off_time'];
    }

    /**
     * Sets friday_off_time
     *
     * @param string|null $friday_off_time friday_off_time
     *
     * @return self
     */
    public function setFridayOffTime($friday_off_time)
    {
        $this->container['friday_off_time'] = $friday_off_time;

        return $this;
    }

    /**
     * Gets saturday_on
     *
     * @return int|null
     */
    public function getSaturdayOn()
    {
        return $this->container['saturday_on'];
    }

    /**
     * Sets saturday_on
     *
     * @param int|null $saturday_on saturday_on
     *
     * @return self
     */
    public function setSaturdayOn($saturday_on)
    {
        $this->container['saturday_on'] = $saturday_on;

        return $this;
    }

    /**
     * Gets saturday_on_time
     *
     * @return string|null
     */
    public function getSaturdayOnTime()
    {
        return $this->container['saturday_on_time'];
    }

    /**
     * Sets saturday_on_time
     *
     * @param string|null $saturday_on_time saturday_on_time
     *
     * @return self
     */
    public function setSaturdayOnTime($saturday_on_time)
    {
        $this->container['saturday_on_time'] = $saturday_on_time;

        return $this;
    }

    /**
     * Gets saturday_off
     *
     * @return int|null
     */
    public function getSaturdayOff()
    {
        return $this->container['saturday_off'];
    }

    /**
     * Sets saturday_off
     *
     * @param int|null $saturday_off saturday_off
     *
     * @return self
     */
    public function setSaturdayOff($saturday_off)
    {
        $this->container['saturday_off'] = $saturday_off;

        return $this;
    }

    /**
     * Gets saturday_off_time
     *
     * @return string|null
     */
    public function getSaturdayOffTime()
    {
        return $this->container['saturday_off_time'];
    }

    /**
     * Sets saturday_off_time
     *
     * @param string|null $saturday_off_time saturday_off_time
     *
     * @return self
     */
    public function setSaturdayOffTime($saturday_off_time)
    {
        $this->container['saturday_off_time'] = $saturday_off_time;

        return $this;
    }

    /**
     * Gets sunday_on
     *
     * @return int|null
     */
    public function getSundayOn()
    {
        return $this->container['sunday_on'];
    }

    /**
     * Sets sunday_on
     *
     * @param int|null $sunday_on sunday_on
     *
     * @return self
     */
    public function setSundayOn($sunday_on)
    {
        $this->container['sunday_on'] = $sunday_on;

        return $this;
    }

    /**
     * Gets sunday_on_time
     *
     * @return string|null
     */
    public function getSundayOnTime()
    {
        return $this->container['sunday_on_time'];
    }

    /**
     * Sets sunday_on_time
     *
     * @param string|null $sunday_on_time sunday_on_time
     *
     * @return self
     */
    public function setSundayOnTime($sunday_on_time)
    {
        $this->container['sunday_on_time'] = $sunday_on_time;

        return $this;
    }

    /**
     * Gets sunday_off
     *
     * @return int|null
     */
    public function getSundayOff()
    {
        return $this->container['sunday_off'];
    }

    /**
     * Sets sunday_off
     *
     * @param int|null $sunday_off sunday_off
     *
     * @return self
     */
    public function setSundayOff($sunday_off)
    {
        $this->container['sunday_off'] = $sunday_off;

        return $this;
    }

    /**
     * Gets sunday_off_time
     *
     * @return string|null
     */
    public function getSundayOffTime()
    {
        return $this->container['sunday_off_time'];
    }

    /**
     * Sets sunday_off_time
     *
     * @param string|null $sunday_off_time sunday_off_time
     *
     * @return self
     */
    public function setSundayOffTime($sunday_off_time)
    {
        $this->container['sunday_off_time'] = $sunday_off_time;

        return $this;
    }

    /**
     * Gets total_monthly_hours_saved
     *
     * @return float|null
     */
    public function getTotalMonthlyHoursSaved()
    {
        return $this->container['total_monthly_hours_saved'];
    }

    /**
     * Sets total_monthly_hours_saved
     *
     * @param float|null $total_monthly_hours_saved total_monthly_hours_saved
     *
     * @return self
     */
    public function setTotalMonthlyHoursSaved($total_monthly_hours_saved)
    {
        $this->container['total_monthly_hours_saved'] = $total_monthly_hours_saved;

        return $this;
    }

    /**
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}

