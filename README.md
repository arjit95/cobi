[![Build Status][build-shield]][build-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![Code Coverage][coverage-shield]][coverage-url]
<p align="center">
  <h3 align="center">cobi</h3>

  <p align="center">
    cobi (cobra Interactive) is a small wrapper on top of <a href="https://github.com/spf13/cobra">cobra</a> and <a href="https://github.com/rivo/tview">tview</a> to build interactive cli applications
    <br />
    <a href="https://github.com/arjit95/cobi#docs"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/arjit95/_examples"> Examples </a>
    ·
    <a href="https://github.com/arjit95/cobi/issues">Report Bug</a>
    ·
    <a href="https://github.com/arjit95/cobi/issues">Request Feature</a>
  </p>
</p>

## Table of Contents

* [About the Project](#about-the-project)
* [Overview](#overview)
* [Installing](#installing)
* [Migrating from cobra](#migrating-from-cobra)
* [Contributions](#contributions)
* [License](#license)
* [Acknowledgements](#acknowledgements)

## About The Project
![cobi Screenshot](_images/screenshot.svg)

Cobra provides a great way to build cli applications whereas go prompt provides powerful interactive prompts. But there are scenarios where the cli application needs to execute a long running task, for eg _port forwarding in kubernetes_. This could be solved with an interactive prompt, whereas normal operations can still work with the default cli application.

## Overview
cobi works by using command completion provided by cobra. These completions are propagated to go prompt providing almost the same experience in both interactive and cli modes. Since cobi implements the same interface as cobra, it becomes very easy to port your existing project to cobi.

## Installing
First, use go get to install the latest version of the library. This command will download cobi with all its dependencies:

```bash
go get -u github.com/arjit95/cobi
```

Next, include cobi in your app:

```go
import "github.com/arjit95/cobi"
```

## Migrating from cobra
Suppose you have an existing cobra command:

```go
import "github.com/spf13/cobra"

var cmd = &cobra.Command{
  Use:   "demo",
  Short: "This is a demo cobra command",
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}
```

This would be re-written as

```go
import (
    "github.com/arjit95/cobi"
    "github.com/spf13/cobra"
)

var cmd = &cobi.Command {
    Command: &cobra.Command{
        Use:   "demo",
        Short: "This is a demo cobra command",
        Run: func(cmd *cobra.Command, args []string) {
            // Do Stuff Here
        }, 
    },
}

// Execute the command normally
cmd.Execute()

// Alternaitvely run the command in interactive mode
// Ctrl+C to exit
cmd.RunInteractive()
```
You only need to wrap the top most command, to use the interactive mode.

## Contributions
Contributions are welcome, this was my first experience with golang so a lot of things may not be up to the mark. Feel free to open an issue or raise a pull request.

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Acknowledgements
- [cobra](https://github.com/spf13/cobra)
- [tview](https://github.com/rivo/tview)

[build-shield]: https://travis-ci.com/arjit95/cobi.svg?branch=main
[build-url]: https://travis-ci.com/arjit95/cobi
[issues-shield]: https://img.shields.io/github/issues/arjit95/cobi.svg
[issues-url]: https://github.com/arjit95/cobi/issues
[license-shield]: https://img.shields.io/github/license/arjit95/cobi.svg
[license-url]: https://github.com/arjit95/cobi/blob/main/LICENSE
[coverage-shield]: https://codecov.io/gh/arjit95/cobi/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/arjit95/cobi