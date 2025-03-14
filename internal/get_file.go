package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// CreateFolders crea una estructura de directorios que consiste en una carpeta principal
// y dos subcarpetas: una para archivos procesados y otra para logs.
//
// Parámetros:
//   - mainRouteName: Nombre de la carpeta principal
//   - filesProcessedName: Nombre de la subcarpeta para archivos procesados
//   - filesLogsName: Nombre de la subcarpeta para logs
//
// Returns:
//   - bool: true si todas las carpetas se crearon correctamente o ya existen,
//     false si ocurrió algún error durante el proceso
//
// La función realiza las siguientes operaciones:
//  1. Verifica si la carpeta principal existe
//  2. Si no existe:
//     - Crea la carpeta principal con permisos 0755
//     - Crea la subcarpeta para archivos procesados
//     - Crea la subcarpeta para logs
//  3. Si la carpeta principal existe:
//     - Verifica y crea la subcarpeta de archivos procesados si no existe
//     - Verifica y crea la subcarpeta de logs si no existe
//
// La función registra mensajes informativos y de error usando el paquete log.
// Todos los directorios se crean con permisos 0755 (lectura y ejecución para todos,
// escritura solo para el propietario).
func CreateFolders(mainRouteName, filesProcessedName, filesLogsName string) bool {

	exist, err := os.Stat(mainRouteName)
	if os.IsNotExist(err) {
		log.Printf("La carpeta %s no existe", mainRouteName)
		log.Println("Creando carpeta principal")

		err := os.Mkdir(mainRouteName, 0755)
		if err != nil {
			log.Printf("error : %v", err.Error())
			log.Printf("Error al crear la carpeta %s", mainRouteName)
			return false
		}

		log.Println("Creando carpeta de archivos procesados")

		path := fmt.Sprintf("%s/%s", mainRouteName, filesProcessedName)
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Printf("Error al crear la carpeta %s", mainRouteName)
			return false
		}

		log.Println("Creando carpeta para los logs")

		path = fmt.Sprintf("%s/%s", mainRouteName, filesLogsName)
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Panicf("Error al crear la carpeta %s", mainRouteName)
			return false
		}

		log.Println("Todas las carpetas fueron creadas con exito")
		return true
	}

	if exist.IsDir() {
		log.Println("Las  ruta principal ya existe")
		pathProcessed := fmt.Sprintf("%s/%s", mainRouteName, filesProcessedName)
		exist, err := os.Stat(pathProcessed)
		if os.IsNotExist(err) || !exist.IsDir() {
			log.Println("Creando carpeta de archivos procesados")
			err = os.Mkdir(pathProcessed, 0755)
			if err != nil {
				log.Printf("Error al crear la carpeta %s", mainRouteName)
				return false
			}
		}

		pathLogs := fmt.Sprintf("%s/%s", mainRouteName, filesLogsName)
		exist, err = os.Stat(pathLogs)
		if os.IsNotExist(err) || !exist.IsDir() {
			log.Println("Creando carpeta para los logs")
			err = os.Mkdir(pathLogs, 0755)
			if err != nil {
				log.Printf("Error al crear la carpeta %s", mainRouteName)
				return false
			}
		}

		return true
	}

	err = os.Mkdir(mainRouteName, 0755)
	if err != nil {
		log.Printf("Error al crear la carpeta %s", mainRouteName)
		return false
	}

	log.Println("Creando carpeta de archivos procesados")

	path := fmt.Sprintf("%s/%s", mainRouteName, filesProcessedName)
	err = os.Mkdir(path, 0755)
	if err != nil {
		log.Printf("Error al crear la carpeta %s", mainRouteName)
		return false
	}

	log.Println("Creando carpeta para los logs")

	path = fmt.Sprintf("%s/%s", mainRouteName, filesLogsName)
	err = os.Mkdir(path, 0755)
	if err != nil {
		log.Printf("Error al crear la carpeta %s", mainRouteName)
		return false
	}

	log.Println("Todas las carpetas fueron creadas con exito")
	return true

}

func MoveDownloadedFiles(origin, newAddress string) bool {
	if err := os.Rename(origin, newAddress); err != nil {
		log.Printf("Error al mover el archivo %s", origin)
		return false
	}
	return true
}

func FindFiles(NameFileMav, NameFileRofex, NameFileMae, NameFileByma, dir, newAddress string) {
	log.Printf("Buscando archivos en: %s", dir)
	log.Printf("Nombres a buscar: %s, %s, %s", NameFileMav, NameFileRofex, NameFileMae)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("Error al leer la carpeta %s", dir)
	}
	for _, file := range files {
		if file.Name() == NameFileRofex {
			log.Println("Se encontro el archivo de Rofex")
			o := filepath.Join(dir, NameFileRofex)
			a := filepath.Join(newAddress, NameFileRofex)
			result := MoveDownloadedFiles(o, a)
			if !result {
				log.Println("El  no se pudo mover el archivo de roffex")
			}
			log.Println("El archivo de roffex se movio con exito")
		}
		if file.Name() == NameFileMae {
			log.Println("Se encontro el archivo Mae")
			o := filepath.Join(dir, NameFileMae)
			a := filepath.Join(newAddress, NameFileMae)
			result := MoveDownloadedFiles(o, a)
			if !result {
				log.Println("El  no se pudo mover el archivo de Mae")
			}
			log.Println("El archivo de mae se movio con exito")
		}
		if file.Name() == NameFileMav {
			log.Println("Se encontro el archivo Nomina_Agentes_MAV")
			o := filepath.Join(dir, NameFileMav)
			a := filepath.Join(newAddress, NameFileMav)
			result := MoveDownloadedFiles(o, a)
			if !result {
				log.Println("El  no se pudo mover el archivo mav")
			}
			log.Println("El archivo de mav se movio con exito")
		}
		if file.Name() == NameFileByma {
			log.Println("Se encontro el archivo Byma")
			o := filepath.Join(dir, NameFileByma)
			a := filepath.Join(newAddress, NameFileByma)
			result := MoveDownloadedFiles(o, a)
			if !result {
				log.Println("no se pudo mover el archivo byma")
			}
			log.Println("El archivo de byma se movio con exito")
		}

	}

}

func MoveOneFile(dir, nameFile string) {
	files, err := os.ReadDir(dir)
	route := fmt.Sprintf("%s\\procesados", dir)
	if err != nil {
		log.Printf("Error al leer la carpeta : %v", dir)
	}
	for _, file := range files {
		if file.Name() == nameFile {
			log.Println("Se encontro el archivo a mover ")
			o := filepath.Join(dir, nameFile)
			a := filepath.Join(route, nameFile)
			result := MoveDownloadedFiles(o, a)

			if !result {
				log.Println(" no se pudo mover el archivo ")
			}
			log.Println("El archivo de se movio con exito")
			newName := fmt.Sprintf("%s-%s", time.Now().Local().Format("02-01-2006"), nameFile)
			newPath := filepath.Join(route, newName)

			if err := os.Rename(a, newPath); err != nil {
				log.Printf("error al renombrar el archivo %s : error %v", nameFile, err.Error())
			}
		}
	}
}
