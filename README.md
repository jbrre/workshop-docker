# Workshop: Docker

## Introduction
Dans ce Workshop, nous allons nous initier à Docker par la création de plusieurs [services](https://docs.docker.com/engine/reference/commandline/service/). Pour ceci, nous créerons un container basique sur Docker, puis nous la ferons communiquer avec un autre container contenant une base de données, et enfin avec une autre API contenue dans un autre container.

## Mise en place
Pour ce workshop, il vous faudra installer Go et Docker.
[Installer Go](https://go.dev/doc/install)
[Installer Docker](https://docs.docker.com/get-docker/)

## Etape 1: Mise en place de Docker
Clonez le repo de ce Workshop.
La première étape va être de créer un [Dockerfile](https://docs.docker.com/engine/reference/builder/) qui fait compiler "server".
Pour votre information, la commande afin de lancer server est `go run server.go`.

Une fois le Dockerfile créé, vous pouvez le tester de cette avec ces 3 commandes.

    docker build -t workshop-docker ./server/
    docker run -dp 8080:8080 workshop-docker
    curl localhost:8080
`docker build -t workshop-docker .` va build votre Dockerfile sous le nom de container `workshop-docker`.
`docker run -dp 8080:8080 workshop-docker` va lancer le container `workshop-docker` en fond (`-d`) et sur le port 8080:8080 (`-p`)
Enfin, `curl localhost:8080` va lancer une requête sur le port 8080, sur lequel une route a été pré-configurée afin de pouvoir tester une requête.
![Output attendu](https://i.imgur.com/Htb8tz8.png)

## Etape 2: docker-compose.yml
A la racine du repo, créez un fichier `docker-compose.yml`. Dans ce fichier, vous allez devoir créer un [docker-compose](https://docs.docker.com/compose/) qui crée un service `server` en se basant sur le Dockerfile contenu dans `./server`. Il exposera ensuite sur le port `8080:8080`.

Une fois votre docker-compose créé, vous pouvez le tester en lançant les commandes :

    docker-compose build
    docker-compose up -d
    curl localhost:8080
   
   ![Output attendu](https://i.imgur.com/HBblmHF.png)

## Etape 3 : Deux autres containers
Passons maintenant à des choses un peu plus compliquées !

Créez un service `db` qui aura pour [container_name](https://docs.docker.com/compose/compose-file/compose-file-v3/#container_name) `workshop_db`. Il se basera sur l'[image](https://docs.docker.com/compose/compose-file/compose-file-v3/#image) `mysql`. Il lancera également la [commande](https://docs.docker.com/compose/compose-file/compose-file-v3/#command) `--default-authentication-plugin=mysql_native_password` et [redémarrera](https://docs.docker.com/compose/compose-file/compose-file-v3/#restart) toujours jusqu'à ce qu'il se lance. Il exposera le [port](https://docs.docker.com/compose/compose-file/compose-file-v3/#ports) `3306:3306`. Il prendra en variable d'[environnement](https://docs.docker.com/compose/compose-file/compose-file-v3/#environment) `MYSQL_ROOT_PASSWORD: example`. Enfin, il [initialisera](https://docs.docker.com/compose/compose-file/compose-file-v3/#volumes) cette base de données avec le fichier `./server/mysql/init.sql`.

Afin de pouvoir tester ce service, nous créerons également un service `adminer`. Il se basera sur l'image `adminer` et redémarrera également toujours jusqu'à ce qu'il se lance. Il exposera le port `5050:8080` et [dépendra](https://docs.docker.com/compose/compose-file/compose-file-v3/#depends_on) du service `db`.

Pour tester tout ceci, lancez le docker-compose comme indiqué à l'étape précédente et rendez-vous sur `localhost:5050`. Connectez-vous avez `root` en username et `example` en password. Si tout s'est bien passé, vous devriez avoir une table `users` avec ces informations une fois sélectionnée :
![output attendu](https://i.imgur.com/G6Lg0bR.png)
**Attention:** Vu qu'il est très peu probable que vous réussissiez cette étape du premier coup, il faut que vous sachiez que même quand vous rebuildez un docker, si le volume a été déclaré en dehors du service (ce que vous devrez faire pour mysql), il restera de manière permanente. Pour afficher la liste de vos volumes, faites `docker volume ls`, et pour en supprimer un, `docker volume rm nomduvolume`. Si l'étape a été réussie, après un `docker volume rm`, vous pourrez rebuild votre mysql et elle aura les même informations que sur l'image ci-dessus.

## Etape 4 : Accès à la BDD depuis le service server
Pour accéder à la BDD depuis le service server, vous allez tout d'abord devoir dé-commenter les lignes de connexion à la BDD dans le fichier main.go
![lignes à décommenter](https://i.imgur.com/Psx75ij.png)

Une fois ceci fait, vous devrez passer les variables d'environnement suivantes dans `docker-compose.yml`
![variables à passer](https://imgur.com/3XLi0Kx.png)
DBUSER: `root`
DBPASS: `example`
DBADRESS: Adresse de la BDD, à vous de trouver !
DBNAME: `area_database`

Une fois ceci fait, vous pourrez faire une requête à `localhost:8080/user_list`, qui devrait vous retourner ceci :
![Output attendu](https://imgur.com/MIE0nsq.png)

## Etape 5: Votre propre service
Nous voici arrivés à l'étape finale ! Dans cette dernière étape, vous devrez créer votre propre service. Celui-ci devra communiquer avec server, récupérer le json de la liste des utilisateurs avec la route vue ci-dessus et la printer dans un format plus joli que celui vu précédemment.

Vous pouvez utiliser le langage de votre choix, mais le programme doit **impérativement** être exécuté dans un service Docker.

Exemple d'output attendu:
![exemple d'output attendu](https://imgur.com/Wep8NCS.png)