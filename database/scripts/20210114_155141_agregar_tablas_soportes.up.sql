
-- object: inscripcion.tipo_documento_programa | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.tipo_documento_programa CASCADE;
CREATE TABLE inscripcion.tipo_documento_programa (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20) NOT NULL,
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_documento_programa PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.nombre IS E'Nombre del tipo de documento requerido por el programa';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.descripcion IS E'Descripcion del tipo de documento requerido por el programa';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.codigo_abreviacion IS E'Codigo de abreviacion del tipo de documento';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.activo IS E'Marcador que indica si el tipo de documento se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.numero_orden IS E'Numero de orden del tipo de documento';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.fecha_creacion IS E'Fecha de creacion del tipo de documento en la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_documento_programa.fecha_modificacion IS E'Fecha de modificacion del tipo de documento en la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_documento_programa ON inscripcion.tipo_documento_programa  IS E'Identificador unico de cada registro de la tabla';
-- ddl-end --



-- object: inscripcion.documento_programa | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.documento_programa CASCADE;
CREATE TABLE inscripcion.documento_programa (
	id serial NOT NULL,
	activo boolean NOT NULL,
	periodo_id integer NOT NULL,
	programa_id integer,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_documento_programa_id integer NOT NULL,
	CONSTRAINT pk_documento_programa PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.documento_programa IS E'Tabla que almacena los documentos requeridos por el programa para la inscripcion';
-- ddl-end --
COMMENT ON COLUMN inscripcion.documento_programa.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.documento_programa.activo IS E'Marcador que indica si el documento requerido por el programa se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.documento_programa.periodo_id IS E'Identificador del periodo academico para el que se solicita el documento';
-- ddl-end --
COMMENT ON COLUMN inscripcion.documento_programa.programa_id IS E'identificador del programa academico que tiene este documento como requisito';
-- ddl-end --
COMMENT ON COLUMN inscripcion.documento_programa.fecha_creacion IS E'Fecha de registro del documento requerido por el programa';
-- ddl-end --
COMMENT ON COLUMN inscripcion.documento_programa.fecha_modificacion IS E'Fecha de modificacion del documento requerido por el programa';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento_programa ON inscripcion.documento_programa  IS E'Identificadorúnico de cada registro de documento solicitado por el programa';
-- ddl-end --



-- object: inscripcion.soporte_documento_programa | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.soporte_documento_programa CASCADE;
CREATE TABLE inscripcion.soporte_documento_programa (
	id serial NOT NULL,
	documento_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	documento_programa_id integer NOT NULL,
	inscripcion_id integer NOT NULL,
	CONSTRAINT pk_soporte_documento_programa PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.soporte_documento_programa IS E'Esta tabla almacena los documentos relacionados en la inscripcion';
-- ddl-end --
COMMENT ON COLUMN inscripcion.soporte_documento_programa.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.soporte_documento_programa.documento_id IS E'Identificador del documento relacionado en la api de Documentos';
-- ddl-end --
COMMENT ON COLUMN inscripcion.soporte_documento_programa.fecha_creacion IS E'Fecha de creacion del registro del soporte';
-- ddl-end --
COMMENT ON COLUMN inscripcion.soporte_documento_programa.fecha_modificacion IS E'Fecha de modificacion del registro del soporte';
-- ddl-end --
COMMENT ON CONSTRAINT pk_soporte_documento_programa ON inscripcion.soporte_documento_programa  IS E'Identificador unico de cada registro de la tabla';
-- ddl-end --


-- object: fk_tipo_documento_programa | type: CONSTRAINT --
-- ALTER TABLE inscripcion.documento_programa DROP CONSTRAINT IF EXISTS fk_tipo_documento_programa CASCADE;
ALTER TABLE inscripcion.documento_programa ADD CONSTRAINT fk_tipo_documento_programa FOREIGN KEY (tipo_documento_programa_id)
REFERENCES inscripcion.tipo_documento_programa (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_documento_programa | type: CONSTRAINT --
-- ALTER TABLE inscripcion.soporte_documento_programa DROP CONSTRAINT IF EXISTS fk_documento_programa CASCADE;
ALTER TABLE inscripcion.soporte_documento_programa ADD CONSTRAINT fk_documento_programa FOREIGN KEY (documento_programa_id)
REFERENCES inscripcion.documento_programa (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_inscripcion | type: CONSTRAINT --
-- ALTER TABLE inscripcion.soporte_documento_programa DROP CONSTRAINT IF EXISTS fk_inscripcion CASCADE;
ALTER TABLE inscripcion.soporte_documento_programa ADD CONSTRAINT fk_inscripcion FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

