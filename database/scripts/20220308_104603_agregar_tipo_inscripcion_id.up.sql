ALTER TABLE inscripcion.documento_programa ADD COLUMN tipo_inscripcion_id INTEGER;

COMMENT ON COLUMN inscripcion.documento_programa.tipo_inscripcion_id IS 'Relacion de Id con tabla tipo_inscripcion';
