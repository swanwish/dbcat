# dbcat

This is a db client, currently support sqlite3.

## Connect to sqlite database

```shell
dbcat shell -dbPath=<sqlite db path>

# or ignore shell
dbcat -dbPath=<sqlite db path>
```

## Show tables

```shell
.tables

# Filter table with pattern
.tables <pattern>
```

## Show schemas

```shell
.schema

.schema <pattern>
```

## Show columns

```shell
.columns

.columns <table name>
```

## Query data

```shell
.query <query sql>

# or just type query sql
<query sql>
```

## Example

```
$ dbcat -dbPath=test.db
dbcat> create table tbl1(one text, two int);
dbcat> insert into tbl1 values('hello!',10);
dbcat> insert into tbl1 values('goodbye', 20);
dbcat> select * from tbl1;
one             two       
----------      ----------
hello!          10        
goodbye         20        
dbcat> CREATE TABLE tbl2 (
  ...>   f1 varchar(30) primary key,
  ...>   f2 text,
  ...>   f3 real
  ...> );
dbcat> .tables
name      
----------
tbl1      
tbl2      
dbcat> .schema
CREATE TABLE tbl1(one text, two int)

CREATE TABLE tbl2 ( f1 varchar(30) primary key, f2 text, f3 real )

dbcat> .schema tbl1
CREATE TABLE tbl1(one text, two int)

dbcat> .columns tbl2
cid             name            type            notnull         dflt_value      pk        
----------      ----------      ----------      ----------      ----------      ----------
0               f1              varchar(30)     0               <nil>           1         
1               f2              TEXT            0               <nil>           0         
2               f3              REAL            0               <nil>           0         
dbcat> .q
```