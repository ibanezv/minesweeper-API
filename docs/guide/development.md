# acq-refund-notifier - Development

## Golang

### Get Starting
See which is the Go version used, into `go.mod` file. 
Example
```
go 1.16
```
means we are using Go go 1.16


Although this project already uses Go Modules to manage dependencies, it is recommended to clone the project into a $GOPATH directory.
It is highly recommended that you run: 
```
$ go mod tidy
```

Some tools and dependencies the projects uses are below. 
Important: before to run them, please verify you aren't in a directory that has a `go.mod` file (because if you are there, the dependencies will be added to it, making the proyect dependent of them). 

### Which IDE can we use ?
We are using VS Code which is free or IntelliJ IDEA + Go Plugin, but sometimes you could not get any licence.

### To Run the Tests
To run the tests use the command `go test ./...` or  `fury test`

### Contribution Guidelines

Please be sure to read the [Contribution Guidelines](./../../CONTRIBUTE.md). It contains important information and practices the team agreed.


The followings are the general steps to correctly deploy this application:
1. Once you have all the features tested that you want to deploy inside develop branch, you must document them in the [Changelog](./../../CHANGELOG.md) file
2. Create release branch with the according version. Example of branch's name: `release/1.0.1`
3. Create the Pull Request to trigger the Continuous Integration process. This will create the candidate version in Fury. Example candidate version's name: `1.0.1-rc-1`
4. Merge the Pull Request into master. Fury will create the productive version when this has done. Example of productive version's name: `1.0.1` 
5. Deploy this productive version into `production` scope and do regression tests plus new features tests.
6. It's advisable that you watch the {app.name}'s Dashboards to make sure that the deploy has worked.

## Project Structure
This project follows [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html). Please read this blog post in order to understand 

[Back home](/README.md)