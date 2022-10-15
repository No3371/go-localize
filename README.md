# go-localize

A fork of [go-localize](github.com/m1/go-localize).

## Differences

- The generated string value is wrapped with backticks(`) instead of quotes(").
  - =>Multi-line content support
- Global Keys: Source files named with "_" (ex: _.yaml) will not be generated with path key. For example, keys of entries in `en/a/b.yaml` will be "en.a.b.key", but the ones in `en/a/_.yaml` will be "en.key".