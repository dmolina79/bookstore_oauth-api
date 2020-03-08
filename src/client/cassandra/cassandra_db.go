package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra connection  established")
	defer session.Close()
}

func GetSession() *gocql.Session {
	return session
}
