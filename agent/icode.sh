# $1 user_id
# $2 host_port

IMAGE_NAME="codercom/code-server"
ROOT_PATH="/root/.eCloud"
CONTAINER_IDS_FILE="i_code_ids"
CONTAINER_NAME=icode-"${1}"-"${2}"

cd ${ROOT_PATH} || mkdir -p ${ROOT_PATH} > /dev/null
docker pull ${IMAGE_NAME} > /dev/null
docker run -d --name "${CONTAINER_NAME}" -p "${2}":8080 ${IMAGE_NAME} >> ${CONTAINER_IDS_FILE}
sleep 1
# 输出 container_id
tail -n 1 ${CONTAINER_IDS_FILE}
# 输出 password
docker exec -i "${CONTAINER_NAME}" cat /home/coder/.config/code-server/config.yaml | grep "password:" | awk '{print $2}'
# 输出 container_ip
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "${CONTAINER_NAME}"
