package viper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Viper struct {
	*viper.Viper
}
type Container struct {
	// Databese contiene la configuración relacionada con la base de datos
	Database Database

	Files Files

	EmailConfig EmailConfig
}
type Files struct {
	Dir string `mapstructure:"DIR_FILES"`

	DirDownload string `mapstructure:"DIR_DOWNLOAD"`
}
type EmailConfig struct {
	MicrosoftUserId       string `mapstructure:"MICROSOFT_USER_ID"`
	MicrosoftClientId     string `mapstructure:"MICROSOFT_CLIENT_ID"`
	MicrosoftClientSecret string `mapstructure:"MICROSOFT_CLIENT_SECRET"`
	MicrosoftTenantId     string `mapstructure:"MICROSOFT_TENANT_ID"`
}
type Database struct {
	// Conn especifica la cadena de conexión al servidor SQL
	Conn string `mapstructure:"SQL_SERVER"`
	// MaxLifeTime especifica el tiempo máximo de vida para las conexiones en la pool
	MaxLifeTime int `mapstructure:"DB_MAX_LIFE_TIME"`
	// MaxOpenConn especifica el número máximo de conexiones abiertas permitidas
	MaxOpenConn int `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
}

func NewViper() (*Viper, error) {
	v := viper.New()

	path, err := os.Executable()
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(path)

	v.AddConfigPath(dir)
	fmt.Println(dir)
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	err = v.ReadInConfig()

	if err != nil {
		return nil, err
	}
	return &Viper{
		v,
	}, nil
}

// Load carga la configuración de la aplicación y retorna un contenedor con las configuraciones específicas
//
//	base de datos.
//
// Retorna:
//   - *Container: Una estructura que contiene las configuraciones cargadas.
//   - error: Cualquier error encontrado durante el proceso.
//
// Ejemplo de uso:
//
//	container, err := config.Load()
//	if err != nil {
//	    // Manejar el error
//	}
func (v *Viper) Load() (*Container, error) {

	db, err := v.newDBConfig()
	if err != nil {
		return nil, err
	}

	files, err := v.newFilesConfig()

	if err != nil {
		return nil, err
	}

	emailConfig, err := v.newEmailConfig()

	if err != nil {
		return nil, err
	}

	return &Container{
		Database:    db,
		Files:       files,
		EmailConfig: emailConfig,
	}, nil
}

// newDBConfig carga la configuración relacionada con la base de datos.
//
// Retorna:
//   - Database: Una estructura que contiene los parámetros de configuración de la base de datos.
//   - error: Cualquier error encontrado durante el proceso.
//
// Ejemplo de uso:
//
//	dbConfig, err := config.newDBConfig()
//	if err != nil {
//	    // Manejar el error
//	}
func (v *Viper) newDBConfig() (Database, error) {
	var db Database

	err := v.Viper.Unmarshal(&db)
	if err != nil {
		return Database{}, err
	}
	return db, nil
}

// newFilesConfig crea y retorna una nueva configuración de archivos a partir de los valores
// almacenados en Viper. La función deserializa la configuración actual de Viper
// en una estructura Files.
//
// Returns:
//   - Files: La configuración de archivos deserializada
//   - error: Un error si la deserialización falla, nil en caso contrario
//
// La función utiliza Viper.Unmarshal internamente para mapear los valores de configuración
// a los campos de la estructura Files. Si hay algún error durante este proceso,
// retorna una estructura Files vacía junto con el error encontrado.
func (v *Viper) newFilesConfig() (Files, error) {
	var files Files

	err := v.Viper.Unmarshal(&files)
	if err != nil {
		return Files{}, err
	}
	return files, nil
}

// newEmailConfig crea y retorna una nueva configuración de email a partir de los valores
// almacenados en Viper. La función deserializa la configuración actual de Viper
// en una estructura EmailConfig.
//
// Returns:
//   - EmailConfig: La configuración de email deserializada
//   - error: Un error si la deserialización falla, nil en caso contrario
//
// La función utiliza Viper.Unmarshal internamente para mapear los valores de configuración
// a los campos de la estructura EmailConfig. Si hay algún error durante este proceso,
// retorna una estructura EmailConfig vacía junto con el error encontrado.
func (v *Viper) newEmailConfig() (EmailConfig, error) {
	var emailConfig EmailConfig

	err := v.Viper.Unmarshal(&emailConfig)
	if err != nil {
		return EmailConfig{}, err
	}
	return emailConfig, nil
}
