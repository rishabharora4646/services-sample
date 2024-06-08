SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `services_sample`
--

-- --------------------------------------------------------

--
-- Table structure for table `orgs`
--

DROP TABLE IF EXISTS `orgs`;
CREATE TABLE IF NOT EXISTS `orgs` (
  `org_id` varchar(30) NOT NULL,
  `org_name` tinytext NOT NULL,
  `created_at` varchar(30) NOT NULL,
  PRIMARY KEY (`org_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `orgs`
--

INSERT INTO `orgs` (`org_id`, `org_name`, `created_at`) VALUES
('cphjjrtaas1c714mre7i', 'Test Org 1', '2024-06-08T16:50:47Z');

-- --------------------------------------------------------

--
-- Table structure for table `org_users`
--

DROP TABLE IF EXISTS `org_users`;
CREATE TABLE IF NOT EXISTS `org_users` (
  `user_id` varchar(30) NOT NULL,
  `org_id` varchar(30) NOT NULL,
  KEY `user_id` (`user_id`),
  KEY `org_id` (`org_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `org_users`
--

INSERT INTO `org_users` (`user_id`, `org_id`) VALUES
('cphjjrtaas1c714mre7h', 'cphjjrtaas1c714mre7i');

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
CREATE TABLE IF NOT EXISTS `services` (
  `service_id` varchar(30) NOT NULL,
  `service_name` tinytext NOT NULL,
  `service_description` text NOT NULL,
  `user_id` varchar(30) NOT NULL,
  `org_id` varchar(30) NOT NULL,
  `created_at` varchar(30) NOT NULL,
  `updated_at` varchar(30) NOT NULL,
  PRIMARY KEY (`service_id`),
  KEY `user_id` (`user_id`),
  KEY `org_id` (`org_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `services`
--

INSERT INTO `services` (`service_id`, `service_name`, `service_description`, `user_id`, `org_id`, `created_at`, `updated_at`) VALUES
('cphjjrtaas1c714mre7j', 'Locate Us', 'This is Locate Us service', 'cphjjrtaas1c714mre7h', 'cphjjrtaas1c714mre7i', '2024-06-08T16:51:47Z', '2024-06-08T16:51:47Z'),
('cphjjrtaas1c714mre7k', 'Contact Us', 'Contact Us service description.', 'cphjjrtaas1c714mre7h', 'cphjjrtaas1c714mre7i', '2024-06-08T16:52:47Z', '2024-06-08T16:52:47Z');

-- --------------------------------------------------------

--
-- Table structure for table `services_versions`
--

DROP TABLE IF EXISTS `services_versions`;
CREATE TABLE IF NOT EXISTS `services_versions` (
  `version_id` varchar(30) NOT NULL,
  `service_id` varchar(30) NOT NULL,
  `version_name` tinytext NOT NULL,
  `service_host` tinytext NOT NULL,
  `service_port` int(11) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `created_at` varchar(30) NOT NULL,
  `updated_at` varchar(30) NOT NULL,
  PRIMARY KEY (`version_id`),
  KEY `service_id` (`service_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `services_versions`
--

INSERT INTO `services_versions` (`version_id`, `service_id`, `version_name`, `service_host`, `service_port`, `is_active`, `created_at`, `updated_at`) VALUES
('cphjjrtaas1c714mre7l', 'cphjjrtaas1c714mre7j', 'version 1', 'locate-v1.abc.com', 8080, 0, '2024-06-08T16:53:47Z', '2024-06-08T16:53:47Z'),
('cphjjrtaas1c714mre7m', 'cphjjrtaas1c714mre7j', 'version 2', 'locate-v2.abc.com', 8081, 1, '2024-06-08T16:54:47Z', '2024-06-08T16:54:47Z'),
('cphjjrtaas1c714mre7n', 'cphjjrtaas1c714mre7k', 'contact version 1', 'contact-v1.abc.com', 1234, 1, '2024-06-08T16:55:47Z', '2024-06-08T16:55:47Z');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` varchar(30) NOT NULL,
  `email` varchar(320) NOT NULL,
  `password` varchar(80) NOT NULL,
  `created_at` varchar(30) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`user_id`, `email`, `password`, `created_at`) VALUES
('cphjjrtaas1c714mre7h', 'test@test.com', '$2a$10$dtQ7juw89Vb0i4e9H3h8n.CNYDAZkSqLccjrbsmWxVTMs.hUCWOj6', '2024-06-08T16:48:47Z');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
