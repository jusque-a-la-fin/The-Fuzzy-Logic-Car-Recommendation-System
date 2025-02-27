-- скрипт для создания базы данных "vehicles"

BEGIN;
-- страны
CREATE TABLE countries (
  id SERIAL PRIMARY KEY,
  -- страна
  country VARCHAR(100) UNIQUE NOT NULL
);

-- марки
CREATE TABLE makes (
  id SERIAL PRIMARY KEY,
  -- марка
  make VARCHAR(100) UNIQUE NOT NULL,
  country_id INT,
  FOREIGN KEY (country_id) REFERENCES countries(id)
);

-- модели
CREATE TABLE models (
  id SERIAL PRIMARY KEY,
  -- модель
  model VARCHAR(100) UNIQUE NOT NULL,
  make_id INT,
  FOREIGN KEY (make_id) REFERENCES makes(id)
);

-- поколения автомобилей
CREATE TABLE generations (
  id SERIAL PRIMARY KEY,
  model_id INT,
  -- поколение
  generation VARCHAR(100) UNIQUE NOT NULL,
  FOREIGN KEY (model_id) REFERENCES models(id)
);

-- позиции руля
CREATE TYPE steering_wheel_position_enum AS ENUM ('Левый руль', 'Правый руль', 'Неизвестно');
CREATE TABLE steering_wheel_positions (
  id SERIAL PRIMARY KEY,
  -- позиция руля
  position steering_wheel_position_enum UNIQUE NOT NULL
);

-- типы усилителей рулевого управления
CREATE TYPE power_steering_types_enum AS ENUM ('Электроусилитель руля', 'Гидроусилитель руля', 'Электрогидроусилитель руля', 'Нет', 'Неизвестно');
CREATE TABLE power_steering_types (
  id SERIAL PRIMARY KEY,
  -- тип усилителя рулевого управления
  power_steering power_steering_types_enum UNIQUE NOT NULL
);


-- типы кузовов
CREATE TABLE body_types (
  id SERIAL PRIMARY KEY,
  -- тип кузова
  body VARCHAR(30) UNIQUE NOT NULL
);


CREATE TYPE bool_enum AS ENUM ('Есть', 'Нет', 'Неизвестно', 'Опция производителя');

-- типы подвесок
CREATE TABLE suspensions (
  id SERIAL PRIMARY KEY,
  -- передний стабилизатор
  front_stabilizer bool_enum,
  -- задний стабилизатор
  back_stabilizer bool_enum,
  -- тип передней подвески
  front_suspension VARCHAR(100),
  -- тип задней подвески
  back_suspension VARCHAR(100)
);


-- спецификации
CREATE TABLE specifications (
  id SERIAL PRIMARY KEY,
  generation_id INT,
  steering_wheel_position_id INT,
  power_steering_type_id INT,
  body_type_id INT,
  suspensions_id INT,
  -- длина автомобиля, мм
  length FLOAT,
  -- ширина, мм
  width FLOAT,
  -- высота, мм
  height FLOAT,
  -- дорожный просвет или клиренс, мм
  ground_clearance FLOAT,
  -- коэффициент аэродинамического сопротивления, cW
  drag_coefficient FLOAT,
  -- ширина передней колеи, мм
  front_track_width FLOAT,
  -- ширина задней колеи, мм"
  back_track_width FLOAT,
  -- колесная база, мм
  wheelbase FLOAT,
  -- баллы за краш-тест
  crash_test_estimate FLOAT,
  -- год выпуска
  year INT,
  FOREIGN KEY (generation_id) REFERENCES generations(id),
  FOREIGN KEY (body_type_id) REFERENCES body_types(id),
  FOREIGN KEY (steering_wheel_position_id) REFERENCES steering_wheel_positions(id),
  FOREIGN KEY (power_steering_type_id) REFERENCES power_steering_types(id),
  FOREIGN KEY (suspensions_id) REFERENCES suspensions(id)
);

-- двигатели
CREATE TABLE engines (
  id SERIAL PRIMARY KEY,
  -- использумое топливо
  fuel_used VARCHAR(50),
  -- тип двигателя
  engine_type VARCHAR(50),
  -- объем двигателя, куб.см
  capacity FLOAT,
  -- максимальная мощность, л.с.
  power FLOAT,
  -- максимальный крутящий момент, Н*м (кг*м) при об./мин.
  max_torque VARCHAR(50)
);

-- типы трансмиссии
CREATE TABLE gearboxes (
  id SERIAL PRIMARY KEY,
  -- тип трансмиссии
  gearbox VARCHAR(50) UNIQUE NOT NULL
);

-- типы привода
CREATE TABLE drive_types (
  id SERIAL PRIMARY KEY,
  -- тип привода
  drive VARCHAR(50) UNIQUE NOT NULL
);


-- шины
CREATE TABLE tires (
  id SERIAL PRIMARY KEY,
  -- ширина передних шин, мм
  back_tires_width INT,
  -- ширина задних шин, мм
  front_tires_width INT,
  -- процентное соотношение высоты профиля передних шин к их ширине
  front_tires_aspect_ratio INT,
  -- процентное соотношение высоты профиля задних шин к их ширине
  back_tires_aspect_ratio INT,
  -- диаметр обода передних шин, мм
  front_tires_rim_diameter INT,
  -- диаметр обода задних шин, мм
  back_tires_rim_diameter INT
);

-- тормоза
CREATE TABLE brakes (
  id SERIAL PRIMARY KEY,
  -- тип передних тормозов
  front_brakes VARCHAR(50),
  -- тип задних тормозов
  back_brakes VARCHAR(50),
  -- тип стояночного тормоза
  parking_brake VARCHAR(50)
);


-- электронные системы безопасности и контроля движения
CREATE TABLE safety_and_motion_control_systems (
  id SERIAL PRIMARY KEY,
  -- наличие антиблокировочной системы (ABS)
  abs_system bool_enum,
  -- наличие системы электронного контроля устойчивости (ESP)
  esp_system bool_enum,
  -- наличие системы распределения тормозного усилия (EBD)
  ebd_system bool_enum,
  -- наличие вспомогательной системы торможения (BAS)
  bas_system bool_enum,
  -- наличие антипробуксовочной системы (TCS)
  tcs_system bool_enum,
  -- наличие переднего парктроника
  front_parking_sensor bool_enum, 
  -- наличие заднего парктроника
  back_parking_sensor bool_enum,
  -- наличие камеры заднего обзора
  rear_view_camera bool_enum,
  -- наличие круиз-контроля
  cruise_control bool_enum
);

-- цвета кузова
CREATE TABLE colors (
  id SERIAL PRIMARY KEY,
  -- цвет кузова
  color VARCHAR(40) UNIQUE NOT NULL
);


-- фонари и фары
CREATE TABLE lights (
  id SERIAL PRIMARY KEY,
  -- тип передних фар
  headlights VARCHAR(50),
  -- наличие светодиодных ходовых огней
  led_running_lights bool_enum,
  -- наличие светодиодных задних фонарей
  led_tail_lights bool_enum,
  -- наличие датчика света
  light_sensor bool_enum,
  -- наличие передних противотуманных фар
  front_fog_lights bool_enum,
  -- наличие задних противотуманных фонарей
  back_fog_lights bool_enum
);

-- интерьер
CREATE TABLE interior_design(
 id SERIAL PRIMARY KEY,
 -- тип обивки салона
 upholstery VARCHAR(100) UNIQUE NOT NULL
);

