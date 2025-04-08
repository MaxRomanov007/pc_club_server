-- Производители процессоров
INSERT INTO processor_producers ([name])
VALUES ('Intel'),
       ('AMD');

-- Производители видеокарт
INSERT INTO video_card_producers ([name])
VALUES ('NVIDIA'),
       ('AMD'),
       ('ASUS'),
       ('MSI');

-- Производители мониторов
INSERT INTO monitor_producers ([name])
VALUES ('Samsung'),
       ('LG'),
       ('Acer'),
       ('ASUS');

-- Типы оперативной памяти
INSERT INTO ram_types ([name])
VALUES ('DDR3'),
       ('DDR4'),
       ('DDR5');

-- Процессоры
INSERT INTO processors (processor_producer_id, model)
VALUES (1, 'Core i5-12400F'),
       (1, 'Core i7-12700K'),
       (1, 'Core i9-12900K'),
       (2, 'Ryzen 5 5600X'),
       (2, 'Ryzen 7 5800X'),
       (2, 'Ryzen 9 5950X');

-- Видеокарты
INSERT INTO video_cards (video_card_producer_id, model)
VALUES (1, 'GeForce RTX 3060'),
       (1, 'GeForce RTX 3070'),
       (1, 'GeForce RTX 3080'),
       (2, 'Radeon RX 6600'),
       (2, 'Radeon RX 6700 XT'),
       (3, 'ROG Strix RTX 3060'),
       (4, 'GeForce RTX 3060 Gaming X');

-- Мониторы
INSERT INTO monitors (monitor_producer_id, model)
VALUES (1, 'Odyssey G5'),
       (1, 'Odyssey G7'),
       (2, 'UltraGear 27GN800-B'),
       (3, 'Nitro XV272U'),
       (4, 'TUF Gaming VG249Q');

-- Оперативная память
INSERT INTO ram (ram_type_id, capacity)
VALUES (2, 8),
       (2, 16),
       (2, 32),
       (3, 16),
       (3, 32);

-- Типы ПК (10 различных конфигураций)
INSERT INTO pc_types (processor_id, video_card_id, monitor_id, ram_id, [name], [description], hour_cost)
VALUES (1, 1, 1, 2, N'Базовый игровой',
        N'Отличный ПК для начинающих геймеров с хорошим соотношением цены и производительности', 50.00),
       (4, 4, 3, 2, N'Средний игровой', N'Сбалансированная сборка для комфортной игры в Full HD', 70.00),
       (2, 2, 2, 3, N'Продвинутый игровой', N'Мощная система для требовательных игр в 2K разрешении', 90.00),
       (5, 3, 4, 4, N'Профессиональный игровой', N'Топовая конфигурация для 4K игр и стриминга', 120.00),
       (3, 5, 5, 5, N'Элитный игровой', N'Максимальная производительность для киберспорта и профессионального гейминга',
        150.00),
       (1, 6, 1, 1, N'Офисный', N'Подходит для работы и нетребовательных задач', 30.00),
       (4, 7, 3, 2, N'Мультимедийный', N'Идеален для просмотра видео, работы с графикой и нетребовательных игр', 40.00),
       (2, 1, 2, 2, N'Разработческий', N'Хороший выбор для программистов и веб-разработчиков', 60.00),
       (6, 3, 4, 5, N'Дизайнерский', N'Мощная рабочая станция для 3D-моделирования и видеомонтажа', 100.00),
       (5, 2, 5, 3, N'Стримерский', N'Оптимизирован для стриминга и контент-креации', 80.00);

-- Компьютерные помещения
INSERT INTO pc_rooms ([name], [rows], places, [description])
VALUES (N'Альфа', 5, 10, N'Основной игровой зал с комфортными креслами'),
       (N'Бета', 4, 8, N'Зал для турниров и киберспортивных мероприятий'),
       (N'Гамма', 3, 6, N'Тихий зал для работы и учебы');

INSERT INTO pc_statuses
VALUES ('available');

-- Конкретные ПК (10 ПК каждого типа, всего 100 ПК)
-- Тип 1 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 1, 1, 1, 1, N'ПК у окна'),
       (1, 1, 1, 1, 2, NULL),
       (1, 1, 1, 1, 3, N'Возле розетки'),
       (1, 1, 1, 2, 1, NULL),
       (1, 1, 1, 2, 2, N'С наушниками'),
       (1, 1, 1, 2, 3, NULL),
       (2, 1, 1, 1, 1, N'Турнирный ПК #1'),
       (2, 1, 1, 1, 2, N'Турнирный ПК #2'),
       (3, 1, 1, 1, 1, N'Тихий уголок'),
       (3, 1, 1, 1, 2, NULL);

