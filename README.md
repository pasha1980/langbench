# LangBench - JFF benchmark of some programming languages

Just For Fun :)

## Background

Everyone was constantly trying to convince me of certain stereotypes about the performance and speed of programming languages. At some point, I decided everyone was wrong and decided to test how *REALLY* true this was, as was commonly believed.

## How does it work?

The `competitors`` directory contains simple API implementations based on the requirements (the requirements are described below).

I tried (and I encourage anyone wanted to contribuet) to implement them optimally and at the same time in a way that a typical developer would implement them on a typical project (without ASM, following common frameworks, without `php://input`, etc.).

Each `competitor` should also provide valid Dockerfile in root directory.

### API requirements

API should contain route POST /api/calc with following input

``` json
{
    "request_id": "some_string",
    "items": [
        {
            "id": "unique_id",
            "value": 3.14,
            "tags": ["foo", "bar"]
        },
        ...
    ]
}
```

Fields:
- `request_id`: random string, in test generates andomly to prevent caching
- `items.id`: random string. Must be unique withing items list
- `items.value`: any float value. Can overflow the limits. In this case, consider it's equal 0
- `items.tags`: list of random strings

``` json
{
    "request_id": "some_string",
    "checksum": "M2ZiNWI2ZDY2YjMzODY1YTliZDE4YTg0MTgxYjNjY2RlM2I1ZjA1MWQxZTk5ODc3ZTZkN2ZkNjYxOWIyYWVmNg==",
    "stats": {
        "foo": {
            "count": 1,
            "sum": 3.14
        },
        "bar": {
            "count": 1,
            "sum": 3.14
        },
        ...
    }
}
```

Fields:
- `request_id`: the same string, provided in input
- `checksum`: sha256 hash of json with only `request_id` and `stats`
- `stats.*`: each tag mentioned in `items.tags`
- `stats.*.count`: how many times that particular tag was mentioned in input
- `stats.*.sum`: summary of values in items containing this particular tag

## Contribution

If you found some bottleneck, don't agree with testing approach, etc - feel free to contribute :)

## TODOs

- [] Workflow for automatic comparison
- [] Add JS-based API
- [] who-knows
