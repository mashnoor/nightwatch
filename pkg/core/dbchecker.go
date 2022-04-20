package core

import (
	"database/sql"
	"fmt"
	"github.com/mashnoor/nightwatch/pkg/helpers"
	"github.com/mashnoor/nightwatch/pkg/settings"
	"github.com/mashnoor/nightwatch/pkg/strcts"
	"sync"
	"time"
)

func execute(service *strcts.Cluster, wg *sync.WaitGroup) {
	for true {
		checkLag(service)
		time.Sleep(time.Second * settings.SystemAppConfig.EvaluateInterval)

	}
	wg.Done()
}

func checkLag(cluster *strcts.Cluster) {

	currentTimeUTC := time.Now().UTC()
	currentTimeBDT := currentTimeUTC.Add(time.Hour * 6)

	currentTimeBDTStr := currentTimeBDT.Format("2006-01-02 15:04:05")
	var DB_DSN = fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", cluster.DbUser, cluster.DbPassword, cluster.DbHost, cluster.DbPort)
	//fmt.Println(DB_DSN)
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		settings.Log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	// Create an empty user and make the sql query (using $1 for the parameter)
	var dbResult strcts.DbResult
	userSql := "SELECT CASE WHEN pg_last_wal_receive_lsn() = pg_last_wal_replay_lsn() THEN 0 ELSE EXTRACT (EPOCH FROM now() - pg_last_xact_replay_timestamp()) END AS log_delay;"

	err = db.QueryRow(userSql).Scan(&dbResult.LogDelay)
	if err != nil {
		settings.Log.Fatal("Failed to execute query: ", err)
	} else {
		logMsg := fmt.Sprintf("name = %s  log_delay = %f", cluster.ClusterName, dbResult.LogDelay)
		settings.Log.Info(logMsg)

		if dbResult.LogDelay >= cluster.LogDelayThreshold {
			slackMsg := fmt.Sprintf("*:rocket: Night Watch*\n*Cluster name:* %s\n*Log delay:* %.2f\n*Time:* %s", cluster.ClusterName, dbResult.LogDelay, currentTimeBDTStr)
			settings.Log.Info("Sent slack notification ", cluster.ClusterName)
			helpers.SendSlackMessage(slackMsg)
			//settings.Log.Info(slackMsg)
		}

		//fmt.Println(slackMsg)
	}

}
