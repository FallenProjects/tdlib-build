# Automated TDLib Build & JSON Generator 

![Build Status](https://github.com/FallenProjects/tdlib-build/actions/workflows/build-tdlib.yml/badge.svg)
![Go Version](https://img.shields.io/github/go-mod/go-version/FallenProjects/tdlib-build)
![License](https://img.shields.io/github/license/FallenProjects/tdlib-build)

An automated workflow to build [TDLib](https://github.com/tdlib/td) (Telegram Database Library) and generate the essential [tdlib.json](tdlib.json) schema file for use in various wrappers and bots.

## Features 

*   **Automatic `tdlib.json` Generation**:
    *   Fetches the latest `td_api.tl` schema.
    *   Parses it into a JSON format compatible with `tdlib-json` wrappers.
*   **Cross-Platform Builds**:
    *   **Linux** (x86_64, arm64) - Built with `cmake`, `gperf`, `zlib`, `openssl`.
    *   **macOS** (x86_64, arm64) - Built with `brew` dependencies.
    *   **Windows** (x64) - Built with `vcpkg` dependencies.
*   **GitHub Releases Integration**:
    *   Automatically creates a new Release tagged with the TDLib version (e.g., `v1.8.0`).
    *   Uploads compiled shared libraries (`libtdjson.so`, `libtdjson.dylib`, `tdjson.dll`) as release assets.
