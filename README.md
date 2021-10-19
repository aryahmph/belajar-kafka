## Getting Started

1. Start with cloning this repo on your local machine :

```
$ git clone git clone https://github.com/aryahmph/belajar_kafka
$ cd belajar_kafka
```

2. Install apache kafka, and switch to apache kafka folder
3. Run zookeeper

```
$ ./bin/zookeeper-server-start.sh config/zookeeper.properties
```

4. Run apache kafka server

```
$ ./bin/kafka-server-start.sh config/server.properties
```

5. Create mysql/mariadb database named `belajar_kafka` with table :

```mysql
CREATE TABLE IF NOT EXISTS students
(
    id    INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name  VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
```

6.Connect to http://localhost:3000 
