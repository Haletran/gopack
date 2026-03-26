describe({
    name    = "ffmpeg",
    version = "8.1",
    url     =
    "https://github.com/BtbN/FFmpeg-Builds/releases/download/autobuild-2026-03-26-13-16/ffmpeg-N-123625-gfd9f1e9c52-linux64-gpl-shared.tar.xz",
    sha256  = "sha256:69f39cec93742c7dc7be26b641e2c814a22ce31a0a92d4d8c919ace9601be7f2",
})

Download(url)
Extract("/tmp/ffmpeg-N-123625-gfd9f1e9c52-linux64-gpl-shared.tar.xz", "/tmp/")
