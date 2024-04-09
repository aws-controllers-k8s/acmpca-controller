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

"""Integration tests for the ACMPCA Certificate resource.
"""

import time
import logging
import base64
import pytest

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from acktest.k8s.resource import _get_k8s_api_client
from kubernetes import client
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_acmpca_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.helper import ACMPCAValidator
from e2e.fixtures import k8s_secret

RESOURCE_PLURAL = "certificates"

CREATE_WAIT_AFTER_SECONDS = 10
UPDATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

@pytest.fixture(scope="module")
def create_secret(k8s_secret):
    secret = k8s_secret(
        "default",
        random_suffix_name("certificate-secret", 50),
        "certificate",
        "value"
    )
    yield secret

@pytest.fixture(scope="module")
def simple_certificate_authority():
    ca_name = random_suffix_name("certificate-authority", 50)
    replacements = {}
    suffix = random_suffix_name("", 2)
    replacements["NAME"] = ca_name
    replacements["COMMON_NAME"] = "www.example" + suffix + ".com"
    replacements["COUNTRY"] = "US"
    replacements["LOCALITY"] = "Arlington"
    replacements["ORG"] = "Example Organization " + suffix
    replacements["STATE"] = "Virginia"

    # Load CA CR
    ca_resource_data = load_acmpca_resource(
        "certificate_authority",
        additional_replacements=replacements,
    )

    # Create k8s resource
    ca_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificateauthorities",
        ca_name, namespace="default",
    )
    k8s.create_custom_resource(ca_ref, ca_resource_data)
    ca_cr = k8s.wait_resource_consumed_by_controller(ca_ref)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)

    assert ca_cr is not None
    assert k8s.get_resource_exists(ca_ref)
    logging.info(ca_cr)
    logging.info(ca_ref)

    ca_resource_arn =  k8s.get_resource_arn(ca_cr)
    assert ca_resource_arn is not None

    yield (ca_cr, ca_name)

    #Delete CA k8s resource
    _, deleted = k8s.delete_custom_resource(ca_ref)
    assert deleted is True

@pytest.fixture(scope="module")
def simple_root_certificate(acmpca_client, create_secret, simple_certificate_authority):
    (ca_cr, ca_name) = simple_certificate_authority
    ca_arn = ca_cr['status']['ackResourceMetadata']['arn']

    cert_name = random_suffix_name("certificate", 30)

    secret = create_secret
    logging.info(secret)
    
    replacements = {}
    replacements["NAME"] = cert_name
    replacements["CA_ARN"] = ca_arn
    replacements["CA_NAME"] = ca_name
    replacements["CERTIFICATE_SEC_NS"] = secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = secret.name
    replacements["CERTIFICATE_SEC_KEY"] = secret.key
    replacements["TEMPLATE_ARN"] = "arn:aws:acm-pca:::template/RootCACertificate/V1"

    # Load Certificate CR
    resource_data = load_acmpca_resource(
        "certificate",
        additional_replacements=replacements,
    )

    # Create k8s resource
    ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        cert_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)

    # Check CA status is PENDING_CERTIFICATE
    acmpca_validator = ACMPCAValidator(acmpca_client)
    acmpca_validator.assert_certificate_authority(ca_arn, "PENDING_CERTIFICATE")

    assert cr is not None
    assert k8s.get_resource_exists(ref)
    logging.info(cr)

    resource_arn =  k8s.get_resource_arn(cr)
    assert resource_arn is not None

    yield (ca_arn, resource_arn, cr, secret)

    #Delete Certificate k8s resource
    _, deleted = k8s.delete_custom_resource(ref)
    assert deleted is True

@pytest.fixture(scope="module")
def simple_root_certificate_with_ref(acmpca_client, create_secret, simple_certificate_authority):
    (ca_cr, ca_name) = simple_certificate_authority
    ca_arn = ca_cr['status']['ackResourceMetadata']['arn']

    cert_name = random_suffix_name("certificate", 30)

    secret = create_secret
    logging.info(secret)
    
    replacements = {}
    replacements["NAME"] = cert_name
    replacements["CA_NAME"] = ca_name
    replacements["CERTIFICATE_SEC_NS"] = secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = secret.name
    replacements["CERTIFICATE_SEC_KEY"] = secret.key
    replacements["TEMPLATE_ARN"] = "arn:aws:acm-pca:::template/RootCACertificate/V1"

    # Load Certificate CR
    resource_data = load_acmpca_resource(
        "certificate_ref",
        additional_replacements=replacements,
    )

    # Create k8s resource
    ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        cert_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)

    # Check CA status is PENDING_CERTIFICATE
    acmpca_validator = ACMPCAValidator(acmpca_client)
    acmpca_validator.assert_certificate_authority(ca_arn, "PENDING_CERTIFICATE")

    assert cr is not None
    assert k8s.get_resource_exists(ref)
    logging.info(cr)

    resource_arn =  k8s.get_resource_arn(cr)
    assert resource_arn is not None

    yield (ca_arn, resource_arn, secret)

    #Delete Certificate k8s resource
    _, deleted = k8s.delete_custom_resource(ref)
    assert deleted is True

