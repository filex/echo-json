echo-json
==================

`echo-json` is a tiny go cli utility that accepts name/value pairs as command
line arguments and prints them as a JSON object.


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

### Nested Objects

There is no `object` type for arguments. However, as `echo-json` is intended to be a command line tool, we can use a sub shell with a `raw` type to create a JSON with nested objects:

```
$ echo-json type request timing:raw "$( echo-json ttfb:float 0.023 total:float 1.04)"
{"timing":{"total":1.04,"ttfb":0.023},"type":"request"}
```
