DROP TABLE IF EXISTS `discountfactor`;

/*!40101 SET @saved_cs_client     = @@character_set_client */;

/*!40101 SET character_set_client = utf8 */;

CREATE TABLE
  `discountfactor` (
    `id` int (11) NOT NULL AUTO_INCREMENT,
    `type` varchar(64) NOT NULL,
    `value` varchar(20) NOT NULL,
    `factor` float NOT NULL,
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB AUTO_INCREMENT = 70 DEFAULT CHARSET = latin1;

INSERT INTO
  `discountfactor`
VALUES
  (8, 'time', '23:00-23:59', 0.4),
  (9, 'day_of_week', '3', 0.6),
  (10, 'day_of_week', '4', 0.4),
  (11, 'time', '01:00-03:00', 0.2),
  (12, 'time', '00:00-01:00', 0.4),
  (13, 'time', '03:00-07:00', 0.02),
  (40, 'default', 'default', 1),
  (54, 'date', '2019-05-27', 0.4),
  (59, 'date', '2019-08-20', 0.4),
  (60, 'date', '2019-09-09', 0.4),
  (61, 'date', '2019-09-10', 0.4),
  (62, 'date', '2019-10-19', 0.4),
  (63, 'date', '2019-10-27', 0.4),
  (64, 'date', '2019-10-29', 0.4),
  (65, 'date', '2019-11-06', 0.4),
  (66, 'date', '2020-01-29', 0.4),
  (67, 'date', '2020-02-11', 0.4),
  (68, 'date', '2020-03-08', 0.4),
  (69, 'date', '2020-03-19', 0.4);
