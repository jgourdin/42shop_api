# API Go et db Postgres
​
​Pour faire fonctionner le back il faut suivre ces quelques étapes.

-Il faut avoir installer Go et Postgres sur son ordinateur.
-Pour avoir Postgresql sur un container :
````
docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=12345 postgres
psql -h localhost -p 5432 -U postgres -W
```

Ensuite il faut initialiser 3 variable d'environement
````
export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD=12345
export APP_DB_NAME=postgres​
```

Pour lancer le serveur il faut rentrer deux commandes : 
````
go run main.go // Pour lancer le serveur.
go get github.com/jgourdin/42shop_api/model à chaque changement.

```

Il ne faut pas oublier de creer les tables users et products avant de commencer a dev.

Si vous rencontrer une erreur etrange utiliser cette commande faut d'avoir une meilleure solution :
```
open -n -a /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --args --user-data-dir="/tmp/chrome_dev_test" --disable-web-security
```