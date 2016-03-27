# jenkins_list_plugins

A command-line tool to list Jenkins plugins in `name:version` format.

The primary use of this tool is to migrate from manual to automatic setup of
Jenkins masters. The Jenkins Docker image [accepts a plugins.txt](https://github.com/jenkinsci/docker/blob/master/README.md#installing-more-tools) file to bootstrap plugin installation.
This tool generates a "plugins.txt"-style list for your existing Jenkins install.

Inactive plugins are red, outdated plugins are yellow, updated plugins are green.

![Screenshot](https://i.imgur.com/nj0NO95.png)

## Usage

Install by cloning this project and running `go build ./...`.

Then run it, pointing to your Jenkins URL:

```
./jenkins_list_plugins https://jenkins.myorg.com
```

```
NAME:
   jenkins_list_plugins - List Jenkins plugins in shortname:version format

USAGE:
   jenkins_list_plugins [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --username, -u   Username for Jenkins authentication [$JENKINS_USERNAME]
   --password, -p   Password for Jenkins authentication [$JENKINS_PASSWORD]
   --insecure, -k Allow connections to SSL sites without certs [$SKIP_TLSVERIFY]
   --help, -h   show help
   --version, -v  print the version
```
