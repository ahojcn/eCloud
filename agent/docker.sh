function downloadDocker() {
    curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
}

downloadDocker
