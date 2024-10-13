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

-- Grant privileges if not already granted
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.role_table_grants 
                   WHERE grantee = 'myuser' AND table_name = 'myappdb') THEN
        EXECUTE 'GRANT ALL PRIVILEGES ON DATABASE myappdb TO myuser';
        RAISE NOTICE 'Privileges granted to myuser on myappdb.';
    ELSE
        RAISE NOTICE 'Privileges already granted to myuser on myappdb.';
    END IF;
END $$;

