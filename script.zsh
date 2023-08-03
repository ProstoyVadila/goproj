

# Golang project init
goproj() {
    if [ ! -n "$1" ]; then
        echo "Enter a project name"
    else
        mkcd $1
        go mod init github.com/ProstoyVadila/$1
        git init
        touch README.md LICENSE Makefile .gitignore Dockerfile .dockerignore .env
        mkdir config cmd/app pkg internal
    fi
}