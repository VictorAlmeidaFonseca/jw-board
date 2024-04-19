package factory

import (
	config "github.com/VictorAlmeidaFonseca/jw-board/internal/infra/db/config"
	protocol "github.com/VictorAlmeidaFonseca/jw-board/internal/protocol"
)

func ConnectionFactory() protocol.DatabaseHandlerAdapter[config.Pool] {
	conn := config.NewConnection()
	return protocol.NewDatabaseHandlerAdapter(conn)
}
