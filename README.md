Spell manager for a new TTRPG.

For now, this will be a CLI app to help work with spells. We'll extend it into other delivery methods and game areas later.

# Build Notes

Since we're using SQLite, we need a C compiler for the Go compiler to call when building the project.

I'm cross-compiling this for macOS and Windows right now.

Since I'm developing on a Mac, `go build` works fine out of the box since I have gcc installed locally already.

For Windows, I installed MinGW through Homebrew and modified the build command to

```
CGO_ENABLED=1 CC="/usr/local/bin/x86_64-w64-mingw32-gcc" CXX="/usr/local/bin/x86_64-w64-mingw32-g++" CC_FOR_TARGET="/usr/local/bin/x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -o wws.exe
```

I haven't tested which of the extra options are required. CGO_ENABLED and CC are. I'll test the others out when I get a sec.

