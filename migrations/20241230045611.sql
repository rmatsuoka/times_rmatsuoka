-- Create "deleted_users" table
CREATE TABLE `deleted_users` (`user_id` integer NOT NULL, `code` varchar NOT NULL, `name` varchar NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, `deleted_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP));
-- Create "usercodes" table
CREATE TABLE `usercodes` (`code` varchar NOT NULL);
-- Create index "usercodes_code" to table: "usercodes"
CREATE UNIQUE INDEX `usercodes_code` ON `usercodes` (`code`);
