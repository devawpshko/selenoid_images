# Selenoid Browser Images

This repository is a fork designed to provide developers and QA engineers the ability to build `chrome` (and potentially other) browser images independently. This is necessary due to Aerokube discontinuing the publication of pre-built browser images.

The original repository has not been updated to accommodate recent changes made by Google in Chrome browser distribution. Specifically, Chrome can no longer be installed on Linux systems using the `apt-get` command, which was previously used to build Chrome Docker images. This repository addresses that issue by allowing custom image builds.

---

## Table of Contents

1. [Requirements](#requirements)
2. [Getting Started](#getting-started)
3. [Building Chrome Images](#building-chrome-images)
4. [Additional Notes](#additional-notes)
5. [Prebuilt Images](#prebuilt-images)

---

## Requirements

Before you begin, ensure the following prerequisites are met:

- [Go (Golang)](https://go.dev/) is installed on your system.

---

## Getting Started

Follow these steps to set up the repository and build the necessary tools:

1. Clone this repository:
   ```bash
   git clone https://github.com/devawpshko/selenoid_images.git
   ```
2. Install the `pkger` tool for Go:
   ```bash
   go install github.com/markbates/pkger/cmd/pkger@latest
   ```
3. Navigate into the `selenoid_images` directory:
   ```bash
   cd selenoid_images
   ```
4. Generate required files:
   ```bash
   go generate
   ```
5. Build the executable:
   ```bash
   go build
   ```

---

## Building Chrome Images

To build a Chrome image, follow these steps:

1. Visit the [Chrome for Testing](https://googlechromelabs.github.io/chrome-for-testing/) page and find the version of Chrome you need.
2. Use the command below to build the image, replacing `<base-version>` and `<driver-version>` with the appropriate Chrome and driver versions, respectively:
   ```bash
   ./images chrome -b <base-version> -d <driver-version> -t selenoid/chrome:<tagname>
   ```
   Example:
   ```bash
   ./images chrome -b 132.0.6834.159 -d 132.0.6834.159 -t selenoid/chrome:132.0
   ```

   **Note:** The `-b` (base version) and the `-d` (driver version) options must match.

---

## Additional Notes

- Ensure you replace placeholders in commands with the actual version numbers obtained from the [Chrome for Testing](https://googlechromelabs.github.io/chrome-for-testing/) page.
- The build process has been tested only with Chrome browser images.

---

## Prebuilt Images

If you want to use prebuilt Chrome images instead, they are available on Docker Hub:  
[Docker Hub - devawpshko/selenoid_chrome](https://hub.docker.com/r/devawpshko/selenoid_chrome/tags)

---

We hope this repository helps simplify your testing setup. Feel free to contribute or raise issues if you encounter any problems!