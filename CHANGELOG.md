# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased
### Added
- format should write zero value of boolean in settings, by [@bradrydzewski](https://github.com/bradrydzewski).
- Support for windows server 1903, by [@bradrydzewski](https://github.com/bradrydzewski).

## [1.2.2] - 2019-07-29
### Added
- Ability to configure network_mode, by [@bradrydzewski](https://github.com/bradrydzewski).
- Convert legacy branch filter to ref trigger, by [@bradrydzewski](https://github.com/bradrydzewski).
- Convert legacy deployment event to promotion event, by [@bradrydzewski](https://github.com/bradrydzewski).

## [1.2.1] - 2019-07-17
### Added
- Pull if-not-exists when converting legacy yaml files, by [@bradrydzewski](https://github.com/bradrydzewski).
- Improve workspace support when converting legacy yaml files, by [@bradrydzewski](https://github.com/bradrydzewski).
- Improve registry secret support when converting legacy yaml files, by [@bradrydzewski](https://github.com/bradrydzewski).

## [1.2.0] - 2019-07-16
### Added
- Added Action field to trigger and when clause, by [@bradrydzewski](https://github.com/bradrydzewski).
- Improve escaping when marshaling to yaml, by [@bradrydzewski](https://github.com/bradrydzewski).
- Handle duplicate step names when converting legacy configurations, by [@bradrydzewski](https://github.com/bradrydzewski).
- Handle dot workspace path when converting legacy configurations, by [@bradrydzewski](https://github.com/bradrydzewski).

## [1.1.1] - 2019-05-30
### Fixed
- Retain order of steps when converting legacy pipelines with merge keys, by [@bradrydzewski](https://github.com/bradrydzewski).


## [1.1.0] - 2019-05-30
### Fixed
- Support for yaml merge keys, by [@bradrydzewski](https://github.com/bradrydzewski).
- Improve how colon characters are escaped, by [@bradrydzewski](https://github.com/bradrydzewski). Issue [#45](https://github.com/drone/drone-yaml/issues/45).
- Improve how pipe and caret characters are escaped, by [@bradrydzewski](https://github.com/bradrydzewski). Issue [#44](https://github.com/drone/drone-yaml/issues/44).
- Error when empty document or missing kind attribute, by [@bradrydzewski](https://github.com/bradrydzewski). Issue [#42](https://github.com/drone/drone-yaml/issues/42).

## [1.0.9] - 2019-05-20
### Added
- Only lint resources of kind pipeline and of type docker, by [@bradrydzewski](https://github.com/bradrydzewski).

## [1.0.8] - 2019-04-13
### Added
- Support Cron job name in When clause, by [@bradrydzewski](https://github.com/bradrydzewski).

## [1.0.7] - 2019-04-10
### Added
- Optionally set the Docker container User, by [@bradrydzewski](https://github.com/bradrydzewski).
