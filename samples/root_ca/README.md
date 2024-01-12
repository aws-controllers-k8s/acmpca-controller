# Root CA Activation Sample

This sample demonstrates how to create and activate a root CA using the AWS Controllers for Kubernetes (ACK) service controller for AWS Private Certificate Authority.

## Create and Activate a root CA

To create your root CA, modify the placeholder values in `certificate_authority.yaml` and apply the manifest to your Kubernetes cluster using the following command:

```
kubectl apply -f certificate_authority.yaml
```

## Issue a root certificate

To issue the root certificate and create a Secret (which the certificate will be stored in), modify the placeholder values in `certificate.yaml` and apply the specification to your Kubernetes cluster using the following command:

```
kubectl apply -f certificate.yaml
```
You will be in charge of deleting the Secret when you no longer have use for it.

## Activate the root CA

### Get the root CA ARN

The ARN of the root CA is needed to create the CertificateAuthorityActivation resource. To get the ARN of the root CA, describe the root CA using the following command:

```
kubectl describe .....
```

### Create CertificateAuthorityActivation resource

To create the CertificateAuthorityActivation resource, modify the placeholder values in `certificate_authority_activation.yaml` and apply the specification to your Kubernetes cluster using the following command:

```
kubectl apply -f certificate_authority_activation.yaml
```

Describing the CA should now show it as active.

## Delete resources

Disable and delete by doing:

Add a field in the root CA spec called 'status' with a value of "DISABLED"

then delete 
kubectl delete -f bucket.yaml

# verify the bucket no longer exists
kubectl get bucket/$BUCKET_NAME