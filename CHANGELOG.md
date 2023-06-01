<!-- markdownlint-disable MD024 -->
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org).

## [v0.3.0](https://github.com/chelnak/ysmrr/tree/v0.3.0) - 2023-06-01

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.2.1...v0.3.0)

### Added

- Stop handling SIGINT [#55](https://github.com/chelnak/ysmrr/pull/55) ([chelnak](https://github.com/chelnak))

## [v0.2.1](https://github.com/chelnak/ysmrr/tree/v0.2.1) - 2022-11-11

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.2.0...v0.2.1)

### Fixed

- fix race condition in `(*Spinner).Print` [#46](https://github.com/chelnak/ysmrr/pull/46) ([fortytw2](https://github.com/fortytw2))

## [v0.2.0](https://github.com/chelnak/ysmrr/tree/v0.2.0) - 2022-09-01

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.1.1...v0.2.0)

### Added

- Add default spinner speeds [#43](https://github.com/chelnak/ysmrr/pull/43) ([chelnak](https://github.com/chelnak))
- Add UpdateMessagef() method [#42](https://github.com/chelnak/ysmrr/pull/42) ([chelnak](https://github.com/chelnak))
- Enable realtime message updates [#41](https://github.com/chelnak/ysmrr/pull/41) ([chelnak](https://github.com/chelnak))

## [v0.1.1](https://github.com/chelnak/ysmrr/tree/v0.1.1) - 2022-08-26

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.1.0...v0.1.1)

### Fixed

- Fix line updates [#35](https://github.com/chelnak/ysmrr/pull/35) ([chelnak](https://github.com/chelnak))
- Fix incorrect line counts [#34](https://github.com/chelnak/ysmrr/pull/34) ([chelnak](https://github.com/chelnak))

## [v0.1.0](https://github.com/chelnak/ysmrr/tree/v0.1.0) - 2022-08-14

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.8...v0.1.0)

### Added

- Improve TTY detection [#30](https://github.com/chelnak/ysmrr/pull/30) ([chelnak](https://github.com/chelnak))

## [v0.0.8](https://github.com/chelnak/ysmrr/tree/v0.0.8) - 2022-08-13

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.7...v0.0.8)

### Added

- Simplify enums [#28](https://github.com/chelnak/ysmrr/pull/28) ([chelnak](https://github.com/chelnak))
- Rename spinner package [#27](https://github.com/chelnak/ysmrr/pull/27) ([chelnak](https://github.com/chelnak))
- Better charmap selection [#25](https://github.com/chelnak/ysmrr/pull/25) ([chelnak](https://github.com/chelnak))

## [v0.0.7](https://github.com/chelnak/ysmrr/tree/v0.0.7) - 2022-07-21

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.6...v0.0.7)

### Added

- Update NewSpinnerManager constructor [#22](https://github.com/chelnak/ysmrr/pull/22) ([chelnak](https://github.com/chelnak))
- Set frameDuration default [#21](https://github.com/chelnak/ysmrr/pull/21) ([chelnak](https://github.com/chelnak))

### Fixed

- Fix line rendering [#23](https://github.com/chelnak/ysmrr/pull/23) ([chelnak](https://github.com/chelnak))

## [v0.0.6](https://github.com/chelnak/ysmrr/tree/v0.0.6) - 2022-07-20

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.5...v0.0.6)

### Added

- Add Windows support [#18](https://github.com/chelnak/ysmrr/pull/18) ([chelnak](https://github.com/chelnak))
- Manager updates [#17](https://github.com/chelnak/ysmrr/pull/17) ([chelnak](https://github.com/chelnak))
- Improve test coverage for color pkg [#16](https://github.com/chelnak/ysmrr/pull/16) ([chelnak](https://github.com/chelnak))
- Improve test coverage for ysmrr pkg [#15](https://github.com/chelnak/ysmrr/pull/15) ([chelnak](https://github.com/chelnak))
- Update examples [#14](https://github.com/chelnak/ysmrr/pull/14) ([chelnak](https://github.com/chelnak))
- Handle SIGINT and SIGTERM [#13](https://github.com/chelnak/ysmrr/pull/13) ([chelnak](https://github.com/chelnak))
- Add WithMessageColor option [#12](https://github.com/chelnak/ysmrr/pull/12) ([chelnak](https://github.com/chelnak))
- Revert spinner interface [#11](https://github.com/chelnak/ysmrr/pull/11) ([chelnak](https://github.com/chelnak))

## [v0.0.5](https://github.com/chelnak/ysmrr/tree/v0.0.5) - 2022-07-17

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.4...v0.0.5)

### Added

- Refactor [#9](https://github.com/chelnak/ysmrr/pull/9) ([chelnak](https://github.com/chelnak))

## [v0.0.4](https://github.com/chelnak/ysmrr/tree/v0.0.4) - 2022-07-17

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.3...v0.0.4)

### Added

- Add some tests [#7](https://github.com/chelnak/ysmrr/pull/7) ([chelnak](https://github.com/chelnak))

### Fixed

- Fix linter [#6](https://github.com/chelnak/ysmrr/pull/6) ([chelnak](https://github.com/chelnak))

## [v0.0.3](https://github.com/chelnak/ysmrr/tree/v0.0.3) - 2022-07-17

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.2...v0.0.3)

### Added

- Enhanced configuration [#4](https://github.com/chelnak/ysmrr/pull/4) ([chelnak](https://github.com/chelnak))

## [v0.0.2](https://github.com/chelnak/ysmrr/tree/v0.0.2) - 2022-07-17

[Full Changelog](https://github.com/chelnak/ysmrr/compare/v0.0.1...v0.0.2)

### Fixed

- Fix Cnorm behaviour [#2](https://github.com/chelnak/ysmrr/pull/2) ([chelnak](https://github.com/chelnak))

## [v0.0.1](https://github.com/chelnak/ysmrr/tree/v0.0.1) - 2022-07-17

[Full Changelog](https://github.com/chelnak/ysmrr/compare/634c76085ea0215b5e9629847cc94995bc7575f6...v0.0.1)
