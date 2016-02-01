# BOSH Certification Receipt Generator

Generates receipts for certified BOSH release + stemcell combinations.

## Usage

```
./certify -v
  # prints version information

./certify --release bosh/250 --release bosh-aws-cpi/42 --stemcell bosh-aws-xen-hvm-ubuntu-trusty-go_agent/3184.1
  # flags:
  #   --release value (multiple flags allowed)
  #     where value is name/version
  #   --stemcell value (single flag allowed)
  #     where value is name/version

  # output
  {
    "releases": [
      {
        "name": "bosh",
        "version": "250"
      },
      {
        "name": "bosh-aws-cpi",
        "version": "42"
      }
    ],
    "stemcell": {
      "name": "bosh-aws-xen-hvm-ubuntu-trusty-go_agent",
      "version": "3184.1"
    }
  }
```
