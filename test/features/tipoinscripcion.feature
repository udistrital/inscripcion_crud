Feature: Validate API responses
    INSCRIPCION_CRUD
    probe JSON responses

Scenario Outline: To probe route code response /tipo_inscripcion
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"

    Examples:
    |method |route                         |bodyreq                       |codres         |
    |GET    |/v1/tipo_inscripcion          |./assets/requests/empty.json  |200 OK         |
    |GET    |/v1/tipo_inscripcio           |./assets/requests/empty.json  |404 Not Found  |
    |POST   |/v1/tipo_inscripcion/0        |./assets/requests/empty.json  |404 Not Found  |
    |POST   |/v1/tipo_inscripcion          |./assets/requests/empty.json  |400 Bad Request|
    |PUT    |/v1/tipo_inscripcio           |./assets/requests/empty.json  |404 Not Found  |
    |PUT    |/v1/tipo_inscripcion          |./assets/requests/empty.json  |400 Bad Request|
    |DELETE |/v1/tipo_inscripcio           |./assets/requests/empty.json  |404 Not Found  |
    |DELETE |/v1/tipo_inscripcion          |./assets/requests/empty.json  |404 Not Found  |


Scenario Outline: To probe response route /tipo_inscripcion       
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method |route                 |bodyreq                           |codres           |bodyres                    |
    |GET    |/v1/tipo_inscripcion  |./assets/requests/empty.json      |200 OK           |./assets/res0/Vok2.json    |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/empty.json      |400 Bad Request  |./assets/res0/Ierr6.json   |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/BodyGen1.json   |201 Created      |./assets/res0/Vok1.json    |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/Nt1.json        |400 Bad Request  |./assets/res0/Ierr1.json   |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/Nt2.json        |400 Bad Request  |./assets/res0/Ierr2.json   |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/Nt3.json        |400 Bad Request  |./assets/res0/Ierr3.json   |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/Nt4.json        |400 Bad Request  |./assets/res0/Ierr4.json   |
    |POST   |/v1/tipo_inscripcion  |./assets/requests/Nt5.json        |400 Bad Request  |./assets/res0/Ierr5.json   | 
    |PUT    |/v1/tipo_inscripcion  |./assets/requests/BodyRec2.json   |200 OK           |./assets/res0/Vok1.json    |
    |GETID  |/v1/tipo_inscripcion  |./assets/requests/BodyGen1.json   |200 OK           |./assets/res0/Vok1.json    |
    |DELETE |/v1/tipo_inscripcion  |./assets/requests/empty.json      |200 OK           |./assets/res0/Ino.json     |
