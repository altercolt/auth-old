package postgres

import (
	"context"
	"github.com/altercolt/auth/internal/core/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
	"time"
)

type myQueryTracer struct {
	log *log.Logger
}

func (tracer *myQueryTracer) TraceQueryStart(
	ctx context.Context,
	_ *pgx.Conn,
	data pgx.TraceQueryStartData) context.Context {
	tracer.log.Println("Executing command", "sql", data.SQL, "args", data.Args)

	return ctx
}

func (tracer *myQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}

func TestUserRepository_Fetch(t *testing.T) {
	ctx := context.Background()
	conf, err := pgxpool.ParseConfig("postgres://altercolt:1952@localhost:5432/auth")
	if err != nil {
		t.Fatalf("db error : %v \n", err)
	}

	conf.ConnConfig.Tracer = &myQueryTracer{log.New(os.Stdout, "DB_TEST : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)}

	db, err := pgxpool.NewWithConfig(ctx, conf)
	repo := userRepository{
		db: db,
	}

	f := user.Filter{
		ID:    []int{2},
		Email: []string{"aybarrel@gmail.com"},
	}

	users, err := repo.Fetch(ctx, f)
	if err != nil {
		t.Fatalf("fetch error : %v", err)
	}

	for _, v := range users {
		t.Logf("result : %v", *v)
	}

}

func TestUserRepository_FetchOne(t *testing.T) {
	ctx := context.Background()
	conf, err := pgxpool.ParseConfig("postgres://altercolt:1952@localhost:5432/auth")
	if err != nil {
		t.Fatalf("db error : %v \n", err)
	}

	conf.ConnConfig.Tracer = &myQueryTracer{log.New(os.Stdout, "DB_TEST : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)}

	db, err := pgxpool.NewWithConfig(ctx, conf)
	repo := userRepository{
		db: db,
	}

	f := user.Filter{
		Username: []string{"aybarrel"},
	}

	usr, err := repo.FetchOne(ctx, f)
	if err != nil {
		t.Fatalf("fetch error : %v", err)
	}

	t.Logf("result : %v", *usr)

}

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	conf, err := pgxpool.ParseConfig("postgres://altercolt:1952@localhost:5432/auth")
	if err != nil {
		t.Fatalf("db error : %v \n", err)
	}

	conf.ConnConfig.Tracer = &myQueryTracer{log.New(os.Stdout, "DB_TEST : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)}

	db, err := pgxpool.NewWithConfig(ctx, conf)
	repo := userRepository{
		db: db,
	}

	model := user.Model{
		Email:     toPtr[string]("helloworld@gmail.com"),
		Username:  toPtr[string]("helloworld"),
		Firstname: toPtr[string]("Hello"),
		Lastname:  toPtr[string]("World"),
		BirthDate: toPtr[time.Time](time.Now()),
		Salt:      toPtr[string]("methsalt"),
		PassHash:  toPtr[string]("randomhash"),
	}

	if err := repo.Create(ctx, &model); err != nil {
		t.Fatalf("db error : %v", err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	ctx := context.Background()
	conf, err := pgxpool.ParseConfig("postgres://altercolt:1952@localhost:5432/auth")
	if err != nil {
		t.Fatalf("db error : %v \n", err)
	}

	conf.ConnConfig.Tracer = &myQueryTracer{log.New(os.Stdout, "DB_TEST : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)}

	db, err := pgxpool.NewWithConfig(ctx, conf)
	repo := userRepository{
		db: db,
	}

	if err := repo.Delete(ctx, 4); err != nil {
		t.Fatalf("db error : %v", err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	ctx := context.Background()
	conf, err := pgxpool.ParseConfig("postgres://altercolt:1952@localhost:5432/auth")
	if err != nil {
		t.Fatalf("db error : %v \n", err)
	}

	conf.ConnConfig.Tracer = &myQueryTracer{log.New(os.Stdout, "DB_TEST : ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)}

	db, err := pgxpool.NewWithConfig(ctx, conf)
	repo := userRepository{
		db: db,
	}

	model := user.Model{
		ID:        toPtr[int](5),
		Email:     nil,
		Username:  toPtr[string]("updatedUSERNAMESECONDVERSION"),
		Firstname: nil,
		Lastname:  nil,
		BirthDate: nil,
		Salt:      nil,
		PassHash:  nil,
	}

	if err := repo.Update(ctx, &model); err != nil {
		t.Fatalf("db error : %v", err)
	}
}

func toPtr[T any](value T) *T {
	return &value
}
