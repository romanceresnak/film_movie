package address

import (
	"context"
	"database/sql"
	actor2 "github.com/romanceresnak/film_project/pkg/api/v1/actor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type addressServiceServer struct {
	db *sql.DB
}

// connect returns SQL database connection from the pool
func (s *addressServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func NewAddressServiceServer(db *sql.DB) actor2.ActorServiceServer {
	return &addressServiceServer{db: db}
}

func (a *addressServiceServer) Create(ctx context.Context, req *actor2.CreateRequest) (*actor2.CreateResponse, error) {
	// insert ToDo entity data
	// get SQL connection from pool
	c, err := a.connect(ctx)
	if err != nil {
		return nil, err
	}

}

func (a *addressServiceServer) Read(context.Context, *actor2.ReadRequest) (*actor2.ReadResponse, error) {
	panic("implement me")
}

func (a *addressServiceServer) Update(context.Context, *actor2.UpdateRequest) (*actor2.UpdateResponse, error) {
	panic("implement me")
}

func (a *addressServiceServer) Delete(context.Context, *actor2.DeleteRequest) (*actor2.DeleteResponse, error) {
	panic("implement me")
}

func (a *addressServiceServer) ReadAll(context.Context, *actor2.ReadAllRequest) (*actor2.ReadAllResponse, error) {
	panic("implement me")
}