-- микроклимат салона
CREATE TABLE cabin_microclimate (
  id SERIAL PRIMARY KEY,
  -- наличие кондиционера
  air_conditioner bool_enum,
  -- наличие климат-контроля
  climate_control bool_enum
);

-- пакет электрических опций
CREATE TABLE electric_options (
  id SERIAL PRIMARY KEY, 
  -- наличие электрических передних стеклоподъемников
  electric_front_side_windows_lifts bool_enum,
  -- наличие электрических задних стеклоподъемников
  electric_back_side_windows_lifts bool_enum,
  -- наличие электроподогрева передних сидений
  electric_heating_of_front_seats bool_enum,
  -- наличие электроподогрева задних сидений
  electric_heating_of_back_seats bool_enum,
  -- наличие электроподогрева рулевого колеса
  electric_heating_of_steering_wheel bool_enum,
  -- наличие электроподогрева лобового стекла
  electric_heating_of_windshield bool_enum,
  -- наличие обогрева заднего стекла
  electric_heating_of_rear_window bool_enum,
  -- наличие электроподогрева боковых зеркал
  electric_heating_of_side_mirrors bool_enum,
  -- наличие электропривода водительского сидения
  electric_drive_of_driver_seat bool_enum,
  -- наличие электропривода передних сидений
  electric_drive_of_front_seats bool_enum,
  -- наличие электропривода боковых зеркал
  electric_drive_of_side_mirrors bool_enum,
  -- наличие электропривода багажника
  electric_trunk_opener bool_enum,
  -- наличие датчика дождя
  rain_sensor bool_enum
);

-- подушки безопасности
CREATE TABLE airbags (
  id SERIAL PRIMARY KEY,
  -- наличие водительской подушки безопасности
  driver_airbag bool_enum,
  -- наличие подушки безопасности переднего пассажира
  front_passenger_airbag bool_enum,
  -- наличие боковых подушек безопасности
  side_airbags bool_enum,
  -- наличие подушек безопасности-шторок
  curtain_airbags bool_enum
);

-- системы мультимедиа
CREATE TABLE multimedia_systems (
  id SERIAL PRIMARY KEY,
  -- наличие бортового компьютера
  on_board_computer bool_enum,
  -- наличие поддержки MP3
  mp3_support bool_enum,
  -- наличие поддержки Hands free
  hands_free_support bool_enum
);

-- комплектации
CREATE TABLE trim_levels (
  id SERIAL PRIMARY KEY,
  engine_id INT,
  gearbox_id INT,
  drive_type_id INT,
  color_id INT,
  specification_id INT,
  tires_id INT,
  brakes_id INT,
  safety_and_motion_control_systems_id INT,
  lights_id INT,
  interior_design_id INT,
  cabin_microclimate_id INT,
  electric_options_id INT,
  airbags_id INT,
  multimedia_systems_id INT, 
  -- название комплектации
  trim_level VARCHAR(100) UNIQUE NOT NULL,
  -- время разгона 0-100 км/ч, с
  acceleration_0_to_100 FLOAT,
  -- максимальная скорость, км/ч
  max_speed FLOAT,
  -- расход топлива в городском цикле, л/100 км
  city_fuel_consumption FLOAT,
  -- расход топлива за городом, л/100 км
  highway_fuel_consumption FLOAT,
  -- расход топлива в смешанном цикле, л/100 км
  mixed_fuel_consumption FLOAT,
  -- число мест
  number_of_seats INT,
  -- объем багажника, литры
  trunk_volume FLOAT,
  -- масса, кг
  mass FLOAT,
  -- наличие сигнализации
  car_alarm bool_enum,
  FOREIGN KEY (engine_id) REFERENCES engines(id),
  FOREIGN KEY (gearbox_id) REFERENCES gearboxes(id),
  FOREIGN KEY (drive_type_id) REFERENCES drive_types(id),
  FOREIGN KEY (color_id) REFERENCES colors(id),
  FOREIGN KEY (specification_id) REFERENCES specifications(id),
  FOREIGN KEY (tires_id) REFERENCES tires(id),
  FOREIGN KEY (brakes_id) REFERENCES brakes(id),
  FOREIGN KEY (safety_and_motion_control_systems_id) REFERENCES safety_and_motion_control_systems(id),
  FOREIGN KEY (lights_id) REFERENCES lights(id),
  FOREIGN KEY (interior_design_id) REFERENCES interior_design(id),
  FOREIGN KEY (cabin_microclimate_id) REFERENCES cabin_microclimate(id),
  FOREIGN KEY (electric_options_id) REFERENCES electric_options(id),
  FOREIGN KEY (airbags_id) REFERENCES airbags(id),
  FOREIGN KEY (multimedia_systems_id) REFERENCES multimedia_systems(id)
);

-- сведения для покупателя
CREATE TABLE offerings (
  id SERIAL PRIMARY KEY,
  trim_level_id INT,
  -- описание
  description TEXT,
  -- цена, руб
  price FLOAT,
  -- пробег, км
  kilometrage INT,
  -- фотографии
  photo_urls TEXT[] UNIQUE NOT NULL,
  FOREIGN KEY (trim_level_id) REFERENCES trim_levels(id)
);

COMMIT;


BEGIN;

INSERT INTO steering_wheel_positions (position)
VALUES ('Левый руль'), ('Правый руль'), ('Неизвестно');

INSERT INTO power_steering_types (power_steering)
VALUES ('Электроусилитель руля'), ('Гидроусилитель руля'), ('Электрогидроусилитель руля');

INSERT INTO countries (country) VALUES ('Германия');
INSERT INTO countries (country) VALUES ('Япония');
INSERT INTO countries (country) VALUES ('Россия');
INSERT INTO countries (country) VALUES ('Китай');
INSERT INTO countries (country) VALUES ('США');
INSERT INTO countries (country) VALUES ('Великобритания');
INSERT INTO countries (country) VALUES ('Франция');
INSERT INTO countries (country) VALUES ('Южная_Корея');
INSERT INTO countries (country) VALUES ('Чехия');


INSERT INTO makes (make, country_id) 
SELECT 'Volkswagen', id FROM countries WHERE country = 'Германия';


INSERT INTO models (model, make_id) 
SELECT 'Polo', id FROM makes WHERE make = 'Volkswagen';

INSERT INTO generations (model_id, generation) 
SELECT id, '5 поколение (MK5)' FROM models WHERE model = 'Polo';

INSERT INTO body_types (body)
VALUES ('Седан');

  INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
  VALUES ('Есть', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Полузависимая, торсионная балка');

  INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id,
                              length, width, height, ground_clearance, drag_coefficient, front_track_width,
                              back_track_width, wheelbase, crash_test_estimate, year)
  VALUES (
    (SELECT id FROM generations WHERE generation = '5 поколение (MK5)'),
    (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
    (SELECT id FROM power_steering_types WHERE power_steering = 'Электроусилитель руля'),
    (SELECT id FROM body_types WHERE body = 'Седан'),
    1, -- suspensions_id
    4384, -- length
    1699, -- width
    1465, -- -height
    170, -- ground_clearance
    0.327, -- drag_coefficient
    1460, -- front_track_width
    1498, -- back_track_width
    2552, -- wheelbase
	2, -- crach_test_estimate
    2011 -- year
  );


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1598, 105, '153 (16) /3800');

INSERT INTO gearboxes (gearbox)
VALUES ('АКПП');

