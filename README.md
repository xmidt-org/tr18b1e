# TR-181 library

[![Build Status](https://travis-ci.com/xmidt-org/tr18b1e.svg?branch=master)](https://travis-ci.com/xmidt-org/tr18b1e)
[![codecov.io](http://codecov.io/github/xmidt-org/tr18b1e/coverage.svg?branch=master)](http://codecov.io/github/xmidt-org/tr18b1e?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/xmidt-org/tr18b1e)](https://goreportcard.com/report/github.com/xmidt-org/tr18b1e)

Simple CRUD interface that sits on a "box" next to the `kratos` websocket wrapper that we can use
to answer queries about attributes of the box. For now we have the "honest" part of the interface,
found in tr18b1e.go, and we also have the "fake" part of the interface which mimics expected behavior
to satisfy tests and the like.

