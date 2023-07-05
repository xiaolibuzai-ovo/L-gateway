-- version 4.7.0
--
-- Host: 127.0.0.1
-- Generation Time: 2020-04-27 13:59:47
-- 服务器版本： 5.6.25-log
-- PHP Version: 7.0.18

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `L_gateway`
--

-- --------------------------------------------------------

--
-- 表的结构 `gateway_admin`
--

CREATE TABLE `gateway_admin` (
  `id` bigint(20) NOT NULL COMMENT '自增id',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员表';

--
-- 表的结构 `gateway_app`
--

CREATE TABLE `gateway_app` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '自增id',
  `tenant_id` varchar(255) NOT NULL DEFAULT '' COMMENT '租户id',
  `tenant_type` varchar(255) NOT NULL DEFAULT '' COMMENT '租户类型',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '租户名称',
  `secret` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥',
  `white_ips` varchar(1000) NOT NULL DEFAULT '' COMMENT 'ip白名单，支持前缀匹配',
  `qpd` bigint(20) NOT NULL DEFAULT '0' COMMENT '日请求量限制',
  `qps` bigint(20) NOT NULL DEFAULT '0' COMMENT '每秒请求量限制',
  `create_at` datetime NOT NULL COMMENT '添加时间',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 1=删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关租户表';

--
-- 表的结构 `gateway_service_access_control`
--

CREATE TABLE `gateway_service_access_control` (
  `id` bigint(20) NOT NULL COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务id',
  `open_auth` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否开启权限 1=开启',
  `black_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '黑名单ip',
  `white_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单ip',
  `white_host_name` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单主机',
  `clientip_flow_limit` int(11) NOT NULL DEFAULT '0' COMMENT '客户端ip限流',
  `service_flow_limit` int(20) NOT NULL DEFAULT '0' COMMENT '服务端限流'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关权限控制表';

--
-- 转存表中的数据 `gateway_service_access_control`
--

--
-- 表的结构 `gateway_service_grpc_rule`
--

CREATE TABLE `gateway_service_grpc_rule` (
  `id` bigint(20) NOT NULL COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务id',
  `port` int(5) NOT NULL DEFAULT '0' COMMENT '端口',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关路由匹配表';

--
-- 表的结构 `gateway_service_http_rule`
--

CREATE TABLE `gateway_service_http_rule` (
  `id` bigint(20) NOT NULL COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL COMMENT '服务id',
  `rule_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '匹配类型 0=url前缀url_prefix 1=域名domain ',
  `rule` varchar(255) NOT NULL DEFAULT '' COMMENT 'type=domain表示域名，type=url_prefix时表示url前缀',
  `need_https` tinyint(4) NOT NULL DEFAULT '0' COMMENT '支持https 1=支持',
  `need_strip_uri` tinyint(4) NOT NULL DEFAULT '0' COMMENT '启用strip_uri 1=启用',
  `need_websocket` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否支持websocket 1=支持',
  `url_rewrite` varchar(5000) NOT NULL DEFAULT '' COMMENT 'url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关路由匹配表';

--
-- 表的结构 `gateway_service_info`
--

CREATE TABLE `gateway_service_info` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '自增主键',
  `load_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '负载类型 0=http 1=tcp 2=grpc',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称 6-128 数字字母下划线',
  `service_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '服务描述',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '添加时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint(4) DEFAULT '0' COMMENT '是否删除 1=删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关基本信息表';

--
-- 表的结构 `gateway_service_load_balance`
--

CREATE TABLE `gateway_service_load_balance` (
  `id` bigint(20) NOT NULL COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务id',
  `check_method` tinyint(20) NOT NULL DEFAULT '0' COMMENT '检查方法 0=tcpchk,检测端口是否握手成功',
  `check_timeout` int(10) NOT NULL DEFAULT '0' COMMENT 'check超时时间,单位s',
  `check_interval` int(11) NOT NULL DEFAULT '0' COMMENT '检查间隔, 单位s',
  `round_type` tinyint(4) NOT NULL DEFAULT '2' COMMENT '轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash',
  `ip_list` varchar(2000) NOT NULL DEFAULT '' COMMENT 'ip列表',
  `weight_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '权重列表',
  `forbid_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '禁用ip列表',
  `upstream_connect_timeout` int(11) NOT NULL DEFAULT '0' COMMENT '建立连接超时, 单位s',
  `upstream_header_timeout` int(11) NOT NULL DEFAULT '0' COMMENT '获取header超时, 单位s',
  `upstream_idle_timeout` int(10) NOT NULL DEFAULT '0' COMMENT '链接最大空闲时间, 单位s',
  `upstream_max_idle` int(11) NOT NULL DEFAULT '0' COMMENT '最大空闲链接数'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关负载表';


--
-- 表的结构 `gateway_service_tcp_rule`
--

CREATE TABLE `gateway_service_tcp_rule` (
  `id` bigint(20) NOT NULL COMMENT '自增主键',
  `service_id` bigint(20) NOT NULL COMMENT '服务id',
  `port` int(5) NOT NULL DEFAULT '0' COMMENT '端口号'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网关路由匹配表';