INSERT INTO drive_types (drive)
VALUES ('Передний FF');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (195, 195, 55, 55, 15, 15);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Барабанные', 'Ручной');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Опция производителя', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Нет', 'Опция производителя', 'Неизвестно', 'Нет');


INSERT INTO colors (color)
VALUES ('Черный');

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Опция производителя', 'Есть', 'Есть');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая');

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Есть');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Опция производителя',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть', -- electric_heating_of_side_mirrors
     'Нет',-- electric_drive_of_driver_seat 
     'Нет',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Нет',-- electric_trunk_opener
     'Опция производителя'-- rain_sensor
  );


INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Опция производителя', 'Нет');
  
INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Есть', 'Есть');


INSERT INTO trim_levels (
  engine_id,
  gearbox_id,
  drive_type_id,
  color_id,
  specification_id,
  tires_id,
  brakes_id,
  safety_and_motion_control_systems_id,
  lights_id,
  interior_design_id,
  cabin_microclimate_id,
  electric_options_id,
  airbags_id,
  multimedia_systems_id, 
  trim_level,
  acceleration_0_to_100,
  max_speed,
  city_fuel_consumption,
  highway_fuel_consumption,
  mixed_fuel_consumption,
  number_of_seats,
  trunk_volume,
  mass,
  car_alarm)
VALUES
    ((SELECT id FROM engines WHERE id = 1),
     (SELECT id FROM gearboxes WHERE gearbox = 'АКПП'),
     (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
     (SELECT id FROM colors WHERE color = 'Черный'),
     (SELECT id FROM specifications WHERE id =  1),
     (SELECT id FROM tires WHERE id = 1),
     (SELECT id FROM brakes WHERE id = 1),
     (SELECT id FROM safety_and_motion_control_systems WHERE id =  1),
     (SELECT id FROM lights WHERE id = 1),
     (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
     (SELECT id FROM cabin_microclimate WHERE id =  1),
     (SELECT id FROM electric_options WHERE id =  1),
     (SELECT id FROM airbags WHERE id = 1),
     (SELECT id FROM multimedia_systems WHERE id =  1),
      '1.6 MPI Tiptronic Highline', -- trim_level
    12.1, -- acceleration_0_to_100
    187, -- max_speed
    9.8, -- city_fuel_consumption
    5.4, -- highway_fuel_consumption
    7.0, -- mixed_fuel_consumption
    5, -- number_of_seats
    460, -- trunk_volume
    1217, -- mass
    'Есть' -- car_alarm
     );


INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.6 MPI Tiptronic Highline'),
  '',
  675000,
  256178,
  ARRAY['/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo1.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo2.jpg', 
        '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo3.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo4.jpg', 
        '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo5.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo6.jpg', 
        '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo7.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo8.jpg', 
        '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo9.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo10.jpg', 
        '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo11.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo12.jpg', 
        '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo13.jpg', '/userID=45e6f2b0-d91c-4751-9a6c-5f6702080361&carID=a9e0c8a6-536c-4eee-9237-c873b16c41de&carName=volkswagen_polo/polo14.jpg']);

COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Renault', id FROM countries WHERE country = 'Франция';

INSERT INTO models (model, make_id) 
SELECT 'Megane', id FROM makes WHERE make = 'Renault';


INSERT INTO generations (model_id, generation) 
SELECT id, '3 поколение (Megane III)' FROM models WHERE model = 'Megane';


INSERT INTO body_types (body)
VALUES ('Хэтчбек')
ON CONFLICT DO NOTHING;


INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Неизвестно', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Полузависимая, торсионная балка');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '3 поколение (Megane III)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Электроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Хэтчбек'),
  2, -- suspensions_id
  4295, -- length
  1808, -- width
  1471, -- height
  165, -- ground_clearance
  0.324, -- drag_coefficient
  1546, -- front_track_width
  1547, -- back_track_width
  2641, -- wheelbase
  3.33, -- crach_test_estimate
  2012 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1598, 106, '145 (15) / 4250');

INSERT INTO gearboxes (gearbox)
VALUES ('МКПП');


INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (205, 205, 65, 65, 15, 15);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Нет', 'Есть', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Нет');


INSERT INTO colors (color)
VALUES ('Белый');

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Нет', 'Нет', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Нет');



INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Опция производителя',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Неизвестно',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Неизвестно',-- electric_drive_of_driver_seat 
     'Неизвестно',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Нет'-- rain_sensor
  );
  
  

INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Нет', 'Нет');

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Опция производителя', 'Нет');


INSERT INTO trim_levels (
  engine_id,
  gearbox_id,
  drive_type_id,
  color_id,
  specification_id,
  tires_id,
  brakes_id,
  safety_and_motion_control_systems_id,
  lights_id,
  interior_design_id,
  cabin_microclimate_id,
  electric_options_id,
  airbags_id,
  multimedia_systems_id, 
  trim_level,
  acceleration_0_to_100,
  max_speed,
  city_fuel_consumption,
  highway_fuel_consumption,
  mixed_fuel_consumption,
  number_of_seats,
  trunk_volume,
  mass,
  car_alarm)
VALUES
    ((SELECT id FROM engines WHERE id = 2),
     (SELECT id FROM gearboxes WHERE gearbox = 'МКПП'),
     (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
     (SELECT id FROM colors WHERE color = 'Белый'),
     (SELECT id FROM specifications WHERE id = 2),
     (SELECT id FROM tires WHERE id = 2),
     (SELECT id FROM brakes WHERE id = 2),
     (SELECT id FROM safety_and_motion_control_systems WHERE id = 2),
     (SELECT id FROM lights WHERE id = 2),
     (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
     (SELECT id FROM cabin_microclimate WHERE id = 2),
     (SELECT id FROM electric_options WHERE id = 2),
     (SELECT id FROM airbags WHERE id = 2),
     (SELECT id FROM multimedia_systems WHERE id = 2),
     '1.6 MT Authentique', -- trim_level
     11.7, -- acceleration_0_to_100
     185, -- max_speed
     9.2, -- city_fuel_consumption
     5.4, -- highway_fuel_consumption
     6.8, -- mixed_fuel_consumption
     5, -- number_of_seats
     368, -- trunk_volume
     1260, -- mass
     'Неизвестно' -- car_alarm
     );

INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.6 MT Authentique'),
  '',
  670000,
  207529,
  ARRAY['/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane1.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane2.jpg', 
        '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane3.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane4.jpg', 
        '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane5.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane6.jpg',
        '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane7.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane8.jpg', 
        '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane9.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane10.jpg', 
        '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane11.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane12.jpg', 
        '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane13.jpg', '/userID=b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f&carID=e634c6c1-fa5b-4e70-bde9-79e45af65d46&carName=renault_megane/megane14.jpg']);
        
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Toyota', id FROM countries WHERE country = 'Япония';

INSERT INTO models (model, make_id) 
SELECT 'Avensis', id FROM makes WHERE make = 'Toyota';


INSERT INTO generations (model_id, generation) 
SELECT id, '2 поколение рестайлинг (T250)' FROM models WHERE model = 'Avensis';

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Есть', 'Есть', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, на двойных поперечных рычагах');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '2 поколение рестайлинг (T250)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Электроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Седан'),
  3, -- suspensions_id
  4645, -- length
  1760, -- width
  1480, -- height
  155, -- ground_clearance
  0.28, -- drag_coefficient
  1505, -- front_track_width
  1500, -- back_track_width
  2700, -- wheelbase
  3.66,  -- crach_test_estimate
  2008 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1598, 105, '153 (16) /3800');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (205, 205, 55, 55, 16, 16);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Есть', 'Есть', 'Нет', 'Есть', 'Нет', 'Опция производителя', 'Нет', 'Нет');


