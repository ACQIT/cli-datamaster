package save

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joskeinerG/cli-datamaster/internal"
	"github.com/joskeinerG/cli-datamaster/internal/email"
	"github.com/joskeinerG/cli-datamaster/pkg/db"
	"github.com/shakinm/xlsReader/xls"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func SaveMae(sql db.Mssql, dir, fileName, userId, toUser, tenantId, clienteId, clientScret string) {
	path := filepath.Join(dir, fileName)
	file, err := xls.OpenFile(path)
	if err != nil {
		log.Fatal("Error al abrir el archivo Excel:", err)
	}

	// Obtener el número total de filas
	sheet, _ := file.GetSheet(0)
	totalRows := sheet.GetNumberRows()

	// Procesar cada fila (excepto la primera que es el encabezado)
	for i := 1; i < totalRows; i++ {

		rows, err := sheet.GetRow(i)
		if err != nil {
			fmt.Printf("error al obtener la fila %v", err.Error())
		}
		descripcion, err := rows.GetCol(0)
		if err != nil {
			log.Printf("error al obtener el valor descripcion %v", err)
		}
		direccion, err := rows.GetCol(1)
		if err != nil {
			log.Printf("error al obtener el valor  direcion %v", err)
		}
		telefono, err := rows.GetCol(2)
		if err != nil {
			log.Printf("error al obtener el valor telefono %v", err)
		}
		email, err := rows.GetCol(3)
		if err != nil {
			log.Printf("error al obtener el valor   email %v", err)
		}
		uRLEntidad, err := rows.GetCol(4)
		if err != nil {
			log.Printf("error al obtener el valor  url entidad%v", err)
		}
		categoria, err := rows.GetCol(5)
		if err != nil {
			log.Printf("error al obtener el valor categoria%v", err)
		}
		uRLCNV, err := rows.GetCol(6)
		if err != nil {
			log.Printf("error al obtener el valor url cnv%v", err)
		}
		nroRegistro, err := rows.GetCol(7)
		if err != nil {
			log.Printf("error al obtener el valor nro registro %v", err)
		}
		matriculaCNV, err := rows.GetCol(8)
		if err != nil {
			log.Printf("error al obtener el valor  matricula cnv %v", err)
		}
		tipo, err := rows.GetCol(9)
		if err != nil {
			log.Printf("error al obtener el valor  tipo %v", err)
		}
		// 	// Obtener valores de cada celda usando coordenadas específicas
		agente := db.Mae{
			Descripcion:  convertToUTF8(descripcion.GetString()),
			Direccion:    convertToUTF8(direccion.GetString()),
			Telefono:     convertToUTF8(telefono.GetString()),
			Email:        convertToUTF8(email.GetString()),
			URLEntidad:   convertToUTF8(uRLEntidad.GetString()),
			Categoria:    convertToUTF8(categoria.GetString()),
			URLCNV:       convertToUTF8(uRLCNV.GetString()),
			NroRegistro:  convertToUTF8(nroRegistro.GetString()),
			MatriculaCNV: convertToUTF8(matriculaCNV.GetString()),
			Tipo:         convertToUTF8(tipo.GetString()),
		}

		saveDBMae(sql, &agente)
		// Insertar en la base de datos
		// if err := sql.DB.Table("mae").Create(&agente).Error; err != nil {
		// 	log.Printf("Error al insertar  en Mae : %v", err)
		// } else {
		// 	log.Printf(" se ingresaron a la tabla N: %d", i)
		// }
	}
	internal.MoveOneFile(dir, fileName)
	NotificateMea(sql, userId, toUser, tenantId, clienteId, clientScret)
}

func saveDBMae(sql db.Mssql, mae *db.Mae) {

	query := `
	MERGE INTO mae AS target
USING (VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)) AS source (
   descripcion, direccion, telefono, email, url_entidad, 
   categoria, url_cnv, nro_registro, matricula_cnv, tipo
)
ON target.nro_registro = source. nro_registro 
  AND target.matricula_cnv = source.matricula_cnv
WHEN MATCHED THEN
   UPDATE SET
	target.descripcion = source.descripcion,
       target.direccion = source.direccion,
       target.telefono = source.telefono,
       target.email = source.email,
       target.url_entidad = source.url_entidad,
       target.categoria = source.categoria, 
       target.url_cnv = source.url_cnv,
       target.nro_registro = source.nro_registro,
       target.tipo = source.tipo,
       target.updated_at = GETDATE()
WHEN NOT MATCHED THEN
   INSERT (descripcion, direccion, telefono, email, url_entidad,
           categoria, url_cnv, nro_registro, matricula_cnv, tipo,
           created_at)
   VALUES (source.descripcion, source.direccion, source.telefono,
           source.email, source.url_entidad, source.categoria,
           source.url_cnv, source.nro_registro, source.matricula_cnv,
           source.tipo, GETDATE());
	`
	err := sql.DB.Exec(query, mae.Descripcion, mae.Direccion, mae.Telefono, mae.Email, mae.URLEntidad, mae.Categoria, mae.URLCNV, mae.NroRegistro, mae.MatriculaCNV, mae.Tipo).Error

	if err != nil {
		log.Printf(" error al guardar los datos de mae %v", err.Error())
	}
	if err == nil {
		log.Println(" se cargaron o actualizaron los datos exitosamente de mae")
	}

}
func getNameMae() string {
	day := strconv.Itoa(time.Now().Day())
	month := time.Now().Month()
	monthS := strconv.Itoa(int(month))
	if len(day) <= 1 {
		day = fmt.Sprintf("0%s", day)
	}
	if len(monthS) <= 1 {
		monthS = fmt.Sprintf("0%s", monthS)
	}

	NameFileMae := fmt.Sprintf("miembros-y-adherentes%d%s%s.xls", time.Now().Year(), monthS, day)
	return NameFileMae
}
func convertToUTF8(text string) string {
	reader := transform.NewReader(strings.NewReader(text), charmap.Windows1252.NewDecoder())
	result, _ := io.ReadAll(reader)
	return string(result)
}
func parserMarketToTemplateMae(market db.Mae, data email.TempaleteData) email.TempaleteData {

	data.NombreALYC = market.Descripcion
	data.NombreCliente = " Javier "
	data.NombreEmpresa = " Acqit "
	data.URLDetalle = market.URLEntidad
	data.Phone = market.Telefono
	data.EmailALYC = market.Email
	data.Market = "Mae"

	return data
}
func NotificateMea(sql db.Mssql, userId, toUser, tenantId, clienteId, clientScret string) {
	var (
		mae      []db.Mae
		htmlBody email.TempaleteData
	)

	if err := sql.DB.Table("mae").Where("created_at >= DATEADD(day, DATEDIFF(day, 0, GETDATE()), 0) AND created_at < DATEADD(day, DATEDIFF(day, 0, GETDATE()) + 1, 0)").Find(&mae).Error; err != nil {
		log.Printf("Error al obtener datos de la tabla byma: %v", err)
		return
	}

	if len(mae) > 0 {
		for _, market := range mae {
			htmlBody = parserMarketToTemplateMae(market, htmlBody)

			email.SendEmail(htmlBody, userId, toUser, tenantId, clienteId, clientScret)
		}

	} else {
		log.Println("No hay nuevos registros")
	}

}
