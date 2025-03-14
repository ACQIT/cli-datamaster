package db

import (
	"github.com/joskeinerG/cli-datamaster/pkg/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mssql struct {
	DB  *gorm.DB
	Cfg viper.Database
}

// newMssql crea una nueva instancia de conexión a la base de datos SQL Server utilizando GORM.
//
// Parámetros:
//   - cfg: Una estructura `env.Database` que contiene la configuración de conexión (cadena de conexión, etc.).
//
// Retorna:
//   - DB: Una implementación de la interfaz DB que representa la conexión inicializada.
//   - error: Cualquier error encontrado durante la inicialización.
//
// Ejemplo de uso:
//
//	db, err := newMssql(cfg)
//	if err != nil {
//	    log.Fatal("Error al conectar con la base de datos", err)
//	}
//	defer db.Close()
func NewMssql(cfg viper.Database) (*Mssql, error) {

	db, err := gorm.Open(sqlserver.Open(cfg.Conn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return &Mssql{
		DB:  db,
		Cfg: cfg,
	}, nil

}

// DB retorna la instancia subyacente de *gorm.DB para realizar operaciones directas.
//
// Retorna:
//   - *gorm.DB: La conexión actual a la base de datos.
//
// Ejemplo de uso:
//
//	if err := m.conn.DB().Table("mae_env").Where("id=?", env).First(&credentials).Error; err != nil {
//			return nil, err
//		}
func (m *Mssql) Db() *gorm.DB {
	return m.DB
}

// Close cierra la conexión activa con la base de datos.
//
// Retorna:
//   - error: Cualquier error encontrado al intentar cerrar la conexión.
//
// Ejemplo de uso:
//
// defer mssql.Close()
func (m *Mssql) Close() error {
	db, err := m.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Migrate realiza la migración automática de las tablas en la base de datos MSSQL.
// Verifica la existencia de las tablas para los modelos Roffex, Mae, Mav, Byma y Cafci.
// Si alguna tabla no existe, la crea utilizando el esquema definido en su respectivo modelo.
//
// La función ejecuta las siguientes operaciones:
//   - Verifica la existencia de cada tabla usando el Migrator de GORM
//   - Si la tabla no existe, ejecuta AutoMigrate para crearla
//   - Si ocurre algún error durante la migración, retorna inmediatamente
//
// Returns:
//   - error: Retorna nil si todas las migraciones son exitosas,
//     o el primer error encontrado durante el proceso de migración
//
// Este método es útil para la inicialización de la base de datos y
// asegura que todas las tablas necesarias estén creadas antes de
// que la aplicación comience a operar.
func (m *Mssql) Migrate() error {

	models := []interface{}{
		&Roffex{},
		&Mae{},
		&Mav{},
		&Byma{},
		&Cafci{},
	}

	for _, model := range models {
		if !m.DB.Migrator().HasTable(model) {
			if err := m.DB.AutoMigrate(model); err != nil {
				return err
			}
		}
	}
	return nil

}
