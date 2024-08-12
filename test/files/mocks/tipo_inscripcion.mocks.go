package mocks

import "github.com/DATA-DOG/go-sqlmock"

func GetAllTipoInscripcionMock(mock sqlmock.Sqlmock){
	//SELECT ALL MOCK
	mock.ExpectPrepare(`SELECT .* FROM "tipo_inscripcion" .* LIMIT 10`).
		ExpectQuery().
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre", "descripcion", "codigo_abreviacion", "activo", "numero_orden", "nivel_id", "fecha_creacion", "fecha_modificacion", "especial"}).
			AddRow(1, "Nombre1", "Descripción1", "Abreviación1", true, 1.0, 1, "2024-08-09T10:57:41.965807-05:00", "2024-08-09T10:57:41.9662181-05:00", true))
}