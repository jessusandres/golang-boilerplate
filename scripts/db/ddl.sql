DROP TABLE IF EXISTS incidents;

CREATE TABLE incidents
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR NOT NULL,
    description   TEXT,
    incident_type VARCHAR,
    location      VARCHAR,
    image         VARCHAR,
    event_date    DATE,
    created_at    TIMESTAMPTZ,
    updated_at    TIMESTAMPTZ
);

-- Demo data: 10 records without images
INSERT INTO incidents (title, description, incident_type, location, image, event_date, created_at, updated_at)
VALUES ('Power outage in Sector 7', 'Reported blackout affecting multiple blocks in Sector 7. Utility crew dispatched.',
        'Power', 'Sector 7, Neo City', NULL, NOW(), NOW(), NOW()),
       ('Water main break', 'Significant water leak causing low pressure in surrounding area.', 'Water',
        '5th Ave & Pine St', NULL, NOW(), NOW(), NOW()),
       ('Road blockage due to fallen tree', 'Large tree has fallen blocking both lanes. Public works en route.',
        'Infrastructure', 'Oak Street Bridge', NULL, NOW(), NOW(), NOW()),
       ('Minor warehouse fire', 'Contained fire in storage area. No injuries reported. Fire dept monitoring hotspots.',
        'Fire', 'Harbor Industrial Park', NULL, NOW(), NOW(), NOW()),
       ('Transit delay on Line A', 'Signal issue causing delays up to 20 minutes. Repairs in progress.', 'Transit',
        'Line A - Central Station', NULL, NOW(), NOW(), NOW()),
       ('Network outage', 'ISP reports regional network disruption. Estimated resolution within 2 hours.', 'IT',
        'Downtown district', NULL, NOW(), NOW(), NOW()),
       ('Gas leak reported', 'Residents report odor of gas. Area cordoned off as precaution.', 'Hazard',
        'Maple St & 12th', NULL, NOW(), NOW(), NOW()),
       ('Severe weather alert', 'Thunderstorm with heavy rainfall expected. Flood-prone zones on watch.', 'Weather',
        'Metro Area', NULL, NOW(), NOW(), NOW()),
       ('Traffic accident', 'Multi-vehicle collision. Emergency services on site. Expect delays.', 'Traffic',
        'I-95 Northbound, MM 142', NULL, NOW(), NOW(), NOW()),
       ('Elevator outage', 'Elevator offline for maintenance. Use stairs or alternate elevator.', 'Maintenance',
        'Civic Center Building B', NULL, NOW(), NOW(), NOW());
