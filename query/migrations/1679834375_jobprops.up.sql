BEGIN;

-- Table: fivenet_job_props
CREATE TABLE IF NOT EXISTS `fivenet_job_props` (
  `job` varchar(20) NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3),
  `theme` varchar(20) DEFAULT "default",
  `livemap_marker_color` char(6) DEFAULT "5C7AFF",
  `quick_buttons` varchar(255) DEFAULT NULL,
  UNIQUE KEY `idx_fivenet_job_props_unique` (`job`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;
