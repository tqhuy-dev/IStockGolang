function build() {
    docker build . -t "tqhuy1996developer/istock"
    docker push "tqhuy1996developer/istock"
}

build