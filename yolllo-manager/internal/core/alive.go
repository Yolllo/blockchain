package core

import "context"

func (c *Core) GetPostgresqlRepoAlive() (resp bool, err error) {
	ctx := context.Background()
	err = c.Repo.PG.Conn.Ping(ctx)
	if err != nil {

		return false, err
	}

	return true, nil
}

func (c *Core) GetElasticsearchRepoAlive() (resp bool, err error) {
	ctx := context.Background()
	pingRequest := c.Repo.ES.Conn.Ping.WithContext(ctx)
	_, err = c.Repo.ES.Conn.Ping(pingRequest)
	if err != nil {

		return false, err
	}

	return true, nil
}
