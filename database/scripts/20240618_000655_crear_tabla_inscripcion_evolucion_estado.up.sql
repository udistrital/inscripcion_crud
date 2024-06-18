CREATE TABLE inscripcion.inscripcion_evolucion_estado(
	id bigserial NOT NULL,
	tercero_id integer NOT NULL,
	inscripcion_id integer NOT NULL,
	estado_inscripcion_id_anterior integer,
	estado_inscripcion_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_inscripcion_evolucion_estado PRIMARY KEY (id)
);

ALTER TABLE inscripcion.inscripcion_evolucion_estado ADD CONSTRAINT fk_inscripcion_evolucion_estado_inscripcion FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE inscripcion.inscripcion_evolucion_estado ADD CONSTRAINT fk_inscripcion_evolucion_estado_inscripcion_anterior FOREIGN KEY (estado_inscripcion_id_anterior)
REFERENCES inscripcion.estado_inscripcion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;