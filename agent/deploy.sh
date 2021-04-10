PREFIX="eCloud.agent."
LINUX_X86_64="linux-amd64"
LINUX_X86_32="linux-386"
LINUX_ARM="linux-arm"
DARWIN_X86_64="darwin-amd64"
DARWIN_X86_32="darwin-386"
WIN_X86_64="win-amd64.exe"
WIN_X86_32="win-386.exe"

HOSTID=$1
CTR_IP=$2
CTR_PORT=$3

SYSNAME=$(uname -s)
SYSLONG=$(uname -m)

SERVER_ROOT_URL="https://ahojcn.gitee.io/"
ROOT_PATH="/root/.eCloud/"
mkdir ${ROOT_PATH} -p
cd ${ROOT_PATH}

function downloadDksv() {
  echo "${SERVER_ROOT_URL}${PREFIX}${1}" --output "agent"
  curl "${SERVER_ROOT_URL}${PREFIX}${1}" --output "agent"
  chmod +x "agent"
  export ECLOUD_AGENT_HOSTID=$HOSTID
  export ECLOUD_CTR_IP=$CTR_IP
  export ECLOUD_CTR_PORT=$CTR_PORT
  "./agent"
  echo $0 $1 $2 $3
}

if [[ ${SYSNAME} == "Linux" ]]; then
    if [[ ${SYSLONG} == "x86_64" ]]; then
      downloadDksv ${LINUX_X86_64}
    fi
fi

if [[ ${SYSNAME} == "Darwin" ]]; then
    if [[ ${SYSLONG} == "x86_64" ]]; then
      downloadDksv ${DARWIN_X86_64}
    fi
fi