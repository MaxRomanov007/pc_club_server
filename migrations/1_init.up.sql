-- роль пользователя
CREATE TABLE user_roles
(
    user_role_id BIGINT PRIMARY KEY IDENTITY,
    [name]       NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_user_role_name UNIQUE ([name]),
);

-- пользователь
CREATE TABLE users
(
    [user_id]             BIGINT PRIMARY KEY IDENTITY,
    user_role_id          BIGINT         NOT NULL,
    email                 NVARCHAR(255)  NOT NULL,
    [password]            VARBINARY(64)  NOT NULL,
    refresh_token_version BIGINT         NOT NULL DEFAULT 0,
    balance               DECIMAL(10, 2) NOT NULL DEFAULT 0,

    FOREIGN KEY (user_role_id) REFERENCES user_roles (user_role_id),
    CONSTRAINT UQ_user_email UNIQUE (email),
);

--производители комплектующих пк
CREATE TABLE processor_producers
(
    processor_producer_id BIGINT PRIMARY KEY IDENTITY,
    [name]                NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_processor_producer_name UNIQUE ([name]),
);

CREATE TABLE video_card_producers
(
    video_card_producer_id BIGINT PRIMARY KEY IDENTITY,
    [name]                 NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_video_card_producer_name UNIQUE ([name]),
);

CREATE TABLE monitor_producers
(
    monitor_producer_id BIGINT PRIMARY KEY IDENTITY,
    [name]              NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_monitor_producer_name UNIQUE ([name]),
);

-- тип ОЗУ (DDR3, DDR4, DDR5 и др)
CREATE TABLE ram_types
(
    ram_type_id BIGINT PRIMARY KEY IDENTITY,
    [name]      NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_ram_type_name UNIQUE ([name]),
);

-- комплектующие
CREATE TABLE processors
(
    processor_id          BIGINT PRIMARY KEY IDENTITY,
    processor_producer_id BIGINT        NOT NULL,
    model                 NVARCHAR(255) NOT NULL,

    FOREIGN KEY (processor_producer_id) REFERENCES processor_producers (processor_producer_id) ON DELETE CASCADE,
    CONSTRAINT UQ_processor_producer_model UNIQUE (processor_producer_id, model),
);

CREATE TABLE video_cards
(
    video_card_id          BIGINT PRIMARY KEY IDENTITY,
    video_card_producer_id BIGINT        NOT NULL,
    model                  NVARCHAR(255) NOT NULL,

    FOREIGN KEY (video_card_producer_id) REFERENCES video_card_producers (video_card_producer_id) ON DELETE CASCADE,
    CONSTRAINT UQ_video_card_producer_model UNIQUE (video_card_producer_id, model),
);

CREATE TABLE monitors
(
    monitor_id          BIGINT PRIMARY KEY IDENTITY,
    monitor_producer_id BIGINT        NOT NULL,
    model               NVARCHAR(255) NOT NULL,

    FOREIGN KEY (monitor_producer_id) REFERENCES monitor_producers (monitor_producer_id) ON DELETE CASCADE,
    CONSTRAINT UQ_monitor_producer_model UNIQUE (monitor_producer_id, model),
);

CREATE TABLE ram
(
    ram_id      BIGINT PRIMARY KEY IDENTITY,
    ram_type_id BIGINT NOT NULL,
    capacity    INT    NOT NULL,

    FOREIGN KEY (ram_type_id) REFERENCES ram_types (ram_type_id) ON DELETE CASCADE,
    CONSTRAINT UQ_ram_type_capacity UNIQUE (ram_type_id, capacity),
    CONSTRAINT CK_ram_capacity_over_zero CHECK (capacity > 0),
);

-- тип пк

