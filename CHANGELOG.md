## [0.6.1](https://github.com/apathetic-tools/sheave/compare/v0.6.0...v0.6.1) (2026-06-18)


### Bug Fixes

* deterministic registry resolution and settings sync ([987139b](https://github.com/apathetic-tools/sheave/commit/987139b626615aed39feccac9fe0aa954abed984))

# [0.6.0](https://github.com/apathetic-tools/sheave/compare/v0.5.0...v0.6.0) (2026-06-17)


### Features

* **skills:** add new instruct and troubleshooting skills ([f9475ca](https://github.com/apathetic-tools/sheave/commit/f9475ca5c48be2d8d4a737a7b2c18eb52b343ae9))

# [0.5.0](https://github.com/apathetic-tools/sheave/compare/v0.4.3...v0.5.0) (2026-06-17)


### Features

* implement POSIX path traversal and security jailing for sheave-family ([54234c8](https://github.com/apathetic-tools/sheave/commit/54234c839a6eaab86800f131af67b0e04c20a54c))

## [0.4.3](https://github.com/apathetic-tools/sheave/compare/v0.4.2...v0.4.3) (2026-06-17)


### Bug Fixes

* specify main package path for goreleaser ([5a97a29](https://github.com/apathetic-tools/sheave/commit/5a97a29e9bee512adf8d5149037139f02174f8b0))

## [0.4.2](https://github.com/apathetic-tools/sheave/compare/v0.4.1...v0.4.2) (2026-06-17)


### Bug Fixes

* remove invalid github name template in goreleaser ([52c7dd7](https://github.com/apathetic-tools/sheave/commit/52c7dd7aa96f4fd62ba8808dd4212c80c99d0abb))

## [0.4.1](https://github.com/apathetic-tools/sheave/compare/v0.4.0...v0.4.1) (2026-06-17)


### Bug Fixes

* pass release-notes via CLI argument instead of yaml ([278e93c](https://github.com/apathetic-tools/sheave/commit/278e93c04ae8778fa089bd60e1ee2b89e617e376))

# [0.4.0](https://github.com/apathetic-tools/sheave/compare/v0.3.0...v0.4.0) (2026-06-17)


### Bug Fixes

* bump semantic-release dependencies ([b6458e1](https://github.com/apathetic-tools/sheave/commit/b6458e19e266d34a748577846eb298e675e450aa))


### Features

* add active_providers list and openai default provider ([5600dc2](https://github.com/apathetic-tools/sheave/commit/5600dc2322fff2738faa6849fe2e4d30c13ec768))
* complete Phase 1 implementation ([57d186c](https://github.com/apathetic-tools/sheave/commit/57d186cf95f8d3b697263d1bd765194acfc4befc))
* implement data-driven provider config schema ([359e02a](https://github.com/apathetic-tools/sheave/commit/359e02a5251bb56f5d02ebccd15ea3fa5f7beaf2))
* implement double-slash path prefix for project root resolution ([ed0713e](https://github.com/apathetic-tools/sheave/commit/ed0713eff45fa1e0aec965cb54248f3e95ed37db))
* implement dynamic AI provider sync and memory index ([18a6444](https://github.com/apathetic-tools/sheave/commit/18a6444c27f550d8b740aba3912496316a0b3263))
* implement sheave sync command (Phase 2 IDE Integration) ([8e956e5](https://github.com/apathetic-tools/sheave/commit/8e956e595b28f781cbcfbb9ba8acba143827bdca))
* parse all .md* extensions in the registry engine ([ffeb68e](https://github.com/apathetic-tools/sheave/commit/ffeb68eb09483dff36479286b381830301980dc4))
* port python list_project.py to sheave project command ([d49924f](https://github.com/apathetic-tools/sheave/commit/d49924fe968aa1da33f00c8e05915df8ecc311bc))
* prompt user to auto-sync after 'sheave init' ([75e0ecb](https://github.com/apathetic-tools/sheave/commit/75e0ecb0a0ed641cab50c829b6b9c24d461f98ff))
* redesign config format to TOML and introduce .ai structure ([8de1281](https://github.com/apathetic-tools/sheave/commit/8de128175497819b99ae81856488798703195287))
* redesign init/scaffold defaults and introduce interactive guide command ([27894e2](https://github.com/apathetic-tools/sheave/commit/27894e239ab911942af9724f80ff9be859597da8))
* **registry:** scaffold builtin registry and update ROADMAP ([fa74968](https://github.com/apathetic-tools/sheave/commit/fa7496822168eea495703749089701df08b21628))
* **skills:** add frontmatter and generalize check/ci skills ([28c3c05](https://github.com/apathetic-tools/sheave/commit/28c3c055b60393237f902f7f0931592f225a4440))

# CHANGELOG

<!-- version list -->

## v0.3.0 (2025-12-03)

### Features

- Update project metadata and classifiers
  ([`892cfcc`](https://github.com/apathetic-tools/sheave/commit/892cfcc1e965e73ec563e04d0708d4505c9a0ffd))

### Refactoring

- Use cwd instead of __file__ for project root
  ([`cbab414`](https://github.com/apathetic-tools/sheave/commit/cbab4144ba42c2887e104ef603097a4449e6ace4))


## v0.2.0 (2025-12-03)

### Build System

- **deps**: Bump actions/checkout from 5 to 6
  ([`bf1dd3a`](https://github.com/apathetic-tools/sheave/commit/bf1dd3a1a88311fd371da07e77bfc5954ba57212))

### Chores

- Include bin scripts in PyPI package
  ([`767e952`](https://github.com/apathetic-tools/sheave/commit/767e9522f54e2733bebc57c0cd087c23fbab8c50))

### Continuous Integration

- Install root package in PyPI workflow
  ([`dbeb685`](https://github.com/apathetic-tools/sheave/commit/dbeb68514013af2502b872acd6a75d7ae470e117))

### Features

- Add CLI skeleton with main entry point
  ([`271b613`](https://github.com/apathetic-tools/sheave/commit/271b6139ab7c00854b3fe3b9173ae7c020fe338c))


## v0.1.0 (2025-12-01)

- Initial Release
