#/bin/bash
docker inspect $1 &> /dev/null
if [[ $? == 0 ]]; then
    echo "$1 existed"
    continue
else
    echo "$1 not existed"
    docker pull $2
    docker tag  $2 $1
    docker rmi $2
fi


