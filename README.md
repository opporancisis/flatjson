# flatjson

A simple command that flattens input json.

### How to run

```shell
$ cat test.json | go run main.go
```

### Flattening algorithm

Consider the following input object:

```json
{
    "a": 1,
    "b": true,
    "c": {
        "d": 3,
        "e": "test"
    }
}
```

In this example the path to the terminal value 1 is "a" and the path to the
terminal value 3 is "c.d".

The program will output the flattened version:

```json
{
    "a": 1,
    "b": true,
    "c.d": 3,
    "c.e": "test"
}
```