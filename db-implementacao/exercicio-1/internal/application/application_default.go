package application

import (
	"app/internal/application/database"
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewApplicationDefault creates a new default application.
func NewApplicationDefault(addr, filePathStore string) (a *ApplicationDefault) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}
	db := database.GetConnection()

	a = &ApplicationDefault{
		rt:            defaultRouter,
		addr:          defaultAddr,
		filePathStore: filePathStore,
		db:            db,
	}
	return
}

// ApplicationDefault is the default application.
type ApplicationDefault struct {
	// rt is the router.
	rt *chi.Mux
	// addr is the address to listen.
	addr string
	// filePathStore is the file path to store.
	filePathStore string
	// database
	db *sql.DB
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	defer a.db.Close()
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - store
	//st := store.NewStoreProductJSON(a.filePathStore)

	// - repository
	//rp := repository.NewRepositoryProductStore(st)
	rp := repository.NewRepositoryProductDB(a.db)
	// - handler
	hd := handler.NewHandlerProduct(rp)

	// router
	// - middlewares
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PUT /products/{id}
		r.Put("/{id}", hd.UpdateOrCreate())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
