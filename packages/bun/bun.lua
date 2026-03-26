describe({
    name = "bun",
    version = "1.3.11",
    url = "https://github.com/oven-sh/bun/releases/download/bun-v1.3.11/bun-linux-x64.zip",
    sha256 = "sha256:8611ba935af886f05a6f38740a15160326c15e5d5d07adef966130b4493607ed",
})

Download(url, sha256)
Extract("/tmp/bun-linux-x64.zip", "/tmp/")
Install("/tmp/bun-linux-x64/bun", "/usr/local/bin/bun")
