CREATE TABLE  `user` (`id` bigint NOT NULL AUTO_INCREMENT,
                      `user_id` bigint NOT NULL,
                      `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                      `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                      `email`  varchar(64) COLLATE utf8mb4_general_ci,
                      `gender` tinyint NOT NULL DEFAULT '0',
                      `create_time`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                      `update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      PRIMARY KEY (`id` ),
                      UNIQUE KEY `idx_username` (`username`) USING BTREE,
                      UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;