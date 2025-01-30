# TFLint Ruleset for terraform-provider-juju
[![Build Status](https://github.com/canonical/tflint-ruleset-juju/workflows/build/badge.svg?branch=main)](https://github.com/canonical/tflint-ruleset-juju/actions)

TFLint ruleset plugin for [Terraform Provider Juju](https://registry.terraform.io/providers/juju/juju/latest/docs).

## Requirements

- TFLint v0.42+
- Go v1.23

## Installation

You can install the plugin by adding a config to `.tflint.hcl` and running `tflint --init`:

```hcl
plugin "juju" {
    enabled = true
    version = "0.1.0"
    source  = "github.com/canonical/tflint-ruleset-juju"
}
```


## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|juju_application_invalid_name|Disallow using invalid name|ERROR|✔||


## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "juju" {
  enabled = true
}
EOS
$ tflint
```
