DROP TABLE
  IF EXISTS `radcheck`;

/*!40101 SET @saved_cs_client     = @@character_set_client */;

/*!40101 SET character_set_client = utf8 */;

CREATE TABLE
  `radcheck` (
    `id` int (11) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(64) NOT NULL DEFAULT '',
    `attribute` varchar(64) NOT NULL DEFAULT '',
    `op` char(2) NOT NULL DEFAULT '==',
    `value` varchar(253) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    KEY `username` (`username` (32))
  ) ENGINE = InnoDB AUTO_INCREMENT = 1543 DEFAULT CHARSET = latin1;

/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `radcheck`
--
/*!40000 ALTER TABLE `radcheck` DISABLE KEYS */;

INSERT INTO
  `radcheck`
VALUES
  (
    2,
    'masoudtest',
    'Cleartext-Password',
    ':=',
    'masoudtest'
  );

/*!40000 ALTER TABLE `radcheck` ENABLE KEYS */;
