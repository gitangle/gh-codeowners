```ruby
gh extension install github.com/gitcheasy/gh-codeowners
```

<p align="center">
  <strong>✨ Oversee all your CODEOWNERS needs</strong>
</p>

GitHub CLI extension designed to simplify managing and validating CODEOWNERS files directly from your terminal or CI environment.



Like the idea? Give a GitHub star ⭐!


## Quick Start

To install the `codeowners` CLI extension, run:

```sh
gh extension install github.com/gitcheasy/gh-codeowners
```

Navigate to your repository and run:
```sh
# Validate CODEOWNERS across all repos for the GitHub owner taken from the current repository directory.
gh codeowners validate -all

# Validate CODEOWNERS for a specific owner across all their repos
gh codeowners validate -owner "mszostok"  -all
```

## Why?

While GitHub's CODEOWNERS feature is powerful, interpreting the CODEOWNERS data outside of GitHub can be challenging. This tool helps streamline that process.

### Exit Status Codes

The application uses exit status codes to indicate different types of errors:

| Code  | Description                                                                     |
|:-----:|:--------------------------------------------------------------------------------|
| **1** | Application startup failed due to incorrect configuration or an internal error. |
| **3** | CODEOWNERS validation failed due to issues found during checks.                 |

> [!TIP]
> To prevent the CLI from failing when validation issues are detected, you can set the exit code for validation issues to `0`. Run the following command:  
> `gh codeowners validate --issues-exit-code 0`.

## Contributing

We welcome contributions from the community! To contribute, please follow the standard GitHub pull request process.