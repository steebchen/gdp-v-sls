# Go client Prisma Data Proxy example

Install:

```bash
git clone git@github.com:steebchen/gdp-v-sls.git
cd gdp-v-sls
```

You should use the url shortener application in order to get some data returned with this example app.
Create a new data proxy connection string and set it via `DATA_PROXY_DATABASE_URL`:

```bash
export DATA_PROXY_DATABASE_URL="prisma://....."
```

Fetch dependencies and generate the Prisma Go client using go tools:

```bash
go generate ./...
```

Run the app:

```bash
go run .
```
