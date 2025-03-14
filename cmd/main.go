package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/joskeinerG/cli-datamaster/internal"
	"github.com/joskeinerG/cli-datamaster/internal/save"
	"github.com/joskeinerG/cli-datamaster/pkg/db"
	"github.com/joskeinerG/cli-datamaster/pkg/viper"
)

const (
	downloadWait    = 220 * time.Second // Ajustado a 100 segundos como en el script original
	downloadTimeout = 250 * time.Second
	pageLoadTimeout = 45 * time.Second // Ajustado a 30 segundos como en el script original
	urlMae          = "https://www.mae.com.ar/nosotros/quienes-somos/miembros-y-adherentes"
	urlRofex        = "https://matbarofex.com.ar/listado-de-agentes"
	NameFileRofex   = "agentes.xlsx"
	urlByma         = "https://www.byma.com.ar/alycs/"
	NameFileByma    = "alycs.xlsx"
	urlCafci        = "https://api.cafci.org.ar/pb_get"
	nameFileCafci   = "cafci.xlsx"
	urlMav          = "https://www.mav-sa.com.ar/uploads/documentos/Nomina_Agentes_MAV.pdf"
	NameFileMav     = "Nomina_Agentes_MAV.pdf"
	ToUser          = "javier.pombo@acqit.com.ar"
)

func main() {
	cfgN, err := viper.NewViper()
	if err != nil {
		log.Printf(" se cargo las configuracion %v", err.Error())
	}

	cfg, err := cfgN.Load()
	if err != nil {
		log.Printf(" se cargo las configuracion %v", err.Error())
	}

	files := cfg.Files

	configEmail := cfg.EmailConfig

	log.Println(" configuracion carga ")
	db, err := db.NewMssql(cfg.Database)
	if err != nil {

		log.Printf(" fallo la conf de db %v", err.Error())
	}
	defer db.Close()

	err = db.Migrate()
	if err != nil {
		log.Printf(" error al correr las migraciones %v", err)
	}

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
	internal.CreateFolders(files.Dir, "procesados", "logs")
	internal.Download(urlByma, "#print_alycs", files.Dir, downloadTimeout, pageLoadTimeout, downloadTimeout)
	internal.Download(urlMae, "button.md-button.md-accent.md-raised.csv.md-theme-default", files.Dir, downloadTimeout, pageLoadTimeout, downloadWait)
	internal.Download(urlRofex, "a.icon-button", files.Dir, downloadTimeout, pageLoadTimeout, downloadWait)
	internal.DownloadCafci(urlCafci, files.Dir, nameFileCafci)
	procesarMercados(&files, configEmail, ToUser, NameFileMae)

	internal.FindFiles(NameFileMav, NameFileRofex, NameFileMae, NameFileByma, files.DirDownload, files.Dir)
	save.SaveByma(*db, files.Dir, NameFileByma, configEmail.MicrosoftUserId, ToUser, configEmail.MicrosoftTenantId, configEmail.MicrosoftClientId, configEmail.MicrosoftClientSecret)
	save.SaveMae(*db, files.Dir, NameFileMae, configEmail.MicrosoftUserId, ToUser, configEmail.MicrosoftTenantId, configEmail.MicrosoftClientId, configEmail.MicrosoftClientSecret)
	save.SaveRofex(*db, files.Dir, NameFileRofex, configEmail.MicrosoftUserId, ToUser, configEmail.MicrosoftTenantId, configEmail.MicrosoftClientId, configEmail.MicrosoftClientSecret)
	save.SaveCafci(*db, files.Dir, nameFileCafci, configEmail.MicrosoftUserId, ToUser, configEmail.MicrosoftTenantId, configEmail.MicrosoftClientId, configEmail.MicrosoftClientSecret)

}

// procesarMercados descarga información de varios mercados financieros argentinos (BYMA, MAE, ROFEX y CAFCI)
// de forma concurrente utilizando goroutines.
//
// Parámetros:
//   - files *viper.Files: Estructura con información de directorios para guardar los archivos
//   - configEmail viper.EmailConfig: Configuración para envío de correos
//   - toUser string: Dirección de correo del destinatario
//   - nameFileMae string: Nombre del archivo para la descarga del MAE
//
// La función ejecuta 4 goroutines simultáneas, cada una encargada de descargar datos
// de un mercado diferente. Utiliza un WaitGroup para sincronizar y esperar a que todas
// las descargas finalicen antes de continuar.
//
// Las descargas realizadas son:
//   - BYMA: Desde el selector "#print_alycs"
//   - MAE: Desde un botón con clase específica para exportar CSV
//   - ROFEX: Desde un enlace con clase "icon-button"
//   - CAFCI: Mediante una función especializada DownloadCafci
//
// Nota: Los parámetros configEmail y toUser están definidos pero no se utilizan en esta función.
func procesarMercados(files *viper.Files, configEmail viper.EmailConfig, toUser string, nameFileMae string) {
	var wg sync.WaitGroup

	// Inicio de descargas concurrentes
	wg.Add(4)

	// Descarga BYMA
	go func() {
		defer wg.Done()
		internal.Download(urlByma, "#print_alycs", files.Dir, downloadTimeout, pageLoadTimeout, downloadTimeout)
	}()

	// Descarga MAE
	go func() {
		defer wg.Done()
		internal.Download(urlMae, "button.md-button.md-accent.md-raised.csv.md-theme-default", files.Dir, downloadTimeout, pageLoadTimeout, downloadWait)
	}()

	// Descarga ROFEX
	go func() {
		defer wg.Done()
		internal.Download(urlRofex, "a.icon-button", files.Dir, downloadTimeout, pageLoadTimeout, downloadWait)
	}()

	// Descarga CAFCI
	go func() {
		defer wg.Done()
		internal.DownloadCafci(urlCafci, files.Dir, nameFileCafci)
	}()

	// Esperar a que termine todo el procesamiento
	wg.Wait()
}
