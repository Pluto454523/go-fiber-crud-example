export $(grep -v '^#' .env | xargs);
air --build.cmd "go build -o build/main ." --build.bin "./build/main" -tmp_dir "./build"