CREATE TABLE pc_types
(
    pc_type_id    BIGINT PRIMARY KEY IDENTITY,
    processor_id  BIGINT         NOT NULL,
    video_card_id BIGINT         NOT NULL,
    monitor_id    BIGINT         NOT NULL,
    ram_id        BIGINT         NOT NULL,
    [name]        NVARCHAR(255)  NOT NULL,
    [description] NVARCHAR(MAX),
    hour_cost     DECIMAL(10, 2) NOT NULL,

    FOREIGN KEY (processor_id) REFERENCES processors (processor_id) ON DELETE CASCADE,
    FOREIGN KEY (video_card_id) REFERENCES video_cards (video_card_id) ON DELETE CASCADE,
    FOREIGN KEY (monitor_id) REFERENCES monitors (monitor_id) ON DELETE CASCADE,
    FOREIGN KEY (ram_id) REFERENCES ram (ram_id) ON DELETE CASCADE,
    CONSTRAINT UQ_pc_type_components UNIQUE (processor_id, video_card_id, monitor_id, ram_id),
    CONSTRAINT UQ_pc_type_name UNIQUE ([name]),
    CONSTRAINT CK_pc_type_hour_cost_over_zero CHECK (hour_cost > 0)
);

-- картинка к типу пк

CREATE TABLE pc_type_images
(
    pc_type_image_id BIGINT PRIMARY KEY IDENTITY,
    pc_type_id       BIGINT        NOT NULL,
    is_main          BIT           NOT NULL DEFAULT 0,
    [path]           NVARCHAR(255) NOT NULL,

    FOREIGN KEY (pc_type_id) REFERENCES pc_types (pc_type_id) ON DELETE CASCADE,
    CONSTRAINT UQ_pc_type_image_path UNIQUE ([path]),
    CONSTRAINT CK_pc_type_image_is_main_boolean CHECK (is_main IN (0, 1)),
);

-- компьютерные помещения

CREATE TABLE pc_rooms
(
    pc_room_id    BIGINT PRIMARY KEY IDENTITY,
    [name]        NVARCHAR(255) NOT NULL,
    [rows]        INT           NOT NULL,
    places        INT           NOT NULL,
    [description] NVARCHAR(MAX),

    CONSTRAINT UQ_pc_room_name UNIQUE ([name]),
    CONSTRAINT CK_pc_room_rows_over_zero CHECK ([rows] > 0),
    CONSTRAINT CK_pc_room_places_over_zero CHECK (places > 0),
);

-- статус пк (занят, свободен, сломан и тд)

CREATE TABLE pc_statuses
(
    pc_status_id BIGINT PRIMARY KEY IDENTITY,
    [name]       NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_pc_status_name UNIQUE ([name]),
);

-- конкретный пк

CREATE TABLE pc
(
    pc_id         BIGINT PRIMARY KEY IDENTITY,
    pc_room_id    BIGINT NOT NULL,
    pc_type_id    BIGINT NOT NULL,
    pc_status_id  BIGINT NOT NULL DEFAULT 1,
    [row]         int    NOT NULL,
    place         int    NOT NULL,
    [description] NVARCHAR(MAX),

    FOREIGN KEY (pc_room_id) REFERENCES pc_rooms (pc_room_id) ON DELETE CASCADE,
    FOREIGN KEY (pc_type_id) REFERENCES pc_types (pc_type_id) ON DELETE CASCADE,
    FOREIGN KEY (pc_status_id) REFERENCES pc_statuses (pc_status_id) ON DELETE CASCADE,
    CONSTRAINT UQ_pc_room_row_place UNIQUE (pc_room_id, [row], place),
    CONSTRAINT CK_pc_row_over_zero CHECK ([row] > 0),
    CONSTRAINT CK_pc_place_over_zero CHECK ([place] > 0),
);

-- статус брони пк (будущ, в действии, окончена)

CREATE TABLE pc_order_statuses
(
    pc_order_status_id BIGINT PRIMARY KEY IDENTITY,
    [name]             NVARCHAR(255),

    CONSTRAINT UQ_pc_order_status_name UNIQUE ([name]),
);

-- бронь пк

