set -x
PREFIX="eCloud.agent."
LINUX_X86_64="linux-amd64"
DARWIN_X86_64="darwin-amd64"

SERVER_ROOT_URL=$1
CTR_URL=$2
ROOT_PATH="/root/.eCloud/"
NGINX_PREFIX="nginx-1.20.0"
NGINX="nginx-1.20.0.tar.gz"
LOGSTASH_PREFIX="logstash-7.12.0"
LOGSTASH="logstash-7.12.0-linux-x86_64.tar.gz"

SYSNAME=$(uname -s)
SYSLONG=$(uname -m)

function downloadAgent() {
  curl "${SERVER_ROOT_URL}${PREFIX}${1}" --output "agent"
  chmod +x "agent"
}

# 下载（更新）agent
if [[ ${SYSNAME} == "Linux" ]]; then
    if [[ ${SYSLONG} == "x86_64" ]]; then
      downloadAgent ${LINUX_X86_64}
    fi
fi

if [[ ${SYSNAME} == "Darwin" ]]; then
    if [[ ${SYSLONG} == "x86_64" ]]; then
      downloadAgent ${DARWIN_X86_64}
    fi
fi

# 删除 nginx/ logstash/
rm -rf "${ROOT_PATH}/nginx"
rm -rf "${ROOT_PATH}/logstash"

# 下载 nginx，nginx.conf
# 编译 && 启动
mkdir "${ROOT_PATH}/nginx/" -p
cd "${ROOT_PATH}/nginx" || exit
yum install -y pcre pcre-devel make zlib zlib-devel openssl openssl-devel
curl "${SERVER_ROOT_URL}/${NGINX}" --output "${NGINX}"
mkdir -p "${ROOT_PATH}/nginx/conf/conf.d/"
mkdir -p "${ROOT_PATH}/nginx/logs/"
tar xfv "${NGINX}"
cd "${ROOT_PATH}/nginx/${NGINX_PREFIX}" || exit
./configure && make && make install
cp -r "${ROOT_PATH}/nginx/${NGINX_PREFIX}/conf" "${ROOT_PATH}/nginx/"
cd "${ROOT_PATH}/nginx/conf" || exit
curl "${SERVER_ROOT_URL}/nginx.conf" --output "nginx.conf"
/usr/local/nginx/sbin/nginx -s stop
/usr/local/nginx/sbin/nginx -c "${ROOT_PATH}/nginx/conf/nginx.conf"

# kill 掉 logstash
# 不 kill 会出现 Unable to run: daemon: Resource temporarily unavailable
kill -9 `cat ${ROOT_PATH}/logstash.pid`
ps x | grep logstash | awk '{print $1}' | xargs kill -9
# 下载 logstash，logstash.conf
# 使用 ./agent logstash 进行启动
mkdir ${ROOT_PATH}/logstash -p
cd "${ROOT_PATH}/logstash" || exit
curl "${SERVER_ROOT_URL}/${LOGSTASH}" --output "${LOGSTASH}"
tar xvf "${LOGSTASH}"
echo "
input {
	file {
		path => \"/root/.eCloud/nginx/logs/access_json.log\"
			codec => json
			start_position => \"beginning\"
			type => \"nginxlog\"
	}
}
output {
	stdout { codec => rubydebug }
	http {
		http_method => \"post\"
		url => \"${CTR_URL}/router\"
		format => \"json\"
		mapping => { \"uri\" => \"%{[request]}\" \"un\" => \"%{[un]}\" }
	}
}
" > "${ROOT_PATH}/logstash/logstash.conf"
cd "${ROOT_PATH}" && ./agent logstash
