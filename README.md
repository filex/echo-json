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
integers or booleans. Use a _type hint_ in the key arg to define a type:

```
$ echo-json int:age 33 float:score 1.234 bool:active true
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
$ echo-json type request raw:timing "$( echo-json float:ttfb 0.023 float:total 1.04)"
{"timing":{"total":1.04,"ttfb":0.023},"type":"request"}
``` 


