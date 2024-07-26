package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var (
    opt         = godog.Options{Output: colors.Colored(os.Stdout)}
    resStatus   string
    resBody     []byte
    savepostres map[string]interface{}
    IntentosAPI = 1
    Id          float64
    debug       = false
    db          *sql.DB
    mock        sqlmock.Sqlmock
)

// initMockDB inicializa el mock de la base de datos
func initMockDB(t *testing.T) {
    var err error
    db, mock, err = sqlmock.New()
    if err != nil {
        t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
    }
}

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

// @init inicia la aplicacion para realizar los test
func init() {
    fmt.Println("Inicio de pruebas de aceptación al API")

    // Inicializar el mock de la base de datos
    initMockDB(nil)

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

// @extractKeysTypes Extraer las llaves de un json
func extractKeysTypes(data interface{}) map[string]reflect.Type {
    keysTypes := make(map[string]reflect.Type)
    value := reflect.ValueOf(data)
    if value.Kind() == reflect.Map {
        for _, key := range value.MapKeys() {
            val := value.MapIndex(key).Interface()
            if val == nil {
                keysTypes[key.String()] = nil
            } else if reflect.TypeOf(val).Kind() == reflect.Map {
                // Recursively check nested objects
                keysTypes[key.String()] = reflect.TypeOf(extractKeysTypes(val))
            } else {
                keysTypes[key.String()] = reflect.TypeOf(val)
            }
        }
    }
    return keysTypes
}

// @sameStructure comparar dos JSON si su estructura es igual retorna true de lo contrario false
func sameStructure(data1, data2 interface{}) bool {
    if data1 == nil || data2 == nil {
        return false
    }

    type1 := reflect.TypeOf(data1)
    type2 := reflect.TypeOf(data2)

    if type1.Kind() != type2.Kind() {
        return false
    }

    if type1.Kind() == reflect.Slice {
        v1 := reflect.ValueOf(data1)
        v2 := reflect.ValueOf(data2)
        if v1.Len() == 0 || v2.Len() == 0 {
            return false
        }
        return sameStructure(v1.Index(0).Interface(), v2.Index(0).Interface())
    } else if type1.Kind() == reflect.Map {
        keysTypes1 := extractKeysTypes(data1)
        keysTypes2 := extractKeysTypes(data2)
        return reflect.DeepEqual(keysTypes1, keysTypes2)
    }

    return false
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

    baseURL := "http://" + "mock-url" + ":" + "mock-port" + endpoint

    switch method {
    case "GET", "POST":
        url = baseURL

    case "PUT", "DELETE", "GETID":
        str := strconv.FormatFloat(Id, 'f', 0, 64)
        url = baseURL + "/" + str

        if method == "GETID" {
            method = "GET"
        }
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

    pages_s := string(pages)
    body_s := string(resBody)

    var data1, data2 interface{}

    if err := json.Unmarshal([]byte(pages_s), &data1); err != nil {
        fmt.Println("Error unmarshalling JSON1:", err)
        return err
    }

    if err := json.Unmarshal([]byte(body_s), &data2); err != nil {
        fmt.Println("Error unmarshalling JSON2:", err)
        return err
    }

    prefix := div[3]

    switch {
    case strings.HasPrefix(prefix, "V"):
        if sameStructure(data1, data2) {
            return nil
        } else {
            return fmt.Errorf("Errores: La estructura del objeto recibido no es la que se esperaba %s != %s", pages_s, body_s)
        }

    case strings.HasPrefix(prefix, "I"):
        areEqual, _ := AreEqualJSON(pages_s, body_s)
        if areEqual {
            return nil
        } else {
            return fmt.Errorf("Se esperaba el body de respuesta %s y se obtuvo %s", pages_s, resBody)
        }
    }

    return fmt.Errorf("Respuesta no validada")
}

// @FeatureContext Define los steps de los escenarios a ejecutar
func FeatureContext(s *godog.ScenarioContext) {
    s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
    s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
    s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)
}
