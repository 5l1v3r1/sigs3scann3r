# s3

![made with go](https://img.shields.io/badge/made%20with-Go-0040ff.svg) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/drsigned/s3.svg?style=flat&color=0040ff)](https://github.com/drsigned/s3/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/drsigned/s3.svg?style=flat&color=0040ff)](https://github.com/drsigned/s3/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/license-MIT-gray.svg?colorB=0040FF)](https://github.com/drsigned/s3/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@drsigned-0040ff.svg)](https://twitter.com/drsigned)

## Resources

* [Installation](#installation)
    * [From Binary](#from-binary)
    * [From source](#from-source)
    * [From github](#from-github)
* [Usage](#usage)
* [Credits](#credits)
* [Contribution](#contribution)

## Installation

#### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/drsigned/s3/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

#### From Source

s3 requires **go1.14+** to install successfully. Run the following command to get the repo

```bash
$ GO111MODULE=on go get -u -v github.com/drsigned/s3/cmd/s3
```

#### From Github

```bash
$ git clone https://github.com/drsigned/s3.git; cd s3/cmd/s3/; go build; mv s3 /usr/local/bin/; s3 -h
```

## Usage

To display help message for s3 use the `-h` flag:

```
$ s3 -h
     _____
 ___|___ /
/ __| |_ \
\__ \___) |
|___/____/ v1.0.0

USAGE:
  s3 [MODE] [OPTIONS]

MODES:
  format            convert various s3 buckets format (see formats below)

FORMAT OPTIONS (used with mode 'format'):
  path              to path-style (e.g. https://s3.amazonaws.com/github-cloud)
  name              to bucket name (e.g. github-cloud)
  url               to s3 url (e.g. s3://github-cloud)
  vhost             to virtual-hosted-style (e.g. github-cloud.s3.amazonaws.com)
```

## Credits

All credits to [HAHWUL](https://github.com/hahwul), I took the initial code from his [s3reverse](https://github.com/hahwul/s3reverse).

## Contribution

[Issues](https://github.com/drsigned/s3/issues) and [Pull Requests](https://github.com/drsigned/s3/pulls) are welcome!