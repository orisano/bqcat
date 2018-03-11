# bqcat
bqcat is Google BigQuery client for a command line.

query result writes to stdout in CSV. pipe friendly.

## installation
```bash
go get -u github.com/orisano/bqcat
```

## usage
```bash
$ bqcat --help
Usage of bqcat:
  -f string
        query file path
  -p string
        project id
```

## how to use
### give query from arguments
```bash
$ bqcat "SELECT 1;"
1
```
### give query file path from arguments
```bash
$ bqcat -f test.sql
1
```

### give query from stdin
```bash
$ echo "SELECT 1;" | bqcat
1
```

## LICENSE
MIT
