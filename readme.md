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
## Read .env file and set envoriments.
@FOR /F "tokens=*" %%a IN ('FINDSTR /V /B "#" .env') DO SET "%%a"
```

2. Add **`./run.bat && ./build/main.exe`** unti bun un `.air.toml`
```toml
bin = "./run.bat && ./build/main.exe"
```

3. Run air live reload
```bash
air
```


### On **`Linux`** or **`Mac OS X`**
1. Add **`export $(grep -v '^#' .env | xargs);`** into bin  in`.air.toml`
```toml
bin = ";export $(grep -v '^#' .env | xargs); ./build/main"
```

2. Run air live reload
```bash
air
```