ALTER TABLE inscripcion.soporte_documento_programa DROP CONSTRAINT IF EXISTS fk_inscripcion CASCADE;
ALTER TABLE inscripcion.soporte_documento_programa DROP CONSTRAINT IF EXISTS fk_documento_programa CASCADE;
ALTER TABLE inscripcion.documento_programa DROP CONSTRAINT IF EXISTS fk_tipo_documento_programa CASCADE;
DROP TABLE IF EXISTS inscripcion.soporte_documento_programa CASCADE;
DROP TABLE IF EXISTS inscripcion.documento_programa CASCADE;
DROP TABLE IF EXISTS inscripcion.tipo_documento_programa CASCADE;
