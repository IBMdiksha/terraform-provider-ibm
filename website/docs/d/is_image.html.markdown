---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Image"
description: |-
  Manages IBM Cloud infrastructure image.
---

# ibm_is_image
Retrieve information of an existing IBM Cloud Infrastructure image as a read-only data source. For more information, about VPC custom images, see [IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform

data "ibm_is_image" "example" {
  name = "centos-7.x-amd64"
}
```
```terraform
data "ibm_is_image" "example" {
  identifier = ibm_is_image.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `identifier` - (Optional, String) The id of the image.
    ~> **Note:** `name` and `identifier` are mutually exclusive.
- `name` - (Optional, String) The name of the image.
    ~> **Note:** `name` and `identifier` are mutually exclusive.
- `visibility` - (Optional, String) The visibility of the image. Accepted values are `public` or `private`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `architecture` - (String) The architecture of the image.
- `checksum`-  (String) The `SHA256` checksum of the image.
- `crn` - (String) The CRN for this image.
- `encryption` - (String) The type of encryption used of the image.
- `encryption_key`-  (String) The CRN of the Key Protect or Hyper Protect Crypto Service root key for this resource.
- `id` - (String) The unique identifier of the image.
- `os` - (String) The name of the operating system.
- `status` - (String) The status of this image.
- `source_volume` - The source volume id of the image.
