#/bin/bash

file="images"

if [ -f "$file" ]
then

  while IFS='=' read -r key value
  do
    # echo "${key}=${value}"
    docker inspect ${key} &> /dev/null
    if [[ $? == 0 ]]; then
        continue
    else
        docker pull ${value}
        docker tag ${value} ${key}
        docker rmi ${value}
    fi
  done < "$file"

else
  echo "$file not found."
fi
