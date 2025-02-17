package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

func Download(downloadURL, buttonSelector, downloadDir string, downloadTimeout, pageLoadTimeout, downloadWait time.Duration) error {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Headless, // Modo headless como en el script original
		chromedp.Flag("disable-software-rasterizer", true),
		chromedp.Flag("download.default_directory", downloadDir),
		chromedp.Flag("download.prompt_for_download", false),
		chromedp.Flag("download.directory_upgrade", true),
		chromedp.Flag("safebrowsing.enabled", false),
		chromedp.Flag("disable-extensions", true),
	)

	// Crear el contexto con las opciones
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Crear el contexto de Chrome con logging
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Establecer timeout global
	ctx, cancel = context.WithTimeout(ctx, downloadTimeout)
	defer cancel()

	fmt.Println("Iniciando proceso de descarga...")

	// const buttonSelector = "button.md-button.md-accent.md-raised.csv.md-theme-default"

	return chromedp.Run(ctx,
		// Navegar a la URL
		chromedp.Navigate(downloadURL),

		// Esperar a que la página cargue (30 segundos como en el script original)
		chromedp.Sleep(pageLoadTimeout),

		// Verificar que estamos en la página correcta
		chromedp.ActionFunc(func(ctx context.Context) error {
			var title string
			if err := chromedp.Title(&title).Do(ctx); err != nil {
				return fmt.Errorf("error obteniendo título: %w", err)
			}
			log.Printf("Página cargada. Título: %s", title)
			return nil
		}),

		// Esperar a que el botón esté presente y visible
		chromedp.WaitVisible(buttonSelector, chromedp.ByQuery),

		// Log antes del clic
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Botón encontrado, intentando hacer clic...")
			return nil
		}),

		// Hacer clic en el botón usando el selector exacto
		chromedp.Click(buttonSelector, chromedp.ByQuery),

		// Log después del clic
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Clic realizado, esperando descarga...")
			return nil
		}),

		// Esperar el tiempo de descarga (100 segundos como en el script original)
		chromedp.Sleep(downloadWait))
}

func DownloadMav(urlMav, dirFile string) {

	req, err := http.NewRequest("GET", urlMav, nil)

	if err != nil {
		log.Printf(" error al momento de descargar mav %v", err.Error())
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("error al momento de hacer la peticion %v", err)
	}

	if res.StatusCode >= 400 {
		log.Println("Hubo un error al momento de descargar mav ")
	}

	if res.StatusCode == 200 {
		log.Println(" respuesta de la peticion OK !!")
	}

	fileName := filepath.Join(dirFile, "Nomina_Agentes_MAV.pdf")

	file, err := os.Create(fileName)

	if err != nil {
		log.Printf("error al crear el archivo %v", err.Error())
	}

	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Printf(" error al momentod de guardar el archivo %v", err.Error())
	}

	log.Printf("Archivo %s descargado exitosamente", fileName)
}

func DownloadCafci(urlCafci, dir, nameFile string) {
	req, err := http.NewRequest("GET", urlCafci, nil)
	if err != nil {
		log.Printf(" error al momento de descargar cafci %v", err.Error())
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("error al momento de hacer la peticion %v", err)
	}

	if res.StatusCode >= 400 {
		log.Println("Hubo un error al momento de descargar CAFCI ")
	}

	if res.StatusCode == 200 {
		log.Println(" respuesta de la peticion OK !!")
	}

	fileName := filepath.Join(dir, nameFile)

	file, err := os.Create(fileName)

	if err != nil {
		log.Printf("error al crear el archivo %v", err.Error())
	}

	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Printf(" error al momentod de guardar el archivo %v", err.Error())
	}

	log.Println("Archivo de cafci descargado exitosamente")

}
