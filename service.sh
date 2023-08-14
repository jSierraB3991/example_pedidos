sudo docker  run --rm -d -p 5432:5432 --name postgre_zabud  -v /home/lelouch/Source/data/postgre_zabud:/var/lib/postgresql/data  -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -e POSTGRES_PASSWORD=root postgres:14.6-alpine

sudo docker run --rm -d -p 27017:27017  -v /home/lelouch/Source/data/mongo_inscription:/data/db --name mongo-inscription mongo:5.0.19-focal