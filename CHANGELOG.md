# Changelog

All notable changes to this project are documented here.
The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and the project uses [Semantic Versioning](https://semver.org/).

## [1.3.0] - 2026-11-15

### Added
- `sshed show` displays every parameter in the host block (e.g. `LocalForward`,
  `ProxyJump`, `StrictHostKeyChecking`), in file order, instead of only
  `Hostname`, `Port`, `User` and `IdentityFile`.
- Repeated keys such as `LocalForward` and `RemoteForward` are grouped under a
  single label, one value per line.
- Inline comments on parameter lines and on the `Host` line are shown.

- Release builds for `arm64` (Apple Silicon macOS, Linux and Windows).

### Fixed
- `sshed --version` / `-v` now works. The version is taken from goreleaser's
  default ldflags (`main.version`, `main.commit`, `main.date`), falling back to
  the build info Go embeds in the binary; previously the flag was never enabled.

### Changed
- `sshed show` no longer prints empty `Port`/`User` lines for hosts that don't
  set them.

## Earlier releases

See [upstream releases](https://github.com/trntv/sshed/releases) for 1.2.0 and
before.

[1.3.0]: https://github.com/shambles07/sshed/compare/1.2.0...v1.3.0
