-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 10, 2021 at 10:06 AM
-- Server version: 10.4.20-MariaDB
-- PHP Version: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `warehouse`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `product_id` int(11) NOT NULL,
  `product_name` varchar(191) DEFAULT NULL,
  `product_price` int(11) NOT NULL DEFAULT 0,
  `product_description` text DEFAULT NULL,
  `product_quantity` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`product_id`, `product_name`, `product_price`, `product_description`, `product_quantity`) VALUES
(1, 'Samsung Galaxy Note20 Ultra 5G 12/512GB - Mystic Bronze Free E Voucher + Samsung Care+*', 1899000, 'Ukuran layar: 6.9 inci, 1440 x 3088 pixels, Dynamic AMOLED 2X capacitive touchscreen, 16M colors\r\nMemori: RAM 12 GB, ROM 512 GB, MicroSD Slot\r\nSistem operasi: Android 11\r\nCPU: Exynos 990 (7 nm+) Octa-core\r\nGPU: Mali-G77 MP11\r\nKamera: Triple 108 MP, f/1.8; 12 MP, f/3.0; 12 MP, f/2.0, depan 10 MP, f/2.2\r\nSIM: Hybrid Dual SIM (Nano-SIM)\r\nBaterai: Non-removable Li-Ion 4500 mAh\r\nBerat: 208 gram\r\nGaransi Resmi', 100),
(2, 'Huawei MatePad 11 6/128GB - Matte Grey Free Keyboard + Mouse', 7299000, 'Ukuran layar: 10.95 inci, 2560 x 1600 pixels, IPS LCD, 120Hz\r\nMemori: RAM 6 GB, ROM 128 GB, MicroSD up to 1TB\r\nSistem operasi: HarmonyOS 2\r\nCPU: Qualcomm Snapdragon 865 Octa-core (1x2.84 GHz Kryo 585 & 3x2.42 GHz Kryo 585 & 4x1.8 GHz Kryo 585)\r\nGPU: Adreno 650\r\nKamera: 13 MP, f/1.8, PDAF, depan 8 MP, f/2.0\r\nWiFi Only\r\nBaterai: Li-Po 7250 mAh, non-removable\r\nBerat: 485 gram\r\nGaransi Resmi.', 100),
(4, 'Xiaomi Black Shark 4 8/128GB', 8900000, 'Ukuran layar: 6.67 inci, 1080 x 2400 pixels, Super AMOLED, 144Hz', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
