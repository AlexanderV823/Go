-- https://dev.mysql.com/doc/refman/8.0/en/select.html

DROP DATABASE IF EXISTS messenger;
CREATE DATABASE messenger;
USE messenger;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` SERIAL PRIMARY KEY, -- SERIAL = BIGINT UNSIGNED NOT NULL AUTO_INCREMENT UNIQUE
    `firstname` VARCHAR(100),
    `lastname` VARCHAR(100) COMMENT 'Фамилия', 
    `email` VARCHAR(120) UNIQUE,
    `is_deleted` BIT DEFAULT b'0'
);

INSERT INTO `users` (`id`, `firstname`, `lastname`, `email`, `is_deleted`) 
VALUES
(1,'Kristoffer','Okuneva','qcorkery@example.net',b'0'),
(2,'Shana','Stokes','leuschke.domingo@example.org',b'0'),
(3,'Felicia','Rempel','stanton.roselyn@example.org',b'1'),
(4,'Janie','Tromp','marks.eric@example.com',b'0'),
(5,'Alize','Kohler','easton.lakin@example.net',b'0'),
(6,'Sabrina','Rolfson','gboyer@example.net',b'0'),
(7,'Lessie','Gorczany','montana.zboncak@example.net',b'0'),
(8,'Daija','Turner','fbernier@example.com',b'0'),
(9,'Kariane','Roberts','lberge@example.org',b'1'),
(10,'Darion','Homenick','jeffery.yundt@example.net',b'1'),
(11,'Soledad','Prosacco','lbernier@example.com',b'1'),
(12,'Trudie','Watsica','kolby.powlowski@example.net',b'1'),
(13,'Abel','Orn','estelle.ortiz@example.com',b'0'),
(14,'Connie','Bins','margarett02@example.com',b'0'),
(15,'Paige','Schaden','bettye96@example.net',b'0'),
(16,'Emmet','Kuhic','rutherford.abby@example.net',b'0'),
(17,'Jayda','Schaefer','nikolaus.edgar@example.org',b'0'),
(18,'Norbert','Hirthe','kattie08@example.org',b'0'),
(19,'Charlene','Volkman','savion78@example.com',b'0'),
(20,'Margaret','Deckow','ydickinson@example.net',b'1'),
(21,'Heber','Pfeffer','psporer@example.org',b'1'),
(22,'Dina','Leffler','zjerde@example.com',b'1'),
(23,'Alisa','Larson','denesik.wendell@example.net',b'0'),
(24,'Hyman','Murazik','tabitha92@example.net',b'1'),
(25,'Marie','Monahan','hschumm@example.com',b'1'),
(26,'Jon','Mante','dulce03@example.org',b'0'),
(27,'Caitlyn','Balistreri','veda.zulauf@example.org',b'1'),
(28,'Brock','Langworth','adams.mayra@example.net',b'0'),
(29,'Hailey','Trantow','retha.koelpin@example.net',b'0'),
(30,'Mckenna','Macejkovic','cullen07@example.org',b'1'),
(31,'Missouri','Torphy','lschneider@example.org',b'0'),
(32,'Lyric','Feil','lew60@example.com',b'0'),
(33,'Lucious','Breitenberg','maeve45@example.org',b'0'),
(34,'Sophie','Daniel','mariam.jenkins@example.org',b'1'),
(35,'Tianna','Nolan','jakubowski.jamar@example.org',b'0'),
(36,'Lavern','Abshire','hhirthe@example.org',b'0'),
(37,'Guido','Haag','lavina.raynor@example.org',b'0'),
(38,'Howard','Murray','rau.beulah@example.com',b'1'),
(39,'Maximillian','Leannon','gregory04@example.org',b'0'),
(40,'Brannon','Marquardt','sandrine.pfeffer@example.org',b'1'),
(41,'Morgan','Orn','wilford.bailey@example.com',b'0'),
(42,'Brady','Mertz','torp.mario@example.org',b'1'),
(43,'Consuelo','Hackett','ondricka.lizzie@example.net',b'1'),
(44,'Annalise','Zboncak','gerlach.valerie@example.org',b'1'),
(45,'Liana','Conroy','lyric29@example.net',b'1'),
(46,'Gennaro','Carroll','lfay@example.org',b'0'),
(47,'Abraham','Labadie','szboncak@example.net',b'1'),
(48,'Patsy','Ruecker','gretchen.roob@example.org',b'0'),
(49,'Rafaela','Beer','huel.sedrick@example.org',b'1'),
(50,'Yasmine','Konopelski','georgiana27@example.com',b'0'),
(51,'Albina','Balistreri','kameron.dickens@example.com',b'0'),
(52,'Hyman','Cremin','jany12@example.net',b'0'),
(53,'Meagan','Corwin','cturner@example.com',b'0'),
(54,'Justen','Sporer','violette27@example.org',b'1'),
(55,'Ena','Koelpin','kristina.langworth@example.org',b'0'),
(56,'Marilie','Goodwin','pbartoletti@example.net',b'1'),
(57,'Eusebio','Davis','keith.kuvalis@example.com',b'0'),
(58,'Caitlyn','Fritsch','ronny50@example.net',b'1'),
(59,'Treva','Heidenreich','lind.coleman@example.net',b'1'),
(60,'Elnora','Simonis','arvilla22@example.net',b'0'),
(61,'Ulices','Schuppe','zsteuber@example.net',b'1'),
(62,'Candido','Zemlak','lhodkiewicz@example.org',b'1'),
(63,'Julie','Mertz','ilabadie@example.org',b'1'),
(64,'Brandon','Fritsch','astrosin@example.org',b'1'),
(65,'Brad','Roob','lehner.earnestine@example.com',b'1'),
(66,'Lia','Jaskolski','janie15@example.com',b'0'),
(67,'Moshe','Leffler','orunolfsdottir@example.com',b'1'),
(68,'Reyna','Kling','zula.brekke@example.org',b'1'),
(69,'Magali','Hahn','keegan43@example.net',b'1'),
(70,'Alan','Walter','lueilwitz.petra@example.net',b'0'),
(71,'Dale','Gutmann','austin.schinner@example.net',b'1'),
(72,'Javonte','Hayes','runte.emma@example.com',b'1'),
(73,'Yolanda','Williamson','orobel@example.org',b'1'),
(74,'Desiree','Lang','reginald.predovic@example.com',b'1'),
(75,'Blair','Parisian','april.medhurst@example.net',b'1'),
(76,'Trever','Heaney','smitham.alena@example.net',b'0'),
(77,'Janessa','Bergstrom','zbeatty@example.org',b'0'),
(78,'Quinten','Stracke','lesch.dina@example.net',b'1'),
(79,'Alfreda','Shanahan','kirstin.kirlin@example.net',b'1'),
(80,'Harmony','Gislason','nwisoky@example.net',b'1'),
(81,'Laura','Spinka','eileen.emmerich@example.com',b'1'),
(82,'Althea','Wintheiser','ocie.schulist@example.net',b'1'),
(83,'Salvador','Lesch','russel.luther@example.net',b'0'),
(84,'Jarrell','Mitchell','lbernhard@example.net',b'1'),
(85,'Eliane','Pagac','idella69@example.org',b'1'),
(86,'Aletha','Runolfsdottir','schowalter.imogene@example.net',b'1'),
(87,'Domenick','Grant','kenyatta51@example.net',b'0'),
(88,'Loyal','Medhurst','hyatt.deion@example.net',b'0'),
(89,'Candido','Mayer','lavonne52@example.com',b'0'),
(90,'Aurore','Haley','maybelle25@example.net',b'0'),
(91,'Vernice','Renner','narciso.blanda@example.com',b'0'),
(92,'Shana','Gibson','eo\'conner@example.net',b'1'),
(93,'Enrico','Walsh','annette83@example.org',b'0'),
(94,'Kenyatta','Vandervort','schaden.cleveland@example.org',b'0'),
(95,'Forrest','Lubowitz','tre12@example.net',b'1'),
(96,'Ian','Price','iwilderman@example.com',b'1'),
(97,'Chadd','Ritchie','germaine.brakus@example.org',b'1'),
(98,'Antonina','Prohaska','hester.frami@example.net',b'1'),
(99,'Grace','Wintheiser','ruecker.kenneth@example.com',b'1'),
(100,'Kathryn','Herman','sgerhold@example.com',b'1');


-- 1-М
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
	`id` SERIAL PRIMARY KEY,
	`from_user_id` BIGINT UNSIGNED,
    `to_user_id` BIGINT UNSIGNED NOT NULL,
    `body` TEXT,
    `created_at` DATETIME DEFAULT NOW(), 
	`updated_at` DATETIME ON UPDATE NOW(),
    FOREIGN KEY (`from_user_id`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`to_user_id`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO `messages` (`id`, `from_user_id`, `to_user_id`, `body`,`created_at`, `updated_at`)
VALUES
(1,11,29,'Sed eveniet assumenda nisi quis. Quasi quam vitae rerum qui. Debitis earum rerum vitae et. Et hic est quo. Dolore nesciunt nulla suscipit quos sint quis.','2017-04-30 17:53:11','1971-03-28 07:52:05'),
(2,4,5,'Illum provident sit eveniet consequatur. Inventore vitae et ea perferendis recusandae. Quis suscipit nam eveniet reprehenderit veniam sequi qui.','1971-12-11 21:32:47','2008-07-01 17:00:40'),
(3,93,1,'Et qui eligendi quae quis porro ea. Aut consectetur soluta officia nam doloribus ut voluptatum. Soluta dolores itaque sit veritatis et provident. Et optio rerum deserunt incidunt amet.','2025-03-12 17:40:53','1989-04-24 19:57:39'),
(4,32,23,'Mollitia numquam magnam dolorem doloribus. Esse voluptatibus unde et. Et eum voluptas unde eius. Aliquam deleniti iusto laborum qui nulla.','1993-11-12 04:27:57','2008-06-24 00:58:59'),
(5,72,18,'Occaecati ut cum corrupti rerum laborum repudiandae rem. Non praesentium aut fugit. Voluptatum ea ut aliquid et qui id.','2023-07-29 11:31:55','2017-05-09 17:29:28'),
(6,31,47,'Recusandae quaerat exercitationem ipsa culpa aut autem voluptas. Ipsum occaecati qui nam enim veniam tempore. Cum est quisquam optio ipsa. Voluptates dolores dignissimos qui voluptatem et veritatis saepe accusantium.','1984-11-12 09:42:34','1983-12-21 01:29:08'),
(7,NULL,40,'Incidunt blanditiis voluptatibus et at hic enim occaecati dolorem. Itaque ad eum sed iste animi. Iure eos quis non animi in explicabo.','2013-06-15 15:12:55','1979-10-23 23:30:18'),
(8,35,24,'Veritatis veniam porro nostrum. Ex tempora magnam qui atque perferendis inventore impedit tenetur. Impedit sit aut neque nihil cupiditate omnis. Nostrum et nihil voluptas occaecati et nisi officia.','2006-04-28 04:52:26','2020-02-16 15:59:13'),
(9,52,1,'Vel nam mollitia voluptatem. Autem et dolor pariatur atque quia rerum quibusdam. Quo est laboriosam quo atque ea facere provident.','2006-08-30 08:37:23','2005-12-11 20:32:45'),
(10,69,35,'Repellendus itaque pariatur commodi dolores. Odio et quia nobis et dolorum ut. Consectetur consequuntur quam quo.','2023-12-02 23:40:31','2006-07-01 14:53:07'),
(11,NULL,53,'Odit molestiae iusto delectus magnam quasi. Cum fugiat earum voluptatem ratione quis et. Modi temporibus nihil sapiente dolor delectus. Vel non sapiente quibusdam sunt ut.','2009-08-30 15:12:03','2008-12-29 13:07:17'),
(12,16,9,'Ut molestiae magni sed illo. Ab nemo sint in omnis id tenetur. Veniam debitis iste et cumque eveniet quam iusto.','2009-05-20 09:14:36','2000-12-14 09:49:45'),
(13,71,77,'Adipisci corporis repellendus sed blanditiis. Possimus incidunt temporibus aut ut. Id nihil et reprehenderit maiores. Minima exercitationem aliquam inventore ut. Voluptatem non eaque facere voluptatem quia.','1979-10-26 05:30:51','1971-08-10 07:33:02'),
(14,14,52,'Earum sed cum ut recusandae. Eum at fugit vitae officiis nobis excepturi quo. Ea eum corporis sed quia. Molestias dolores officiis voluptate dolores quibusdam.','1983-04-16 22:20:10','1986-11-30 09:03:05'),
(15,39,33,'Atque excepturi qui nobis praesentium. Natus quidem in accusantium ut. Accusamus incidunt fugiat et voluptatem delectus blanditiis quidem. Expedita beatae porro dolor minus assumenda.','1991-05-28 18:01:22','2023-07-30 17:24:35'),
(16,64,46,'Eveniet ut repellendus necessitatibus id ratione nemo iusto est. Et mollitia eum aut velit sunt non autem. Optio aperiam nesciunt aut animi. Unde ad sint nostrum sint id. Qui rerum magnam minima dignissimos quis.','2001-05-26 05:12:20','2011-04-18 20:56:10'),
(17,NULL,81,'Cumque quaerat est ut quia occaecati nam. Sint odit ipsam mollitia tempora. Qui eligendi deleniti velit aut adipisci nihil voluptatem. Ut eum sed sunt molestias.','1980-10-25 03:12:59','1985-07-31 11:15:33'),
(18,52,24,'Sequi possimus voluptatem molestiae quisquam rerum sed dolor. Et non architecto aperiam voluptatem. Temporibus omnis voluptates dolore ut. Officia corporis aut illo repudiandae. Cum error harum vero qui molestias id inventore.','2023-09-10 02:13:11','2023-03-30 03:58:43'),
(19,5,72,'Amet necessitatibus occaecati asperiores omnis ut consequuntur dicta. Et autem reiciendis aut nisi cupiditate. Modi quae sit voluptas repudiandae nulla exercitationem.','1983-09-20 05:03:57','1984-12-12 04:19:20'),
(20,63,70,'Nam quibusdam et et voluptatem et iusto amet. Et eligendi dolor illum omnis nihil quibusdam. Recusandae enim dolores nesciunt voluptas exercitationem. Suscipit aut eos perspiciatis.','1978-11-28 18:19:27','2021-11-13 20:03:00'),
(21,NULL,85,'In quo aut molestiae autem ullam mollitia. Ut quos nam dolor tempora sed voluptatibus. Nobis sunt maxime a delectus. Laboriosam aliquam deserunt dolorem sit neque voluptas commodi.','1989-05-01 18:14:42','2019-04-14 00:44:18'),
(22,40,29,'Modi quae distinctio possimus rerum. Qui voluptatem repellat corrupti porro ducimus cum. At itaque est aut consequuntur voluptatum. Facilis neque suscipit consequatur.','2007-10-29 20:28:27','1988-03-13 15:27:38'),
(23,73,43,'Mollitia distinctio facilis praesentium temporibus. Recusandae recusandae praesentium tempore ratione nulla sit ut. Consequatur qui delectus repellat minima.','2004-07-15 09:34:37','2001-09-27 15:03:23'),
(24,50,58,'Ducimus repellat et impedit maxime. Non aut nostrum labore velit itaque distinctio. Eum libero accusamus quis doloremque. Nihil accusamus labore aut a error repellat.','2014-01-20 06:53:06','2019-01-17 16:24:53'),
(25,24,79,'Corporis voluptatem impedit est mollitia enim repellendus reprehenderit cupiditate. Qui dolorem ut et sit. Voluptatibus sit error enim.','2023-08-23 22:03:34','2014-01-07 19:39:42'),
(26,2,46,'Distinctio inventore commodi cum sint id nihil animi architecto. Assumenda nobis minima nesciunt harum ut eius.','1997-12-29 07:17:49','2004-10-20 02:54:38'),
(27,77,83,'Maxime dolor quasi cupiditate autem eum molestiae illum. Voluptatem incidunt possimus rem. Laboriosam voluptate quia et qui necessitatibus.','1974-11-03 21:49:13','2012-08-15 20:29:19'),
(28,NULL,25,'Sunt harum at omnis et. Totam quia earum ut nihil sapiente. Velit ut voluptates qui doloremque dolorem ipsum possimus. Rem aut explicabo totam dolore.','2019-01-01 17:44:11','1971-09-17 08:33:29'),
(29,7,47,'Eaque voluptatem omnis dolor sint praesentium qui placeat et. Qui ipsum quam est repellendus quia dolorem. Tempora similique nobis ut rerum numquam doloremque quibusdam. Iusto quo hic doloribus.','1971-01-16 16:57:57','2023-06-08 01:58:03'),
(30,50,14,'Fugit aliquam ab recusandae et sint. Est quo cum sed tenetur accusantium. Hic aut rerum tenetur sed suscipit quam corrupti autem. Perspiciatis tempore dolores quas nulla praesentium voluptates delectus.','2003-06-02 11:37:52','1997-07-22 01:09:34'),
(31,48,28,'Enim nihil commodi veniam sint saepe. Sit temporibus vel velit velit dolor aliquid consequuntur. Tenetur non alias voluptas numquam unde optio. Blanditiis dolor enim a quia eum doloremque deleniti.','2009-10-30 01:34:44','2017-09-01 07:15:50'),
(32,36,32,'Enim ullam minima possimus aliquid aut. Quasi earum repellat minima corporis cumque voluptas voluptatem. Voluptas dolores nihil ad est placeat. Excepturi ut sed aut sequi quis explicabo. Earum consequatur aspernatur laborum rerum molestiae.','2014-10-18 09:51:38','1992-08-16 01:47:25'),
(33,61,40,'Voluptatem dolores omnis quasi rerum. Ipsa necessitatibus amet sequi reprehenderit inventore repellendus. Possimus perferendis rerum excepturi rerum esse. Voluptas voluptatem earum nemo ipsam ut maiores velit unde.','2008-11-15 11:57:21','2009-11-04 17:30:27'),
(34,NULL,84,'Quod alias aut voluptas ex nobis fugiat. Molestiae ut eos et eum et. Sint eum voluptatem dignissimos est maxime. Unde aut quam non ullam nostrum.','2024-10-28 18:49:46','1973-12-23 08:04:43'),
(35,48,88,'Debitis et cupiditate pariatur doloribus quia quo. Cupiditate eum et aut maxime ab. Ea iusto pariatur quos officia. Mollitia tempore dolores sapiente autem.','2018-09-16 11:40:10','2024-12-25 18:01:12'),
(36,100,8,'Recusandae et dolores aperiam aut. Vero sit quos voluptatem minima qui. Mollitia ut est ut ad. Quis suscipit eum assumenda dicta dignissimos quo quia. Porro autem quaerat excepturi voluptatum culpa quasi.','2010-10-14 03:58:22','1997-05-02 09:09:23'),
(37,17,71,'Atque et ipsum enim delectus maiores. Consequatur tenetur omnis minus veritatis eum sint. Et commodi voluptatem tenetur reiciendis. Et ut nam aliquid amet harum.','2024-04-15 05:07:03','1998-04-18 11:43:40'),
(38,16,89,'Corrupti dolorum suscipit sit ut. Ullam est suscipit beatae alias possimus doloremque nihil. Asperiores consectetur perspiciatis rerum voluptatem id vel sit quia.','1970-02-12 23:27:29','1999-07-14 22:09:37'),
(39,72,53,'Quia non magnam qui. Porro et tempora et veniam dolore repudiandae porro quos. Enim veniam qui velit dolore a.','2014-07-08 07:36:05','1999-10-08 10:00:21'),
(40,41,84,'Sint nostrum in illo dolores aut odit. Iure nihil voluptatem et fuga.','1992-05-19 20:06:50','2011-05-04 18:00:41'),
(41,70,7,'Est vitae optio laudantium omnis id vero. Ex qui et molestias placeat placeat magni. Praesentium sequi quod molestiae alias et repellat.','1991-08-17 16:37:28','1995-08-19 01:31:00'),
(42,24,67,'Porro rerum tempore aspernatur animi quo dicta et. Nostrum doloribus molestiae sit qui ut qui iure eligendi. Commodi ratione qui alias provident voluptates blanditiis.','2003-07-05 09:26:46','2001-02-14 16:30:50'),
(43,49,19,'Consequatur et unde dolorem sit repellat. Fugiat nesciunt neque in voluptas vel nemo. Aliquid velit atque eaque aliquid. Numquam ullam dolorum dolorum est.','1991-05-18 02:40:05','1996-03-26 03:39:27'),
(44,29,41,'Pariatur sint porro ut. Beatae ad necessitatibus illo at. Quidem minus aut impedit et voluptas modi.','2019-06-06 10:35:05','1985-12-10 12:44:53'),
(45,NULL,81,'Et molestiae facilis alias magnam porro atque animi. Quo architecto tempore saepe voluptatem repudiandae impedit et nihil. Eum ut sit dolore est neque enim quas.','2003-05-12 13:26:25','2015-01-04 06:15:55'),
(46,15,62,'Nobis mollitia impedit voluptatem veniam sunt. Beatae illum minus et nisi consequatur qui. Maxime soluta culpa animi est in ex enim officiis.','1973-11-16 07:21:46','1989-11-23 04:11:04'),
(47,NULL,42,'Itaque sunt ducimus qui enim quas doloribus. Repellendus explicabo velit sint labore qui. Sed non deserunt ea veniam optio omnis voluptatem.','1998-06-27 06:19:31','2003-07-03 15:04:49'),
(48,59,12,'Eum reiciendis illo et eos fugit a. Ipsam et in consequuntur ratione. Ratione omnis non eos qui asperiores assumenda. Repellendus ut ipsa libero totam.','1984-02-13 05:21:15','1984-01-23 03:21:23'),
(49,37,9,'Facere consequuntur reiciendis velit reprehenderit dicta eum sed. Porro ipsam est omnis nulla ut. Vel ipsa fuga tempora et.','1992-09-05 13:45:27','1992-07-25 09:53:19'),
(50,78,2,'Tenetur omnis veritatis veritatis nobis quia autem. Aut ab commodi distinctio voluptatem distinctio occaecati sapiente.','1997-09-16 07:03:16','1991-08-05 03:04:51'),
(51,26,55,'Odit possimus dolores ut rerum quasi eos. Consequuntur dolor quos officiis nostrum ipsa omnis aut. Autem aspernatur veniam repudiandae.','2002-03-18 04:26:33','1989-05-16 04:14:19'),
(52,11,44,'Rerum ipsa quis enim eos. Aut qui iusto laudantium dolore. Qui a omnis error. Omnis maiores incidunt praesentium delectus fuga placeat.','2001-12-08 14:39:47','2010-07-08 22:43:31'),
(53,92,65,'Sed hic unde qui nihil quia et repudiandae. Numquam fuga ut nisi ullam illo. Dicta quis ratione quibusdam et.','2001-05-01 10:06:24','1979-03-29 05:01:38'),
(54,30,56,'Illo aut id omnis itaque et et. Aut distinctio quibusdam reiciendis qui. Vitae culpa laborum qui qui. Iste exercitationem recusandae nostrum voluptas doloribus. Reprehenderit similique minima expedita.','1990-03-09 11:42:00','1977-11-30 15:56:21'),
(55,90,98,'Deserunt ratione et aliquam qui. Consequuntur error vel unde fugit quo perferendis itaque. Qui est asperiores laboriosam.','1971-04-05 13:27:07','1996-07-18 09:14:31'),
(56,NULL,32,'Consequatur vero laudantium recusandae aut. Ut eligendi amet rem similique. Perspiciatis consectetur et eos veniam cupiditate sapiente voluptatem. Facilis eum cum temporibus aperiam ad porro qui.','1974-12-29 00:08:38','2019-11-22 12:34:39'),
(57,21,98,'Vitae impedit optio animi et. Sed sed odio voluptatem vitae. Dolores tempore provident rerum enim omnis maiores molestiae. Maxime animi culpa repellendus aut voluptatem.','1998-11-13 03:05:28','2003-05-13 09:11:07'),
(58,84,98,'Quia deserunt eum consequatur. Et dolor saepe laudantium odit. Eos eum unde recusandae non enim veritatis nihil. Animi voluptatum omnis id magnam.','2018-01-03 05:52:18','1972-04-25 14:36:32'),
(59,1,6,'Molestiae ut reprehenderit odio ab. Et et sit unde est est deleniti. Qui dolores sequi reprehenderit maxime tempore. Repellat commodi fugiat et dignissimos.','2012-04-04 19:33:38','1971-11-24 03:48:13'),
(60,19,17,'Enim aut illo repellendus exercitationem. Aut molestiae rerum et a voluptate enim similique laudantium. Esse accusamus commodi vel et expedita aut ex. Voluptas omnis autem in numquam et. Eum repudiandae rerum fuga cumque expedita iste eligendi.','1996-07-15 06:38:38','1991-09-29 04:33:40'),
(61,NULL,38,'Natus blanditiis illum odio ipsa labore expedita. Officia eos unde minima enim. Dolor velit similique numquam. Dignissimos dolorem expedita ducimus vero. Quos aperiam temporibus est illo vitae blanditiis.','1999-11-29 23:45:51','1979-05-16 13:51:01'),
(62,97,29,'Fugiat et illum neque saepe itaque. Sit ipsum aspernatur et laboriosam aut. Sint quod placeat quam aut asperiores.','1996-10-22 11:33:38','1993-12-04 10:42:14'),
(63,75,51,'Sint nisi aut rerum commodi ut nemo id. Facilis et corporis ea id quasi vel. Ea et soluta voluptate est et deserunt. Ut non nostrum consequatur sapiente aspernatur labore voluptas amet. Dolore aut ut veniam placeat non.','2017-09-05 07:52:37','2019-05-06 04:14:37'),
(64,5,6,'Ullam quo consequatur maiores ut placeat distinctio. Non minus tempora error debitis voluptatibus. Sit minus vel velit et distinctio. Tenetur cum aperiam consequatur unde.','1978-11-19 16:55:12','1980-01-15 10:46:05'),
(65,NULL,75,'Sed sequi alias amet aut. Eum pariatur est enim voluptas ipsum repellat voluptatem a. Eaque voluptas corporis sequi illo similique sequi. Tenetur eveniet dolor vel eveniet.','2009-08-15 01:53:52','1986-02-01 16:30:01'),
(66,67,7,'Assumenda nam accusamus dolores ut sint perspiciatis. Tempore voluptatem enim consectetur incidunt. Omnis est distinctio ipsum veritatis necessitatibus non cum.','1988-04-13 22:16:28','1974-09-24 09:56:31'),
(67,92,14,'Tempora et aut tempore rerum vitae voluptatibus. Sit fugit esse impedit aspernatur. Atque omnis aliquam vel amet quas delectus.','2007-08-15 12:12:35','2007-06-02 20:12:22'),
(68,92,46,'Ut doloremque ut saepe eos neque. Sint vero libero ducimus ut labore asperiores. Magnam quibusdam quo provident reiciendis aut enim. Corporis ut recusandae quo eos ut enim deserunt.','2009-07-15 13:16:06','1998-06-03 03:09:41'),
(69,NULL,5,'Sit nesciunt ea nam corporis. Necessitatibus a vero porro provident in. Provident sed cum nihil et quia sit culpa. Rerum dolor vero molestiae rem qui dolore quam.','1972-12-17 20:48:05','1985-02-21 02:36:53'),
(70,86,93,'Sit ut laudantium cumque. Praesentium excepturi placeat unde est. Vero numquam ullam animi exercitationem veritatis facilis. Autem voluptas repudiandae dolore ratione est ullam molestiae quidem. Cumque quia id at totam.','1990-04-25 06:42:28','1973-08-19 18:22:23'),
(71,24,19,'Eaque laboriosam nam quaerat iste minus quia nesciunt. Perspiciatis sint voluptatum nesciunt natus fugiat voluptatem excepturi nesciunt. Provident tempora voluptatem iure aliquam.','1994-09-10 20:21:00','1974-02-04 22:04:31'),
(72,82,23,'Excepturi unde ea consequatur repellat. Nobis ut non id et. Nulla dignissimos dolor ullam unde unde eveniet.','1996-12-11 20:12:20','1972-07-08 08:26:59'),
(73,31,34,'Beatae placeat dolor sapiente hic voluptatem necessitatibus iure sed. Et perferendis dolorem et ipsam totam. Sapiente non non aut consequatur sequi. Aliquid cumque veniam ab voluptas.','1976-12-13 03:55:20','1991-07-23 05:01:06'),
(74,NULL,69,'Molestiae est qui distinctio et pariatur occaecati quae. Impedit aut in nobis qui distinctio voluptate. Ex ea sit nobis voluptatem earum sit rem. Autem ut optio voluptas quisquam qui.','2007-01-02 11:55:12','2020-01-02 09:51:30'),
(75,75,52,'Eum unde aut voluptatem quo beatae. Tempore in sint in saepe omnis sit unde. Occaecati itaque iusto eligendi soluta vel. Quis est optio rerum.','2024-10-28 11:09:20','2006-11-23 17:18:40'),
(76,6,100,'Ut ut sit debitis et repudiandae. Earum veniam autem maiores quia fugiat optio. Dicta quaerat itaque dolores sint tempore magni omnis. Sequi quaerat est et.','2010-10-31 03:56:09','2015-10-09 00:14:34'),
(77,48,11,'Molestiae at sint expedita fugiat hic corrupti et. Esse accusamus quos et non eum. Rerum tenetur ad veniam labore a iste nobis accusantium.','2022-03-22 23:43:02','1987-01-26 16:25:55'),
(78,36,46,'Sit inventore aspernatur ullam reiciendis odit non. Porro inventore laudantium et voluptas animi est pariatur. Et eaque ab in et voluptatibus molestias aut. Atque et culpa eius deleniti placeat.','1989-08-23 04:25:06','2011-05-29 10:08:00'),
(79,82,85,'Optio consequatur ea sint soluta enim cumque iusto. Sunt ut illo molestiae ut qui aut. Similique voluptate deserunt rem vel fugit aut provident.','1987-02-25 20:09:50','2016-10-13 23:04:17'),
(80,22,46,'Odio ut voluptas asperiores et molestiae provident repudiandae. Animi aperiam et perspiciatis voluptatem. Molestiae impedit id consequatur eos.','1974-10-19 23:54:51','2010-10-02 06:46:28'),
(81,44,85,'Ut minus iste cumque facilis nam distinctio. Aut odio eaque aliquid suscipit. Aut nobis rerum est dolor inventore optio.','2006-08-22 15:22:17','2005-04-27 09:43:53'),
(82,57,32,'Totam sequi quibusdam ut ad. Ut placeat ut quis.','2001-11-27 02:06:34','1987-04-07 23:01:22'),
(83,31,16,'Dicta perferendis placeat esse quae consequatur sit provident. Officiis molestiae quisquam aperiam eligendi. Eos possimus adipisci et distinctio adipisci. Ipsa facilis accusamus et aut eius. Quia quisquam neque sit magnam sint quo similique.','2006-09-02 14:51:11','2002-09-27 17:12:43'),
(84,22,10,'Laborum aut omnis dolor reiciendis non. Nemo esse voluptatum quaerat molestias incidunt expedita rerum. Alias et aut natus vel.','1971-01-16 12:28:58','1991-07-17 12:02:38'),
(85,77,20,'Reiciendis aperiam nihil minima fuga reiciendis. Optio illum est aut non dolorum. Aut officia nobis incidunt nihil aspernatur consequatur. A qui fuga fuga inventore est.','1992-12-15 12:56:24','2001-01-18 12:44:52'),
(86,91,73,'Quia aperiam et sint est blanditiis aut. Accusantium sed quod libero enim tempore perferendis dolorem. Perferendis porro repellendus incidunt autem dolores dolorum laborum.','1975-02-11 04:32:54','1987-02-01 01:53:00'),
(87,95,24,'Et ipsa sed fuga qui eos exercitationem cum reprehenderit. Enim error dolores ipsum debitis corrupti nam et. Saepe illum odio ea aspernatur. Eos vel fugiat possimus eum sint voluptatum. Blanditiis ut molestias quis ut delectus rerum.','1992-09-17 21:45:39','1975-10-13 22:54:21'),
(88,100,15,'Quia minima ut autem corporis deserunt omnis perferendis aperiam. Eaque illo beatae qui aut recusandae. Hic nobis tenetur asperiores aut explicabo sed. Harum atque et voluptatibus maiores.','1971-04-12 10:32:23','1990-04-04 13:32:43'),
(89,19,67,'Velit quam deserunt quaerat et. Id est consequuntur animi aut impedit tempore. Consequatur qui sit placeat vero repudiandae esse quia. Voluptatibus eos ut accusantium repellat suscipit.','1982-07-21 20:34:32','1981-01-14 05:06:16'),
(90,100,27,'Sunt laudantium modi sapiente natus quo animi. Non unde assumenda sit nihil. Veritatis ipsa officia necessitatibus nobis.','1986-01-08 12:25:20','1983-05-29 17:00:24'),
(91,10,28,'Vel ea est sunt ratione. Sint corrupti fugit ad minus eligendi. Nisi et explicabo quia suscipit reiciendis et.','2018-05-21 08:53:11','1984-04-16 12:40:49'),
(92,61,71,'Et dolorem explicabo quisquam. Amet nesciunt laborum alias corrupti vel. Ipsa laboriosam debitis doloribus itaque. Saepe fugiat neque quia magni doloremque.','2022-08-15 05:49:46','1987-01-06 00:26:33'),
(93,40,45,'Sunt quod distinctio quisquam sit necessitatibus. Et eaque amet est non tempore in facilis. Ab et voluptatem voluptate et quaerat.','2019-12-15 02:27:23','2010-12-28 15:32:40'),
(94,93,16,'Quos repellendus aut dolor. Nihil recusandae exercitationem excepturi laboriosam voluptates. Eos in ab cupiditate facilis omnis molestiae exercitationem. Est quidem animi enim quo libero.','1987-02-24 21:00:05','2023-12-11 01:28:25'),
(95,29,93,'Qui labore laudantium culpa cumque officia iure. Perspiciatis dicta iste laborum id quas et est. Vel quis voluptatem dolores odio est et asperiores. Ea ut cum praesentium quod praesentium deserunt. Quos quis dolorem reprehenderit maxime nihil.','1978-07-25 09:31:53','1986-06-21 09:07:58'),
(96,6,8,'Dolores eveniet quae itaque. Porro optio quis dolorum sit nobis. Illum magni et hic soluta eos. Iste dolorum qui vel sint repellat.','2009-07-03 06:43:53','2019-12-30 02:49:25'),
(97,82,18,'Qui officia repellat dolores corporis non enim sequi. Ipsum cumque suscipit sunt. Debitis beatae accusantium et maiores. Expedita impedit et error officiis.','2019-06-21 12:56:01','1990-07-26 09:49:48'),
(98,94,76,'Ipsam voluptatem aspernatur ut laboriosam maxime sunt alias eius. Perferendis laboriosam incidunt sint praesentium. Accusamus non et accusamus nostrum nostrum consectetur quos. Impedit quibusdam sed vel nam illum qui libero.','2001-12-22 02:51:33','1996-09-18 20:22:42'),
(99,29,98,'Facilis est assumenda cupiditate repellat. Et at deserunt architecto sit. Ullam vitae exercitationem et corporis. Nostrum quod sed voluptate veniam eum earum.','2021-02-14 23:31:07','1990-10-23 05:18:12'),
(100,54,40,'Amet enim accusantium at accusantium. Et blanditiis excepturi itaque rem rerum est provident excepturi. Alias repudiandae consectetur rerum ipsa non blanditiis.','2004-05-04 03:44:46','2001-06-30 20:23:50');

-- CROSS
SELECT * FROM users 
JOIN messages;

-- users 100
-- messages 100
SELECT COUNT(*) FROM users 
JOIN messages;

-- выбрать пользователей, которые отправили сообщения. показать их сообщения
SELECT u.*, m.*  FROM users u 
JOIN messages m WHERE u.id=m.from_user_id 
ORDER BY u.id;

-- INNER 
SELECT u.*, m.*  FROM users u 
JOIN messages m ON u.id=m.from_user_id 
ORDER BY u.id;

-- выбрать ВСЕХ пользователей и показать отправленные сообщения, если они есть
-- LEFT
SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id 
ORDER BY u.id;

-- выбрарь пользователей, которые не отправляли сообщения
SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id 
WHERE m.from_user_id IS NULL
ORDER BY u.id;

-- выбрать ВСЕ сообщения и показать отправителей, есть они есть
-- RIGHT 
SELECT m.*, u.*   FROM users u
RIGHT JOIN messages m   ON u.id=m.from_user_id 
ORDER BY u.id;

-- аналог решения с LEFT, но через RIGHT 
SELECT  u.*, m.* FROM messages m
RIGHT JOIN users u   ON u.id=m.from_user_id 
ORDER BY u.id;

-- FULL JOIN
-- выбрать ВСЕХ пользователей и ВСЕ сообщения
/*SELECT u.*, m.* FROM users u
FULL JOIN messages m ON u.id=m.from_user_id
ORDER BY m.id;*/

SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id 
UNION 
SELECT u.*, m.*  FROM users u 
RIGHT JOIN messages m ON u.id=m.from_user_id;

-- одинаковые "схлопываются"
SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id 
UNION 
SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id;

-- показывать ВСЁ
SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id 
UNION ALL
SELECT u.*, m.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id;
 
-- 
DROP TABLE IF EXISTS `media_types`;
CREATE TABLE `media_types`(
	`id` SERIAL PRIMARY KEY,
    `name` VARCHAR(255),
    `created_at` DATETIME DEFAULT NOW(),
    `updated_at` DATETIME ON UPDATE NOW()
);

DROP TABLE IF EXISTS `media`;
CREATE TABLE `media`(
	`id` SERIAL PRIMARY KEY,
    `media_type_id` BIGINT UNSIGNED,
    -- `user_id` BIGINT UNSIGNED NOT NULL, -- создатель media
  	`body` text,
    `filename` VARCHAR(255),
    `created_at` DATETIME DEFAULT NOW(),
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP,
    -- FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (`media_type_id`) REFERENCES `media_types`(`id`) ON UPDATE CASCADE ON DELETE SET NULL
);

DROP TABLE IF EXISTS `profiles`;
CREATE TABLE `profiles` (
	`user_id` SERIAL PRIMARY KEY,
    `gender` CHAR(1),
    `birthday` DATE,
	`hometown` VARCHAR(100),
	`photo_id` BIGINT UNSIGNED,
    `created_at` DATETIME DEFAULT NOW(),
    FOREIGN KEY (`photo_id`) REFERENCES `media`(`id`) ON UPDATE CASCADE ON DELETE SET NULL 
); 

ALTER TABLE `profiles` ADD CONSTRAINT fk_user_id
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON UPDATE CASCADE ON DELETE CASCADE;

INSERT INTO `media_types` (`id`, `name`, `created_at`) 
VALUES 
(1,'Photo','2003-07-09 10:08:05'),
(2,'Music','2009-06-19 20:08:09'),
(3,'Video','1984-04-18 01:55:09'),
(4,'Post','2001-04-17 06:47:52');

INSERT INTO `media` (`id`, `media_type_id`, `body`, `filename`, `created_at`, `updated_at`)
VALUES
(1,1,'Aliquam exercitationem voluptate cupiditate dolor aliquid et. Vitae nobis excepturi velit hic numquam. Dolor cumque sed aliquid dolores aut est.','voluptate','1995-03-13 10:50:31','2010-06-23 21:28:51'),
(2,3,'Modi id nihil veniam sed amet dignissimos. Nemo consequatur aspernatur sit ea. Maiores accusamus aut explicabo cupiditate tenetur est.','necessitatibus','2010-04-14 22:49:50','1987-01-21 07:04:07'),
(3,3,'Quia corrupti et saepe aut sit expedita similique possimus. Voluptatem repudiandae vero fugit assumenda. Aut fugit magnam adipisci eligendi laborum.','fugit','2021-10-27 14:45:18','1974-03-26 01:17:22'),
(4,2,'Explicabo molestiae quibusdam rerum id voluptates neque quaerat. Animi earum assumenda est qui ab repellendus impedit. Quia ipsa ratione facere laudantium provident error architecto.','numquam','1970-02-19 13:58:54','1977-02-19 12:28:09'),
(5,1,'Error nemo quia voluptatibus assumenda quam. Blanditiis harum exercitationem hic. Voluptates et adipisci aperiam amet nulla. Ad a laborum dolorem ut qui.','aliquid','2018-03-07 04:18:53','2020-12-05 00:54:54'),
(6,2,'Temporibus qui unde eos. Quia ut assumenda odio sint praesentium. Architecto voluptatem id ea modi. Molestiae est sit quibusdam enim dolore accusamus.','quae','2017-04-19 19:06:17','1994-12-15 16:54:15'),
(7,1,'Eum modi aut quibusdam repellendus nemo dolores cum maiores. Et explicabo expedita in voluptatem. Modi harum modi cum magni quis enim. Animi aspernatur animi aspernatur et.','voluptas','1986-05-18 11:55:30','1988-03-01 01:04:25'),
(8,3,'Voluptatem ut similique odio qui a aspernatur. Expedita possimus fuga assumenda accusantium. Distinctio autem et nihil asperiores quidem omnis. Veniam voluptates aut eum inventore eligendi qui minima.','quia','1996-01-18 06:54:05','2025-10-09 14:07:30'),
(9,2,'Nulla nihil earum non perferendis aut. Iure distinctio nam quis nihil veniam dignissimos. Voluptatum commodi nemo culpa repellat sint quis ut magni.','aut','2006-03-16 05:23:08','1983-10-15 23:05:19'),
(10,2,'Ullam repellendus impedit asperiores vitae occaecati. Praesentium modi possimus maxime dolorem voluptatem enim dicta quisquam. Assumenda vitae architecto nesciunt quas. Vero qui dolorem dolor sequi perferendis aliquid maxime.','enim','1973-09-05 11:49:34','2013-03-05 18:56:33'),
(11,3,'Ut voluptatem voluptatibus provident quia dolores. Minima libero quam laboriosam laborum accusantium quos.','nesciunt','2012-02-25 07:58:10','1970-09-10 07:35:33'),
(12,4,'Corrupti blanditiis corporis expedita numquam occaecati. Labore tempore quidem nobis omnis. Vero rem inventore aut minus sit perferendis. Ducimus consequuntur non eligendi nam sed optio praesentium.','voluptatem','1993-06-15 22:15:21','1987-09-05 02:24:56'),
(13,4,'Dicta velit et omnis sit et in earum fugit. Et hic aliquam aut sit. Doloribus iusto doloribus minus rerum pariatur.','rerum','2019-01-17 15:54:31','1971-04-19 18:26:20'),
(14,1,'Nobis qui ratione rerum accusantium dolor blanditiis qui. Et et nemo asperiores et mollitia vero et. In libero in corrupti aliquid deleniti ut hic.','beatae','2005-03-18 17:38:55','2009-05-10 18:12:27'),
(15,3,'In voluptatem ad incidunt maiores recusandae quod. Vel impedit doloribus beatae est voluptates. Quam optio sit cumque rerum.','voluptas','1979-06-11 23:20:05','2007-01-10 07:23:38'),
(16,1,'Maiores porro enim eos quia explicabo hic. Eligendi voluptate et culpa. Excepturi soluta perspiciatis quo non eligendi quis. Quam eos nostrum consequatur vitae officia dignissimos.','quia','2001-10-29 16:15:28','2004-10-05 11:52:37'),
(17,3,'Vitae et omnis maxime a omnis ea. Vero ullam consequatur nam voluptatem velit ullam qui. Atque rerum quibusdam delectus eius assumenda ea cupiditate. Voluptatum numquam est qui dolorem necessitatibus. Totam ullam aut voluptatibus quo.','tempora','1982-01-03 02:46:46','2004-07-15 07:49:47'),
(18,2,'Officia recusandae odit eum. Minima maxime neque voluptas. Quisquam voluptatem sunt voluptatem quibusdam quod.','placeat','2013-11-06 08:25:26','2022-11-07 00:16:49'),
(19,3,'Natus voluptate quo porro nulla. Reiciendis ipsa minus magnam similique corporis voluptates. Rerum labore dolore laborum exercitationem.','deserunt','1979-09-07 03:41:00','1983-11-06 17:25:26'),
(20,2,'Molestias eos vitae dolorem totam iure eos. Odio enim error non aut. Voluptatibus cum reiciendis corporis autem eum suscipit.','sint','1974-11-08 21:40:51','2025-07-21 17:14:42'),
(21,3,'Enim ut sit odio itaque vel debitis esse. Animi sint officiis velit quibusdam explicabo ullam. Quo corrupti omnis dolores harum sunt. Architecto iusto est ab reprehenderit ea dolorem accusantium. Qui minus et dolorem.','libero','1986-01-01 04:06:48','2008-06-30 18:42:34'),
(22,3,'In voluptas mollitia fuga corporis. Facere temporibus magni in vitae doloribus voluptatibus. Veritatis aut rerum ipsa qui ipsum quisquam enim.','dolores','1987-10-06 05:02:05','2002-02-02 04:47:00'),
(23,2,'Similique voluptatem vero ullam quia in eum natus itaque. Ut incidunt suscipit doloremque assumenda laudantium architecto dicta. Voluptatibus assumenda magni tenetur quia repellat. Qui tempora labore odio.','dicta','2017-10-24 04:49:41','2006-06-09 11:07:06'),
(24,4,'In non quis sunt ducimus voluptatem. Vel qui veritatis inventore excepturi in cumque. Nulla necessitatibus qui molestiae eveniet voluptate numquam quibusdam. Dolor animi et expedita quod natus enim saepe.','harum','1971-03-20 14:06:56','2011-11-28 18:40:57'),
(25,2,'Sed quia quia et et. Nam tempora earum omnis id. Et adipisci magnam dolor.','mollitia','1997-12-14 00:12:41','1980-04-08 20:37:43'),
(26,1,'Provident impedit repudiandae molestiae tenetur in dolores omnis voluptatem. Itaque laudantium illo perferendis aut. Corrupti minus eum enim beatae.','saepe','1995-02-15 11:47:04','1981-10-20 09:12:10'),
(27,4,'Dolores porro sint deleniti quos ex magni. Dolorem dicta molestias eaque dolore. Rerum possimus nihil est aut mollitia eum occaecati.','sequi','1970-04-22 16:34:20','1987-10-14 12:53:09'),
(28,2,'Quasi aperiam delectus id non nostrum itaque dolores. Alias quod ut dolores saepe et doloremque. Porro et voluptas quia officiis magnam. Eligendi pariatur id praesentium quidem.','vel','2022-12-08 22:33:45','1995-07-07 21:26:28'),
(29,4,'Laudantium ut illum cupiditate omnis. Corporis vel inventore asperiores porro. Laborum repellendus ullam explicabo ut.','laudantium','1992-05-14 08:49:33','1972-12-03 17:24:02'),
(30,4,'Architecto dolorem aut perferendis et et qui voluptatum animi. Illum voluptatibus ipsam consequatur velit dolores. Sunt quidem earum esse animi natus facilis. Voluptate harum eius in culpa ipsa voluptates qui quo.','corporis','2003-10-14 14:50:30','1992-03-12 16:54:31'),
(31,1,'Modi dolorum libero dolor ea sit laborum. Nesciunt ut sequi ut et dignissimos qui architecto. Quidem qui et dolor. Et in quod fuga quo iusto optio.','repellendus','1970-07-13 15:57:59','2010-11-16 19:22:23'),
(32,3,'In aut voluptas est inventore consectetur. Alias eius voluptatibus officia ut. Qui mollitia corrupti rerum id molestias laudantium.','incidunt','2010-11-08 12:26:42','2001-11-24 15:32:48'),
(33,4,'Quaerat enim minus est dolor magni facere recusandae. Exercitationem hic minima vel voluptas fugit. Velit modi quia molestias quibusdam.','nam','2022-02-22 18:21:39','1977-02-17 08:58:12'),
(34,4,'Labore omnis eum ipsam. Id ut est laudantium dolorem quod. Quas quaerat omnis qui.','maiores','2020-12-26 03:55:35','1985-03-25 08:48:46'),
(35,1,'Magni quis alias qui itaque. Nesciunt molestiae et natus est dolores sapiente quis. Dolorem voluptate eius excepturi rerum officia doloremque.','non','1989-12-31 19:53:15','1991-06-14 16:42:20'),
(36,1,'Quaerat alias aut voluptates labore consequatur nobis fugiat. Porro voluptatibus minus vero omnis et autem cum libero.','aut','1999-09-19 15:04:31','1979-03-08 08:35:22'),
(37,3,'A perferendis nisi illo natus. Sapiente non libero sit explicabo aspernatur. Dignissimos error natus assumenda qui.','et','2016-11-10 00:54:32','1985-06-26 05:58:55'),
(38,2,'Illo et eligendi qui perferendis voluptatibus. Reiciendis id non quidem quidem modi quia fugiat dignissimos. Asperiores et vitae iure quas. Enim autem quia neque.','laborum','2009-12-10 01:21:00','1991-12-04 13:38:24'),
(39,1,'Occaecati doloremque labore commodi enim. Et voluptatem vero fugiat soluta tenetur. Tenetur qui est aut necessitatibus. Quos et ut ad sit sit vel.','praesentium','2001-04-14 02:46:04','1980-04-07 15:11:00'),
(40,2,'Eius hic voluptatem aut qui aut consequatur similique. Laudantium aut at ducimus et laudantium corporis quo. Omnis repellat sint qui provident sequi id.','aut','1975-07-29 12:55:35','1988-05-15 03:14:22'),
(41,3,'Quis natus tempore quos reprehenderit molestias est. Accusamus perspiciatis perferendis et aut. Iure quo neque enim est quo consequatur voluptate cum.','quae','1985-05-08 02:52:21','2013-08-02 17:55:54'),
(42,3,'Beatae eveniet odio quod unde quo. Architecto fugit est dolores enim sed. Deserunt voluptate quo impedit eum architecto dolorem assumenda qui. Nihil voluptate omnis aut corrupti consectetur est. Minus incidunt voluptatem nam aspernatur nulla nihil.','nemo','1989-04-08 21:17:13','2022-11-04 19:12:14'),
(43,1,'Eius consequatur qui quis accusamus doloremque. Rerum quia consectetur eos qui magni repellendus expedita praesentium. Iusto reprehenderit qui consequatur dolor id est voluptas. Enim sequi et soluta optio autem porro.','suscipit','1974-09-17 13:22:41','1982-06-11 13:19:42'),
(44,3,'Ut sunt rerum ducimus delectus rerum. Dignissimos veniam sequi voluptatem vel totam sequi architecto. Est excepturi qui voluptatem omnis ut.','sunt','2008-03-17 04:45:52','1989-03-21 12:26:18'),
(45,4,'Odit assumenda est unde facilis. Nesciunt enim dolores enim. Consequatur et reiciendis consequuntur ut nostrum. Atque laborum inventore culpa nostrum quibusdam.','voluptas','1997-09-27 14:17:02','1986-05-13 19:32:37'),
(46,2,'Reprehenderit accusamus qui ut. Consequatur iste libero quasi unde voluptatem quisquam et et. Nesciunt velit fuga enim.','doloremque','1976-10-31 08:13:22','2004-01-31 17:35:44'),
(47,1,'Minus aliquid quia labore consequatur reprehenderit. Iusto a nostrum velit. Soluta rerum nemo sint dolorem. Illo cupiditate rerum impedit.','unde','2021-09-19 06:29:38','2025-01-07 23:55:35'),
(48,2,'Et tenetur minima ipsam quod explicabo recusandae eius. Et aspernatur voluptatem odit. Amet facere dignissimos totam est porro qui laudantium et. Tenetur ullam laborum et reiciendis vel sed.','similique','1985-10-19 00:27:10','2016-10-15 11:18:37'),
(49,1,'Maxime ipsam necessitatibus quia eos hic non beatae. Dolorem veritatis nulla molestiae accusamus voluptatibus eos. Et placeat consequatur molestiae quibusdam quia nam sed. At deserunt dolorem minus possimus rerum ratione.','et','1997-08-16 12:21:12','1992-02-10 17:26:05'),
(50,2,'Et similique dolorem laudantium et ratione. Velit consequatur qui ab perspiciatis unde ut. Aut provident et eius quam. Laborum veniam consequatur sit aliquid enim magnam quam.','libero','1973-10-05 20:36:49','1978-01-30 22:01:05'),
(51,2,'Ex molestiae laudantium et quas aspernatur fugiat. Quo at sunt saepe nulla pariatur. Occaecati animi natus nulla.','enim','1984-12-27 04:45:55','1981-03-01 22:31:02'),
(52,2,'Et rerum molestiae in laborum sint est. Est est corrupti omnis repudiandae. Adipisci ut maxime excepturi facere similique. Aspernatur provident accusamus officiis consequatur commodi libero.','ipsam','2019-10-28 01:55:42','1983-02-15 23:55:32'),
(53,2,'Quia non nobis laboriosam inventore voluptas est hic adipisci. Quidem enim quia nulla et non ut impedit. Reprehenderit beatae illum consequuntur cupiditate nostrum. Dolorum est rerum doloribus cupiditate velit pariatur. Voluptas sunt quam non quibusdam repellendus quo et eligendi.','mollitia','1998-10-29 11:11:53','1988-07-09 09:07:08'),
(54,2,'Placeat veritatis optio libero ex perferendis consequatur. Nesciunt perferendis eaque et sit odio. Sed qui exercitationem dolorum alias ut. Aliquam accusamus doloribus quidem dignissimos dicta.','maiores','1999-12-22 02:57:49','1984-07-20 05:46:18'),
(55,4,'Earum ipsum magni distinctio quis occaecati facere. Dicta explicabo unde non voluptatem. Quam reiciendis quibusdam aperiam deserunt.','similique','1984-07-13 07:10:08','1974-08-30 00:15:11'),
(56,3,'Sequi ipsa dolores quisquam autem. Debitis aut a consequatur dolores. Ut est tenetur eveniet. Et ducimus asperiores velit consequuntur facilis hic.','adipisci','2020-08-30 03:21:44','1987-01-19 20:04:51'),
(57,4,'Illum similique modi ullam et quis sapiente debitis. Dolores quia dolorem officiis dolorem occaecati doloremque sit quia. Deserunt rerum omnis aut eveniet culpa. Fugit nam libero neque.','corrupti','2019-12-16 05:09:05','2024-09-02 07:59:00'),
(58,3,'Suscipit in sed blanditiis cumque sit dolorem. Maxime animi aliquam asperiores. Et sed similique atque aliquid ad numquam. Enim quo quam aut ducimus qui et est.','rem','2017-08-31 20:26:22','2019-10-31 04:24:23'),
(59,1,'Adipisci voluptatem quod est odit. Repudiandae ab necessitatibus dolor at quis. Cum dolores sunt perspiciatis nihil illo laboriosam et.','repudiandae','1985-11-13 14:00:47','1989-10-19 18:57:19'),
(60,3,'Ea ut modi laborum. Omnis voluptate est modi dolorem consectetur est in. Ut corporis expedita eum est exercitationem ea in. Autem enim culpa delectus autem quis molestiae. Qui placeat est qui.','est','1970-12-17 09:54:53','2005-01-09 19:03:09'),
(61,3,'Id necessitatibus temporibus quis dolores. Illum omnis non debitis odio dicta accusamus aut. Expedita commodi molestiae aut adipisci asperiores velit cupiditate labore. Consequatur molestias debitis expedita nisi quam quas.','facilis','1996-08-29 22:12:20','2013-03-02 13:58:30'),
(62,2,'Cupiditate maiores esse est numquam. Inventore possimus consequatur porro consequatur et aut culpa. Delectus voluptatem rerum quis esse praesentium dolor earum explicabo. Dolorem unde aut ex tenetur.','eos','2019-05-22 02:04:30','1992-10-04 00:50:00'),
(63,2,'Non harum ad aut necessitatibus aut quod. Quasi consequatur eum et vel. Quo praesentium numquam quia autem non consectetur doloremque. Tenetur minima ex architecto eveniet similique nihil. Architecto expedita consequuntur vel adipisci.','iusto','1971-03-19 01:11:15','2015-07-25 06:05:27'),
(64,1,'Et sunt repudiandae ex officia tempora quia. Sit fugiat est ad fugiat unde et explicabo at. Sed in consequatur animi dolorum sed rerum nesciunt.','qui','2004-03-12 20:07:37','1977-01-29 13:11:11'),
(65,1,'Odit qui deleniti rerum eaque blanditiis aperiam qui. Quo vero et porro iusto. Quo neque quo quo et.','aut','1975-12-21 09:49:10','1992-09-20 14:34:43'),
(66,1,'Explicabo aut dolores voluptatem cum commodi. Error fugiat cumque id in perferendis voluptas aut. Velit similique qui neque non ab voluptas quia eos.','aut','2018-10-02 01:48:07','2002-04-28 05:55:04'),
(67,4,'Occaecati sint expedita animi in. Quia ipsum consequuntur sunt consequatur.','culpa','2008-04-15 22:24:53','1987-02-08 09:36:36'),
(68,2,'Ipsa odio laudantium neque quibusdam impedit beatae magnam. Officia excepturi voluptas nulla consequuntur autem. Provident nam repellendus distinctio qui voluptatem veniam est. Aut explicabo doloremque nobis ut iusto nulla ducimus.','expedita','1994-09-09 05:50:21','1974-04-03 04:43:43'),
(69,3,'Sint qui unde tenetur in consequatur blanditiis. Voluptas voluptas in incidunt non mollitia quo. Excepturi debitis sunt consequatur doloribus est tenetur. Recusandae distinctio optio sit sed accusamus ut.','deserunt','2002-08-16 08:21:44','2022-03-09 01:05:21'),
(70,3,'Et et odit repudiandae natus ad eveniet. Explicabo aliquam iusto excepturi eos. Vero tenetur neque minus. Est quae dolores reprehenderit.','architecto','2009-03-06 18:59:49','2019-01-19 05:08:04'),
(71,2,'Voluptatum aut nisi eveniet eveniet suscipit. Quasi maiores natus aut optio et. Est error aliquid qui iste veritatis earum aspernatur.','officia','1996-08-01 14:40:56','1998-06-01 15:05:08'),
(72,3,'Et accusantium sed veniam ipsum aut. Cumque at sapiente vero enim quasi et. Sed corrupti sed qui. Optio sequi qui asperiores sequi.','tempore','1973-05-11 08:17:05','2013-11-18 21:19:34'),
(73,3,'Qui debitis velit rem. Omnis debitis ratione consectetur dolorum recusandae quo sed. Nobis dolorem ut voluptas molestias rerum voluptatem animi itaque.','at','1999-11-11 09:48:35','2025-10-23 08:42:53'),
(74,1,'Dolorem debitis qui minus magni dolor neque dolores. Saepe facere et in ullam temporibus dolorem. Adipisci sed et porro.','consequatur','2005-09-10 13:33:02','2004-01-01 23:46:50'),
(75,4,'Alias et aut molestias nihil ipsum. Esse nesciunt adipisci numquam vel quos. Ea soluta accusantium error debitis in voluptatem natus. Sed non quis architecto velit nihil et.','sed','1982-06-18 00:19:27','1994-11-12 00:56:27'),
(76,3,'Vel corrupti qui quaerat occaecati optio. Porro sint sapiente molestias maiores. Quam dicta et odit. Rem aut sequi occaecati dolorem corrupti.','tempora','1999-05-19 11:25:35','1991-10-23 14:20:37'),
(77,2,'Ut eaque sit necessitatibus. Sit sint beatae ad ex facilis doloremque exercitationem. Assumenda molestias dolore aperiam. Molestiae ratione qui at vel incidunt.','voluptas','2019-04-26 23:32:43','2018-05-18 03:43:49'),
(78,2,'Dolor inventore nisi necessitatibus eius explicabo voluptates. Vero corporis sunt non impedit beatae quos et. Voluptas quis unde et non officiis.','blanditiis','1972-11-27 09:26:46','1989-03-24 01:16:12'),
(79,1,'Enim inventore minima non recusandae. Ad sed eum voluptatibus. Consequatur debitis dolorem tempore odio necessitatibus earum.','a','2009-05-15 09:46:21','1975-09-29 09:50:39'),
(80,1,'Nam autem id vel et dignissimos. Numquam incidunt error possimus a nemo cum veniam libero. Et porro aut ab aut.','aut','2008-08-29 05:00:41','1981-12-13 03:25:42'),
(81,1,'Eius nobis saepe quos et occaecati dolores. Praesentium odio unde beatae voluptatibus ea vel dolorem nesciunt. Natus aut molestiae fugit atque optio. Modi est vel consequatur et.','a','2019-12-14 04:00:59','2017-11-02 13:46:10'),
(82,4,'Corporis excepturi tempora commodi eum praesentium dolore quo debitis. Et sit voluptates ratione in. Nihil minus non veniam officia voluptatem. Ea laborum quam corrupti placeat voluptatum nihil.','perferendis','1987-03-28 12:55:07','2011-08-12 03:47:54'),
(83,4,'Minus tempora accusamus a amet vero. Culpa aliquam ut vero dolor dolores rem beatae. Possimus aut et dicta et sint consequatur atque. Iure maxime consectetur dolorum et voluptate. Accusantium iure similique veniam velit aut aut.','ut','1980-02-25 12:27:04','2023-04-28 17:47:59'),
(84,4,'Modi unde sunt et. Dicta nihil consequatur officia nemo. Eius quidem atque hic error et quos. Quibusdam dolor animi enim.','quas','1997-12-02 20:39:39','1996-01-21 06:30:36'),
(85,3,'Voluptate commodi sapiente quas quibusdam dolorum ducimus. Nulla laboriosam necessitatibus odit enim illum. Necessitatibus doloremque quaerat occaecati quos nisi dicta ducimus. Quos non corrupti ut repellendus.','eaque','1979-02-04 15:09:10','2010-09-24 22:39:19'),
(86,4,'Voluptas quo incidunt non quam tenetur ut. At ullam aut aut unde ut nobis. Animi voluptatibus nam et quam aliquam.','voluptate','2007-01-13 13:41:44','1977-10-14 17:02:59'),
(87,4,'Ea placeat suscipit eligendi est soluta perspiciatis. Magnam aliquid assumenda porro in debitis minus ut. Praesentium aspernatur id quos modi iste sit.','et','1974-11-06 12:41:24','1989-06-08 11:39:47'),
(88,2,'Mollitia eaque aut et quo. Cupiditate necessitatibus consequatur ipsam ducimus dolore. Rem non est soluta quae laboriosam. Earum quis at aut ea reprehenderit libero. Quo explicabo magnam consequatur voluptatibus ad ab sint ex.','quidem','2017-09-16 22:27:53','1998-10-02 07:33:22'),
(89,4,'Dicta accusantium eligendi quisquam molestias veritatis nesciunt non. Ea necessitatibus ut sunt vel voluptatem et. Illo laborum sunt vitae. Iure omnis eum enim.','illo','2023-07-19 18:56:19','1991-07-16 15:11:10'),
(90,1,'Et harum aut id totam veritatis vel eligendi sapiente. Harum sint reiciendis mollitia dignissimos vel. Voluptatibus ea est voluptas explicabo quaerat impedit.','aut','2015-10-23 03:42:31','2021-05-16 21:29:37'),
(91,4,'Sit illum ut iure corporis aspernatur excepturi fuga. Eveniet ut illum facere.','exercitationem','1977-02-21 19:06:39','1979-03-09 02:25:45'),
(92,1,'Sed laboriosam et placeat numquam natus. Veniam voluptas ex at. Nesciunt cum eos earum a. Qui voluptatibus ut et dolorem exercitationem.','vel','1985-01-31 17:32:57','2006-06-08 22:57:46'),
(93,1,'Voluptatem reprehenderit architecto nobis autem inventore voluptates. Dolor voluptates perspiciatis vel ut minima. Porro quia nihil consequuntur qui et magni. Optio corporis assumenda voluptatem laudantium.','id','2006-05-04 05:56:33','2024-03-06 11:30:14'),
(94,1,'Iste impedit laboriosam omnis et. Voluptatum occaecati quam qui mollitia aliquam a. Esse veritatis modi voluptatum molestias voluptates voluptatem vel dolorem.','excepturi','2022-02-04 22:35:23','2021-10-12 22:08:36'),
(95,4,'Nesciunt nihil est sunt nisi est officiis. Quia consequatur possimus quod aliquam. Et est est veniam nihil. Corrupti ratione autem quasi odit. Quidem et qui dignissimos voluptatum illum iste.','aut','2000-09-01 10:37:29','1997-10-04 01:19:36'),
(96,2,'Quia ipsum ea facilis vitae maxime. Praesentium in eum amet temporibus sint id ipsa. Ut et et aperiam omnis aut.','nihil','2011-08-30 10:44:29','1996-12-12 10:01:45'),
(97,4,'Ipsam maxime incidunt ut. Molestiae et exercitationem esse soluta eos autem. Soluta qui quidem placeat sequi magni.','voluptatem','1971-08-07 06:54:47','1990-02-02 01:53:45'),
(98,2,'Fugit quam non et est. Provident asperiores dolore rerum in reiciendis et. Omnis labore iure eum incidunt debitis et aut.','quam','1976-08-18 12:06:31','1982-08-08 02:04:13'),
(99,2,'Sapiente excepturi dolor unde itaque blanditiis quaerat vero. Quo nihil aut sit vitae rerum omnis. Quisquam ullam soluta unde itaque et. Omnis dolorum amet cum autem rem.','provident','1985-07-29 05:41:50','2021-03-07 16:42:26'),
(100,2,'Sed expedita mollitia voluptatem quo porro possimus dignissimos. Eveniet qui voluptatem vero consectetur at.','ipsam','2000-03-22 04:11:48','2003-04-19 01:54:46');


INSERT INTO `profiles` (`user_id`, `gender`, `birthday`, `hometown`, `photo_id`, `created_at`)
VALUES
(1,'f','2005-06-20','New Cory',5,'1984-07-31 09:31:18'),
(2,'f','1976-08-13','Norbertofurt',62,'2011-09-16 19:38:00'),
(3,'m','2022-05-14','Berneiceberg',32,'1993-05-30 23:14:28'),
(4,'m','2019-07-08','South Eltatown',41,'1985-08-13 07:01:40'),
(5,'f','2002-09-22','North Kanestad',44,'1978-12-09 19:28:47'),
(6,'m','2025-01-22','Port Jacynthe',83,'1983-06-25 06:38:18'),
(7,'m','1981-01-15','Manteview',35,'2021-05-29 06:41:08'),
(8,'f','1999-04-23','Deanview',92,'2003-09-26 21:57:16'),
(9,'f','1995-06-05','Anikaborough',14,'1978-03-01 06:54:51'),
(10,'m','2010-09-29','West Arnoton',NULL,'1997-07-25 10:12:29'),
(11,'f','1992-09-12','North Helene',39,'1999-08-15 22:19:53'),
(12,'f','1998-12-29','Manteton',69,'2002-01-08 16:45:44'),
(13,'f','1987-10-21','Morissetteside',33,'2002-12-23 00:21:44'),
(14,'m','1984-02-20','West Mark',18,'1993-01-29 13:48:53'),
(15,'f','1992-02-27','Jakubowskiport',91,'2013-07-25 03:24:52'),
(16,'m','2000-04-23','East Valentin',6,'1984-03-18 16:11:51'),
(17,'f','1986-04-05','East Veronaview',75,'1986-07-27 05:24:08'),
(18,'m','1983-11-19','Port Julienstad',95,'1998-08-21 10:41:53'),
(19,'m','2018-03-24','Schmidtland',8,'1995-06-17 17:44:29'),
(20,'f','2015-08-15','East Caterina',74,'1971-09-05 16:32:51'),
(21,'f','2008-03-09','North Shayna',68,'2015-06-05 18:15:04'),
(22,'m','1971-03-20','West Davonteburgh',53,'1974-06-13 08:10:12'),
(23,'f','2024-11-13','Considinehaven',88,'1979-08-16 12:48:54'),
(24,'f','2001-03-28','Rudyfurt',83,'2011-12-10 09:20:15'),
(25,'m','2024-04-06','Port Floy',7,'1986-12-18 07:13:15'),
(26,'f','1970-10-18','Evanport',95,'1971-12-30 00:35:21'),
(27,'f','2013-10-05','Gretastad',80,'2008-12-09 14:55:10'),
(28,'f','2014-01-31','Stantonstad',17,'2000-04-03 21:13:37'),
(29,'f','1997-06-21','North Nikki',83,'2000-11-19 07:26:50'),
(30,'f','2022-02-03','Carleeport',51,'2004-09-06 07:39:16'),
(31,'m','1989-01-18','Kamronport',68,'1987-09-05 11:09:27'),
(32,'f','1998-01-23','Aniyaton',46,'1974-09-05 23:31:50'),
(33,'f','1974-06-18','Port Meggie',94,'2017-03-16 21:37:28'),
(34,'m','2002-05-04','East Mara',2,'2013-05-22 06:33:27'),
(35,'m','1971-12-28','Hilpertport',27,'2015-05-17 17:31:16'),
(36,'f','2020-06-22','Abelside',3,'2007-09-06 03:28:25'),
(37,'f','2007-01-18','Anniestad',74,'1983-11-23 22:48:48'),
(38,'m','2011-08-25','Halvorsonmouth',16,'1978-06-17 16:50:19'),
(39,'m','2010-09-06','West Delphiatown',25,'2017-10-27 09:01:36'),
(40,'f','2009-04-07','East Reese',13,'1982-10-31 21:58:35'),
(41,'f','1997-01-02','Toystad',50,'2002-08-04 06:00:27'),
(42,'m','2017-03-11','East Jalynshire',14,'2006-01-14 18:18:03'),
(43,'f','1986-02-22','North Grahamshire',28,'2017-02-27 21:09:34'),
(44,'f','2006-06-22','Amparohaven',56,'1986-01-16 18:02:03'),
(45,'m','2016-11-10','Aidenland',52,'1983-12-03 04:11:48'),
(46,'m','1999-03-21','New Leonard',76,'2012-12-21 06:58:59'),
(47,'m','1996-06-02','South Kyleigh',95,'2003-06-01 18:15:20'),
(48,'f','1999-08-13','Ladariusport',57,'1991-07-06 20:39:58'),
(49,'f','2019-06-27','Jastton',65,'2016-06-05 16:53:18'),
(50,'m','2014-03-28','Rodgerstad',NULL,'2024-05-08 17:04:10'),
(51,'f','2021-11-26','Port Arianeburgh',94,'1973-10-17 03:48:09'),
(52,'m','2007-08-10','West Esmeralda',42,'1992-02-03 23:11:19'),
(53,'f','1970-07-16','Lake Brycen',87,'2008-12-16 12:26:37'),
(54,'f','1976-04-02','Gerlachstad',86,'2001-05-17 14:00:55'),
(55,'m','1986-07-26','Hilariomouth',84,'2019-03-15 06:55:25'),
(56,'f','2010-11-20','Lake Abetown',19,'1970-04-16 12:30:11'),
(57,'f','1982-01-01','North Cortneymouth',39,'2025-10-30 23:54:22'),
(58,'m','1988-06-07','Haleighborough',3,'2008-04-01 22:57:11'),
(59,'f','2002-02-14','Antoniettaton',18,'1989-05-25 16:16:55'),
(60,'m','1986-03-18','Krajcikberg',34,'1973-11-17 09:09:08'),
(61,'m','1998-09-24','Port Zachary',83,'2014-06-18 05:49:02'),
(62,'m','1986-10-02','Port Genovevaville',65,'1999-11-10 00:53:00'),
(63,'f','1986-02-26','Jacobsfort',7,'2023-06-28 22:53:49'),
(64,'f','1983-06-13','North Estellachester',49,'1999-04-13 05:08:36'),
(65,'f','1974-01-27','Port Carmella',25,'2012-07-31 04:20:38'),
(66,'f','1997-07-01','South Ellen',10,'1999-12-16 08:32:23'),
(67,'m','2018-09-02','West Caliview',86,'1980-10-14 04:20:58'),
(68,'f','1990-04-26','Lake Albinaberg',94,'2013-01-09 14:51:27'),
(69,'f','1981-10-21','Rennertown',54,'1973-07-21 13:46:10'),
(70,'m','2013-05-10','East Taureanbury',66,'2002-01-03 17:14:00'),
(71,'m','2024-09-02','West Orinland',9,'1976-08-28 11:13:36'),
(72,'m','1999-02-19','East Ryann',75,'1971-10-12 14:41:24'),
(73,'f','2020-02-18','Lake Adrain',90,'1970-10-25 05:48:29'),
(74,'m','1992-03-02','Port Eldridgeburgh',6,'2022-08-20 07:41:59'),
(75,'m','1977-09-12','North Rosachester',92,'1993-10-07 23:33:55'),
(76,'f','2024-10-20','Johnsmouth',82,'1991-11-21 16:19:54'),
(77,'f','1995-12-01','Elinorechester',47,'1993-03-16 18:51:34'),
(78,'f','2010-09-23','Willisburgh',28,'2024-01-12 07:42:41'),
(79,'f','1972-04-10','Swiftfort',52,'1981-03-16 20:44:09'),
(80,'f','1996-10-25','Herzogton',44,'2016-08-12 22:09:38'),
(81,'f','2019-10-02','South Lurline',6,'1974-02-14 16:32:35'),
(82,'m','1972-10-15','Danykamouth',69,'2005-01-22 20:12:11'),
(83,'f','2016-05-25','New Ethelyn',81,'2023-03-30 00:32:17'),
(84,'m','1986-07-09','West Hendersonberg',68,'1987-08-02 15:27:14'),
(85,'f','2006-05-16','Alanisshire',92,'1978-05-24 05:00:46'),
(86,'m','1986-06-02','North Ismael',28,'2004-05-30 12:03:32'),
(87,'f','1978-03-06','Sawaynfurt',96,'2018-02-15 16:37:39'),
(88,'f','1972-10-22','East Euna',38,'2006-11-16 07:14:26'),
(89,'m','1981-08-30','North D\'angelo',27,'1976-12-09 01:47:28'),
(90,'f','2014-07-29','Johnstonshire',NULL,'1971-09-25 08:07:15'),
(91,'m','2012-06-10','Lyricstad',31,'1992-04-12 20:44:28'),
(92,'f','1983-12-10','North Louvenia',90,'1977-01-05 01:28:35'),
(93,'f','1979-06-21','South Marioport',85,'2004-11-06 02:10:40'),
(94,'f','2015-04-06','Hermannmouth',96,'2004-06-10 23:30:24'),
(95,'f','1989-04-15','North Anais',30,'1971-01-15 17:16:28'),
(96,'m','2016-03-09','Fidelberg',92,'2018-08-28 06:29:45'),
(97,'f','2014-03-30','Johanberg',NULL,'2007-11-04 05:33:45'),
(98,'f','1982-02-22','Laishachester',89,'2010-06-04 17:48:03'),
(99,'f','1989-07-13','East Susanahaven',18,'1985-02-17 21:22:35'),
(100,'f','1991-06-03','South Oralberg',NULL,'2023-01-23 12:29:59');

-- Агрегатные функции 
-- COUNT, SUM, MIN, MAX, AVG

SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM `profiles`;
SELECT COUNT(photo_id) FROM `profiles`;
SELECT SUM(id) FROM users; -- как пример
SELECT MIN(id), MAX(id) FROM users;
SELECT AVG(id) FROM users;

-- GROUP BY
-- сгрупировать все media по типам
SELECT media_type_id 
FROM media
GROUP BY media_type_id;

-- посчитать кол-во mediа КАЖДОГО типа
SELECT media_type_id, COUNT(*) 
FROM media
GROUP BY media_type_id;

-- HAVING
-- вывести тип media, у которых кол-во больше 25
SELECT media_type_id, COUNT(*) AS cnt 
FROM media
GROUP BY media_type_id
HAVING cnt>25;

SELECT media_type_id
FROM media
GROUP BY media_type_id
HAVING COUNT(*)>25;


-- вложенные запросы 
-- выбрать всех пользователей, указав их id, имя и фамилию, город и аватарку
-- используя вложенные запросы
SELECT 
	id,
	CONCAT(firstname, ' ', lastname) AS 'Пользователь', 
	(SELECT hometown FROM profiles WHERE user_id = users.id) AS 'Город',
	(SELECT filename FROM media WHERE id = 
	    (SELECT photo_id FROM profiles WHERE user_id = users.id)) AS 'Аватарка'
FROM users;

-- используя JOIN
SELECT 
	u.id,
	u.firstname,
	u.lastname,
	p.hometown AS city,
	m.filename AS `main_photo`
FROM users u
JOIN profiles p ON u.id=p.user_id 
LEFT JOIN media m ON m.id=p.photo_id 
ORDER BY u.id;

-- выбрать пользователей, которые отправили сообщения
-- используя JOIN
SELECT DISTINCT u.*  
FROM users u 
JOIN messages m ON u.id=m.from_user_id 
ORDER BY u.id;

-- используя вложенные запросы
SELECT * FROM users 
WHERE id IN (SELECT from_user_id FROM messages) 

-- выбрарь пользователей, которые не отправляли сообщения
-- используя JOIN
SELECT u.*  FROM users u 
LEFT JOIN messages m ON u.id=m.from_user_id 
WHERE m.from_user_id IS NULL
ORDER BY u.id;

-- используя вложенные запросы
SELECT * FROM users 
WHERE id NOT IN 
(SELECT DISTINCT from_user_id FROM messages WHERE from_user_id IS NOT NULL ) 

-- посчитать кол-во mediа КАЖДОГО типа (БОЛЕЕ ТОЧНОЕ РЕШЕНИЕ)
-- SELECT media_type_id, COUNT(*) 
-- FROM media
-- GROUP BY media_type_id

-- вложенный запрос
SELECT id,
	name,
	(SELECT COUNT(*) FROM media WHERE media_type_id = media_types.id) AS cnt
FROM media_types;

-- JOIN
SELECT mt.id,
	mt.name,
	COUNT(m.media_type_id)AS cnt
FROM media_types mt
LEFT JOIN media m ON m.media_type_id = mt.id
GROUP BY mt.id;

-- оператор CASE 
-- вывести id пользователей и указать их пол (мужской, женский)
-- также вывести их возраст 
SELECT user_id, 
    CASE (gender)
         WHEN 'm' THEN 'мужской'
         WHEN 'f' THEN 'женский'
         ELSE 'нет данных'
    END AS gender, 
    TIMESTAMPDIFF(YEAR, birthday, NOW()) AS age -- функция определяет разницу между датами в выбранных единицах (YEAR)
  FROM profiles;

-- разбить пользователей на группы (дети, молодежь, взрослые)
SELECT 
	user_id, 
    CASE 
         WHEN age <18 THEN 'дети'
         WHEN age BETWEEN 18 AND 35 THEN 'молодёжь'
         WHEN age > 35 THEN 'взрослые'
         ELSE 'не определено'
    END AS group_users
FROM 
(SELECT 
	user_id, 
    TIMESTAMPDIFF(YEAR, birthday, NOW()) AS age -- функция определяет разницу между датами в выбранных единицах (YEAR)
FROM profiles) AS list;

-- посчитать кол-во пользователей в каждой группе
SELECT 
    CASE 
         WHEN age <18 THEN 'дети'
         WHEN age BETWEEN 18 AND 35 THEN 'молодёжь'
         WHEN age > 35 THEN 'взрослые'
         ELSE 'не определено'
    END AS group_users,
    COUNT(*) AS cnt
FROM 
(SELECT 
	user_id, 
    TIMESTAMPDIFF(YEAR, birthday, NOW()) AS age -- функция определяет разницу между датами в выбранных единицах (YEAR)
FROM profiles) AS list
GROUP BY group_users; 

-- функция if
-- вывести id пользователей и указать их пол (мужской, женский)
SELECT user_id, 
    if(gender='m',  'мужской',
    	if(gender='f', 'женский', 'нет данных')
    	) AS gender 
FROM profiles;

-- IFNULL
-- получить id пользователя и его автарку
SELECT
	p.user_id,
	IFNULL(m.filename, 'нет аватарки')
FROM profiles p
LEFT JOIN media m ON m.id=p.photo_id
ORDER BY p.user_id;

-- COALESCE
-- получить id пользователя и его автарку
SELECT
	p.user_id,
	COALESCE(m.filename, NULL, if(p.gender='m', 'Мужчина не имеет аватарки', NULL),
	if(p.gender='f', 'Женщина не имеет аватарки', NULL),
	'нет аватарки') AS 'Аватарка' 
FROM profiles p
LEFT JOIN media m ON m.id=p.photo_id;


