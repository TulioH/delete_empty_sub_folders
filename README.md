# delete_empty_sub_folders

Este archivo contiene un programa en Go que elimina carpetas vacías en un directorio especificado.

El programa utiliza la función filepath.Walk para recorrer el directorio y la función BorrarCarpetasVacias para verificar si cada carpeta encontrada está vacía. Si una carpeta está vacía, se elimina utilizando la función os.Remove.

Aquí hay un ejemplo de uso del programa:

dir := "./"
err := filepath.Walk(dir, BorrarCarpetasVacias)
if err != nil {
    fmt.Printf("Error al recorrer la carpeta: %v\n", err)
}
El programa recorre el directorio actual ("./") y muestra mensajes en la consola indicando si una carpeta está vacía y si se ha eliminado correctamente.

Puedes reemplazar dir := "./" con la ruta de tu carpeta para eliminar las carpetas vacías en ese directorio.
