DELETE FROM pc;
DELETE FROM pc_statuses WHERE name = 'available';
DELETE FROM pc_rooms;
DELETE FROM pc_types;
DELETE FROM ram;
DELETE FROM monitors;
DELETE FROM video_cards;
DELETE FROM processors;
DELETE FROM dish_statuses WHERE name = 'available';
DELETE FROM dishes;
DELETE FROM ram_types;
DELETE FROM monitor_producers;
DELETE FROM video_card_producers;
DELETE FROM processor_producers;
DELETE FROM pc_type_images;
DELETE FROM dish_images;

DELETE FROM user_roles WHERE name = 'user';
DELETE FROM pc_order_statuses WHERE name = 'done';
DELETE FROM dish_order_statuses WHERE name = 'done';

-- Сброс identity счетчиков для всех таблиц
DBCC CHECKIDENT ('processor_producers', RESEED, 0);
DBCC CHECKIDENT ('video_card_producers', RESEED, 0);
DBCC CHECKIDENT ('monitor_producers', RESEED, 0);
DBCC CHECKIDENT ('ram_types', RESEED, 0);
DBCC CHECKIDENT ('processors', RESEED, 0);
DBCC CHECKIDENT ('video_cards', RESEED, 0);
DBCC CHECKIDENT ('monitors', RESEED, 0);
DBCC CHECKIDENT ('ram', RESEED, 0);
DBCC CHECKIDENT ('pc_types', RESEED, 0);
DBCC CHECKIDENT ('pc_rooms', RESEED, 0);
DBCC CHECKIDENT ('pc', RESEED, 0);
DBCC CHECKIDENT ('dishes', RESEED, 0);
DBCC CHECKIDENT ('pc_type_images', RESEED, 0);
DBCC CHECKIDENT ('dish_images', RESEED, 0);