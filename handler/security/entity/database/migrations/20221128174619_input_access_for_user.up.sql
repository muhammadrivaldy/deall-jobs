-- user_type_id meaning
    -- 1 is admin user
    -- 2 is regular user

-- api_id
    -- 2 is create user
    -- 3 is get user by id
    -- 4 is edit user
    -- 5 is edit password user
    -- 6 is remove user
    -- 7 is login
    -- 8 is refresh jwt

INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (2, 3, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (2, 7, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (2, 8, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 2, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 3, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 4, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 5, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 6, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 7, now(), now());
INSERT INTO `mst_access` (`user_type_id`, `api_id`, `created_at`, `updated_at`) VALUES (1, 8, now(), now());