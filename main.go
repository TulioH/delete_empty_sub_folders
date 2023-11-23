package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "./" // Reemplaza con la ruta de tu carpeta
	err := filepath.Walk(dir, BorrarCarpetasVacias)
	if err != nil {
		fmt.Printf("Error al recorrer la carpeta: %v\n", err)
	}
}

func BorrarCarpetasVacias(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("Error al leer el directorio: %v\n", err)
			return err
		}
		if len(files) == 0 {
			fmt.Printf("La carpeta %s está vacía, se eliminará\n", path)
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Error al eliminar la carpeta: %v\n", err)
				return err
			}
			fmt.Printf("La carpeta %s ha sido eliminada\n", path)
		}
	}
	return nil
}
