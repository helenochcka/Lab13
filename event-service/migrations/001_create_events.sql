CREATE TABLE events (
                        id UUID PRIMARY KEY,
                        title TEXT NOT NULL,
                        total_seats INT NOT NULL,
                        available_seats INT NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
