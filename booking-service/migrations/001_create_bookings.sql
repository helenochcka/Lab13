CREATE TABLE bookings (
                          id UUID PRIMARY KEY,
                          account_id UUID NOT NULL,
                          event_id UUID NOT NULL,
                          status TEXT NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