-- Тип 2 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 2, 1, 1, 4, NULL),
       (1, 2, 1, 1, 5, N'С веб-камерой'),
       (1, 2, 1, 2, 4, NULL),
       (1, 2, 1, 2, 5, N'С микрофоном'),
       (1, 2, 1, 3, 1, NULL),
       (2, 2, 1, 1, 3, N'Турнирный ПК #3'),
       (2, 2, 1, 1, 4, N'Турнирный ПК #4'),
       (2, 2, 1, 2, 1, NULL),
       (3, 2, 1, 1, 3, NULL),
       (3, 2, 1, 1, 4, N'У стены');

-- Тип 3 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 3, 1, 3, 2, NULL),
       (1, 3, 1, 3, 3, N'VIP место'),
       (1, 3, 1, 4, 1, NULL),
       (1, 3, 1, 4, 2, N'С охлаждением'),
       (2, 3, 1, 2, 2, N'Турнирный ПК #5'),
       (2, 3, 1, 2, 3, NULL),
       (2, 3, 1, 3, 1, N'Турнирный ПК #6'),
       (3, 3, 1, 2, 1, NULL),
       (3, 3, 1, 2, 2, N'Тихое место'),
       (3, 3, 1, 3, 1, NULL);

-- Тип 4 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 4, 1, 3, 4, N'VIP место с креслом DXRacer'),
       (1, 4, 1, 3, 5, NULL),
       (1, 4, 1, 4, 3, N'С профессиональными наушниками'),
       (1, 4, 1, 4, 4, NULL),
       (1, 4, 1, 5, 1, N'Угловой ПК'),
       (2, 4, 1, 2, 4, N'Турнирный ПК #7'),
       (2, 4, 1, 3, 2, NULL),
       (2, 4, 1, 3, 3, N'С микрофоном HyperX'),
       (3, 4, 1, 2, 3, NULL),
       (3, 4, 1, 3, 2, N'Тихое место у стены');

-- Тип 5 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 5, 1, 4, 5, N'Элитное место с 144Hz монитором'),
       (1, 5, 1, 5, 2, NULL),
       (1, 5, 1, 5, 3, N'С RGB подсветкой'),
       (2, 5, 1, 3, 4, N'Турнирный ПК #8'),
       (2, 5, 1, 4, 1, NULL),
       (2, 5, 1, 4, 2, N'С механической клавиатурой'),
       (3, 5, 1, 3, 3, NULL),
       (3, 5, 1, 3, 4, N'Для профессиональных игроков'),
       (3, 5, 1, 3, 5, NULL),
       (3, 5, 1, 3, 6, N'С ковриком XL');

-- Тип 6 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 6, 1, 5, 4, NULL),
       (1, 6, 1, 5, 5, N'Для офисных задач'),
       (1, 6, 1, 5, 6, NULL),
       (1, 6, 1, 5, 7, N'С принтером рядом'),
       (1, 6, 1, 5, 8, NULL),
       (2, 6, 1, 4, 3, N'Для организаторов турниров'),
       (3, 6, 1, 1, 5, NULL),
       (3, 6, 1, 1, 6, N'Для работы с документами'),
       (3, 6, 1, 2, 4, NULL),
       (3, 6, 1, 2, 5, N'С комфортной клавиатурой');

-- Тип 7 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 7, 1, 3, 6, NULL),
       (1, 7, 1, 3, 7, N'Для просмотра фильмов'),
       (1, 7, 1, 3, 8, NULL),
       (1, 7, 1, 4, 6, N'С аудиосистемой'),
       (1, 7, 1, 4, 7, NULL),
       (2, 7, 1, 4, 4, N'Для мультимедийных задач'),
       (3, 7, 1, 2, 6, NULL),
       (1, 7, 1, 1, 8, N'С графическим планшетом'),
       (1, 7, 1, 1, 9, NULL),
       (1, 7, 1, 1, 7, N'Для дизайнеров');

-- Тип 8 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 8, 1, 2, 6, NULL),
       (1, 8, 1, 4, 9, N'Для программистов'),
       (1, 8, 1, 2, 8, NULL),
       (1, 8, 1, 3, 9, N'С двумя мониторами'),
       (1, 8, 1, 2, 10, NULL),
       (2, 8, 1, 1, 5, N'Для разработчиков'),
       (1, 8, 1, 4, 10, NULL),
       (1, 8, 1, 2, 9, N'С IDE предустановленными'),
       (1, 8, 1, 3, 10, NULL),
       (1, 8, 1, 4, 8, N'Для веб-разработки');

