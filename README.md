## About

This repository provides a GitHub action for running builds and tests
on a [hardenedBSD](https://hardenedbsd.org) virtual machine. It is
inspired by the
[vmactions](https://github.com/vmactions)
project that provides a similar service for the mainstream BSD operating
systems (FreeBSD, OpenBSD, NetBSD, etc). Their work inspired me and it was
adapted for hardenedBSD.

## Usage

#### Workflow

The following is an example GitHub workflow that uses this action to run
tests on a hardenedBSD virtual machine. It checks out the code, boots the
VM, installs the Go programming language, and then runs `make test` on the
virtual machine:

```yaml
name: My workflow
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
  workflow_dispatch:

jobs:
  test:
    name: Build
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Run test
      uses: 0x1eef/hardenedbsd-vm@v1
      with:
        release: '15-STABLE'
        run: |
          pkg-static install -y go
          make test
```

#### Inputs

All GitHub actions accept inputs via the "with" directive. This
action provides a couple of input variables that can be used this
way. In the future, more variables may be supported. Certain variables,
like the CPU architecture and filesystem type are always amd64 and ufs
respectively but might be configurable in the future.

* release<br>
The hardenedBSD release to use. <br>
This can be `16-CURRENT` or `15-STABLE`.
* run<br>
The command to run on the hardenedBSD virtual machine. <br>
This can be any valid shell command(s).

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/hardenedbsd-vm)
* [git.hardenedBSD.org/@0x1eef](https://git.hardenedBSD.org/0x1eef/hardenedbsd-vm)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
