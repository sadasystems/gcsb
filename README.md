# GCSB

It's like YCSB but with more Google.

- [GCSB](#gcsb)
  - [Build](#build)
  - [Test](#test)
  - [Load](#load)
    - [Single table load](#single-table-load)
  
## Build

```sh
make build
```

## Test

```sh
make test
```

## Load

### Single table load

By default, GCSB will detect the table schema and create default random data generators based on the columns it finds. In order to tune the values the generator creates, you must create override configurations in the gcsb.yaml file. Please see that file's documentation for more information.

```sh
gcsb load -t TABLE_NAME -o NUM_ROWS
```

Additionally, please see `gcsb load --help` for additional configuration options.
