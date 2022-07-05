ALTER TABLE inscripcion.documento_programa ADD COLUMN obligatorio BOOLEAN;

COMMENT ON COLUMN inscripcion.documento_programa.obligatorio IS 'Campo para filtrar documentos programa obligatorios';