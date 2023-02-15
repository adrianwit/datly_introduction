## Managed mode



## Reader Service

### Use Cases

#### Invoice Service

```bash
    datly -C='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' -X invoice.sql
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/invoice    
```


#### Trader Service
```bash
    datly -C='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
     -C='dyndb|dynamodb|dynamodb://localhost:8000/us-west-1?key=dummy&secret=dummy' \
      -X trader.sql
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/trader    
```

#### Product Service

```bash
    datly -C='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
   -C='bqdev|bigquery|bigquery://viant-e2e/bqdev'  \
      -X product.sql
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/product    
```

#### Audience Service

```bash
    datly -C='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
      -X audience.sql
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/audience    
```


## Executor Service

