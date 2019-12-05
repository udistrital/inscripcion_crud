
CREATE SEQUENCE sq_tipo_icfes_id;
ALTER TABLE inscripcion.tipo_icfes ALTER COLUMN id SET DEFAULT nextval('sq_tipo_icfes_id');

insert into inscripcion.tipo_icfes (nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values ('AC - Normal', 'Icfes Normal', 'AC', true, 1, now(), now());
insert into inscripcion.tipo_icfes (nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values ('VG - Validante', 'Icfes de validante', 'VG', true, 2, now(), now());
