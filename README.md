# react-catchup-server

## 開発環境起動
```
docker-compose -f ./docker-compose.dev.yaml up --build -d
```

## SQLBoiler
```
sqlboiler mysql -o ./src/infrastructure/dbmodels -c config/sqlboiler.toml
```