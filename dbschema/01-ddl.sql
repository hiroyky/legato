-- MySQL Script generated by MySQL Workbench
-- Sat Jul 24 18:03:09 2021
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema legato_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema legato_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `legato_db` DEFAULT CHARACTER SET utf8 ;
USE `legato_db` ;

-- -----------------------------------------------------
-- Table `legato_db`.`album_artists`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `legato_db`.`album_artists` ;

CREATE TABLE IF NOT EXISTS `legato_db`.`album_artists` (
  `album_artist_id` INT NOT NULL AUTO_INCREMENT,
  `name` TEXT NOT NULL,
  `name_hash` CHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`album_artist_id`),
  UNIQUE INDEX `name_hash_UNIQUE` (`name_hash` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `legato_db`.`albums`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `legato_db`.`albums` ;

CREATE TABLE IF NOT EXISTS `legato_db`.`albums` (
  `album_id` INT NOT NULL AUTO_INCREMENT,
  `name` TEXT NOT NULL,
  `album_artist_id` INT NOT NULL,
  `disc_no` INT NOT NULL,
  `disc_total` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`album_id`),
  INDEX `fk_albums_album_artists_album_artist_id_idx` (`album_artist_id` ASC) VISIBLE,
  CONSTRAINT `fk_albums_album_artists_album_artist_id`
    FOREIGN KEY (`album_artist_id`)
    REFERENCES `legato_db`.`album_artists` (`album_artist_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `legato_db`.`genres`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `legato_db`.`genres` ;

CREATE TABLE IF NOT EXISTS `legato_db`.`genres` (
  `genre_id` INT NOT NULL AUTO_INCREMENT,
  `name` TEXT NOT NULL,
  `name_hash` CHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`genre_id`),
  UNIQUE INDEX `name_hash_UNIQUE` (`name_hash` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `legato_db`.`tracks`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `legato_db`.`tracks` ;

CREATE TABLE IF NOT EXISTS `legato_db`.`tracks` (
  `track_id` INT NOT NULL AUTO_INCREMENT,
  `title` TEXT NOT NULL,
  `artist` TEXT NOT NULL,
  `composer` TEXT NOT NULL,
  `track_no` INT NOT NULL,
  `lyrics` TEXT NOT NULL,
  `comment` TEXT NOT NULL,
  `year` INT NOT NULL,
  `file_path` TEXT NOT NULL,
  `file_path_hash` CHAR(128) NOT NULL,
  `file_hash` CHAR(128) NOT NULL,
  `format` TEXT NOT NULL,
  `file_type` TEXT NOT NULL,
  `album_artist_id` INT NOT NULL,
  `album_id` INT NOT NULL,
  `genre_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`track_id`),
  INDEX `fk_tracks_albums_album_id_idx` (`album_id` ASC) VISIBLE,
  INDEX `fk_tracks_album_artists_album_artist_id_idx` (`album_artist_id` ASC) VISIBLE,
  INDEX `fk_tracks_genres_genre_id_idx` (`genre_id` ASC) VISIBLE,
  UNIQUE INDEX `file_path_md5_hash_UNIQUE` (`file_path_hash` ASC) VISIBLE,
  CONSTRAINT `fk_tracks_albums_album_id`
    FOREIGN KEY (`album_id`)
    REFERENCES `legato_db`.`albums` (`album_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tracks_album_artists_album_artist_id`
    FOREIGN KEY (`album_artist_id`)
    REFERENCES `legato_db`.`album_artists` (`album_artist_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tracks_genres_genre_id`
    FOREIGN KEY (`genre_id`)
    REFERENCES `legato_db`.`genres` (`genre_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
