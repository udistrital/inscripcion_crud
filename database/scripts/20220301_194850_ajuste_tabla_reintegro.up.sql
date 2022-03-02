ALTER TABLE inscripcion.reintegro ADD COLUMN ultimo_semestre_cursado NUMERIC(5,0) NULL;
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.ultimo_semestre_cursado IS 'Último semestre cursado por el estudiante de la carreara de la que proviene.';