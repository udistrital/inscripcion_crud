package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	"github.com/astaxie/beego"
)

// @opt opciones de godog
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

// @resStatus codigo de respuesta a las solicitudes a la api
var resStatus string

// @resBody JSON repuesta Delete
var resDelete string

// @resBody JSON de respuesta a las solicitudesde la api
var resBody []byte

var savepostres map[string]interface{}

var IntentosAPI = 1

var Id float64

type BadRequest struct {
	Development string
	Message     string
	Status      string
	System      any
}

var debug = false

// @exe_cmd ejecuta comandos en la terminal
func exe_cmd(cmd string, wg *sync.WaitGroup) {

	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()

	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
}

// @deleteFile Borrar archivos
func deleteFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		err := fmt.Errorf("no se pudo eliminar el archivo")
		fmt.Println(err.Error())
	}
}

// @run_bee activa el servicio de la api para realizar los test
func run_bee() {
	parametros := "INSCRIPCION_CRUD_HTTP_PORT=" + beego.AppConfig.String("httpport") +
		" INSCRIPCION_CRUD_PGUSER=" + beego.AppConfig.String("PGuser") +
		" INSCRIPCION_CRUD_PGPASS=" + beego.AppConfig.String("PGpass") +
		" INSCRIPCION_CRUD_PGHOST=" + beego.AppConfig.String("PGurls") +
		" INSCRIPCION_CRUD_PGPORT=" + beego.AppConfig.String("PGport") +
		" INSCRIPCION_CRUD_PGDB=" + beego.AppConfig.String("PGdb") +
		" INSCRIPCION_CRUD_PGSCHEMA=" + beego.AppConfig.String("PGschemas") + " bee run"

	file, err := os.Create("script.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "cd ..")
	fmt.Fprintln(file, parametros)

	wg := new(sync.WaitGroup)
	commands := []string{"sh script.sh &"}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	time.Sleep(5 * time.Second)
	deleteFile("script.sh")
	wg.Done()
}

// @init inicia la aplicacion para realizar los test
func init() {
	fmt.Println("Inicio de pruebas de aceptación al API")

	run_bee()
	//pasa las banderas al comando godog
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

// @TestMain para realizar la ejecucion con el comando go test ./test
func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "progress", // Utiliza el formato "pretty" para una salida más detallada, "progress" por default
		Paths:  []string{"features"},
		Output: colors.Colored(os.Stdout),
	}

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: FeatureContext,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)

}

/*------------------------------
  ---- Ejecución de pruebas ----
  ------------------------------*/

// @AreEqualJSON comparar dos JSON si son iguales retorna true de lo contrario false
func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func sameStructure(json1, json2 string) bool {
	var data1, data2 interface{}

	if err := json.Unmarshal([]byte(json1), &data1); err != nil {
		fmt.Println("Error unmarshalling JSON1:", err)
		return false
	}

	if err := json.Unmarshal([]byte(json2), &data2); err != nil {
		fmt.Println("Error unmarshalling JSON2:", err)
		return false
	}

	type1 := reflect.TypeOf(data1)
	type2 := reflect.TypeOf(data2)

	if type1.Kind() == reflect.Slice && type2.Kind() == reflect.Slice {
		if reflect.ValueOf(data1).Len() == 0 || reflect.ValueOf(data2).Len() == 0 {
			return false
		}
		return reflect.DeepEqual(reflect.TypeOf(reflect.ValueOf(data1).Index(0).Interface()), reflect.TypeOf(reflect.ValueOf(data2).Index(0).Interface()))
	}

	return reflect.DeepEqual(type1, type2)
}

// @getPages convierte en un tipo el json
func getPages(ruta string) []byte {
	raw, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}

