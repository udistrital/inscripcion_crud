package mocks

import "github.com/DATA-DOG/go-sqlmock"

//GETALL MOCK
func GetAllTipoInscripcionMock(mock sqlmock.Sqlmock) {
	//SELECT ALL MOCK
	mock.ExpectPrepare(`SELECT .* FROM "tipo_inscripcion" .* LIMIT 10`).
		ExpectQuery().
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre", "descripcion", "codigo_abreviacion", "activo", "numero_orden", "nivel_id", "fecha_creacion", "fecha_modificacion", "especial"}).
			AddRow(1, "Nombre1", "Descripción1", "Abreviación1", true, 1.0, 1, "2024-08-09T10:57:41.965807-05:00", "2024-08-09T10:57:41.9662181-05:00", true))
}

//POST MOCK
func PostTipoInscripcionMock(mock sqlmock.Sqlmock) {
	mock.ExpectPrepare(`INSERT INTO "tipo_inscripcion".*RETURNING "id"`).
		ExpectQuery().
		WithArgs(
			"Nombre",
			"Descripcion",
			"Abreviacion",
			true,
			1.0,
			1,
			"2024-08-09T10:57:41.965807-05:00",
			"2024-08-09T10:57:41.965807-05:00",
			true,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
}

//GET BY ID MOCK
func GetByIdTipoInscripcionMock(mock sqlmock.Sqlmock) {
	mock.ExpectPrepare(`SELECT .* FROM "tipo_inscripcion" WHERE "id" = \$1$`).
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre", "descripcion", "codigo_abreviacion", "activo", "numero_orden", "nivel_id", "fecha_creacion", "fecha_modificacion", "especial"}).
			AddRow(1, "NombreExistente", "DescripciónExistente", "AbreviaciónExistente", true, 1, 1, "2024-08-09T10:57:41.965807-05:00", "2024-08-09T10:57:41.9662181-05:00", true))
}

// PUT MOCK
func PutTipoInscripcion(mock sqlmock.Sqlmock) {

	GetByIdTipoInscripcionMock(mock)

	mock.ExpectQuery(`SELECT .* FROM "tipo_inscripcion" WHERE "id" = \$1$`).
		WithArgs(1). 
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre", "descripcion", "codigo_abreviacion", "activo", "numero_orden", "nivel_id", "fecha_creacion", "fecha_modificacion", "especial"}).
			AddRow(1, "NombreExistente", "DescripciónExistente", "AbreviaciónExistente", true, 1, 1, "2024-08-09T10:57:41.965807-05:00", "2024-08-09T10:57:41.9662181-05:00", true))

	mock.ExpectPrepare(`UPDATE "tipo_inscripcion" SET .* WHERE "id" = \$10`).
		ExpectExec().
		WithArgs("string", "string", "string", true, 1.0, 0, "2024-08-09t10:57:41.965807-05:00Z", "2024-08-09t10:57:41.9662181-05:00Z", true, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

//DELETE MOCK -> no funcional por solapamiento de consultas
func DeleteTipoInscripcion(mock sqlmock.Sqlmock) {

	GetByIdTipoInscripcionMock(mock)

	mock.ExpectPrepare(`DELETE FROM "tipo_inscripcion" WHERE "id" = \$1`).
		ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(`SELECT T0\."id" FROM "cupo_inscripcion" T0 WHERE T0\."tipo_inscripcion_id" IN \(\$1\)`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
}