@pytest.fixture(scope="module")
def simple_root_certificate_without_secret_key(acmpca_client, create_secret, simple_certificate_authority):
    (ca_cr, ca_name) = simple_certificate_authority
    ca_arn = ca_cr['status']['ackResourceMetadata']['arn']

    cert_name = random_suffix_name("certificate", 30)

    secret = create_secret
    logging.info(secret)
    
    replacements = {}
    replacements["NAME"] = cert_name
    replacements["CA_NAME"] = ca_name
    replacements["CERTIFICATE_SEC_NS"] = secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = secret.name
    replacements["TEMPLATE_ARN"] = "arn:aws:acm-pca:::template/RootCACertificate/V1"

    # Load Certificate CR
    resource_data = load_acmpca_resource(
        "certificate_without_secret_key",
        additional_replacements=replacements,
    )

    # Create k8s resource
    ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        cert_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)

    # Check CA status is PENDING_CERTIFICATE
    acmpca_validator = ACMPCAValidator(acmpca_client)
    acmpca_validator.assert_certificate_authority(ca_arn, "PENDING_CERTIFICATE")

    assert cr is not None
    assert k8s.get_resource_exists(ref)
    logging.info(cr)

    resource_arn =  k8s.get_resource_arn(cr)
    assert resource_arn is not None

    yield (ca_arn, resource_arn, secret)

    #Delete Certificate k8s resource
    _, deleted = k8s.delete_custom_resource(ref)
    assert deleted is True

@pytest.fixture(scope="module")
def simple_root_certificate_without_secret_namespace(acmpca_client, create_secret, simple_certificate_authority):
    (ca_cr, ca_name) = simple_certificate_authority
    ca_arn = ca_cr['status']['ackResourceMetadata']['arn']

    cert_name = random_suffix_name("certificate", 30)

    secret = create_secret
    logging.info(secret)
    
    replacements = {}
    replacements["NAME"] = cert_name
    replacements["CA_NAME"] = ca_name
    replacements["CERTIFICATE_SEC_NAME"] = secret.name
    replacements["CERTIFICATE_SEC_KEY"] = secret.key
    replacements["TEMPLATE_ARN"] = "arn:aws:acm-pca:::template/RootCACertificate/V1"

    # Load Certificate CR
    resource_data = load_acmpca_resource(
        "certificate_without_secret_namespace",
        additional_replacements=replacements,
    )

    # Create k8s resource
    ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        cert_name, namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    time.sleep(CREATE_WAIT_AFTER_SECONDS)

    # Check CA status is PENDING_CERTIFICATE
    acmpca_validator = ACMPCAValidator(acmpca_client)
    acmpca_validator.assert_certificate_authority(ca_arn, "PENDING_CERTIFICATE")

    assert cr is not None
    assert k8s.get_resource_exists(ref)
    logging.info(cr)

    resource_arn =  k8s.get_resource_arn(cr)
    assert resource_arn is not None

    yield (ca_arn, resource_arn, secret)

    #Delete Certificate k8s resource
    _, deleted = k8s.delete_custom_resource(ref)
    assert deleted is True

@service_marker
class TestCertificate:

    def test_create_delete(self, acmpca_client, simple_root_certificate):

        (ca_arn, cert_arn, cert_cr, secret) = simple_root_certificate

        # Check certificate is in secret
        _api_client = _get_k8s_api_client()
        api_response = client.CoreV1Api(_api_client).read_namespaced_secret(secret.name, secret.ns).data

        acmpca_validator = ACMPCAValidator(acmpca_client)
        cert = acmpca_validator.get_certificate(ca_arn=ca_arn, cert_arn=cert_arn)

        assert 'certificate' in api_response
        assert base64.b64decode(api_response['certificate']).decode("ascii") == cert

        logging.info(cert_cr['status'].values())
        assert cert not in cert_cr['status'].values()
    
    def test_create_delete_with_ref(self, acmpca_client, simple_root_certificate_with_ref):

        (ca_arn, cert_arn, secret) = simple_root_certificate_with_ref

        # Check certificate is in secret
        _api_client = _get_k8s_api_client()
        api_response = client.CoreV1Api(_api_client).read_namespaced_secret(secret.name, secret.ns).data

        acmpca_validator = ACMPCAValidator(acmpca_client)
        cert = acmpca_validator.get_certificate(ca_arn=ca_arn, cert_arn=cert_arn)

        assert 'certificate' in api_response
        assert base64.b64decode(api_response['certificate']).decode("ascii") == cert
    
    '''def test_create_delete_without_secret_key(self, acmpca_client, simple_root_certificate_without_secret_key):

        (ca_arn, cert_arn, secret) = simple_root_certificate_without_secret_key

        # Check certificate is in secret
        _api_client = _get_k8s_api_client()
        api_response = client.CoreV1Api(_api_client).read_namespaced_secret(secret.name, secret.ns).data

        acmpca_validator = ACMPCAValidator(acmpca_client)
        cert = acmpca_validator.get_certificate(ca_arn=ca_arn, cert_arn=cert_arn)

        assert 'certificate' in api_response
        assert base64.b64decode(api_response['certificate']).decode("ascii") == cert
    
    def test_create_delete_without_secret_namespace(self, acmpca_client, simple_root_certificate_without_secret_namespace):

        (ca_arn, cert_arn, secret) = simple_root_certificate_without_secret_namespace

        # Check certificate is in secret
        _api_client = _get_k8s_api_client()
        api_response = client.CoreV1Api(_api_client).read_namespaced_secret(secret.name, secret.ns).data

        acmpca_validator = ACMPCAValidator(acmpca_client)
        cert = acmpca_validator.get_certificate(ca_arn=ca_arn, cert_arn=cert_arn)

        assert 'certificate' in api_response
        assert base64.b64decode(api_response['certificate']).decode("ascii") == cert'''