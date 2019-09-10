<h1>{{Name}}</h1>

<p align="center">
<a href="https://github.com/{{GithubUser}}/{{Name}}/releases/latest" target="_blank"><img src="https://img.shields.io/github/release/{{GithubUser}}/{{Name}}.svg?style=flat-square"/></a>{{ if TravisCI }}<a href="https://travis-ci.org/{{GithubUser}}/{{Name}}" target="_blank"><img src="https://img.shields.io/travis/{{GithubUser}}/{{Name}}/master.svg?style=flat-square"/></a>{{ if Codecov }}<a href="http://codecov.io/github/{{GithubUser}}/{{Name}}?branch=master" target="_blank"><img src="https://img.shields.io/codecov/c/github/{{GithubUser}}/{{Name}}/master.svg?style=flat-square"/></a>{{ end }}{{ end }}<a href="https://goreportcard.com/report/github.com/{{GithubUser}}/{{Name}}" target="_blank"><img src="https://goreportcard.com/badge/github.com/{{GithubUser}}/{{Name}}?style=flat-square"/></a><a href="https://godoc.org/github.com/{{GithubUser}}/{{Name}}" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"/></a>{{ if eq License "Apache Software License 2.0" }}<a href="./LICENSE" target="_blank"><img src="https://img.shields.io/badge/license-Apache%202.0-blue.svg?style=flat-square"/></a>{{ else if eq License "MIT" }}<a href="./LICENSE" target="_blank"><img src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square"/></a>{{ else if eq License "GNU GPL v3.0" }}<a href="./LICENSE" target="_blank"><img src="https://img.shields.io/badge/license-GPL-blue?style=flat-square"/></a>{{ end }}
</p>

{{ if TravisCI }}
In travis set GITHUB_TOKEN
{{ if Codecov }}In travis set CODECOV_TOKEN{{ end }}
{{ if DockerHub }}In travis set DOCKER_PASSWORD{{ end }}
{{ end }}

## Table of contents

