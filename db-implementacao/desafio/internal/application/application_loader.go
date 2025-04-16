package application

import (
	"database/sql"
	"log"
	"os"

	"app/internal"
	"app/internal/loader"
	"app/internal/migrator"
	"app/internal/repository"

	"github.com/go-sql-driver/mysql"
)

type ConfigApplicationLoader struct {
	Db                *mysql.Config
	FilePathCustomers string
	FilePathInvoices  string
	FilePathProducts  string
	FilePathSales     string
}

func NewApplicationLoader(config *ConfigApplicationLoader) *ApplicationLoader {
	return &ApplicationLoader{
		config: config,
	}
}

type ApplicationLoader struct {
	config    *ConfigApplicationLoader
	database  *sql.DB
	migrators []internal.Migrator

	customersFile *os.File
	invoicesFile  *os.File
	productsFile  *os.File
	salesFile     *os.File
}

func (a *ApplicationLoader) TearDown() {
	if a.customersFile != nil {
		a.customersFile.Close()
	}
	if a.invoicesFile != nil {
		a.invoicesFile.Close()
	}
	if a.productsFile != nil {
		a.productsFile.Close()
	}
	if a.salesFile != nil {
		a.salesFile.Close()
	}

	if a.database != nil {
		a.database.Close()
	}
}

func (a *ApplicationLoader) SetUp() (err error) {
	a.customersFile, err = os.Open(a.config.FilePathCustomers)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	a.invoicesFile, err = os.Open(a.config.FilePathInvoices)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	a.productsFile, err = os.Open(a.config.FilePathProducts)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	a.salesFile, err = os.Open(a.config.FilePathSales)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	a.database, err = sql.Open("mysql", a.config.Db.FormatDSN())
	if err != nil {
		return
	}

	// ler arquivo
	ldCustomers := loader.NewCustomersJSON(a.customersFile)
	ldInvoices := loader.NewInvoicesJSON(a.invoicesFile)
	ldProducts := loader.NewProductsJSON(a.productsFile)
	ldSales := loader.NewSalesJSON(a.salesFile)

	rpCustomers := repository.NewCustomersMySQL(a.database)
	rpInvoices := repository.NewInvoicesMySQL(a.database)
	rpProducts := repository.NewProductsMySQL(a.database)
	rpSales := repository.NewSalesMySQL(a.database)

	mgCustomers := migrator.NewCustomerMigrator(ldCustomers, rpCustomers)
	mgInvoices := migrator.NewInvoiceMigrator(ldInvoices, rpInvoices)
	mgProducts := migrator.NewProductMigrator(ldProducts, rpProducts)
	mgSales := migrator.NewSaleMigrator(ldSales, rpSales)

	a.migrators = []internal.Migrator{
		mgCustomers,
		mgInvoices,
		mgProducts,
		mgSales,
	}

	return
}

func (a *ApplicationLoader) Run() (err error) {
	for _, v := range a.migrators {
		err = v.Migrate()
		if err != nil {
			return
		}
	}
	return
}
