-- object: organizaciones | type: SCHEMA --
-- DROP SCHEMA IF EXISTS organizaciones CASCADE;
CREATE SCHEMA organizaciones;
-- ddl-end --

SET search_path TO pg_catalog,public,organizaciones;
-- ddl-end --

-- object: organizaciones.organizacion | type: TABLE --
-- DROP TABLE IF EXISTS organizaciones.organizacion CASCADE;
CREATE TABLE organizaciones.organizacion (
	id serial NOT NULL,
	nombre character varying(200) NOT NULL,
	ente integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_organizacion integer NOT NULL,
	CONSTRAINT pk_organizacion PRIMARY KEY (id),
	CONSTRAINT uq_ente_organizacion UNIQUE (ente)

)
WITH ( OIDS = TRUE );
-- ddl-end --
COMMENT ON TABLE organizaciones.organizacion IS 'Tabla que almacena la información básica general de las organizaciones';
-- ddl-end --
COMMENT ON COLUMN organizaciones.organizacion.id IS 'Identificador de la tabla organizacion';
-- ddl-end --
COMMENT ON COLUMN organizaciones.organizacion.nombre IS 'Nombre de la organizacion';
-- ddl-end --
COMMENT ON COLUMN organizaciones.organizacion.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN organizaciones.organizacion.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN organizaciones.organizacion.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT uq_ente_organizacion ON organizaciones.organizacion  IS 'Restricción unique para el ente de la organizacion';
-- ddl-end --

-- object: organizaciones.relacion_organizaciones | type: TABLE --
-- DROP TABLE IF EXISTS organizaciones.relacion_organizaciones CASCADE;
CREATE TABLE organizaciones.relacion_organizaciones (
	id serial NOT NULL,
	organizacion_padre integer NOT NULL,
	organizacion_hija integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_relacion_organizaciones integer NOT NULL,
	CONSTRAINT pk_relacion_organizaciones PRIMARY KEY (id),
	CONSTRAINT uq_organizacion_hija UNIQUE (organizacion_hija),
	CONSTRAINT uq_organizacion_padre UNIQUE (organizacion_padre)

);
-- ddl-end --
COMMENT ON COLUMN organizaciones.relacion_organizaciones.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN organizaciones.relacion_organizaciones.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN organizaciones.relacion_organizaciones.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --

-- object: organizaciones.tipo_organizacion | type: TABLE --
-- DROP TABLE IF EXISTS organizaciones.tipo_organizacion CASCADE;
CREATE TABLE organizaciones.tipo_organizacion (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_organizacion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_organizacion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE organizaciones.tipo_organizacion IS 'Tabla paramétrica que almacena los tipos de organización';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.id IS 'Identificador de la tabla tipo_organizacion';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro de tipo_organizacion';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.activo IS 'Campo de tipo boolean que indica si el parámetro está activo';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_organizacion.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_organizacion ON organizaciones.tipo_organizacion  IS 'Llave primaria de la tabla tipo_organizacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_organizacion ON organizaciones.tipo_organizacion  IS 'Constraint unique para que el nombre de tipo_organizacion no se repita';
-- ddl-end --

-- object: organizaciones.tipo_relacion_organizaciones | type: TABLE --
-- DROP TABLE IF EXISTS organizaciones.tipo_relacion_organizaciones CASCADE;
CREATE TABLE organizaciones.tipo_relacion_organizaciones (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_relacion_organizaciones PRIMARY KEY (id),
	CONSTRAINT uq_tipo_relacion_organizaciones_nombre UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_relacion_organizaciones.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN organizaciones.tipo_relacion_organizaciones.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --

-- object: fk_organizacion_tipo_organizacion | type: CONSTRAINT --
-- ALTER TABLE organizaciones.organizacion DROP CONSTRAINT IF EXISTS fk_organizacion_tipo_organizacion CASCADE;
ALTER TABLE organizaciones.organizacion ADD CONSTRAINT fk_organizacion_tipo_organizacion FOREIGN KEY (tipo_organizacion)
REFERENCES organizaciones.tipo_organizacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_relacion_organizaciones_tipo_relacion_organizaciones | type: CONSTRAINT --
-- ALTER TABLE organizaciones.relacion_organizaciones DROP CONSTRAINT IF EXISTS fk_relacion_organizaciones_tipo_relacion_organizaciones CASCADE;
ALTER TABLE organizaciones.relacion_organizaciones ADD CONSTRAINT fk_relacion_organizaciones_tipo_relacion_organizaciones FOREIGN KEY (tipo_relacion_organizaciones)
REFERENCES organizaciones.tipo_relacion_organizaciones (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_relacion_organizaciones_organizacion_padre | type: CONSTRAINT --
-- ALTER TABLE organizaciones.relacion_organizaciones DROP CONSTRAINT IF EXISTS fk_relacion_organizaciones_organizacion_padre CASCADE;
ALTER TABLE organizaciones.relacion_organizaciones ADD CONSTRAINT fk_relacion_organizaciones_organizacion_padre FOREIGN KEY (organizacion_padre)
REFERENCES organizaciones.organizacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_relacion_organizaciones_organizacion_hija | type: CONSTRAINT --
-- ALTER TABLE organizaciones.relacion_organizaciones DROP CONSTRAINT IF EXISTS fk_relacion_organizaciones_organizacion_hija CASCADE;
ALTER TABLE organizaciones.relacion_organizaciones ADD CONSTRAINT fk_relacion_organizaciones_organizacion_hija FOREIGN KEY (organizacion_hija)
REFERENCES organizaciones.organizacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --