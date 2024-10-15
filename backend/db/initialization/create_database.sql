-- Create database if it does not exist
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'myappdb') THEN EXECUTE 'CREATE DATABASE myappdb';
        RAISE NOTICE 'Database myappdb created.';
    ELSE
        RAISE NOTICE 'Database myappdb already exists.';
    END IF;
END $$;

-- Create user if it does not exist
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'myuser') THEN
        EXECUTE 'CREATE USER myuser WITH ENCRYPTED PASSWORD ''mypassword''';
        RAISE NOTICE 'User myuser created.';
    ELSE
        RAISE NOTICE 'User myuser already exists.';
    END IF;
END $$;