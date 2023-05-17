## FOR DATABASE
### Docker
    - cd ./sql/
    - docker volume create dbtuto
    - docker-compose up -d 
#### To List Data From DB
    - docker exec -it pg-container bash 
    - psql -h 0.0.0.0 -p 5432 -U test -W golang //password:12345
### Goose
    - goose postgres postgres://test:12345@0.0.0.0:5432/golang up    
    - goose postgres postgres://test:12345@0.0.0.0:5432 golang status
    - goose postgres postgres://test:12345@0.0.0.0:5432/golang down 
### SQLC
    -sqlc generate    