INSERT INTO colors (color)
VALUES ('Бежевый');

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Ксеноновые фары', 'Нет', 'Нет', 'Есть', 'Есть', 'Есть');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Есть');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Неизвестно',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Есть',-- electric_drive_of_driver_seat 
     'Опция производителя',-- electric_drive_of_front_seats
     'Опция производителя',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Нет'-- rain_sensor
  );
  
  
INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Есть', 'Есть'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Неизвестно', 'Есть');


INSERT INTO trim_levels (
  engine_id,
  gearbox_id,
  drive_type_id,
  color_id,
  specification_id,
  tires_id,
  brakes_id,
  safety_and_motion_control_systems_id,
  lights_id,
  interior_design_id,
  cabin_microclimate_id,
  electric_options_id,
  airbags_id,
  multimedia_systems_id, 
  trim_level,
  acceleration_0_to_100,
  max_speed,
  city_fuel_consumption,
  highway_fuel_consumption,
  mixed_fuel_consumption,
  number_of_seats,
  trunk_volume,
  mass,
  car_alarm)
VALUES
    ((SELECT id FROM engines WHERE id = 3),
     (SELECT id FROM gearboxes WHERE gearbox = 'МКПП'),
     (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
     (SELECT id FROM colors WHERE color = 'Бежевый'),
     (SELECT id FROM specifications WHERE id = 3),
     (SELECT id FROM tires WHERE id = 3),
     (SELECT id FROM brakes WHERE id = 3),
     (SELECT id FROM safety_and_motion_control_systems WHERE id = 3),
     (SELECT id FROM lights WHERE id = 3),
     (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
     (SELECT id FROM cabin_microclimate WHERE id = 3),
     (SELECT id FROM electric_options WHERE id = 3),
     (SELECT id FROM airbags WHERE id = 3),
     (SELECT id FROM multimedia_systems WHERE id = 3),
     '1.8 MT Executive', -- trim_level
     10, -- acceleration_0_to_100
     200, -- max_speed
     9.4, -- city_fuel_consumption
     5.8, -- highway_fuel_consumption
     7.2, -- mixed_fuel_consumption
     5, -- number_of_seats
     520, -- trunk_volume
     1355, -- mass
     'Неизвестно' -- car_alarm
     );

  
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.8 MT Executive'),
  '',
  677000,
  255416,
  ARRAY['/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis1.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis2.jpg', 
        '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis3.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis4.jpg', 
        '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis5.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis6.jpg', 
        '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis7.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis8.jpg', 
        '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis9.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis10.jpg', 
        '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis11.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis12.jpg', 
        '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis13.jpg', '/userID=67256291-50ac-4559-be40-8a3abf3e2577&carID=2d323aa2-4f54-455f-a422-12416ed5fa58&carName=toyota_avensis/avensis14.jpg']);
        
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Kia', id FROM countries WHERE country = 'Южная_Корея';

INSERT INTO models (model, make_id) 
SELECT 'Rio', id FROM makes WHERE make = 'Kia';


INSERT INTO generations (model_id, generation) 
SELECT id, '3 поколение рестайлинг (QB)' FROM models WHERE model = 'Rio';


INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Есть', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Полузависимая, торсионная балка');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '3 поколение рестайлинг (QB)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Седан'),
  4, -- suspensions_id
  4370, -- length
  1700, -- width
  1470, -- height
  160, -- ground_clearance
  0.31, -- drag_coefficient
  1495, -- front_track_width
  1502, -- back_track_width
  2570, -- wheelbase
  3.66, -- crach_test_estimate
  2012 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1591, 123, '155 (16) / 4200');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (185, 185, 65, 65, 15, 15);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Нет', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Нет', 'Нет', 'Неизвестно');


INSERT INTO colors (color)
VALUES ('Белый')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Неизвестно', 'Есть', 'Есть');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Есть');

INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Есть',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Неизвестно',-- electric_drive_of_driver_seat 
     'Неизвестно',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Неизвестно'-- rain_sensor
  );
  
  
INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Есть', 'Есть'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Неизвестно', 'Нет');

  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 4),
    (SELECT id FROM gearboxes WHERE gearbox = 'МКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
    (SELECT id FROM colors WHERE color = 'Белый'),
    (SELECT id FROM specifications WHERE id = 4),
    (SELECT id FROM tires WHERE id = 4),
    (SELECT id FROM brakes WHERE id = 4),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 4),
    (SELECT id FROM lights WHERE id = 4),
    (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
    (SELECT id FROM cabin_microclimate WHERE id = 4),
    (SELECT id FROM electric_options WHERE id = 4),
    (SELECT id FROM airbags WHERE id = 4),
    (SELECT id FROM multimedia_systems WHERE id = 4),
    '1.6 MT Prestige', -- trim_level
    10.3, -- acceleration_0_to_100
    190, -- max_speed
    7.9, -- city_fuel_consumption
    4.9, -- highway_fuel_consumption
    6.0, -- mixed_fuel_consumption
    5, -- number_of_seats
    500, -- trunk_volume
    1155, -- mass
    'Неизвестно' -- car_alarm
  );




INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.6 MT Prestige'),
  '',
  670000,
  201362,
  ARRAY['/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio1.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio2.jpg', 
        '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio3.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio4.jpg', 
        '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio5.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio6.jpg', 
        '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio7.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio8.jpg', 
        '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio9.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio10.jpg', 
        '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio11.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio12.jpg', 
        '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio13.jpg', '/userID=4eee2a93-5ed3-4799-b9d6-62fbdd13fa73&carID=d6fc0fbe-9582-402e-b5cc-4e9f684f0ecf&carName=kia_rio/rio14.jpg']);
  
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'LADA', id FROM countries WHERE country = 'Россия';

INSERT INTO models (model, make_id) 
SELECT '4x4 2121 Нива', id FROM makes WHERE make = 'LADA';


INSERT INTO generations (model_id, generation) 
SELECT id, '1 поколение 4x4 2121 Нива' FROM models WHERE model = '4x4 2121 Нива';

INSERT INTO body_types (body)
VALUES ('Внедорожник')
ON CONFLICT DO NOTHING;

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Есть', 'Есть', 'Независимая, на двойных поперечных рычагах', 'Зависимая, пружинная');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '1 поколение 4x4 2121 Нива'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Внедорожник'),
  5, -- suspensions_id
  3740, -- length
  1680, -- width
  1640, -- height
  205, -- ground_clearance
  0.42, -- drag_coefficient
  1440, -- front_track_width
  1420, -- back_track_width
  2200, -- wheelbase
  2, -- crach_test_estimate
  2018 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1690, 83, '129 (13) / 4000');

INSERT INTO drive_types (drive)
VALUES ('Полный 4WD');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (175, 175, 80, 80, 16, 16);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые', 'Барабанные', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Неизвестно', 'Есть', 'Есть', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Неизвестно');


INSERT INTO colors (color)
VALUES ('Белый')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Неизвестно', 'Неизвестно', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Комбинированная')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Неизвестно');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Нет', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Неизвестно',-- electric_drive_of_driver_seat 
     'Неизвестно',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Неизвестно'-- rain_sensor
  );
  
  

INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Нет', 'Нет', 'Нет', 'Неизвестно'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Неизвестно', 'Неизвестно', 'Неизвестно');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 5),
    (SELECT id FROM gearboxes WHERE gearbox = 'МКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Полный 4WD'),
    (SELECT id FROM colors WHERE color = 'Белый'),
    (SELECT id FROM specifications WHERE id = 5),
    (SELECT id FROM tires WHERE id = 5),
    (SELECT id FROM brakes WHERE id = 5),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 5),
    (SELECT id FROM lights WHERE id = 5),
    (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
    (SELECT id FROM cabin_microclimate WHERE id = 5),
    (SELECT id FROM electric_options WHERE id = 5),
    (SELECT id FROM airbags WHERE id = 5),
    (SELECT id FROM multimedia_systems WHERE id = 5),
    '1.7 MT Luxe + Кондиционер', -- trim_level
    17, -- acceleration_0_to_100
    142, -- max_speed
    12.1, -- city_fuel_consumption
    8.3, -- highway_fuel_consumption
    9.9, -- mixed_fuel_consumption
    4, -- number_of_seats
    265, -- trunk_volume
    1285, -- mass
    'Есть' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.7 MT Luxe + Кондиционер'),
  '',
  685000,
  71538,
  ARRAY['/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva1.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva2.jpg', 
        '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva3.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva4.jpg', 
        '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva5.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva6.jpg', 
        '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva7.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva8.jpg', 
        '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva9.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva10.jpg', 
        '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva11.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva12.jpg', 
        '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva13.jpg', '/userID=3aa016ca-6582-4b2b-99ce-67160f589bee&carID=0ed1b714-d6cb-4acb-b18d-0469c6e91802&carName=lada_niva/niva14.jpg']);
  
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Great Wall', id FROM countries WHERE country = 'Китай';

INSERT INTO models (model, make_id) 
SELECT 'Hover H5', id FROM makes WHERE make = 'Great Wall';


INSERT INTO generations (model_id, generation) 
SELECT id, '1 поколение Hover H5' FROM models WHERE model = 'Hover H5';

INSERT INTO body_types (body)
VALUES ('Внедорожник')
ON CONFLICT DO NOTHING;

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Есть', 'Есть', 'Независимая, на двойных поперечных рычагах', 'Зависимая, пружинная');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '1 поколение Hover H5'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Внедорожник'),
  6, -- suspensions_id
  4649, -- length
  1810, -- width
  1735, -- height
  240, -- ground_clearance
  0.35, -- drag_coefficient
  1515, -- front_track_width
  1520, -- back_track_width
  2700, -- wheelbase
  2,  -- crach_test_estimate
  2011 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Дизель', 'Рядный, 4-цилиндровый', 1996, 150, '310 (32) / 2800');


INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (235, 235, 65, 65, 17, 17);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые вентилируемые', 'Ручной');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES (
    'Есть', -- abs_system
    'Неизвестно', -- esp_system
    'Есть',       -- ebd_system
    'Неизвестно',-- bas_system
    'Неизвестно', -- tcs_system
    'Неизвестно', -- front_parking_sensor
    'Есть', -- back_parking_sensor
    'Есть', -- rear_view_camera
    'Есть'  -- cruise_control
    );


INSERT INTO colors (color)
VALUES ('Серый')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Есть', 'Есть', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Кожаная')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Неизвестно', 'Есть');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Есть',-- electric_drive_of_driver_seat 
     'Неизвестно',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Есть'-- rain_sensor
  );
  


INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Неизвестно', 'Неизвестно'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Есть', 'Есть');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 6),
    (SELECT id FROM gearboxes WHERE gearbox = 'АКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Полный 4WD'),
    (SELECT id FROM colors WHERE color = 'Белый'),
    (SELECT id FROM specifications WHERE id = 6),
    (SELECT id FROM tires WHERE id = 6),
    (SELECT id FROM brakes WHERE id = 6),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 6),
    (SELECT id FROM lights WHERE id = 6),
    (SELECT id FROM interior_design WHERE upholstery = 'Кожаная'),
    (SELECT id FROM cabin_microclimate WHERE id = 6),
    (SELECT id FROM electric_options WHERE id = 6),
    (SELECT id FROM airbags WHERE id = 6),
    (SELECT id FROM multimedia_systems WHERE id = 6),
    '2.0 D AT Luxe', -- trim_level
    11, -- acceleration_0_to_100
    170, -- max_speed
    9.1, -- city_fuel_consumption
    7.8, -- highway_fuel_consumption
    8.6, -- mixed_fuel_consumption
    5, -- number_of_seats
    810, -- trunk_volume
    1880, -- mass
    'Неизвестно' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '2.0 D AT Luxe'),
  '',
  705000,
  190829,
  ARRAY['/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_1.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_2.jpg', 
        '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_3.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_4.jpg', 
        '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_5.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_6.jpg', 
        '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_7.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_8.jpg', 
        '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_9.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_10.jpg', 
        '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_11.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_12.jpg', 
        '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_13.jpg', '/userID=a58cb049-fd12-443c-9e5d-cfb16f5360ae&carID=780d8c76-3c0f-4001-81ae-4d29fa4efd84&carName=great_wall_hover_h5/hover_h5_14.jpg']);
  
COMMIT;

BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Land Rover', id FROM countries WHERE country = 'Великобритания';

INSERT INTO models (model, make_id) 
SELECT 'Freelander', id FROM makes WHERE make = 'Land Rover';


INSERT INTO generations (model_id, generation) 
SELECT id, '1 поколение рестайлинг Freelander' FROM models WHERE model = 'Freelander';

INSERT INTO body_types (body)
VALUES ('Внедорожник')
ON CONFLICT DO NOTHING;

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Неизвестно', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, амортизационная стойка типа МакФерсон');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '1 поколение рестайлинг Freelander'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Внедорожник'),
  7, -- suspensions_id
  4445, -- length
  1809, -- width
  1828, -- height
  185, -- ground_clearance
  0.39, -- drag_coefficient
  1545, -- front_track_width
  1545, -- back_track_width
  2557, -- wheelbase
  2, -- crash_test_estimate
  2005 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'V-образный, 6-цилиндровый', 2497, 177, '240 (24) / 4000');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (225, 225, 55, 55, 17, 17);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Барабанные', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES (
    'Есть', -- abs_system
    'Есть', -- esp_system
    'Есть',       -- ebd_system
    'Неизвестно',-- bas_system
    'Неизвестно', -- tcs_system
    'Неизвестно', -- front_parking_sensor
    'Есть', -- back_parking_sensor
    'Неизвестно', -- rear_view_camera
    'Неизвестно'  -- cruise_control
    );


INSERT INTO colors (color)
VALUES ('Красный')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Неизвестно', 'Есть', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Кожаная')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Неизвестно');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Неизвестно',-- electric_front_side_windows_lifts
     'Неизвестно', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Неизвестно',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Есть',-- electric_drive_of_driver_seat 
     'Есть',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Неизвестно'-- rain_sensor
  );
  


INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Неизвестно', 'Неизвестно'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Неизвестно', 'Неизвестно', 'Неизвестно');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 7),
    (SELECT id FROM gearboxes WHERE gearbox = 'АКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Полный 4WD'),
    (SELECT id FROM colors WHERE color = 'Красный'),
    (SELECT id FROM specifications WHERE id = 7),
    (SELECT id FROM tires WHERE id = 7),
    (SELECT id FROM brakes WHERE id = 7),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 7),
    (SELECT id FROM lights WHERE id = 7),
    (SELECT id FROM interior_design WHERE upholstery = 'Кожаная'),
    (SELECT id FROM cabin_microclimate WHERE id = 7),
    (SELECT id FROM electric_options WHERE id = 7),
    (SELECT id FROM airbags WHERE id = 7),
    (SELECT id FROM multimedia_systems WHERE id = 7),
    '2.5 AT 4WD HSE', -- trim_level
    10.1, -- acceleration_0_to_100
    182, -- max_speed
    17.2, -- city_fuel_consumption
    9.7, -- highway_fuel_consumption
    12.4, -- mixed_fuel_consumption
    5, -- number_of_seats
    546, -- trunk_volume
    1650, -- mass
    'Есть' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '2.5 AT 4WD HSE'),
  '',
  678000,
  259381,
  ARRAY['/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander1.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander2.jpg', 
        '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander3.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander4.jpg', 
        '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander5.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander6.jpg',
        '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander7.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander8.jpg', 
        '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander9.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander10.jpg', 
        '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander11.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander12.jpg', 
        '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander13.jpg', '/userID=f1135004-d575-4376-99c7-06bb5c06d0e5&carID=7420a923-038c-4ce9-9c48-734674b5e21c&carName=land_rover_freelander/freelander14.jpg']);
  
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Skoda', id FROM countries WHERE country = 'Чехия';

INSERT INTO models (model, make_id) 
SELECT 'Octavia', id FROM makes WHERE make = 'Skoda';


INSERT INTO generations (model_id, generation) 
SELECT id, '2 поколение рестайлинг (Octavia II)' FROM models WHERE model = 'Octavia';

INSERT INTO body_types (body)
VALUES ('Лифтбек');

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Неизвестно', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, многорычажная');


INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '2 поколение рестайлинг (Octavia II)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Электроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Лифтбек'),
  8, -- suspensions_id
  4569, -- length
  1769, -- width
  1462, -- height
  164, -- ground_clearance
  0.3, -- drag_coefficient
  1541, -- front_track_width
  1514, -- back_track_width
  2578, -- wheelbase
  3.78, -- crash_test_estimate
  2012 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1595, 102, '148 (15) / 3800');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (195, 195, 65, 65, 15, 15);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Ручной');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Опция производителя', 'Нет', 'Нет', 'Нет', 'Нет', 'Опция производителя', 'Неизвестно', 'Опция производителя');


INSERT INTO colors (color)
VALUES ('Серебристый')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Неизвестно', 'Опция производителя', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Есть');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Нет',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Нет',-- electric_drive_of_driver_seat 
     'Нет',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Нет'-- rain_sensor
  );
  
  

INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Нет', 'Нет'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Нет', 'Нет');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 8),
    (SELECT id FROM gearboxes WHERE gearbox = 'АКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
    (SELECT id FROM colors WHERE color = 'Серебристый'),
    (SELECT id FROM specifications WHERE id = 8),
    (SELECT id FROM tires WHERE id = 8),
    (SELECT id FROM brakes WHERE id = 8),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 8),
    (SELECT id FROM lights WHERE id = 8),
    (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
    (SELECT id FROM cabin_microclimate WHERE id = 8),
    (SELECT id FROM electric_options WHERE id = 8),
    (SELECT id FROM airbags WHERE id = 8),
    (SELECT id FROM multimedia_systems WHERE id = 8),
    '1.6 MPI AT Ambition', -- trim_level
    14.1, -- acceleration_0_to_100
    184, -- max_speed
    11.2, -- city_fuel_consumption
    6.1, -- highway_fuel_consumption
    7.9, -- mixed_fuel_consumption
    5, -- number_of_seats
    560, -- trunk_volume
    1315, -- mass
    'Опция производителя' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.6 MPI AT Ambition'),
  '',
  687000,
  297638,
  ARRAY['/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia1.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia2.jpg', 
        '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia3.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia4.jpg', 
        '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia5.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia6.jpg', 
        '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia7.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia8.jpg', 
        '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia9.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia10.jpg', 
        '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia11.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia12.jpg', 
        '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia13.jpg', '/userID=e9c72e1b-5b3c-465a-9732-eedc91082c9f&carID=045d740c-1da5-4d5e-ae23-8e1b03abc11a&carName=skoda_octavia/octavia14.jpg']);
  
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Ford', id FROM countries WHERE country = 'США';

INSERT INTO models (model, make_id) 
SELECT 'Mondeo', id FROM makes WHERE make = 'Ford';


INSERT INTO generations (model_id, generation) 
SELECT id, '4 поколение (Mk IV)' FROM models WHERE model = 'Mondeo';

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Неизвестно', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, многорычажная');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '4 поколение (Mk IV)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Седан'),
  9, -- suspensions_id
  4850, -- length
  1886, -- width
  1500, -- height
  130, -- ground_clearance
  0.31, -- drag_coefficient
  1588, -- front_track_width
  1605, -- back_track_width
  2850, -- wheelbase
  5, -- crash_test_estimate
  2010 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1596, 125, '166 (17) / 4100');


INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (205, 205, 55, 55, 16, 16);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Опция производителя', 'Есть', 'Есть', 'Есть', 'Опция производителя', 'Опция производителя', 'Нет', 'Нет');


INSERT INTO colors (color)
VALUES ('Серый')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Нет', 'Опция производителя', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Нет');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Нет', -- electric_back_side_windows_lifts
     'Опция производителя',-- electric_heating_of_front_seats
     'Нет',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Нет',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Нет',-- electric_drive_of_driver_seat 
     'Нет',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Нет'-- rain_sensor
  );
  
  

INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Есть', 'Есть'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Есть', 'Нет');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 9),
    (SELECT id FROM gearboxes WHERE gearbox = 'МКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
    (SELECT id FROM colors WHERE color = 'Серый'),
    (SELECT id FROM specifications WHERE id = 9),
    (SELECT id FROM tires WHERE id = 9),
    (SELECT id FROM brakes WHERE id = 9),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 9),
    (SELECT id FROM lights WHERE id = 9),
    (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
    (SELECT id FROM cabin_microclimate WHERE id = 9),
    (SELECT id FROM electric_options WHERE id = 9),
    (SELECT id FROM airbags WHERE id = 9),
    (SELECT id FROM multimedia_systems WHERE id = 9),
    '1.6 MT Ambiente', -- trim_level
    12.3, -- acceleration_0_to_100
    195, -- max_speed
    10.3, -- city_fuel_consumption
    5.7, -- highway_fuel_consumption
    7.4, -- mixed_fuel_consumption
    5, -- number_of_seats
    493, -- trunk_volume
    1435, -- mass
    'Опция производителя' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.6 MT Ambiente'),
  '',
  710000,
  143925,
  ARRAY['/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo1.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo2.jpg', 
        '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo3.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo4.jpg', 
        '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo5.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo6.jpg', 
        '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo7.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo8.jpg', 
        '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo9.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo10.jpg', 
        '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo11.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo12.jpg', 
        '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo13.jpg', '/userID=aef25a1c-998c-43ed-a720-92bfbbb0c6ed&carID=9efc00ce-c363-48a5-8768-fa9d4efe1e94&carName=ford_mondeo/mondeo14.jpg']);
  
COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'BMW', id FROM countries WHERE country = 'Германия';

INSERT INTO models (model, make_id) 
SELECT '7 Series', id FROM makes WHERE make = 'BMW';


