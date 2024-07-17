Feature: tipo_inscripcion controller
  Test all endpoints exposed by the controller.

  Scenario Outline: To probe route code response /tipo_inscripcion
    Given I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"

    Examples:
    |method|route               |bodyreq               |codres       |
    |GET   |/v1/tipo_inscripcion|./files/req/Vacio.json|200 OK       |
    |GET   |/v1/tipo_inscripcio |./files/req/Vacio.json|404 Not Found|
    |POST  |/v1/tipo_inscripcio |./files/req/Vacio.json|404 Not Found|
    |PUT   |/v1/tipo_inscripcio |./files/req/Vacio.json|404 Not Found|
    |DELETE|/v1/tipo_inscripcio |./files/req/Vacio.json|404 Not Found|

  Scenario Outline: To probe response route /tipo_inscripcion       
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method|route               |bodyreq                                  |codres         |bodyres                                  |
    |GET   |/v1/tipo_inscripcion|./files/req/Vacio.json                   |200 OK         |./files/res/Vtipo_inscripcion_list.json  |
    |POST  |/v1/tipo_inscripcion|./files/req/Vacio.json                   |400 Bad Request|./files/res/Ibad_request.json            |
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create1.json|201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create2.json|201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create3.json|201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create4.json|201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create5.json|201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create6.json|201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |PUT   |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_update.json |200 OK         |./files/res/Vtipo_inscripcion_single.json|
    |GETID |/v1/tipo_inscripcion|./files/req/Vacio.json                   |200 OK         |./files/res/Vtipo_inscripcion_single.json|
    |DELETE|/v1/tipo_inscripcion|./files/req/Vacio.json                   |200 OK         |./files/res/Vtipo_inscripcion_id.json    |
    |DELETE|/v1/tipo_inscripcion|./files/req/Vacio.json                   |404 Not Found  |./files/res/Inot_found.json              |