# jsonw

A lightweight Go library providing a higher-level wrapper around [EasyJSON](https://github.com/mailru/easyjson)'s `jwriter.Writer` for manual, low-level JSON generation.  

Supports nested objects and arrays, all JSON value types, and proper string escaping. Ideal for building JSON payloads dynamically based on field masks or custom serialization logic.

## Features

- Handles all JSON value types:  
  - `string`, `number`, `integer`, `float`, `boolean`, `null`  
- Nested arrays and objects  
- Automatic comma management  
- Proper JSON escaping, including special characters and Unicode line separators  
- Low-level, efficient writing with minimal allocations using EasyJSON buffers  

## Installation

```bash
go get github.com/asgari-hamid/jsonw
