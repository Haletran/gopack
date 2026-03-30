describe({
    name = "yt-dlp",
    version = "2026.03.17",
    url = "https://github.com/yt-dlp/yt-dlp/releases/download/2026.03.17/yt-dlp_linux",
    sha256 = "sha256:c2b0189f581fe4a2ddd41954f1bcb7d327db04b07ed0dea97e4f1b3e09b5dd8e",
})

Download(url)
Install("/tmp/yt-dlp_linux", "/usr/local/bin/yt-dlp")
