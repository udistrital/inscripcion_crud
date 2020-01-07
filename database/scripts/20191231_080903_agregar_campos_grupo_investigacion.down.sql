ALTER TABLE inscripcion.propuesta 
	DROP COLUMN grupo_investigacion_id;
ALTER TABLE inscripcion.propuesta 
	DROP COLUMN linea_investigacion_id;
ALTER TABLE inscripcion.propuesta
    ADD COLUMN grupo_investigacion_linea_invetigacion_id INTEGER NOT NULL DEFAULT 0;