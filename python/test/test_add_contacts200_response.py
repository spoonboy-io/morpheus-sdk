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
from openapi_client.model.add_contacts200_response_all_of import AddContacts200ResponseAllOf
from openapi_client.model.contact import Contact
from openapi_client.model.model200_success import Model200Success
globals()['AddContacts200ResponseAllOf'] = AddContacts200ResponseAllOf
globals()['Contact'] = Contact
globals()['Model200Success'] = Model200Success
from openapi_client.model.add_contacts200_response import AddContacts200Response


class TestAddContacts200Response(unittest.TestCase):
    """AddContacts200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAddContacts200Response(self):
        """Test AddContacts200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AddContacts200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
