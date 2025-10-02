# jsonw

[![Go Reference](https://pkg.go.dev/badge/github.com/asgari-hamid/jsonw.svg)](https://pkg.go.dev/github.com/asgari-hamid/jsonw)
[![Go Report Card](https://goreportcard.com/badge/github.com/asgari-hamid/jsonw)](https://goreportcard.com/report/github.com/asgari-hamid/jsonw)
![CI](https://github.com/asgari-hamid/jsonw/actions/workflows/jsonw.yml/badge.svg)

A small Go library that provides **low-level JSON writing** with a clean API.  
It is built on top of [`easyjson/jwriter`](https://pkg.go.dev/github.com/mailru/easyjson/jwriter) for high performance and zero allocations when writing JSON.

Unlike `encoding/json`, this package lets you **manually control JSON output**, which is useful when:
- You want to respect a field mask provided by a client.
- You need fine-grained control over escaping, formatting, and ordering.
- You care about performance and memory allocations.

---

## 📦 Installation

```bash
go get github.com/asgari-hamid/jsonw
