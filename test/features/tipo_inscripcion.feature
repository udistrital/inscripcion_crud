Feature: tipo_inscripcion controller
  Test all endpoints exposed by the controller.

  Scenario Outline: To probe response route /tipo_inscripcion       
    When I create mock whit "<method>"
    And I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method|route               |bodyreq                                   |codres         |bodyres                                  |
    |GET   |/v1/tipo_inscripcion|./files/req/Vacio.json                    |200 OK         |./files/res/Vtipo_inscripcion_list.json  |
    |POST  |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_create2.json |201 Created    |./files/res/Vtipo_inscripcion_single.json|
    |GETID |/v1/tipo_inscripcion|./files/req/Vacio.json                    |200 OK         |./files/res/Vtipo_inscripcion_single.json|
    |PUT   |/v1/tipo_inscripcion|./files/req/tipo_inscripcion_update.json  |200 OK         |./files/res/Vtipo_inscripcion_single.json|
    