CREATE TABLE inscripcion.cupo_inscripcion (
	id SERIAL NOT NULL,
	cupos_habilitados integer NOT NULL,
	cupos_opcionados integer NOT NULL,
	periodo_id integer NOT NULL,
	proyecto_academico_id integer NOT NULL,
	tipo_inscripcion_id integer NOT NULL,
	cupo_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_cupo_inscripcion PRIMARY KEY (id)
);

CREATE TABLE inscripcion.documento_cupo (
	id SERIAL NOT NULL,
	uid character varying(250) NOT NULL,
	comentario character varying(250) NOT NULL,
	cupo_inscripcion_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_documento_cupo PRIMARY KEY (id)
);

ALTER TABLE inscripcion.cupo_inscripcion ADD CONSTRAINT fk_tipo_inscripcion_cupo_inscripcion FOREIGN KEY (tipo_inscripcion_id)
REFERENCES inscripcion.tipo_inscripcion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;

ALTER TABLE inscripcion.documento_cupo ADD CONSTRAINT fk_cupo_inscripcion_documento_cupo FOREIGN KEY (cupo_inscripcion_id)
REFERENCES inscripcion.cupo_inscripcion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
