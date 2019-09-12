# Boilr Go Template

This is a highly opinionated project template for creating go projects using [boilr](https://github.com/tmrts/boilr)

## Features

This template has the following features:

* License file
* Libraries
  * Cobra for CLI creation
  * Viper for configuration loading
  * Logrus for logging
  * Testify for unit testing
* VS Code
  * Support for remote containers
  * Some common tasks
  * Launch configuration
* CI/CD
  * Travis builds
  * Build fail on badly fomatted code
  * Codecov reports
  * Goreleaser releases
    * Binaries to Github
    * Docker images
* Badges
  * Latest release
  * Build status
  * Code coverage
  * Go report
  * Go doc
  * License

## Install

`boilr template download gbbirkisson/boilr-go-template gbbirkisson-go`

## Usage

`boilr template use gbbirkisson-go <my-project-dir>`

## Builds failing because of formatting

This template is very strict when it comes to builds. Run the VS Code task `Report` on your project to figure out what the problem.

## More Information

- [Go Standards - Project Layout](https://github.com/golang-standards/project-layout)
- [boilr](https://github.com/tmrts/boilr)
