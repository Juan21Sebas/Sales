-- +goose Up
-- +goose StatementBegin

-- Tabla para módulos de encuesta
CREATE TABLE survey_modules (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    display_order INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabla para preguntas de encuesta
CREATE TABLE survey_questions (
    id SERIAL PRIMARY KEY,
    module_id INTEGER NOT NULL,
    question_text TEXT NOT NULL,
    field_name VARCHAR(255) NOT NULL,
    field_type VARCHAR(50) NOT NULL CHECK (field_type IN (
        'selection', 'country_selector', 'textfield', 'picker',
        'slider', 'text_to_chip', 'chip_selection'
    )),
    field_options JSONB,
    display_notes TEXT,
    is_required BOOLEAN NOT NULL DEFAULT TRUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    display_order INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_survey_questions_module FOREIGN KEY (module_id) 
        REFERENCES survey_modules(id) ON DELETE CASCADE
);

-- Tabla para opciones de preguntas
CREATE TABLE survey_question_options (
    id SERIAL PRIMARY KEY,
    question_id INTEGER NOT NULL,
    option_text VARCHAR(255) NOT NULL,
    option_value VARCHAR(255) NOT NULL,
    display_order INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_survey_question_options_question FOREIGN KEY (question_id) 
        REFERENCES survey_questions(id) ON DELETE CASCADE
);

-- Tabla para respuestas de usuarios
CREATE TABLE survey_responses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    option_id INTEGER,
    response_value TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_survey_responses_question FOREIGN KEY (question_id) 
        REFERENCES survey_questions(id) ON DELETE CASCADE,
    CONSTRAINT fk_survey_responses_option FOREIGN KEY (option_id) 
        REFERENCES survey_question_options(id) ON DELETE SET NULL
);

-- Índices para optimización
CREATE INDEX idx_survey_modules_is_active ON survey_modules(is_active);
CREATE INDEX idx_survey_modules_display_order ON survey_modules(display_order);
CREATE INDEX idx_survey_questions_module_id ON survey_questions(module_id);
CREATE INDEX idx_survey_questions_is_active ON survey_questions(is_active);
CREATE INDEX idx_survey_questions_display_order ON survey_questions(display_order);
CREATE INDEX idx_survey_question_options_question_id ON survey_question_options(question_id);
CREATE INDEX idx_survey_question_options_is_active ON survey_question_options(is_active);
CREATE INDEX idx_survey_question_options_display_order ON survey_question_options(display_order);
CREATE INDEX idx_survey_responses_user_id ON survey_responses(user_id);
CREATE INDEX idx_survey_responses_question_id ON survey_responses(question_id);
CREATE INDEX idx_survey_responses_option_id ON survey_responses(option_id);

-- Función para actualización automática de timestamps
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ language 'plpgsql';

-- Triggers para actualización automática
CREATE TRIGGER update_survey_modules_timestamp
BEFORE UPDATE ON survey_modules
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_survey_questions_timestamp
BEFORE UPDATE ON survey_questions
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_survey_question_options_timestamp
BEFORE UPDATE ON survey_question_options
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_survey_responses_timestamp
BEFORE UPDATE ON survey_responses
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

-- Comentarios descriptivos
COMMENT ON TABLE survey_modules IS 'Módulos de la encuesta (Profile, Experience, Preferences)';
COMMENT ON TABLE survey_questions IS 'Preguntas asociadas a cada módulo';
COMMENT ON TABLE survey_question_options IS 'Opciones para preguntas de selección';
COMMENT ON TABLE survey_responses IS 'Respuestas de los usuarios a las preguntas';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_survey_responses_timestamp ON survey_responses;
DROP TRIGGER IF EXISTS update_survey_question_options_timestamp ON survey_question_options;
DROP TRIGGER IF EXISTS update_survey_questions_timestamp ON survey_questions;
DROP TRIGGER IF EXISTS update_survey_modules_timestamp ON survey_modules;

DROP FUNCTION IF EXISTS update_timestamp;

DROP TABLE IF EXISTS survey_responses;
DROP TABLE IF EXISTS survey_question_options;
DROP TABLE IF EXISTS survey_questions;
DROP TABLE IF EXISTS survey_modules;
-- +goose StatementEnd