-- Тип 9 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (1, 9, 1, 2, 7, N'Для 3D-моделирования'),
       (1, 9, 1, 1, 10, NULL),
       (1, 9, 1, 5, 9, N'С графическим планшетом Wacom'),
       (1, 9, 1, 5, 10, NULL),
       (2, 9, 1, 1, 6, N'Для видеомонтажа'),
       (2, 9, 1, 2, 5, N'Для дизайнеров'),
       (2, 9, 1, 2, 6, NULL),
       (2, 9, 1, 1, 7, N'Для CAD-программ'),
       (2, 9, 1, 1, 8, NULL),
       (2, 9, 1, 2, 7, N'С профессиональным цветокорректором');

-- Тип 10 (10 ПК)
INSERT INTO pc (pc_room_id, pc_type_id, pc_status_id, [row], place, [description])
VALUES (2, 10, 1, 2, 8, N'Для стримеров'),
       (2, 10, 1, 3, 5, NULL),
       (2, 10, 1, 3, 6, N'С зеленым экраном'),
       (2, 10, 1, 3, 7, NULL),
       (2, 10, 1, 3, 8, N'Для контент-креаторов'),
       (2, 10, 1, 4, 5, NULL),
       (2, 10, 1, 4, 6, N'С профессиональной камерой'),
       (2, 10, 1, 4, 7, NULL),
       (2, 10, 1, 4, 8, N'Для записи подкастов'),
       (1, 10, 1, 1, 6, N'С микрофоном Blue Yeti');

INSERT INTO dish_statuses
VALUES ('available');

-- Блюда (10 различных блюд)
INSERT INTO dishes (dish_status_id, [name], calories, cost, [description])
VALUES (1, N'Пицца Пепперони', 850, 350.00, N'Классическая пицца с колбасками пепперони и сыром моцарелла'),
       (1, N'Чизбургер', 550, 200.00, N'Аппетитный бургер с говяжьей котлетой и сыром'),
       (1, N'Картофель фри', 320, 120.00, N'Хрустящий картофель с соусами на выбор'),
       (1, N'Салат Цезарь', 280, 250.00, N'Свежий салат с курицей, сухариками и соусом цезарь'),
       (1, N'Куриные крылышки', 450, 280.00, N'Хрустящие куриные крылышки в остром соусе'),
       (1, N'Спагетти Болоньезе', 620, 300.00, N'Итальянская паста с мясным соусом'),
       (1, N'Шоколадный мусс', 380, 180.00, N'Нежный десерт из темного шоколада'),
       (1, N'Молочный коктейль', 420, 150.00, N'Охлаждающий напиток с ванильным мороженым'),
       (1, N'Чай черный', 5, 80.00, N'Ароматный черный чай с сахаром по вкусу'),
       (1, N'Кофе латте', 120, 150.00, N'Кофе с молоком и пенкой');

INSERT INTO user_roles
VALUES ('user');
INSERT INTO pc_order_statuses
VALUES ('done');
INSERT INTO dish_order_statuses
VALUES ('done');

INSERT INTO pc_type_images (pc_type_id, is_main, [path])
VALUES (1, 1, '1.jpg'),
       (2, 1, '2.jpg'),
       (3, 1, '3.jpg'),
       (4, 1, '4.jpg'),
       (5, 1, '5.jpg'),
       (6, 1, '6.jpg'),
       (7, 1, '7.jpg'),
       (8, 1, '8.jpg'),
       (9, 1, '9.jpg'),
       (10, 1, '10.jpg'),
       (5, 1, '41.jpg'),
       (6, 1, '31.jpg'),
       (7, 1, '21.jpg'),
       (8, 1, '11.jpg'),
       (9, 1, '51.jpg'),
       (10, 1, '61.jpg'),
       (4, 1, '71.jpg'),
       (3, 1, '81.jpg'),
       (2, 1, '91.jpg'),
       (1, 1, '101.jpg');

INSERT INTO dish_images (dish_id, is_main, [path])
VALUES (1, 1, 'd1.jpg'),
       (2, 1, 'd2.jpg'),
       (3, 1, 'd3.jpg'),
       (4, 1, 'd4.jpg'),
       (5, 1, 'd5.jpg'),
       (6, 1, 'd6.jpg'),
       (7, 1, 'd7.jpg'),
       (8, 1, 'd8.jpg'),
       (9, 1, 'd9.jpg'),
       (10, 1, 'd10.jpg');
