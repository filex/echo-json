echo-json
==================

`echo-json` is a tiny go cli utility that accepts name/value pairs as command
line arguments and prints them as a JSON object.

[![Build Status](https://travis-ci.com/filex/echo-json.svg?branch=master)](https://travis-ci.com/filex/echo-json)
[![Docker Image](https://img.shields.io/docker/pulls/filex/echo-json)](https://hub.docker.com/r/filex/echo-json)

```
$ echo-json foo bar baz '"quux"'
{"baz":"\"quux\"","foo":"bar"}
```

While this seems silly, it can be very helpful to output JSON formatted log
messages in cli tools.

Typed Values
-------------------

Besides strings, `echo-json` can create typed values such as floats,
integers or booleans. Append a _type hint_ to the key arg to set the value's type:

```
$ echo-json age:int 33 score:float 1.234 active:bool true
{"active":true,"age":33,"score":1.234}
```

Supported type hints:

* `int`
* `float`
* `bool` (accepts `true`, `false`, `1` and `0`)
* `string` (default)
* `raw` (accepts a literal JSON string)

You can still use "namespaced" keys with colons. The last such part is
considered the type:

```
$ echo-json u:age:int 33
{"u:age":33}
```

If the type hint does not match one of the types listed above, it is left
as part of the key name:

```
$ echo-json u:name alice u:role admin d:timestamp 1562520403
{"d:timestamp":"1562520403","u:name":"alice","u:role":"admin"}
```

### Nested Objects

There is no `object` type for arguments. However, as `echo-json` is intended to be a command line tool, we can use a sub shell with a `raw` type to create a JSON with nested objects:

```
$ echo-json type request timing:raw "$( echo-json ttfb:float 0.023 total:float 1.04)"
{"timing":{"total":1.04,"ttfb":0.023},"type":"request"}
```


## Default Values

A value is missing if

* it is the empty string `""` or
* an uneven number of arguments was given.

If a field's value is missing, the default value for the corresponding type is used.

This avoids failure if the command line has to rely on optional data, such as variables or sub shells:

```shell
$ echo-json id:int "$VAR"
{"id":0}
```

If `$VAR` is not set value argument will expand to `""`. While this is not an `int`, the default value of `0` is used instead of bailing out with an error.

Note that the variable must be quoted: `"$VAR"`. Without quotes, `echo-json` would not see the variable at all if it is not set. (Then a following _key_ would be used as _value_.)

Default Values:

* `int` →`0`
* `float` → `0.0`
* `bool` → `false`
* `raw` → `null`
