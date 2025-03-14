package save

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/joskeinerG/cli-datamaster/internal"
	"github.com/joskeinerG/cli-datamaster/internal/email"
	"github.com/joskeinerG/cli-datamaster/pkg/db"
	"github.com/xuri/excelize/v2"
)

func SaveByma(sql db.Mssql, dir, fileName, userId, toUser, tenantId, clienteId, clientScret string) {
	path := filepath.Join(dir, fileName)
	file, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatal("Error al abrir el archivo Excel :", err)
	}
	defer file.Close()

	// Obtener el número total de filas
	rows, _ := file.GetRows("alycs")
	totalRows := len(rows)
	log.Println(totalRows)
	// Procesar cada fila (excepto la primera que es el encabezado)
	for i := 2; i <= totalRows; i++ {
		// Obtener cada valor usando coordenadas específicas

		// Obtener valores de cada celda usando coordenadas específicas
		agente := db.Byma{
			Titulo: getCellValue(file, "alycs", fmt.Sprintf("A%d", i)),

			Matricula:    convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("B%d", i))),
			Participante: convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("C%d", i))),
			Categoria:    convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("D%d", i))),
			Direccion:    convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("E%d", i))),
			Phone:        convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("F%d", i))),
			Fax:          convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("G%d", i))),
			Email:        convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("H%d", i))),
			Web:          convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("I%d", i))),
			Leyenda:      convertToUTF8(getCellValue(file, "alycs", fmt.Sprintf("J%d", i))),
		}

		saveDBByma(sql, &agente)

	}

	internal.MoveOneFile(dir, "alycs.xlsx")
	NotificateByma(sql, userId, toUser, tenantId, clienteId, clientScret)

}

func NotificateByma(sql db.Mssql, userId, toUser, tenantId, clienteId, clientScret string) {
	var (
		byma     []db.Byma
		htmlBody email.TempaleteData
	)

	if err := sql.DB.Table("byma").Where("created_at >= DATEADD(day, DATEDIFF(day, 0, GETDATE()), 0) AND created_at < DATEADD(day, DATEDIFF(day, 0, GETDATE()) + 1, 0)").Find(&byma).Error; err != nil {

		log.Printf("Error al obtener datos de la tabla byma: %v", err)
		return
	}
	if len(byma) > 0 {
		for _, market := range byma {

			htmlBody = parserMarketToTemplateByma(market, htmlBody)
			email.SendEmail(htmlBody, userId, toUser, tenantId, clienteId, clientScret)
		}

	} else {
		log.Println("No hay nuevos registros")
	}

}

func parserMarketToTemplateByma(byma db.Byma, data email.TempaleteData) email.TempaleteData {

	data.NombreALYC = byma.Titulo
	data.NombreCliente = "Javier"
	data.NombreEmpresa = " Acqit "
	data.URLDetalle = byma.Web
	data.Phone = byma.Phone
	data.EmailALYC = byma.Email
	data.Market = "Byma"

	return data
}

func saveDBByma(sql db.Mssql, byma *db.Byma) {
	query := `
	MERGE INTO byma AS target
USING (VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)) AS source (
   titulo, matricula, participante, categoria, direccion,
   phone, fax, email, web, leyenda
)
ON target.matricula = source.matricula
WHEN MATCHED THEN
   UPDATE SET
       target.titulo = source.titulo,
       target.participante = source.participante,
       target.categoria = source.categoria,
       target.direccion = source.direccion,
       target.phone = source.phone,
       target.fax = source.fax,
       target.email = source.email,
       target.web = source.web,
       target.leyenda = source.leyenda,
       target.updated_at = GETDATE()
WHEN NOT MATCHED THEN
   INSERT (titulo, matricula, participante, categoria, direccion,
           phone, fax, email, web, leyenda, created_at)
   VALUES (source.titulo, source.matricula, source.participante,
           source.categoria, source.direccion, source.phone,
           source.fax, source.email, source.web, source.leyenda,
           GETDATE());
	`
	err := sql.DB.Exec(query, byma.Titulo, byma.Matricula, byma.Participante, byma.Categoria, byma.Direccion, byma.Phone, byma.Fax, byma.Email, byma.Web, byma.Leyenda).Error
	if err != nil {
		log.Printf(" error al guardar los datos byma %v", err)
	}
	if err == nil {
		log.Println(" se cargaron o actualizaron los datos exitosamente byma")
	}
}
