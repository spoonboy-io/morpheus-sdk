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
from openapi_client.model.add_tenant200_response_all_of import AddTenant200ResponseAllOf
from openapi_client.model.model200_success import Model200Success
from openapi_client.model.tenant import Tenant
globals()['AddTenant200ResponseAllOf'] = AddTenant200ResponseAllOf
globals()['Model200Success'] = Model200Success
globals()['Tenant'] = Tenant
from openapi_client.model.add_tenant200_response import AddTenant200Response


class TestAddTenant200Response(unittest.TestCase):
    """AddTenant200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAddTenant200Response(self):
        """Test AddTenant200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AddTenant200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()