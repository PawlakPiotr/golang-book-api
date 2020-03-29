BASEPATH=`pwd`
FILE=book_api.out

export GOPATH=$HOME/go:${BASEPATH}

cd src/golang-book-api

if [[ -f "$FILE" ]]; then
    rm book_api.out
fi

go build -o book_api.out
./book_api.out