package event

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"0chain.net/smartcontract/dbs"
)

func TestCuratorEvent(t *testing.T) {
	t.Skip("only for local debugging, requires local postgresql")
	access := dbs.DbAccess{
		Enabled:         true,
		Name:            "events_db",
		User:            os.Getenv("POSTGRES_USER"),
		Password:        os.Getenv("POSTGRES_PASSWORD"),
		Host:            os.Getenv("POSTGRES_HOST"),
		Port:            os.Getenv("POSTGRES_PORT"),
		MaxIdleConns:    100,
		MaxOpenConns:    200,
		ConnMaxLifetime: 20 * time.Second,
	}
	eventDb, err := NewEventDb(access)
	require.NoError(t, err)
	defer eventDb.Close()
	err = eventDb.AutoMigrate()
	require.NoError(t, err)

	c := Curator{
		AllocationID: "allocation_id",
		CuratorID:    "curator_id",
	}

	err = eventDb.addOrOverwriteCurator(c)
	require.NoError(t, err, "Error while inserting Curator to event Database")

	var count int64
	eventDb.Get().Table("curators").Count(&count)
	require.Equal(t, int64(1), count, "Curator not getting inserted")

	err = eventDb.removeCurator(c)
	require.NoError(t, err, "Error while removing Curator to event Database")

	eventDb.Get().Table("curators").Count(&count)
	require.Equal(t, int64(0), count, "Curator not getting deleted")

	err = eventDb.drop()
	require.NoError(t, err)
}
