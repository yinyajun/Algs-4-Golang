# prepend `GOPATH` and wrap `go run` command
DIR="$( cd "$( dirname "$0"  )" && pwd  )"

GOPATH=`go env GOPATH`
export GOPATH=${DIR}:${GOPATH}

GOROOT=`go env GOROOT`
$GOROOT/bin/go run $*
echo "== Current GOPATH:${GOPATH} =="
