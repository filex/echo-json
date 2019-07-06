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