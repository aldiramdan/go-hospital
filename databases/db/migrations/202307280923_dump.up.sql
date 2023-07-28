-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: db_hospital
-- ------------------------------------------------------
-- Server version	8.0.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `disease`
--

DROP TABLE IF EXISTS `disease`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `disease` (
  `id` char(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `medicine_id` char(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `medicine_fk` (`medicine_id`),
  CONSTRAINT `medicine_fk` FOREIGN KEY (`medicine_id`) REFERENCES `medicine` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `disease`
--

LOCK TABLES `disease` WRITE;
/*!40000 ALTER TABLE `disease` DISABLE KEYS */;
INSERT INTO `disease` VALUES ('0c2417cc-4aa1-4229-88d2-898da41a8630','Sore Throat','150799a6-464b-4079-878f-ded155eb12d4','2023-07-28 03:29:35','2023-07-28 03:29:35'),('183bab79-c225-46e6-972f-62ca00603f65','Allergies','6b1d55a9-5216-48c1-972a-d0a1ccac7cfa','2023-07-28 03:30:09','2023-07-28 03:30:09'),('50b49fe8-4560-4982-9616-c5277fc91442','Cold','d58b9fd1-a819-42b7-a67b-b9538bf616dd','2023-07-28 03:30:33','2023-07-28 03:30:33'),('c4303521-032a-4f60-a047-f2774deea8eb','Arthritis','8ebc4436-f564-45a9-b27b-00a2114742be','2023-07-28 03:30:19','2023-07-28 03:30:19'),('e55232d2-3542-4183-8689-c566110a17c4','Fever','c974a096-5f9e-463f-973c-140ce67bfdc3','2023-07-28 03:29:59','2023-07-28 03:29:59');
/*!40000 ALTER TABLE `disease` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `doctor`
--

DROP TABLE IF EXISTS `doctor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `doctor` (
  `id` char(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `specialization` varchar(100) NOT NULL,
  `consultation_fee` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `doctor`
--

LOCK TABLES `doctor` WRITE;
/*!40000 ALTER TABLE `doctor` DISABLE KEYS */;
INSERT INTO `doctor` VALUES ('1e96cdba-5eb8-4df3-b2b4-54ffb3b165ac','Dr. Michael Johnson','Dermatologist',71922,'2023-07-28 03:20:02','2023-07-28 03:20:02'),('a6c398d4-9f8c-4c50-909a-b91050f11976','Dr. Emily Brown','Orthopedic Surgeon',74179,'2023-07-28 03:20:07','2023-07-28 03:20:07'),('c49491e3-06e4-439b-8191-e3d80e31ec00','Dr. John Doe','Pediatrician',75549,'2023-07-28 03:20:05','2023-07-28 03:20:05'),('f6632238-6617-4a7f-9154-c0904f760ecd','Dr. Jane Smith','General Physician',83672,'2023-07-28 03:20:04','2023-07-28 03:20:04');
/*!40000 ALTER TABLE `doctor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medicine`
--

DROP TABLE IF EXISTS `medicine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `medicine` (
  `id` char(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `price` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medicine`
--

LOCK TABLES `medicine` WRITE;
/*!40000 ALTER TABLE `medicine` DISABLE KEYS */;
INSERT INTO `medicine` VALUES ('150799a6-464b-4079-878f-ded155eb12d4','Metoprolol',3598,'2023-07-28 03:22:50','2023-07-28 03:22:50'),('6a8b5759-2066-4c39-9044-c34bc9140c4f','Ibuprofen',3892,'2023-07-28 03:22:49','2023-07-28 03:22:49'),('6b1d55a9-5216-48c1-972a-d0a1ccac7cfa','Amoxicillin',3339,'2023-07-28 03:22:51','2023-07-28 03:22:51'),('8ebc4436-f564-45a9-b27b-00a2114742be','Omeprazole',4180,'2023-07-28 03:22:47','2023-07-28 03:22:47'),('9adf0c70-6773-4fab-8e41-523abe1eb5cb','Lisinopril',4209,'2023-07-28 03:22:46','2023-07-28 03:22:46'),('c974a096-5f9e-463f-973c-140ce67bfdc3','Paracetamol',3199,'2023-07-28 03:22:52','2023-07-28 03:22:52'),('d58b9fd1-a819-42b7-a67b-b9538bf616dd','Aspirin',4243,'2023-07-28 03:22:45','2023-07-28 03:22:45');
/*!40000 ALTER TABLE `medicine` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `patient`
--

DROP TABLE IF EXISTS `patient`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `patient` (
  `id` char(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `age` int NOT NULL,
  `address` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `patient`
--

LOCK TABLES `patient` WRITE;
/*!40000 ALTER TABLE `patient` DISABLE KEYS */;
INSERT INTO `patient` VALUES ('6b3bf6cb-e060-4062-a9ed-f39eb22761f7','Emily',26,'987 Pine Rd','2023-07-28 03:21:31','2023-07-28 03:21:31'),('cca7927e-1bb5-403a-ac6f-6a1f4a767f14','James',21,'123 Main St','2023-07-28 03:21:33','2023-07-28 03:21:33'),('d493a289-bb6e-4cc7-acf3-5a1f0390455b','Emma',37,'789 Oak Ave','2023-07-28 03:21:34','2023-07-28 03:21:34'),('da602506-69c0-4101-9191-7644eda0a296','John',36,'321 Maple Ave','2023-07-28 03:21:32','2023-07-28 03:21:32');
/*!40000 ALTER TABLE `patient` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `treatment`
--

DROP TABLE IF EXISTS `treatment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `treatment` (
  `id` char(100) NOT NULL,
  `patient_id` char(100) NOT NULL,
  `doctor_id` char(100) NOT NULL,
  `service_fee` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `patient_fk` (`patient_id`),
  KEY `doctor_fk` (`doctor_id`),
  CONSTRAINT `doctor_fk` FOREIGN KEY (`doctor_id`) REFERENCES `doctor` (`id`),
  CONSTRAINT `patient_fk` FOREIGN KEY (`patient_id`) REFERENCES `patient` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `treatment`
--

LOCK TABLES `treatment` WRITE;
/*!40000 ALTER TABLE `treatment` DISABLE KEYS */;
/*!40000 ALTER TABLE `treatment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `treatment_disease`
--

DROP TABLE IF EXISTS `treatment_disease`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `treatment_disease` (
  `treatment_id` char(100) NOT NULL,
  `disease_id` char(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `treatment_fk` (`treatment_id`),
  KEY `disease_fk` (`disease_id`),
  CONSTRAINT `disease_fk` FOREIGN KEY (`disease_id`) REFERENCES `disease` (`id`),
  CONSTRAINT `treatment_fk` FOREIGN KEY (`treatment_id`) REFERENCES `treatment` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `treatment_disease`
--

LOCK TABLES `treatment_disease` WRITE;
/*!40000 ALTER TABLE `treatment_disease` DISABLE KEYS */;
/*!40000 ALTER TABLE `treatment_disease` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'db_hospital'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-28 10:31:15
