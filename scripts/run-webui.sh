#IMAGE=gebug-ui
#PORT=3030
#PROJECT_LOCATION=/Users/moshe/Dev/cpp-gebug
#docker run -p $PORT:$PORT -e PORT=$PORT -e VUE_APP_GEBUG_PROJECT_LOCATION=$PROJECT_LOCATION -v $PROJECT_LOCATION:$PROJECT_LOCATION -it $IMAGE

docker-compose -f .gebug/webui-docker-compose.yml up