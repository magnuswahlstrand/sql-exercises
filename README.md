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