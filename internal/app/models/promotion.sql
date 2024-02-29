CREATE TABLE promotion_table (
  id SERIAL PRIMARY KEY,
  promotion_id VARCHAR(10) NOT NULL UNIQUE,
  promotion_name VARCHAR(255) NOT NULL,
  discount_type VARCHAR(50) NOT NULL,
  discount_value NUMERIC(10,2) NOT NULL,
  promotion_start_date TIMESTAMP WITH TIME ZONE NOT NULL,
  promotion_end_date TIMESTAMP WITH TIME ZONE NOT null,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP WITH TIME ZONE
);

INSERT INTO promotion_table (promotion_id, promotion_name, discount_type, discount_value, promotion_start_date, promotion_end_date)
VALUES
    ('cae8651b', 'Ramadhan Sale', 'Percentage', 10.50, '2024-03-01 00:00:00', '2024-03-31 23:59:59'),
    ('137ce1cf', 'New Year Sale', 'Fixed', 20.00, '2024-03-05 00:00:00', '2024-04-05 23:59:59'),
    ('dc6ef7bf', 'White Christmas', 'Percentage', 15.75, '2024-03-10 00:00:00', '2024-04-10 23:59:59'),
    ('e1d63455', 'Idul Fitri Day', 'Fixed', 30.00, '2024-03-15 00:00:00', '2024-04-15 23:59:59'),
    ('c373b046', 'Idul Adha', 'Percentage', 12.25, '2024-03-20 00:00:00', '2024-04-20 23:59:59'),
    ('1c5b25c1', 'Independence Day', 'Fixed', 25.00, '2024-03-25 00:00:00', '2024-04-25 23:59:59'),
    ('7cb19c9d', 'Special 70 Days', 'Percentage', 8.75, '2024-03-30 00:00:00', '2024-04-30 23:59:59'),
    ('e0206bb0', 'Warrior Day', 'Fixed', 35.00, '2024-04-01 00:00:00', '2024-05-01 23:59:59'),
    ('c78221e3', 'Summer Day', 'Percentage', 18.50, '2024-04-05 00:00:00', '2024-05-05 23:59:59'),
    ('7791abc3', 'Winter Day', 'Fixed', 40.00, '2024-04-10 00:00:00', '2024-05-10 23:59:59');