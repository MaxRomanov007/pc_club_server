IF EXISTS (SELECT * FROM sys.triggers WHERE name = 'TRG_pc_check')
    DROP TRIGGER TRG_pc_check;