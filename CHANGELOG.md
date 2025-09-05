# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Error handling now uses the internal `errors` package, not `github.com/giantswarm/microerror`

## [0.1.0] - 2024-06-03

### Added

- Add [pre-commit](https://pre-commit.com/) hooks
- Add support for loading external schema from URL.

## [0.0.7] - 2024-04-02

### Changed

- Change linear output to avoid HTML, use Markdown formatting only.
- Improve `pkg/generate` unit tests to support updating golden files via `go test ./pkg/generate -update`.

## [0.0.6] - 2024-03-06

### Changed

- Make sure HTML is escaped in linear output.

## [0.0.5] - 2024-02-27

### Added

- New flag layout to allow multiple rendering options

## [0.0.4] - 2023-12-06

### Fixed

- Sort sections by title-then-path to fix unstable order

## [0.0.3] - 2023-11-17

### Fixed

- Fixed rendering of global properties

## [0.0.2] - 2023-11-16

### Fixed

- Fixed root command name in the help text
- Added CI pipeline & adjusted the code to pass the checks
- Render separate sections for global properties like for top-level properties

## [0.0.1] - 2023-04-28

- Initial release

[Unreleased]: https://github.com/giantswarm/schemadocs/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/giantswarm/schemadocs/compare/v0.0.7...v0.1.0
[0.0.7]: https://github.com/giantswarm/schemadocs/compare/v0.0.6...v0.0.7
[0.0.6]: https://github.com/giantswarm/schemadocs/compare/v0.0.5...v0.0.6
[0.0.5]: https://github.com/giantswarm/schemadocs/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/giantswarm/schemadocs/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/giantswarm/schemadocs/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/giantswarm/schemadocs/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/schemadocs/releases/tag/v0.0.1
