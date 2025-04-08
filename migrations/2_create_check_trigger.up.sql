CREATE TRIGGER TRG_pc_check
    ON pc
    FOR INSERT, UPDATE
    AS
BEGIN
    IF EXISTS (SELECT 1
               FROM inserted
               WHERE [row] > (SELECT [rows] FROM pc_rooms WHERE pc_rooms.pc_room_id = inserted.pc_room_id)
                  OR [place] > (SELECT [places] FROM pc_rooms WHERE pc_rooms.pc_room_id = inserted.pc_room_id))
        BEGIN
            ROLLBACK TRANSACTION;
            THROW 54332,
                'Row or place inserted pc bigger than pc room rows or places',
                1;
        END
END