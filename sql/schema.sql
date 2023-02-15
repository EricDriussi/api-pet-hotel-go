CREATE TABLE bookings
(
    id       VARCHAR(255) NOT NULL,
    pet_name     VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