// @iSendRequestToWhereBodyIsJson realiza la solicitud a la API
func iSendRequestToWhereBodyIsJson(method, endpoint, bodyreq string) error {
	if debug {
		fmt.Println("Step: iSendRequestToWhereBodyIsJson")
	}

	var url string

	if method == "GET" || method == "POST" {
		url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + endpoint

	} else {
		if method == "PUT" || method == "DELETE" {
			str := strconv.FormatFloat(Id, 'f', 5, 64)
			url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + endpoint + "/" + str

		}
	}
	if method == "GETID" {
		method = "GET"
		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + endpoint + "/" + str

	}
	if method == "DELETE" {
		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + endpoint + "/" + str
		resDelete = "{\"Id\":" + str + "}"
		os.WriteFile("./assets/responses/Ino.json", []byte(resDelete), 0644)
	}

	if debug {
		fmt.Println("Test: " + method + " to " + url)
	}

	pages := getPages(bodyreq)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(pages))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyr, _ := io.ReadAll(resp.Body)

	resStatus = resp.Status
	resBody = bodyr

	if method == "POST" && resStatus == "201 Created" {
		os.WriteFile("./assets/requests/BodyRec2.json", resBody, 0644)
		json.Unmarshal([]byte(bodyr), &savepostres)
		Id = savepostres["Id"].(float64)

	}
	return nil

}

// @theResponseCodeShouldBe valida el codigo de respuesta
func theResponseCodeShouldBe(arg1 string) error {
	if debug {
		fmt.Println("Step: theResponseCodeShouldBe")
	}

	if resStatus != arg1 {
		return fmt.Errorf("se esperaba el codigo de respuesta .. %s .. y se obtuvo el codigo de respuesta .. %s .. ", arg1, resStatus)
	}
	return nil
}

// @theResponseShouldMatchJson valida el JSON de respuesta
func theResponseShouldMatchJson(arg1 string) error {
	if debug {
		fmt.Println("Step: theResponseShouldMatchJson")
	}

	div := strings.Split(arg1, "/")
	pages := getPages(arg1)
	if strings.HasPrefix(div[3], "V") {
		if sameStructure(string(pages), string(resBody)) {
			return nil
		} else {
			return fmt.Errorf("Errores: La estructura del objeto recibiod no es la que se esperaba")
		}
	}
	if strings.HasPrefix(div[3], "I") {
		areEqual, _ := AreEqualJSON(string(pages), string(resBody))
		if areEqual {
			return nil
		} else {
			return fmt.Errorf(" se esperaba el body de respuesta %s y se obtuvo %s", string(pages), resBody)
		}

	}
	return nil
}

// @iSendRequestToWhereBodyIsMultipartformdataWithThisParamsAndTheFileLocatedAt realiza la solicitud a la API
func iSendRequestToWhereBodyIsMultipartformdataWithThisParamsAndTheFileLocatedAt(method, endpoint, bodyreq string, filename string, bodyfile string) error {
	if debug {
		fmt.Println("Step: iSendRequestToWhereBodyIsMultipartformdataWithThisParamsAndTheFileLocatedAt")
	}

	var url string

	if method == "GET" || method == "POST" {
		url = "http://" + beego.AppConfig.String("appurl") + ":" + beego.AppConfig.String("httpport") + endpoint
	} else {
		if method == "PUT" || method == "DELETE" {
			url = "http://" + beego.AppConfig.String("appurl") + ":" + beego.AppConfig.String("httpport") + endpoint
		}
	}

	extraParams := getPages(bodyreq)
	var params map[string]string
	err := json.Unmarshal(extraParams, &params)
	if err != nil {
		return err
	}

	path, _ := os.Getwd()
	path += "/"
	path += bodyfile

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filename, filepath.Base(path))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bodyr, _ := io.ReadAll(resp.Body)

	resStatus = resp.Status
	resBody = bodyr

	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is multipart\/form-data with this params "([^"]*)" and the file "([^"]*)" located at "([^"]*)"$`, iSendRequestToWhereBodyIsMultipartformdataWithThisParamsAndTheFileLocatedAt)
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
	s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)
}
