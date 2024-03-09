# **GO-FIBER-CRUD-EXAMPLE**

## Install and Usage Air (How reload)

### ☁️ Air - Live reload for Go apps.
1. Install `Air - Live reload` with Go installer
```bash
go install github.com/cosmtrek/air@latest
```
2. Initialize `.air.toml` file for configuration
```bash
air init
```

## Setup ☁️ Air configuration.


### On **`Windows`**
1. Create a batch script (run.bat) with the following content
```bash
@FOR /F "tokens=*" %%a IN ('FINDSTR /V /B "#" .env') DO SET "%%a"
```

2. Execute **`run.bat`** with modify `.air.toml` file
```toml
bin = "tmp\\run && tmp\\main.exe"
```

3. Run air live reload
```bash
air
```


### On **`Linux`** or **`Mac OS X`**
1. Add **`export $(grep -v '^#' .env | xargs);`** into bin  in`.air.toml`
```toml
bin = ";export $(grep -v '^#' .env | xargs); ./tmp/main"
```

2. Run air live reload
```bash
air
```