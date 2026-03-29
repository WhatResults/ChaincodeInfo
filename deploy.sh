docker rm -f $(docker ps -aq)
docker network prune
docker volume prune
cd fixtures
docker-compose up -d
cd ..
rm ChaincodeInfo.exe
go build
./ChaincodeInfo.exe