package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

func main() {

	//readFiles("./2")
	readDir("./enron_mail_20110402/maildir")
}

func readFiles(path string, lastDirName string) {

	archivo, err := os.Open(path)

	if err != nil {
		fmt.Println("Hubo un error")
	}

	scanner := bufio.NewScanner(archivo)

	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	readEmail(content, lastDirName)
}

func JSONMarshal(v map[string]string, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func readEmail(content string, lastDirName string) {

	msg := content
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)

	if err != nil {
		log.Fatal(err)
	}

	/**/
	m1 := map[string]string{}

	headers := m.Header
	/*fmt.Println("Date:", header.Get("Date"))*/

	for key := range headers {
		//fmt.Println("Key:", key, " ====> ", "Value:", headers[key][0])
		m1[key] = headers[key][0]
	}

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}

	m1["body"] = string(body)
	m1["directory"] = lastDirName

	responseJson, _ := JSONMarshal(m1, true)

	//fmt.Println(string(responseJson))
	writeJsonFile("{ \"index\" : { \"_index\" : \"emails\" } }" + "\n" + string(responseJson))
}

func writeJsonFile(jsonValue string) {

	var fileName = "./index.ndjson"
	if _, err := os.Stat(fileName); err == nil {
		// path/to/whatever exists
		//si archivo existe se agrega el JSON al final del contenido

		content, errRead := os.ReadFile(fileName)
		if errRead != nil {
			log.Fatal(errRead)
		}

		val := jsonValue
		data := []byte(string(content) + val + "\n")

		err := os.WriteFile(fileName, data, 0)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("done")

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		//si archivo no existe, se crea el archivo
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		//volver a llamar a la funcion para que guarde el primer registro
		writeJsonFile(jsonValue)
		fmt.Println(f.Name())
	}

}

func readDir(dirname string) {
	var lastDirName = ""
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		lastDirName = dirname
		if file.IsDir() {
			fmt.Println(dirname+"/"+file.Name(), file.IsDir())
			readDir(dirname + "/" + file.Name())
		} else {
			fmt.Println(dirname+"/"+file.Name(), file.IsDir())
			readFiles(dirname+"/"+file.Name(), lastDirName)
		}
	}
	//fmt.Println("---------->", lastDirName)
}
