CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username TEXT UNIQUE NOT NULL,
                                     password TEXT NOT NULL,
                                     role TEXT NOT NULL CHECK (role IN ('admin', 'user'))
    );

INSERT INTO users (username, password, role) VALUES
                                        ('admin', '$2a$10$Daqs1UJH2WPPAoYDj4wwl.HLP4cGevCH8D8ux1GcWX1zX3Y65Dt3O', 'admin'),
                                        ('user', '$2a$10$JXppXPEnzJJT1vR98TK8Pu.5Iv7PlFsE1IjBMbN/Xv1lVsF627sbK', 'user');
