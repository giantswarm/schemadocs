[![CircleCI](https://circleci.com/gh/giantswarm/template.svg?style=shield)](https://circleci.com/gh/giantswarm/template)

# schemadocs - the Giant Swarm JSON Schema docs generator 

- Generates Markdown documentation from JSON schema files and injects it into text files
- Validates generated Markdown documentation in text files and checks whether it is aligned with the supplied JSON schema file

## Usage

```nohighlight
schemadoc generate path/to/myschema.json

schemadoc validata --schema path/to/myschema.json path/to/text-file.md
```
