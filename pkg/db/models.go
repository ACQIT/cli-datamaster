package db

import (
	"time"

	"gorm.io/gorm"
)

type Mav struct {
	ID        uint           `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`

	MatCNV       int    `gorm:"column:mat_cnv;uniqueIndex"`
	OperadorMAV  int    `gorm:"column:operador_mav"`
	DeposCVSA    string `gorm:"column:depos_cvsa;type:varchar(10)"`
	Categoria    string `gorm:"column:categoria;type:varchar(10)"`
	TipoAgente   string `gorm:"column:tipo_agente;type:varchar(50)"`
	Denominacion string `gorm:"column:denominacion;type:varchar(100)"`
	Condicion    string `gorm:"column:condicion;type:varchar(20)"`
	Domicilio    string `gorm:"column:domicilio;type:varchar(100)"`
	Localidad    string `gorm:"column:localidad;type:varchar(50)"`
	Telefono     string `gorm:"column:telefono;type:varchar(50)"`
}

func (Mav) TableName() string {
	return "mav"
}

type Byma struct {
	ID        uint           `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`

	Titulo       string `gorm:"column:titulo;type:varchar(100);not null"`
	Matricula    string `gorm:"column:matricula;type:varchar(20);uniqueIndex;not null"`
	Participante string `gorm:"column:participante;type:varchar(20)"`
	Categoria    string `gorm:"column:categoria;type:varchar(50);index"`
	Direccion    string `gorm:"column:direccion;type:varchar(200)"`
	Phone        string `gorm:"column:phone;type:varchar(100);index"`
	Fax          string `gorm:"column:fax;type:varchar(100)"`
	Email        string `gorm:"column:email;type:varchar(500);index"`
	Web          string `gorm:"column:web;type:varchar(100);index"`
	Leyenda      string `gorm:"column:leyenda;type:text"`
}

func (Byma) TableName() string {
	return "byma"
}

type Mae struct {
	ID        uint           `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`

	Descripcion  string `gorm:"column:descripcion;type:varchar(200);not null;index"`
	Direccion    string `gorm:"column:direccion;type:varchar(255)"`
	Telefono     string `gorm:"column:telefono;type:varchar(50);index"`
	Email        string `gorm:"column:email;type:varchar(500);index"`
	URLEntidad   string `gorm:"column:url_entidad;type:varchar(255)"`
	Categoria    string `gorm:"column:categoria;type:varchar(50);index"`
	URLCNV       string `gorm:"column:url_cnv;type:varchar(255)"`
	NroRegistro  string `gorm:"column:nro_registro;type:varchar(20);index"`
	MatriculaCNV string `gorm:"column:matricula_cnv;type:varchar(50);"`
	Tipo         string `gorm:"column:tipo;type:varchar(50);index"`
}

func (Mae) TableName() string {
	return "mae"
}

type Roffex struct {
	ID        uint           `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`

	RazonSocial           string    `gorm:"column:razon_social;type:varchar(255);not null;index"`
	NumeroRegistroCNV     int       `gorm:"column:numero_registro_cnv"`
	NumeroParticipanteMtR int       `gorm:"column:numero_participante_mtr"`
	CategoriaCNV          string    `gorm:"column:categoria_cnv;type:varchar(255)"`
	FechaAlta             time.Time `gorm:"column:fecha_alta;type:date"`
	Circular              int       `gorm:"column:circular"`
	Direccion             string    `gorm:"column:direccion;type:text"`
	Telefono              string    `gorm:"column:telefono;type:varchar(100);index"`
	Web                   string    `gorm:"column:web;type:varchar(255);index"`
	CorreoElectronico     string    `gorm:"column:correo_electronico;type:varchar(255);index"`
}

func (Roffex) TableName() string {
	return "roffex"
}

type Cafci struct {
	ID        uint           `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`

	CodigoCnv       int    `gorm:"column:codigo_cnv"`
	CodigoCafci     int    `gorm:"column:codigo_cafci"`
	CodigoDeSocDep  int    `gorm:"column:codigo_de_soc_dep"`
	CodigoDeSocGte  int    `gorm:"column:codigo_de_soc_get"`
	SociedadGerente string `gorm:"column:sociedad_gerente"`
}

func (Cafci) TableName() string {
	return "cafci"
}
