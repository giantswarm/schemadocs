[![CircleCI](https://circleci.com/gh/giantswarm/template.svg?style=shield)](https://circleci.com/gh/giantswarm/template)

# schemadocs - the Giant Swarm JSON Schema docs generator 

schemadocs helps you generate documentation from your JSON schema and store it in README and other documentation files

## Features

- Generate Markdown documentation from a JSON schema file and inject it into an output text file
- Validate generated Markdown documentation in a text file and check whether it is aligned with documentation generated from the supplied JSON schema file

## Installation

```nohighlight
go install github.com/giantswarm/schemadocs@latest
```

## Usage

### Generating documentation

Executing `schemadocs generate` without any options will generate Markdown documentation from a JSON schema file and store it in a `README.md` file in the current working directory. 

It is required that the README file contains exactly one pair of placeholder strings, which mark the start and end of the documentation. 
By default, the placeholders are `<!-- DOCS_START -->` and `<!-- DOCS_END -->`.

If the README file does not exist or in case it does not contain one pair of valid placeholder strings, the execution will fail.

```nohighlight
$ schemadocs generate schema.json

Generating documentation
Schema: schema.json
Destination file: README.md
Using default placeholders '<!-- DOCS_START -->' and '<!-- DOCS_END -->'

[SUCCESS] Documentation generated successfully!
```

To generate the documentation to a custom file, apply the `--output-path` option.
To use a pair of custom placeholder strings, apply the `--doc-placeholder-start` and `--doc-placeholder-end`.

```nohighlight
$ schemadocs generate schema.json --output-path README.md --doc-placeholder-start <!-- DOCS_START --> --doc-placeholder-end <!-- DOCS_END -->
```

**Note:** The `--doc-placeholder-start` and `--doc-placeholder-end` need to be provided as a pair. So, if one of them is specified, the other one needs to be specified, too.

Use `--help` to learn about more options.

### Validation

Validate documentation in a text file by comparing it to documentation generated from a JSON schema file specified in the required `--schema` option.
The validation passes in case both documentations match exactly.

It is required that the text file contains exactly one pair of placeholder strings, which mark the start and end of the documentation.
By default, the placeholders are `<!-- DOCS_START -->` and `<!-- DOCS_END -->`.

If the text file or the schema file do not exist or in case the text file does not contain one pair of valid placeholder strings, the execution will fail.

```nohighlight
$ schemadocs validate README.md --schema schema.json

Validating documentation
Source file: README.md
Schema file: schema.json
Using default placeholders '<!-- DOCS_START -->' and '<!-- DOCS_END -->'

[SUCCESS] Documentation is valid!
```

To use a pair of custom placeholder strings, apply the `--doc-placeholder-start` and `--doc-placeholder-end`.

```nohighlight
$ schemadocs validate README.md --schema schema.json --doc-placeholder-start [DOCS_START] --doc-placeholder-end [DOCS_END]
```

**Note:** The `--doc-placeholder-start` and `--doc-placeholder-end` need to be provided as a pair. So, if one of them is specified, the other one needs to be specified, too.

Use `--help` to learn about more options.
