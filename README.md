# gRPC

Ce projet a pour objectif d'approfondir mes connaissances sur le **fonctionnement de gRPC**, en explorant ses concepts, l'implémentation de services, et la communication client-serveur.

## Objectifs

- Comprendre les bases de **gRPC** (Remote Procedure Call).
- Implémenter des services gRPC et des clients pour consommer ces services.
- Explorer les différents types de communication supportés par gRPC (Unary, Streaming, Bidirectional).
- Expérimenter la génération de stubs à partir de fichiers **.proto**.

## Technologies utilisées

- **Vue 3**
- **Golang**
- **gRPC** et Protocol Buffers (**.proto**)
- **Envoy**
- **Docker** et **Docker-compose**

## Instructions d'installation et d'exécution

### Cloner le dépôt

```bash
git clone https://github.com/rijenth/grpc.git
```

### Lancer les conteneurs
Ne pas oublier de créer un fichier **.env** à partir du **.env.example** avant d'executer la commande qui suit.
```bash
cd ./grpc && docker-compose up -d
```

### Migrer la base de donnée
```bash
docker-compose exec go migrate -path internal/infrastructure/database/migrations -database "mysql://root:root@tcp(mysql:3306)/database" up
```

## Outils

### 1. Protobuf

Le fichier `.proto` utilise un langage appelé **Protobuf**, qui permet de décrire les services et les messages (les données échangées).

- **Installation** :
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```

### 2. Protolint

Permet de **lint** les fichiers `.proto` pour détecter les erreurs de syntaxe et appliquer les bonnes pratiques.

- **Installation** :
  ```bash
  go install github.com/yoheimuta/protolint/cmd/protolint@latest
  ```

- **Usage** :
  ```bash
  protolint lint ./proto
  ```

### 3. gRPCurl

Permet de **simuler des requêtes** vers un serveur gRPC et d'interagir avec les services gRPC.

- **Installation** :
  ```bash
  brew install grpcurl
  ```

- **Usage** :
  
  1. **Lister les services disponibles sur un serveur gRPC** :
     
     ```bash
     grpcurl -plaintext localhost:8000 list
     ```

  2. **Décrire un service gRPC** :
     
     ```bash
     grpcurl -plaintext localhost:8000 describe user.UserService
     ```

  3. **Appeler une méthode gRPC avec des paramètres** :
     
     ```bash
     grpcurl -plaintext -d '{"id": "alice_d"}' localhost:8000 user.UserService/GetUserByUsername
     ```

  4. **Lister les méthodes disponibles d'un service gRPC** :
     
     ```bash
     grpcurl -plaintext localhost:8000 list user.UserService
     ```

  5. **Tester une méthode sans paramètres** :
     
     ```bash
     grpcurl -plaintext localhost:8000 user.UserService/ServiceSansParametre
     ```

  Ces commandes permettent de tester rapidement les services gRPC exposés par le serveur et de simuler des appels client-serveur.

### 4. Migration base de données (golang-migrate)

Permet de créer des fichiers de migration à l'aide du package `golang-migrate`.

- **Installation** :
  ```bash
  go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```

- **Usage** :

  1. **Création des fichiers de migration SQL** (`.up` et `.down`) :
     
     ```bash
     migrate create -ext sql -dir internal/infrastructure/database/migrations -seq <nom_de_la_migration>
     ```

  2. **Application de toutes les migrations** (`.up`) :
     
     ```bash
     docker-compose exec go migrate -path internal/infrastructure/database/migrations -database "mysql://root:root@tcp(mysql:3306)/database" up
     ```

  3. **Application de toutes les migrations** (`.down`) :
     
     ```bash
     docker-compose exec go migrate -path internal/infrastructure/database/migrations -database "mysql://root:root@tcp(mysql:3306)/database" down
     ```

## Génération des fichiers .proto

Les commandes suivantes servent à générer les stubs à partir du fichier **user.proto**. Elles doivent être exécutées depuis le répertoire principal du projet.

- **Fichier backend** :
  
  ```bash
  protoc --go_out=. --go-grpc_out=. proto/user.proto
  ```

- **Fichier frontend** :
  
  ```bash
  protoc --ts_out ./web/src/generated --proto_path "./proto" ./proto/user.proto
  ```

- **Régénérer l'ensemble des fichiers** :
  
  ```bash
  protoc --go_out=. --go-grpc_out=. proto/user.proto && protoc --ts_out ./web/src/generated --proto_path "./proto" ./proto/user.proto
  ```

