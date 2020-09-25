ALTER TABLE inscripcion.propuesta 
	ADD COLUMN grupo_investigacion_id INTEGER NOT NULL DEFAULT 0;
ALTER TABLE inscripcion.propuesta 
	ADD COLUMN linea_investigacion_id INTEGER NOT NULL DEFAULT 0;
ALTER TABLE inscripcion.propuesta
    DROP COLUMN grupo_investigacion_linea_invetigacion_id;