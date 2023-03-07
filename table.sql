CREATE TABLE `mahasiswa`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `nama` VARCHAR(100) NOT NULL,
    `jurusan` VARCHAR(25) NOT NULL,
    `nim` CHAR(9) NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB;
