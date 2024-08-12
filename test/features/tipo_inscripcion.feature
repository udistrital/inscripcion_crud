Feature: tipo_inscripcion controller
  Test all endpoints exposed by the controller.

  Scenario Outline: To probe response route /tipo_inscripcion       
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method|route               |bodyreq                                   |codres         |bodyres                                  |
    |DELETE|/v1/tipo_inscripcion|./files/req/Vacio.json                    |200 OK         |./files/res/Vtipo_inscripcion_id.json    |

    