for i in $(which -a docker); do
    if [ -e "$i" ]
    then
      export DOCKER_PATH="$i"
      break    
    fi
done

docker_entry() {    
     case "$1" in
     "push" )
        shift 1
        echo oras push
        echo docker push $@
        ;;
    "pull" )
        shift 1
        echo verify signature
        echo docker pull $@
        ;;
    *)     
        $DOCKER_PATH $@
        ;;
    esac
}

alias docker='docker_entry $@'
