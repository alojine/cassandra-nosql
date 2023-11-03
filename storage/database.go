package storage

import (
	"log"

	"github.com/gocql/gocql"
)

type Database struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func SetupDBConnection() (*Database, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "factory"
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	log.Println("Connected to Cassandra on port :", cluster.Port)

	return &Database{
		cluster: cluster,
		session: session,
	}, nil
}
