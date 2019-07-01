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
