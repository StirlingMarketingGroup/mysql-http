# MySQL HTTP

A small MySQL UDF library for making HTTP requests written in Golang.

## Usage

### `http_touch`

Makes an HTTP Get request to the given URL and returns nothing.

```sql
`http_touch` ( `URL` )
```

 - `` `URL` ``
   - The URL to request.

## Examples

We can quickly build a test table of random numbers by running these queries

```sql
-- hit our test endpoint
select`http_touch`('http://localhost:48642/');
```


## Dependencies

You will need Golang, which you can get from here https://golang.org/doc/install.

Debian / Ubuntu

```shell
sudo apt update
sudo apt install libmysqlclient-dev
```

## Installing

You can find your MySQL plugin directory by running this MySQL query

```sql
select @@plugin_dir;
```

then replace `/usr/lib/mysql/plugin` below with your MySQL plugin directory.

```shell
cd ~ # or wherever you store your git projects
git clone https://github.com/StirlingMarketingGroup/mysql-http.git
cd mysql-http
go build -buildmode=c-shared -o mysql_http.so
sudo cp mysql_http.so /usr/lib/mysql/plugin/mysql_http.so # replace plugin dir here if needed
```

Enable the function in MySQL by running this MySQL query

```sql
create function`http_touch`returns int soname'mysql_http.so';
```