ALTER TABLE inscripcion.transferencia ADD COLUMN cantidad_creditos NUMERIC(3) NULL;
ALTER TABLE inscripcion.transferencia ADD COLUMN documento_id INTEGER;
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.cantidad_creditos IS 'Número de creditos obtenidos en la carrera que proviene';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.documento_id IS 'Identificador del documento relacionado en la api de Documentos';


ALTER TABLE inscripcion.reintegro ADD COLUMN cantidad_creditos NUMERIC(3) NULL;
ALTER TABLE inscripcion.reintegro ADD COLUMN documento_id INTEGER;
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.cantidad_creditos IS 'Número de creditos obtenidos en la carrera que proviene';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.documento_id IS 'Identificador del documento relacionado en la api de Documentos';