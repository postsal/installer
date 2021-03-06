#/bin/bash
file="deploy/images"

if [ -f "$file" ]
then
  echo "$file found."

  while IFS='=' read -r key value
  do
    echo "${key}=${value}"
    docker inspect ${key} &> /dev/null
    if [[ $? == 0 ]]; then
        echo "${key} existed"
        continue
    else
        echo "${key} not existed"
        docker pull ${value}
        docker tag ${value} ${key}
        docker rmi ${value}
    fi
  done < "$file"

else
  echo "$file not found."
fi