INSERT INTO generations (model_id, generation) 
SELECT id, '4 поколение рестайлинг (E65)' FROM models WHERE model = '7 Series';

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Неизвестно', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, многорычажная');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '4 поколение рестайлинг (E65)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Седан'),
  10, -- suspensions_id
  5179, -- length
  1902, -- width
  1484, -- height
  147, -- ground_clearance
  0.3, -- drag_coefficient
  1579, -- front_track_width
  1596, -- back_track_width
  3128, -- wheelbase
  2, -- crash_test_estimate
  2006 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'V-образный, 8-цилиндровый', 4799, 367, '490 (50) / 3400');

INSERT INTO drive_types (drive)
VALUES ('Задний FR');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (245, 245, 55, 55, 17, 17);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Электронный');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Есть', 'Есть', 'Есть', 'Неизвестно', 'Есть', 'Есть', 'Неизвестно', 'Есть');


INSERT INTO colors (color)
VALUES ('Черный')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Биксеноновые фары', 'Нет', 'Нет', 'Есть', 'Есть', 'Есть');

INSERT INTO interior_design (upholstery)
VALUES ('Кожаная')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Есть');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Есть',-- electric_heating_of_back_seats
     'Опция производителя',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Неизвестно',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Есть',-- electric_drive_of_driver_seat 
     'Есть',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Опция производителя',-- electric_trunk_opener
     'Есть'-- rain_sensor
  );
  
  

INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Есть', 'Есть'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Неизвестно', 'Есть');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 10),
    (SELECT id FROM gearboxes WHERE gearbox = 'АКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Задний FR'),
    (SELECT id FROM colors WHERE color = 'Серебристый'),
    (SELECT id FROM specifications WHERE id = 10),
    (SELECT id FROM tires WHERE id = 10),
    (SELECT id FROM brakes WHERE id = 10),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 10),
    (SELECT id FROM lights WHERE id = 10),
    (SELECT id FROM interior_design WHERE upholstery = 'Кожаная'),
    (SELECT id FROM cabin_microclimate WHERE id = 10),
    (SELECT id FROM electric_options WHERE id = 10),
    (SELECT id FROM airbags WHERE id = 10),
    (SELECT id FROM multimedia_systems WHERE id = 10),
    '750Li AT', -- trim_level
    6, -- acceleration_0_to_100
    250, -- max_speed
    16.9, -- city_fuel_consumption
    8.3, -- highway_fuel_consumption
    11.4, -- mixed_fuel_consumption
    5, -- number_of_seats
    501, -- trunk_volume
    2025, -- mass
    'Есть' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '750Li AT'),
  '',
  725000,
  293168,
   ARRAY['/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series1.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series2.jpg',
         '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series3.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series4.jpg', 
         '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series5.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series6.jpg', 
         '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series7.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series8.jpg', 
         '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series9.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series10.jpg', 
         '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series11.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series12.jpg', 
         '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series13.jpg', '/userID=6d3fdc5d-2e8b-4536-a04c-86c68fe38e10&carID=c372d273-233e-425c-9804-48a62731c525&carName=bmw_7_series/7_series14.jpg']);
  
COMMIT;

BEGIN;

INSERT INTO makes (make, country_id) 
SELECT 'Mitsubishi', id FROM countries WHERE country = 'Япония';


INSERT INTO models (model, make_id) 
SELECT 'Lancer', id FROM makes WHERE make = 'Mitsubishi';

INSERT INTO generations (model_id, generation) 
SELECT id, '10 поколение (Evolution X)' FROM models WHERE model = 'Lancer';

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Есть', 'Есть', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, многорычажная');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '10 поколение (Evolution X)'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Седан'),
  11, -- suspensions_id
  4570, -- length
  1760, -- width
  1490, -- height
  150, -- ground_clearance
  0.35, -- drag_coefficient
  1530, -- front_track_width
  1530, -- back_track_width
  2635, -- wheelbase
  4.5, -- crach_test_estimate
  2008 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 1798, 143, '178 (18) / 4250');

INSERT INTO gearboxes (gearbox)
VALUES ('Вариатор');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (215, 215, 45, 45, 18, 18);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES ('Есть', 'Есть', 'Есть', 'Неизвестно', 'Есть', 'Неизвестно', 'Неизвестно', 'Неизвестно', 'Есть');


INSERT INTO colors (color)
VALUES ('Черный')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Есть', 'Есть', 'Есть');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Есть');

INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Нет',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Есть',-- electric_heating_of_rear_window
     'Есть', -- electric_heating_of_side_mirrors
     'Нет',-- electric_drive_of_driver_seat 
     'Нет',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Нет',-- electric_trunk_opener
     'Есть'-- rain_sensor
  );


INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Есть', 'Есть');
  
INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Есть', 'Есть', 'Есть');

INSERT INTO trim_levels (
  engine_id,
  gearbox_id,
  drive_type_id,
  color_id,
  specification_id,
  tires_id,
  brakes_id,
  safety_and_motion_control_systems_id,
  lights_id,
  interior_design_id,
  cabin_microclimate_id,
  electric_options_id,
  airbags_id,
  multimedia_systems_id, 
  trim_level,
  acceleration_0_to_100,
  max_speed,
  city_fuel_consumption,
  highway_fuel_consumption,
  mixed_fuel_consumption,
  number_of_seats,
  trunk_volume,
  mass,
  car_alarm)
VALUES
    ((SELECT id FROM engines WHERE id = 11),
     (SELECT id FROM gearboxes WHERE gearbox = 'Вариатор'),
     (SELECT id FROM drive_types WHERE drive = 'Передний FF'),
     (SELECT id FROM colors WHERE color = 'Черный'),
     (SELECT id FROM specifications WHERE id = 11),
     (SELECT id FROM tires WHERE id = 11),
     (SELECT id FROM brakes WHERE id = 11),
     (SELECT id FROM safety_and_motion_control_systems WHERE id = 11),
     (SELECT id FROM lights WHERE id = 11),
     (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
     (SELECT id FROM cabin_microclimate WHERE id = 11),
     (SELECT id FROM electric_options WHERE id = 11),
     (SELECT id FROM airbags WHERE id = 11),
     (SELECT id FROM multimedia_systems WHERE id = 11),
      '1.8 CVT Intense', -- trim_level
    11.2, -- acceleration_0_to_100
    192, -- max_speed
    10.9, -- city_fuel_consumption
    6.2, -- highway_fuel_consumption
    7.9, -- mixed_fuel_consumption
    5, -- number_of_seats
    377, -- trunk_volume
    1395, -- mass
    'Неизвестно' -- car_alarm
     );
  
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '1.8 CVT Intense'),
  '',
  657000,
  258617,
  ARRAY['/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer1.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer2.jpg', 
        '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer3.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer4.jpg', 
        '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer5.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer6.jpg', 
        '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer7.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer8.jpg', 
        '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer9.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer10.jpg', 
        '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer11.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer12.jpg', 
        '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer13.jpg', '/userID=537d0b8a-4037-40c1-8a66-32a497baa9f0&carID=6575d86a-76d4-4a26-a748-fe3d4054974e&carName=mitsubishi_lancer/lancer14.jpg']);

COMMIT;


BEGIN;
INSERT INTO makes (make, country_id) 
SELECT 'Opel', id FROM countries WHERE country = 'Германия';

INSERT INTO models (model, make_id) 
SELECT 'Antara', id FROM makes WHERE make = 'Opel';


