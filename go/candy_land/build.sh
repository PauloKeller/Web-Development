#! /bin/sh

build_service() {
  sub_folder="../"
  echo "########################################"
  echo "#########   BUILDING SERVICES   ########"
  echo "########################################"
  for i in "addresses_service" "users_service" "carts_service" "orders_service"
  do
    cd $i
    echo "Building service $i..."
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/main .
    echo "Building done!"
    cd $sub_folder
  done
}

run_docker_compose() {
  echo "########################################"
  echo "########   RUN DOCKER COMPOSE   ########"
  echo "########################################"
  docker-compose up --build
  echo "Runing done!"
}

build_service
run_docker_compose
