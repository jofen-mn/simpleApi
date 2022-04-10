use simpledb

CREATE TABLE `user` (
    `id` bigint(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(12) NOT NULL,
    `number` varchar(18) NOT NULL,
    `phone` varchar(11) NOT NULL,
    `gender` smallint(2) NOT NULL,
    `age` smallint(3) NOT NULL,
    `stature` smallint(3) NOT NULL,
    `weight` smallint(3) NOT NULL,
    `address` varchar(30),
    `occupation` varchar(10),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

CREATE TABLE `relation` (
    `id` bigint(11) NOT NULL AUTO_INCREMENT,
    `origin`  varchar(18) NOT NULL,
    `target` varchar(18) NOT NULL,
    `relationship` smallint(3) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;