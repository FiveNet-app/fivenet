BEGIN;

-- Table: fivenet_job_props
CREATE TABLE IF NOT EXISTS `fivenet_job_props` (
  `job` varchar(20) NOT NULL,
  `theme` varchar(20) DEFAULT "default",
  UNIQUE KEY `idx_fivenet_job_props_unique` (`job`),
  CONSTRAINT `fk_fivenet_job_props_job` FOREIGN KEY (`job`) REFERENCES `jobs` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;