CREATE TABLE pc_orders
(
    pc_order_id        BIGINT PRIMARY KEY IDENTITY,
    [user_id]          BIGINT         NOT NULL,
    pc_id              BIGINT         NOT NULL,
    pc_order_status_id BIGINT         NOT NULL,
    code               NVARCHAR(255)  NOT NULL,
    cost               DECIMAL(10, 2) NOT NULL,
    start_time         SMALLDATETIME  NOT NULL,
    end_time           SMALLDATETIME  NOT NULL,
    actual_end_time    SMALLDATETIME  NOT NULL,
    order_date         DATETIME       NOT NULL DEFAULT GETDATE(),
    duration           INT            NOT NULL DEFAULT 0,

    FOREIGN KEY ([user_id]) REFERENCES users ([user_id]) ON DELETE CASCADE,
    FOREIGN KEY (pc_id) REFERENCES pc (pc_id) ON DELETE CASCADE,
    FOREIGN KEY (pc_order_status_id) REFERENCES pc_order_statuses (pc_order_status_id) ON DELETE CASCADE,
);

-- состояние блюда (доступно, нет и тд)

CREATE TABLE dish_statuses
(
    dish_status_id BIGINT PRIMARY KEY IDENTITY,
    [name]         NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_dish_status_name UNIQUE ([name]),
);

-- блюдо

CREATE TABLE dishes
(
    dish_id        BIGINT PRIMARY KEY IDENTITY,
    dish_status_id BIGINT         NOT NULL,
    [name]         NVARCHAR(255)  NOT NULL,
    calories       SMALLINT       NOT NULL,
    cost           DECIMAL(10, 2) NOT NULL,
    [description]  NVARCHAR(MAX),

    FOREIGN KEY (dish_status_id) REFERENCES dish_statuses (dish_status_id) ON DELETE CASCADE,
    CONSTRAINT UQ_dish_name UNIQUE ([name]),
    CONSTRAINT CK_dish_calories_over_zero CHECK (calories > 0),
    CONSTRAINT CK_dish_cost_over_zero CHECK (cost > 0),
);

-- картинка на блюдо

CREATE TABLE dish_images
(
    dish_image_id BIGINT PRIMARY KEY IDENTITY,
    dish_id       BIGINT NOT NULL,
    is_main       BIT    NOT NULL DEFAULT 0,
    [path]        NVARCHAR(255),

    FOREIGN KEY (dish_id) REFERENCES dishes (dish_id) ON DELETE CASCADE,
    CONSTRAINT UQ_dish_image_path UNIQUE ([path]),
    CONSTRAINT CK_dish_image_is_main_boolean CHECK (is_main IN (0, 1)),
);

-- статус заказа еды (в корзине, готовится, готов и тд)

CREATE TABLE dish_order_statuses
(
    dish_order_status_id BIGINT PRIMARY KEY IDENTITY,
    [name]               NVARCHAR(255) NOT NULL,

    CONSTRAINT UQ_dish_order_status_name UNIQUE ([name]),
);

-- заказ еды

CREATE TABLE dish_orders
(
    dish_order_id        BIGINT PRIMARY KEY IDENTITY,
    dish_order_status_id BIGINT         NOT NULL,
    [user_id]            BIGINT         NOT NULL,
    cost                 DECIMAL(10, 2) NOT NULL,
    order_date           DATETIME       NOT NULL DEFAULT GETDATE(),

    FOREIGN KEY (dish_order_status_id) REFERENCES dish_order_statuses (dish_order_status_id) ON DELETE CASCADE,
    FOREIGN KEY ([user_id]) REFERENCES users ([user_id]) ON DELETE CASCADE,

    CONSTRAINT CK_dish_order_cost_over_zero CHECK (cost > 0),
);

CREATE TABLE dish_order_list
(
    dish_order_id BIGINT   NOT NULL,
    dish_id       BIGINT   NOT NULL,
    [count]       SMALLINT NOT NULL,

    FOREIGN KEY (dish_order_id) REFERENCES dish_orders (dish_order_id) ON DELETE CASCADE,
    FOREIGN KEY (dish_id) REFERENCES dishes (dish_id) ON DELETE CASCADE,
);