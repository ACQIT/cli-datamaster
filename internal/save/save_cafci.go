package save

import (
	"fmt"
	"log"

	"github.com/joskeinerG/cli-datamaster/internal"
	"github.com/joskeinerG/cli-datamaster/pkg/db"
	"github.com/xuri/excelize/v2"
)

func SaveCafci(sql db.Mssql) {
	file, err := excelize.OpenFile("C:\\Users\\joskeiner.simosa\\Desktop\\crm\\descargas\\cafci.xlsx")
	if err != nil {
		log.Fatal("Error al abrir el archivo Excel:", err)
	}
	defer file.Close()

	rows, err := file.GetRows("Sheet 1")
	if err != nil {
		log.Printf("error al obtener las filas de el archivo de cafci %v", err.Error())
	}
	totalRows := len(rows) - 5

	for i := 12; i < totalRows; i++ {
		flag := 0
		codCnv := parseIntSafe(getCellValue(file, "Sheet 1", fmt.Sprintf("S%d", i)))
		codCafci := parseIntSafe(getCellValue(file, "Sheet 1", fmt.Sprintf("U%d", i)))
		codSocD := parseIntSafe(getCellValue(file, "Sheet 1", fmt.Sprintf("W%d", i)))
		codSocG := parseIntSafe(getCellValue(file, "Sheet 1", fmt.Sprintf("V%d", i)))
		socGerente := convertToUTF8(getCellValue(file, "Sheet 1", fmt.Sprintf("X%d", i)))

		if codCnv != 0 && codCnv != flag {
			fondo := db.Cafci{
				CodigoCnv:       codCnv,
				CodigoCafci:     codCafci,
				CodigoDeSocDep:  codSocD,
				CodigoDeSocGte:  codSocG,
				SociedadGerente: socGerente,
			}
			log.Println(fondo)
			saveDBCafci(sql, &fondo)
			flag = codCnv
		}
	}
	internal.MoveOneFile("cafci.xlsx")

}

func saveDBCafci(db db.Mssql, fondo *db.Cafci) {
	query := `
	MERGE INTO cafci AS target 
        USING (VALUES (?, ?, ?, ?, ?)) AS source (
            codigo_cafci, codigo_cnv, codigo_de_soc_dep, codigo_de_soc_get, sociedad_gerente)
        ON target.codigo_cnv = source.codigo_cnv
        WHEN MATCHED THEN
            UPDATE SET    
                target.codigo_cnv = source.codigo_cnv,
                target.codigo_cafci = source.codigo_cafci,
                target.codigo_de_soc_dep = source.codigo_de_soc_dep,
                target.codigo_de_soc_get = source.codigo_de_soc_get,
                target.sociedad_gerente = source.sociedad_gerente,
                target.updated_at = GETDATE()
        WHEN NOT MATCHED THEN
            INSERT (codigo_cnv, codigo_cafci, codigo_de_soc_dep, codigo_de_soc_get, sociedad_gerente, created_at)
            VALUES (
                source.codigo_cnv, 
                source.codigo_cafci, 
                source.codigo_de_soc_dep, 
                source.codigo_de_soc_get, 
                source.sociedad_gerente,
                GETDATE()
            );`

	err := db.DB.Exec(query, fondo.CodigoCafci, fondo.CodigoCnv, fondo.CodigoDeSocDep, fondo.CodigoDeSocGte, fondo.SociedadGerente)
	if err != nil {
		log.Printf(" error al guardar los datos de cafci %v", err)
	}
	if err == nil {
		log.Println(" se cargaron o actualizaron los datos exitosamente cafci")
	}
}
