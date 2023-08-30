"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.instance_create_instance_instance_type import InstanceCreateInstanceInstanceType
from openapi_client.model.instance_create_instance_layout import InstanceCreateInstanceLayout
from openapi_client.model.instance_create_instance_plan import InstanceCreateInstancePlan
from openapi_client.model.instance_create_instance_site import InstanceCreateInstanceSite
globals()['InstanceCreateInstanceInstanceType'] = InstanceCreateInstanceInstanceType
globals()['InstanceCreateInstanceLayout'] = InstanceCreateInstanceLayout
globals()['InstanceCreateInstancePlan'] = InstanceCreateInstancePlan
globals()['InstanceCreateInstanceSite'] = InstanceCreateInstanceSite
from openapi_client.model.instance_create_instance import InstanceCreateInstance


class TestInstanceCreateInstance(unittest.TestCase):
    """InstanceCreateInstance unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testInstanceCreateInstance(self):
        """Test InstanceCreateInstance"""
        # FIXME: construct object with mandatory attributes with example values
        # model = InstanceCreateInstance()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()