INSERT INTO generations (model_id, generation) 
SELECT id, '1 поколение Antara' FROM models WHERE model = 'Antara';

INSERT INTO body_types (body)
VALUES ('Внедорожник')
ON CONFLICT DO NOTHING;

INSERT INTO suspensions (front_stabilizer, back_stabilizer, front_suspension, back_suspension) 
VALUES ('Неизвестно', 'Неизвестно', 'Независимая, амортизационная стойка типа МакФерсон', 'Независимая, многорычажная');

INSERT INTO specifications (generation_id, steering_wheel_position_id, power_steering_type_id, body_type_id, suspensions_id, length, width, height, ground_clearance, drag_coefficient, front_track_width, back_track_width, wheelbase, crash_test_estimate, year)
VALUES (
  (SELECT id FROM generations WHERE generation = '1 поколение Antara'),
  (SELECT id FROM steering_wheel_positions WHERE position = 'Левый руль'),
  (SELECT id FROM power_steering_types WHERE power_steering = 'Гидроусилитель руля'),
  (SELECT id FROM body_types WHERE body = 'Внедорожник'),
  12, -- suspensions_id
  4575, -- length
  1850, -- width
  1704, -- height
  200, -- ground_clearance
  0.3, -- drag_coefficient
  1578, -- front_track_width
  1574, -- back_track_width
  2707, -- wheelbase
  4, -- crash_test_estimate
  2007 -- year
);


INSERT INTO engines (fuel_used, engine_type, capacity, power, max_torque)
VALUES ('Бензин', 'Рядный, 4-цилиндровый', 2405, 140, '220 (22) / 2400');

INSERT INTO tires (back_tires_width, front_tires_width, front_tires_aspect_ratio, back_tires_aspect_ratio, front_tires_rim_diameter, back_tires_rim_diameter)
VALUES (235, 235, 60, 60, 17, 17);

INSERT INTO brakes (front_brakes, back_brakes, parking_brake)
VALUES ('Дисковые вентилируемые', 'Дисковые', 'Неизвестно');


INSERT INTO safety_and_motion_control_systems (abs_system, esp_system, ebd_system, bas_system, tcs_system, front_parking_sensor, back_parking_sensor, rear_view_camera, cruise_control)
VALUES (
    'Есть', -- abs_system
    'Есть', -- esp_system
    'Есть',       -- ebd_system
    'Неизвестно',-- bas_system
    'Есть', -- tcs_system
    'Опция производителя', -- front_parking_sensor
    'Опция производителя', -- back_parking_sensor
    'Есть', -- rear_view_camera
    'Опция производителя'  -- cruise_control
    );


INSERT INTO colors (color)
VALUES ('Серебристый')
ON CONFLICT DO NOTHING;

INSERT INTO lights (headlights, led_running_lights, led_tail_lights, light_sensor, front_fog_lights, back_fog_lights)
VALUES ('Галогенные фары', 'Нет', 'Нет', 'Неизвестно', 'Есть', 'Неизвестно');

INSERT INTO interior_design (upholstery)
VALUES ('Тканевая')
ON CONFLICT DO NOTHING;

INSERT INTO cabin_microclimate(air_conditioner, climate_control)
VALUES ('Есть', 'Неизвестно');


INSERT INTO electric_options(
  electric_front_side_windows_lifts,
  electric_back_side_windows_lifts,
  electric_heating_of_front_seats,
  electric_heating_of_back_seats,
  electric_heating_of_steering_wheel,
  electric_heating_of_windshield,
  electric_heating_of_rear_window,
  electric_heating_of_side_mirrors,
  electric_drive_of_driver_seat,
  electric_drive_of_front_seats,
  electric_drive_of_side_mirrors,
  electric_trunk_opener,
  rain_sensor)
  VALUES (
     'Есть',-- electric_front_side_windows_lifts
     'Есть', -- electric_back_side_windows_lifts
     'Есть',-- electric_heating_of_front_seats
     'Неизвестно',-- electric_heating_of_back_seats
     'Неизвестно',-- electric_heating_of_steering_wheel
     'Неизвестно',-- electric_heating_of_windshield
     'Неизвестно',-- electric_heating_of_rear_window
     'Есть',-- electric_heating_of_side_mirrors
     'Нет',-- electric_drive_of_driver_seat 
     'Нет',-- electric_drive_of_front_seats
     'Есть',-- electric_drive_of_side_mirrors
     'Неизвестно',-- electric_trunk_opener
     'Нет'-- rain_sensor
  );
  


INSERT INTO airbags (driver_airbag, front_passenger_airbag, side_airbags, curtain_airbags)
VALUES ('Есть', 'Есть', 'Есть', 'Есть'); 

INSERT INTO multimedia_systems (on_board_computer, mp3_support, hands_free_support)
VALUES ('Опция производителя', 'Есть', 'Опция производителя');


  INSERT INTO trim_levels (
    engine_id,
    gearbox_id,
    drive_type_id,
    color_id,
    specification_id,
    tires_id,
    brakes_id,
    safety_and_motion_control_systems_id,
    lights_id,
    interior_design_id,
    cabin_microclimate_id,
    electric_options_id,
    airbags_id,
    multimedia_systems_id, 
    trim_level,
    acceleration_0_to_100,
    max_speed,
    city_fuel_consumption,
    highway_fuel_consumption,
    mixed_fuel_consumption,
    number_of_seats,
    trunk_volume,
    mass,
    car_alarm
  )
  VALUES (
    (SELECT id FROM engines WHERE id = 12),
    (SELECT id FROM gearboxes WHERE gearbox = 'АКПП'),
    (SELECT id FROM drive_types WHERE drive = 'Полный 4WD'),
    (SELECT id FROM colors WHERE color = 'Серебристый'),
    (SELECT id FROM specifications WHERE id = 12),
    (SELECT id FROM tires WHERE id = 12),
    (SELECT id FROM brakes WHERE id = 12),
    (SELECT id FROM safety_and_motion_control_systems WHERE id = 12),
    (SELECT id FROM lights WHERE id = 12),
    (SELECT id FROM interior_design WHERE upholstery = 'Тканевая'),
    (SELECT id FROM cabin_microclimate WHERE id = 12),
    (SELECT id FROM electric_options WHERE id = 12),
    (SELECT id FROM airbags WHERE id = 12),
    (SELECT id FROM multimedia_systems WHERE id = 12),
    '2.4 AT Enjoy', -- trim_level
    12.4, -- acceleration_0_to_100
    170, -- max_speed
    14.1, -- city_fuel_consumption
    7.7, -- highway_fuel_consumption
    10.1, -- mixed_fuel_consumption
    5, -- number_of_seats
    420, -- trunk_volume
    1865, -- mass
    'Неизвестно' -- car_alarm
  );


 
INSERT INTO offerings (trim_level_id, description, price, kilometrage, photo_urls)
VALUES (
  (SELECT id FROM trim_levels WHERE trim_level = '2.4 AT Enjoy'),
  '',
  710000,
  238794,
  ARRAY['/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara1.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara2.jpg', 
        '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara3.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara4.jpg', 
        '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara5.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara6.jpg', 
        '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara7.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara8.jpg', 
        '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara9.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara10.jpg', 
        '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara11.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara12.jpg', 
        '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara13.jpg', '/userID=a8c7c916-1021-4587-9ead-fab83ba4a3b6&carID=ac550c52-6016-400d-bbf3-89efcf3e3efe&carName=opel_antara/antara14.jpg']);
  
COMMIT;