@FOR /F "tokens=*" %%a IN ('FINDSTR /V /B "#" .env') DO SET "%%a"
air --build.cmd "go build -o build\main.exe ." --build.bin ".\build\main.exe" -tmp_dir ".\build\"