package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joskeinerG/cli-datamaster/internal/email"
	"github.com/joskeinerG/cli-datamaster/pkg/viper"
)

const (
	downloadWait    = 220 * time.Second // Ajustado a 100 segundos como en el script original
	downloadTimeout = 300 * time.Second
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
	ToUser          = "joskeiner.simosa@acqit.com.ar"
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

	// log.Println(" configuracion carga ")
	// db, err := db.NewMssql(cfg.Database)
	// if err != nil {

	// 	log.Printf(" fallo la conf de db %v", err.Error())
	// }
	// defer db.Close()

	// err = db.Migrate()
	// if err != nil {
	// }

	day := strconv.Itoa(time.Now().Day())
	month := time.Now().Month()
	monthS := strconv.Itoa(int(month))
	if len(day) <= 1 {
		day = fmt.Sprintf("0%s", day)
	}
	if len(monthS) <= 1 {
		monthS = fmt.Sprintf("0%s", monthS)
	}

	// NameFileMae := fmt.Sprintf("miembros-y-adherentes%d%s%s.xls", time.Now().Year(), monthS, day)
	// internal.CreateFolders(dir, "procesados", "logs")
	// internal.Download(urlByma, "#print_alycs", dir, downloadTimeout, pageLoadTimeout, downloadTimeout)
	// internal.Download(urlMae, "button.md-button.md-accent.md-raised.csv.md-theme-default", dir, downloadTimeout, pageLoadTimeout, downloadWait)
	// internal.Download(urlRofex, "a.icon-button", dir, downloadTimeout, pageLoadTimeout, downloadWait)
	// //internal.DownloadMav(urlMav, dir)

	// internal.DownloadCafci(urlCafci, dir, nameFileCafci)
	// internal.FindFiles(NameFileMav, NameFileRofex, NameFileMae, NameFileByma, files.DirDownload, files.Dir)
	// save.SaveByma(*db, files.Dir, NameFileByma)
	// save.SaveMae(*db)
	// save.SaveRofex(*db)
	// // save.SaveMav(db)
	// save.SaveCafci(*db)
	data := email.TempaleteData{
		NombreCliente: "Joskeiner",
		NombreALYC:    "Byma",
		URLDetalle:    "google.com",
		NombreEmpresa: "Acqit",
	}
	email.SendEmail(data, configEmail.MicrosoftUserId, ToUser, configEmail.MicrosoftTenantId, configEmail.MicrosoftClientId, configEmail.MicrosoftClientSecret)
}
