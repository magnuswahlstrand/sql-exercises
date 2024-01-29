# sql-exercises


### Features
* [x] Show checkmark on success/error on failure
* [x] Store success forever
* [ ] Store current query
* [ ] Add more exercises
* [ ] Add formatter
* [ ] Code highlighting


### Prepare for production
* [ ] Fix tailwind production setup
* [x] Add timeout for request
  * Set a generous 20 ms
* [x] Readonly DB
  * `no such table error` - https://github.com/mattn/go-sqlite3/blob/master/README.md#faq

### Deployment
* [ ] Deploy to production
* [ ] New domain name

### Other
* [ ] Close db connection on shutdown




# Query to end all queries
```sql
WITH RECURSIVE temp AS (
  SELECT 1 AS n
  UNION ALL
  SELECT n + 1 FROM temp
  WHERE n < 1000000
)
SELECT * FROM temp;
```


## Lessons learnt
- Apparently you can't have different access modes to an in-memory sqlite database. So if you want to have a readonly database, you need to create a file and then open it with `ro` mode. I had to fallback to using a transaction with `defer db.Rollback()` to make sure the database is not modified.
- When using an in-memory database, you need to keep the connection open, otherwise the database will be lost. This is because the database is stored in memory, and when the connection is closed, the memory is freed.
- sqlite3 needs to be compiled with `CGO_ENABLED=1` to work with go. This is because sqlite3 is written in C, and go needs to be able to call C functions.
  - It is not enabled by default. 
    - I found this blog post - https://adrianhesketh.com/2022/12/14/go-sqlite3-on-lambda/
    - And this issue SST issue - https://github.com/sst/sst/pull/3116

# Fix build
```
docker build -t go-lambda-app .
```

```
docker create --name temp-container go-lambda-app
docker cp temp-container:/app/bootstrap ./bootstrap
docker rm temp-container
```


...
#11 28.52 # runtime/cgo
#11 28.52 gcc: error: unrecognized command line option '-m64'
https://github.com/confluentinc/confluent-kafka-go/issues/898

