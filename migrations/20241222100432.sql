-- Create "channels" table
CREATE TABLE `channels` (`id` integer NOT NULL, `code` varchar NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
-- Create index "channels_code" to table: "channels"
CREATE UNIQUE INDEX `channels_code` ON `channels` (`code`);
-- Create "members" table
CREATE TABLE `members` (`channel_id` integer NOT NULL, `user_id` integer NOT NULL, `role` integer NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), CONSTRAINT `0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "members_channel_id_user_id" to table: "members"
CREATE UNIQUE INDEX `members_channel_id_user_id` ON `members` (`channel_id`, `user_id`);
-- Create index "members_channel_id" to table: "members"
CREATE INDEX `members_channel_id` ON `members` (`channel_id`);
-- Create index "members_user_id" to table: "members"
CREATE INDEX `members_user_id` ON `members` (`user_id`);
-- Create "messages" table
CREATE TABLE `messages` (`id` integer NOT NULL, `user_id` integer NOT NULL, `channel_id` integer NOT NULL, `text` text NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), `modified_at` timestamp NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT `1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "messages_created_at" to table: "messages"
CREATE INDEX `messages_created_at` ON `messages` (`created_at`);
-- Create index "messages_channel_id" to table: "messages"
CREATE INDEX `messages_channel_id` ON `messages` (`channel_id`);
-- Create "users" table
CREATE TABLE `users` (`id` integer NOT NULL, `code` varchar NOT NULL, `name` varchar NOT NULL, `created_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), `updated_at` timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP), PRIMARY KEY (`id`));
-- Create index "users_code" to table: "users"
CREATE UNIQUE INDEX `users_code` ON `users` (`code`);
