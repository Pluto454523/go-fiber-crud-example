# **GO-FIBER-CRUD-EXAMPLE**

## Install and Usage

### ☁️ Air - Live reload for Go apps.
First, ensure you have Air installed. You can install it using:
```bash
go install github.com/cosmtrek/air@latest
```

### Windows:
- Create a batch script (let's name it run.bat) with the following content:
```bash
## Read .env file and set envoriments.
@FOR /F "tokens=*" %%a IN ('FINDSTR /V /B "#" .env') DO SET "%%a"

## Run air and argurment without air config.
air --build.cmd "go build -o build\main.exe ." --build.bin ".\build\main.exe" -tmp_dir ".\build\"
```
- Run the script
```bash
./run.bat
```


### Unix:
- Create a bash script (let's name it run.sh) with the following content:
```bash
## Read .env file and set envoriments.
export $(grep -v '^#' .env | xargs);

## Run air and argurment without air config.
air --build.cmd "go build -o build/main ." --build.bin "./build/main" -tmp_dir "./build"
```

- Make the script executable:
```bash
chmod +x run.sh
```

- Run the script
```bash
./run.sh
```