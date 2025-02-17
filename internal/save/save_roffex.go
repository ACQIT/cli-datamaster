package save

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/joskeinerG/cli-datamaster/internal"
	"github.com/joskeinerG/cli-datamaster/internal/email"
	"github.com/joskeinerG/cli-datamaster/pkg/db"
	"github.com/xuri/excelize/v2"
)

func SaveRofex(sql db.Mssql, dir, fileName, userId, toUser, tenantId, clienteId, clientScret string) {

	path := filepath.Join(dir, fileName)
	file, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal("Error al abrir el archivo Excel:", err)
	}
	defer file.Close()

	// Obtener el número total de filas
	rows, _ := file.GetRows("Listado de Agentes")
	totalRows := len(rows)
	totalRows = totalRows - 3

	// Procesar cada fila (excepto la primera que es el encabezado)
	for i := 2; i < totalRows; i++ {
		// Obtener cada valor usando coordenadas específicas
		razonSocial := getCellValue(file, "Listado de Agentes", fmt.Sprintf("A%d", i+1))

		// Si la razón social está vacía, saltamos esta fila
		if strings.TrimSpace(razonSocial) == "" {
			continue
		}

		// Obtener valores de cada celda usando coordenadas específicas
		agente := db.Roffex{
			RazonSocial:           convertToUTF8(razonSocial),
			NumeroRegistroCNV:     parseIntSafe(getCellValue(file, "Listado de Agentes", fmt.Sprintf("B%d", i+1))),
			NumeroParticipanteMtR: parseIntSafe(getCellValue(file, "Listado de Agentes", fmt.Sprintf("C%d", i+1))),
			CategoriaCNV:          convertToUTF8(getCellValue(file, "Listado de Agentes", fmt.Sprintf("D%d", i+1))),
			FechaAlta:             parseFechaSafe(getCellValue(file, "Listado de Agentes", fmt.Sprintf("E%d", i+1))),
			Circular:              parseIntSafe(getCellValue(file, "Listado de Agentes", fmt.Sprintf("F%d", i+1))),
			Direccion:             convertToUTF8(getCellValue(file, "Listado de Agentes", fmt.Sprintf("G%d", i+1))),
			Telefono:              convertToUTF8(getCellValue(file, "Listado de Agentes", fmt.Sprintf("H%d", i+1))),
			Web:                   convertToUTF8(getCellValue(file, "Listado de Agentes", fmt.Sprintf("I%d", i+1))),
			CorreoElectronico:     convertToUTF8(getCellValue(file, "Listado de Agentes", fmt.Sprintf("J%d", i+1))),
		}
		saveDBRoffex(sql, &agente)
	}
	internal.MoveOneFile(dir, "agentes.xlsx")
	NotificateRoffex(sql, userId, toUser, tenantId, clienteId, clientScret)
}

// Función auxiliar para obtener el valor de una celda de manera segura
func getCellValue(file *excelize.File, sheet, cell string) string {
	value, err := file.GetCellValue(sheet, cell)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(value)
}

// Función auxiliar para parsear enteros de manera segura
func parseIntSafe(value string) int {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}
	num := 0
	fmt.Sscanf(value, "%d", &num)
	return num
}

// Función auxiliar para parsear fechas de manera segura
func parseFechaSafe(value string) time.Time {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}
	}
	fecha, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}
	}
	return fecha
}
func saveDBRoffex(sql db.Mssql, roffex *db.Roffex) {
	query := `MERGE INTO roffex AS target
USING (VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)) AS source (
   razon_social, numero_registro_cnv, numero_participante_mtr, categoria_cnv,
   fecha_alta, circular, direccion, telefono, web, correo_electronico
)
ON target.numero_registro_cnv = source.numero_registro_cnv
WHEN MATCHED THEN
   UPDATE SET
       target.razon_social = source.razon_social,
       target.numero_participante_mtr = source.numero_participante_mtr,
       target.categoria_cnv = source.categoria_cnv,
       target.fecha_alta = source.fecha_alta,
       target.circular = source.circular,
       target.direccion = source.direccion,
       target.telefono = source.telefono,
       target.web = source.web,
       target.correo_electronico = source.correo_electronico,
       target.updated_at = GETDATE()
WHEN NOT MATCHED THEN
   INSERT (razon_social, numero_registro_cnv, numero_participante_mtr, 
           categoria_cnv, fecha_alta, circular, direccion, telefono, 
           web, correo_electronico, created_at)
   VALUES (source.razon_social, source.numero_registro_cnv, 
           source.numero_participante_mtr, source.categoria_cnv,
           source.fecha_alta, source.circular, source.direccion,
           source.telefono, source.web, source.correo_electronico,
           GETDATE());`

	err := sql.DB.Exec(query, roffex.RazonSocial, roffex.NumeroRegistroCNV, roffex.NumeroParticipanteMtR, roffex.CategoriaCNV, roffex.FechaAlta, roffex.Circular, roffex.Direccion, roffex.Telefono, roffex.Web, roffex.CorreoElectronico).Error
	if err != nil {
		log.Printf("hubo un error a guardar los datos roffex %v", err)
	}
	if err == nil {
		log.Println("se cargo o actualizo todo con exito roffex")
	}
}
func parserMarketToTemplateRoffex(market db.Roffex, data email.TempaleteData) email.TempaleteData {

	data.NombreALYC = market.RazonSocial
	data.NombreCliente = "Javier "
	data.NombreEmpresa = " Acqit "
	data.URLDetalle = market.Web
	data.Phone = market.Telefono
	data.EmailALYC = market.CorreoElectronico
	data.Market = "Roffex"

	return data
}
func NotificateRoffex(sql db.Mssql, userId, toUser, tenantId, clienteId, clientScret string) {
	var (
		roffex   []db.Roffex
		htmlBody email.TempaleteData
	)

	if err := sql.DB.Table("roffex").Where("created_at >= GETDATE()").Find(&roffex).Error; err != nil {
		log.Printf("Error al obtener datos de la tabla byma: %v", err)
		return
	}
	if err := sql.DB.Table("roffex").Where("created_at >= DATEADD(day, DATEDIFF(day, 0, GETDATE()), 0) AND created_at < DATEADD(day, DATEDIFF(day, 0, GETDATE()) + 1, 0)").Find(&roffex).Error; err != nil {
		log.Printf("Error al obtener datos de la tabla byma: %v", err)
		return
	}
	if len(roffex) > 0 {
		for _, market := range roffex {

			htmlBody = parserMarketToTemplateRoffex(market, htmlBody)
			email.SendEmail(htmlBody, userId, toUser, tenantId, clienteId, clientScret)
		}

	} else {
		log.Println("No hay nuevos registros")
	}

}
