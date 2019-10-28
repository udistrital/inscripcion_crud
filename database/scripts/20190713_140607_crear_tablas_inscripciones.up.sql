-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2-alpha1
-- PostgreSQL version: 11.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- object: inscripcion | type: SCHEMA --
DROP SCHEMA IF EXISTS inscripcion CASCADE;
CREATE SCHEMA inscripcion;
-- ddl-end --

-- SET search_path TO pg_catalog,public,inscripcion;
-- ddl-end --

-- object: inscripcion.inscripcion | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.inscripcion CASCADE;
CREATE TABLE inscripcion.inscripcion (
	id serial NOT NULL,
	persona_id integer NOT NULL,
	programa_academico_id integer NOT NULL,
	recibo_matricula_id integer,
	recibo_inscripcion_id integer,
	periodo_id integer NOT NULL,
	enfasis_id integer,
	acepta_terminos boolean NOT NULL,
	fecha_acepta_terminos date NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	estado_inscripcion_id integer NOT NULL,
	tipo_inscripcion_id integer NOT NULL,
	CONSTRAINT pk_inscripcion PRIMARY KEY (id),
	CONSTRAINT uq_aspirante UNIQUE (persona_id,programa_academico_id,periodo_id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.inscripcion IS 'Tabla que almacena la información del proceso de inscripción y admisión de un aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.id IS 'Identificador de la tabla de admisión';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.programa_academico_id IS 'Programa academico al que se presenta el aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.recibo_matricula_id IS 'Recibo de matricula del aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.recibo_inscripcion_id IS 'Recibo de inscripción del aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.periodo_id IS 'Periodo en el que se presenta el  aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.enfasis_id IS 'Enfasis del proyecto curricular escogido por el aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.acepta_terminos IS 'Flag que indica si el aspirante acepta terminos y condiciones';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.fecha_acepta_terminos IS 'Fecha en la que el aspirante acepta los terminos.';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.activo IS 'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.fecha_creacion IS 'Fecha de creacion de una inscripcion';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion.fecha_modificacion IS 'Fecha de modifiación de una inscripcion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_inscripcion ON inscripcion.inscripcion  IS 'Llave primaria';
-- ddl-end --
COMMENT ON CONSTRAINT uq_aspirante ON inscripcion.inscripcion  IS 'Restricción para que un aspirante no pueda registrarse a un mismo programa académico durante el periodo académico';
-- ddl-end --

-- object: inscripcion.estado_inscripcion | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.estado_inscripcion CASCADE;
CREATE TABLE inscripcion.estado_inscripcion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20) NOT NULL,
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_estado_inscripicion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.estado_inscripcion IS 'Tabla que almacena los diferentes estados de admision';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.nombre IS 'Nombre del estado de admision';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.descripcion IS 'Descripción del estado de admisión';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.codigo_abreviacion IS 'Codigo de abreviación del estado de admisión';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.activo IS 'Flag que indica si el estado de admisión esta activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.numero_orden IS 'Número de orden en que se deben mostrar los estados de admisión';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.fecha_creacion IS 'Fecha de creacion de un  estado inscripcion';
-- ddl-end --
COMMENT ON COLUMN inscripcion.estado_inscripcion.fecha_modificacion IS 'Fecha de modifiación de un estado inscripicion';
-- ddl-end --

-- object: inscripcion.propuesta | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.propuesta CASCADE;
CREATE TABLE inscripcion.propuesta (
	id serial NOT NULL,
	nombre character varying(250) NOT NULL,
	resumen text,
	grupo_investigacion_linea_invetigacion_id integer NOT NULL,
	documento_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	inscripcion_id integer NOT NULL,
	tipo_proyecto_id integer NOT NULL,
	CONSTRAINT pk_propuesta PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.propuesta IS 'Propuesta que presenta el aspirante.';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.nombre IS 'Nombre de la propuesta';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.resumen IS 'Resumen  de la propuesta';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.grupo_investigacion_linea_invetigacion_id IS 'Grupo de investigacion al que pertenece el proyecto.';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.documento_id IS 'Formato del proyecto con el que el estudiante presenta la propuesta';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.activo IS 'Flag que indica si la propuesta esta activa o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.fecha_creacion IS 'Fecha de creacion de una propuesta';
-- ddl-end --
COMMENT ON COLUMN inscripcion.propuesta.fecha_modificacion IS 'Fecha de modifiación de una propuesta';
-- ddl-end --

-- object: inscripcion.tipo_proyecto | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.tipo_proyecto CASCADE;
CREATE TABLE inscripcion.tipo_proyecto (
	id integer NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_proyecto PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.tipo_proyecto IS 'Tabla que almacena los tipos de proyectos de grado';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.nombre IS 'Nombre del tipo de proyecto';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.descripcion IS 'Descripcion del tipo de proyecto';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.codigo_abreviacion IS 'Codigo de abreviacion del tipo de proyecto';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.activo IS 'Flag que indica si el tipo de proyecto se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.numero_orden IS 'Numero de orden del tipo de trabajo de proyecto';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_proyecto.fecha_creacion IS 'Fecha de creacion de un tipo proyecto';
-- ddl-end --

-- object: fk_inscripcion_estado_inscripcion | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion DROP CONSTRAINT IF EXISTS fk_inscripcion_estado_inscripcion CASCADE;
ALTER TABLE inscripcion.inscripcion ADD CONSTRAINT fk_inscripcion_estado_inscripcion FOREIGN KEY (estado_inscripcion_id)
REFERENCES inscripcion.estado_inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_propuesta_inscripcion | type: CONSTRAINT --
-- ALTER TABLE inscripcion.propuesta DROP CONSTRAINT IF EXISTS fk_propuesta_inscripcion CASCADE;
ALTER TABLE inscripcion.propuesta ADD CONSTRAINT fk_propuesta_inscripcion FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_tipo_proyecto_propuesta | type: CONSTRAINT --
-- ALTER TABLE inscripcion.propuesta DROP CONSTRAINT IF EXISTS fk_tipo_proyecto_propuesta CASCADE;
ALTER TABLE inscripcion.propuesta ADD CONSTRAINT fk_tipo_proyecto_propuesta FOREIGN KEY (tipo_proyecto_id)
REFERENCES inscripcion.tipo_proyecto (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: inscripcion.tipo_inscripcion | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.tipo_inscripcion CASCADE;
CREATE TABLE inscripcion.tipo_inscripcion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20) NOT NULL,
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_inscripicion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_inscripcion.fecha_creacion IS 'Fecha de creacion de un  tipo de inscripcion';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_inscripcion.fecha_modificacion IS 'Fecha de modifiación de un tipo de inscripcion';
-- ddl-end --

-- object: fk_tipo_inscripcion_inscripcion | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion DROP CONSTRAINT IF EXISTS fk_tipo_inscripcion_inscripcion CASCADE;
ALTER TABLE inscripcion.inscripcion ADD CONSTRAINT fk_tipo_inscripcion_inscripcion FOREIGN KEY (tipo_inscripcion_id)
REFERENCES inscripcion.tipo_inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: inscripcion.transferencia | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.transferencia CASCADE;
CREATE TABLE inscripcion.transferencia (
	id serial NOT NULL,
	inscripcion_id integer NOT NULL,
	transferencia_interna boolean NOT NULL,
	codigo_estudiante_proviene varchar(50),
	universidad_proviene varchar(250) NOT NULL,
	proyecto_curricular_proviene varchar(250) NOT NULL,
	ultimo_semestre_cursado numeric(5) NOT NULL,
	motivo_retiro text NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_transferencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.transferencia IS 'Tabla que almacena informacion de las transferencias';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.transferencia_interna IS 'Flag que indica si la transferencia es interna o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.codigo_estudiante_proviene IS 'Código del estudiante de la carreara donde proviene';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.universidad_proviene IS 'Universidad de la que proviene el estudiante que solicita la transferencia.';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.proyecto_curricular_proviene IS 'Carrera de la que proviene el aspirante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.ultimo_semestre_cursado IS 'Último semestre cursado por el estudiante de la carreara de la que proviene.';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.motivo_retiro IS 'Motivo de retiro del estudiante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.activo IS 'Flag que indica si el registro es activo';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.fecha_creacion IS 'Fecha de creación de la transferencia';
-- ddl-end --
COMMENT ON COLUMN inscripcion.transferencia.fecha_modificacion IS 'Fecha de modificación de la transferencia';
-- ddl-end --

-- object: inscripcion.reintegro | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.reintegro CASCADE;
CREATE TABLE inscripcion.reintegro (
	id serial NOT NULL,
	codigo_estudiante numeric(11) NOT NULL,
	cancelo_semestre boolean NOT NULL,
	motivo_retiro text NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	inscripcion_id integer NOT NULL,
	CONSTRAINT pk_reintegro PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.reintegro IS 'Tabla que almacena informacion de los reintegros';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.codigo_estudiante IS 'Codigo del estudiante ';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.cancelo_semestre IS 'Flag que indica si el aspirante cancelo el semestre';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.motivo_retiro IS 'Motivo por el cuael estudiante se retiro';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.activo IS 'Flag que indica si el registro esta activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.fecha_creacion IS 'Fecha de creación del reintegro';
-- ddl-end --
COMMENT ON COLUMN inscripcion.reintegro.fecha_modificacion IS 'Fecha de modificación del reintegro';
-- ddl-end --

-- object: fk_inscripcion_transferencia | type: CONSTRAINT --
-- ALTER TABLE inscripcion.transferencia DROP CONSTRAINT IF EXISTS fk_inscripcion_transferencia CASCADE;
ALTER TABLE inscripcion.transferencia ADD CONSTRAINT fk_inscripcion_transferencia FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_inscripcion_transferencia | type: CONSTRAINT --
-- ALTER TABLE inscripcion.transferencia DROP CONSTRAINT IF EXISTS uq_inscripcion_transferencia CASCADE;
ALTER TABLE inscripcion.transferencia ADD CONSTRAINT uq_inscripcion_transferencia UNIQUE (inscripcion_id);
-- ddl-end --

-- object: fk_inscripcion_reintegro | type: CONSTRAINT --
-- ALTER TABLE inscripcion.reintegro DROP CONSTRAINT IF EXISTS fk_inscripcion_reintegro CASCADE;
ALTER TABLE inscripcion.reintegro ADD CONSTRAINT fk_inscripcion_reintegro FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_inscripcion_reintegro | type: CONSTRAINT --
-- ALTER TABLE inscripcion.reintegro DROP CONSTRAINT IF EXISTS uq_inscripcion_reintegro CASCADE;
ALTER TABLE inscripcion.reintegro ADD CONSTRAINT uq_inscripcion_reintegro UNIQUE (inscripcion_id);
-- ddl-end --

-- object: inscripcion.inscripcion_pregrado | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.inscripcion_pregrado CASCADE;
CREATE TABLE inscripcion.inscripcion_pregrado (
	id serial NOT NULL,
	inscripcion_id integer NOT NULL,
	codigo_icfes varchar(15) NOT NULL,
	tipo_documento_icfes integer NOT NULL,
	numero_identificacion_icfes numeric(11) NOT NULL,
	ano_icfes numeric(4) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_icfes_id integer NOT NULL,
	CONSTRAINT pk_inscripcion_pregrado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.inscripcion_pregrado IS 'Tabla que almacena informacion de las inscripciones de pregrado';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.codigo_icfes IS 'Código del icfes del estudiante';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.tipo_documento_icfes IS 'Tipo de documento con el que el aspirante presento el icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.numero_identificacion_icfes IS 'Número de identificación con el qeu el estudiante  presentó el icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.ano_icfes IS 'Año en el que el aspirante presento el icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.activo IS 'Flag que indica si el registro es activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.fecha_creacion IS 'Fecha de creación del registro de inscripción';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_pregrado.fecha_modificacion IS 'Fecha de modificación del registro de inscripción';
-- ddl-end --

-- object: fk_inscripcion_inscripcion_pregrado | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion_pregrado DROP CONSTRAINT IF EXISTS fk_inscripcion_inscripcion_pregrado CASCADE;
ALTER TABLE inscripcion.inscripcion_pregrado ADD CONSTRAINT fk_inscripcion_inscripcion_pregrado FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_inscripcion_inscripcion_pregrado | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion_pregrado DROP CONSTRAINT IF EXISTS uq_inscripcion_inscripcion_pregrado CASCADE;
ALTER TABLE inscripcion.inscripcion_pregrado ADD CONSTRAINT uq_inscripcion_inscripcion_pregrado UNIQUE (inscripcion_id);
-- ddl-end --

-- object: inscripcion.tipo_icfes | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.tipo_icfes CASCADE;
CREATE TABLE inscripcion.tipo_icfes (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_icfes PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.tipo_icfes IS 'Tabla que almacena los tipos de icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.nombre IS 'Nombre del tipo de icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.descripcion IS 'Descripcion del tipo de icfes
';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.codigo_abreviacion IS 'Codigo de abreviacion del tipo de icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.activo IS 'Flag que indica si el tipo de icfes se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.numero_orden IS 'Numero de orden del tipo de icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.fecha_creacion IS 'Fecha de creacion de un tipo icfes';
-- ddl-end --
COMMENT ON COLUMN inscripcion.tipo_icfes.fecha_modificacion IS 'Fecha de modificación de un tipo icfes';
-- ddl-end --

-- object: fk_tipo_icfes_inscripcion_pregrado | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion_pregrado DROP CONSTRAINT IF EXISTS fk_tipo_icfes_inscripcion_pregrado CASCADE;
ALTER TABLE inscripcion.inscripcion_pregrado ADD CONSTRAINT fk_tipo_icfes_inscripcion_pregrado FOREIGN KEY (tipo_icfes_id)
REFERENCES inscripcion.tipo_icfes (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: inscripcion.inscripcion_posgrado | type: TABLE --
-- DROP TABLE IF EXISTS inscripcion.inscripcion_posgrado CASCADE;
CREATE TABLE inscripcion.inscripcion_posgrado (
	id serial NOT NULL,
	idioma integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	inscripcion_id integer NOT NULL,
	CONSTRAINT pk_inscripcion_posgrado PRIMARY KEY (id),
	CONSTRAINT uq_idioma UNIQUE (idioma)

);
-- ddl-end --
COMMENT ON TABLE inscripcion.inscripcion_posgrado IS 'Tabla que almacena informacion de las inscripciones de posgrado';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_posgrado.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_posgrado.idioma IS 'Campo que relaciona la tabla parametrica idioma del esquema idiomas';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_posgrado.activo IS 'Flag que indica si el registro es activo o no';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_posgrado.fecha_creacion IS 'Fecha de creación del registro de inscripción';
-- ddl-end --
COMMENT ON COLUMN inscripcion.inscripcion_posgrado.fecha_modificacion IS 'Fecha de modificación del registro de inscripción';
-- ddl-end --
COMMENT ON CONSTRAINT uq_idioma ON inscripcion.inscripcion_posgrado  IS 'Constraint para el campo idioma pues el aspirante solo puede seleccionar un unico idioma para realizar la prueba';
-- ddl-end --

-- object: fk_inscripcion_inscripcion_posgrado | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion_posgrado DROP CONSTRAINT IF EXISTS fk_inscripcion_inscripcion_posgrado CASCADE;
ALTER TABLE inscripcion.inscripcion_posgrado ADD CONSTRAINT fk_inscripcion_inscripcion_posgrado FOREIGN KEY (inscripcion_id)
REFERENCES inscripcion.inscripcion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_inscripcion_posgrado | type: CONSTRAINT --
-- ALTER TABLE inscripcion.inscripcion_posgrado DROP CONSTRAINT IF EXISTS uq_inscripcion_posgrado CASCADE;
ALTER TABLE inscripcion.inscripcion_posgrado ADD CONSTRAINT uq_inscripcion_posgrado UNIQUE (inscripcion_id);
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA inscripcion TO test;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA inscripcion TO test;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA inscripcion TO test;
