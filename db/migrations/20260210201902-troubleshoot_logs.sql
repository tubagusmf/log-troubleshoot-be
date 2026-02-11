
-- +migrate Up
CREATE TABLE troubleshoot_logs (
    id SERIAL PRIMARY KEY,
    ticket_number VARCHAR(50) UNIQUE,
    trouble_date DATE NOT NULL,
    trouble_time TIME NOT NULL,
    done_date DATE,
    done_time TIME,
    duration INTERVAL,
    user_id INT REFERENCES users(id),
    project_id INT REFERENCES projects(id),
    location_id INT REFERENCES locations(id),
    device_id INT REFERENCES devices(id),
    work_type_id INT REFERENCES work_types(id),
    device_number VARCHAR(20),
    part VARCHAR(100),
    issue TEXT,                      
    solution TEXT,
    status VARCHAR(20) DEFAULT 'OPEN',
    whatsapp_sender VARCHAR(50),
    whatsapp_message TEXT,
    sheet_id VARCHAR(120),
    sheet_row INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS troubleshoot_logs;
