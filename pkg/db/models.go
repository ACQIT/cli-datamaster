package db

import (
	"time"

	"gorm.io/gorm"
)

// Mav representa la estructura de un Mercado Abierto de Valores en la base de datos.
//
// La estructura contiene campos que siguen las convenciones de GORM para ORM,
// incluyendo campos para seguimiento de fechas de creación/actualización y borrado lógico.
//
// Campos principales:
//   - MatCNV: Número de matrícula en CNV, identificador único
//   - OperadorMAV: Número de operador en el MAV
//   - DeposCVSA: Código de depósito en Caja de Valores S.A.
//   - Categoria: Categoría del agente MAV
//   - TipoAgente: Tipo de agente
//   - Denominacion: Nombre o denominación social
//   - Condicion: Condición del agente
//   - Domicilio: Dirección física
//   - Localidad: Localidad/ciudad
//   - Telefono: Número telefónico de contacto
//
// La estructura implementa el método TableName() para especificar
// el nombre de la tabla en la base de datos.
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

// Byma representa la estructura de un agente del Mercado de Bolsas y Mercados Argentinos (BYMA) en la base de datos.
//
// La estructura utiliza convenciones de GORM para ORM, incluyendo campos de auditoría
// para seguimiento de creación, actualización y borrado lógico.
//
// Campos principales:
//   - Titulo: Nombre o título del agente BYMA
//   - Matricula: Número de matrícula, identificador único del agente
//   - Participante: Código de participante en el mercado
//   - Categoria: Categoría o tipo de agente
//   - Direccion: Dirección física del agente
//   - Phone: Número telefónico de contacto
//   - Fax: Número de fax (si aplica)
//   - Email: Dirección de correo electrónico de contacto
//   - Web: Sitio web del agente
//   - Leyenda: Información adicional o descripción
//
// La estructura implementa el método TableName() para definir
// explícitamente el nombre de la tabla en la base de datos.

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

// Mae representa la estructura de un agente del Mercado Abierto Electrónico (MAE) en la base de datos.
//
// La estructura implementa convenciones de GORM para ORM, incluyendo los campos estándar
// para el seguimiento de creación, actualización y borrado lógico de registros.
//
// Campos principales:
//   - Descripcion: Nombre o descripción del agente MAE
//   - Direccion: Dirección física del agente
//   - Telefono: Número telefónico de contacto
//   - Email: Dirección de correo electrónico
//   - URLEntidad: Sitio web oficial del agente
//   - Categoria: Categoría o clasificación del agente
//   - URLCNV: Enlace a la información del agente en el sitio de la CNV
//   - NroRegistro: Número de registro del agente
//   - MatriculaCNV: Número de matrícula asignado por la CNV
//   - Tipo: Tipo o modalidad del agente
//
// La estructura implementa el método TableName() para especificar
// explícitamente el nombre de la tabla en la base de datos.
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

// Roffex representa la estructura de un agente del Mercado ROFEX (Rosario Futures Exchange) en la base de datos.
//
// La estructura implementa las convenciones de GORM para ORM, incluyendo campos automáticos
// para seguimiento de creación, actualización y borrado lógico de registros.
//
// Campos principales:
//   - RazonSocial: Nombre o razón social del agente ROFEX
//   - NumeroRegistroCNV: Número de registro en la Comisión Nacional de Valores
//   - NumeroParticipanteMtR: Número de participante en el Mercado a Término de Rosario
//   - CategoriaCNV: Categoría asignada por la CNV
//   - FechaAlta: Fecha de inicio de operaciones en el mercado
//   - Circular: Número de circular relacionada
//   - Direccion: Dirección física del agente
//   - Telefono: Número telefónico de contacto
//   - Web: Sitio web del agente
//   - CorreoElectronico: Dirección de correo electrónico de contacto
//
// La estructura implementa el método TableName() para definir
// explícitamente el nombre de la tabla en la base de datos.
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

// Cafci representa la estructura de un registro de la Cámara Argentina de Fondos Comunes de Inversión (CAFCI) en la base de datos.
//
// La estructura implementa las convenciones de GORM para ORM, incluyendo campos estándar
// para el seguimiento de creación, actualización y borrado lógico de registros.
//
// Campos principales:
//   - CodigoCnv: Código asignado por la Comisión Nacional de Valores
//   - CodigoCafci: Código identificador dentro de la CAFCI
//   - CodigoDeSocDep: Código de la sociedad depositaria
//   - CodigoDeSocGte: Código de la sociedad gerente
//   - SociedadGerente: Nombre de la sociedad gerente del fondo
//
// La estructura implementa el método TableName() para especificar
// explícitamente el nombre de la tabla en la base de datos.
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
