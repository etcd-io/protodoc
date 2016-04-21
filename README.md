
# protodoc

[![Build Status](https://img.shields.io/travis/coreos/protodoc.svg?style=flat-square)][cistat] [![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][protodoc-godoc]

protodoc generates Protocol Buffer documentation.

```
go get -v -u github.com/coreos/protodoc

protodoc --directory=./parse/testdata \
	--parse="service,message" \
	--languages="Go,C++,Java,Python" \
	--title=testdata \
	--output=sample.md

# to combine multiple directories into one
protodoc --directories=./parse/testdata=message,dirA=message_service \
	--languages="Go,C++,Java,Python" \
	--title=testdata \
	--output=sample.md
```

Note that parser only understands the minimum syntax
of Protocol Buffer (just enough to generate documentation).

For full featured parser, please check out https://github.com/golang/protobuf.

[cistat]: https://travis-ci.org/coreos/protodoc
[protodoc-godoc]: https://godoc.org/github.com/coreos/protodoc

