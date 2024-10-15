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
