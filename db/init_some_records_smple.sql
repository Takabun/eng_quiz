LOCK TABLES `questions` WRITE;
/*!40000 ALTER TABLE `questions` DISABLE KEYS */;

INSERT INTO `questions` (`id`, `text`, `user`, `default_image`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'TESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTESTTEST','test',1,'2020-09-06 16:12:39','2020-09-06 16:12:39',NULL),
	(2,'222','test',3,'2020-09-06 16:16:28','2020-09-06 16:16:28',NULL),
	(3,'3333','test',2,'2020-09-06 16:16:16','2020-09-06 15:15:15',NULL),
	(4,'444','test',1,'2020-09-16 16:16:16','2020-09-16 16:16:16',NULL),
	(5,'testestestestestestestestestestestestestestestestestestestes','test2',1,'2020-09-16 16:16:16','2020-09-16 16:16:16',NULL);

/*!40000 ALTER TABLE `questions` ENABLE KEYS */;
UNLOCK TABLES;


LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;

INSERT INTO `tags` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'Java','2020-09-06 16:12:39','2020-09-06 16:16:28',NULL),
	(2,'Javascript','2020-09-06 16:12:39','2020-09-06 16:12:39',NULL),
	(3,'Vue.js','2020-09-06 16:14:57','2020-09-06 16:14:57',NULL),
	(4,NULL,NULL,NULL,NULL);

/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;


LOCK TABLES `question_tags` WRITE;
/*!40000 ALTER TABLE `question_tags` DISABLE KEYS */;

INSERT INTO `question_tags` (`id`, `question_id`, `tag_id`)
VALUES
	(1,1,1),
	(2,1,2),
	(3,2,1),
	(4,3,3),
	(5,4,2);

/*!40000 ALTER TABLE `question_tags` ENABLE KEYS */;
UNLOCK TABLES;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
