<br/>
<br/>
<p align="center">
  <strong>✨ Oversee all your CODEOWNERS needs</strong>
</p>
<br/>

## Overview

<a href="https://twitter.com/m_szostok"><img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/m_szostok?color=a&label=Follow%20%40m_szostok%20for%20updates&style=social"></a>

**Codeowners** is a GitHub CLI extension designed to simplify managing and validating CODEOWNERS files directly from your terminal or CI environment. While GitHub's CODEOWNERS feature is powerful, interpreting the CODEOWNERS data outside of GitHub can be challenging. This tool helps streamline that process.

Like the idea? Give a GitHub star ⭐!

## Installation

To install the `codeowners` CLI extension, use the following command:

```sh
gh extension install github.com/gitcheasy/gh-codeowners
```

## Usage

Here are some common commands for using the `codeowners` extension:

```sh
# Validate CODEOWNERS across all repositories
gh codeowners validate -all

# Validate CODEOWNERS for the current repository
gh codeowners validate
```

### Exit Status Codes

The application uses exit status codes to indicate different types of errors:

| Code  | Description                                                                     |
|:-----:|:--------------------------------------------------------------------------------|
| **1** | Application startup failed due to incorrect configuration or an internal error. |
| **3** | CODEOWNERS validation failed due to issues found during checks.                 |

> **Tip:**  
> To prevent the CLI from failing when validation issues are detected, you can set the exit code for validation issues to `0`. Run the following command:  
> `gh codeowners validate --issues-exit-code 0`.

## Contributing

We welcome contributions from the community! To contribute, please follow the standard GitHub pull request process.