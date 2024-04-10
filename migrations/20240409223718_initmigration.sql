-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Banners (
                        id SERIAL PRIMARY KEY,
                        data VARCHAR(255),
                        is_enabled BOOLEAN,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP,
                        feature_id INT
);

CREATE TABLE Tags (
                         id SERIAL PRIMARY KEY,
                         title VARCHAR(255),
                         created_at TIMESTAMP,
                         updated_at TIMESTAMP
);

CREATE TABLE Features (
                         id SERIAL PRIMARY KEY,
                         title VARCHAR(255),
                         created_at TIMESTAMP,
                         updated_at TIMESTAMP
);

CREATE TABLE BannerTags (
                        id SERIAL PRIMARY KEY,
                        banner_id INT NOT NULL,
                        tag_id INT NOT NULL
);

CREATE TABLE Users (
                        id SERIAL PRIMARY KEY,
                        email VARCHAR(255),
                        name VARCHAR(255),
                        hashed_password VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP,
                        updated_at TIMESTAMP

);
--
ALTER TABLE Banners ADD CONSTRAINT "fk_feature" FOREIGN KEY("feature_id") REFERENCES Features("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE BannerTags ADD CONSTRAINT "fk_banner" FOREIGN KEY("banner_id") REFERENCES Banners("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE BannerTags ADD CONSTRAINT "fk_tag" FOREIGN KEY("tag_id") REFERENCES Tags("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE Banners DROP CONSTRAINT "fk_feature";
ALTER TABLE BannerTags DROP CONSTRAINT "fk_banner";
ALTER TABLE BannerTags DROP CONSTRAINT "fk_tag";
DROP TABLE BannerTags;
DROP TABLE Features;
DROP TABLE Tags;
DROP TABLE Banners;
-- +goose StatementEnd
