$VERSION = "1.1.0"

# Recreate dist directory
Remove-Item -Force -Recurse -Path dist -ErrorAction Ignore
New-Item -ItemType Directory -Force -Path dist

# Generate Windows binary
$Env:GOOS = "windows"; $Env:GOARCH = "amd64"
go build

# Zip Windows binary
tar.exe -acf "dist/SEQ-Search-$VERSION-Windows.zip" seq_search.exe

# Generate Mac binary
$Env:GOOS = "darwin"; $Env:GOARCH = "amd64"
go build

# Zip Mac binary
tar.exe -acf "dist/SEQ-Search-$VERSION-Mac.zip" seq_search

# Generate Linux binary
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"
go build

# Zip Linux binary
tar.exe -acf "dist/SEQ-Search-$VERSION-Linux.zip" seq_search
