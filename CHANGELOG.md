# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added

### Fixed

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
