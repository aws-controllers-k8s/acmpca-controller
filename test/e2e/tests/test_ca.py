# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the ACMPCA CertificateAuthority resource.
"""

import time
import logging
import base64
import pytest
import re

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from acktest.k8s.resource import _get_k8s_api_client
from kubernetes import client
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_acmpca_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.helper import ACMPCAValidator
from e2e.fixtures import k8s_secret
from acktest import tags

RESOURCE_PLURAL = "certificateauthorities"

CREATE_WAIT_AFTER_SECONDS = 30
UPDATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

@pytest.fixture(scope="module")
def simple_certificate_authority(acmpca_client):
    ca_name = random_suffix_name("certificate-authority", 50)
    replacements = {}
    suffix = random_suffix_name("", 10)
    replacements["NAME"] = ca_name
    replacements["COMMON_NAME"] = "www.example" + suffix + ".com"
    replacements["COUNTRY"] = "US"
    replacements["LOCALITY"] = "Arlington"
    replacements["ORG"] = "Example Organization " + suffix
    replacements["STATE"] = "Virginia"

    # Load CA CR
    ca_resource_data = load_acmpca_resource(
        "certificate_authority_defaults",
        additional_replacements=replacements,
    )

    # Create k8s resource
    ca_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        ca_name, namespace="default",
    )
    k8s.create_custom_resource(ca_ref, ca_resource_data)
    ca_cr = k8s.wait_resource_consumed_by_controller(ca_ref)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)

    assert ca_cr is not None
    assert k8s.get_resource_exists(ca_ref)
    logging.info(ca_cr)

    ca_resource_arn =  k8s.get_resource_arn(ca_cr)
    assert ca_resource_arn is not None

    yield (ca_cr, ca_ref, ca_resource_arn)

    #Delete CA k8s resource
    _, deleted = k8s.delete_custom_resource(ca_ref)
    assert deleted is True

    time.sleep(DELETE_WAIT_AFTER_SECONDS) 

    # Check CA status is DELETED
    acmpca_validator = ACMPCAValidator(acmpca_client)
    acmpca_validator.assert_certificate_authority(ca_resource_arn, "DELETED")

@service_marker
class TestCertificateAuthority:

    def test_ca_crud(self, acmpca_client, simple_certificate_authority):
        
        (ca_cr, ca_ref, ca_resource_arn) = simple_certificate_authority

        # Check CA status is PENDING_CERTIFICATE
        acmpca_validator = ACMPCAValidator(acmpca_client)
        ca = acmpca_validator.assert_certificate_authority(ca_resource_arn, "PENDING_CERTIFICATE")

        # Check CA Spec fields
        assert ca["Type"] == "ROOT"
        assert re.search("^www[.]example.{10}[.]com$", ca["CertificateAuthorityConfiguration"]["Subject"]["CommonName"])
        assert ca["CertificateAuthorityConfiguration"]["Subject"]["Country"] == "US"
        assert ca["CertificateAuthorityConfiguration"]["Subject"]["Locality"] == "Arlington"
        assert re.search("^Example Organization .{10}$", ca["CertificateAuthorityConfiguration"]["Subject"]["Organization"])
        assert ca["CertificateAuthorityConfiguration"]["Subject"]["State"] == "Virginia"
        assert ca["CertificateAuthorityConfiguration"]["KeyAlgorithm"] == "RSA_2048"
        assert ca["CertificateAuthorityConfiguration"]["SigningAlgorithm"] == "SHA256WITHRSA"

        # Check Tags
        expected_tags = [
            {
                "key": "tag1",
                "value": "val1"
            },
        ]
        observed_tags = acmpca_validator.get_ca_tags(ca_resource_arn)
        tags_dict = tags.to_dict(
            expected_tags,
            key_member_name="key",
            value_member_name="value"
        )
        tags.assert_equal_without_ack_tags(
            expected=tags_dict,
            actual=observed_tags,
        )

        ca_cr = k8s.patch_custom_resource(ca_ref, {})
        logging.info(ca_cr)

        # Check CA Status fields
        assert 'status' in ca_cr
        assert 'certificateSigningRequest' in ca_cr['status']
        csr = acmpca_validator.get_csr(ca_resource_arn)
        assert ca_cr['status']['certificateSigningRequest'] == csr

        assert 'status' in ca_cr['status']
        assert ca_cr['status']['status'] == "PENDING_CERTIFICATE"

        # Update CA tags
        new_tags = [
            {
                "key": "tag2",
                "value": "val2"
            },
        ]

        updates = {
            "spec": {
                "tags": new_tags
            },
        }
        patch_res = k8s.patch_custom_resource(ca_ref, updates)
        logging.info(patch_res)
        time.sleep(UPDATE_WAIT_AFTER_SECONDS) 

        # Check Tags
        observed_tags = acmpca_validator.get_ca_tags(ca_resource_arn)
        tags_dict = tags.to_dict(
            new_tags,
            key_member_name="key",
            value_member_name="value"
        )
        logging.info(observed_tags)
        logging.info(tags_dict)
        tags.assert_equal_without_ack_tags(
            expected=tags_dict,
            actual=observed_tags,
        )